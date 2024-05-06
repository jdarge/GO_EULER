package main

import (
	"fmt"
	"strconv"
)

func problem6() {
	
}

func main() {

	//problem1()
	//problem2()
	//problem3()
	//problem4()
	//problem5()
	problem6()

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

func problem3() {

	var factors, index = primeFactors(600851475143)
	var largestFactor = factors[index-1]
	fmt.Printf("%d", largestFactor)
}

func problem4() {
	var answer = 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			x := i * j
			if x < answer {
				break
			}
			a := strconv.Itoa(x)
			b := reverseString(a)
			if a == b && x > answer {
				answer = x
			}
		}
	}
	fmt.Printf("%d\n", answer)
}

func problem5() {
	answer := 1
	for i := 1; i <= 20; i++ {
		answer = lcm(answer, i)
	}
	fmt.Printf("%d\n", answer)
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

func lcm(a, b int) int { // least common multiple
	return a * b / gcd(a, b)
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

func reverseString(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}
