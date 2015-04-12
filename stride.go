package imago

import (
	"image"
)

//Stride tells you the stride of a common image
func Stride(i image.Image) int {
	switch img := i.(type) {
	case *image.Alpha:
		return img.Stride
	case *image.Alpha16:
		return img.Stride
	case *image.Gray:
		return img.Stride
	case *image.Gray16:
		return img.Stride
	case *image.NRGBA:
		return img.Stride
	case *image.NRGBA64:
		return img.Stride
	case *image.Paletted:
		return img.Stride
	case *image.RGBA:
		return img.Stride
	case *image.RGBA64:
		return img.Stride
	default:
		panic(UnknownImageType.Error())
	}
}
