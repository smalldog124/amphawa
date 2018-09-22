package add

import (
	"strconv"
)

func Add(number1, number2 string) string {
	lengthNumber1 := len(number1) - 1
	var suffer int
	var minus int
	var result string
	for index := lengthNumber1; index >= 0; index-- {
		convertNumber1, _ := strconv.Atoi(string(number1[index]))
		convertNumber2, _ := strconv.Atoi(string(number2[index]))
		plus := convertNumber1 + convertNumber2 + suffer
		minus, suffer = isSuffer(plus)
		result = strconv.Itoa(minus) + result
	}
	if suffer == 1 {
		result = "1" + result
	}
	return result
}

func isSuffer(result int) (int, int) {
	minus := result - 10
	if minus >= 0 {
		return minus, 1
	}
	return result, 0
}
