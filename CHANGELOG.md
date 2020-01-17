# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/).

> **Types of changes:**
>
> -   **Added**: for new features.
> -   **Changed**: for changes in existing functionality.
> -   **Deprecated**: for soon-to-be removed features.
> -   **Removed**: for now removed features.
> -   **Fixed**: for any bug fixes.
> -   **Security**: in case of vulnerabilities.

[UNRELEASED](https://github.com/sysflow-telemetry/sf-apis/compare/0.1-rc3...HEAD)
====================================================================

[0.1-rc3](https://github.com/sysflow-telemetry/sf-apis/compare/0.1-rc2...0.1-rc3) - 2020-01-17
===============================================================================

Added
-------

- Added support for Pandas Dataframe conversion.
- Query language and support for filtering SysFlow records (Python).

Changed
-------

- Changed sysprint's base image to ubi8/python36.
- Refactored and improved JSON converters; new JSON schema [breaking change].
- Several bug fixes in formatting API.
- Increased `sf-apis` version to the latest release candidate 0.1-rc2.


[0.1-rc2](https://github.com/sysflow-telemetry/sf-apis/compare/0.1-rc1...0.1-rc2) - 2019-11-08
===============================================================================

Changed
-------

- Increased `sf-apis` version to the latest release candidate 0.1-rc2.

0.1-rc1 - 2019-10-31
===============================================================================

Added
-------

- First release candidate with basic set of SysFlow APIs (C++ and Python).
