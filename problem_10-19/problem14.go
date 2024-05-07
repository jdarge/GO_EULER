package problem1019

import "fmt"

func CollatzChain(n int, c int) int {
	// returns the amount of chains

	if(n <= 1) {
		return c
	}

	if(n%2==0) {
		return CollatzChain(n/2, c+1)
	} else {
		return CollatzChain(3*n+1, c+1)
	}
	
}

func Problem14() {

	value, chain := 0, 0
	for i:=1; i < 1_000_000; i++ {
		c:=CollatzChain(i, 1)
		if c > chain {
			chain = c 
			value = i
		}
	}
	fmt.Printf("Problem 14:\t%d\n", value)
}