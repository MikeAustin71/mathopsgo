package main

import (
	"MikeAustin71/mathopsgo/mathops"
	"fmt"
)

func main() {

	numStr := "0.000003451234"
	//mantissaLen := uint(2)
	//expectedStr := "3.45e-6"

	nDto, _ := mathops.NumStrDto{}.NewNumStr(numStr)

	intRunes := nDto.GetAbsIntRunes()

	intLength := nDto.GetAbsIntRunesLength()

	fmt.Println(" intRunes: ", string(intRunes))
	fmt.Println("intLength: ", intLength)

	/*

	iaNum, _ := mathops.IntAry{}.NewNumStr(numStr)

	sciNotNum, err := iaNum.GetSciNotationNumber(mantissaLen)

	if err != nil {
		fmt.Printf("Error returned by iaNum.GetSciNotationNumber(mantissaLen). " +
			"Error='%v'\n", err.Error())
		return
	}

	actualSciNotStr := sciNotNum.GetNumStr()
	//actualSignificand := sciNotNum.GetSignificand()
	actualExponent := sciNotNum.GetExponent()

	//fmt.Println("     Original iaNum: ", iaNum.GetNumStr())
	fmt.Println("           Expected: ", expectedStr)
	fmt.Println("        Sci Not Num: ", actualSciNotStr)
	//fmt.Println("Sci Not Significand: ", actualSignificand.GetNumStr())
	fmt.Println("Sci Not    Exponent: ", actualExponent.GetNumStr())


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
			fmt.Println("    iaNum.GetNumStr()= ", actualNumStr)
			fmt.Println(" Scientific Notation = ", sciNotStr)
			fmt.Println("Expected Notation Str= ", expectedStr)
			*/

}



