package args

import (
	"os"
)

var (
	// DefaultBucket is default arguments bucket
	DefaultBucket *Bucket
)

func init() {
	defaultPrefix := "--"
	DefaultBucket = &Bucket{prefix: &defaultPrefix}
}

// Buckets is array of Bucket pointers
type Buckets []*Bucket

// Bucket is argument bucket
type Bucket struct {
	prefix *string
	name   string
}

// WithPrefix is bucket with given prefix
func WithPrefix(name, prefix string) *Bucket {
	return &Bucket{name: name, prefix: &prefix}
}

// With is bucket with default prefix factory
func With(name string) *Bucket {
	return WithPrefix(name, *DefaultBucket.prefix)
}

// Bucket is full name (with prefix) of bucket
func (m *Bucket) Bucket() string {
	if m.prefix == nil {
		return *DefaultBucket.prefix + m.name
	}
	return *m.prefix + m.name
}

// Args is exported arguments type
type Args struct {
	*arguments
}

// arguments in unexported arguments type
type arguments struct {
	defaultBucket *Bucket
	buckets       Buckets
	parsed        map[string][]string
}

// FromArgv is parser of given args slice
func (a *Args) FromArgv() *Args {
	argv := os.Args[1:]
	return a.From(&argv)
}

// append value to bucket
func (a *Args) append(value, bucket *string) *Args {
	a.parsed[*bucket] = append(a.parsed[*bucket], *value)
	return a
}

// From is parser of given args slice
func (a *Args) From(args *[]string) *Args {
	a.Clear()
	stop := false
	current := *DefaultBucket.prefix
	for _, arg := range *args {
		if stop {
			a.append(&arg, &current)
			continue
		} else if arg == *DefaultBucket.prefix {
			stop = !stop
		}
		if _, ok := a.parsed[arg]; ok {
			current = arg
			continue
		}
		a.append(&arg, &current)
	}
	return a
}

// Parser is factory of Arguments
func Parser(bucket ...*Bucket) *Args {
	args := arguments{
		defaultBucket: DefaultBucket,
		buckets:       append(bucket, DefaultBucket),
	}
	return &Args{&args}
}

// Clear argument values
func (a *Args) Clear() *Args {
	values := map[string][]string{}
	for _, m := range a.buckets {
		values[m.Bucket()] = []string{}
	}
	a.parsed = values
	return a
}

// Values returns values for argument bucketName
func (a *Args) Values(bucketName string) *[]string {
	if args, ok := a.parsed[bucketName]; ok {
		return &args
	}
	return nil
}

// Has returns presence of arguments for name
func (a *Args) Has(name string) (has bool) {
	if args, has := a.parsed[name]; has {
		has = len(args) > 0
	}
	return
}
