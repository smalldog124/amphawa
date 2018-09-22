package add

import "testing"

func Test_Add_Input_Number1_Number2_Should_Be_1232(t *testing.T) {
	expected := "1232"
	number1 := "981"
	number2 := "251"

	actual := Add(number1, number2)

	if expected != actual {
		t.Errorf("expect %s but it got %v", expected, actual)
	}
}
func Test_Add_Input_Number1_Number2_64(t *testing.T) {
	expected := "64"
	number1 := "48"
	number2 := "16"

	actual := Add(number1, number2)

	if expected != actual {
		t.Errorf("expect %s but it got %v", expected, actual)
	}
}
func Test_Add_Input_Number1_Number2_xxx(t *testing.T) {
	expected := "xxx"
	number1 := "995144900824012405238259418456758068524921"
	number2 := "879161138176194604342893336019619756243251"

	actual := Add(number1, number2)

	if expected != actual {
		t.Errorf("expect %s but it got %v", expected, actual)
	}
}
