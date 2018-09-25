package main

import (
	"../mathops"
	"fmt"
)

func main() {

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

}
