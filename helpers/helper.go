package helper

func ABS(n int) int { // absolute value |x|
	if n >= 0 {
		return n
	}
	return -n
}

func GCD(a, b int) int { // greatest common divisor
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int { // least common multiple
	return a * b / GCD(a, b)
}

func RhoG(input, n int) int {
	return (input*input + 1) % n
}

// https://en.wikipedia.org/wiki/Pollard's_rho_algorithm
func Rho(n int) int {
	var x, y, d = 2, 2, 1

	for d == 1 {
		x = RhoG(x, n)
		y = RhoG(RhoG(y, n), n)
		d = GCD(ABS(x-y), n)
	}

	return d
}

func PrimeFactors(n int) ([]int, int) {

	var factors []int
	var factor = 0
	var index = 0

	for n > 1 {
		factor = Rho(n)
		factors = append(factors, factor)
		n /= factor
		index++
	}

	return factors, index
}

func ReverseString(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}

func IsPerfectSquare(n int) bool {

	d, i := 1, n
	for i > 0 {
		i -= d
		d += 2
	}

	return i == 0
}

func SumOfNumbers(n int) int {
	// https://en.wikipedia.org/wiki/Triangular_number#Formula
	return n * (n + 1) / 2
	// return BinomialCoefficient(n+1, 2)
}

func SumOfSquares(n int) int {
	// https://www.cuemath.com/algebra/sum-of-squares/
	return n * (n + 1) * (2*n + 1) / 6
}

func CollatzChain(n int) int {

	length := 1

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		length++
	}

	return length
}

func FindLongestCollatzChain(start, end int) (int, int) {

	length, val := 0, 0

	for i := start; i < end; i++ {
		c := CollatzChain(i)
		if c > length {
			length = c
			val = i
		}
	}

	return length, val
}

func BinomialCoefficient(c int, k int) int {

	answer := 1

	if k > c/2 {
		k = c - k
	}

	for i := 1; i <= k; i++ {
		answer = (c - k + i) * answer / i
	}

	return answer
}
