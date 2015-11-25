package json2struct

import (
	"bytes"
	"encoding/json"
	"testing"
)

var basic = []byte(`{
	"foo": "fooer",
	"bar": "bars",
	"biz": 1,
	"baz": 42.1,
	"foo_bar": "frood"
}`)
var expectedBasicFields = "\tBar\tstring `json:\"bar\"`\n\tBaz\tfloat64 `json:\"baz\"`\n\tBiz\tint `json:\"biz\"`\n\tFoo\tstring `json:\"foo\"`\n\tFooBar\tstring `json:\"foo_bar\"`\n"
var expectedBasicStruct = "type Basic struct {\n\tBar\tstring `json:\"bar\"`\n\tBaz\tfloat64 `json:\"baz\"`\n\tBiz\tint `json:\"biz\"`\n\tFoo\tstring `json:\"foo\"`\n\tFooBar\tstring `json:\"foo_bar\"`\n}"

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

var intermediate = []byte(`{
	"name": "Marvin",
	"id": 42,
	"bot": true,
	"quotes": [
		"I think you ought to know I'm feeling very depressed",
		"Life? Don't talk to me about life!"
	],
	"ints": [
		1,
		2,
		3,
		4
	],
	"floats": [
		1.01,
		2.01,
		3.01
	],
	"bools": [
		true,
		false
	],
	"date": "Fri Jan 23 13:02:46 +0000 2015"
}`)
var expectedIntermediateFields = "\tBools\t[]bool `json:\"bools\"`\n\tBot\tbool `json:\"bot\"`\n\tDate\tstring `json:\"date\"`\n\tFloats\t[]float64 `json:\"floats\"`\n\tId\tint `json:\"id\"`\n\tInts\t[]int `json:\"ints\"`\n\tName\tstring `json:\"name\"`\n\tQuotes\t[]string `json:\"quotes\"`\n"
var expectedIntermediateStruct = "type Intermediate strict {\n\tBools\t[]bool `json:\"bools\"`\n\tBot\tbool `json:\"bot\"`\n\tDate\tstring `json:\"date\"`\n\tFloats\t[]float64 `json:\"floats\"\n\tId\tint `json:\"id\"`\n\tInts\t[]int `json:\"ints\"`\n\tName\tstring `json:\"name\"`\n\tQuotes\t[]string `json:\"quotes\"`\n}"

func TestIntermediateStruct(t *testing.T) {
 	var tst interface{}
	err := json.Unmarshal(intermediate, &tst)
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
	if buff.String() != expectedIntermediateFields {
		t.Errorf("expected:\n%q, got:\n%q", expectedIntermediateFields, buff.String())
	}
}
