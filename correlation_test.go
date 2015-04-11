package imago_test

import (
	"github.com/azr/imago"

	"testing"

	"image"
)

func TestCrossCorrelateGray(t *testing.T) {
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

func TestCrossCorrelateRGBA(t *testing.T) {
	img := get4x4IncrementalRGBA()

	autocorr := imago.NormalizedAutoCorrelation(img).(*image.RGBA)

	i := 0
	for ; i < 4; i++ {
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
