package problem0109

import (
	helper "euler/helpers"
	"fmt"
)

func Problem3() {

	var factors, index = helper.PrimeFactors(600851475143)
	var largestFactor = factors[index-1]
	fmt.Printf("Problem 3:\t%d\n", largestFactor)
}
