package examples

import (
	"../mathops"
	"fmt"
)

func ExampleDecPowInt_01(baseStr string, exponent int, maxPrecision uint, expectedStr string ) {
	d1, err := mathops.Decimal{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by Decimal{}.NewNumStr(numStrDto) " +
			"numStrDto='%v' Error = '%v' \n",baseStr, err.Error())
	}

	d2, err := d1.PowInt(exponent, maxPrecision)

	if err != nil {
		fmt.Printf("Error returned by d1.PowInt(exp, maxPrecision) " +
			"exp='%v' maxPrecision='%v' Error = '%v' \n",
			exponent, maxPrecision, err.Error())
	}

	actualResult := d2.GetNumStr()

	fmt.Println("   Decimal.PowInt()")
	fmt.Println("------------------------")
	if expectedStr == actualResult {
		fmt.Println("*** SUCCESS ***")
	} else {
		fmt.Println("@@@ FAILURE @@@")
	}
	fmt.Println("------------------------")
	fmt.Println("           Base: ", baseStr)
	fmt.Println("       Exponent: ", exponent)
	fmt.Println("   maxPrecision: ", maxPrecision)
	fmt.Println("  Actual Result: ", actualResult)
	fmt.Println("Expected Result: ", expectedStr)
}
