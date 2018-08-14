package main

import (
	"fmt"
	"MikeAustin71/mathopsgo/mathops"
)

func main() {

	numStr := "0.000000034"
	mantissaLen := uint(2)
	expectedStr := "3.40e-8"


	bINum, _ := mathops.BigIntNum{}.NewNumStr(numStr)

	sciNotStr := bINum.GetSciNotationStr(mantissaLen)

	sciNotNum := bINum.GetSciNotationNumber()

	actualSciNotStr := sciNotNum.GetSciNotationStr(2)

	fmt.Println(" Sci Notation: ", sciNotStr)
	fmt.Println("     Expected: ", expectedStr)
	fmt.Println("  Sci Not Num: ", actualSciNotStr)

/*
	actualNumStr := bIBigNum.GetNumStr()
	sciNotStr := bIBigNum.GetSciNotationStr(maxPrecision)

	fmt.Println("                 nStr= ", nStr)
	fmt.Println("    bINum.GetNumStr()= ", actualNumStr)
	fmt.Println(" Scientific Notation = ", sciNotStr)
	fmt.Println("Expected Notation Str= ", expectedStr)

*/
}



