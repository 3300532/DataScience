package stat

// To be elaborated.....

// This function returns the Average of the given slice of numbers
func Average(data Data) float64 {
	var sum float64

	sum = data.SumFloat64()

	var res float64

	for _, el := range numbers {
		sum += el
	}
	res = sum / float64(len(numbers))
	return res
}
