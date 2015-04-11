package imago

//AddRow will add each Pixel values of r2 to r1.
func (r1 *Row) AddRow(r2 *Row) {
	for n, _ := range r1.Pix {
		r1.Pix[n] += r2.Pix[n]
	}
}
