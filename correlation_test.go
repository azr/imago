package imago_test

import (
	"github.com/azr/imago"

	"testing"
)

func TestCrossCorrelate(t *testing.T) {
	i := get4x4IncrementalGray()

	autocorr := imago.AutoCorrelateUInt8(i.Pix)

	for i := 0; i < 1; i++ {
		if autocorr[i] == 0 {
			t.Errorf("autocorr[%d] == 0", i)
		}
	}
}
