package examples

import (
	"fmt"
	"math/big"

	"github.com/mikeaustin71/mathops"
)

func ExampleBigIntMagnitude01(target *big.Int) {

	ePrefix := "bigIntExamples.ExampleBigIntMagnitude01"

	magnitude, err := new(mathops.BigIntMath).GetMagnitude(target)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"magnitude, err := new(mathops.BigIntMath).\n"+
			" GetMagnitude(target)\n"+
			"target= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			target.Text(10),
			err.Error())
		return
	}

	bigTen := big.NewInt(10)

	val := big.NewInt(0).Exp(bigTen, magnitude, nil)

	fmt.Println("   target= ", target.Text(10))
	fmt.Println("magnitude= ", magnitude.Text(10))
	fmt.Println("ten Value= ", val.Text(10))

	bigTarget, err := new(mathops.BigIntNum).NewBigInt(target, 0)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bigTarget, err := new(mathops.BigIntNum).\n"+
			"  NewBigInt(target, 0)\n"+
			"target= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			target.Text(10),
			err.Error())
		return
	}

	bigMagnitude, err := new(mathops.BigIntNum).NewBigInt(magnitude, 0)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bigMagnitude, err := new(mathops.BigIntNum).\n"+
			"  NewBigInt(magnitude, 0)\n"+
			"magnitude= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			magnitude.Text(10),
			err.Error())
		return
	}

	bigTenValue, err := new(mathops.BigIntNum).NewBigInt(val, 0)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bigTenValue, err := new(mathops.BigIntNum).\n"+
			"  NewBigInt(val, 0)\n"+
			"val= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			val.Text(10),
			err.Error())
		return
	}

	bigTargetFmtThousandsStr, err := bigTarget.FormatThousandsStr(mathops.LEADMINUSNEGVALFMTMODE)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bigTargetFmtThousandsStr, err := bigTarget.\n"+
			" FormatThousandsStr(mathops.LEADMINUSNEGVALFMTMODE)\n"+
			"mathops.LEADMINUSNEGVALFMTMODE= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			mathops.LEADMINUSNEGVALFMTMODE.String(),
			err.Error())
		return
	}

	bigTenValueFmtThousandsStr, err := bigTenValue.FormatThousandsStr(mathops.LEADMINUSNEGVALFMTMODE)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bigTenValueFmtThousandsStr, err := bigTenValue.\n"+
			" FormatThousandsStr(mathops.LEADMINUSNEGVALFMTMODE)\n"+
			"mathops.LEADMINUSNEGVALFMTMODE= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			mathops.LEADMINUSNEGVALFMTMODE.String(),
			err.Error())
		return
	}

	bigMagnitudeNumStr, err := bigMagnitude.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bigMagnitudeNumStr, err := bigMagnitude.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fmt.Println("========================================")
	fmt.Println("   Target Str: ", bigTargetFmtThousandsStr)
	fmt.Println("ten Value Str: ", bigTenValueFmtThousandsStr)
	fmt.Println("    Magnitude: ", bigMagnitudeNumStr)

}

func ExampleBigIntExpTest01(baseStr, exponentStr, expectedStr string) {

	ePrefix := "bigIntExamples.ExampleBigIntExpTest01"

	bINumBase, err := new(mathops.BigIntNum).NewNumStr(baseStr)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(mathops.BigIntNum).\n"+
			" NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			baseStr,
			err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	bIBaseBigInt, err := bINumBase.GetBigInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bIBaseBigInt, err := bINumBase.GetBigInt()\n"+
			"bINumBase= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			err.Error())
		return
	}

	bINumExponent, err := new(mathops.BigIntNum).NewNumStr(exponentStr)

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bINumExponent, err := new(mathops.BigIntNum).\n"+
			" NewNumStr(exponentStr)\n"+
			"exponentStr= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			exponentStr,
			err.Error())
		return
	}

	bINumExponentNumStr, err := bINumExponent.GetNumStr()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bIExponentNumStr, err :=  bINumExponent.GetNumStr()\n"+
			"Error='%v'\n\n",
			ePrefix, err.Error())
		return
	}

	bIExponentBigInt, err := bINumExponent.GetBigInt()

	if err != nil {
		fmt.Printf("%v\n"+
			"Error returned by:\n"+
			"bIExponentBigInt, err := bINumExponent.GetBigInt()\n"+
			"bINumExponent= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			bINumExponentNumStr,
			err.Error())
		return
	}

	modM := big.NewInt(0)
	result := big.NewInt(0).Exp(bIBaseBigInt, bIExponentBigInt, modM)

	fmt.Println("big.NewInt(0).Exp(...)")
	fmt.Println("-------------------------------")
	fmt.Println("    Base: ", bIBaseBigInt.Text(10))
	fmt.Println("Exponent: ", bIExponentBigInt.Text(10))
	fmt.Println("-------------------------------")
	fmt.Println("Expected: ", expectedStr)
	fmt.Println("  Result: ", result.Text(10))
	fmt.Println("-------------------------------")
	fmt.Println("    ModM: ", modM.Text(10))
}
