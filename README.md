json2struct
===========
[![Build Status](https://travis-ci.org/mohae/json2struct.png)](https://travis-ci.org/mohae/json2struct)

Generate Go struct(s) from example JSON.  The result is a Go source file with the struct definition(s).  If the JSON includes objects, those objects will be created as their own struct and embedded.

By default, the package name used is `main`, but the package name can be set to a custom value.  Optionally, the import statement for `encoding/json` can be added to the generated Go source.

This package reads from the provided reader and writes to the provided writer.  The JSON used to generate the struct(s) can also be written to a writer.

Keys with underscores, `_`, are converted to MixedCase.  Keys starting with characters that are invalid for Go variable names have those characters discarded, unless they are a number, `0-9`, which are converted to their word equivalents. All fields are exported and the JSON field tag for the field is generated using the original JSON key value.

If a field's value is null, the field's type will be `interface{}`, as that field's type is not determinable.

The CLI is in the `cmd/json2struct` subdirectory.  See that [README](https://github.com/mohae/json2struct/tree/master/cmd/json2struct) for more info.
