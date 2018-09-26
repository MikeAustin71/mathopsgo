package main

import (
	"MikeAustin71/mathopsgo/mathops"
	"fmt"
	"math/big"
)

func main() {

	num1 := big.NewInt(902)
	precision1 := uint(2)
	rotations := uint(3)

	BigIntMultiplyBy2(num1, rotations, precision1)

}

func BigIntDivideBy2(num1 *big.Int, rotations, precision1 uint) {

	fmt.Println("BigIntDivideBy2")

	fmt.Println("   num1 before: ", num1.Text(10))

	result := big.NewInt(0).Rsh(num1, rotations)

	fmt.Println(" num1 after: ", num1.Text(10))
	fmt.Println(" precision1: ", precision1)
	fmt.Println("  rotations: ", rotations)
	fmt.Println("  result: ", result.Text(10))


}

func BigIntMultiplyBy2(num1 *big.Int, rotations, precision1 uint) {

	fmt.Println("BigIntMultiplyBy2")
	fmt.Println("   num1 before: ", num1.Text(10))

	result := big.NewInt(0).Lsh(num1, rotations)

	fmt.Println("      num1 after: ", num1.Text(10))
	fmt.Println("       rotations: ", rotations)
	fmt.Println("          result: ", result.Text(10))
	bINum := mathops.BigIntNum{}.NewBigInt(result, precision1)
	fmt.Println("BigIntNum Result: ", bINum.GetNumStr())

}

func BigIntDivide(num1, num2 *big.Int, precision1, precision2 uint) {

	xRemainder := big.NewInt(0)
	quotient, remainder := big.NewInt(0).QuoRem(num1, num2, xRemainder)

	// productPrecision := precision1 + precision2

	//biNum := mathops.BigIntNum{}.NewBigInt(product, productPrecision)
	fmt.Println("              num1: ", num1.Text(10))
	fmt.Println("              num2: ", num2.Text(10))
	fmt.Println("          quotient: ", quotient.Text(10))
	fmt.Println("         remainder: ", remainder.Text(10))
	fmt.Println("        xRemainder: ", xRemainder.Text(10))
	//fmt.Println(" product precision: ", productPrecision)
	//fmt.Println(" BigIntNum Product: ", biNum.GetNumStr())
	fmt.Println("        precision1: ", precision1)
	fmt.Println("        precision2: ", precision2)

}

/*
  // Logarithm test code
	xNumInt := 1500000
	baseInt := 4
	maxPrecision := uint(15)
	expectedLogValue := "10.2582655350227"
	// closer expectedLogValue = "10.258265535022667"
	xNum := mathops.BigIntNum{}.NewInt(xNumInt, 0)
	base := mathops.BigIntNum{}.NewInt(baseInt, 0)

	logValue, err := mathops.BigIntMathLogarithms{}.LogBaseOfX(base, xNum, maxPrecision)

	if err != nil {
		fmt.Printf("Error: error returned from BigIntMathLogarithms{}.GetIntDigits(" +
			"base, xNum, maxPrecision) "+
			"base='%v' xNum='%v' Error='%v'",
			base.GetNumStr(), xNum.GetNumStr(), err.Error())

		return
	}

	fmt.Println("GetNextDecimalDigit() ")
	fmt.Println("          log Value: ", logValue.GetNumStr())
	fmt.Println(" Expected log Value: ", expectedLogValue)
	fmt.Println("               base: ", base.GetNumStr())
	fmt.Println("               xNum: ", xNum.GetNumStr())

 */