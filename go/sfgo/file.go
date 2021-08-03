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

type File struct {
	State SFObjectState `json:"state"`

	Oid FOID `json:"oid"`

	Ts int64 `json:"ts"`

	Restype int32 `json:"restype"`

	Path string `json:"path"`

	ContainerId *UnionNullString `json:"containerId"`
}

const FileAvroCRC64Fingerprint = "g(EKٶ2\x1a"

func NewFile() File {
	r := File{}
	r.ContainerId = NewUnionNullString()

	return r
}

func DeserializeFile(r io.Reader) (File, error) {
	t := NewFile()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFileFromSchema(r io.Reader, schema string) (File, error) {
	t := NewFile()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFile(r File, w io.Writer) error {
	var err error
	err = writeSFObjectState(r.State, w)
	if err != nil {
		return err
	}
	err = writeFOID(r.Oid, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Ts, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Restype, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Path, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContainerId, w)
	if err != nil {
		return err
	}
	return err
}

func (r File) Serialize(w io.Writer) error {
	return writeFile(r, w)
}

func (r File) Schema() string {
	return "{\"fields\":[{\"name\":\"state\",\"type\":{\"name\":\"SFObjectState\",\"namespace\":\"sysflow.type\",\"symbols\":[\"CREATED\",\"MODIFIED\",\"REUP\"],\"type\":\"enum\"}},{\"name\":\"oid\",\"type\":{\"name\":\"FOID\",\"namespace\":\"sysflow.type\",\"size\":20,\"type\":\"fixed\"}},{\"name\":\"ts\",\"type\":\"long\"},{\"name\":\"restype\",\"type\":\"int\"},{\"name\":\"path\",\"type\":\"string\"},{\"name\":\"containerId\",\"type\":[\"null\",\"string\"]}],\"name\":\"sysflow.entity.File\",\"type\":\"record\"}"
}

func (r File) SchemaName() string {
	return "sysflow.entity.File"
}

func (_ File) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ File) SetInt(v int32)       { panic("Unsupported operation") }
func (_ File) SetLong(v int64)      { panic("Unsupported operation") }
func (_ File) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ File) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ File) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ File) SetString(v string)   { panic("Unsupported operation") }
func (_ File) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *File) Get(i int) types.Field {
	switch i {
	case 0:
		return &SFObjectStateWrapper{Target: &r.State}
	case 1:
		return &FOIDWrapper{Target: &r.Oid}
	case 2:
		return &types.Long{Target: &r.Ts}
	case 3:
		return &types.Int{Target: &r.Restype}
	case 4:
		return &types.String{Target: &r.Path}
	case 5:
		r.ContainerId = NewUnionNullString()

		return r.ContainerId
	}
	panic("Unknown field index")
}

func (r *File) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *File) NullField(i int) {
	switch i {
	case 5:
		r.ContainerId = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ File) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ File) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ File) Finalize()                        {}

func (_ File) AvroCRC64Fingerprint() []byte {
	return []byte(FileAvroCRC64Fingerprint)
}

func (r File) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["state"], err = json.Marshal(r.State)
	if err != nil {
		return nil, err
	}
	output["oid"], err = json.Marshal(r.Oid)
	if err != nil {
		return nil, err
	}
	output["ts"], err = json.Marshal(r.Ts)
	if err != nil {
		return nil, err
	}
	output["restype"], err = json.Marshal(r.Restype)
	if err != nil {
		return nil, err
	}
	output["path"], err = json.Marshal(r.Path)
	if err != nil {
		return nil, err
	}
	output["containerId"], err = json.Marshal(r.ContainerId)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *File) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["state"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.State); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for state")
	}
	val = func() json.RawMessage {
		if v, ok := fields["oid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Oid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for oid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ts"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ts); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ts")
	}
	val = func() json.RawMessage {
		if v, ok := fields["restype"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Restype); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for restype")
	}
	val = func() json.RawMessage {
		if v, ok := fields["path"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Path); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for path")
	}
	val = func() json.RawMessage {
		if v, ok := fields["containerId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContainerId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for containerId")
	}
	return nil
}
