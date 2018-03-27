package main

import (

	"fmt"
	"../mathops"
)

func main() {

	ExampleMathOps_01()

}

func ExampleMathOps_01() {
	nStr1 := "-457.3"

	n1IntAry, err := mathops.IntAry{}.NewNumStr(nStr1)

	if err != nil {
		fmt.Printf("Error returned by mathops.IntAry{}.NewNumStr(nStr1). " +
			"Error='%v' \n", err.Error())
	}

	nStr2 := "0"

	n2IntAry, err := mathops.IntAry{}.NewNumStr(nStr2)

	if err != nil {
		fmt.Printf("Error returned by mathops.IntAry{}.NewNumStr(nStr2). " +
			"Error='%v' \n", err.Error())
		return
	}

	nResultIntAry := mathops.IntAry{}

	expected := "0.0"
	/*
	nRunes := []rune("00")
	eIAry := []int{0, 0}
	lNRunes := len(nRunes)
	lEArray := len(eIAry)
	*/
	ePrecision := 1
	eSignVal := 1

	err = n1IntAry.Multiply(&n1IntAry, &n2IntAry, &nResultIntAry, 1, -1)

	if err!=nil {
		fmt.Printf("Error returned by n1IntAry.Multiply(&n1IntAry, &n2IntAry, &nResultIntAry, 1, -1) " +
			"Error='%v'\n", err.Error())
		return
	}

	if expected != nResultIntAry.GetNumStr() {
		fmt.Printf("Expected NumStr='%v'. Instead, NumStr='%v' \n",
			expected, nResultIntAry.GetNumStr())
		return
	}

	if ePrecision != nResultIntAry.GetPrecision() {
		fmt.Printf("Expected Precision='%v'. Instead, Precision='%v'",
			ePrecision, nResultIntAry.GetPrecision())
		return
	}

	if eSignVal != nResultIntAry.GetSign() {
		fmt.Printf("Expected Sign='%v'. Instead, Sign='%v' \n",
			eSignVal, nResultIntAry.GetSign())
	}

	fmt.Println("Successful Completion!")
}