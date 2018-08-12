package main

import (
	"fmt"
	"MikeAustin71/mathopsgo/mathops"
)

func main() {

	numStr := "1"

	expectedStr := "1.00e+10000000"
	maxPrecision:= uint(2)
	exponent := uint64(10000000)

	bINum, _ := mathops.BigIntNum{}.NewNumStr(numStr)

	bIBigNum := mathops.BigIntMathMultiply{}.MultiplyBigIntNumByTenToIntPower(bINum, exponent, 0)

	sciNotStr := bIBigNum.GetSciNotationStr(maxPrecision)

	fmt.Println(" Sci Notation: ", sciNotStr)
	fmt.Println("     Expected: ", expectedStr)

/*
	actualNumStr := bIBigNum.GetNumStr()
	sciNotStr := bIBigNum.GetSciNotationStr(maxPrecision)

	fmt.Println("                 nStr= ", nStr)
	fmt.Println("    bINum.GetNumStr()= ", actualNumStr)
	fmt.Println(" Scientific Notation = ", sciNotStr)
	fmt.Println("Expected Notation Str= ", expectedStr)

*/
}



