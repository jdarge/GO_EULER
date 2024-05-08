package problem1019

import (
	"fmt"
	"math/big" // required to use "big integers" to avoid overflows
)

func BinomialCoefficient(c int, k int) uint64 {
	top := big.NewInt(1)
	bottom := big.NewInt(1)

	for i := 0; i < k; i++ {
		top.Mul(top, big.NewInt(int64(c-i)))
		bottom.Mul(bottom, big.NewInt(int64(i+1)))
	}

	result := new(big.Int).Div(top, bottom)

	if result.BitLen() <= 64 {
		return result.Uint64()
	}
	return 0
}

func Problem15() {
	answer := BinomialCoefficient(40, 20)
	fmt.Printf("Problem 15:\t%d\n", answer)
}
