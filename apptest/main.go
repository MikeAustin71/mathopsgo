package main

import (
	"MikeAustin71/mathopsgo/examples"
	"math/big"
)

func main() {

	/*
	bigFloat := big.NewFloat(float64(32.129))

	bigI, accuracy := bigFloat.Int(nil)

	fmt.Println("       bigI: ", bigI.Text(10))
	fmt.Println("   accuracy: ", accuracy)

	rat, ratAccuracy := bigFloat.Rat(nil)

	fmt.Println("         rat: ", rat.FloatString(2))
	fmt.Println("rat accuracy: ", ratAccuracy)
	*/

	bigFloat := big.NewFloat(float64(32.129))
	expectedNumStr := "32.129"
	maxPrecision := uint(3)

	examples.ExampleSetBigFloat_01(bigFloat, maxPrecision, expectedNumStr)
}



