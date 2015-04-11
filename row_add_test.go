package imago_test

import (
	"github.com/azr/imago"

	"testing"
)

func TestRowAdd(t *testing.T) {
	i := get4x4IncrementalGray()

	r1 := imago.GetRow(i, 0)

	r1.AddRow(r1)

	for i := 0; i < 4; i++ {
		v := int(r1.Pix[i])
		if v != i*2 {
			t.Errorf("r1.[%d] != %d : %d", i, i*2, v)
		}
	}
}
