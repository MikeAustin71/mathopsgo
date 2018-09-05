package examples

import (
	"../mathops"
	"fmt"
	"math/big"
)

func ExampleBigIntMagnitude_01(target *big.Int) {

	magnitude, err := mathops.BigIntMath{}.GetMagnitude(target)

	if err != nil {
		fmt.Printf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ", target.Text(10), err.Error())
		return
	}

	bigTen := big.NewInt(10)
	val := big.NewInt(0).Exp(bigTen, magnitude, nil)

	fmt.Println("   target= ", target.Text(10))
	fmt.Println("magnitude= ", magnitude.Text(10))
	fmt.Println("Ten Value= ", val.Text(10))

	bigTarget := mathops.BigIntNum{}.NewBigInt(target, 0)
	bigMagnitude := mathops.BigIntNum{}.NewBigInt(magnitude, 0)
	bigTenValue := mathops.BigIntNum{}.NewBigInt(val, 0)
	fmt.Println("========================================")
	fmt.Println("   Target Str: ", bigTarget.FormatThousandsStr(mathops.LEADMINUSNEGVALFMTMODE))
	fmt.Println("Ten Value Str: ", bigTenValue.FormatThousandsStr(mathops.LEADMINUSNEGVALFMTMODE))
	fmt.Println("    Magnitude: ", bigMagnitude.GetNumStr())

}

func ExampleBigIntExpTest_01(baseStr, exponentStr, expectedStr string) {

	bINumBase, err := mathops.BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return
	}

	bIBase, err := bINumBase.GetBigInt()

	if err != nil {
		fmt.Printf("Error returned by bINumBase.GetBigInt(). "+
			"bINumBase='%v' Error='%v' \n", bINumBase.GetNumStr(), err.Error())
		return
	}

	bINumExponent, err := mathops.BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return
	}

	bIExponent, err := bINumExponent.GetBigInt()

	if err != nil {
		fmt.Printf("Error returned by bINumExponent.GetBigInt(). "+
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
