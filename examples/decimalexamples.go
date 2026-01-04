package examples

import (
	"fmt"

	"github.com/mikeaustin71/mathops"
)

func ExampleDecPowInt01(baseStr string, exponent int, maxPrecision uint, expectedStr string) {

	ePrefix := "ExampleDecPowInt01"

	d1, err := new(mathops.Decimal).NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"d1, err := new(mathops.Decimal).NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			baseStr,
			err.Error())
		return
	}

	d2, err := d1.PowInt(exponent, maxPrecision)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"d2, err := d1.PowInt(exponent, maxPrecision)\n"+
			"exponent= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			exponent,
			maxPrecision,
			err.Error())
		return
	}

	actualResult, err := d2.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"actualResult, err := d2.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

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
