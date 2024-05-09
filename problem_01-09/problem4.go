package problem0109

import (
	helper "euler/helpers"
	"strconv"
)

func Problem4() int {

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

	return answer
}
