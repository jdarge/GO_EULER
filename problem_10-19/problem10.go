package problem1019

import (
	"euler/helper"
)

func Problem10() int {

	var sum int

	primes := helper.SieveOfEratosthenes(2_000_000)

	for i := 0; i < len(primes); i++ {
		sum += primes[i]
	}

	return sum
}
