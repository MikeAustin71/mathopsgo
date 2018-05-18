package main

import (
	"MikeAustin71/mathopsgo/mathops"
	"fmt"
	"math/big"
)

func main() {

	baseStr := "0.027"
	nthRootStr := "3"
	maximumPrecision := uint(6)
	expectedStr := "0.300000"

	ExampleBigIntNumNthRoot_01(baseStr, nthRootStr, maximumPrecision, expectedStr)

}

func ExampleBigIntExpTest_01(baseStr, exponentStr, expectedStr string) {

	bINumBase, err := mathops.BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return
	}

	bIBase, err := bINumBase.GetBigInt()

	if err != nil {
		fmt.Printf("Error returned by bINumBase.GetBigInt(). " +
			"bINumBase='%v' Error='%v' \n", bINumBase.GetNumStr(), err.Error())
		return
	}


	bINumExponent, err := mathops.BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return
	}

	bIExponent, err := bINumExponent.GetBigInt()

	if err != nil {
		fmt.Printf("Error returned by bINumExponent.GetBigInt(). " +
			"bINumExponent='%v' Error='%v' \n", bINumExponent.GetNumStr(), err.Error())
		return
	}
	modM := big.NewInt(0)
	result := big.NewInt(0).Exp(bIBase, bIExponent, modM)

	fmt.Println("big.NewInt(0).Exp(...)")
	fmt.Println("-------------------------------")
	fmt.Println("    Base: ", bIBase.Text(10))
	fmt.Println("Exponent: ", bIExponent.Text(10))
	fmt.Println("-------------------------------")
	fmt.Println("Expected: ", expectedStr)
	fmt.Println("  Result: ", result.Text(10))
	fmt.Println("-------------------------------")
	fmt.Println("    ModM: ", modM.Text(10))
}

func ExampleBigIntNumNthRoot_01(
				baseStr, nthRootStr string,
					maxPrecision uint,
						expectedNumStr string) {


	base, err := mathops.BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return
	}


	nthRoot, err := mathops.BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(nthRootStr). " +
			"nthRootStr='%v' Error='%v' \n", nthRootStr, err.Error())
		return
	}


	mathNthRootOp := mathops.BigIntMathNthRoot{}

	result, err := mathNthRootOp.GetNthRoot(base, nthRoot, maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by mathNthRootOp.GetNthRoot(base, nthRoot, maxPrecision). " +
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