package imago_test

import (
	"github.com/azr/imago"

	"testing"
)

func TestRowDivide(t *testing.T) {

	i := get4x4IncrementalGray()

	r1 := imago.GetRow(i, 0)

	r2 := imago.GetRow(i, 0)
	r2.AddRow(r1)
	r2.DivideValues(2)

	for i := 0; i < 4; i++ {
		v := int(r1.Pix[i])
		if v != i {
			t.Errorf("r1.Pix[%d] != %d : %d", i, i, v)
		}

		if r2.Pix[i] != r1.Pix[i] {
			t.Errorf("r2.Pix[%d] != r1.Pix[%d] : %d, %d", i, i, r2.Pix[i], r1.Pix[i])
		}
	}

}
