package expensetracking

import (
	"io/ioutil"
	"testing"
)

func Test_ExpenseAVGpreDay_Input_7_Days_Should_Be_Total_Price_185dot71(t *testing.T) {
	expected := 185.71
	file, _ := ioutil.ReadFile("dataex.txt")
	inputString := string(file)
	actual := ExpenseAVGpreDay(inputString)

	if expected != actual {
		t.Errorf("expect %f but it got %f", expected, actual)
	}
}

func Test_split_Input_String_And_Sysbol_Shoud_Be_Array(t *testing.T) {
	expected := 3
	inputString := "a\nb\nc"
	sysbol := "\n"

	actual := split(inputString, sysbol)

	if expected != len(actual) {
		t.Errorf("expect %d but it got %d", expected, len(actual))
	}
}
func Test_ExpenseAVGpreDay_Input_31_Days_Should_Be_Total_Price_xxx(t *testing.T) {
	expected := 0.00
	file, _ := ioutil.ReadFile("data.txt")
	inputString := string(file)
	actual := ExpenseAVGpreDay(inputString)

	if expected != actual {
		t.Errorf("expect %f but it got %f", expected, actual)
	}
}
