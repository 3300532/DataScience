package mat

type mat struct {
	NumsInt   []int
	NumsFloat []float64
}

func NewMath() mat {

	var matStruct mat

	return matStruct
}

// SumInt func receives a slice of integers and returns the sum of these integers
func (m *mat) sumInt() int {
	var res int

	for _, el := range m.NumsInt {
		res += el
	}
	return res
}

// SumFloat64 func receives a slice of float64 and returns the sum of these float64 numbers
func (m *mat) sumFloat64() float64 {
	var res float64

	for _, el := range m.NumsFloat {
		res += el
	}
	return res
}
