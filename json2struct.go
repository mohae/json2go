package json2struct

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// stringValues is a slice of reflect.Value holding *reflect.StringValue.
// It implements the methods to sort by string.
type stringValues []reflect.Value

func (sv stringValues) Len() int           { return len(sv) }
func (sv stringValues) Swap(i, j int)      { sv[i], sv[j] = sv[j], sv[i] }
func (sv stringValues) Less(i, j int) bool { return sv.get(i) < sv.get(j) }
func (sv stringValues) get(i int) string   { return sv[i].String() }

// Gen accepts an interface and
//func Gen

// gen generates a byte representation of the interface{} fields, which is
// assumed to be the result of a json.Unmarshal,
func gen(data interface{}, buff *bytes.Buffer) error {
	fmt.Printf("%#v\n", data)
	// TODO should pointers to interface be supported or generate an error?
	var datum reflect.Value
	switch data.(type) {
	case map[string]interface{}:
		datum = reflect.ValueOf(data)
	default:
		return fmt.Errorf("%q not supported", reflect.TypeOf(data).Kind())
	}
	toStruct(datum, buff)
	return nil
}

// takes a map[string]interface and a buffer, and populates the buffer
// with the struct def.
func toStruct(datum reflect.Value, buff *bytes.Buffer) error {
	var sv stringValues = datum.MapKeys()
	sort.Sort(sv)
	for _, key := range sv {
		k, tag := getFieldName(key)
		_, err := buff.WriteString(fmt.Sprintf("\t%s\t", k))
		if err != nil {
			return err
		}
		val := datum.MapIndex(key)
		typ := getValueKind(val)
		buff.WriteString(typ.String())
		buff.WriteString(fmt.Sprintf(" `json:%q`", tag))
		buff.WriteString("\n")
	}
	return nil
}

func getValueKind(val reflect.Value) reflect.Kind {
	if val.Elem().Type().Kind() == reflect.Float64 {
		v := val.Elem().Float()
		if v == float64(int64(v)) {
			return reflect.Int
		}
		return reflect.Float64
	}
	return val.Elem().Type().Kind()
}

func getFieldName(key reflect.Value) (name, tag string) {
	tag = key.String()
	vals := strings.Split(tag, "_")
	for _, v := range vals {
		name = fmt.Sprintf("%s%s", name, strings.Title(v))
	}
	return name, tag
}
