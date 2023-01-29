package mf

import (
	"github.com/andviro/fluffy/num"
)

type Singleton struct {
	A num.Num
}

func (f Singleton) MarshalYAML() (interface{}, error) {
	return struct {
		Type string  `yaml:"type"`
		A    num.Num `yaml:"a"`
	}{"Singleton", f.A}, nil
}

func (f *Singleton) Value(x num.Num) num.Num {
	if x.Equal(f.A) {
		return one
	}
	return num.ZERO
}
