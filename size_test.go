package imago_test

import (
	"github.com/azr/imago"

	"image"
	"testing"
)

func TestWidth(t *testing.T) {
	i := image.NewGray(image.Rect(0, 0, 5, 2))

	if imago.Width(i) != 5 {
		t.Errorf("imago.Width(i) != 5: %d", imago.Width(i))
	}
}

func TestHeight(t *testing.T) {
	i := image.NewNRGBA64(image.Rect(0, 0, 2, 5))

	if imago.Height(i) != 5 {
		t.Errorf("imago.Height(i) != 5: %d", imago.Height(i))
	}
}
