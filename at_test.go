package imago_test

import (
	"github.com/azr/imago"
	"image/color"

	"testing"
)

func TestAt(t *testing.T) {
	i := get4x4IncrementalGray()

	r, _ := imago.GetRow(i, 0)

	c := r.At(3).(color.Gray).Y
	if c != 3 {
		t.Errorf("c != 4 : %d", c)
	}
}

func TestAtOutOfBounds(t *testing.T) {
	r, _ := imago.GetRow(get4x4IncrementalGray(), 0)

	defer func() {
		expected := "At: out of bounds"
		if r := recover(); r == nil || r != expected {
			t.Errorf("At: out of bound expected '%s', got '%s' instead", expected, r)
		}
	}()
	r.At(4)
}
