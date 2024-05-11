package main

import (
	problem0109 "euler/problem_01-09"
	problem1019 "euler/problem_10-19"
	"fmt"
	"time"
)

func main() {

	problem0109Test()
	problem1019Test()

}

func problem0109Test() {

	problems := []struct {
		name     string
		problem  func() int
		expected int
	}{
		{"Problem 1", problem0109.Problem1, 233168},
		{"Problem 2", problem0109.Problem2, 4613732},
		{"Problem 3", problem0109.Problem3, 6857},
		{"Problem 4", problem0109.Problem4, 906609},
		{"Problem 5", problem0109.Problem5, 232792560},
		{"Problem 6", problem0109.Problem6, 25164150},
		// {"Problem 7", problem0109.Problem7, 104743},
		{"Problem 8", problem0109.Problem8, 23514624000},
		{"Problem 9", problem0109.Problem9, 31875000},
	}

	executeProblem(problems)
}

func problem1019Test() {

	problems := []struct {
		name     string
		problem  func() int
		expected int
	}{
		{"Problem 10", problem1019.Problem10, 142913828922},
		{"Problem 11", problem1019.Problem11, -1}, // TODO: not implemented
		{"Problem 12", problem1019.Problem12, 76576500},
		{"Problem 13", problem1019.Problem13, 5537376230},
		{"Problem 14", problem1019.Problem14, 837799},
		{"Problem 15", problem1019.Problem15, 137846528820},
		{"Problem 16", problem1019.Problem16, 1366},
	}

	executeProblem(problems)
}

func executeProblem(problems []struct {
	name     string
	problem  func() int
	expected int
}) {
	for _, problem := range problems {
		start := time.Now()
		answer := problem.problem()
		elapsed := time.Since(start).Seconds()

		var color string

		if elapsed > 0.1 {
			color = "\x1b[31m"
		} else if elapsed > 0.01 {
			color = "\x1b[38;2;255;165;0m"
		} else {
			color = "\x1b[0m"
		}

		fmt.Printf("%-30s%-30d ? %v\t(%s%.6f seconds\x1b[0m)\n", problem.name, answer, answer == problem.expected, color, elapsed)
	}
}
