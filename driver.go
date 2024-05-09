package main

import (
	problem0109 "euler/problem_01-09"
	problem1019 "euler/problem_10-19"
	"fmt"
)

func main() {

	problem0109_test()
	problem1019_test()

}

func problem0109_test() {

	problems := []struct {
		name    string
		problem func() int
	}{
		{"Problem 1", problem0109.Problem1},
		{"Problem 2", problem0109.Problem2},
		{"Problem 3", problem0109.Problem3},
		{"Problem 4", problem0109.Problem4},
		{"Problem 5", problem0109.Problem5},
		{"Problem 6", problem0109.Problem6},
		{"Problem 7", problem0109.Problem7},
		{"Problem 8", problem0109.Problem8},
		{"Problem 9", problem0109.Problem9},
	}

	for _, problem := range problems {
		answer := problem.problem()
		fmt.Printf("%s:\t%d\n", problem.name, answer)
	}
}

func problem1019_test() {

	problems := []struct {
		name    string
		problem func() int
	}{
		{"Problem 14", problem1019.Problem14},
		{"Problem 15", problem1019.Problem15},
	}

	for _, problem := range problems {
		answer := problem.problem()
		fmt.Printf("%s:\t%d\n", problem.name, answer)
	}
}
