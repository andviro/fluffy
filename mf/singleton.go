package mf

import (
	"github.com/andviro/fluffy/v2/num"
)

type Singleton[T num.Num[T]] struct {
	A T
}

func (f Singleton[T]) MarshalYAML() (any, error) {
	return struct {
		Type string `yaml:"type"`
		A    string `yaml:"a"`
	}{"Singleton", f.A.String()}, nil
}

func (f *Singleton[T]) Value(x T) T {
	if x.Equal(f.A) {
		return num.One[T]()
	}
	return num.ZERO[T]()
}
