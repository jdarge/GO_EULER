package main

import (
	"fmt"
)

func main() {

	problem1()
	problem2()
	problem3()

}

func problem1() {
	var sum = 0

	for i := 1; i < 1000; i++ {

		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Printf("%v\n", sum)
}

func problem2() {
	var sum, tmp, f, s = 0, 0, 1, 2

	for {
		if s > 4_000_000 {
			break
		}

		if s%2 == 0 {
			sum += s
		}

		tmp = f + s
		f = s
		s = tmp
	}

	fmt.Printf("%v\n", sum)
}

func abs(n int) int { // absolute value |x|
	if n >= 0 {
		return n
	}
	return -n
}

func gcd(a, b int) int { // greatest common divisor
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func rhoG(input, n int) int {
	return (input*input + 1) % n
}

// https://en.wikipedia.org/wiki/Pollard's_rho_algorithm
func rho(n int) int {
	var x, y, d = 2, 2, 1

	for d == 1 {
		x = rhoG(x, n)
		y = rhoG(rhoG(y, n), n)
		d = gcd(abs(x-y), n)
	}

	return d
}

func primeFactors(n int) ([]int, int) {

	var factors []int
	var factor = 0
	var index = 0

	for n > 1 {
		factor = rho(n)
		factors = append(factors, factor)
		n /= factor
		index++
	}

	return factors, index
}

func problem3() {

	var factors, index = primeFactors(600851475143)
	var largestFactor = factors[index-1]
	fmt.Printf("%d", largestFactor)
}
