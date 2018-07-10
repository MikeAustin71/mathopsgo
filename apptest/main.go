package main

import (
	"../examples"

	)

func main() {

	//                        12345648901234567890123456789012


	baseStr := "2"
	exponentStr := "4"
	expectedNumStr := "16"
	maxPrecision := 17
	minPrecision := 0

	examples.ExampleIntAryPwrByTwos_01(baseStr, exponentStr, expectedNumStr, minPrecision, maxPrecision)

	examples.ExampleBigIntNumPower_01(baseStr, exponentStr, expectedNumStr, uint(maxPrecision))
}



