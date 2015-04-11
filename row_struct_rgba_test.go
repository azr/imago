package imago_test

import (
	"github.com/azr/imago"

	"image"
	"image/color"
	"testing"
)

func getRGBA() *image.RGBA {
	i := image.NewRGBA(image.Rect(0, 0, 4, 4))
	i.SetRGBA(0, 0, color.RGBA{0, 0, 0, 0})
	i.SetRGBA(1, 0, color.RGBA{1, 1, 1, 1})
	i.SetRGBA(2, 0, color.RGBA{2, 2, 2, 2})
	i.SetRGBA(3, 0, color.RGBA{3, 3, 3, 3})

	i.SetRGBA(0, 1, color.RGBA{4, 4, 4, 4})
	i.SetRGBA(1, 1, color.RGBA{5, 5, 5, 5})
	i.SetRGBA(2, 1, color.RGBA{6, 6, 6, 6})
	i.SetRGBA(3, 1, color.RGBA{7, 7, 7, 7})

	return i
}

func TestRowRGBA(t *testing.T) {

	i := get4x4IncrementalGray()

	r := imago.GetRow(i, 0)
	if r.Stride != i.Stride {
		t.Fatal("r.Stride != i.Stride")
	}

	for i := 0; i < r.Stride; i++ {
		p := r.Pix[i]
		if p != uint8(i) {
			t.Fatalf("r.Pix[%d] != %d, : %d ", p)
		}
	}
}

func TestHeightenRGBA(t *testing.T) {
	r := imago.GetRow(get4x4IncrementalGray(), 0)
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
