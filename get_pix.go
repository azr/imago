package imago

import (
	"image"
)

func GetPix(i image.Image) []uint8 {
	switch img := i.(type) {
	case *image.Alpha:
		return img.Pix
	case *image.Alpha16:
		return img.Pix
	case *image.Gray:
		return img.Pix
	case *image.Gray16:
		return img.Pix
	case *image.NRGBA:
		return img.Pix
	case *image.NRGBA64:
		return img.Pix
	case *image.Paletted:
		return img.Pix
	case *image.RGBA:
		return img.Pix
	case *image.RGBA64:
		return img.Pix
	default:
		panic(UnknownImageType.Error())
	}
}
