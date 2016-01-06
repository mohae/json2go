package json2go

import (
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
var expectedBasic = "type Basic struct {\n\tBar    string  `json:\"bar\"`\n\tBaz    float64 `json:\"baz\"`\n\tBiz    int     `json:\"biz\"`\n\tFoo    string  `json:\"foo\"`\n\tFooBar string  `json:\"foo_bar\"`\n}\n"
var expectedBasicPkg = fmt.Sprintf("package main\n\n%s", expectedBasic)

func TestBasicStruct(t *testing.T) {
	// create reader
	r := bytes.NewReader(basic)
	var buff bytes.Buffer
	calvin := NewTransmogrifier("basic", r, &buff)
	err := calvin.Gen()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if buff.String() != expectedBasicPkg {
		t.Errorf("expected %q got %q", expectedBasicPkg, buff.String())
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
var expectedIntermediate = "type Intermediate struct {\n\tBools  []bool    `json:\"bools\"`\n\tBot    bool      `json:\"bot\"`\n\tDate   string    `json:\"date\"`\n\tFloats []float64 `json:\"floats\"`\n\tID     int       `json:\"id\"`\n\tInts   []int     `json:\"ints\"`\n\tName   string    `json:\"name\"`\n\tQuotes []string  `json:\"quotes\"`\n}\n"
var expectedIntermediatePkg = fmt.Sprintf("package main\n\n%s", expectedIntermediate)

func TestIntermediateStruct(t *testing.T) {
	// create reader
	r := bytes.NewReader(intermediate)
	var buff bytes.Buffer
	calvin := NewTransmogrifier("intermediate", r, &buff)
	err := calvin.Gen()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if buff.String() != expectedIntermediatePkg {
		t.Errorf("expected %q got %q", expectedIntermediatePkg, buff.String())
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

var expectedWidget = "type TestW struct {\n\tWidget `json:\"widget\"`\n}\n\ntype Widget struct {\n\tDebug  string `json:\"debug\"`\n\tImage  `json:\"image\"`\n\tText   `json:\"text\"`\n\tWindow `json:\"window\"`\n}\n\ntype Image struct {\n\tAlignment string `json:\"alignment\"`\n\tHOffset   int    `json:\"hOffset\"`\n\tName      string `json:\"name\"`\n\tSrc       string `json:\"src\"`\n\tVOffset   int    `json:\"vOffset\"`\n}\n\ntype Text struct {\n\tAlignment string `json:\"alignment\"`\n\tData      string `json:\"data\"`\n\tHOffset   int    `json:\"hOffset\"`\n\tName      string `json:\"name\"`\n\tOnMouseUp string `json:\"onMouseUp\"`\n\tSize      int    `json:\"size\"`\n\tStyle     string `json:\"style\"`\n\tVOffset   int    `json:\"vOffset\"`\n}\n\ntype Window struct {\n\tHeight int    `json:\"height\"`\n\tName   string `json:\"name\"`\n\tTitle  string `json:\"title\"`\n\tWidth  int    `json:\"width\"`\n}\n"
var expectedWidgetPkg = fmt.Sprintf("package main\n\n%s", expectedWidget)

func TestWidget(t *testing.T) {
	// create reader
	r := bytes.NewReader(widget)
	var buff bytes.Buffer
	calvin := NewTransmogrifier("TestW", r, &buff)
	err := calvin.Gen()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if buff.String() != expectedWidgetPkg {
		t.Errorf("expected %q got %q", expectedWidgetPkg, buff.String())
	}
}

var wnull = []byte(`{
	"foo": "fooer",
	"bar": null,
	"biz": 1,
	"baz": 42.1,
	"foo_bar": "frood"
}`)

var expectedWNull = "type WNull struct {\n\tBar    interface{} `json:\"bar\"`\n\tBaz    float64     `json:\"baz\"`\n\tBiz    int         `json:\"biz\"`\n\tFoo    string      `json:\"foo\"`\n\tFooBar string      `json:\"foo_bar\"`\n}\n"
var expectedWNullPkg = fmt.Sprintf("package main\n\n%s", expectedWNull)

func TestWNull(t *testing.T) {
	// create reader
	r := bytes.NewReader(wnull)
	var buff bytes.Buffer
	calvin := NewTransmogrifier("WNull", r, &buff)
	err := calvin.Gen()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if buff.String() != expectedWNullPkg {
		t.Errorf("expected %q got %q", expectedWNullPkg, buff.String())
	}
}

var basicArr = []byte(`[{
	"foo": "fooer",
	"bar": "bars",
	"biz_id": 1,
	"baz": 42.1,
	"foo_bar": "frood"
}]`)

var expectedBasicArr = "type BasicArr struct {\n\tBar    string  `json:\"bar\"`\n\tBaz    float64 `json:\"baz\"`\n\tBizID  int     `json:\"biz_id\"`\n\tFoo    string  `json:\"foo\"`\n\tFooBar string  `json:\"foo_bar\"`\n}\n"
var expectedBasicArrPkg = fmt.Sprintf("package main\n\n%s", expectedBasicArr)

func TestBasicArr(t *testing.T) {
	// create reader
	r := bytes.NewReader(basicArr)
	var buff bytes.Buffer
	calvin := NewTransmogrifier("BasicArr", r, &buff)
	err := calvin.Gen()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if buff.String() != expectedBasicArrPkg {
		t.Errorf("expected %q got %q", expectedBasicArrPkg, buff.String())
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

var expectedSliceMap = "type SliceMap struct {\n\tFoos []Foo `json:\"foo\"`\n}\n\ntype Foo struct {\n\tBar    string `json:\"bar\"`\n\tFooBar string `json:\"foo_bar\"`\n}\n"
var expectedSliceMapPkg = fmt.Sprintf("package main\n\n%s", expectedSliceMap)

func TestArrays(t *testing.T) {
	// create reader
	r := bytes.NewReader(sliceMap)
	var buff bytes.Buffer
	calvin := NewTransmogrifier("SliceMap", r, &buff)
	err := calvin.Gen()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if buff.String() != expectedSliceMapPkg {
		t.Errorf("expected %q got %q", expectedSliceMapPkg, buff.String())
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

var expectedMapTypeStruct = "type Zone map[string]Struct\n\ntype Struct struct {\n\tContent string `json:\"content\"`\n\tName    string `json:\"name\"`\n\tTTL     int    `json:\"ttl\"`\n\tType    string `json:\"type\"`\n}\n"
var expectedMapTypeStructPkg = fmt.Sprintf("package main\n\n%s", expectedMapTypeStruct)
var expectedMapTypeDomain = "type Zone map[string]Domain\n\ntype Domain struct {\n\tContent string `json:\"content\"`\n\tName    string `json:\"name\"`\n\tTTL     int    `json:\"ttl\"`\n\tType    string `json:\"type\"`\n}\n"
var expectedMapTypeDomainPkg = fmt.Sprintf("package main\n\n%s", expectedMapTypeDomain)

func TestMapType(t *testing.T) {
	// create reader
	tests := []struct {
		structName string
		expected   string
	}{
		{"", expectedMapTypeStructPkg},
		{"domain", expectedMapTypeDomainPkg},
	}
	for _, test := range tests {
		r := bytes.NewReader(mapType)
		var buff bytes.Buffer
		calvin := NewTransmogrifier("Zone", r, &buff)
		calvin.MapType = true
		if test.structName != "" {
			calvin.SetStructName(test.structName)
		}
		err := calvin.Gen()
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if buff.String() != test.expected {
			t.Errorf("expected %q got %q", test.expected, buff.String())
		}
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
var expectedMapSliceTypeStruct = "type Zone map[string][]Struct\n\ntype Struct struct {\n\tContent string `json:\"content\"`\n\tName    string `json:\"name\"`\n\tTTL     int    `json:\"ttl\"`\n\tType    string `json:\"type\"`\n}\n"
var expectedMapSliceTypeStructkg = fmt.Sprintf("package main\n\n%s", expectedMapSliceTypeStruct)
var expectedMapSliceTypeDomain = "type Zone map[string][]Domain\n\ntype Domain struct {\n\tContent string `json:\"content\"`\n\tName    string `json:\"name\"`\n\tTTL     int    `json:\"ttl\"`\n\tType    string `json:\"type\"`\n}\n"
var expectedMapSliceTypeDomainPkg = fmt.Sprintf("package main\n\n%s", expectedMapSliceTypeDomain)

func TestMapSliceType(t *testing.T) {
	// create reader
	tests := []struct {
		structName string
		expected   string
	}{
		{"", expectedMapSliceTypeStructkg},
		{"domain", expectedMapSliceTypeDomainPkg},
	}
	for _, test := range tests {
		r := bytes.NewReader(mapSliceType)
		var buff bytes.Buffer
		calvin := NewTransmogrifier("Zone", r, &buff)
		calvin.MapType = true
		if test.structName != "" {
			calvin.SetStructName(test.structName)
		}
		err := calvin.Gen()
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if buff.String() != test.expected {
			t.Errorf("expected %q got %q", test.expected, buff.String())
		}
	}
}

func TestDefineFieldTags(t *testing.T) {
	tests := []struct {
		keys []string
		value string
		expected string
	}{
		{nil, "field", "`json:\"field\"`"},
		{[]string{}, "field", "`json:\"field\"`"},
		{[]string{"xml"}, "field", "`json:\"field\" xml:\"field\"`"},
		{[]string{"xml", "yaml", "db"}, "field", "`json:\"field\" xml:\"field\" yaml:\"field\" db:\"field\"`"},
	}
	for i, test := range tests {
		tag := defineFieldTags(test.value, test.keys)
		if tag != test.expected {
			t.Errorf("%d: got %q, want %q", i, tag, test.expected)
		}
	}
}

func TestTransmogrify(t *testing.T) {
	tests := []struct {
		pkg        string
		name       string
		structName string
		MapType    bool
		importJSON bool
		json       []byte
		expected   string
	}{
		{"", "Basic", "", false, false, basic, fmt.Sprintf("package main\n\n%s", expectedBasic)},
		{"test", "Basic", "", false, false, basic, fmt.Sprintf("package test\n\n%s", expectedBasic)},
		{"", "Basic", "", false, true, basic, fmt.Sprintf("package main\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedBasic)},
		{"test", "Basic", "", false, true, basic, fmt.Sprintf("package test\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedBasic)},
		{"", "zone", "", true, false, mapType, fmt.Sprintf("package main\n\n%s", expectedMapTypeStruct)},
		{"test", "zone", "", true, false, mapType, fmt.Sprintf("package test\n\n%s", expectedMapTypeStruct)},
		{"", "zone", "", true, true, mapType, fmt.Sprintf("package main\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedMapTypeStruct)},
		{"test", "zone", "", true, true, mapType, fmt.Sprintf("package test\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedMapTypeStruct)},
		{"", "zone", "domain", true, false, mapType, fmt.Sprintf("package main\n\n%s", expectedMapTypeDomain)},
		{"test", "zone", "domain", true, false, mapType, fmt.Sprintf("package test\n\n%s", expectedMapTypeDomain)},
		{"", "zone", "domain", true, true, mapType, fmt.Sprintf("package main\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedMapTypeDomain)},
		{"test", "zone", "domain", true, true, mapType, fmt.Sprintf("package test\n\nimport (\n\t\"encoding/json\"\n)\n\n%s", expectedMapTypeDomain)},
	}
	for i, test := range tests {
		var b bytes.Buffer
		// create reader
		r := bytes.NewReader(test.json)

		calvin := NewTransmogrifier(test.name, r, &b)
		if test.pkg != "" {
			calvin.SetPkg(test.pkg)
		}
		calvin.ImportJSON = test.importJSON
		if test.MapType {
			calvin.MapType = test.MapType
			if test.structName != "" {
				calvin.SetStructName(test.structName)
			}
		}
		err := calvin.Gen()
		if err != nil {
			t.Errorf("%d: unexpected error %q", i, err)
			continue
		}
		if b.String() != test.expected {
			t.Errorf("%d: expected %q, got %q", i, test.expected, b.String())
		}
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
