package imago

import (
	"image"
)

//PixelSize tells you the size of a pixel for the main image types
func PixelSize(i image.Image) int {
	switch i.(type) {
	case *image.Alpha:
		return 1
	case *image.Alpha16:
		return 2
	case *image.Gray:
		return 1
	case *image.Gray16:
		return 1
	case *image.NRGBA:
		return 4
	case *image.NRGBA64:
		return 8
	case *image.Paletted:
		return 1
	case *image.RGBA:
		return 4
	case *image.RGBA64:
		return 8
	default:
		panic(UnknownImageType.Error())
	}
}
