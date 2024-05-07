package problem0109

import (
	helper "euler/helpers"
	"fmt"
)

func Problem5() {
	answer := 1
	for i := 1; i <= 20; i++ {
		answer = helper.LCM(answer, i)
	}
	fmt.Printf("Problem 5:\t%d\n", answer)
}
