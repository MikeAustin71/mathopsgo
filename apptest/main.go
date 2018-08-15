package main

import (
	"MikeAustin71/mathopsgo/mathops"
	"fmt"
)

func main() {

	numStr := "345123"
	mantissaLen := uint(2)
	expectedStr := "3.45e+5"

	bINum, _ := mathops.BigIntNum{}.NewNumStr(numStr)

	sciNotNum, err := bINum.GetSciNotationNumber(mantissaLen)

	if err != nil {
		fmt.Printf("Error returned by bINum.GetSciNotationNumber(mantissaLen). " +
			"Error='%v'\n", err.Error())
		return
	}

	actualSciNotStr := sciNotNum.GetNumStr()
	actualSignificand := sciNotNum.GetSignificand()
	actualExponent := sciNotNum.GetExponent()

	fmt.Println("     Original bINum: ", bINum.GetNumStr())
	fmt.Println("           Expected: ", expectedStr)
	fmt.Println("        Sci Not Num: ", actualSciNotStr)
	fmt.Println("Sci Not Significand: ", actualSignificand.GetNumStr())
	fmt.Println("Sci Not    Exponent: ", actualExponent.GetNumStr())

	/*

	base := uint(987654321)
	exponent := 3
	expectedValue := "987654321.000"
	baseBigINum := mathops.BigIntNum{}.NewUintExponent(base, exponent)

	fmt.Println("       baseStr: ", base)
	fmt.Println("   baseBigINum: ", baseBigINum.GetNumStr())
	fmt.Println("Expected Value: ", expectedValue)


	baseBigINum.DivideByTenToPower(exponent)

	fmt.Println("      exponent: ", exponent)
	fmt.Println("     New Value: ", baseBigINum.GetNumStr())
	fmt.Println("Expected Value: ", expectedValue)





			actualNumStr := bIBigNum.GetNumStr()
			sciNotStr := bIBigNum.GetSciNotationStr(maxPrecision)

			fmt.Println("                 nStr= ", nStr)
			fmt.Println("    bINum.GetNumStr()= ", actualNumStr)
			fmt.Println(" Scientific Notation = ", sciNotStr)
			fmt.Println("Expected Notation Str= ", expectedStr)
			*/

}



