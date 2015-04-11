package imago_test

import (
	"github.com/azr/imago"

	"math"

	"testing"
)

func TestDegToRadConversion(t *testing.T) {
	if imago.DegToRad*180 != math.Pi {
		t.Errorf("imago.DegToRad * 180 == ", imago.DegToRad*180)
	}
}

func TestRotate(t *testing.T) {
	i := get4x4IncrementalGray()

	i2 := get4x4IncrementalGray()

	imago.RotateDeg(i2, i, 90)
}
