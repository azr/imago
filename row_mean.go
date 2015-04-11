package imago

//AverageRow creates a new row containing average Pix values pixel by pixel.
//
//Copies first row and sets Pix values into it.
//
//A uint32 array is used to not lose values of additions going over 255.
func AverageRow(rows ...*Row) *Row {
	nbRow := len(rows)
	if nbRow == 0 {
		return nil
	}
	if nbRow == 1 {
		return rows[0]
	}
	w := len(rows[0].Pix)
	buffer := make([]uint32, w)
	for _, row := range rows {
		for i := 0; i < w; i++ {
			buffer[i] += uint32(row.Pix[i])
		}
	}

	pix := make([]uint8, w)
	for n, v := range buffer {
		pix[n] = uint8(v / uint32(nbRow))
	}
	r := *rows[0]
	r.Pix = pix
	return &r
}

func (r1 *Row) MeanRow(r2 *Row) {
	for n, _ := range r1.Pix {
		r1.Pix[n] = uint8((float64(r1.Pix[n]) + float64(r2.Pix[n])) / 2)
	}
}
