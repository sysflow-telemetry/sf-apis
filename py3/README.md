# SysFlow SDK and Utilities

This package includes the SDK and command-line utilities for [SysFlow](https://sysflow.io).

## Minimum requirements

Python 3.7 or higher.

## Installation

```bash
pip3 install sysflow-tools
```

## About this package

This Python package includes:  

- **sysprint**, a command-line utility that reads, prints, and converts SysFlow traces to human-readale outputs, including console, JSON, and CSV formats. It supports reading traces from local disk and from S3-compliant object stores. 

- **sysflow library**, a Python package for programming data exploration and analytics with SysFlow. It includes data manipulation using Pandas dataframes and a native query language (`sfql`) with macro support.

Check [Sysflow APIs](https://sysflow.readthedocs.io/en/latest/api-utils.html) for programmatic APIs and more information about sysprint.

## How to use sysprint

The following command shows how to run sysprint with trace files located in `/mnt/data` on the host.

```bash
sysprint /mnt/data/<trace>
```

For help and advanced options, run:

```bash
sysprint -h
```

## What is SysFlow?

The SysFlow Telemetry Pipeline is a framework for monitoring cloud workloads and for creating performance and security analytics. The goal of this project is to build all the plumbing required for system telemetry so that users can focus on writing and sharing analytics on a scalable, common open-source platform. The backbone of the telemetry pipeline is a new data format called SysFlow, which lifts raw system event information into an abstraction that describes process behaviors, and their relationships with containers, files, and network. This object-relational format is highly compact, yet it provides broad visibility into container clouds. We have also built several APIs that allow users to process SysFlow with their favorite toolkits. Learn more about SysFlow in the [SysFlow specification document](https://sysflow.readthedocs.io/en/latest/spec.html).

The SysFlow framework consists of the following sub-projects:

- [sf-apis](https://github.com/sysflow-telemetry/sf-apis) provides the SysFlow schema and programatic APIs in go, python, and C++.
- [sf-collector](https://github.com/sysflow-telemetry/sf-collector) monitors and collects system call and event information from hosts and exports them in the SysFlow format using Apache Avro object serialization.
- [sf-processor](https://github.com/sysflow-telemetry/sf-processor) provides a performance optimized policy engine for processing, enriching, filtering SysFlow events, generating alerts, and exporting the processed data to various targets.
- [sf-exporter](https://github.com/sysflow-telemetry/sf-exporter) exports SysFlow traces to S3-compliant storage systems for archival purposes.
- [sf-deployments](https://github.com/sysflow-telemetry/sf-deployments) contains deployment packages for SysFlow, including Docker, Helm, and OpenShift.
- [sysflow](https://github.com/sysflow-telemetry/sysflow) is the documentation repository and issue tracker for the SysFlow framework.


