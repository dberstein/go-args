package main

import (
	"fmt"

	"github.com/dberstein/go-args"
	"github.com/dberstein/go-args/flag"
)

func main() {
	flagSet := flag.Set{
		flag.DefaultFlag,
		flag.New(flag.Named("cmd"), flag.Has(2, 2)),
		flag.New(flag.Named("cmd"), flag.Has(2, 0), flag.Prefixed("%%")),
	}

	arguments, err := args.New(flagSet...).FromArgv()
	if err != nil {
		panic(err)
	}

	for _, f := range flagSet {
		fmt.Printf("Marker '%s': %q\n", f.Canonic(), *arguments.Values(f))
	}
}
