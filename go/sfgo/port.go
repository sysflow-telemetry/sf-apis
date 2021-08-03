// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     SysFlow.avsc
 */
package sfgo

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
)

var _ = fmt.Printf

type Port struct {
	Port int32 `json:"port"`

	TargetPort int32 `json:"targetPort"`

	NodePort int32 `json:"nodePort"`

	Proto string `json:"proto"`
}

const PortAvroCRC64Fingerprint = "\nY\x88\xe2\xe7\xed\xef`"

func NewPort() Port {
	r := Port{}
	return r
}

func DeserializePort(r io.Reader) (Port, error) {
	t := NewPort()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePortFromSchema(r io.Reader, schema string) (Port, error) {
	t := NewPort()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePort(r Port, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Port, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.TargetPort, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.NodePort, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Proto, w)
	if err != nil {
		return err
	}
	return err
}

func (r Port) Serialize(w io.Writer) error {
	return writePort(r, w)
}

func (r Port) Schema() string {
	return "{\"fields\":[{\"name\":\"port\",\"type\":\"int\"},{\"name\":\"targetPort\",\"type\":\"int\"},{\"name\":\"nodePort\",\"type\":\"int\"},{\"name\":\"proto\",\"type\":\"string\"}],\"name\":\"sysflow.entity.Port\",\"type\":\"record\"}"
}

func (r Port) SchemaName() string {
	return "sysflow.entity.Port"
}

func (_ Port) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Port) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Port) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Port) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Port) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Port) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Port) SetString(v string)   { panic("Unsupported operation") }
func (_ Port) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Port) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Int{Target: &r.Port}
	case 1:
		return &types.Int{Target: &r.TargetPort}
	case 2:
		return &types.Int{Target: &r.NodePort}
	case 3:
		return &types.String{Target: &r.Proto}
	}
	panic("Unknown field index")
}

func (r *Port) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Port) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Port) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Port) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Port) Finalize()                        {}

func (_ Port) AvroCRC64Fingerprint() []byte {
	return []byte(PortAvroCRC64Fingerprint)
}

func (r Port) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["port"], err = json.Marshal(r.Port)
	if err != nil {
		return nil, err
	}
	output["targetPort"], err = json.Marshal(r.TargetPort)
	if err != nil {
		return nil, err
	}
	output["nodePort"], err = json.Marshal(r.NodePort)
	if err != nil {
		return nil, err
	}
	output["proto"], err = json.Marshal(r.Proto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Port) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["port"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Port); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for port")
	}
	val = func() json.RawMessage {
		if v, ok := fields["targetPort"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TargetPort); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for targetPort")
	}
	val = func() json.RawMessage {
		if v, ok := fields["nodePort"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NodePort); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nodePort")
	}
	val = func() json.RawMessage {
		if v, ok := fields["proto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Proto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for proto")
	}
	return nil
}
