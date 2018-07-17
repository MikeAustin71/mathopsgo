package main

import "../examples"

func main() {


	baseStr := "123456.789"
	shiftPlacesLeft := uint(3)
	expectedResult := "123.456789"

	examples.ExampleShiftPrecisionLeft(baseStr, shiftPlacesLeft, expectedResult)

}



