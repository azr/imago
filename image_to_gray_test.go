package imago_test

import (
	"github.com/azr/imago"

	"image/color"

	"testing"
)

func TestImageToGray(t *testing.T) {
	i := get4x4IncrementalRGBA()

	igrey := imago.ImageToGray(i)

	rgba := i.At(0, 0).(color.RGBA)
	grey := igrey.At(0, 0).(color.Gray)

	if rgba.B != grey.Y {
		t.Errorf("Color differs: %d != %d", rgba.B, grey.Y)
	}

}
