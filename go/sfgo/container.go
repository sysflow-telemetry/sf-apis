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

type Container struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Image string `json:"image"`

	Imageid string `json:"imageid"`

	Type ContainerType `json:"type"`

	Privileged bool `json:"privileged"`

	PodId *UnionNullString `json:"podId"`
}

const ContainerAvroCRC64Fingerprint = "\xbav\xfc\f\x9bU\xc8\xcd"

func NewContainer() Container {
	r := Container{}
	r.PodId = NewUnionNullString()

	return r
}

func DeserializeContainer(r io.Reader) (Container, error) {
	t := NewContainer()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeContainerFromSchema(r io.Reader, schema string) (Container, error) {
	t := NewContainer()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeContainer(r Container, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Image, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Imageid, w)
	if err != nil {
		return err
	}
	err = writeContainerType(r.Type, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Privileged, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PodId, w)
	if err != nil {
		return err
	}
	return err
}

func (r Container) Serialize(w io.Writer) error {
	return writeContainer(r, w)
}

func (r Container) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"image\",\"type\":\"string\"},{\"name\":\"imageid\",\"type\":\"string\"},{\"name\":\"type\",\"type\":{\"name\":\"ContainerType\",\"namespace\":\"sysflow.type\",\"symbols\":[\"CT_DOCKER\",\"CT_LXC\",\"CT_LIBVIRT_LXC\",\"CT_MESOS\",\"CT_RKT\",\"CT_CUSTOM\",\"CT_CRI\",\"CT_CONTAINERD\",\"CT_CRIO\",\"CT_BPM\"],\"type\":\"enum\"}},{\"name\":\"privileged\",\"type\":\"boolean\"},{\"name\":\"podId\",\"type\":[\"null\",\"string\"]}],\"name\":\"sysflow.entity.Container\",\"type\":\"record\"}"
}

func (r Container) SchemaName() string {
	return "sysflow.entity.Container"
}

func (_ Container) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Container) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Container) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Container) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Container) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Container) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Container) SetString(v string)   { panic("Unsupported operation") }
func (_ Container) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Container) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Id}
	case 1:
		return &types.String{Target: &r.Name}
	case 2:
		return &types.String{Target: &r.Image}
	case 3:
		return &types.String{Target: &r.Imageid}
	case 4:
		return &ContainerTypeWrapper{Target: &r.Type}
	case 5:
		return &types.Boolean{Target: &r.Privileged}
	case 6:
		r.PodId = NewUnionNullString()

		return r.PodId
	}
	panic("Unknown field index")
}

func (r *Container) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Container) NullField(i int) {
	switch i {
	case 6:
		r.PodId = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Container) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Container) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Container) Finalize()                        {}

func (_ Container) AvroCRC64Fingerprint() []byte {
	return []byte(ContainerAvroCRC64Fingerprint)
}

func (r Container) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["image"], err = json.Marshal(r.Image)
	if err != nil {
		return nil, err
	}
	output["imageid"], err = json.Marshal(r.Imageid)
	if err != nil {
		return nil, err
	}
	output["type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["privileged"], err = json.Marshal(r.Privileged)
	if err != nil {
		return nil, err
	}
	output["podId"], err = json.Marshal(r.PodId)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Container) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["image"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Image); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for image")
	}
	val = func() json.RawMessage {
		if v, ok := fields["imageid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Imageid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for imageid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["privileged"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Privileged); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for privileged")
	}
	val = func() json.RawMessage {
		if v, ok := fields["podId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PodId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for podId")
	}
	return nil
}
