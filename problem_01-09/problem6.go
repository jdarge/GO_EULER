package problem0109

import (
	helper "euler/helpers"
	"fmt"
)

func Problem6() int {
	var tetrahedral, triangle, answer int

	n := 100

	triangle = helper.SumOfNumbers(n)
	tetrahedral = helper.SumOfSquares(n)
	answer = (triangle * triangle) - tetrahedral

	fmt.Printf("%v::\n%v::\n", triangle, tetrahedral)

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
