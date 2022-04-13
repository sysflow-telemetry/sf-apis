#!/bin/bash
java -jar avro-tools/avro-tools-1.11.0.jar idl avdl/sysflow.avdl ./avpr/sysflow.avpr
java -jar avro-tools/avro-tools-1.11.0.jar idl2schemata ./avdl/sysflow.avdl avsc/

# cpp stub generation
#avrogencpp -i ./avsc/ContainerID.avsc  -o ./src/sysflow/container_id.hh -n sysflow.type
#avrogencpp -i ./avsc/EventType.avsc  -o ./src/sysflow/event_type.hh -n sysflow.type
#avrogencpp -i ./avsc/OID.avsc  -o ./src/sysflow/oid.hh  -n sysflow.type
#avrogencpp -i ./avsc/Process.avsc  -o ./src/sysflow/process.hh  -n sysflow.entity
#avrogencpp -i ./avsc/File.avsc  -o ./src/sysflow/file.hh  -n sysflow.entity
#avrogencpp -i ./avsc/Container.avsc  -o ./src/sysflow/container.hh  -n sysflow.entity
#avrogencpp -i ./avsc/ProcessFlow.avsc  -o ./src/sysflow/proc_flow.hh -n sysflow.flow
#avrogencpp -i ./avsc/FileFlow.avsc  -o ./src/sysflow/file_flow.hh -n sysflow.flow
#avrogencpp -i ./avsc/NetworkFlow.avsc  -o ./src/sysflow/network_flow.hh -n sysflow.flow
#avrogencpp -i ./avsc/ActionType.avsc  -o ./src/sysflow/action_type.hh -n sysflow.type
avrogencpp -i ./avsc/SysFlow.avsc  -o ../c++/sysflow/sysflow.hh -n sysflow
echo "#ifndef __AVSC_SYSFLOW${1}" > ../c++/sysflow/avsc_sysflow${1}.hh
echo "#define __AVSC_SYSFLOW${1}" >> ../c++/sysflow/avsc_sysflow${1}.hh
echo "#include <string>" >> ../c++/sysflow/avsc_sysflow${1}.hh
echo -n "extern const std::string AVSC_SF = " >> ../c++/sysflow/avsc_sysflow${1}.hh
#AVSC=`cat avsc/sysflow${1}/TCCDMDatum.avsc`
#sed -e "s/\"/\\\\\"/gi" avsc/sysflow${1}/TCCDMDatum.avsc | tr -d '\n' >> ../c++/sysflow/avsc_sysflow${1}.hh
cat ./avsc/SysFlow.avsc | python3 -c 'import json,sys; print(json.dumps(sys.stdin.read()))' | tr -d '\n' >> ../c++/sysflow/avsc_sysflow${1}.hh
#printf "%q" $AVSC   >> ../c++/sysflow/avsc_sysflow${1}.hh
echo  ";" >> ../c++/sysflow/avsc_sysflow${1}.hh
echo "#endif" >> ../c++/sysflow/avsc_sysflow${1}.hh
