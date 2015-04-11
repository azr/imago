package imago_test

import (
	"github.com/azr/imago"

	"testing"
)

func TestRowToImageRow(t *testing.T) {
	ig := get4x4IncrementalGray()
	ig2 := get4x4IncrementalGray()

	chosenRow := 1

	r1 := imago.GetRow(ig, chosenRow)
	r2 := imago.GetRow(ig2, chosenRow)

	r1.ToImageRow(ig, 0)
	r1.ToImageRow(ig, 1)
	r1.ToImageRow(ig, 2)
	r1.ToImageRow(ig, 3)

	for i := 0; i < 4; i++ {
		if r1.Pix[i] != ig.Pix[i+r1.Stride] {
			t.Errorf("r1.Pix[%d] != ig.Pix[%d+%d] : %d , %d", i, i, r1.Stride, ig.Pix[i], ig.Pix[i+r1.Stride])
		}
		r3 := imago.GetRow(ig, i)
		for n, _ := range r3.Pix {
			if r3.Pix[n] != r2.Pix[n] {
				t.Errorf("r3.Pix[%d] != r2.Pix[%d]: %d, %d", n, n, r3.Pix[n], r2.Pix[n])
			}
		}
	}
}

func TestRowToImageCol(t *testing.T) {
	ig := get4x4IncrementalGray()
	ig2 := get4x4IncrementalGray()

	chosenCol := 1

	r1 := imago.GetRow(ig, chosenCol)

	r1.ToImageCol(ig2, 0)
	r1.ToImageCol(ig2, 1)
	r1.ToImageCol(ig2, 2)
	r1.ToImageCol(ig2, 3)

	for i := 0; i < 4; i++ {
		v1, v2 := ig.At(i, chosenCol), ig2.At(chosenCol, i)
		if v1 != v2 {
			t.Errorf("ig.At(%d, %d) !=  ig2.At(%d, %d)", i, chosenCol, chosenCol, i, v1, v2)
		}
	}
}
