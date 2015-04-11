package imago

import (
	"image"
)

type Row struct {
	Pix       []uint8
	PixelSize int
	Stride    int
	From      image.Image
}

func (r *Row) PixOffset(x int) int {
	return x * r.PixelSize
}

func GetRow(i image.Image, y int) (*Row, error) {
	switch img := i.(type) {
	case *image.Alpha:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 1,
			Stride:    img.Stride,
			From:      i,
		}, nil
	case *image.Alpha16:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 2,
			Stride:    img.Stride,
			From:      i,
		}, nil
	case *image.Gray:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 1,
			Stride:    img.Stride,
			From:      i,
		}, nil
	case *image.Gray16:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 1,
			Stride:    img.Stride,
			From:      i,
		}, nil
	case *image.NRGBA:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 4,
			Stride:    img.Stride,
			From:      i,
		}, nil
	case *image.NRGBA64:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 8,
			Stride:    img.Stride,
			From:      i,
		}, nil
	case *image.Paletted:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 1,
			Stride:    img.Stride,
			From:      i,
		}, nil
	case *image.RGBA:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 4,
			Stride:    img.Stride,
			From:      i,
		}, nil
	case *image.RGBA64:
		return &Row{
			Pix:       img.Pix[img.PixOffset(0, y) : img.PixOffset(0, y)+img.Stride],
			PixelSize: 8,
			Stride:    img.Stride,
			From:      i,
		}, nil
	default:
		return &Row{}, CantDoThat
	}
}