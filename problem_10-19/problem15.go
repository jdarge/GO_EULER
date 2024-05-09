package problem1019

import helper "euler/helpers"

/*

LOGIC:
	In problem 15 we are tasked with going from the top left to bottom right of a square using only right and down
		movements
	In an NxN `matrix` we are required to move N times down and N time right to get to our destination
	In other words...
		( 40 ) binomial coefficient
		( 20 )
	Here 40 represents the total amount of moves and 20 represents the amount of times we can move in a single direction
	A binomial coefficient of
		( c )
		( k )
	... can be represented as (c!)/(k!(c-k)!)
	This is the logic used to find the answer

	https://en.wikipedia.org/wiki/Binomial_coefficient
	https://www.youtube.com/watch?v=AOsWph2FNLw

*/

func Problem15() int {

	answer := helper.BinomialCoefficient(40, 20)

	return answer
}
