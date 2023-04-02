package stat

type stat struct {
	numsInt   []int
	sumInt    int
	numsFloat []float64
	sumFloat  float64
}

func NewStat() stat {

	var statStruct stat

	return statStruct
}

// This function returns the Average of the given slice of numbers
func (s *stat) Average() float64 {

	// s.sumFloat =

	var res float64

	for _, el := range numbers {
		sum += el
	}
	res = sum / float64(len(numbers))
	return res
}
