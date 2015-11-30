// json2struct generates Go struct definitions from JSON encoded objects as
// defined in RFC 4627.  If the source JSON contains other JSON objects,
// separate structs are defined and the object is embedded in the defintion.
// The result is a Go source, for the provided package name, with the
// struct definiton(s) for the source JSON.
//
// The source file's name will be the same as the provided struct's name.
//
// Optionally, the source json can be written out to a file.  This file's
// name will be the same as the Go source file's name, ending in '.json'.
//
// Errors are written to stderr
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mohae/json2struct"
)

var (
	pkg        string
	input      string
	output     string
	writeJSON  bool
	importJSON bool
)

func init() {
	flag.StringVar(&pkg, "pkg", "", "the name of the package")
	flag.StringVar(&pkg, "p", "", "the short flag for -pkg")
	flag.StringVar(&input, "input", "stdin", "the path to the input file; if not specified stdin is used")
	flag.StringVar(&input, "i", "stdin", "the short flag for -input")
	flag.StringVar(&output, "output", "stdout", "path to the output file; if not specified stdout is used")
	flag.StringVar(&output, "o", "stdout", "the short flag for -output")
	flag.BoolVar(&writeJSON, "writejson", false, "write the source JSON to file; if the output destination is stdout, this flag will be ignored")
	flag.BoolVar(&writeJSON, "w", false, "the short flag for -writejson")
	flag.BoolVar(&importJSON, "import", false, "add import statement for encoding/json")
	flag.BoolVar(&importJSON, "m", false, "the shortflat for -import")
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "struct2json error: name of struct must be provided")
		return 1
	}
	var in, out, jsn *os.File
	var err error
	//var in io.Reader
	//var out io.Writer
	in = os.Stdin
	// if it
	if input != "stdin" {
		in, err = os.Open(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	}
	defer in.Close()
	out = os.Stdout
	if output != "stdout" {
		//
		out, err = os.OpenFile(output, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		if writeJSON {
			jsn, err = os.OpenFile(fmt.Sprintf("%s.json", strings.TrimSuffix(output, path.Ext(output))), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return 1
			}
			defer jsn.Close()
		}
	}
	defer out.Close()
	// json is only written out if the output isn't stdout
	// TODO
	// get the output filename - ext
	// add .json ext
	// crete output file
	// pipe input to output
	//if writeJSON && output != "stdout" {

	//}

	t := json2struct.NewTransmogrifier(args[0], in, out)
	if jsn != nil {
		t.SetWriteJSON(writeJSON)
		t.SetJSONWriter(jsn)
	}
	t.SetImportJSON(importJSON)
	if pkg != "" {
		t.SetPkg(pkg)
	}
	err = t.Gen()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
