package problem1019

import (
	"math/big"
	"strconv"
)

func Problem16() int {

	// Yeah... I used math/big for this solution.
	// I forgot to commit my original solution and now we're here.

	base := big.NewInt(2)
	exponent := big.NewInt(1000)
	result := new(big.Int).Exp(base, exponent, nil).String()

	sum := 0
	for _, digit := range result {
		n, _ := strconv.Atoi(string(digit))
		sum += n
	}

	return sum
}
