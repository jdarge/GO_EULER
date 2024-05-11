package problem1019

import (
	"euler/helper"
)

func Problem12() int {

	// https://www.mathsisfun.com/prime-factorization.html

	maxDivisors := 500

	/*for i := 1; ; i++ {
		a := helper.SumOfNumbers(i)
		b := helper.DivisorsOfNumber(a)
		if len(b) > maxDivisors {
			return a
		}
	}*/

	/*	var list []int
		for i := 1; i < maxDivisors; i++ {
			tmp := helper.FindNumberWithNDivisors(i)
			fmt.Printf("\r%d", i)
			list = append(list, tmp)
		}

		var strArray []string
		for _, num := range list {
			strArray = append(strArray, strconv.Itoa(num))
		}

		// Join the string slice with newline characters
		outputStr := strings.Join(strArray, "\n")

		// Write the string to a file
		err := ioutil.WriteFile("analysis/output12.txt", []byte(outputStr), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return -1
		}*/

	answer := helper.FindNumberWithNDivisors(maxDivisors)
	return answer
}
