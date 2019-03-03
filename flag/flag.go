package flag

import (
	"errors"
	"fmt"
)

var (
	// DefaultFlag is default arguments flag
	DefaultFlag = New(Prefixed("--"))
)

// Set is array of Flag pointers
type Set []*Flag

type valueValidator func([]string) error

// Flag is argument values separator
type Flag struct {
	prefix   *string
	name     string
	min, max uint
	kind     uint
}

// New returns new Flag
func New(opts ...Option) *Flag {
	options := &Options{}
	for _, o := range opts {
		o(options)
	}

	return &Flag{
		prefix: options.prefix,
		name:   options.name,
		min:    options.min,
		max:    options.max,
		kind:   options.kind,
	}
}

// Canonic clag name
func (f *Flag) Canonic() string {
	if f.prefix == nil {
		return *DefaultFlag.prefix + f.name
	}
	return *f.prefix + f.name
}

// Prefix pointer
func (f *Flag) Prefix() *string {
	return f.prefix
}

// Valid check validity of argument values
func (f *Flag) Valid(values []string) error {
	validators := []valueValidator{
		f.validHas,
		f.validKind,
	}
	for _, v := range validators {
		err := v(values)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Flag) validHas(values []string) error {
	nValues := uint(len(values))
	if f.min > 0 && nValues < f.min {
		return fmt.Errorf("%s expects at least %d values, received %d", f.Canonic(), f.min, nValues)
	}
	if f.max > 0 && nValues > f.max {
		return fmt.Errorf("%s expects at most %d values, received %d", f.Canonic(), f.max, nValues)
	}
	return nil
}

func (f *Flag) validKind(values []string) error {
	return nil
}

// Options of Flag
type Options struct {
	prefix   *string
	name     string
	min, max uint
	kind     uint
}

// Option for Flag type
type Option func(*Options)

// Named Flag options
func Named(name string) Option {
	return func(options *Options) {
		options.name = name
	}
}

// Prefixed Flag options
func Prefixed(prefix string) Option {
	return func(options *Options) {
		if len(prefix) == 0 {
			panic(errors.New("prefix cannot be of zero length"))
		}
		options.prefix = &prefix
	}
}

// Has Flag options
func Has(min, max uint) Option {
	return func(options *Options) {
		if min > 0 && min > max && max > 0 {
			panic(fmt.Errorf("Has: non-zero min cannot be greater than max (%d > %d)", min, max))
		}
		options.min = min
		options.max = max
	}
}

func Kind(kind uint) Option {
	return func(options *Options) {
		options.kind = kind
	}
}
