package imago_test

import (
	"github.com/azr/imago"

	"testing"

	"image"
)

func TestCrossCorrelate(t *testing.T) {
	img := get4x4IncrementalGray()

	autocorr := imago.NormalizedAutoCorrelation(img).(*image.Gray)

	i := 0
	for ; i < 1; i++ {
		if autocorr.Pix[i] == 0 {
			t.Errorf("autocorr.Pix[%d] == 0", i)
		}
	}
	for ; i < 16; i++ {
		if autocorr.Pix[i] != 0 {
			t.Errorf("autocorr.Pix[%d] != 0", i)
		}
	}
}
