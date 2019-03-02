package args_test

import (
	"fmt"

	"github.com/dberstein/go-args"
)

func Example() {
	// Create argument bucket for "--name"
	namesBucket := args.With("name")
	// Create parser with single bucket (note that default bucket "--" is created automatically)
	arguments := args.Parser(namesBucket).FromArgv()
	// Print values found in default and "name" buckets
	for _, bucket := range []string{args.DefaultBucket.Bucket(), namesBucket.Bucket()} {
		fmt.Printf("Bucket %s: %q\n", bucket, *arguments.Values(bucket))
	}

	// Output: go run example.go p1 --name n1 n2 -- p2
	// Bucket --: ["p1", "p2"]
	// Bucket --name: ["n1", "n2"]
}
