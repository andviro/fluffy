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
		fmt.Fprintf(buf, "(%d): %s\n", n, strings.Join(res, ", "))
	}
	return buf.String()
}
