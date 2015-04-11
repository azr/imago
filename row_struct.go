package imago

import (
	"image"
)

//Row represents a row from an image.
//
// Use it when you need per-rown calculus and/or when
// you want to split calculus in more a readable way.
type Row struct {
	//Slice reference to image row
	Pix       []uint8
	PixelSize int
	//Stride of image, which is also lenght of row
	Stride int
	//Image the Row is from
	From image.Image
}

//GetRow returns you a Rown from a known image type or panics.
//
//Panic will in most cases happen during dev.
func GetRow(i image.Image, y int) *Row {
	switch img := i.(type) {
	case *image.Alpha:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 1,
			Stride:    img.Stride,
			From:      i,
		}
	case *image.Alpha16:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 2,
			Stride:    img.Stride,
			From:      i,
		}
	case *image.Gray:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 1,
			Stride:    img.Stride,
			From:      i,
		}
	case *image.Gray16:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 1,
			Stride:    img.Stride,
			From:      i,
		}
	case *image.NRGBA:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 4,
			Stride:    img.Stride,
			From:      i,
		}
	case *image.NRGBA64:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 8,
			Stride:    img.Stride,
			From:      i,
		}
	case *image.Paletted:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 1,
			Stride:    img.Stride,
			From:      i,
		}
	case *image.RGBA:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 4,
			Stride:    img.Stride,
			From:      i,
		}
	case *image.RGBA64:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 8,
			Stride:    img.Stride,
			From:      i,
		}
	default:
		panic(UnknownImageType.Error())
	}
}
