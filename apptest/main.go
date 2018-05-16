package main

import (
	"MikeAustin71/mathopsgo/mathops"
	"fmt"
	"math/big"
)

func main() {

	base := mathops.BigIntNum{}.NewBigInt(big.NewInt(8), 0)
	nthRoot := mathops.BigIntNum{}.NewBigInt(big.NewInt(3), 0)
	expectedNumStr := "2"
	maxPrecision := uint(30)

	ExampleBigIntNumNthRoot_01(base, nthRoot, maxPrecision, expectedNumStr)

}

func ExampleBigIntNumNthRoot_01(
				base, nthRoot mathops.BigIntNum,
					maxPrecision uint,
						expectedNumStr string) {

	mathNthRoot := mathops.BigIntMathNthRoot{}

	result, err := mathNthRoot.GetNthRootIntAry(base, nthRoot, maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by mathNthRoot.GetNthRootIntAry(base, nthRoot, maxPrecision). " +
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("*** BigIntMathNthRoot ***")
	fmt.Println("Expected Result: ", expectedNumStr)
	fmt.Println("  Actual Result: ", result.GetNumStr())
	fmt.Println("           Base: ", base.GetNumStr())
	fmt.Println("        NthRoot: ", nthRoot.GetNumStr())

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