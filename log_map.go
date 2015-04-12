package imago

import (
	"image"
	"math"
)

func LogMap(i image.Image) []float64 {
	return LogMapFloat64(ToFloat64(i))
}

func LogMapFloat64(fs []float64) []float64 {
	for i := 0; i < len(fs); i++ {
		fs[i] = math.Log(fs[i])
	}
	return fs
}
