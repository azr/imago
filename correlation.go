package imago

import (
	"image"
)

//NormalizedCrossCorrUint8 calculates the normalized cross-correlation
//[]uint8 between two []uint8.
//
//To Avoid blue colors beeing cross correlated with red colors, PixelSize must be set to
//the size a pixel is in your []uint8.
func NormalizedCrossCorrUint8(a, b []uint8, PixelSize int) []uint8 {
	N := len(a)
	M := min(len(a), len(b)) // M -> desired length output
	R := make([]uint8, M)
	for k := 0; k < M; k += PixelSize {
		sum := make([]float64, PixelSize)
		for i := 0; i < N-k; i += PixelSize {
			for pix := 0; pix < PixelSize; pix++ {
				I := i + pix
				if float64(b[I]) != 0 {
					sum[pix] += float64(a[I+k]) * float64(b[I]) /
						float64(a[I+k]) * float64(a[I+k])
				}
			}
		}
		for pix := 0; pix < PixelSize; pix++ {
			R[k+pix] = uint8(sum[pix])
		}
	}

	return R
}

//NormalizedCrossCorrelation Calculates the normalized cross-correlation image
//between two images
func NormalizedCrossCorrelation(a, b image.Image) image.Image {
	switch imgA := a.(type) {
	case *image.Alpha:
		res := CloneType(imgA, imgA.Rect).(*image.Alpha)
		res.Pix = NormalizedCrossCorrUint8(imgA.Pix, b.(*image.Alpha).Pix, PixelSize(imgA))
		return res
	case *image.Alpha16:
		res := CloneType(imgA, imgA.Rect).(*image.Alpha16)
		res.Pix = NormalizedCrossCorrUint8(imgA.Pix, b.(*image.Alpha16).Pix, PixelSize(imgA))
		return res
	case *image.Gray:
		res := CloneType(imgA, imgA.Rect).(*image.Gray)
		res.Pix = NormalizedCrossCorrUint8(imgA.Pix, b.(*image.Gray).Pix, PixelSize(imgA))
		return res
	case *image.Gray16:
		res := CloneType(imgA, imgA.Rect).(*image.Gray16)
		res.Pix = NormalizedCrossCorrUint8(imgA.Pix, b.(*image.Gray16).Pix, PixelSize(imgA))
		return res
	case *image.NRGBA:
		res := CloneType(imgA, imgA.Rect).(*image.NRGBA)
		res.Pix = NormalizedCrossCorrUint8(imgA.Pix, b.(*image.NRGBA).Pix, PixelSize(imgA))
		return res
	case *image.NRGBA64:
		res := CloneType(imgA, imgA.Rect).(*image.NRGBA64)
		res.Pix = NormalizedCrossCorrUint8(imgA.Pix, b.(*image.NRGBA64).Pix, PixelSize(imgA))
		return res
	case *image.Paletted:
		res := CloneType(imgA, imgA.Rect).(*image.Paletted)
		res.Pix = NormalizedCrossCorrUint8(imgA.Pix, b.(*image.Paletted).Pix, PixelSize(imgA))
		return res
	case *image.RGBA:
		res := CloneType(imgA, imgA.Rect).(*image.RGBA)
		res.Pix = NormalizedCrossCorrUint8(imgA.Pix, b.(*image.RGBA).Pix, PixelSize(imgA))
		return res
	case *image.RGBA64:
		res := CloneType(imgA, imgA.Rect).(*image.RGBA64)
		res.Pix = NormalizedCrossCorrUint8(imgA.Pix, b.(*image.RGBA64).Pix, PixelSize(imgA))
		return res
	default:
		panic(UnknownImageType.Error())
	}
}

//NormalizedAutoCorrelation calculates normalized cross correlation of a signal with itself
func NormalizedAutoCorrelation(i image.Image) image.Image {
	return NormalizedCrossCorrelation(i, i)
}
