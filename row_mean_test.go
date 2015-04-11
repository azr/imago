package imago_test

import (
	"github.com/azr/imago"

	"testing"
)

func TestAverageRow(t *testing.T) {
	i := get4x4IncrementalGray()

	r1 := imago.GetRow(i, 0)

	rr := imago.AverageRow(r1, r1, r1, r1, r1, r1, r1)
	for n, _ := range r1.Pix {
		if rr.Pix[n] != r1.Pix[n] {
			t.Errorf("rr.Pix[%d] != r1.Pix[%d], %d %d", n, n, rr.Pix[n], r1.Pix[n])
		}
	}

	r2 := &imago.Row{
		Pix: []uint8{7},
	}
	r3 := &imago.Row{
		Pix: []uint8{9},
	}
	r4 := &imago.Row{
		Pix: []uint8{10},
	}

	rr = imago.AverageRow(r2, r3, r4)
	if rr.Pix[0] != 8 {
		t.Errorf("rr.Pix[0] != 8 : %d", rr.Pix[0])
	}
}

func TestRowMean(t *testing.T) {
	i := get4x4IncrementalGray()

	r1 := imago.GetRow(i, 0)
	r1.MeanRow(r1)
	for n, v := range r1.Pix {
		if int(v) != n {
			t.Errorf("int(v) != n : %d != %d", int(v), n)
		}
	}

}
