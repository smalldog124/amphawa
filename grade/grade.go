package grade

import (
	"fmt"
	"strconv"
	"strings"
)

func GradingSummary(scor string) string {
	var A, Aplus, B, Bplus, C, Cplus, D, Dplus, F int
	arrayScor := strings.Split(scor, "\n")
	for _, pont := range arrayScor {
		pontTypeInt, _ := strconv.Atoi(pont)
		if pontTypeInt >= 85 {
			Aplus++
		} else if pontTypeInt >= 80 {
			A++
		} else if pontTypeInt >= 75 {
			Bplus++
		} else if pontTypeInt >= 70 {
			B++
		} else if pontTypeInt >= 65 {
			Cplus++
		} else if pontTypeInt >= 60 {
			C++
		} else if pontTypeInt >= 55 {
			Dplus++
		} else if pontTypeInt >= 50 {
			D++
		} else {
			F++
		}
	}
	return fmt.Sprintf("A+ %d\nA %d\nB+ %d\nB %d\nC+ %d\nC %d\nD+ %d\nD %d\nF %d", Aplus, A, Bplus, B, Cplus, C, Dplus, D, F)

}
