package problem1019

import (
	"euler/helper"
)

func Problem13() int {

	size := 10
	strList := helper.ReadFileLBL("problem_10-19/input13.txt")
	strList = helper.StripStringToFirstN(strList, size+1) // +1 b/c we need the extra data for accurate sum
	intList := helper.StringArrayToIntArray(strList)
	sum := helper.SumIntArray(intList)
	answer := helper.NLengthMSBInt(sum, size)

	return answer
}
