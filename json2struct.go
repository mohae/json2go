package json2struct

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"

	"github.com/mohae/firkin/queue"
)

// stringValues is a slice of reflect.Value holding *reflect.StringValue.
// It implements the methods to sort by string.
type stringValues []reflect.Value

func (sv stringValues) Len() int           { return len(sv) }
func (sv stringValues) Swap(i, j int)      { sv[i], sv[j] = sv[j], sv[i] }
func (sv stringValues) Less(i, j int) bool { return sv.get(i) < sv.get(j) }
func (sv stringValues) get(i int) string   { return sv[i].String() }

type structDef struct {
	name string
	val reflect.Value
	buff bytes.Buffer
}

func newStructDef(name string, val reflect.Value) structDef {
	s := structDef{name: name, val: val}
	s.buff.WriteString(fmt.Sprintf("type %s struct {\n", name))
	return s
}

func (s *structDef) Bytes() []byte {
	s.buff.WriteString("}\n")
	return s.buff.Bytes()
}

// Gen unmarshals JSON-encoded data and returns its struct definition(s) using
// the name as the struct's name.  If the JSON includes other maps, the field
// will be an embedded struct, with that struct's definition also being
// generated.  If an error occurs during unmarshalling of the data, it will
// be returned.  If an error occurs while writing to the buffer, that error
// will be returned.
func Gen(name string, data []byte) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("no name")
	}
	// unmarshal the JSON-encoded data
	var datum interface{}
	err := json.Unmarshal(data, &datum)
	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer
	var wg sync.WaitGroup
	q := queue.NewQ(2)
	result := make(chan []byte)
	// start the worker
	// send initial work item
	q.Enqueue(newStructDef(name, reflect.ValueOf(datum)))
	go func() {
		defineStruct(q, result, &wg)
	}()
	// collect the results until the resCh is closed
	var i int
	for {
		i++
		val, ok := <-result
		if !ok {
			break
		}
		// TODO handle error/short read
		buff.Write(val)
	}
	return buff.Bytes(), nil
}

func defineStruct(q *queue.Queue, result chan []byte, wg *sync.WaitGroup) {
	for {
		if q.IsEmpty() {
			break
		}
		tmp, ok := q.Dequeue()
		if !ok {
			break
		}
		s := tmp.(structDef)
		var sv stringValues = s.val.MapKeys()
		sort.Sort(sv)
		for _, key := range sv {
			k, tag := getFieldName(key)
			val := s.val.MapIndex(key)
			typ := getValueKind(val)
			// maps are embedded structs
			if typ == reflect.Map.String() {
				tmp := newStructDef(k, val.Elem())
				q.Enqueue(tmp)
				s.buff.WriteString(fmt.Sprintf("\t%s `json:%q`\n", k, tag))
				continue
			}
			s.buff.WriteString(fmt.Sprintf("\t%s %s `json:%q`\n", k, typ, tag))
		}
		result <- s.Bytes()
	}
	close(result)
}

func getValueKind(val reflect.Value) string {
	switch val.Elem().Type().Kind() {
	case reflect.Float64:
		v := val.Elem().Float()
		if v == float64(int64(v)) {
			return reflect.Int.String()
		}
		return reflect.Float64.String()
	case reflect.Slice:
		v := val.Elem().Index(0).Elem()
		if v.Type().Kind() == reflect.Float64 {
			vv := v.Float()
			if vv == float64(int64(vv)) {
				return fmt.Sprintf("[]%s", reflect.Int.String())
			}
			return fmt.Sprintf("[]%s", reflect.Float64.String())
		}
		return fmt.Sprintf("[]%s", v.Type().Kind().String())
	case reflect.Struct:
		fmt.Println("struct not handled")
		return ""

	}
	return val.Elem().Type().Kind().String()
}

func getFieldName(key reflect.Value) (name, tag string) {
	tag = key.String()
	vals := strings.Split(tag, "_")
	for _, v := range vals {
		name = fmt.Sprintf("%s%s", name, strings.Title(v))
	}
	return name, tag
}
