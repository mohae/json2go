// json2struct generates Go type definitions from JSON encoded objects as
// defined in RFC 4627.  The type can be one of the following:
//    map[string]Type
//    map[string][]Type
//    Type
//
// By default a struct Type will be generated, unless the -ismap flag is
// used.
//
// If a Type contains other JSON objects, separate structs are defined
// and the struct is embedded in the defintion.  The result is Go source
// for the provided package name, with the Type definiton(s) for the source
// JSON.
//
// If an output destination is specified, the generated Go source will be
// written to the specified destination, otherwise it will be written to
// stdout.
//
// Optionally, if there is an output destination that isn't stdout, the
// source json can be written out to a file.  This file's name will be the
// same as the Go source file's name, ending in '.json'.  This may be
// useful when grabbing the JSON from a remote source and piping it into
// json2struct via stdin.
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
	name       string
	pkg        string
	input      string
	output     string
	structName string
	writeJSON  bool
	importJSON bool
	isMap      bool
	help       bool
)

func init() {
	flag.BoolVar(&help, "help", false, "json2struct help")
	flag.BoolVar(&help, "h", false, "the short flag for -help")
	flag.StringVar(&name, "name", "", "the name of the type")
	flag.StringVar(&name, "n", "", "the short flag for -name")
	flag.StringVar(&structName, "structname", "", "the name of the struct; only used with -ismap")
	flag.StringVar(&structName, "s", "", "the short flag for -structname")
	flag.StringVar(&pkg, "pkg", "", "the name of the package")
	flag.StringVar(&pkg, "p", "", "the short flag for -pkg")
	flag.StringVar(&input, "input", "stdin", "the path to the input file; if not specified stdin is used")
	flag.StringVar(&input, "i", "stdin", "the short flag for -input")
	flag.StringVar(&output, "output", "stdout", "path to the output file; if not specified stdout is used")
	flag.StringVar(&output, "o", "stdout", "the short flag for -output")
	flag.BoolVar(&writeJSON, "writejson", false, "write the source JSON to file; if the output destination is stdout, this flag will be ignored")
	flag.BoolVar(&writeJSON, "w", false, "the short flag for -writejson")
	flag.BoolVar(&importJSON, "import", false, "add import statement for encoding/json")
	flag.BoolVar(&importJSON, "m", false, "the short flag for -import")
	flag.BoolVar(&isMap, "ismap", false, "the provided json is a map type; not a struct type")
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()
	args := flag.Args()
	// the only arg we care about is help.  This is in case the user uses
	// just help instead of -help or -h
	for _, arg := range args {
		if arg == "help" {
			help = true
			break
		}
	}
	if help {
		Help()
		return 0
	}
	if name == "" {
		fmt.Fprintln(os.Stderr, "\nstruct2json error: name of struct must be provided using the -n or -name flag.\nUse the '-h', '-help', or 'help' flag for more information about json2struct flags.")
		return 1
	}
	var in, out, jsn *os.File
	var err error
	// set input
	in = os.Stdin
	if input != "stdin" {
		in, err = os.Open(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	}
	defer in.Close()
	// set output
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
	// create the transmogrifier and configure it.
	t := json2struct.NewTransmogrifier(name, in, out)
	if jsn != nil {
		t.SetWriteJSON(writeJSON)
		t.SetJSONWriter(jsn)
	}
	t.SetImportJSON(importJSON)
	if pkg != "" {
		t.SetPkg(pkg)
	}
	t.SetIsMap(isMap)
	t.SetStructName(structName)
	// Generate the Go Types
	err = t.Gen()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func Help() {
	helpText := `
Usage: json2go [options]

Go struct definitions will be generated from the input JSON.  The
generated Go code will be part of package main, unless a different
package is specified using either the -p or -pkg flag.

A JSON source and the name for the struct must be specified.  The JSON
can either be piped in via stdin or a file with the JSON can be
specified with either the -i or -input flag.  The name of the struct
is specified with either the -n or -name flag.

Minimal examples:
    $ curl http://example.com/source.json | json2strct -n example

or

    $ json2struct -i example.json -n example

Options:

flag              default   description
---------------   -------   ------------------------------------------
-n  -name                   The name of the Type: required.
-s  -structname             The name of the struct; only used in
                            conjunction with -ismap.
-p  -pkg          main      The name of the package.
-i  -input        stdin     The JSON input source.
-o  -output       stdout    The Go srouce code output destination.
-w  -writejson    false     Write the source JSON to file; only valid
                            when the output is a file.
-m  -import       false     Add import statement for 'encoding/json'.
    -ismap        false     Interpret the JSON as a map type instead
                            of a struct type.
-h  -help         false     Print the help text; 'help' is also valid.
`
	fmt.Println(helpText)
}
