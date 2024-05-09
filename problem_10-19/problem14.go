package problem1019

import (
	"euler/helper"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// var mutex sync.Mutex

func Problem14() int {

	numWorkers := runtime.NumCPU()
	chunkSize := 1_000_000 / numWorkers

	wg.Add(numWorkers)
	results := make(chan struct {
		length int
		number int
	})

	// give each thread a range to run FindLongestCollatzChain()
	for i := 0; i < numWorkers; i++ {
		go func(start, end int) {
			length, number := helper.FindLongestCollatzChain(start, end)
			results <- struct {
				length int
				number int
			}{length, number}
			wg.Done()
		}(i*chunkSize+1, (i+1)*chunkSize)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	longestChain, longestNum := 0, 0
	for result := range results {
		if result.length > longestChain {
			longestChain = result.length
			longestNum = result.number
		}
	}

	return longestNum
}

/*
func Problem14() {

	value, chain := 0, 0
	for i:=1; i < 1_000_000; i++ {
		c:=CollatzChain(i)
		if c > chain {
			chain = c
			value = i
		}
	}
	fmt.Printf("Problem 14:\t%d\n", value)
}
*/
