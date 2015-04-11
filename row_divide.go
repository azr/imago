package imago

//AddRow will divide each Pixel values by divider.
func (r *Row) DivideValues(divider float64) {
	for n, pix := range r.Pix {
		r.Pix[n] = byte(float64(pix) / divider)
	}
}
