[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sysflowtelemetry/sysprint)](https://hub.docker.com/r/sysflowtelemetry/sf-apis/builds)
[![Docker Pulls](https://img.shields.io/docker/pulls/sysflowtelemetry/sf-apis)](https://hub.docker.com/r/sysflowtelemetry/sf-exporter)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/sysflow-telemetry/sf-apis)
[![Documentation Status](https://readthedocs.org/projects/sysflow/badge/?version=latest)](https://sysflow.readthedocs.io/en/latest/?badge=latest)
[![GitHub](https://img.shields.io/github/license/sysflow-telemetry/sf-apis)](https://github.com/sysflow-telemetry/sf-apis/blob/master/LICENSE.md)

# Supported tags and respective `Dockerfile` links

-	[`0.1`, `latest`](https://github.com/sysflow-telemetry/sf-exporter/blob/0.1/Dockerfile)

# Quick reference

-	**Documentation**:  
	[the SysFlow Documentation](https://sysflow.readthedocs.io)
  
-	**Where to get help**:  
	[the SysFlow Community Slack](https://join.slack.com/t/sysflow-telemetry/shared_invite/enQtODA5OTA3NjE0MTAzLTlkMGJlZDQzYTc3MzhjMzUwNDExNmYyNWY0NWIwODNjYmRhYWEwNGU0ZmFkNGQ2NzVmYjYxMWFjYTM1MzA5YWQ)

-	**Where to file issues**:  
	[the github issue tracker](https://github.com/sysflow-telemetry/sf-docs/issues) (include the `sf-apis` tag)

-	**Source of this description**:  
	[repo's readme](https://github.com/sysflow-telemetry/sf-apis/edit/master/README.md) ([history](https://github.com/sysflow-telemetry/sf-apis/commits/master))

# What is SysFlow?

The SysFlow Telemetry Pipeline is a framework for monitoring cloud workloads and for creating performance and security analytics. The goal of this project is to build all the plumbing required for system telemetry so that users can focus on writing and sharing analytics on a scalable, common open-source platform. The backbone of the telemetry pipeline is a new data format called SysFlow, which lifts raw system event information into an abstraction that describes process behaviors, and their relationships with containers, files, and network. This object-relational format is highly compact, yet it provides broad visibility into container clouds. We have also built several APIs that allow users to process SysFlow with their favorite toolkits. Learn more about SysFlow in the [SysFlow specification document](https://sysflow.readthedocs.io/en/latest/spec.html).

# About This Image

This image packages sysprint, which reads, prints, and converts SysFlow traces to human-readale outputs, including console, JSON, and CSV formats. It supports reading traces from local disk and from S3-compliant object stores. Please check [Sysflow APIs](https://sysflow.readthedocs.io/en/latest/api-utils.html) for programmatic APIs and more information about sysprint.

# How to use this image

The easiest way to run sysprint is from a Docker container, with host mount for the directories from where to read trace files. The following command shows how to run sysprint with trace files located in `/mnt/data` on the host.

```
docker run --rm -v /mnt/data:/mnt/data sysflowtelemetry/sysprint /mnt/data/<trace>    
```
For help, run:
```
docker run --rm -v /mnt/data:/mnt/data sysflowtelemetry/sysprint -h
```

# License

View [license information](https://github.com/sysflow-telemetry/sf-exporter/blob/master/LICENSE.md) for the software contained in this image.

As with all Docker images, these likely also contain other software which may be under other licenses (such as Bash, etc from the base distribution, along with any direct or indirect dependencies of the primary software being contained).

As for any pre-built image usage, it is the image user's responsibility to ensure that any use of this image complies with any relevant licenses for all software contained within.
