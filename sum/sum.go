package sum

func Sum(numbers []uint64) uint64 {
	var result uint64
	for _, value := range numbers {
		result = result + value
	}
	return result
}

func SumRecursion(numbers []uint64) uint64 {
	length := len(numbers) - 1
	for index := 0; index <= length; index++ {
		if index+1 <= length {
			numbers[index+1] = add(numbers[index], numbers[index+1])
		}
	}
	return numbers[length]
}
func add(number1, number2 uint64) uint64 {
	return number1 + number2
}
