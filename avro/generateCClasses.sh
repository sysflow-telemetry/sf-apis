#!/bin/bash

if [ -z "$1" ]
then
    echo "Missing argument. Usage: ./geberateCClasses.sh <version>"
    exit
fi

source ./manifest

if ! command -v java &> /dev/null
then
    echo "java could not be found"
    exit
fi

wget -N -P avro-tools/ ${AVRO_TOOLS_URL}

if ! command -v avrogencpp &> /dev/null 
then
    sudo apt-get install -y build-essential libboost-all-dev libsnappy-dev
    wget -N -P avro-cpp/ ${AVRO_GENCPP_URL}
    cd avro-cpp && tar -xzf avro-cpp-${AVRO_VERSION}.tar.gz && cd avro-cpp-${AVRO_VERSION} && sudo ./build.sh install && cd ../..
fi

java -jar avro-tools/avro-tools-${AVRO_VERSION}.jar idl avdl/sysflow.avdl ./avpr/sysflow.avpr
java -jar avro-tools/avro-tools-${AVRO_VERSION}.jar idl2schemata ./avdl/sysflow.avdl avsc/

# cpp stub generation
avrogencpp -i ./avsc/SysFlow.avsc  -o ../c++/sysflow/sysflow.hh -n sysflow
echo "#ifndef __AVSC_SYSFLOW${1}" > ../c++/sysflow/avsc_sysflow${1}.hh
echo "#define __AVSC_SYSFLOW${1}" >> ../c++/sysflow/avsc_sysflow${1}.hh
echo "#include <string>" >> ../c++/sysflow/avsc_sysflow${1}.hh
echo -n "extern const std::string AVSC_SF = " >> ../c++/sysflow/avsc_sysflow${1}.hh
cat ./avsc/SysFlow.avsc | python3 -c 'import json,sys; print(json.dumps(sys.stdin.read()))' | tr -d '\n' >> ../c++/sysflow/avsc_sysflow${1}.hh
echo  ";" >> ../c++/sysflow/avsc_sysflow${1}.hh
echo "#endif" >> ../c++/sysflow/avsc_sysflow${1}.hh
