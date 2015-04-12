package imago

import (
	"image"
)

//Row represents a row from an image.
//
// Use it when you need per-row calculus and/or when
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
//Panic will mostly happen during dev !
func GetRow(i image.Image, y int) *Row {
	row := &Row{
		PixelSize: PixelSize(i),
		Stride:    Stride(i),
		From:      i,
	}
	switch img := i.(type) {
	case *image.Alpha:
		row.Pix = img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride]
	case *image.Alpha16:
		row.Pix = img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride]
	case *image.Gray:
		row.Pix = img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride]
	case *image.Gray16:
		row.Pix = img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride]
	case *image.NRGBA:
		row.Pix = img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride]
	case *image.NRGBA64:
		row.Pix = img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride]
	case *image.Paletted:
		row.Pix = img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride]
	case *image.RGBA:
		row.Pix = img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride]
	case *image.RGBA64:
		row.Pix = img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride]
	default:
		panic(UnknownImageType.Error())
	}
	return row
}
