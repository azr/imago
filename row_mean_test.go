package imago_test

import (
	"github.com/azr/imago"

	"testing"
)

func TestRowMean(t *testing.T) {
	i := get4x4IncrementalGray()

	r1, _ := imago.GetRow(i, 0)
	r1.MeanRow(r1)
	for n, v := range r1.Pix {
		if int(v) != n {
			t.Errorf("int(v) != n : %d != %d", int(v), n)
		}
	}

}
