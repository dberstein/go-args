package args

import (
	"os"

	"github.com/dberstein/go-args/flag"
)

// arguments in unexported arguments type
type arguments struct {
	defaultFlag *flag.Flag
	flags       []*flag.Flag
	values      map[string][]string
}

// Args is exported arguments type
type Args struct {
	*arguments
}

// append argument to flag values
func (a *Args) append(arg, flagName *string) *Args {
	a.values[*flagName] = append(a.values[*flagName], *arg)
	return a
}

// From is parser of given args slice
func (a *Args) From(args *[]string) (*Args, error) {
	stop := false
	defaultPrefix := *flag.DefaultFlag.Prefix()
	current := defaultPrefix
	for _, arg := range *args {
		if stop {
			a.append(&arg, &current)
			continue
		} else if arg == defaultPrefix {
			stop = !stop
		}
		if _, ok := a.values[arg]; ok {
			current = arg
			continue
		}
		a.append(&arg, &current)
	}
	for _, f := range a.flags {
		values := a.values[f.Canonic()]
		if err := f.Valid(values); err != nil {
			return nil, err
		}
	}
	return a, nil
}

// FromArgv is parser of given args slice
func (a *Args) FromArgv() (*Args, error) {
	argv := os.Args[1:]
	return a.From(&argv)
}

// New Arguments from flags
func New(f ...*flag.Flag) *Args {
	a := &Args{&arguments{
		defaultFlag: flag.DefaultFlag,
		flags:       f,
		values:      make(map[string][]string),
	}}
	for _, fl := range f {
		a.values[fl.Canonic()] = []string{}
	}
	return a
}

// Values of arguments of Flag
func (a *Args) Values(f *flag.Flag) *[]string {
	if args, ok := a.values[f.Canonic()]; ok {
		return &args
	}
	return nil
}

// Has returns presence of arguments for name
func (a *Args) Has(name string) bool {
	if _, ok := a.values[name]; ok {
		return true
	}
	return false
}
