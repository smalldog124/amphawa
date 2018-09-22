package sum

func Sum(numbers []uint64) uint64 {
	var result uint64
	for _, value := range numbers {
		result = result + value
	}
	return result
}
