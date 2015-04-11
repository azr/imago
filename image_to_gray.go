package imago

import (
	"image"
	"image/color"
)

//ImageToGray returns the greyscale of an image
func ImageToGray(src image.Image) *image.Gray {
	// Create a new grayscale image
	bounds := src.Bounds()
	gray := image.NewGray(bounds)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			gray.Set(x, y, color.GrayModel.Convert(src.At(x, y)))
		}
	}
	return gray
}
