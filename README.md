# GO_EULER
## What is this repository?
This is a collection of potential solutions for Project Euler. I like to use euler as a way to force myself to learn and experiment with different languages.

## What's completed?
```
Problems 1 - 9
```

## What's next?
```
Problems 10 - 19
```

## How to compile?
You can modify driver.go to only run specific problems. Currently, it runs them all.
``` 
$ go build && ./euler
```
Expected output:
```
Problem 1:	233168
Problem 2:	4613732
Problem 3:	6857
Problem 4:	906609
Problem 5:	232792560
Problem 6:	25164150
Problem 7:	104759
Problem 8:	23514624000
Problem 9:	31875000
Problem 14:	837799
Problem 15:	137846528820
```

## How optimized are the solutions?
No idea... but they seem pretty efficient. The only slow ones seem to be problems that require the calculation of a large amount of primes.
```
$ go build && time ./euler
real	0m0.461s
user	0m0.663s
sys	0m0.016s
```