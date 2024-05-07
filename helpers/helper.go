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

func IsPerfectSquare (n int) bool {

	d,i := 1, n
	for ; i > 0; {
		i -= d
		d += 2
	}

	return i == 0
}