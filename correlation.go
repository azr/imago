package imago

//CrossCorrelate Calculates the cross-correlation function between two vectors
func CrossCorrelateUInt8(a []uint8, b []uint8) []uint8 {
	N := len(a)
	M := min(len(a), len(b)) // M -> desired length output
	R := make([]uint8, M)
	for k := 0; k < M; k++ {
		sum := 0.0
		for i := 0; i < N-k; i++ {
			if float64(b[i]) != 0 {
				sum += float64(a[i+k]) * float64(b[i]) /
					float64(a[i+k]) * float64(a[i+k])
				// float64(N) * float64(N) ?
				// float64(b[i]) * float64(b[i]) ?
			}
		}
		R[k] = uint8(sum)
	}

	return R
}

func AutoCorrelateUInt8(pix []uint8) []uint8 {
	return CrossCorrelateUInt8(pix, pix)
}
