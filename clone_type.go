package imago

import (
	"image"
	"image/draw"
)

//CloneType clones an image's type and size without copying Pix
func CloneType(i image.Image, rect image.Rectangle) draw.Image {
	switch img := i.(type) {
	case *image.Alpha:
		return image.NewAlpha(rect)
	case *image.Alpha16:
		return image.NewAlpha16(rect)
	case *image.Gray:
		return image.NewGray(rect)
	case *image.Gray16:
		return image.NewGray16(rect)
	case *image.NRGBA:
		return image.NewNRGBA(rect)
	case *image.NRGBA64:
		return image.NewNRGBA64(rect)
	case *image.Paletted:
		return image.NewPaletted(rect, img.Palette)
	case *image.RGBA:
		return image.NewRGBA(rect)
	case *image.RGBA64:
		return image.NewRGBA64(rect)
	default:
		panic(UnknownImageType.Error())
	}
}
