package mat

type Data struct {
	NumsInt   []int
	NumsFloat []float64
}

func (d *Data) SumInt() int {
	var res int

	for _, el := range d.NumsInt {
		res += el
	}
	return res
}

func (d *Data) SumFlo64() float64 {
	var res float64

	for _, el := range d.NumsFloat {
		res += el
	}
	return res
}
