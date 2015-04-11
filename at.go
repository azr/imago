package imago

import (
	"image"
	"image/color"
)

//PixOffset gives you the beginning of a pixel from offset
func (r *Row) PixOffset(x int) int {
	return x * r.PixelSize
}

//At gives you a color from a position.
func (r *Row) At(x int) color.Color {
	i := r.PixOffset(x)
	if i > r.Stride-1 {
		panic("At: out of bounds")
	}
	switch img := r.From.(type) {
	case *image.Alpha:
		return color.Alpha{r.Pix[i]}
	case *image.Alpha16:
		return color.Alpha16{uint16(r.Pix[i+0])<<8 | uint16(r.Pix[i+1])}
	case *image.Gray:
		return color.Gray{r.Pix[i]}
	case *image.Gray16:
		return color.Gray16{uint16(r.Pix[i+0])<<8 | uint16(r.Pix[i+1])}
	case *image.NRGBA:
		return color.NRGBA{r.Pix[i+0], r.Pix[i+1], r.Pix[i+2], r.Pix[i+3]}
	case *image.NRGBA64:
		return color.NRGBA64{
			uint16(r.Pix[i+0])<<8 | uint16(r.Pix[i+1]),
			uint16(r.Pix[i+2])<<8 | uint16(r.Pix[i+3]),
			uint16(r.Pix[i+4])<<8 | uint16(r.Pix[i+5]),
			uint16(r.Pix[i+6])<<8 | uint16(r.Pix[i+7]),
		}
	case *image.Paletted:
		return img.Palette[r.Pix[i]]
	case *image.RGBA:
		return color.RGBA{r.Pix[i+0], r.Pix[i+1], r.Pix[i+2], r.Pix[i+3]}
	case *image.RGBA64:
		return color.RGBA64{
			uint16(r.Pix[i+0])<<8 | uint16(r.Pix[i+1]),
			uint16(r.Pix[i+2])<<8 | uint16(r.Pix[i+3]),
			uint16(r.Pix[i+4])<<8 | uint16(r.Pix[i+5]),
			uint16(r.Pix[i+6])<<8 | uint16(r.Pix[i+7]),
		}
	default:
		return nil
	}
}
