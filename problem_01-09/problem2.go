package problem0109

import "fmt"

func Problem2() {
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

	fmt.Printf("Problem 2:\t%v\n", sum)
}