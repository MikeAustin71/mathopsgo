package main

import "../examples"

func main() {


	baseStr := "123456.789"
	shiftPlacesLeft := uint(6)
	expectedResult := "123456789000"

	examples.ExampleShiftPrecisionRight(baseStr, shiftPlacesLeft, expectedResult)

}



