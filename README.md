[![Build Status](https://img.shields.io/github/workflow/status/sysflow-telemetry/sf-apis/ci)](https://github.com/sysflow-telemetry/sf-apis/actions)
[![Docker Pulls](https://img.shields.io/docker/pulls/sysflowtelemetry/sysprint)](https://hub.docker.com/r/sysflowtelemetry/sysprint)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/sysflow-telemetry/sf-apis)
[![Documentation Status](https://readthedocs.org/projects/sysflow/badge/?version=latest)](https://sysflow.readthedocs.io/en/latest/?badge=latest)
[![GitHub](https://img.shields.io/github/license/sysflow-telemetry/sf-apis)](https://github.com/sysflow-telemetry/sf-apis/blob/master/LICENSE.md)

# Supported tags and respective `Dockerfile` links

-	[`0.4.3`, `latest`](https://github.com/sysflow-telemetry/sf-apis/blob/0.4.3/Dockerfile), [`edge`](https://github.com/sysflow-telemetry/sf-apis/blob/master/Dockerfile), [`dev`](https://github.com/sysflow-telemetry/sf-apis/blob/dev/Dockerfile)

# Quick reference

-	**Documentation**:
	[the SysFlow Documentation](https://sysflow.readthedocs.io)

-	**Where to get help**:
	[the SysFlow Community Slack](https://join.slack.com/t/sysflow-telemetry/shared_invite/enQtODA5OTA3NjE0MTAzLTlkMGJlZDQzYTc3MzhjMzUwNDExNmYyNWY0NWIwODNjYmRhYWEwNGU0ZmFkNGQ2NzVmYjYxMWFjYTM1MzA5YWQ)

-	**Where to file issues**:
	[the github issue tracker](https://github.com/sysflow-telemetry/sysflow/issues) (include the `sf-apis` tag)

-	**Source of this description**:
	[repo's readme](https://github.com/sysflow-telemetry/sf-apis/edit/master/README.md) ([history](https://github.com/sysflow-telemetry/sf-apis/commits/master))

-	**Docker images**:
	[docker hub](https://hub.docker.com/u/sysflowtelemetry) | [GHCR](https://github.com/orgs/sysflow-telemetry/packages)

# What is SysFlow?

The SysFlow Telemetry Pipeline is a framework for monitoring cloud workloads and for creating performance and security analytics. The goal of this project is to build all the plumbing required for system telemetry so that users can focus on writing and sharing analytics on a scalable, common open-source platform. The backbone of the telemetry pipeline is a new data format called SysFlow, which lifts raw system event information into an abstraction that describes process behaviors, and their relationships with containers, files, and network. This object-relational format is highly compact, yet it provides broad visibility into container clouds. We have also built several APIs that allow users to process SysFlow with their favorite toolkits. Learn more about SysFlow in the [SysFlow specification document](https://sysflow.readthedocs.io/en/latest/spec.html).

The SysFlow framework consists of the following sub-projects:

- [sf-apis](https://github.com/sysflow-telemetry/sf-apis) provides the SysFlow schema and programatic APIs in go, python, and C++.
- [sf-collector](https://github.com/sysflow-telemetry/sf-collector) monitors and collects system call and event information from hosts and exports them in the SysFlow format using Apache Avro object serialization.
- [sf-processor](https://github.com/sysflow-telemetry/sf-processor) provides a performance optimized policy engine for processing, enriching, filtering SysFlow events, generating alerts, and exporting the processed data to various targets.
- [sf-exporter](https://github.com/sysflow-telemetry/sf-exporter) exports SysFlow traces to S3-compliant storage systems for archival purposes.
- [sf-deployments](https://github.com/sysflow-telemetry/sf-deployments) contains deployment packages for SysFlow, including Docker, Helm, and OpenShift.
- [sysflow](https://github.com/sysflow-telemetry/sysflow) is the documentation repository and issue tracker for the SysFlow framework.

# About this repository

This repository packages two images:

- **sysprint**, which reads, prints, and converts SysFlow traces to human-readale outputs, including console, JSON, and CSV formats. It supports reading traces from local disk and from S3-compliant object stores. Please check [Sysflow APIs](https://sysflow.readthedocs.io/en/latest/api-utils.html) for programmatic APIs and more information about sysprint.

- **sfnb**, a Jupyter Notebook for performing data exploration and analytics with SysFlow. It includes data manipulation using Pandas dataframes  and a native query language (`sfql`) with macro support.

Please check [Sysflow APIs](https://sysflow.readthedocs.io/en/latest/api-utils.html) for programmatic APIs and more information about sysprint.

# How to use sysprint

The easiest way to run sysprint is from a Docker container, with host mount for the directories from where to read trace files. The following command shows how to run sysprint with trace files located in `/mnt/data` on the host.

```
docker run --rm -v /mnt/data:/mnt/data sysflowtelemetry/sysprint /mnt/data/<trace>
```
For help, run:
```
docker run --rm -v /mnt/data:/mnt/data sysflowtelemetry/sysprint -h
```

# How to use sfnb

The following command shows how to run sfnb.

```
docker run --rm -d --name sfnb -p 8888:8888 sysflowtelemetry/sfnb
```

To mount example notebooks and data files into Jupyter's `work` directory, git clone this repository locally, cd into it, and run:

```
docker run --rm -d --name sfnb -v $(pwd)/pynb:/home/jovyan/work -p 8888:8888 sysflowtelemetry/sfnb
```

Then, open a web browser and point it to `http://localhost:8888` (alternatively, the remote server name or IP where the notebook is hosted). To obtain the notebook authentication token, run `docker logs sfnb`.

Note: If classic Jupyter notebook is preferred, run:

```
docker run --rm -d --name sfnb -v $(pwd)/pynb:/home/jovyan/work -p 8888:8888 -e DOCKER_STACKS_JUPYTER_CMD=notebook sysflowtelemetry/sfnb
```

# License

View [license information](https://github.com/sysflow-telemetry/sf-apis/blob/master/LICENSE.md) for the software contained in this image.

As with all Docker images, these likely also contain other software which may be under other licenses (such as Bash, etc from the base distribution, along with any direct or indirect dependencies of the primary software being contained).

As for any pre-built image usage, it is the image user's responsibility to ensure that any use of this image complies with any relevant licenses for all software contained within.
