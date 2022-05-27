#  SysFlow APIs and Utilities

SysFlow uses [Apache Avro](https://avro.apache.org/) serialization to create compact records that can be processed by a wide variety of programming languages, and big data analytics platforms such as [Apache Spark](https://spark.apache.org/). Avro enables a user to generate programming stubs for serializing and deserializing data, using either [Apache Avro IDL](https://avro.apache.org/docs/1.9.1/idl.html) or [Apache schema files](https://avro.apache.org/docs/1.9.1/spec.html).  

## Cloning source

The sf-apis project has been tested primarily on Ubuntu 16.04 and 18.04.  The project will be tested on other flavors of UNIX in the future. This document describes how to build and run the application both on a linux host. 

To build the project, first pull down the source code:
```
git clone git@github.com:sysflow-telemetry/sf-apis.git 
```

## Avro IDL and schema files

The Avro IDL files for SysFlow are available in the repository under `sf-apis/avro/avdl`, while the schema files are available under `sf-apis/avro/avsc`.   The `avrogen` tool can be used to generate classes using the schema.  See `sf-apis/avro/generateCClasses.sh` for an example of how to generate C++ headers from apache schema files.  

##  SysFlow Avro C++

SysFlow C++ SysFlow objects and encoders/decoders are all available in `sf-apis/c++/sysflow/sysflow.hh`.  `sf-collector/src/sysreader.cpp` provides a good example of how to read and process different SysFlow avro objects in C++.   Note that one must install [Apache Avro 1.9.1 cpp](https://avro.apache.org/releases.html) to run an application that includes `sysflow.hh`.  The library file `-lavrocpp` must also be linked during compilation. 

## SysFlow Avro Python 3

SysFlow Python 3 APIs are generated with the avro-gen Python package. These classes are available in `sf-apis/py3`.

In order to install the SysFlow Python package:

```
cd sf-apis/py3
sudo python3 setup.py install
```

Please see the SysFlow Python API reference documents for more information on the modules and objects in the library.

## SysFlow utilities

### sysprint

`sysprint` is a tool written using the SysFlow Python API that will print out SysFlow traces from a file into several different formats including JSON, CSV, and tabular pretty print form.  Not only will sysprint help you interact with SysFlow, it is also a good example for how to write new analytics tools using the SysFlow API.   

```
usage: sysprint [-h] [-i {local,s3}] [-o {str,flatjson,json,csv}] [-w FILE]
                [-c FIELDS] [-f FILTER] [-l] [-e S3ENDPOINT] [-p S3PORT]
                [-a S3ACCESSKEY] [-s S3SECRETKEY] [-k] [-A]
                [--secure [SECURE]]
                path [path ...]

sysprint: a human-readable printer and format converter for Sysflow captures.

positional arguments:
  path                  list of paths or bucket names from where to read trace
                        files

optional arguments:
  -h, --help            show this help message and exit
  -i {local,s3}, --input {local,s3}
                        input type
  -o {str,flatjson,json,csv}, --output {str,flatjson,json,csv}
                        output format
  -w FILE, --file FILE  output file path
  -c FIELDS, --fields FIELDS
                        comma-separated list of sysflow fields to be printed
  -f FILTER, --filter FILTER
                        filter expression
  -l, --list            list available record attributes
  -e S3ENDPOINT, --s3endpoint S3ENDPOINT
                        s3 server address from where to read sysflows
  -p S3PORT, --s3port S3PORT
                        s3 server port
  -a S3ACCESSKEY, --s3accesskey S3ACCESSKEY
                        s3 access key
  -s S3SECRETKEY, --s3secretkey S3SECRETKEY
                        s3 secret key
  -k, --k8s             add pod related fields to output
  -A, --allfields       add all available fields to output
  --secure [SECURE]     indicates if SSL connection
```
