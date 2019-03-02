# go-args
argv processor for Go

## Installation ##
```
go get github.com/dberstein/go-args
```
... or to add as requirement to go.mod ...
```
go mod edit -require github.com/dberstein/go-args@master
```
## Usage ##

```go
package main

import (
	"fmt"

	"github.com/dberstein/go-args"
)

func main() {
	// Create argument bucket for "--name"
	namesBucket := args.With("name")
	// Create parser with single bucket (note that default bucket "--" is created automatically)
	arguments := args.Parser(namesBucket).FromArgv()
	// Print values found in default and "name" buckets
	for _, bucket := range []string{args.DefaultBucket.Bucket(), namesBucket.Bucket()} {
		fmt.Printf("Bucket %s: %q\n", bucket, *arguments.Values(bucket))
	}
}
```
    go run example.go p1 --name n1 n2 -- p2
would produce:

    Bucket --: ["p1", "p2"]
    Bucket --name: ["n1", "n2"]

### Create arguments ###
```go
// Create argument "--name"
arg1 := args.With("name")

// Create argument with non-standard prefix
arg2 := args.WithPrefix("name", "%")
```
### Parse arguments ###
```
// Create parser for both arguments
parser := args.Parser(arg1, arg2)
```
### Read argument values ###
```go
// Create []string for "d0 --name n1 n2 %name n3 --name n4 -- d1"
params := []string{"d0", "--name", "n1", "n2", "%name", "n3", "--name", "n4", "--", "d1"}
parser := args.Parser(args.With("name"), args.WithPrefic("name", "%"))
parsed := parser.From(*params)
// For convenience args.FromArgv() processed os.Argv[1:] (drops os.Arv[0])
parsedArgv := parser.FromArgv()

// names1 will be []string{"n1", "n2", "n4"}
names1 := parsed.Values("--name")

// names2 will be []string{"n3"}
names2 := parsed.Values("%name")

// default will be []string{"d0", "d1"}
default := parsed.Values("--")

// For convenience args.FromArgv() processes os.Argv[1:] (drops os.Arv[0])
parsed := parser.FromArgv()
```

Godoc: https://godoc.org/github.com/dberstein/go-args
