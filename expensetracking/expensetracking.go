package expensetracking

import (
	"math"
	"strconv"
	"strings"
)

func ExpenseAVGpreDay(orders string) float64 {
	listOrder := split(orders, "\n")
	var totalPrice int
	for _, value := range listOrder {
		orderDay := split(value, " ")

		for index := 1; index <= len(orderDay)-1; index++ {
			temp := orderDay[index]
			priceString := temp[1:]
			totalPrice = sum(priceString, totalPrice)
		}
	}
	avg := float64(totalPrice) / float64(len(listOrder))
	return math.Round(avg*100) / 100
}

func split(dataString, sysbol string) []string {
	return strings.Split(dataString, sysbol)
}

func sum(priceString string, totalPrice int) int {
	price, _ := strconv.Atoi(priceString)
	return price + totalPrice
}
