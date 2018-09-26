package main

import (
	"MikeAustin71/mathopsgo/mathops"
	"fmt"
	"math/big"
)

func main() {

	num1 := big.NewInt(10)
	precision1 := uint(0)
	num2 := big.NewInt(200)
	precision2 := uint(1)

	BigIntMultiply(num1, num2, precision1, precision2)

}

func BigIntMultiply(num1, num2 *big.Int, precision1, precision2 uint) {

	product := big.NewInt(0).Mul(num1, num2)
	productPrecision := precision1 + precision2
	biNum := mathops.BigIntNum{}.NewBigInt(product, productPrecision)
	fmt.Println("              num1: ", num1.Text(10))
	fmt.Println("              num2: ", num2.Text(10))
	fmt.Println("           product: ", product.Text(10))
	fmt.Println(" product precision: ", productPrecision)
	fmt.Println(" BigIntNum Product: ", biNum.GetNumStr())
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