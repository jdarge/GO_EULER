package problem0109

import (
	"euler/helper"
)

func Problem6() int {

	var sos, son, answer int

	n := 100

	son = helper.SumOfNumbers(n)
	sos = helper.SumOfSquares(n)
	answer = (son * son) - sos

	return answer
}

func Problem6_original() int {

	sum, squares := 0, 0

	for i := 1; i <= 100; i++ {
		sum += i
		squares += i * i
	}

	answer := (sum * sum) - squares

	return answer
}
