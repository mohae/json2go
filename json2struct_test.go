package json2struct

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

var basic = []byte(`{
	"foo": "fooer",
	"bar": "bars",
	"biz": 1,
	"baz": 42.1,
	"foo_bar": "frood"
}`)
var expectedBasic = "type Basic struct {\n\tBar string `json:\"bar\"`\n\tBaz float64 `json:\"baz\"`\n\tBiz int `json:\"biz\"`\n\tFoo string `json:\"foo\"`\n\tFooBar string `json:\"foo_bar\"`\n}\n\n"
var expectedFmtBasic = "type Basic struct {\n\tBar    string  `json:\"bar\"`\n\tBaz    float64 `json:\"baz\"`\n\tBiz    int     `json:\"biz\"`\n\tFoo    string  `json:\"foo\"`\n\tFooBar string  `json:\"foo_bar\"`\n}\n"

func TestBasicStruct(t *testing.T) {
	def, err := Gen("basic", basic)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(def) != expectedBasic {
		t.Errorf("expected %q got %q", expectedBasic, string(def))
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
var expectedIntermediate = "type Intermediate struct {\n\tBools []bool `json:\"bools\"`\n\tBot bool `json:\"bot\"`\n\tDate string `json:\"date\"`\n\tFloats []float64 `json:\"floats\"`\n\tID int `json:\"id\"`\n\tInts []int `json:\"ints\"`\n\tName string `json:\"name\"`\n\tQuotes []string `json:\"quotes\"`\n}\n\n"

func TestIntermediateStruct(t *testing.T) {
	def, err := Gen("Intermediate", intermediate)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(def) != expectedIntermediate {
		t.Errorf("expected %q got %q", expectedIntermediate, string(def))
	}
}

// widget example from json.org/example.html
var widget = []byte(`{
	"widget": {
		"debug": "on",
		"window": {
			"title": "Sample Konfabulator Widget",
        		"name": "main_window",
        		"width": 500,
        		"height": 500
		},
		"image": {
        		"src": "Images/Sun.png",
        		"name": "sun1",
        		"hOffset": 250,
        		"vOffset": 250,
        		"alignment": "center"
		},
		"text": {
			"data": "Click Here",
        		"size": 36,
        		"style": "bold",
        		"name": "text1",
        		"hOffset": 250,
        		"vOffset": 100,
        		"alignment": "center",
        		"onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
		}
	}
}`)

var expectedWidget = "type TestW struct {\n\tWidget `json:\"widget\"`\n}\n\ntype Widget struct {\n\tDebug string `json:\"debug\"`\n\tImage `json:\"image\"`\n\tText `json:\"text\"`\n\tWindow `json:\"window\"`\n}\n\ntype Image struct {\n\tAlignment string `json:\"alignment\"`\n\tHOffset int `json:\"hOffset\"`\n\tName string `json:\"name\"`\n\tSrc string `json:\"src\"`\n\tVOffset int `json:\"vOffset\"`\n}\n\ntype Text struct {\n\tAlignment string `json:\"alignment\"`\n\tData string `json:\"data\"`\n\tHOffset int `json:\"hOffset\"`\n\tName string `json:\"name\"`\n\tOnMouseUp string `json:\"onMouseUp\"`\n\tSize int `json:\"size\"`\n\tStyle string `json:\"style\"`\n\tVOffset int `json:\"vOffset\"`\n}\n\ntype Window struct {\n\tHeight int `json:\"height\"`\n\tName string `json:\"name\"`\n\tTitle string `json:\"title\"`\n\tWidth int `json:\"width\"`\n}\n\n"

func TestWidget(t *testing.T) {
	s, err := Gen("TestW", widget)
	if err != nil {
		t.Errorf("unexpected error %q", err)
	}
	if string(s) != expectedWidget {
		t.Errorf("expected:\n%q, got:\n%q", expectedWidget, string(s))
	}
}

var wnull = []byte(`{
	"foo": "fooer",
	"bar": null,
	"biz": 1,
	"baz": 42.1,
	"foo_bar": "frood"
}`)

var expectedWNull = "type WNull struct {\n\tBar interface{} `json:\"bar\"`\n\tBaz float64 `json:\"baz\"`\n\tBiz int `json:\"biz\"`\n\tFoo string `json:\"foo\"`\n\tFooBar string `json:\"foo_bar\"`\n}\n\n"

func TestWNull(t *testing.T) {
	def, err := Gen("WNull", wnull)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(def) != expectedWNull {
		t.Errorf("expected %q got %q", expectedWNull, string(def))
	}
}

var basicArr = []byte(`[{
	"foo": "fooer",
	"bar": "bars",
	"biz_id": 1,
	"baz": 42.1,
	"foo_bar": "frood"
}]`)

var expectedBasicArr = "type BasicArr struct {\n\tBar string `json:\"bar\"`\n\tBaz float64 `json:\"baz\"`\n\tBizID int `json:\"biz_id\"`\n\tFoo string `json:\"foo\"`\n\tFooBar string `json:\"foo_bar\"`\n}\n\n"

func TestBasicArr(t *testing.T) {
	def, err := Gen("BasicArr", basicArr)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(def) != expectedBasicArr {
		t.Errorf("expected %q got %q", expectedBasicArr, string(def))
	}
}

var sliceMap = []byte(`{
	"foo": [
		{
			"bar": "biz",
			"foo_bar": "frood"
		},
		{
			"bar": "baz",
			"foo_bar": "hoopy"
		}
	]
}`)

var expectedSliceMap = "type SliceMap struct {\n\tFoos []Foo `json:\"foo\"`\n}\n\ntype Foo struct {\n\tBar string `json:\"bar\"`\n\tFooBar string `json:\"foo_bar\"`\n}\n\n"

func TestArrays(t *testing.T) {
	def, err := Gen("SliceMap", sliceMap)
	if err != nil {
		t.Errorf("unexpected error: %s", sliceMap)
	}
	if string(def) != expectedSliceMap {
		t.Errorf("expected %q got %q", expectedSliceMap, string(def))
	}
}

var mapType = []byte(`{
  "example.com": {
        "name": "example.com",
        "type": "SOA",
        "ttl":  300,
        "content": "ns1.example.com. hostmaster.example.com. 1299682996 300 1800 604800 300"
    }
}`)

var expectedMapTypeStruct = "type Zone map[string]Struct\n\ntype Struct struct {\n\tContent string `json:\"content\"`\n\tName string `json:\"name\"`\n\tTTL int `json:\"ttl\"`\n\tType string `json:\"type\"`\n}\n\n"
var expectedFmtMapTypeStruct = "type Zone map[string]Struct\n\ntype Struct struct {\n\tContent string `json:\"content\"`\n\tName    string `json:\"name\"`\n\tTTL     int    `json:\"ttl\"`\n\tType    string `json:\"type\"`\n}\n"
var expectedMapTypeDomain = "type Zone map[string]Domain\n\ntype Domain struct {\n\tContent string `json:\"content\"`\n\tName string `json:\"name\"`\n\tTTL int `json:\"ttl\"`\n\tType string `json:\"type\"`\n}\n\n"
var expectedFmtMapTypeDomain = "type Zone map[string]Domain\n\ntype Domain struct {\n\tContent string `json:\"content\"`\n\tName    string `json:\"name\"`\n\tTTL     int    `json:\"ttl\"`\n\tType    string `json:\"type\"`\n}\n"
func TestMapType(t *testing.T) {
	def, err := GenMapType("Zone", "", mapType)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(def) != expectedMapTypeStruct {
		t.Errorf("expected %q got %q", expectedMapTypeStruct, string(def))
	}

	def, err = GenMapType("Zone", "domain", mapType)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(def) != expectedMapTypeDomain {
		t.Errorf("expected %q got %q", expectedMapTypeDomain, string(def))
	}
}

var mapSliceType = []byte(`{
  "example.com": [
    {
        "name": "example.com",
        "type": "SOA",
        "ttl":  300,
        "content": "ns1.example.com. hostmaster.example.com. 1299682996 300 1800 604800 300"
    }
  ]
}`)

var expectedMapSliceTypeStruct = "type Zone map[string][]Struct\n\ntype Struct struct {\n\tContent string `json:\"content\"`\n\tName string `json:\"name\"`\n\tTTL int `json:\"ttl\"`\n\tType string `json:\"type\"`\n}\n\n"
var expectedMapSliceTypeDomain = "type Zone map[string][]Domain\n\ntype Domain struct {\n\tContent string `json:\"content\"`\n\tName string `json:\"name\"`\n\tTTL int `json:\"ttl\"`\n\tType string `json:\"type\"`\n}\n\n"
func TestMapSliceType(t *testing.T) {
	def, err := GenMapType("Zone", "", mapSliceType)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(def) != expectedMapSliceTypeStruct {
		t.Errorf("expected %q got %q", expectedMapSliceTypeStruct, string(def))
	}

	def, err = GenMapType("Zone", "domain", mapSliceType)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(def) != expectedMapSliceTypeDomain {
		t.Errorf("expected %q got %q", expectedMapSliceTypeDomain, string(def))
	}
}

func TestTransmogrify(t *testing.T) {
	tests := []struct {
		pkg        string
		name string
		structName string
		isMap bool
		importJSON bool
		json	[]byte
		expected   string
	}{
		{"", "Basic", "", false, false, basic, fmt.Sprintf("package main\n\n%s", expectedFmtBasic)},
		{"test", "Basic", "", false, false, basic, fmt.Sprintf("package test\n\n%s", expectedFmtBasic)},
		{"", "Basic", "", false, true, basic, fmt.Sprintf("package main\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedFmtBasic)},
		{"test", "Basic", "", false, true, basic, fmt.Sprintf("package test\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedFmtBasic)},
		{"", "zone", "", true, false, mapType, fmt.Sprintf("package main\n\n%s", expectedFmtMapTypeStruct)},
		{"test", "zone", "", true, false, mapType, fmt.Sprintf("package test\n\n%s", expectedFmtMapTypeStruct)},
		{"", "zone", "", true, true, mapType, fmt.Sprintf("package main\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedFmtMapTypeStruct)},
		{"test", "zone", "", true, true, mapType, fmt.Sprintf("package test\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedFmtMapTypeStruct)},
		{"", "zone", "domain", true, false, mapType, fmt.Sprintf("package main\n\n%s", expectedFmtMapTypeDomain)},
		{"test", "zone", "domain", true, false, mapType, fmt.Sprintf("package test\n\n%s", expectedFmtMapTypeDomain)},
		{"", "zone", "domain", true, true, mapType, fmt.Sprintf("package main\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedFmtMapTypeDomain)},
		{"test", "zone", "domain", true, true, mapType, fmt.Sprintf("package test\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedFmtMapTypeDomain)},

	}
	var b bytes.Buffer
	for i, test := range tests {
		// create reader
		r := bytes.NewReader(test.json)
		// create Writer
		w := bufio.NewWriter(&b)

		calvin := NewTransmogrifier(test.name, r, w)
		if test.pkg != "" {
			calvin.SetPkg(test.pkg)
		}
		calvin.SetImportJSON(test.importJSON)
		if test.isMap {
			calvin.SetIsMap(test.isMap)
			calvin.SetStructName(test.structName)
		}
		err := calvin.Gen()
		if err != nil {
			t.Errorf("%d: unexpected error %q", i, err)
			continue
		}
		w.Flush()
		if b.String() != test.expected {
			t.Errorf("%d: expected %q, got %q", i, test.expected, b.String())
		}
		b.Reset()
	}
}

func TestNumToAlpha(t *testing.T) {
	tests := []struct {
		char     rune
		expected string
	}{
		{'0', "Zero"},
		{'1', "One"},
		{'2', "Two"},
		{'3', "Three"},
		{'4', "Four"},
		{'5', "Five"},
		{'6', "Six"},
		{'7', "Seven"},
		{'8', "Eight"},
		{'9', "Nine"},
	}
	for i, test := range tests {
		s := numToAlpha(test.char)
		if s != test.expected {
			t.Errorf("%d: expected %s got %s", i, test.expected, s)
		}
	}
}

func TestShouldDiscard(t *testing.T) {
	tests := []struct {
		char     rune
		expected bool
	}{
		{'~', true},
		{'!', true},
		{'@', true},
		{'#', true},
		{'$', true},
		{'%', true},
		{'^', true},
		{'&', true},
		{'-', true},
		{'_', true},
		{'*', true},
		{'=', true},
		{'+', true},
		{':', true},
		{'.', true},
		{'<', true},
		{'>', true},
		{'a', false},
		{'z', false},
		{'你', false},
	}
	for i, test := range tests {
		b := shouldDiscard(test.char)
		if b != test.expected {
			t.Errorf("%d: expected %t, got %t", i, test.expected, b)
		}
	}
}

func TestCleanFieldName(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"abdcn", "Abdcn"},
		{">asb", "Asb"},
		{"_asdf", "Asdf"},
		{"<>field", "Field"},
		{"日本語a", "日本語a"},
		{"бsdf", "Бsdf"},
		{"6dsaf", "Sixdsaf"},
	}
	for i, test := range tests {
		s := cleanFieldName(test.name)
		if s != test.expected {
			t.Errorf("%d: expected %q, got %q", i, test.expected, s)
		}
	}
}
