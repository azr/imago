package imago

import (
	"image"
)

func (r *Row) Heighten(to int) image.Image {
	rect := image.Rect(0, 0, r.Stride, to)
	p := r.Pix[:r.Stride*r.PixelSize]
	switch r.From.(type) {
	case *image.Alpha:
		img := image.Alpha{
			Stride: r.Stride,
			Rect:   rect,
		}
		for i := 0; i < to; i++ {
			img.Pix = append(img.Pix, p...)
		}
		return &img
	case *image.Alpha16:
		img := image.Alpha16{
			Stride: r.Stride,
			Rect:   rect,
		}
		for i := 0; i < to; i++ {
			img.Pix = append(img.Pix, p...)
		}
		return &img
	case *image.Gray:
		img := image.Gray{
			Stride: r.Stride,
			Rect:   rect,
		}
		for i := 0; i < to; i++ {
			img.Pix = append(img.Pix, p...)
		}
		return &img
	case *image.Gray16:
		img := image.Gray16{
			Stride: r.Stride,
			Rect:   rect,
		}
		for i := 0; i < to; i++ {
			img.Pix = append(img.Pix, p...)
		}
		return &img
	case *image.NRGBA:
		img := image.NRGBA{
			Stride: r.Stride,
			Rect:   rect,
		}
		for i := 0; i < to; i++ {
			img.Pix = append(img.Pix, p...)
		}
		return &img
	case *image.NRGBA64:
		img := image.NRGBA64{
			Stride: r.Stride,
			Rect:   rect,
		}
		for i := 0; i < to; i++ {
			img.Pix = append(img.Pix, p...)
		}
		return &img
	case *image.Paletted:
		img := image.Paletted{
			Stride: r.Stride,
			Rect:   rect,
		}
		for i := 0; i < to; i++ {
			img.Pix = append(img.Pix, p...)
		}
		return &img
	case *image.RGBA:
		img := image.RGBA{
			Stride: r.Stride,
			Rect:   rect,
		}
		for i := 0; i < to; i++ {
			img.Pix = append(img.Pix, p...)
		}
		return &img
	case *image.RGBA64:
		img := image.RGBA64{
			Stride: r.Stride,
			Rect:   rect,
		}
		for i := 0; i < to; i++ {
			img.Pix = append(img.Pix, p...)
		}
		return &img
	default:
		return nil
	}
}
