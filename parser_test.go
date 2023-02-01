package fluffy_test

import (
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/andviro/goldie"

	"github.com/andviro/fluffy"
)

func TestParser(t *testing.T) {
	r, err := fluffy.ParseRules(`
    1 : a = b & c = d | a = e -> z = y
    1.5 : x = y & (x = d | x = f) -> z = t
`)
	if err != nil {
		t.Fatal(err)
	}
	yd, err := yaml.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	goldie.Assert(t, "parse", yd)
}
