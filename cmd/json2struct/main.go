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
	name       string
	pkg        string
	input      string
	output     string
	writeJSON  bool
	importJSON bool
	help bool
)

func init() {
	flag.BoolVar(&help, "help", false, "json2struct help")
	flag.BoolVar(&help, "h", false, "the short flag for -help")
	flag.StringVar(&name, "name", "", "the name of the struct")
	flag.StringVar(&name, "n", "", "the short flag for -name")
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

	t := json2struct.NewTransmogrifier(name, in, out)
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

flag             default   description
--------------   -------   ------------------------------------------
-n  -name                  The name of the struct: required.
-p  -pkg         main      The name of the package.
-i  -input       stdin     The JSON input source.
-o  -output      stdout    The Go srouce code output destination.
-w  -writejson   false     Write the source JSON to file; only valid
                           when the output is a file.
-m  -import      false     Add import statement for 'encoding/json'.
-h  -help        false     Print the help text; 'help' is also valid.
`
	fmt.Println(helpText)
}
