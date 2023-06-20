#!/usr/bin/env python3
#
# Copyright (C) 2022 IBM Corporation.
#
# Authors:
# Frederico Araujo <frederico.araujo@ibm.com>
# Teryl Taylor <terylt@ibm.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
import hashlib, json, os, sys
import urllib.request
from functools import reduce
from collections import OrderedDict
import sysflow.utils as utils
import sysflow.opflags as opflags
from sysflow.formatter import _fields, SFFormatter
from sysflow.objtypes import ObjectTypes, OBJECT_MAP
from sysflow.reader import FlattenedSFReader
from sysflow.sfql import SfqlInterpreter
from graphviz import Digraph
import matplotlib.pylab as plt
import matplotlib.dates as mdates
import numpy as np
import pandas as pd
import mitreattack.attackToExcel.attackToExcel as attackToExcel
import mitreattack.attackToExcel.stixToDf as stixToDf

# To silence info logging and progress bar from mitreattack module.
from loguru import logger

logger.remove()
logger.add(sys.stderr, level='ERROR')
from tqdm import tqdm
from functools import partialmethod

tqdm.__init__ = partialmethod(tqdm.__init__, disable=True)

"""
.. module:: sysflow.graphlet
   :synopsis: This module re-interprets SysFlow traces in a compact provenance graph representation
.. moduleauthor:: Frederico Araujo, Teryl Taylor
"""

INFSYMB = '&infin;'

FLOW_FIELDS = [
    'ts_uts',
    'endts_uts',
    'type',
    'opflags',
    'proc.pid',
    'proc.tid',
    'pproc.pid',
    'proc.exe',
    'proc.args',
    'pproc.exe',
    'pproc.args',
    'res',
    'flow.rbytes',
    'flow.rops',
    'flow.wbytes',
    'flow.wops',
    'container.id',
    'tags',
]
EVT_FIELDS = [
    'ts_uts',
    'type',
    'opflags',
    'proc.pid',
    'proc.tid',
    'pproc.pid',
    'proc.exe',
    'proc.args',
    'pproc.exe',
    'tags',
]


class Graphlet(object):
    """
    **Graphlet**

    This class takes a path pointing to a sysflow trace or a directory containing sysflow traces.

    Example Usage::

         # basic usage
         g1 = Graphlet('data/')
         g1.view()

         # filtering and enrichment with policies
         ioc1 = 'proc.exe = /usr/bin/scp'
         g1 = Graphlet('data/', ioc1, ['policies/ttps.yaml'])
         g1.view()

    :param graphlet: A compact provenance graph representation based on sysflow traces.
    :type graphlet: sysflow.Graphlet
    """

    attackdata = None
    techniques_normalized = pd.DataFrame()
    mitigations_normalized = pd.DataFrame()
    associated_mitigations_normalized = pd.DataFrame()
    defend_data = {}

    def __init__(self, path, expr=None, defs=[]):
        """Create graphlet object from raw sysflow with optional filters and policy taggers.

        :param path: a path to a sysflow trace or directory containing sysflow traces.
        :type path: str

        :param expr: sfql style filter.
        :type expr: str

        :param defs: a list of paths for yaml policies that enrich graph nodes.
        :type defs: list
        """
        if os.path.isfile(path):
            self.readers = [FlattenedSFReader(path, retEntities=True)]
        elif os.path.isdir(path):
            self.readers = [FlattenedSFReader(f, retEntities=True) for f in _files(path)]
        self.nodes = OrderedDict()
        self.edges = set()
        self.sfqlint = SfqlInterpreter(paths=defs)
        self.fmt = SFFormatter(None)
        for reader in self.readers:
            self.reader = reader
            self.__create(expr)

    def __create(self, expr=None):
        for objtype, header, pod, cont, pproc, proc, files, evt, flow in self.sfqlint.filter(self.reader, expr):

            tags = self.sfqlint.enrich((objtype, header, pod, cont, pproc, proc, files, evt, flow))

            if objtype == ObjectTypes.PROC_EVT:
                if proc.oid.hpid != evt.tid or not pproc:
                    continue
                r = self.fmt._flatten(objtype, header, pod, cont, pproc, proc, files, evt, flow, None, tags=tags)
                opflag = utils.getOpFlagsStr(evt.opFlags)

                filt = lambda v: (v.exe, v.args, v.uid, v.user, v.gid, v.group, v.tty) == (
                    proc.exe,
                    proc.exeArgs,
                    proc.uid,
                    proc.userName,
                    proc.gid,
                    proc.groupName,
                    proc.tty,
                ) and v.hasProc(pproc.oid.hpid, pproc.oid.createTS)
                if opflag == utils.getOpFlagsStr(opflags.OP_CLONE) and (
                    proc.exe,
                    proc.exeArgs,
                    proc.uid,
                    proc.userName,
                    proc.gid,
                    proc.groupName,
                ) == (
                    pproc.exe,
                    pproc.exeArgs,
                    pproc.uid,
                    pproc.userName,
                    pproc.gid,
                    pproc.groupName,
                ):
                    self.__addProcEvtEdge(opflag, proc, pproc, r, filt)

                filt = (
                    lambda v: (v.exe, v.args, v.tty)
                    != (
                        proc.exe,
                        proc.exeArgs,
                        proc.tty,
                    )
                    and (v.uid, v.user, v.gid, v.group)
                    == (
                        proc.uid,
                        proc.userName,
                        proc.gid,
                        proc.groupName,
                    )
                    and v.hasProc(proc.oid.hpid, proc.oid.createTS)
                )
                if opflag == utils.getOpFlagsStr(opflags.OP_EXEC):
                    self.__addProcEvtEdge(opflag, proc, pproc, r, filt)

                filt = (
                    lambda v: (v.exe, v.args, v.tty)
                    == (
                        proc.exe,
                        proc.exeArgs,
                        proc.tty,
                    )
                    and (v.uid, v.user, v.gid, v.group)
                    != (
                        proc.uid,
                        proc.userName,
                        proc.gid,
                        proc.groupName,
                    )
                    and v.hasProc(proc.oid.hpid, proc.oid.createTS)
                )
                if opflag == utils.getOpFlagsStr(opflags.OP_SETUID):
                    self.__addProcEvtEdge(opflag, proc, pproc, r, filt)

                filt = lambda v: (v.exe, v.args, v.uid, v.user, v.gid, v.group, v.tty) == (
                    proc.exe,
                    proc.exeArgs,
                    proc.uid,
                    proc.userName,
                    proc.gid,
                    proc.groupName,
                    proc.tty,
                ) and v.hasProc(proc.oid.hpid, proc.oid.createTS)
                if opflag == utils.getOpFlagsStr(opflags.OP_EXIT):
                    self.__addProcEvtEdge(opflag, proc, pproc, r, filt)

            if objtype == ObjectTypes.FILE_FLOW and files[0].path != proc.exe:
                filt = lambda v: (v.exe, v.args, v.uid, v.user, v.gid, v.group, v.tty) == (
                    proc.exe,
                    proc.exeArgs,
                    proc.uid,
                    proc.userName,
                    proc.gid,
                    proc.groupName,
                    proc.tty,
                ) and v.hasProc(proc.oid.hpid, proc.oid.createTS)
                r = self.fmt._flatten(objtype, header, pod, cont, pproc, proc, files, evt, flow, None, tags=tags)
                self.__addFileFlowEdge(proc, pproc, r, filt)

            if objtype == ObjectTypes.NET_FLOW:
                filt = lambda v: (v.exe, v.args, v.uid, v.user, v.gid, v.group, v.tty) == (
                    proc.exe,
                    proc.exeArgs,
                    proc.uid,
                    proc.userName,
                    proc.gid,
                    proc.groupName,
                    proc.tty,
                ) and v.hasProc(proc.oid.hpid, proc.oid.createTS)
                r = self.fmt._flatten(objtype, header, pod, cont, pproc, proc, files, evt, flow, None, tags=tags)
                self.__addNetFlowEdge(proc, pproc, r, filt)

    def __addProcEvtEdge(self, opflag, proc, pproc, r, filt):
        n1_k = _hash(
            (
                proc.exe,
                proc.exeArgs,
                proc.uid,
                proc.userName,
                proc.gid,
                proc.groupName,
                proc.tty,
                pproc.exe,
                pproc.exeArgs,
            )
        )
        if n1_k in self.nodes:
            n1_v = self.nodes[n1_k]
        else:
            n1_v = ProcessNode(
                n1_k, proc.exe, proc.exeArgs, proc.uid, proc.userName, proc.gid, proc.groupName, proc.tty
            )
        n1_v.addProc(proc.oid.hpid, proc.oid.createTS, r)
        self.nodes[n1_k] = n1_v

        if opflag == utils.getOpFlagsStr(opflags.OP_EXIT):
            p = proc
            n2_k, n2_v = self.__findNode(filt)
        else:  # OP_CLONE, OP_EXEC, OP_SETUID
            p = pproc
            n2_k, n2_v = self.__findNode(filt)

        if not n2_k:
            key = self.reader.getProcessKey(p.poid) if p.poid else None
            pp = None
            if key in self.reader.processes:
                pp = self.reader.processes[key]
                n2_k = _hash((p.exe, p.exeArgs, p.uid, p.userName, p.gid, p.groupName, p.tty, pp.exe, pp.exeArgs))
            else:
                n2_k = _hash((p.exe, p.exeArgs, p.uid, p.userName, p.gid, p.groupName, p.tty))

            if n2_k in self.nodes:
                n2_v = self.nodes[n2_k]
            else:
                n2_v = ProcessNode(n2_k, p.exe, p.exeArgs, p.uid, p.userName, p.gid, p.groupName, p.tty)
            n2_v.addProc(
                p.oid.hpid,
                p.oid.createTS,
                self.fmt._flatten(ObjectTypes.PROC, None, None, None, pp, p, None, None, None, None, None),
            )
            self.nodes[n2_k] = n2_v
        self.edges.add(EvtEdge(n2_k, n1_k, opflag))

    def __addFileFlowEdge(self, proc, pproc, r, filt):
        if pproc:
            n1_k = _hash(
                (
                    proc.exe,
                    proc.exeArgs,
                    proc.uid,
                    proc.userName,
                    proc.gid,
                    proc.groupName,
                    proc.tty,
                    OBJECT_MAP[ObjectTypes.FILE_FLOW],
                    pproc.exe,
                    pproc.exeArgs,
                )
            )
        else:
            n1_k = _hash(
                (
                    proc.exe,
                    proc.exeArgs,
                    proc.uid,
                    proc.userName,
                    proc.gid,
                    proc.groupName,
                    proc.tty,
                    OBJECT_MAP[ObjectTypes.FILE_FLOW],
                )
            )
        new = False
        if n1_k in self.nodes:
            n1_v = self.nodes[n1_k]
        else:
            n1_v = FileFlowNode(n1_k, proc.exe, proc.exeArgs)
            new = True
        n1_v.addFlow(r)
        self.nodes[n1_k] = n1_v
        if new:
            n2_k, n2_v = self.__findNode(filt)
            if not n2_k:
                key = self.reader.getProcessKey(proc.poid) if proc.poid else None
                pp = None
                if key in self.reader.processes:
                    pp = self.reader.processes[key]
                    n2_k = _hash(
                        (
                            proc.exe,
                            proc.exeArgs,
                            proc.uid,
                            proc.userName,
                            proc.gid,
                            proc.groupName,
                            proc.tty,
                            pp.exe,
                            pp.exeArgs,
                        )
                    )
                else:
                    n2_k = _hash((proc.exe, proc.exeArgs, proc.uid, proc.userName, proc.gid, proc.groupName, proc.tty))

                if n2_k in self.nodes:
                    n2_v = self.nodes[n2_k]
                else:
                    n2_v = ProcessNode(
                        n2_k, proc.exe, proc.exeArgs, proc.uid, proc.userName, proc.gid, proc.groupName, proc.tty
                    )
                n2_v.addProc(
                    proc.oid.hpid,
                    proc.oid.createTS,
                    self.fmt._flatten(ObjectTypes.PROC, None, None, None, pp, proc, None, None, None, None, None),
                )
                self.nodes[n2_k] = n2_v
            self.edges.add(FlowEdge(n2_k, n1_k, OBJECT_MAP[ObjectTypes.FILE_FLOW]))

    def __addNetFlowEdge(self, proc, pproc, r, filt):
        if pproc:
            n1_k = _hash(
                (
                    proc.exe,
                    proc.exeArgs,
                    proc.uid,
                    proc.userName,
                    proc.gid,
                    proc.groupName,
                    proc.tty,
                    OBJECT_MAP[ObjectTypes.NET_FLOW],
                    pproc.exe,
                    pproc.exeArgs,
                )
            )
        else:
            n1_k = _hash(
                (
                    proc.exe,
                    proc.exeArgs,
                    proc.uid,
                    proc.userName,
                    proc.gid,
                    proc.groupName,
                    proc.tty,
                    OBJECT_MAP[ObjectTypes.NET_FLOW],
                )
            )
        new = False
        if n1_k in self.nodes:
            n1_v = self.nodes[n1_k]
        else:
            n1_v = NetFlowNode(n1_k, proc.exe, proc.exeArgs)
            new = True
        n1_v.addFlow(r)
        self.nodes[n1_k] = n1_v
        if new:
            n2_k, n2_v = self.__findNode(filt)
            if not n2_k:
                key = self.reader.getProcessKey(proc.poid) if proc.poid else None
                pp = None
                if key in self.reader.processes:
                    pp = self.reader.processes[key]
                    n2_k = _hash(
                        (
                            proc.exe,
                            proc.exeArgs,
                            proc.uid,
                            proc.userName,
                            proc.gid,
                            proc.groupName,
                            proc.tty,
                            pp.exe,
                            pp.exeArgs,
                        )
                    )
                else:
                    n2_k = _hash((proc.exe, proc.exeArgs, proc.uid, proc.userName, proc.gid, proc.groupName, proc.tty))
                if n2_k in self.nodes:
                    n2_v = self.nodes[n2_k]
                else:
                    n2_v = ProcessNode(
                        n2_k, proc.exe, proc.exeArgs, proc.uid, proc.userName, proc.gid, proc.groupName, proc.tty
                    )
                n2_v.addProc(
                    proc.oid.hpid,
                    proc.oid.createTS,
                    self.fmt._flatten(ObjectTypes.PROC, None, None, None, pp, proc, None, None, None, None, None),
                )
                self.nodes[n2_k] = n2_v
            self.edges.add(FlowEdge(n2_k, n1_k, OBJECT_MAP[ObjectTypes.NET_FLOW]))

    def __findNode(self, filt):
        for k, v in reversed(self.nodes.items()):
            # print('Find Node ({0}, {1}): {2} {3}'.format(v.exe, v.args, isinstance(v, ProcessNode), filt(v)))
            if isinstance(v, ProcessNode) and filt(v):
                return (k, v)
        return (None, None)

    def __loadAttackTechniquesOnce(self):
        if not self.techniques_normalized.empty:
            return
        if not self.attackdata:
            self.attackdata = attackToExcel.get_stix_data('enterprise-attack')
        techniques_data = stixToDf.techniquesToDf(self.attackdata, 'enterprise-attack')
        self.techniques_normalized = techniques_data['techniques']

    def __loadAttackMitigationsOnce(self):
        if not self.mitigations_normalized.empty:
            return
        if not self.attackdata:
            self.attackdata = attackToExcel.get_stix_data('enterprise-attack')
        mitigations_data = stixToDf.mitigationsToDf(self.attackdata)
        self.mitigations_normalized = mitigations_data['mitigations']

    def __loadAssociatedAttackMitigationsOnce(self):
        if not self.associated_mitigations_normalized.empty:
            return
        if not self.attackdata:
            self.attackdata = attackToExcel.get_stix_data('enterprise-attack')
        stixToDf.relationshipsToDf(self.attackdata, 'technique')
        associated_mitigations_data = stixToDf.relationshipsToDf(self.attackdata, 'technique')
        self.associated_mitigations_normalized = associated_mitigations_data['associated mitigations']

    def __loadDefendData(self, techniqueID):
        if techniqueID in self.defend_data.keys():
            return
        d3f_url = 'https://d3fend.mitre.org/api/offensive-technique/attack/{0}.json'.format(techniqueID)
        d3f_vars = [
            'def_tactic_label.value',
            'def_tech_parent_label.value',
            'def_tech_label.value',
            'def_artifact_rel_label.value',
            'def_artifact_label.value',
            'off_tech_id.value',
            'off_artifact_label.value',
            'off_artifact_rel_label.value',
            'off_tech_label.value',
            'off_tactic_rel_label.value',
            'off_tactic_label.value',
        ]
        try:
            with urllib.request.urlopen(d3f_url) as url:
                data = json.loads(url.read().decode())
            # normalize data
            df = pd.json_normalize(data["off_to_def"]["results"]["bindings"])
            if df.empty:
                return
            # fill in missing bindings
            for c in d3f_vars:
                if c not in df:
                    df[c] = None
            # rename columns
            defend_data = df[d3f_vars].rename(
                columns={
                    'def_tactic_label.value': 'def_tactic',
                    'def_tech_label.value': 'def_tech',
                    'def_tech_parent_label.value': 'def_tech_parent',
                    'def_artifact_rel_label.value': 'def_artifact_rel',
                    'def_artifact_label.value': 'def_artifact',
                    'off_tech_id.value': 'technique_id',
                    'off_artifact_label.value': 'artifact',
                    'off_artifact_rel_label.value': 'artifact_rel',
                    'off_tech_label.value': 'technique',
                    'off_tactic_rel_label.value': 'tactic_rel',
                    'off_tactic_label.value': 'tactic',
                }
            )
            self.defend_data[techniqueID] = defend_data.dropna(subset=['def_tactic', 'def_tech'])
        except Exception:
            return

    def df(self, oid=None):
        """Returns a dataframe containing a summary of the graph node IDs and process metadata associated with them.

        :param oid: a node ID filter.
        :type oid: object ID string
        """
        data = OrderedDict()
        for idx, r in enumerate(self.nodes.items()):
            if not oid or oid == r[1].oid:
                data[idx] = r
        return pd.DataFrame.from_dict(data, orient='index', columns=['id', 'name'])

    def data(self, oid=None):
        """Returns a dataframe containing the underlying data (sysflow records) of the graph.

        :param oid: a node ID filter.
        :type oid: object ID string
        """
        df = pd.DataFrame()
        for k, r in self.nodes.items():
            if not oid or oid == r.oid:
                df = pd.concat([df, r.df()])
        df.reindex()
        df.sort_values(by=['ts_uts'], inplace=True, ignore_index=True)
        return df

    def tags(self, oid=None):
        """Returns a dataframe containing the set of (enrichment) tags in the graph.

        :param oid: a node ID filter.
        :type oid: object ID string
        """
        df = pd.DataFrame()
        for k, r in self.nodes.items():
            if not oid or oid == r.oid:
                tdf = r.df()[['tags']]
                df = pd.concat([df, tdf[tdf['tags'].map(lambda d: len(d)) > 0].reset_index(drop=True)])
                # df['tmp'] = df.apply(lambda row: str(row.tags[1]), axis=1)
        # df = df.drop_duplicates(subset='tmp')[['tags']]
        return df

    def ttps(self, oid=None, details=False):
        """Returns a dataframe containing the set of MITRE TTP tags in the graph (e.g., as enriched by the ttps.yaml policy provided with the SysFlow processor).

        :param oid: a node ID filter.
        :type oid: object ID string

        :param details: indicates whether to include complete TTP metadata in the dataframe.
        :type details: boolean
        """
        df = self.tags(oid)
        ttps = pd.DataFrame()
        for e in df['tags']:
            for s in e[1]:
                t = s.split(':')
                if t[0] == 'mitre':
                    self.__loadAttackTechniquesOnce()
                    tdf = self.techniques_normalized[self.techniques_normalized.ID == t[1]]
                    if not details:
                        tdf = tdf.reindex(['ID', 'name', 'url', 'tactics', 'platforms'], axis=1)
                    ttps = pd.concat([ttps, tdf[tdf.ID == t[1]]])
        return ttps.drop_duplicates(subset='ID').reset_index(drop=True)

    def associatedMitigations(self, oid=None):
        """Returns a dataframe containing the set of MITRE mitigations associated with TTPs annotated in the graph.

        :param oid: a node ID filter.
        :type oid: object ID string
        """
        mitigations = pd.DataFrame()
        ttps = self.ttps(oid)
        for e in ttps['ID']:
            self.__loadAssociatedAttackMitigationsOnce()
            mdf = self.associated_mitigations_normalized[self.associated_mitigations_normalized['target ID'] == e]
            mitigations = pd.concat([mitigations, mdf[mdf['target ID'] == e]])
        return mitigations

    def mitigations(self, oid=None, details=False):
        """Returns a dataframe containing the summary set of MITRE mitigations associated with TTPs annotated in the graph.

        :param oid: a node ID filter.
        :type oid: object ID string
        """
        mitigations = pd.DataFrame()
        rdf = self.associatedMitigations(oid)
        for e in rdf['source ID']:
            self.__loadAttackMitigationsOnce()
            mdf = self.mitigations_normalized[self.mitigations_normalized['ID'] == e]
            if not details:
                mdf = mdf.reindex(['ID', 'name', 'url'], axis=1)
            mitigations = pd.concat([mitigations, mdf[mdf['ID'] == e]])
        return mitigations.drop_duplicates(subset='ID').reset_index(drop=True)

    def countermeasures(self, oid=None):
        """Returns a dataframe containing the set of MITRE d3fend defenses associated with TTPs annotated in the graph.

        :param oid: a node ID filter.
        :type oid: object ID string
        """
        countermeasures = pd.DataFrame()
        ttps = self.ttps(oid)
        for e in ttps['ID']:
            self.__loadDefendData(e)
            if e in self.defend_data.keys():
                def_data = self.defend_data[e]
                cmdf = def_data[def_data['technique_id'] == e]
            countermeasures = pd.concat([countermeasures, cmdf[cmdf['technique_id'] == e]])
        return countermeasures

    def view(self, withoid=False, peek=True, peeksize=3, flows=True, ttps=False):
        """Visualizes the graph in dot format.

        :param withoid: indicates whether to show the node ID.
        :type withoid: boolean

        :param peek: indicates whether to show details about the records associated with the nodes.
        :type peek: boolean

        :param peeksize: the number of underlying sysflow records to show in the node.
        :type peeksize: integer

        :param flows: indicates whether to show flow nodes.
        :type flows: boolean

        :param ttps: indicates whether to show tags.
        :type ttps: boolean
        """
        graph_attr = {'splines': 'true', 'overlap': 'scale', 'rankdir': 'TD'}
        node_attr = {'shape': 'Mrecord', 'fontsize': '9'}
        edge_attr = {'fontsize': '8'}
        g = Digraph('graphlet', directory='/tmp/.sf/', node_attr=node_attr, edge_attr=edge_attr, graph_attr=graph_attr)
        for k, v in self.nodes.items():
            if flows and (isinstance(v, FileFlowNode) or isinstance(v, NetFlowNode)) and len(v.df()) > 0:
                if ttps and v.score() > 0:
                    g.node(
                        str(k),
                        v.dot(withoid, peek, peeksize, ttps),
                        style='filled,bold',
                        color='red',
                        fontcolor='red',
                        fillcolor='#ff000010',
                    )
                else:
                    g.node(str(k), v.dot(withoid, peek, peeksize, ttps), style='bold')
            if isinstance(v, ProcessNode):
                if ttps and v.score() > 0:
                    g.node(
                        str(k),
                        v.dot(withoid, peek, peeksize, ttps),
                        style='filled',
                        color='red',
                        fontcolor='red',
                        fillcolor='#ff000010',
                    )
                else:
                    g.node(str(k), v.dot(withoid, peek, peeksize, ttps))
        for e in self.edges:
            t = self.nodes[e.nto()].interval()
            label = '    {0} ({1},{2})'.format(e.op(), t[0], t[1])
            if isinstance(e, EvtEdge):
                g.edge(str(e.nfrom()), str(e.nto()), label=label)
            if flows and isinstance(e, FlowEdge) and len(self.nodes[e.n2].df()) > 0:
                g.edge(str(e.nfrom()), str(e.nto()), label=label, style='dashed')
        return g

    def compare(self, other, withoid=False, peek=True, peeksize=3, flows=True, ttps=False):
        """Compares the graph to another graph (using a simple graph difference), returning a graph slice.

        :param withoid: indicates whether to show the node ID.
        :type withoid: boolean

        :param peek: indicates whether to show details about the records associated with the nodes.
        :type peek: boolean

        :param peeksize: the number of node records to show.
        :type peeksize: integer

        :param flows: indicates whether to show flow nodes.
        :type flows: boolean

        :param ttps: indicates whether to show tags.
        :type ttps: boolean
        """
        lndiff = set(self.nodes) - set(other.nodes)
        lediff = set(self.edges) - set(other.edges)
        graph_attr = {'splines': 'true', 'overlap': 'scale', 'rankdir': 'TD'}
        node_attr = {'shape': 'Mrecord', 'fontsize': '9'}
        edge_attr = {'fontsize': '8'}
        g = Digraph('graphlet', directory='/tmp/.sf/', node_attr=node_attr, edge_attr=edge_attr, graph_attr=graph_attr)
        for k, v in self.nodes.items():
            if flows and (isinstance(v, FileFlowNode) or isinstance(v, NetFlowNode)) and len(v.df()) > 0:
                if k in lndiff:
                    if ttps and v.score() > 0:
                        g.node(
                            str(k),
                            v.dot(withoid, peek, peeksize, ttps),
                            style='filled,bold',
                            color='red',
                            fontcolor='red',
                            fillcolor='#ff000020',
                        )
                    else:
                        g.node(str(k), v.dot(withoid, peek, peeksize, ttps), style='bold', color='red', fontcolor='red')
                else:
                    if ttps and v.score() > 0:
                        g.node(
                            str(k),
                            v.dot(withoid, peek, peeksize, ttps),
                            style='filled,bold',
                            color='red',
                            fontcolor='red',
                            fillcolor='#ff000020',
                        )
                    else:
                        g.node(str(k), v.dot(withoid, peek, peeksize, ttps), style='bold')
            if isinstance(v, ProcessNode):
                if k in lndiff:
                    if ttps and v.score() > 0:
                        g.node(
                            str(k),
                            v.dot(withoid, peek, peeksize, ttps),
                            style='filled,bold',
                            color='red',
                            fontcolor='red',
                            fillcolor='#ff000020',
                        )
                    else:
                        g.node(str(k), v.dot(withoid, peek, peeksize, ttps), style='bold', color='red', fontcolor='red')
                else:
                    if ttps and v.score() > 0:
                        g.node(
                            str(k),
                            v.dot(withoid, peek, peeksize, ttps),
                            style='filled,bold',
                            color='red',
                            fontcolor='red',
                            fillcolor='#ff000020',
                        )
                    else:
                        g.node(str(k), v.dot(withoid, peek, peeksize, ttps), style='bold')
        for e in self.edges:
            t = self.nodes[e.nto()].interval()
            label = '{0} ({1},{2})'.format(e.op(), t[0], t[1])
            if isinstance(e, EvtEdge):
                if e in lediff:
                    g.edge(str(e.nfrom()), str(e.nto()), label=label, color='red', fontcolor='red')
                else:
                    g.edge(str(e.nfrom()), str(e.nto()), label=label)
            if flows and isinstance(e, FlowEdge) and len(self.nodes[e.n2].df()) > 0:
                if e in lediff:
                    g.edge(str(e.nfrom()), str(e.nto()), label=label, style='dashed', color='red', fontcolor='red')
                else:
                    g.edge(str(e.nfrom()), str(e.nto()), label=label, style='dashed')
        return g

    def intersection(self, other, withoid=False, peek=True, peeksize=3, flows=True, ttps=False):
        """Computes the intersection of a graph with another graph, returning the graph corresponding to the intersection of the two graphs.

        :param other: the other graphlet to compute the intersection.
        :type other: Graphlet
        """
        lndiff = set(self.nodes) - set(other.nodes)
        lediff = set(self.edges) - set(other.edges)
        g = self.__clone()
        for n in lndiff:
            del g.nodes[n]
        for e in lediff:
            g.edges.remove(e)
        return g

    def __tag(self, n, label):
        if not label:
            return
        if len(n.data) > 0:
            tags = n.data[0]['tags']
            t0 = tags[0] if tags else []
            t1 = tags[1] if tags else set()
            t2 = tags[2] if tags else 0
            hasTag = any(e for e in [t == label for t in t0])
            if not hasTag:
                n.data[0]['tags'] = (t0 + [label], t1.union(set([label])), max(t2, 1))

    def __prune(self):
        for e in self.edges.copy():
            nto = self.nodes[e.nto()]
            nfrom = self.nodes[e.nfrom()]
            if not nto.tainted or not nfrom.tainted:
                self.edges.remove(e)
        for k, v in self.nodes.copy().items():
            if not v.tainted:
                del self.nodes[k]

    def __bt(self, fringe, label):
        tnodes = set()
        for n in fringe:
            for e in self.edges:
                if e.nto() == n:
                    nfrom = self.nodes[e.nfrom()]
                    self.__tag(nfrom, label)
                    if not nfrom.tainted:
                        nfrom.tainted = True
                        tnodes.add(e.nfrom())
        if len(tnodes) > 0:
            self.__bt(tnodes, label)

    def bt(self, cond, prune=True, label=None):
        """Performs a backward traversal on the graph from nodes matching a condition.

        Example Usage::
            def passwd(df):
                return len(df[(df['file.path'].str.contains('passwd'))])>0
            cond = lambda n: passwd(n.df())
            g.bt(cond, prune=True, label='discovery').view()

        :param cond: a lambda describing a predicate over node properties.
        :type cond: a lambda predicate that received a node object as argument and returns True or False.

        :param prune: if true, nodes outside the dominance paths of matched nodes are pruned.
        :type prune: boolean

        :param label: if set, add a label to nodes in the dominance paths of matched nodes.
        :type label: string
        """
        g = self.__clone()
        for k, v in g.nodes.items():
            v.tainted = False
        for k, v in reversed(g.nodes.items()):
            if cond(v):
                if label:
                    g.__tag(v, label)
                v.tainted = True
                self.__bt({v.oid}, label)
        if prune:
            g.__prune()
        return g

    def __ft(self, fringe, label):
        tnodes = set()
        for n in fringe:
            for e in self.edges:
                if e.nfrom() == n:
                    nto = self.nodes[e.nto()]
                    self.__tag(nto, label)
                    if not nto.tainted:
                        nto.tainted = True
                        tnodes.add(e.nto())
        if len(tnodes) > 0:
            self.__ft(tnodes, label)

    def ft(self, cond, prune=True, label=None):
        """Performs a forward traversal on the graph from nodes matching a condition.

        Example Usage::
            def scp(df):
                return len(df[(df['proc.exe'].str.contains('scp'))])>0
            cond = lambda n: scp(n.df())
            g.ft(cond, prune=True, label='remotecopy').view()

        :param cond: a lambda describing a predicate over node properties.
        :type cond: a lambda predicate that received a node object as argument and returns True or False.

        :param prune: if true, nodes outside the dominated paths of matched nodes are pruned.
        :type prune: boolean

        :param label: if set, add a label to nodes in the dominated paths of matched nodes.
        :type label: string
        """
        g = self.__clone()
        for k, v in g.nodes.items():
            v.tainted = False
        for k, v in g.nodes.items():
            if cond(v):
                g.__tag(v, label)
                v.tainted = True
                self.__ft({v.oid}, label)
        if prune:
            g.__prune()
        return g

    def findPaths(self, source, sink, prune=True, label=None):
        """Finds paths from source to sink nodes matching conditions.

        Example Usage::
            def scp(df):
                return len(df[(df['proc.exe'].str.contains('scp'))])>0
            source = lambda n: scp(n.df())
            def passwd(df):
                return len(df[(df['file.path'].str.contains('passwd'))])>0
            sink = lambda n: passwd(n.df())
            g.findPaths(source, sink, prune=True, label='exfil').view()

        :param source: a lambda describing a predicate over node properties.
        :type source: a lambda predicate that received a node object as argument and returns True or False.

        :param sink: a lambda describing a predicate over node properties.
        :type sink: a lambda predicate that received a node object as argument and returns True or False.

        :param prune: if true, nodes outside the paths connecting matched nodes are pruned.
        :type prune: boolean

        :param label: if set, add a label to nodes in the paths connecting matched nodes.
        :type label: string
        """
        g1 = self.bt(sink, prune=True)
        g2 = self.ft(source, prune=True)
        g = g2.intersection(g1)
        if label:
            for k, v in g.nodes.items():
                g.__tag(v, label)
        return g if prune else self

    def __clone(self):
        g = Graphlet.__new__(Graphlet)
        g.nodes = self.nodes.copy()
        g.edges = self.edges.copy()
        return g

    def __str__(self):
        nodes = reduce(lambda s1, s2: str(s1) + str(s2) + '\n', self.nodes.values(), '')
        edges = reduce(lambda s1, s2: str(s1) + str(s2) + '\n', self.edges, '')
        return 'nodes:\n' + nodes + 'edges:\n ' + edges


class Edge(object):
    """
    **Edge**

    This class represents a graph edge, and acts as a super class for specific edges.

    :param edge: an abstract edge object.
    :type edge: sysflow.Edge
    """

    def __init__(self, n1, n2, label):
        super().__init__()
        self.n1 = n1
        self.n2 = n2
        self.label = label


class EvtEdge(Edge):
    """
    **EvtEdge**

    This class represents a graph event edge. It is used
    for sysflow event objects and subclasses Edge.

    :param evtedge: an edge object representing a sysflow evt.
    :type evtedge: sysflow.EvtEdge
    """

    def __init__(self, n1, n2, label):
        super().__init__(n1, n2, label)

    def nfrom(self):
        return self.n1

    def nto(self):
        return self.n2

    def op(self):
        return self.label

    def __key(self):
        return (self.n1, self.n2, self.label)

    def __hash__(self):
        return _hash(self.__key())

    def __eq__(self, other):
        if isinstance(other, EvtEdge):
            return self.__key() == other.__key()
        return NotImplemented

    def __str__(self):
        return 'edge: [{0}] -- {1} --> [{2}]'.format(self.n1, self.label, self.n2)


class FlowEdge(Edge):
    """
    **FlowEdge**

    This class represents a graph flow edge. It is used
    for sysflow flow objects and subclasses Edge.

    :param flowedge: an edge object representing a sysflow flow.
    :type flowedge: sysflow.FlowEdge
    """

    def __init__(self, n1, n2, label):
        super().__init__(n1, n2, label)

    def nfrom(self):
        return self.n1

    def nto(self):
        return self.n2

    def op(self):
        return self.label

    def __key(self):
        return (self.n1, self.n2, self.label)

    def __hash__(self):
        return _hash(self.__key())

    def __eq__(self, other):
        if isinstance(other, FlowEdge):
            return self.__key() == other.__key()
        return NotImplemented

    def __str__(self):
        return 'edge: [{0}] -- {1} --> [{2}]'.format(self.n1, self.label, self.n2)


class Node(object):
    """
    **Node**

    This class represents a graph node, and acts as a super class for specific nodes.

    :param node: an abstract node object.
    :type node: sysflow.Node
    """

    def __init__(self, oid):
        super().__init__()
        self.oid = oid
        self.tainted = False


class ProcessNode(Node):
    """
    **ProcessNode**

    This class represents a process node.

    :param proc: a process node object.
    :type proc: sysflow.ProcessNode
    """

    def __init__(self, oid, exe, args, uid, user, gid, group, tty):
        super().__init__(oid)
        self.type = 'P'
        self.exe = exe
        self.args = args
        self.uid = uid
        self.user = user
        self.gid = gid
        self.group = group
        self.tty = tty
        self.procs = set()
        self.data = list()

    def addProc(self, pid, createTS, r):
        self.procs.add((pid, createTS))
        if r:
            self.data.append(r)

    def hasProc(self, pid, createTS):
        return (pid, createTS) in self.procs

    def df(self):
        data = OrderedDict()
        if len(self.data) > 0:
            for idx, r in enumerate(self.data):
                data[idx] = r.values()
            return pd.DataFrame.from_dict(data, orient='index', columns=r.keys() if r else None)
        return pd.DataFrame(columns=_fields)

    def interval(self):
        ts = str(self.df()[['ts_uts']].min().to_string(index=False)).strip()
        te = str(self.df()[['ts_uts']].max().to_string(index=False)).strip()  # INFSYMB
        return (ts, te)

    def score(self):
        for r in self.data:
            if len(r['tags']) > 0:
                return r['tags'][2]
        return 0

    def tags(self):
        tags = set()
        for r in self.data:
            if len(r['tags']) > 0:
                for t in r['tags'][1]:
                    tags.add(str(t))
        return tags if len(tags) > 0 else None

    def dot(self, withoid=False, peek=True, peeksize=3, showtags=False):
        node = 'P|{{{0} [{1}]|{{{2}|{3}|{4}|{5}|{6}}}{7}}}'
        peeknode = 'P|{{{0} [{1}]|{{{2}}}|{{{3}|{4}|{5}|{6}|{7}}}{8}}}'
        oidnode = 'P|{{{0}|{{{1} [{2}]}}|{{{3}|{4}|{5}|{6}|{7}}}{8}}}'
        peekoidnode = 'P|{{{0}|{{{1} [{2}]}}|{{{3}}}|{{{4}|{5}|{6}|{7}|{8}}}{9}}}'
        reslist = ['{0}, {1}'.format(p[0], p[1]) for p in self.procs]
        details = reslist[-peeksize:] + (reslist[peeksize:] and ['...'])
        peekstr = '\\n'.join(details)
        exe = _escape(self.exe)
        args = _escape(self.args)
        tags = '|{{{0}}}'.format(str(self.tags())) if showtags and self.tags() else ''
        if peek:
            return (
                peekoidnode.format(
                    self.oid, exe, args, peekstr, self.user, self.uid, self.group, self.gid, self.tty, tags
                )
                if withoid
                else peeknode.format(exe, args, peekstr, self.user, self.uid, self.group, self.gid, self.tty, tags)
            )
        else:
            return (
                oidnode.format(self.oid, exe, args, self.user, self.uid, self.group, self.gid, self.tty, tags)
                if withoid
                else node.format(exe, args, self.user, self.uid, self.group, self.gid, self.tty, tags)
            )

    def __key(self):
        return (self.exe, self.args)

    def __hash__(self):
        return _hash(self.__key())

    def __eq__(self, other):
        if isinstance(other, ProcessNode):
            return self.__key() == other.__key()
        return NotImplemented

    def __str__(self):
        return str(self.__key())


class FileFlowNode(Node):
    """
    **FileFlowNode**

    This class represents a fileflow node.

    :param ff: a fileflow node object.
    :type ff: sysflow.FileFlow
    """

    def __init__(self, oid, exe, args):
        super().__init__(oid)
        self.type = 'FF'
        self.exe = exe
        self.args = args
        self.data = list()

    def addFlow(self, r):
        self.data.append(r)

    def hasProc(self, pid, createTS):
        return True

    def df(self):
        data = OrderedDict()
        for idx, r in enumerate(self.data):
            data[idx] = r.values()
        df = pd.DataFrame.from_dict(data, orient='index', columns=r.keys() if r else None)
        # return df[(df['file.path'] != '') & ((df['flow.rops'] > 0) | (df['flow.wops'] > 0))]
        return df[(df['file.path'] != '')]

    def interval(self):
        ts = str(self.df()[['ts_uts']].min().to_string(index=False)).strip()
        te = str(self.df()[['endts_uts']].max().to_string(index=False)).strip()  # INFSYMB
        return (ts, te)

    def plot(self):
        df = self.df()
        flows = df[(df.type.isin(['FF']))]
        ax = flows[['ts_uts', 'flow.rbytes', 'flow.wbytes']].plot.bar(
            x='ts_uts', y=['flow.rbytes', 'flow.wbytes'], rot=45, figsize=(20, 5)
        )
        ax.xaxis.set_major_locator(mdates.AutoDateLocator())
        plt.gcf().autofmt_xdate()
        plt.show()

    def describe(self):
        # Rank ordering file operations
        _df = self.df().replace('', np.nan).dropna(axis=0, how='any', subset=['file.path'])
        paths = _df.groupby(['file.path']).count()[['ts_uts']].rename(columns={"ts_uts": "count"})
        return paths.sort_values(by='count', ascending=False)

    def score(self):
        for r in self.data:
            if len(r['tags']) > 0:
                return r['tags'][2]
        return 0

    def tags(self):
        tags = set()
        for r in self.data:
            if len(r['tags']) > 0:
                for t in r['tags'][1]:
                    tags.add(str(t))
        return tags if len(tags) > 0 else None

    def dot(self, withoid=False, peek=True, peeksize=3, showtags=False):
        node = 'FF|{{{0}|{{{1}, {2}, {3}, {4}}}{5}}}'
        peeknode = 'FF|{{{{{0}|{1}, {2}, {3}, {4}}}|{{{5}}}{6}}}'
        oidnode = 'FF|{{{0}|{{{1}|{2}, {3}, {4}, {5}}}{6}}}'
        peekoidnode = 'FF|{{{0}|{{{1}|{2}, {3}, {4}, {5}}}|{{{6}}}{7}}}'
        flowstats = self.df()[['flow.rbytes', 'flow.wbytes', 'flow.rops', 'flow.wops']].sum(axis=0, skipna=True)
        rb = flowstats['flow.rbytes']
        rop = flowstats['flow.rops']
        wb = flowstats['flow.wbytes']
        wop = flowstats['flow.wops']
        ufiles = len(self.df()['file.path'].unique())
        res = self.df()[['res', 'ts_uts']].groupby(['res']).count()[['ts_uts']].rename(columns={'ts_uts': 'count'})
        reslist = res.index.tolist()
        details = reslist[-peeksize:] + (reslist[peeksize:] and ['...'])
        peekstr = _escape('\\n'.join(details))
        tags = '|{{{0}}}'.format(str(self.tags())) if showtags and self.tags() else ''
        if peek:
            return (
                peekoidnode.format(self.oid, ufiles, rb, rop, wb, wop, peekstr, tags)
                if withoid
                else peeknode.format(ufiles, rb, rop, wb, wop, peekstr, tags)
            )
        else:
            return (
                oidnode.format(self.oid, ufiles, rb, rop, wb, wop, tags)
                if withoid
                else node.format(ufiles, rb, rop, wb, wop, tags)
            )

    def __key(self):
        return (self.exe, self.args, self.type)

    def __hash__(self):
        return _hash(self.__key())

    def __eq__(self, other):
        if isinstance(other, FileFlowNode):
            return self.__key() == other.__key()
        return NotImplemented

    def __str__(self):
        return str(self.__key())


class NetFlowNode(Node):
    """
    **NetFlowNode**

    This class represents a netflow node.

    :param nf: a netflow node object.
    :type nf: sysflow.NetFlow
    """

    def __init__(self, oid, exe, args):
        super().__init__(oid)
        self.type = 'NF'
        self.exe = exe
        self.args = args
        self.data = list()

    def addFlow(self, r):
        self.data.append(r)

    def hasProc(self, pid, createTS):
        return True

    def df(self):
        data = OrderedDict()
        for idx, r in enumerate(self.data):
            data[idx] = r.values()
        df = pd.DataFrame.from_dict(data, orient='index', columns=r.keys() if r else None)
        # return df
        return df[((df['flow.rops'] > 0) | (df['flow.wops'] > 0))]

    def plot(self):
        df = self.df()
        flows = df[(df.type.isin(['NF']))]
        ax = flows[['ts_uts', 'flow.rbytes', 'flow.wbytes']].plot.bar(
            x='ts_uts', y=['flow.rbytes', 'flow.wbytes'], rot=45, figsize=(20, 5)
        )
        ax.xaxis.set_major_locator(mdates.AutoDateLocator())
        plt.gcf().autofmt_xdate()
        plt.show()

    def interval(self):
        ts = str(self.df()[['ts_uts']].min().to_string(index=False)).strip()
        te = str(self.df()[['endts_uts']].max().to_string(index=False)).strip()  # INFSYMB
        return (ts, te)

    def score(self):
        for r in self.data:
            if len(r['tags']) > 0:
                return r['tags'][2]
        return 0

    def tags(self):
        tags = set()
        for r in self.data:
            if len(r['tags']) > 0:
                for t in r['tags'][1]:
                    tags.add(str(t))
        return tags if len(tags) > 0 else None

    def dot(self, withoid=False, peek=True, peeksize=3, showtags=False):
        node = 'NF|{{{{{0}|{1}|{2}}}|{{{3}, {4}, {5}, {6}}}{7}}}'
        peeknode = 'NF|{{{{{0}|{1}|{2}}}|{{{3}, {4}, {5}, {6}}}|{{{7}}}{8}}}'
        oidnode = 'NF|{{{0}|{{{1}|{2}|{3}}}|{{{4}, {5}, {6}, {7}}}{8}}}'
        peekoidnode = 'NF|{{{0}|{{{1}|{2}|{3}}}|{{{4}, {5}, {6}, {7}}}|{{{8}}}{9}}}'
        flowstats = self.df()[['flow.rbytes', 'flow.wbytes', 'flow.rops', 'flow.wops']].sum(axis=0, skipna=True)
        rb = flowstats['flow.rbytes']
        rop = flowstats['flow.rops']
        wb = flowstats['flow.wbytes']
        wop = flowstats['flow.wops']
        uips = len(pd.unique(self.df()[['net.sip', 'net.dip']].values.ravel('K')))
        uports = len(pd.unique(self.df()[['net.sport', 'net.dport']].values.ravel('K')))
        uprotos = self.df()['net.proto'].unique()
        res = self.df()[['res', 'ts_uts']].groupby(['res']).count()[['ts_uts']].rename(columns={'ts_uts': 'count'})
        reslist = res.index.tolist()
        details = reslist[-peeksize:] + (reslist[peeksize:] and ['...'])
        peekstr = _escape('\\n'.join(details))
        tags = '|{{{0}}}'.format(str(self.tags())) if showtags and self.tags() else ''
        if peek:
            return (
                peekoidnode.format(self.oid, uips, uports, uprotos, rb, rop, wb, wop, peekstr, tags)
                if withoid
                else peeknode.format(uips, uports, uprotos, rb, rop, wb, wop, peekstr, tags)
            )
        else:
            return (
                oidnode.format(self.oid, uips, uports, uprotos, rb, rop, wb, wop, tags)
                if withoid
                else node.format(uips, uports, uprotos, rb, rop, wb, wop, tags)
            )

    def __key(self):
        return (self.exe, self.args, self.type)

    def __hash__(self):
        return _hash(self.__key())

    def __eq__(self, other):
        if isinstance(other, NetFlowNode):
            return self.__key() == other.__key()
        return NotImplemented

    def __str__(self):
        return str(self.__key())


def _hash(o):
    return int(hashlib.md5(json.dumps(o).encode('utf-8')).hexdigest(), 16)


def _escape(s):
    return s.replace('<', '\<').replace('>', '\>').replace('{', '\{').replace('}', '\}')


def _files(path):
    """list files in dir path"""
    for file in os.listdir(path):
        if os.path.isfile(os.path.join(path, file)):
            yield os.path.join(path, file)
