package main

import (

	"fmt"
	"../mathops"
)

func main() {

	ExampleDecimalRelevantPrecision_01()

}

func ExampleDecimalRelevantPrecision_01() {
	str1 := "-2.0105000"
	expected := uint(4)
	d1, err := mathops.Decimal{}.NewNumStr(str1)

	if err != nil {
		fmt.Printf("Error returned by Decimal{}.NewNumStr(str1) " +
			"str1='%v' Error = '%v' \n",str1, err.Error())
		return
	}

	rP := d1.GetRelevantPrecision()

	if rP != expected {
		fmt.Printf("Expected Relevant Precision = %v. Instead, got %v\n",
				expected, rP)
		return
	}

	fmt.Println("Successful Completion")
}

func ExampleDecimalSetPrecisionRound_01() {

	str1 := "2.0105500"
	precision := uint(4)
	expected := "2.0106"

	d1, err := mathops.Decimal{}.NewNumStr(str1)

	if err != nil {
		fmt.Printf("Error returned by Decimal{}.NewNumStr(str1) " +
			"str1='%v' Error = '%v' \n", str1, err.Error())
		return
	}

	if str1 != d1.GetNumStr() {
		fmt.Printf("Expected NumStr= '%v'. Instead, got %v.\n", str1, d1.GetNumStr())
		return
	}

	d1.SetPrecisionRound(precision)

	if expected != d1.GetNumStr() {
		fmt.Printf("Expected 2nd NumStr= '%v'. Instead, got %v .\n", expected, d1.GetNumStr())
		return
	}

	fmt.Println("Successful Completion")

}

func ExampleNumStrThousands_01(){
	rawNumStr := "4613"
	expectedNStr := "4,613"
	ns, err := mathops.NumStrDto{}.NewNumStr(rawNumStr)

	if err != nil {
		fmt.Printf("Error returne by NewNumStr(rawNumStr). rawNumStr='%v' Error='%v'",
			rawNumStr, err.Error())
		return
	}

	thousandsNStr := ns.GetThouStr()

	fmt.Println("Thousands Delimited Number")
	fmt.Println("   Expected NStr: ", expectedNStr)
	fmt.Println("ns.GetThouStr() : ", thousandsNStr)

}

func ExampleNumStrCurrency_01(){
	rawNumStr := "-0.23"
	expectedNStr := "-$0.23"
	ns, err := mathops.NumStrDto{}.NewNumStr(rawNumStr)

	if err != nil {
		fmt.Printf("Error returne by NewNumStr(rawNumStr). rawNumStr='%v' Error='%v'",
			rawNumStr, err.Error())
		return
	}

	thousandsNStr := ns.GetCurrencyStr()

	fmt.Println("Thousands Currency Number")
	fmt.Println("     		Expected NStr: ", expectedNStr)
	fmt.Println("		ns.GetCurrencyStr(): ", thousandsNStr)
	fmt.Println("  Length of NStr: ", len(expectedNStr))
	fmt.Println("Length of OutStr: ", len(thousandsNStr))
}

func ExampleDecimal_04() {
	numStr1 := "71.0159"
	requestedPrecision := uint(4)
	expected := "71.0159"
	dec := mathops.Decimal{}

	d2, err := dec.NumStrPrecisionToDecimal(numStr1, requestedPrecision, true)

	if err != nil {
		fmt.Printf("Error returned by dec.NumStrPrecisionToDecimal(numStr1, " +
			"requestedPrecision, true). numStr1='%v', requestedPrecision='%v' Error='%v'",
			requestedPrecision, numStr1, err.Error())
		return
	}

	if expected != d2.GetNumStr() {
		fmt.Printf("Error: Expected NumStr='%v'. Instead, NumStr='%v'", expected, d2.GetNumStr())
		return
	}

	fmt.Println("NumStr: ", d2.GetNumStr())

	fmt.Println("Successful Completion!")
}


func ExampleDecimal_03() {
	numStr1 := "71.0159"
	requestedPrecision := uint(3)
	expected := "71.016"
	dec := mathops.Decimal{}

	d2, err := dec.NumStrPrecisionToDecimal(numStr1, requestedPrecision, true)

	if err != nil {
		fmt.Printf("Error returned by dec.NumStrPrecisionToDecimal(numStr1, " +
			"requestedPrecision, true). numStr1='%v', requestedPrecision='%v' Error='%v'",
			requestedPrecision, numStr1, err.Error())
		return
	}

	if expected != d2.GetNumStr() {
		fmt.Printf("Error: Expected NumStr='%v'. Instead, NumStr='%v'", expected, d2.GetNumStr())
		return
	}

	fmt.Println("NumStr: ", d2.GetNumStr())

	fmt.Println("Successful Completion!")
}


func ExampleDecimal_02() {
	numStr1 := "71.01"
	requestedPrecision := uint(3)
	expected := "71.010"
	dec := mathops.Decimal{}

	d2, err := dec.NumStrPrecisionToDecimal(numStr1, requestedPrecision, true)

	if err != nil {
		fmt.Printf("Error returned by dec.NumStrPrecisionToDecimal(numStr1, " +
			"requestedPrecision, true). numStr1='%v', requestedPrecision='%v' Error='%v'",
				requestedPrecision, numStr1, err.Error())
		return
	}

	if expected != d2.GetNumStr() {
		fmt.Printf("Error: Expected NumStr='%v'. Instead, NumStr='%v'", expected, d2.GetNumStr())
		return
	}

	fmt.Println("NumStr: ", d2.GetNumStr())

	fmt.Println("Successful Completion!")
}


func ExampleDecimal_01() {
	numStr1 := "35.50"
	numStr2 := "35.51"
	expected := "71.01"
	nu := mathops.NumStrUtility{}

	dec1, err := nu.ConvertNumStrToDecimal(numStr1)

	if err != nil {
		fmt.Printf("Received error from nu.ConvertNumStrToDecimal(numStr1). numStr1= '%v'. Error= %v\n", numStr1, err)
		return
	}

	dec2, _ := nu.ConvertNumStrToDecimal(numStr2)

	if err != nil {
		fmt.Printf("Received error from nu.ConvertNumStrToDecimal(numStr2). numStr2= '%v'. Error= %v", numStr2, err)
		return
	}

	dec3, err := dec1.Add(dec2)

	if err != nil {
		fmt.Printf("Received error from dec1.Add(dec2). Error= %v", err)
		return
	}

	if expected != dec3.GetNumStr() {
		fmt.Printf("Expected NumStrOut='%v'. Instead, got '%v'", expected, dec3.GetNumStr())
		return
	}

	if !dec3.GetIsValid() {
		fmt.Print("Expected dec3.isValid='true'. Instead, got 'false'")
		return
	}

	fmt.Println("Successful Completion!")
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