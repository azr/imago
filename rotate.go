package imago

import (
	"image"
	"image/draw"

	"code.google.com/p/graphics-go/graphics"

	"math"
)

const (
	DegToRad = math.Pi / 180
)

func RotateDeg(dst image.Image, i image.Image, degrees float64) error {
	dstd /*, ok*/ := dst.(draw.Image)
	// if !ok { // Let's panic !
	// 	return errors.New("RotateDeg: dst image cannot be drawed in")
	// }
	return graphics.Rotate(dstd, i, &graphics.RotateOptions{
		Angle: degrees * DegToRad,
	})
}
