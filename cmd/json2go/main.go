// json2go generates Go type definitions from JSON encoded objects as
// defined in RFC 4627.  The type can be one of the following:
//    map[string]T
//    map[string][]T
//    Type
//
// The result is Go source for the provided package name, with the type
// definiton(s) for the JSON.  If no package name is provided, the name
// will either be the name of the output directory, if the output is a
// file, or the working directory.
//
// If the JSON is an array of elements, e.g. []T or []map[string]T, the
// first element will be used to generate the definitions.
//
// By default a struct type will be generated, unless the -maptype flag is
// used.
//
// If a type contains other JSON objects, separate structs are defined
// and the struct is embedded in the defintion.
//
// If an output destination is specified, the generated Go source will be
// written to the specified destination, otherwise it will be written to
// stdout.
//
// Optionally, if there is an output destination that isn't stdout, the
// source JSON can be written out to a file.  This file's name will be the
// same as the Go source file's name, ending in '.json'.  This may be
// useful when grabbing the JSON from a remote source and piping it into
// json2go via stdin.
//
// Errors are written to stderr
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mohae/json2go"
)

var (
	name       string
	pkg        string
	input      string
	output     string
	structName string
	writeJSON  bool
	importJSON bool
	mapType    bool
	help       bool
)

func init() {
	flag.StringVar(&name, "name", "", "the name of the type")
	flag.StringVar(&name, "n", "", "the short flag for -name")
	flag.StringVar(&input, "input", "stdin", "the path to the input file; if not specified stdin is used")
	flag.StringVar(&input, "i", "stdin", "the short flag for -input")
	flag.StringVar(&output, "output", "stdout", "path to the output file; if not specified stdout is used")
	flag.StringVar(&output, "o", "stdout", "the short flag for -output")
	flag.StringVar(&pkg, "pkg", "", "the name of the package")
	flag.StringVar(&pkg, "p", "", "the short flag for -pkg")
	flag.StringVar(&structName, "structname", "Struct", "the name of the struct; only used with -maptype")
	flag.StringVar(&structName, "s", "Struct", "the short flag for -structname")
	flag.BoolVar(&writeJSON, "writejson", false, "write the source JSON to file; if the output destination is stdout, this flag will be ignored")
	flag.BoolVar(&writeJSON, "w", false, "the short flag for -writejson")
	flag.BoolVar(&importJSON, "addimport", false, "add import statement for encoding/json")
	flag.BoolVar(&importJSON, "a", false, "the short flag for -addimport")
	flag.BoolVar(&mapType, "maptype", false, "the provided json is a map type; not a struct type")
	flag.BoolVar(&mapType, "m", false, "the short flag for -maptype")
	flag.BoolVar(&help, "help", false, "json2go help")
	flag.BoolVar(&help, "h", false, "the short flag for -help")
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
		fmt.Fprintln(os.Stderr, "\nstruct2json error: name of struct must be provided using the -n or -name flag.\nUse the '-h', '-help', or 'help' flag for more information about json2go flags.")
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
		defer out.Close()
		// set the package name, if one isn't set
		if len(pkg) == 0 {
			// get the rooted path to the output
			output, err := filepath.Abs(output)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return 1
			}
			base := filepath.Base(filepath.Dir(output))
			if base != string(os.PathSeparator) && base != "." {
				pkg = base
			}
		}
		// write the source json if applicable
		if writeJSON {
			jsn, err = os.OpenFile(fmt.Sprintf("%s.json", strings.TrimSuffix(output, filepath.Ext(output))), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return 1
			}
			defer jsn.Close()
		}
	}
	// there is a chance pkg hasn't been set; get the wd and set pkg to
	// the parent element of cwd
	if len(pkg) == 0 {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		// get the parent dir name
		base := filepath.Base(dir)
		if base != string(os.PathSeparator) && base != "." {
			pkg = base
		}
	}
	// create the transmogrifier and configure it.
	t := json2go.NewTransmogrifier(name, in, out)
	if writeJSON {
		t.WriteJSON = writeJSON
		t.SetJSONWriter(jsn)
	}
	t.ImportJSON = importJSON
	if pkg != "" {
		t.SetPkg(pkg)
	}
	t.MapType = mapType
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

Go type definitions will be generated from the input JSON.  The
generated Go code will be part of package main, unless a different
package is specified using either the -p or -pkg flags.

The name for the struct must be specified using either the -n or
-name flags; name is required.

A JSON source file can be specified with either the -i or -input
flags.  If none is specified, the JSON is expected to come from
stdin.

The output file of the generated Go source code is specified
with either the -o or -output flags.  If none is specified, the
output will be written to stdout.

Errors are written to stderr.

Minimal examples:

    $ curl http://example.com/source.json | json2strct -n example

or

    $ json2go -i example.json -o example.go -n example

Options:

flag              default   description
---------------   -------   ------------------------------------------
-n  -name                   The name of the type: required.
-i  -input        stdin     The JSON input source.
-o  -output       stdout    The Go srouce code output destination.
-w  -writejson    false     Write the source JSON to file; only valid
                            when the output is a file.
-p  -pkg                    The name of the package.
-a  -addimport    false     Add import statement for 'encoding/json'.
-m  -maptype      false     Interpret the JSON as a map type instead
                            of a struct type.
-s  -structname   Struct    The name of the struct; only used in
                            conjunction with -maptype.
-h  -help         false     Print the help text; 'help' is also valid.
`
	fmt.Println(helpText)
}
