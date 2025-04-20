package fluffy

import (
	"github.com/andviro/fluffy/v2/num"
)

type MembershipFunc[T num.Num[T]] interface {
	Value(x T) T
	MarshalYAML() (interface{}, error)
}

type Term[T num.Num[T]] struct {
	Name           TermName          `yaml:"name"`
	MembershipFunc MembershipFunc[T] `yaml:"mf"`
}

func (t *Term[T]) MembershipValue(x T) T {
	if t.MembershipFunc == nil {
		return num.NaN[T]()
	}
	return t.MembershipFunc.Value(x)
}
