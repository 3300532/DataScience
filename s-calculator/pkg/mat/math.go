package mat

type Data struct {
	NumsInt   []int
	NumsFloat []float64
}

// SumInt func receives a slice of integers and returns the sum of these integers
func (d *Data) SumInt() int {
	var res int

	for _, el := range d.NumsInt {
		res += el
	}
	return res
}

// SumFloat64 func receives a slice of float64 and returns the sum of these float64 numbers
func (d *Data) SumFloat64() float64 {
	var res float64

	for _, el := range d.NumsFloat {
		res += el
	}
	return res
}
