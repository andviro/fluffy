package num_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/andviro/fluffy/v2/num"
	"github.com/andviro/fluffy/v2/num/fixed"
	"github.com/andviro/goldie"
)

func TestNum_StringN(t *testing.T) {
	type tc struct {
		val float64
		n   int
	}
	buf := new(bytes.Buffer)
	for _, tc := range []tc{
		{0, 3},
		{0.1, 3},
		{-0.01, 3},
		{0.0001, 3},
		{10000, 3},
		{10000, 2},
		{10000.0001, 4},
		{10000.0001, 5},
		{-10000.0001, 5},
	} {
		fmt.Fprintf(buf, "%f %d: %s\n", tc.val, tc.n, num.StringN(num.NewF[fixed.Fixed](tc.val), tc.n))
	}
	goldie.Assert(t, "stringn", buf.Bytes())
}
