package problem1019

import (
	"euler/helper"
	"fmt"
)

func Problem12() int {

	maxDivisors := 500

	for i := 1; ; i++ {
		if i%1_000_000 == 0 {
			fmt.Printf("%d ", i)
		}
		a := helper.SumOfNumbers(i)
		b := helper.DivisorsOfNumber(a)
		if len(b) > maxDivisors {
			return a
		}
	}
}
