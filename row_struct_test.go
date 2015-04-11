package imago_test

import (
	"github.com/azr/imago"

	// "image"
	// "image/color"
	"testing"
)

func TestGetRow(t *testing.T) {
	i := get4x4IncrementalGray()

	r, _ := imago.GetRow(i, 0)
	if len(r.Pix) != imago.Width(i) {
		t.Errorf("len(r.Pix) != imago.Width(i): %d %d", len(r.Pix), imago.Width(i))
	}
}
