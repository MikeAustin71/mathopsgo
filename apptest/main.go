package main

import (
	"../examples"

	)

func main() {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "-45.632"
	exponentStr := "-1.01579"
	//                      12345678901234567890123456789012345
	expectedNumStr := "-45.1"
	maxPrecision := 15
	minPrecision := 0

	examples.ExampleIntAryPwrByTwos_01(baseStr, exponentStr, expectedNumStr, minPrecision, maxPrecision)

	examples.ExampleIntAryMultiplyPower_01(baseStr, exponentStr, expectedNumStr, minPrecision, maxPrecision)

	examples.ExampleBigIntNumPower_01(baseStr, exponentStr, expectedNumStr, uint(maxPrecision))
}



