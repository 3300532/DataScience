package data

type data struct {
	numsInt   []int
	numsFloat []float64
}

var numsInt []int
var numsFloat []float64

func NewData() *data {

	mainData := data{
		numsInt:   numsInt,
		numsFloat: numsFloat,
	}

	return &mainData

}
