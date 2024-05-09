package problem0109

import (
	"euler/helper"
)

func Problem3() int {

	var factors, index = helper.PrimeFactors(600851475143)
	var largestFactor = factors[index-1]
	return largestFactor
}
