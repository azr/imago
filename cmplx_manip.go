package imago

func Complex128UpperBounds(in []complex128, stride, X, Y int) []complex128 {
	out := make([]complex128, X*Y)
	n := 0
	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			out[n] = in[x*stride+y]
			n++
		}
	}
	return out
}
