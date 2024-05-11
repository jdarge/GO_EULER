package problem1019

import (
	"euler/helper"
	"math/big"
)

func Problem16() int {

	// Yeah... I used math/big for this solution.
	// I forgot to commit my original solution and now we're here.

	base := big.NewInt(0).SetInt64(2)
	exponent := big.NewInt(0).SetInt64(1000)
	result := new(big.Int).Exp(base, exponent, nil).String()

	sum := helper.SumStringDigits(result)

	return sum
}
