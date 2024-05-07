package problem0109

import (
	helper "euler/helpers"
	"fmt"
	"strconv"
)

func Problem4() {
	var answer = 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			x := i * j
			if x < answer {
				break
			}
			a := strconv.Itoa(x)
			b := helper.ReverseString(a)
			if a == b && x > answer {
				answer = x
			}
		}
	}
	fmt.Printf("Problem 4:\t%d\n", answer)
}