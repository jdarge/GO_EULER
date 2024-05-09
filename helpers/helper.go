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

// https://en.wikipedia.org/wiki/Triangular_number#Formula
func SumOfNumbers(n int) int {

	return n * (n + 1) / 2
	// return BinomialCoefficient(n+1, 2)
}

// https://www.cuemath.com/algebra/sum-of-squares/
func SumOfSquares(n int) int {

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

func UnpreciseSqrt(n int) (result int) {

	if n <= 0 {
		return -1
	}

	if n == 1 {
		return n
	}

	left, right := 1, n
	result = 0

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == n {
			return mid
		} else if mid*mid < n {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1
		}
	}

	return
}

func NthPrime(n int) int {

	primes := SieveOfEratosthenes(n * n)

	return primes[n-1]
}

// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func SieveOfEratosthenes(limit int) []int {

	sieve := make([]bool, limit+1)
	primes := make([]int, 0, limit/2)

	for i := 2 * 2; i <= limit; i += 2 { // i = 2 was taken out so we can improve step performance
		sieve[i] = true // in the bulk of the loop by using j += 2*i
	}
	for i := 3; i <= limit; i++ {
		if !sieve[i] {
			primes = append(primes, i)
			for j := i * i; j <= limit; j += 2 * i {
				sieve[j] = true
			}
		}
	}
	/*
	   	for i := 2; i <= limit; i++ {
	           if !sieve[i] {
	               primes = append(primes, i)
	               for j := i*i; j <= limit; j += i {
	                   sieve[j] = true
	               }
	           }
	       }
	*/

	return primes
}
