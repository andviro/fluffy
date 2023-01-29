package fis

import (
	"bytes"
	"fmt"
	"strings"
)

func (f *TSK) Dump() string {
	buf := new(bytes.Buffer)
	for n, i := range f.Inputs {
		var res []string
		for k, v := range i.GetTermValues() {
			res = append(res, fmt.Sprintf("%s:%v", k, v))
		}
		fmt.Fprintf(buf, "in %d: %s\n", n, strings.Join(res, ", "))
	}
	for n, i := range f.Outputs {
		for _, t := range i.Terms {
			fmt.Fprintf(buf, "out %d/%s: %v\n", n, t.Name, t.z)
		}
		fmt.Fprintf(buf, "out %d=%v %v\n", n, i.GetValue(), i.evaluations)
	}
	return buf.String()
}
