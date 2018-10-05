package grade_test

import (
	"amphawa/grade"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GradingSummary_Input_Scor_Example_Should_Be_B_1_Dplus_1_F_3(t *testing.T) {
	scor := "25\n4\n57\n73\n1"
	expected := "A+ 0\nA 0\nB+ 0\nB 1\nC+ 0\nC 0\nD+ 1\nD 0\nF 3"
	actual := grade.GradingSummary(scor)

	assert.Equal(t, expected, actual, "they should be equal")
}
func Test_GradingSummary_Input_Scor_Should_Be_XXXX(t *testing.T) {
	scor, _ := ioutil.ReadFile("grad.txt")
	expected := "xxx"
	actual := grad.GradingSummary(string(scor))

	assert.Equal(t, expected, actual)
}
