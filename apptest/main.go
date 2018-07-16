package main

import "../examples"

func main() {


	baseStr := "2.125"
	exponent := "-5"
	maxPrecision := uint(32)
	//                   12345678901234567890123456789012
	expectedResult := "0.02307838042845159759046157465153"

	//examples.ExampleDecPowInt_01(baseStr, exponent, maxPrecision, expectedResult)

	examples.ExampleBigIntNumPower_01(baseStr, exponent, expectedResult, maxPrecision)

}



