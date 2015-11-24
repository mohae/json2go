package json2struct

import (
	"bytes"
	"encoding/json"
	"testing"
)

var basic = []byte(`{"foo": "fooer", "bar": "bars", "biz": 1, "baz": 42.1, "foo_bar": "frood"}`)
var expectedBasicFields = "\tBar\tstring `json:\"bar\"`\n\tBaz\tfloat64 `json:\"baz\"`\n\tBiz\tint `json:\"biz\"`\n\tFoo\tstring `json:\"foo\"`\n\tFooBar\tstring `json:\"foo_bar\"`\n"

func TestBasicStruct(t *testing.T) {
 	var tst interface{}
	err := json.Unmarshal(basic, &tst)
	if err != nil {
		t.Errorf("unexpected error unmarshaling basic test data: %q", err)
		return
	}
	var buff bytes.Buffer
	err = gen(tst, &buff)
	if err != nil {
		t.Errorf("unexpected error generating struct: %q", err)
		return
	}
	if buff.String() != expectedBasicFields {
		t.Errorf("expected:\n%q, got:\n%q", expectedBasicFields, buff.String())
	}
}
