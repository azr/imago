package imago

import (
	"image"
	"image/draw"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//ToImageRow writes row to image's row y (-)
func (r *Row) ToImageRow(dst image.Image, y int) {
	dstd /*, ok*/ := dst.(draw.Image)
	// if !ok { // Let's panic !
	// 	return errors.New("ToImageX: dst image cannot be drawed in")
	// }
	stop := min(Width(dstd), r.Stride)

	for x := 0; x < stop; x++ {
		dstd.Set(x, y, r.At(x))
	}
}

//ToImageRow writes row to image's columns x (|)
func (r *Row) ToImageCol(dst image.Image, x int) {
	dstd /*, ok*/ := dst.(draw.Image)
	// if !ok { // Let's panic !
	// 	return errors.New("ToImageY: dst image cannot be drawed in")
	// }
	stop := min(Height(dstd), r.Stride)

	for y := 0; y < stop; y++ {
		dstd.Set(x, y, r.At(y))
	}
}
