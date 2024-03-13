/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */


#ifndef __SRC_SYSFLOW_SYSFLOW_HH_2119762948__H_
#define __SRC_SYSFLOW_SYSFLOW_HH_2119762948__H_


#include <sstream>
#include "boost/any.hpp"
#include "avro/Specific.hh"
#include "avro/Encoder.hh"
#include "avro/Decoder.hh"

namespace sysflow {
struct SFHeader {
    int64_t version;
    std::string exporter;
    std::string ip;
    std::string filename;
    SFHeader() :
        version(int64_t()),
        exporter(std::string()),
        ip(std::string()),
        filename(std::string())
        { }
};

enum class ContainerType: unsigned {
    CT_DOCKER,
    CT_LXC,
    CT_LIBVIRT_LXC,
    CT_MESOS,
    CT_RKT,
    CT_CUSTOM,
    CT_CRI,
    CT_CONTAINERD,
    CT_CRIO,
    CT_BPM,
};

struct _SysFlow_avsc_Union__0__ {
private:
    size_t idx_;
    boost::any value_;
public:
    size_t idx() const { return idx_; }
    bool is_null() const {
        return (idx_ == 0);
    }
    void set_null() {
        idx_ = 0;
        value_ = boost::any();
    }
    std::string get_string() const;
    void set_string(const std::string& v);
    _SysFlow_avsc_Union__0__();
};

struct Container {
    typedef _SysFlow_avsc_Union__0__ podId_t;
    std::string id;
    std::string name;
    std::string image;
    std::string imageid;
    ContainerType type;
    bool privileged;
    podId_t podId;
    Container() :
        id(std::string()),
        name(std::string()),
        image(std::string()),
        imageid(std::string()),
        type(ContainerType()),
        privileged(bool()),
        podId(podId_t())
        { }
};

enum class SFObjectState: unsigned {
    CREATED,
    MODIFIED,
    REUP,
};

struct OID {
    int64_t createTS;
    int64_t hpid;
    OID() :
        createTS(int64_t()),
        hpid(int64_t())
        { }
};

struct _SysFlow_avsc_Union__1__ {
private:
    size_t idx_;
    boost::any value_;
public:
    size_t idx() const { return idx_; }
    bool is_null() const {
        return (idx_ == 0);
    }
    void set_null() {
        idx_ = 0;
        value_ = boost::any();
    }
    OID get_OID() const;
    void set_OID(const OID& v);
    _SysFlow_avsc_Union__1__();
};

struct _SysFlow_avsc_Union__2__ {
private:
    size_t idx_;
    boost::any value_;
public:
    size_t idx() const { return idx_; }
    bool is_null() const {
        return (idx_ == 0);
    }
    void set_null() {
        idx_ = 0;
        value_ = boost::any();
    }
    std::string get_string() const;
    void set_string(const std::string& v);
    _SysFlow_avsc_Union__2__();
};

struct Process {
    typedef _SysFlow_avsc_Union__1__ poid_t;
    typedef _SysFlow_avsc_Union__2__ containerId_t;
    SFObjectState state;
    OID oid;
    poid_t poid;
    int64_t ts;
    std::string exe;
    std::string exeArgs;
    int32_t uid;
    std::string userName;
    int32_t gid;
    std::string groupName;
    bool tty;
    containerId_t containerId;
    bool entry;
    std::string cwd;
    std::vector<std::string > env;
    int64_t sid;
    std::string selabel;
    Process() :
        state(SFObjectState()),
        oid(OID()),
        poid(poid_t()),
        ts(int64_t()),
        exe(std::string()),
        exeArgs(std::string()),
        uid(int32_t()),
        userName(std::string()),
        gid(int32_t()),
        groupName(std::string()),
        tty(bool()),
        containerId(containerId_t()),
        entry(bool()),
        cwd(std::string()),
        env(std::vector<std::string >()),
        sid(int64_t()),
        selabel(std::string())
        { }
};

struct _SysFlow_avsc_Union__3__ {
private:
    size_t idx_;
    boost::any value_;
public:
    size_t idx() const { return idx_; }
    bool is_null() const {
        return (idx_ == 0);
    }
    void set_null() {
        idx_ = 0;
        value_ = boost::any();
    }
    std::string get_string() const;
    void set_string(const std::string& v);
    _SysFlow_avsc_Union__3__();
};

struct File {
    typedef _SysFlow_avsc_Union__3__ containerId_t;
    SFObjectState state;
    std::array<uint8_t, 20> oid;
    int64_t ts;
    int32_t restype;
    std::string path;
    containerId_t containerId;
    int64_t sid;
    std::string selabel;
    File() :
        state(SFObjectState()),
        oid(std::array<uint8_t, 20>()),
        ts(int64_t()),
        restype(int32_t()),
        path(std::string()),
        containerId(containerId_t()),
        sid(int64_t()),
        selabel(std::string())
        { }
};

struct ProcessEvent {
    OID procOID;
    int64_t ts;
    int64_t tid;
    int32_t opFlags;
    std::vector<std::string > args;
    int32_t ret;
    ProcessEvent() :
        procOID(OID()),
        ts(int64_t()),
        tid(int64_t()),
        opFlags(int32_t()),
        args(std::vector<std::string >()),
        ret(int32_t())
        { }
};

struct NetworkFlow {
    OID procOID;
    int64_t ts;
    int64_t tid;
    int32_t opFlags;
    int64_t endTs;
    int32_t sip;
    int32_t sport;
    int32_t dip;
    int32_t dport;
    int32_t proto;
    int32_t fd;
    int64_t numRRecvOps;
    int64_t numWSendOps;
    int64_t numRRecvBytes;
    int64_t numWSendBytes;
    NetworkFlow() :
        procOID(OID()),
        ts(int64_t()),
        tid(int64_t()),
        opFlags(int32_t()),
        endTs(int64_t()),
        sip(int32_t()),
        sport(int32_t()),
        dip(int32_t()),
        dport(int32_t()),
        proto(int32_t()),
        fd(int32_t()),
        numRRecvOps(int64_t()),
        numWSendOps(int64_t()),
        numRRecvBytes(int64_t()),
        numWSendBytes(int64_t())
        { }
};

struct FileFlow {
    OID procOID;
    int64_t ts;
    int64_t tid;
    int32_t opFlags;
    int32_t openFlags;
    int64_t endTs;
    std::array<uint8_t, 20> fileOID;
    int32_t fd;
    int64_t numRRecvOps;
    int64_t numWSendOps;
    int64_t numRRecvBytes;
    int64_t numWSendBytes;
    FileFlow() :
        procOID(OID()),
        ts(int64_t()),
        tid(int64_t()),
        opFlags(int32_t()),
        openFlags(int32_t()),
        endTs(int64_t()),
        fileOID(std::array<uint8_t, 20>()),
        fd(int32_t()),
        numRRecvOps(int64_t()),
        numWSendOps(int64_t()),
        numRRecvBytes(int64_t()),
        numWSendBytes(int64_t())
        { }
};

struct _SysFlow_avsc_Union__4__ {
private:
    size_t idx_;
    boost::any value_;
public:
    size_t idx() const { return idx_; }
    bool is_null() const {
        return (idx_ == 0);
    }
    void set_null() {
        idx_ = 0;
        value_ = boost::any();
    }
    std::array<uint8_t, 20> get_FOID() const;
    void set_FOID(const std::array<uint8_t, 20>& v);
    _SysFlow_avsc_Union__4__();
};

struct FileEvent {
    typedef _SysFlow_avsc_Union__4__ newFileOID_t;
    OID procOID;
    int64_t ts;
    int64_t tid;
    int32_t opFlags;
    std::array<uint8_t, 20> fileOID;
    int32_t ret;
    newFileOID_t newFileOID;
    FileEvent() :
        procOID(OID()),
        ts(int64_t()),
        tid(int64_t()),
        opFlags(int32_t()),
        fileOID(std::array<uint8_t, 20>()),
        ret(int32_t()),
        newFileOID(newFileOID_t())
        { }
};

struct NetworkEvent {
    OID procOID;
    int64_t ts;
    int64_t tid;
    int32_t opFlags;
    int32_t sip;
    int32_t sport;
    int32_t dip;
    int32_t dport;
    int32_t proto;
    int32_t ret;
    NetworkEvent() :
        procOID(OID()),
        ts(int64_t()),
        tid(int64_t()),
        opFlags(int32_t()),
        sip(int32_t()),
        sport(int32_t()),
        dip(int32_t()),
        dport(int32_t()),
        proto(int32_t()),
        ret(int32_t())
        { }
};

struct ProcessFlow {
    OID procOID;
    int64_t ts;
    int64_t numThreadsCloned;
    int32_t opFlags;
    int64_t endTs;
    int64_t numThreadsExited;
    int64_t numCloneErrors;
    ProcessFlow() :
        procOID(OID()),
        ts(int64_t()),
        numThreadsCloned(int64_t()),
        opFlags(int32_t()),
        endTs(int64_t()),
        numThreadsExited(int64_t()),
        numCloneErrors(int64_t())
        { }
};

struct Port {
    int32_t port;
    int32_t targetPort;
    int32_t nodePort;
    std::string proto;
    Port() :
        port(int32_t()),
        targetPort(int32_t()),
        nodePort(int32_t()),
        proto(std::string())
        { }
};

struct Service {
    std::string name;
    std::string id;
    std::string namespace_;
    std::vector<Port > portList;
    std::vector<int64_t > clusterIP;
    Service() :
        name(std::string()),
        id(std::string()),
        namespace_(std::string()),
        portList(std::vector<Port >()),
        clusterIP(std::vector<int64_t >())
        { }
};

struct Pod {
    int64_t ts;
    std::string id;
    std::string name;
    std::string nodeName;
    std::vector<int64_t > hostIP;
    std::vector<int64_t > internalIP;
    std::string namespace_;
    int64_t restartCount;
    std::map<std::string, std::string > labels;
    std::map<std::string, std::string > selectors;
    std::vector<Service > services;
    Pod() :
        ts(int64_t()),
        id(std::string()),
        name(std::string()),
        nodeName(std::string()),
        hostIP(std::vector<int64_t >()),
        internalIP(std::vector<int64_t >()),
        namespace_(std::string()),
        restartCount(int64_t()),
        labels(std::map<std::string, std::string >()),
        selectors(std::map<std::string, std::string >()),
        services(std::vector<Service >())
        { }
};

enum class K8sComponent: unsigned {
    K8S_NODES,
    K8S_NAMESPACES,
    K8S_PODS,
    K8S_REPLICATIONCONTROLLERS,
    K8S_SERVICES,
    K8S_EVENTS,
    K8S_REPLICASETS,
    K8S_DAEMONSETS,
    K8S_DEPLOYMENTS,
    K8S_UNKNOWN,
};

enum class K8sAction: unsigned {
    K8S_COMPONENT_ADDED,
    K8S_COMPONENT_MODIFIED,
    K8S_COMPONENT_DELETED,
    K8S_COMPONENT_ERROR,
    K8S_COMPONENT_NONEXISTENT,
    K8S_COMPONENT_UNKNOWN,
};

struct K8sEvent {
    K8sComponent kind;
    K8sAction action;
    int64_t ts;
    std::string message;
    K8sEvent() :
        kind(K8sComponent()),
        action(K8sAction()),
        ts(int64_t()),
        message(std::string())
        { }
};

struct _SysFlow_avsc_Union__5__ {
private:
    size_t idx_;
    boost::any value_;
public:
    size_t idx() const { return idx_; }
    SFHeader get_SFHeader() const;
    void set_SFHeader(const SFHeader& v);
    Container get_Container() const;
    void set_Container(const Container& v);
    Process get_Process() const;
    void set_Process(const Process& v);
    File get_File() const;
    void set_File(const File& v);
    ProcessEvent get_ProcessEvent() const;
    void set_ProcessEvent(const ProcessEvent& v);
    NetworkFlow get_NetworkFlow() const;
    void set_NetworkFlow(const NetworkFlow& v);
    FileFlow get_FileFlow() const;
    void set_FileFlow(const FileFlow& v);
    FileEvent get_FileEvent() const;
    void set_FileEvent(const FileEvent& v);
    NetworkEvent get_NetworkEvent() const;
    void set_NetworkEvent(const NetworkEvent& v);
    ProcessFlow get_ProcessFlow() const;
    void set_ProcessFlow(const ProcessFlow& v);
    Pod get_Pod() const;
    void set_Pod(const Pod& v);
    K8sEvent get_K8sEvent() const;
    void set_K8sEvent(const K8sEvent& v);
    _SysFlow_avsc_Union__5__();
};

struct SysFlow {
    typedef _SysFlow_avsc_Union__5__ rec_t;
    rec_t rec;
    SysFlow() :
        rec(rec_t())
        { }
};

inline
std::string _SysFlow_avsc_Union__0__::get_string() const {
    if (idx_ != 1) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<std::string >(value_);
}

inline
void _SysFlow_avsc_Union__0__::set_string(const std::string& v) {
    idx_ = 1;
    value_ = v;
}

inline
OID _SysFlow_avsc_Union__1__::get_OID() const {
    if (idx_ != 1) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<OID >(value_);
}

inline
void _SysFlow_avsc_Union__1__::set_OID(const OID& v) {
    idx_ = 1;
    value_ = v;
}

inline
std::string _SysFlow_avsc_Union__2__::get_string() const {
    if (idx_ != 1) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<std::string >(value_);
}

inline
void _SysFlow_avsc_Union__2__::set_string(const std::string& v) {
    idx_ = 1;
    value_ = v;
}

inline
std::string _SysFlow_avsc_Union__3__::get_string() const {
    if (idx_ != 1) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<std::string >(value_);
}

inline
void _SysFlow_avsc_Union__3__::set_string(const std::string& v) {
    idx_ = 1;
    value_ = v;
}

inline
std::array<uint8_t, 20> _SysFlow_avsc_Union__4__::get_FOID() const {
    if (idx_ != 1) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<std::array<uint8_t, 20> >(value_);
}

inline
void _SysFlow_avsc_Union__4__::set_FOID(const std::array<uint8_t, 20>& v) {
    idx_ = 1;
    value_ = v;
}

inline
SFHeader _SysFlow_avsc_Union__5__::get_SFHeader() const {
    if (idx_ != 0) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<SFHeader >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_SFHeader(const SFHeader& v) {
    idx_ = 0;
    value_ = v;
}

inline
Container _SysFlow_avsc_Union__5__::get_Container() const {
    if (idx_ != 1) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<Container >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_Container(const Container& v) {
    idx_ = 1;
    value_ = v;
}

inline
Process _SysFlow_avsc_Union__5__::get_Process() const {
    if (idx_ != 2) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<Process >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_Process(const Process& v) {
    idx_ = 2;
    value_ = v;
}

inline
File _SysFlow_avsc_Union__5__::get_File() const {
    if (idx_ != 3) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<File >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_File(const File& v) {
    idx_ = 3;
    value_ = v;
}

inline
ProcessEvent _SysFlow_avsc_Union__5__::get_ProcessEvent() const {
    if (idx_ != 4) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<ProcessEvent >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_ProcessEvent(const ProcessEvent& v) {
    idx_ = 4;
    value_ = v;
}

inline
NetworkFlow _SysFlow_avsc_Union__5__::get_NetworkFlow() const {
    if (idx_ != 5) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<NetworkFlow >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_NetworkFlow(const NetworkFlow& v) {
    idx_ = 5;
    value_ = v;
}

inline
FileFlow _SysFlow_avsc_Union__5__::get_FileFlow() const {
    if (idx_ != 6) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<FileFlow >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_FileFlow(const FileFlow& v) {
    idx_ = 6;
    value_ = v;
}

inline
FileEvent _SysFlow_avsc_Union__5__::get_FileEvent() const {
    if (idx_ != 7) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<FileEvent >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_FileEvent(const FileEvent& v) {
    idx_ = 7;
    value_ = v;
}

inline
NetworkEvent _SysFlow_avsc_Union__5__::get_NetworkEvent() const {
    if (idx_ != 8) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<NetworkEvent >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_NetworkEvent(const NetworkEvent& v) {
    idx_ = 8;
    value_ = v;
}

inline
ProcessFlow _SysFlow_avsc_Union__5__::get_ProcessFlow() const {
    if (idx_ != 9) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<ProcessFlow >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_ProcessFlow(const ProcessFlow& v) {
    idx_ = 9;
    value_ = v;
}

inline
Pod _SysFlow_avsc_Union__5__::get_Pod() const {
    if (idx_ != 10) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<Pod >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_Pod(const Pod& v) {
    idx_ = 10;
    value_ = v;
}

inline
K8sEvent _SysFlow_avsc_Union__5__::get_K8sEvent() const {
    if (idx_ != 11) {
        throw avro::Exception("Invalid type for union");
    }
    return boost::any_cast<K8sEvent >(value_);
}

inline
void _SysFlow_avsc_Union__5__::set_K8sEvent(const K8sEvent& v) {
    idx_ = 11;
    value_ = v;
}

inline _SysFlow_avsc_Union__0__::_SysFlow_avsc_Union__0__() : idx_(0) { }
inline _SysFlow_avsc_Union__1__::_SysFlow_avsc_Union__1__() : idx_(0) { }
inline _SysFlow_avsc_Union__2__::_SysFlow_avsc_Union__2__() : idx_(0) { }
inline _SysFlow_avsc_Union__3__::_SysFlow_avsc_Union__3__() : idx_(0) { }
inline _SysFlow_avsc_Union__4__::_SysFlow_avsc_Union__4__() : idx_(0) { }
inline _SysFlow_avsc_Union__5__::_SysFlow_avsc_Union__5__() : idx_(0), value_(SFHeader()) { }
}
namespace avro {
template<> struct codec_traits<sysflow::SFHeader> {
    static void encode(Encoder& e, const sysflow::SFHeader& v) {
        avro::encode(e, v.version);
        avro::encode(e, v.exporter);
        avro::encode(e, v.ip);
        avro::encode(e, v.filename);
    }
    static void decode(Decoder& d, sysflow::SFHeader& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.version);
                    break;
                case 1:
                    avro::decode(d, v.exporter);
                    break;
                case 2:
                    avro::decode(d, v.ip);
                    break;
                case 3:
                    avro::decode(d, v.filename);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.version);
            avro::decode(d, v.exporter);
            avro::decode(d, v.ip);
            avro::decode(d, v.filename);
        }
    }
};

template<> struct codec_traits<sysflow::ContainerType> {
    static void encode(Encoder& e, sysflow::ContainerType v) {
        if (v > sysflow::ContainerType::CT_BPM)
        {
            std::ostringstream error;
            error << "enum value " << static_cast<unsigned>(v) << " is out of bound for sysflow::ContainerType and cannot be encoded";
            throw avro::Exception(error.str());
        }
        e.encodeEnum(static_cast<size_t>(v));
    }
    static void decode(Decoder& d, sysflow::ContainerType& v) {
        size_t index = d.decodeEnum();
        if (index > static_cast<size_t>(sysflow::ContainerType::CT_BPM))
        {
            std::ostringstream error;
            error << "enum value " << index << " is out of bound for sysflow::ContainerType and cannot be decoded";
            throw avro::Exception(error.str());
        }
        v = static_cast<sysflow::ContainerType>(index);
    }
};

template<> struct codec_traits<sysflow::_SysFlow_avsc_Union__0__> {
    static void encode(Encoder& e, sysflow::_SysFlow_avsc_Union__0__ v) {
        e.encodeUnionIndex(v.idx());
        switch (v.idx()) {
        case 0:
            e.encodeNull();
            break;
        case 1:
            avro::encode(e, v.get_string());
            break;
        }
    }
    static void decode(Decoder& d, sysflow::_SysFlow_avsc_Union__0__& v) {
        size_t n = d.decodeUnionIndex();
        if (n >= 2) { throw avro::Exception("Union index too big"); }
        switch (n) {
        case 0:
            d.decodeNull();
            v.set_null();
            break;
        case 1:
            {
                std::string vv;
                avro::decode(d, vv);
                v.set_string(vv);
            }
            break;
        }
    }
};

template<> struct codec_traits<sysflow::Container> {
    static void encode(Encoder& e, const sysflow::Container& v) {
        avro::encode(e, v.id);
        avro::encode(e, v.name);
        avro::encode(e, v.image);
        avro::encode(e, v.imageid);
        avro::encode(e, v.type);
        avro::encode(e, v.privileged);
        avro::encode(e, v.podId);
    }
    static void decode(Decoder& d, sysflow::Container& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.id);
                    break;
                case 1:
                    avro::decode(d, v.name);
                    break;
                case 2:
                    avro::decode(d, v.image);
                    break;
                case 3:
                    avro::decode(d, v.imageid);
                    break;
                case 4:
                    avro::decode(d, v.type);
                    break;
                case 5:
                    avro::decode(d, v.privileged);
                    break;
                case 6:
                    avro::decode(d, v.podId);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.id);
            avro::decode(d, v.name);
            avro::decode(d, v.image);
            avro::decode(d, v.imageid);
            avro::decode(d, v.type);
            avro::decode(d, v.privileged);
            avro::decode(d, v.podId);
        }
    }
};

template<> struct codec_traits<sysflow::SFObjectState> {
    static void encode(Encoder& e, sysflow::SFObjectState v) {
        if (v > sysflow::SFObjectState::REUP)
        {
            std::ostringstream error;
            error << "enum value " << static_cast<unsigned>(v) << " is out of bound for sysflow::SFObjectState and cannot be encoded";
            throw avro::Exception(error.str());
        }
        e.encodeEnum(static_cast<size_t>(v));
    }
    static void decode(Decoder& d, sysflow::SFObjectState& v) {
        size_t index = d.decodeEnum();
        if (index > static_cast<size_t>(sysflow::SFObjectState::REUP))
        {
            std::ostringstream error;
            error << "enum value " << index << " is out of bound for sysflow::SFObjectState and cannot be decoded";
            throw avro::Exception(error.str());
        }
        v = static_cast<sysflow::SFObjectState>(index);
    }
};

template<> struct codec_traits<sysflow::OID> {
    static void encode(Encoder& e, const sysflow::OID& v) {
        avro::encode(e, v.createTS);
        avro::encode(e, v.hpid);
    }
    static void decode(Decoder& d, sysflow::OID& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.createTS);
                    break;
                case 1:
                    avro::decode(d, v.hpid);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.createTS);
            avro::decode(d, v.hpid);
        }
    }
};

template<> struct codec_traits<sysflow::_SysFlow_avsc_Union__1__> {
    static void encode(Encoder& e, sysflow::_SysFlow_avsc_Union__1__ v) {
        e.encodeUnionIndex(v.idx());
        switch (v.idx()) {
        case 0:
            e.encodeNull();
            break;
        case 1:
            avro::encode(e, v.get_OID());
            break;
        }
    }
    static void decode(Decoder& d, sysflow::_SysFlow_avsc_Union__1__& v) {
        size_t n = d.decodeUnionIndex();
        if (n >= 2) { throw avro::Exception("Union index too big"); }
        switch (n) {
        case 0:
            d.decodeNull();
            v.set_null();
            break;
        case 1:
            {
                sysflow::OID vv;
                avro::decode(d, vv);
                v.set_OID(vv);
            }
            break;
        }
    }
};

template<> struct codec_traits<sysflow::_SysFlow_avsc_Union__2__> {
    static void encode(Encoder& e, sysflow::_SysFlow_avsc_Union__2__ v) {
        e.encodeUnionIndex(v.idx());
        switch (v.idx()) {
        case 0:
            e.encodeNull();
            break;
        case 1:
            avro::encode(e, v.get_string());
            break;
        }
    }
    static void decode(Decoder& d, sysflow::_SysFlow_avsc_Union__2__& v) {
        size_t n = d.decodeUnionIndex();
        if (n >= 2) { throw avro::Exception("Union index too big"); }
        switch (n) {
        case 0:
            d.decodeNull();
            v.set_null();
            break;
        case 1:
            {
                std::string vv;
                avro::decode(d, vv);
                v.set_string(vv);
            }
            break;
        }
    }
};

template<> struct codec_traits<sysflow::Process> {
    static void encode(Encoder& e, const sysflow::Process& v) {
        avro::encode(e, v.state);
        avro::encode(e, v.oid);
        avro::encode(e, v.poid);
        avro::encode(e, v.ts);
        avro::encode(e, v.exe);
        avro::encode(e, v.exeArgs);
        avro::encode(e, v.uid);
        avro::encode(e, v.userName);
        avro::encode(e, v.gid);
        avro::encode(e, v.groupName);
        avro::encode(e, v.tty);
        avro::encode(e, v.containerId);
        avro::encode(e, v.entry);
        avro::encode(e, v.cwd);
        avro::encode(e, v.env);
        avro::encode(e, v.sid);
        avro::encode(e, v.selabel);
    }
    static void decode(Decoder& d, sysflow::Process& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.state);
                    break;
                case 1:
                    avro::decode(d, v.oid);
                    break;
                case 2:
                    avro::decode(d, v.poid);
                    break;
                case 3:
                    avro::decode(d, v.ts);
                    break;
                case 4:
                    avro::decode(d, v.exe);
                    break;
                case 5:
                    avro::decode(d, v.exeArgs);
                    break;
                case 6:
                    avro::decode(d, v.uid);
                    break;
                case 7:
                    avro::decode(d, v.userName);
                    break;
                case 8:
                    avro::decode(d, v.gid);
                    break;
                case 9:
                    avro::decode(d, v.groupName);
                    break;
                case 10:
                    avro::decode(d, v.tty);
                    break;
                case 11:
                    avro::decode(d, v.containerId);
                    break;
                case 12:
                    avro::decode(d, v.entry);
                    break;
                case 13:
                    avro::decode(d, v.cwd);
                    break;
                case 14:
                    avro::decode(d, v.env);
                    break;
                case 15:
                    avro::decode(d, v.sid);
                    break;
                case 16:
                    avro::decode(d, v.selabel);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.state);
            avro::decode(d, v.oid);
            avro::decode(d, v.poid);
            avro::decode(d, v.ts);
            avro::decode(d, v.exe);
            avro::decode(d, v.exeArgs);
            avro::decode(d, v.uid);
            avro::decode(d, v.userName);
            avro::decode(d, v.gid);
            avro::decode(d, v.groupName);
            avro::decode(d, v.tty);
            avro::decode(d, v.containerId);
            avro::decode(d, v.entry);
            avro::decode(d, v.cwd);
            avro::decode(d, v.env);
            avro::decode(d, v.sid);
            avro::decode(d, v.selabel);
        }
    }
};

template<> struct codec_traits<sysflow::_SysFlow_avsc_Union__3__> {
    static void encode(Encoder& e, sysflow::_SysFlow_avsc_Union__3__ v) {
        e.encodeUnionIndex(v.idx());
        switch (v.idx()) {
        case 0:
            e.encodeNull();
            break;
        case 1:
            avro::encode(e, v.get_string());
            break;
        }
    }
    static void decode(Decoder& d, sysflow::_SysFlow_avsc_Union__3__& v) {
        size_t n = d.decodeUnionIndex();
        if (n >= 2) { throw avro::Exception("Union index too big"); }
        switch (n) {
        case 0:
            d.decodeNull();
            v.set_null();
            break;
        case 1:
            {
                std::string vv;
                avro::decode(d, vv);
                v.set_string(vv);
            }
            break;
        }
    }
};

template<> struct codec_traits<sysflow::File> {
    static void encode(Encoder& e, const sysflow::File& v) {
        avro::encode(e, v.state);
        avro::encode(e, v.oid);
        avro::encode(e, v.ts);
        avro::encode(e, v.restype);
        avro::encode(e, v.path);
        avro::encode(e, v.containerId);
        avro::encode(e, v.sid);
        avro::encode(e, v.selabel);
    }
    static void decode(Decoder& d, sysflow::File& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.state);
                    break;
                case 1:
                    avro::decode(d, v.oid);
                    break;
                case 2:
                    avro::decode(d, v.ts);
                    break;
                case 3:
                    avro::decode(d, v.restype);
                    break;
                case 4:
                    avro::decode(d, v.path);
                    break;
                case 5:
                    avro::decode(d, v.containerId);
                    break;
                case 6:
                    avro::decode(d, v.sid);
                    break;
                case 7:
                    avro::decode(d, v.selabel);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.state);
            avro::decode(d, v.oid);
            avro::decode(d, v.ts);
            avro::decode(d, v.restype);
            avro::decode(d, v.path);
            avro::decode(d, v.containerId);
            avro::decode(d, v.sid);
            avro::decode(d, v.selabel);
        }
    }
};

template<> struct codec_traits<sysflow::ProcessEvent> {
    static void encode(Encoder& e, const sysflow::ProcessEvent& v) {
        avro::encode(e, v.procOID);
        avro::encode(e, v.ts);
        avro::encode(e, v.tid);
        avro::encode(e, v.opFlags);
        avro::encode(e, v.args);
        avro::encode(e, v.ret);
    }
    static void decode(Decoder& d, sysflow::ProcessEvent& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.procOID);
                    break;
                case 1:
                    avro::decode(d, v.ts);
                    break;
                case 2:
                    avro::decode(d, v.tid);
                    break;
                case 3:
                    avro::decode(d, v.opFlags);
                    break;
                case 4:
                    avro::decode(d, v.args);
                    break;
                case 5:
                    avro::decode(d, v.ret);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.procOID);
            avro::decode(d, v.ts);
            avro::decode(d, v.tid);
            avro::decode(d, v.opFlags);
            avro::decode(d, v.args);
            avro::decode(d, v.ret);
        }
    }
};

template<> struct codec_traits<sysflow::NetworkFlow> {
    static void encode(Encoder& e, const sysflow::NetworkFlow& v) {
        avro::encode(e, v.procOID);
        avro::encode(e, v.ts);
        avro::encode(e, v.tid);
        avro::encode(e, v.opFlags);
        avro::encode(e, v.endTs);
        avro::encode(e, v.sip);
        avro::encode(e, v.sport);
        avro::encode(e, v.dip);
        avro::encode(e, v.dport);
        avro::encode(e, v.proto);
        avro::encode(e, v.fd);
        avro::encode(e, v.numRRecvOps);
        avro::encode(e, v.numWSendOps);
        avro::encode(e, v.numRRecvBytes);
        avro::encode(e, v.numWSendBytes);
    }
    static void decode(Decoder& d, sysflow::NetworkFlow& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.procOID);
                    break;
                case 1:
                    avro::decode(d, v.ts);
                    break;
                case 2:
                    avro::decode(d, v.tid);
                    break;
                case 3:
                    avro::decode(d, v.opFlags);
                    break;
                case 4:
                    avro::decode(d, v.endTs);
                    break;
                case 5:
                    avro::decode(d, v.sip);
                    break;
                case 6:
                    avro::decode(d, v.sport);
                    break;
                case 7:
                    avro::decode(d, v.dip);
                    break;
                case 8:
                    avro::decode(d, v.dport);
                    break;
                case 9:
                    avro::decode(d, v.proto);
                    break;
                case 10:
                    avro::decode(d, v.fd);
                    break;
                case 11:
                    avro::decode(d, v.numRRecvOps);
                    break;
                case 12:
                    avro::decode(d, v.numWSendOps);
                    break;
                case 13:
                    avro::decode(d, v.numRRecvBytes);
                    break;
                case 14:
                    avro::decode(d, v.numWSendBytes);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.procOID);
            avro::decode(d, v.ts);
            avro::decode(d, v.tid);
            avro::decode(d, v.opFlags);
            avro::decode(d, v.endTs);
            avro::decode(d, v.sip);
            avro::decode(d, v.sport);
            avro::decode(d, v.dip);
            avro::decode(d, v.dport);
            avro::decode(d, v.proto);
            avro::decode(d, v.fd);
            avro::decode(d, v.numRRecvOps);
            avro::decode(d, v.numWSendOps);
            avro::decode(d, v.numRRecvBytes);
            avro::decode(d, v.numWSendBytes);
        }
    }
};

template<> struct codec_traits<sysflow::FileFlow> {
    static void encode(Encoder& e, const sysflow::FileFlow& v) {
        avro::encode(e, v.procOID);
        avro::encode(e, v.ts);
        avro::encode(e, v.tid);
        avro::encode(e, v.opFlags);
        avro::encode(e, v.openFlags);
        avro::encode(e, v.endTs);
        avro::encode(e, v.fileOID);
        avro::encode(e, v.fd);
        avro::encode(e, v.numRRecvOps);
        avro::encode(e, v.numWSendOps);
        avro::encode(e, v.numRRecvBytes);
        avro::encode(e, v.numWSendBytes);
    }
    static void decode(Decoder& d, sysflow::FileFlow& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.procOID);
                    break;
                case 1:
                    avro::decode(d, v.ts);
                    break;
                case 2:
                    avro::decode(d, v.tid);
                    break;
                case 3:
                    avro::decode(d, v.opFlags);
                    break;
                case 4:
                    avro::decode(d, v.openFlags);
                    break;
                case 5:
                    avro::decode(d, v.endTs);
                    break;
                case 6:
                    avro::decode(d, v.fileOID);
                    break;
                case 7:
                    avro::decode(d, v.fd);
                    break;
                case 8:
                    avro::decode(d, v.numRRecvOps);
                    break;
                case 9:
                    avro::decode(d, v.numWSendOps);
                    break;
                case 10:
                    avro::decode(d, v.numRRecvBytes);
                    break;
                case 11:
                    avro::decode(d, v.numWSendBytes);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.procOID);
            avro::decode(d, v.ts);
            avro::decode(d, v.tid);
            avro::decode(d, v.opFlags);
            avro::decode(d, v.openFlags);
            avro::decode(d, v.endTs);
            avro::decode(d, v.fileOID);
            avro::decode(d, v.fd);
            avro::decode(d, v.numRRecvOps);
            avro::decode(d, v.numWSendOps);
            avro::decode(d, v.numRRecvBytes);
            avro::decode(d, v.numWSendBytes);
        }
    }
};

template<> struct codec_traits<sysflow::_SysFlow_avsc_Union__4__> {
    static void encode(Encoder& e, sysflow::_SysFlow_avsc_Union__4__ v) {
        e.encodeUnionIndex(v.idx());
        switch (v.idx()) {
        case 0:
            e.encodeNull();
            break;
        case 1:
            avro::encode(e, v.get_FOID());
            break;
        }
    }
    static void decode(Decoder& d, sysflow::_SysFlow_avsc_Union__4__& v) {
        size_t n = d.decodeUnionIndex();
        if (n >= 2) { throw avro::Exception("Union index too big"); }
        switch (n) {
        case 0:
            d.decodeNull();
            v.set_null();
            break;
        case 1:
            {
                std::array<uint8_t, 20> vv;
                avro::decode(d, vv);
                v.set_FOID(vv);
            }
            break;
        }
    }
};

template<> struct codec_traits<sysflow::FileEvent> {
    static void encode(Encoder& e, const sysflow::FileEvent& v) {
        avro::encode(e, v.procOID);
        avro::encode(e, v.ts);
        avro::encode(e, v.tid);
        avro::encode(e, v.opFlags);
        avro::encode(e, v.fileOID);
        avro::encode(e, v.ret);
        avro::encode(e, v.newFileOID);
    }
    static void decode(Decoder& d, sysflow::FileEvent& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.procOID);
                    break;
                case 1:
                    avro::decode(d, v.ts);
                    break;
                case 2:
                    avro::decode(d, v.tid);
                    break;
                case 3:
                    avro::decode(d, v.opFlags);
                    break;
                case 4:
                    avro::decode(d, v.fileOID);
                    break;
                case 5:
                    avro::decode(d, v.ret);
                    break;
                case 6:
                    avro::decode(d, v.newFileOID);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.procOID);
            avro::decode(d, v.ts);
            avro::decode(d, v.tid);
            avro::decode(d, v.opFlags);
            avro::decode(d, v.fileOID);
            avro::decode(d, v.ret);
            avro::decode(d, v.newFileOID);
        }
    }
};

template<> struct codec_traits<sysflow::NetworkEvent> {
    static void encode(Encoder& e, const sysflow::NetworkEvent& v) {
        avro::encode(e, v.procOID);
        avro::encode(e, v.ts);
        avro::encode(e, v.tid);
        avro::encode(e, v.opFlags);
        avro::encode(e, v.sip);
        avro::encode(e, v.sport);
        avro::encode(e, v.dip);
        avro::encode(e, v.dport);
        avro::encode(e, v.proto);
        avro::encode(e, v.ret);
    }
    static void decode(Decoder& d, sysflow::NetworkEvent& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.procOID);
                    break;
                case 1:
                    avro::decode(d, v.ts);
                    break;
                case 2:
                    avro::decode(d, v.tid);
                    break;
                case 3:
                    avro::decode(d, v.opFlags);
                    break;
                case 4:
                    avro::decode(d, v.sip);
                    break;
                case 5:
                    avro::decode(d, v.sport);
                    break;
                case 6:
                    avro::decode(d, v.dip);
                    break;
                case 7:
                    avro::decode(d, v.dport);
                    break;
                case 8:
                    avro::decode(d, v.proto);
                    break;
                case 9:
                    avro::decode(d, v.ret);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.procOID);
            avro::decode(d, v.ts);
            avro::decode(d, v.tid);
            avro::decode(d, v.opFlags);
            avro::decode(d, v.sip);
            avro::decode(d, v.sport);
            avro::decode(d, v.dip);
            avro::decode(d, v.dport);
            avro::decode(d, v.proto);
            avro::decode(d, v.ret);
        }
    }
};

template<> struct codec_traits<sysflow::ProcessFlow> {
    static void encode(Encoder& e, const sysflow::ProcessFlow& v) {
        avro::encode(e, v.procOID);
        avro::encode(e, v.ts);
        avro::encode(e, v.numThreadsCloned);
        avro::encode(e, v.opFlags);
        avro::encode(e, v.endTs);
        avro::encode(e, v.numThreadsExited);
        avro::encode(e, v.numCloneErrors);
    }
    static void decode(Decoder& d, sysflow::ProcessFlow& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.procOID);
                    break;
                case 1:
                    avro::decode(d, v.ts);
                    break;
                case 2:
                    avro::decode(d, v.numThreadsCloned);
                    break;
                case 3:
                    avro::decode(d, v.opFlags);
                    break;
                case 4:
                    avro::decode(d, v.endTs);
                    break;
                case 5:
                    avro::decode(d, v.numThreadsExited);
                    break;
                case 6:
                    avro::decode(d, v.numCloneErrors);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.procOID);
            avro::decode(d, v.ts);
            avro::decode(d, v.numThreadsCloned);
            avro::decode(d, v.opFlags);
            avro::decode(d, v.endTs);
            avro::decode(d, v.numThreadsExited);
            avro::decode(d, v.numCloneErrors);
        }
    }
};

template<> struct codec_traits<sysflow::Port> {
    static void encode(Encoder& e, const sysflow::Port& v) {
        avro::encode(e, v.port);
        avro::encode(e, v.targetPort);
        avro::encode(e, v.nodePort);
        avro::encode(e, v.proto);
    }
    static void decode(Decoder& d, sysflow::Port& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.port);
                    break;
                case 1:
                    avro::decode(d, v.targetPort);
                    break;
                case 2:
                    avro::decode(d, v.nodePort);
                    break;
                case 3:
                    avro::decode(d, v.proto);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.port);
            avro::decode(d, v.targetPort);
            avro::decode(d, v.nodePort);
            avro::decode(d, v.proto);
        }
    }
};

template<> struct codec_traits<sysflow::Service> {
    static void encode(Encoder& e, const sysflow::Service& v) {
        avro::encode(e, v.name);
        avro::encode(e, v.id);
        avro::encode(e, v.namespace_);
        avro::encode(e, v.portList);
        avro::encode(e, v.clusterIP);
    }
    static void decode(Decoder& d, sysflow::Service& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.name);
                    break;
                case 1:
                    avro::decode(d, v.id);
                    break;
                case 2:
                    avro::decode(d, v.namespace_);
                    break;
                case 3:
                    avro::decode(d, v.portList);
                    break;
                case 4:
                    avro::decode(d, v.clusterIP);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.name);
            avro::decode(d, v.id);
            avro::decode(d, v.namespace_);
            avro::decode(d, v.portList);
            avro::decode(d, v.clusterIP);
        }
    }
};

template<> struct codec_traits<sysflow::Pod> {
    static void encode(Encoder& e, const sysflow::Pod& v) {
        avro::encode(e, v.ts);
        avro::encode(e, v.id);
        avro::encode(e, v.name);
        avro::encode(e, v.nodeName);
        avro::encode(e, v.hostIP);
        avro::encode(e, v.internalIP);
        avro::encode(e, v.namespace_);
        avro::encode(e, v.restartCount);
        avro::encode(e, v.labels);
        avro::encode(e, v.selectors);
        avro::encode(e, v.services);
    }
    static void decode(Decoder& d, sysflow::Pod& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.ts);
                    break;
                case 1:
                    avro::decode(d, v.id);
                    break;
                case 2:
                    avro::decode(d, v.name);
                    break;
                case 3:
                    avro::decode(d, v.nodeName);
                    break;
                case 4:
                    avro::decode(d, v.hostIP);
                    break;
                case 5:
                    avro::decode(d, v.internalIP);
                    break;
                case 6:
                    avro::decode(d, v.namespace_);
                    break;
                case 7:
                    avro::decode(d, v.restartCount);
                    break;
                case 8:
                    avro::decode(d, v.labels);
                    break;
                case 9:
                    avro::decode(d, v.selectors);
                    break;
                case 10:
                    avro::decode(d, v.services);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.ts);
            avro::decode(d, v.id);
            avro::decode(d, v.name);
            avro::decode(d, v.nodeName);
            avro::decode(d, v.hostIP);
            avro::decode(d, v.internalIP);
            avro::decode(d, v.namespace_);
            avro::decode(d, v.restartCount);
            avro::decode(d, v.labels);
            avro::decode(d, v.selectors);
            avro::decode(d, v.services);
        }
    }
};

template<> struct codec_traits<sysflow::K8sComponent> {
    static void encode(Encoder& e, sysflow::K8sComponent v) {
        if (v > sysflow::K8sComponent::K8S_UNKNOWN)
        {
            std::ostringstream error;
            error << "enum value " << static_cast<unsigned>(v) << " is out of bound for sysflow::K8sComponent and cannot be encoded";
            throw avro::Exception(error.str());
        }
        e.encodeEnum(static_cast<size_t>(v));
    }
    static void decode(Decoder& d, sysflow::K8sComponent& v) {
        size_t index = d.decodeEnum();
        if (index > static_cast<size_t>(sysflow::K8sComponent::K8S_UNKNOWN))
        {
            std::ostringstream error;
            error << "enum value " << index << " is out of bound for sysflow::K8sComponent and cannot be decoded";
            throw avro::Exception(error.str());
        }
        v = static_cast<sysflow::K8sComponent>(index);
    }
};

template<> struct codec_traits<sysflow::K8sAction> {
    static void encode(Encoder& e, sysflow::K8sAction v) {
        if (v > sysflow::K8sAction::K8S_COMPONENT_UNKNOWN)
        {
            std::ostringstream error;
            error << "enum value " << static_cast<unsigned>(v) << " is out of bound for sysflow::K8sAction and cannot be encoded";
            throw avro::Exception(error.str());
        }
        e.encodeEnum(static_cast<size_t>(v));
    }
    static void decode(Decoder& d, sysflow::K8sAction& v) {
        size_t index = d.decodeEnum();
        if (index > static_cast<size_t>(sysflow::K8sAction::K8S_COMPONENT_UNKNOWN))
        {
            std::ostringstream error;
            error << "enum value " << index << " is out of bound for sysflow::K8sAction and cannot be decoded";
            throw avro::Exception(error.str());
        }
        v = static_cast<sysflow::K8sAction>(index);
    }
};

template<> struct codec_traits<sysflow::K8sEvent> {
    static void encode(Encoder& e, const sysflow::K8sEvent& v) {
        avro::encode(e, v.kind);
        avro::encode(e, v.action);
        avro::encode(e, v.ts);
        avro::encode(e, v.message);
    }
    static void decode(Decoder& d, sysflow::K8sEvent& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.kind);
                    break;
                case 1:
                    avro::decode(d, v.action);
                    break;
                case 2:
                    avro::decode(d, v.ts);
                    break;
                case 3:
                    avro::decode(d, v.message);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.kind);
            avro::decode(d, v.action);
            avro::decode(d, v.ts);
            avro::decode(d, v.message);
        }
    }
};

template<> struct codec_traits<sysflow::_SysFlow_avsc_Union__5__> {
    static void encode(Encoder& e, sysflow::_SysFlow_avsc_Union__5__ v) {
        e.encodeUnionIndex(v.idx());
        switch (v.idx()) {
        case 0:
            avro::encode(e, v.get_SFHeader());
            break;
        case 1:
            avro::encode(e, v.get_Container());
            break;
        case 2:
            avro::encode(e, v.get_Process());
            break;
        case 3:
            avro::encode(e, v.get_File());
            break;
        case 4:
            avro::encode(e, v.get_ProcessEvent());
            break;
        case 5:
            avro::encode(e, v.get_NetworkFlow());
            break;
        case 6:
            avro::encode(e, v.get_FileFlow());
            break;
        case 7:
            avro::encode(e, v.get_FileEvent());
            break;
        case 8:
            avro::encode(e, v.get_NetworkEvent());
            break;
        case 9:
            avro::encode(e, v.get_ProcessFlow());
            break;
        case 10:
            avro::encode(e, v.get_Pod());
            break;
        case 11:
            avro::encode(e, v.get_K8sEvent());
            break;
        }
    }
    static void decode(Decoder& d, sysflow::_SysFlow_avsc_Union__5__& v) {
        size_t n = d.decodeUnionIndex();
        if (n >= 12) { throw avro::Exception("Union index too big"); }
        switch (n) {
        case 0:
            {
                sysflow::SFHeader vv;
                avro::decode(d, vv);
                v.set_SFHeader(vv);
            }
            break;
        case 1:
            {
                sysflow::Container vv;
                avro::decode(d, vv);
                v.set_Container(vv);
            }
            break;
        case 2:
            {
                sysflow::Process vv;
                avro::decode(d, vv);
                v.set_Process(vv);
            }
            break;
        case 3:
            {
                sysflow::File vv;
                avro::decode(d, vv);
                v.set_File(vv);
            }
            break;
        case 4:
            {
                sysflow::ProcessEvent vv;
                avro::decode(d, vv);
                v.set_ProcessEvent(vv);
            }
            break;
        case 5:
            {
                sysflow::NetworkFlow vv;
                avro::decode(d, vv);
                v.set_NetworkFlow(vv);
            }
            break;
        case 6:
            {
                sysflow::FileFlow vv;
                avro::decode(d, vv);
                v.set_FileFlow(vv);
            }
            break;
        case 7:
            {
                sysflow::FileEvent vv;
                avro::decode(d, vv);
                v.set_FileEvent(vv);
            }
            break;
        case 8:
            {
                sysflow::NetworkEvent vv;
                avro::decode(d, vv);
                v.set_NetworkEvent(vv);
            }
            break;
        case 9:
            {
                sysflow::ProcessFlow vv;
                avro::decode(d, vv);
                v.set_ProcessFlow(vv);
            }
            break;
        case 10:
            {
                sysflow::Pod vv;
                avro::decode(d, vv);
                v.set_Pod(vv);
            }
            break;
        case 11:
            {
                sysflow::K8sEvent vv;
                avro::decode(d, vv);
                v.set_K8sEvent(vv);
            }
            break;
        }
    }
};

template<> struct codec_traits<sysflow::SysFlow> {
    static void encode(Encoder& e, const sysflow::SysFlow& v) {
        avro::encode(e, v.rec);
    }
    static void decode(Decoder& d, sysflow::SysFlow& v) {
        if (avro::ResolvingDecoder *rd =
            dynamic_cast<avro::ResolvingDecoder *>(&d)) {
            const std::vector<size_t> fo = rd->fieldOrder();
            for (std::vector<size_t>::const_iterator it = fo.begin();
                it != fo.end(); ++it) {
                switch (*it) {
                case 0:
                    avro::decode(d, v.rec);
                    break;
                default:
                    break;
                }
            }
        } else {
            avro::decode(d, v.rec);
        }
    }
};

}
#endif
