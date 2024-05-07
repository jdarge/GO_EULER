package problem0109

import (
	"fmt"
)

// Problem9
/*
MAX = 3 + 4 + 5 = 12
12 - 3 - 4 = 5
if(9 + 16 == 25)
    answer = 3 * 4 * 5
*/
func Problem9() {
	MAX := 1000
	answer := 0
	for i := 1; i < MAX; i++ {
		for j := i + 1; j < MAX; j++ {
			k := MAX - i - j        // ensures that i + j + k = 1000
			if (i*i)+(j*j) == k*k { // check pythagorean triplet
				answer = i * j * k // finalize product answer
				break
			}
		}
	}
	fmt.Printf("Problem 9:\t%d\n", answer)
}

/*
func problem9_original() {
	answer := 0
	for i:=1; i < 1000; i++ {
		for j:=i+1; j < 1000; j++ {
			k:= (i*i) + (j*j)
			if(!isPerfectSquare(k)) {
				continue
			}
			if(i+j+int(math.Sqrt(float64(k))) == 1000) {
				answer = i * j * int(math.Sqrt(float64(k)))
				break
			}
		}
	}
	fmt.Printf("%d\n",answer)
}
*/
