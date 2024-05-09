package problem0109

import (
	"euler/helper"
)

func Problem5() int {

	answer := 1

	for i := 1; i <= 20; i++ {
		answer = helper.LCM(answer, i)
	}

	return answer
}
