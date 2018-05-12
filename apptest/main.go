package main

import (
	"MikeAustin71/mathopsgo/mathops"
	"fmt"
)

func main() {

	baseStr := "-2"
	exponentStr := "3.8"
	expectedNumStr := "13.928809012737986226180320279676"
	maxPrecision := uint(30)

	ExampleBigIntNumPower_01(baseStr, exponentStr, expectedNumStr, maxPrecision)

}

func ExampleBigIntNumPower_01(baseStr, exponentStr, expectedStr string, maxPrecision uint) {

	bINumBase, err := mathops.BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return
	}

	bINumExponent, err := mathops.BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return
	}

	result, err := mathops.BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). " +
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
		return
	}

	fmt.Println("*** BigIntMathPower{}.Pwr() ***")
	fmt.Println("Expected Result: ", expectedStr)
	fmt.Println("  Actual Result: ", result.GetNumStr())
	fmt.Println("           Base: ", bINumBase.GetNumStr())
	fmt.Println("       Exponent: ", bINumExponent.GetNumStr())

}