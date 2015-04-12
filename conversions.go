package imago

import (
	"image"
)

func ToFloat64(i image.Image) []float64 {
	return PixelsToFloat64(GetPix(i), PixelSize(i))
}

func PixelsToFloat64(Pix []uint8, PixelSize int) []float64 {
	plen := len(Pix)
	res := make([]float64, 0, plen/PixelSize)
	for i := 0; i < plen; i += PixelSize {
		var v uint64
		for n, j := 0, PixelSize-1; j >= 0; j-- {
			v += uint64(Pix[i+j]) << uint64(n*8)
			n++
		}
		res = append(res, float64(v))
	}
	return res
}

func ToComplex(i image.Image) []complex128 {
	return PixelsToComplex(GetPix(i), PixelSize(i))
}

func PixelsToComplex(Pix []uint8, PixelSize int) []complex128 {
	plen := len(Pix)
	res := make([]complex128, 0, plen/PixelSize)
	for i := 0; i < plen; i += PixelSize {
		var v uint64
		for n, j := 0, PixelSize-1; j >= 0; j-- {
			v += uint64(Pix[i+j]) << uint64(n*8)
			n++
		}
		res = append(res, complex(float64(v), 0))
	}
	return res
}
