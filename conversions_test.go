package imago_test

import (
	"github.com/azr/imago"

	"testing"
)

func TestPixelsToComplex(t *testing.T) {
	v := []uint8{255, 255, 255, 255, 255, 255, 255}
	c := imago.PixelsToFloat64(v, 7)

	if c[0] != float64(0xFFFFFFFFFFFFFF) {
		t.Errorf("c[0] != 0xFFFFFFFFFFFFFF : %F, expected %F", c[0], 0xFFFFFFFFFFFFFF)
	}
}

func TestRGBAToFloat64(t *testing.T) {
	rgba := get4x4IncrementalRGBA()
	c := imago.ToFloat64(rgba)

	if len(c) != 16 {
		t.Errorf("len(c) != 16 : ", len(c))
	}

	for i := 0; i < 4; i++ {
		v := float64(i<<24 + i<<16 + i<<8 + i)
		if c[i] != v {
			t.Errorf("c[%d] != v, %F != %F", i, c[i], v)
		}
	}
}
