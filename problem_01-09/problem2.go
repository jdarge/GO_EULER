package problem0109

func Problem2() int {

	var sum, f, s = 0, 1, 2

	for s <= 4_000_000 {

		if s%2 == 0 {
			sum += s
		}

		f, s = s, f+s
	}

	return sum
}
