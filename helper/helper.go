package helper

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

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

// Rho https://en.wikipedia.org/wiki/Pollard's_rho_algorithm
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

// SumOfNumbers https://en.wikipedia.org/wiki/Triangular_number#Formula
func SumOfNumbers(n int) int {

	return n * (n + 1) / 2
	// return BinomialCoefficient(n+1, 2)
}

// SumOfSquares https://www.cuemath.com/algebra/sum-of-squares/
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

// SieveOfEratosthenes https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func SieveOfEratosthenes(limit int) []int {

	sieve := make([]bool, limit+1)
	primes := make([]int, 0, limit/2)

	// i = 2 was taken out so we can improve step performance
	// in the bulk of the loop by using j += 2*i
	primes = append(primes, 2)
	for i := 2 * 2; i <= limit; i += 2 {
		sieve[i] = true
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

func ConvertStrTo1DIntArray(str string) []int {

	var numbers []int

	for _, digit := range str {
		num, err := strconv.Atoi(string(digit))
		if err != nil {
			return nil
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func LargestProductInSeries(numbers []int, size int) int {

	largestProduct := 0
	var tmp int

	for i := 0; i+size < len(numbers); i++ {
		tmp = 1
		for j := i; j < i+size; j++ {
			tmp *= numbers[j]
		}
		if tmp > largestProduct {
			largestProduct = tmp
		}
	}

	return largestProduct
}

func DivisorsOfNumber(n int) []int {

	var divisors []int

	sqrtN := IntSqrt(n)
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i) // Source of logic: https://stackoverflow.com/a/26753963
			}
		}
	}

	return divisors
}

func CountDivisors(n int) int {

	divisors := 1
	sqrtN := IntSqrt(n)
	for i := 2; i <= sqrtN; i++ {
		exponent := 0
		for n%i == 0 {
			n /= i
			exponent++
		}
		if exponent > 0 {
			divisors *= exponent + 1
		}
	}
	if n > 1 {
		divisors *= 2
	}
	return divisors
}

func FindNumberWithNDivisors(divisorCount int) int {
	var n int
	for i := 1; ; i++ {
		n = SumOfNumbers(i)
		if CountDivisors(n) > divisorCount {
			return n
		}
	}
}

func IntSqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

// ReadFileLBL reads a file line by line and returns the resulting string array.
func ReadFileLBL(file string) []string {

	// taken from https://gosamples.dev/read-file w/ extra error handling implemented

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var list []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}

// StripStringToFirstN // TODO
func StripStringToFirstN(s []string, n int) []string {

	var list []string

	for i := 0; i < len(s); i++ {
		line := s[i]
		lastN := line[:n]
		list = append(list, lastN)
	}

	return list
}

// StringArrayToIntArray // TODO
func StringArrayToIntArray(list []string) []int {

	intList := make([]int, len(list))
	for i, str := range list {
		num, _ := strconv.Atoi(str)
		intList[i] = num
	}

	return intList
}

// SumIntArray returns the sum of each value inside an integer array.
func SumIntArray(list []int) int {

	sum := 0
	for _, num := range list {
		sum += num
	}
	return sum
}

// NLengthMSBInt truncates an integer to a specified size by removing values from the LSB.
func NLengthMSBInt(sum, size int) int {

	sumStr := strconv.Itoa(sum)
	sumStr = sumStr[:size]
	truncatedSum, _ := strconv.Atoi(sumStr)

	return truncatedSum
}

// SumStringDigits takes in a string and adds up each character as if it were a number.
func SumStringDigits(result string) int {

	sum := 0
	for _, digit := range result {
		n := int(digit - '0')
		sum += n
	}
	return sum
}
