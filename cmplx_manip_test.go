package imago_test

import (
	"github.com/azr/imago"

	"testing"
)

func TestComplex128UpperBounds(t *testing.T) {
	v := []complex128{254i, 3i, 0, 255i, 2i, 0, 255i, 2i, 0}
	c := imago.Complex128UpperBounds(v, 3, 2, 2)

	if !cmpxEq(c, []complex128{254i, 3i, 255i, 2i}) {
		t.Errorf("c != []complex128{254i, 3i, 255i, 2i}: %#v", c)
	}
}

func cmpxEq(a, b []complex128) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
