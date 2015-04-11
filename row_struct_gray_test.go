package imago_test

import (
	"github.com/azr/imago"

	"image"
	"image/color"
	"testing"
)

func get4x4IncrementalGray() *image.Gray {
	i := image.NewGray(image.Rect(0, 0, 4, 4))
	i.SetGray(0, 0, color.Gray{0})
	i.SetGray(1, 0, color.Gray{1})
	i.SetGray(2, 0, color.Gray{2})
	i.SetGray(3, 0, color.Gray{3})

	i.SetGray(0, 1, color.Gray{4})
	i.SetGray(1, 1, color.Gray{5})
	i.SetGray(2, 1, color.Gray{6})
	i.SetGray(3, 1, color.Gray{7})

	return i
}

func TestRowGray(t *testing.T) {

	i := get4x4IncrementalGray()

	r, _ := imago.GetRow(i, 0)
	if r.Stride != i.Stride {
		t.Fatal("r.Stride != i.Stride")
	}

	for i := 0; i < r.Stride; i++ {
		p := r.Pix[i]
		if p != uint8(i) {
			t.Fatalf("r.Pix[%d] != %d, : %d ", i, i, p)
		}
	}
}

func TestHeightenGray(t *testing.T) {
	r, _ := imago.GetRow(get4x4IncrementalGray(), 0)
	height := 70
	ii := r.Heighten(height)

	img := ii.(*image.Gray)

	if imago.Height(img) != height {
		t.Fatalf("Image height %d != %d", imago.Height(img), height)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < r.Stride; j++ {
			c := img.GrayAt(j, i).Y
			if c != uint8(j) {
				t.Fatalf("img.GrayAt(%d, %d).Y != %d : %d", j, i, j, c)
			}
		}
	}

}
