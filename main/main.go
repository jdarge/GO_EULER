package main

import (
	"fmt"
)



func main() {
	
	problem_1()
	problem_2()
	problem_3()
	
}

func problem_1() {
	var sum = 0

	for i := 1; i < 1000; i++ {
	
		if (i%3 == 0 || i%5 == 0) {
			sum += i
		}
	}

	fmt.Printf("%v\n", sum)
}

func problem_2() {
	var sum, tmp, f, s = 0, 0, 1, 2

	for {
		if(s > 4_000_000) {
			break
		}

		if(s%2==0) {
			sum+=s
		}
		
		tmp = f+s
		f=s
		s=tmp
	}

	fmt.Printf("%v\n", sum)
}

func abs(n int) int { // absolute value |x|
	if n >= 0 {
		return n
	}
	return -n
}

func gcd (a,b int) int { // greatest common divisor
	for b!=0 {
		a,b = b,a%b
	}
	return a
}

func rho_g(input, n int) int {
	return ((input*input + 1) % n)
}

func rho (n int) int {
	var x,y,d = 2,2,1

	for d==1 {
		x = rho_g(x, n)
		y = rho_g(rho_g(y, n), n)
		d = gcd(abs(x-y), n)
	}

	return d
}

func prime_factors(n int) ([]int, int) {

	var factors = []int{}
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

func problem_3() {

	var factors, index = prime_factors(600851475143)
	var largest_factor = factors[index-1]
	fmt.Printf("%d", largest_factor)
}