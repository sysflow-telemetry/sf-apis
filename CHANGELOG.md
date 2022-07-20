# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/).

> **Types of changes:**
>
> - **Added**: for new features.
> - **Changed**: for changes in existing functionality.
> - **Deprecated**: for soon-to-be removed features.
> - **Removed**: for now removed features.
> - **Fixed**: for any bug fixes.
> - **Security**: in case of vulnerabilities.

## [Unreleased]

## [0.5.0] - 2022-07-31

### Added

- Add k8s pod and event object support to sysflow spec and APIs
- Add enumeration of sysflow types in C++ API.

## [0.4.3] - 2022-06-21

### Changed

- Bumped SysFlow version to 0.4.3

## [0.4.2] - 2022-06-13

### Changed

- Bumped SysFlow version to 0.4.2

### Removed

- Removed unused container package from Go APIs

## [0.4.1] - 2022-05-26

### Added

- Updated avdl, c++, and go bindings to support k8s events and metadata (preparation for 0.5.0)

### Changed

- Bumped sysprint UBI to 8.6-751
- Updated pynb requirements.txt
- Updated query/policy language syntax to support rules and tagging

### Fixed

- Minor bug fixes in py3 APIs

## [0.4.0] - 2022-02-18

### Added

- Go: Added set data structures
- Go: Added contextual SysFlow record to capture provenance information
- Pynb: Added new notebook on MITRE ATT&CK tagging and visualization

### Changed

- BREAKING Go: renamed the unions in golang classes
- Pynb: restrictured working directory for SysFlow notebooks
- Pynb: updated base Jupyter image, making Jupyter lab the default environment for SysFlow notebooks (see updated usage in README.md)

### Security

- Update github.com/containers/storage to fix CVE-2021-20291

## [0.3.1] - 2021-09-29

### Changed

- Update(ubi): Bumped UBI-minimal version to 8.4-210 in sysprint.
- Update(py3): Updated log level on warning messages in SysFlow reader API.

## [0.3.0] - 2021-09-20

### Added

- Added secret vault wrapper package to Go API.
- Added hashing utility package to Go API.
- Added trace attribute to SysFlow schema.

### Changed

- Moved away from Dockerhub CI.
- Updated verstions of python API dependencies.
- Fixed lint issues in Python and Go APIs.
- Refactored processor plugin interfaces in Go APIs.

## [0.2.2] - 2020-12-07

### Changed

- Fixed versions of Pandas and numpy in python APIs.

## [0.2.1] - 2020-12-02

### Added

- Adds flattened indices for file OID attributes in go API.

## [0.2.0] - 2020-12-01

### Added

- Implemented ProcessFlow support for sysprint.
- Added mappings for sysdig system calls to support the Falco policy language.

### Changed

- Performance optimizations for golang APIs, including opflag and openflag map caching.

## [0.1.0] - 2020-10-30

### Added

- Implemented caching for opflags and openflags in golang APIs.

### Changed

- Refactored driver and plugin interface
- Refactored golang libraries to use constants for flags.

## [0.1.0-rc4] - 2020-08-10

### Added

- Added `node.id`, `node.ip`, `proc.entry`, and `schema` attributes to query language and export APIs.
- Added golang APIs.

### Changed

- Support for new Avro schema (version 2).
- Added missing EXIT opflag to Python APIs.
- Adding patch level to comply with semnatic versioning.

### Fixed

- Fixed open flags bitmaps.
- Fixed attribute name typo when computing proc and pproc duration.
- Fixed bug in provenance queries.

## [0.1-rc3] - 2020-03-17

### Added

- Added support for Pandas Dataframe conversion.
- Query language and support for filtering SysFlow records (Python).
- Added filter option for sysprint.
- Added SysFlow Jupyter notebook with sample notebooks and data science libraries.

### Changed

- Changed sysprint's base image to Red Hat UBI (ubi8/ubi).
- Updated option list for sysprint, with option renaming [breaking change]
- Refactored and improved JSON converters; new JSON schema [breaking change].
- Increased `sf-apis` version to the latest release candidate 0.1-rc3.

### Fixed

- Several bug fixes in formatting API.
- Proper handling of keyboard interrupts in sysprint.

## [0.1-rc2] - 2019-11-08

### Changed

- Increased `sf-apis` version to the latest release candidate 0.1-rc2.

## [0.1-rc1] - 2019-10-31

### Added

- First release candidate with basic set of SysFlow APIs (C++ and Python).

[Unreleased]: https://github.com/sysflow-telemetry/sf-apis/compare/0.5.0-rc1...HEAD
[0.5.0]: https://github.com/sysflow-telemetry/sf-apis/compare/0.4.3...0.5.0-rc1
[0.4.3]: https://github.com/sysflow-telemetry/sf-apis/compare/0.4.2...0.4.3
[0.4.2]: https://github.com/sysflow-telemetry/sf-apis/compare/0.4.1...0.4.2
[0.4.1]: https://github.com/sysflow-telemetry/sf-apis/compare/0.4.0...0.4.1
[0.4.0]: https://github.com/sysflow-telemetry/sf-apis/compare/0.3.1...0.4.0
[0.3.1]: https://github.com/sysflow-telemetry/sf-apis/compare/0.3.0...0.3.1
[0.3.0]: https://github.com/sysflow-telemetry/sf-apis/compare/0.2.2...0.3.0
[0.2.2]: https://github.com/sysflow-telemetry/sf-apis/compare/0.2.1...0.2.2
[0.2.1]: https://github.com/sysflow-telemetry/sf-apis/compare/0.2.0...0.2.1
[0.2.0]: https://github.com/sysflow-telemetry/sf-apis/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/sysflow-telemetry/sf-apis/compare/0.1.0-rc4...0.1.0
[0.1.0-rc4]: https://github.com/sysflow-telemetry/sf-apis/compare/0.1-rc3...0.1.0-rc4
[0.1-rc3]: https://github.com/sysflow-telemetry/sf-apis/compare/0.1-rc2...0.1-rc3
[0.1-rc2]: https://github.com/sysflow-telemetry/sf-apis/compare/0.1-rc1...0.1-rc2
[0.1-rc1]: https://github.com/sysflow-telemetry/sf-apis/releases/tag/0.1-rc1