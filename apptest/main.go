package main

import (

	"fmt"
	"../mathops"
	"math/big"
)

func main() {

	ExampleNumStrDtoBigIntParse_02()

}

func ExampleNumStrDtoBigIntParse_02() {
	num1Str := "-123456789"
	precision := uint(15)

	bNum1, ok := big.NewInt(0).SetString(num1Str, 10)

	if !ok {
		fmt.Print("Error returned by big.NewInt(0).SetString(num1Str, 10).  ")
		return
	}

	nDto2, err :=  mathops.NumStrDto{}.ParseSignedBigInt(bNum1, precision)
	if err != nil {
		fmt.Printf("Error returned by nDto.ParseBigIntNum(bNum1). Error='%v' ",
			err.Error())
		return
	}


	fmt.Println("Original NumStr: ", num1Str)
	fmt.Println("   nDto2 NumStr: ", nDto2.GetNumStr())
	fmt.Println(" spec precision: ", precision)

}

func ExampleNumStrDtoBigIntNumParse_01() {
	num1Str := "0.000"

	bNum1, err := mathops.BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
		return
	}

	nDto2, err :=  mathops.NumStrDto{}.ParseBigIntNum(bNum1)
	if err != nil {
		fmt.Printf("Error returned by nDto.ParseBigIntNum(bNum1). Error='%v' ",
			err.Error())
		return
	}


	fmt.Println("Original NumStr: ", num1Str)
	fmt.Println("   nDto2 NumStr: ", nDto2.GetNumStr())

}

func ExampleRoundPrecision_01() {
	num1Str := "654.123456"
	expectedNumStr := "654.123"
	newPrecision := uint(3)

	bNum1, err := mathops.BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
		return
	}

	expectedNum, err := mathops.BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
		return
	}

	fmt.Println("Old BNum1: ", bNum1.GetNumStr())

	bNum1.RoundToDecPlace(newPrecision)

	fmt.Println("New BNum1: ", bNum1.GetNumStr())

	if !expectedNum.Equal(bNum1) {
		fmt.Printf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
		return
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		fmt.Printf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
		return
	}

}

func ExampleSetPrecision_01(){

	num1Str := "654.123456"
	expectedNumStr := "654.123"
	newPrecision := uint(3)

	bNum1, err := mathops.BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
		return
	}

	expectedNum, err := mathops.BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
		return
	}

	fmt.Println("Old BNum1: ", bNum1.GetNumStr())

	bNum1.SetPrecision(newPrecision)

	fmt.Println("New BNum1: ", bNum1.GetNumStr())

	if !expectedNum.Equal(bNum1) {
		fmt.Printf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
		return
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		fmt.Printf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
		return
	}

}

func ExampleTestRat_01(num64, denom64 int64, maxPrecision uint) {

	rNum := big.NewRat(num64, denom64)

	numerator := rNum.Num()
	denominator := rNum.Denom()
	fmt.Println("       rNum: ", rNum.FloatString(int(maxPrecision)))
	fmt.Println("  Numerator: ", numerator.Text(10))
	fmt.Println("Denominator: ", denominator.Text(10))

}

func ExampleDecimalConvertToNumStr_04() {

	nStrAry := []string{
		"35.50",
		"36.50",
		"5.5",
		"92.75",
	}

	expected := "170.25"

	d := mathops.Decimal{}.New()

	for i := 0; i < len(nStrAry); i++ {
		dx := mathops.Decimal{}
		dx.SetNumStr(nStrAry[i])
		d.AddToThis(dx)
	}

	if expected != d.GetNumStr() {
		fmt.Printf("Expected NumStrOut='%v'. Instead, got '%v'\n", expected, d.GetNumStr())
		return
	}

	fmt.Println("Successful Completion!")

}

func ExampleDecimalConvertToNumStr_03() {

	str := "123456.654321"

	nDto, err := mathops.NumStrDto{}.NewNumStr(str)

	if err != nil {
		fmt.Printf("Error from NumStrDto{}.NewNumStr(str). " +
			"str='%v' Error='%v' \n",
				str, err.Error())
		return
	}

	err = nDto.IsNumStrDtoValid("")

	if err != nil {
		fmt.Printf("nDto NumStrDto is INVALID! " +
			"Error='%v' ", err.Error())
		return
	}


	nDtoScaleFactor, err := nDto.GetScaleFactor()

	if err != nil {
		fmt.Printf("Error from nDto.GetScaleFactor(). " +
			"Error='%v' ", err.Error())
		return
	}

	fmt.Println("nDto Scale Factor: ", nDtoScaleFactor.Text(10))


	dec := mathops.Decimal{}.New()

	if err != nil {
		fmt.Printf("Error returned by dec.SetNumStr(str). Error='%v\n", err.Error())
		return
	}

	d2, err := dec.MakeDecimalFromNumStrDto(nDto)

	if err != nil {
		fmt.Printf("Error returned by dec.MakeDecimalFromNumStrDto(nDto). Error='%v\n", err.Error())
		return
	}

	scaleVal, err := d2.GetScaleVal()

	if err != nil {
		fmt.Printf("Error returned by d2.GetScaleVal(). Error='%v\n", err.Error())
		return
	}

	fmt.Println("d2.GetScaleFactor() ", scaleVal.Text(10) )
	fmt.Println("d2.GetNumStr() ", d2.GetNumStr())

	dec.CopyIn(d2)
	fmt.Println("dec.GetNumStr() ", dec.GetNumStr())
	decScaleVal, err := dec.GetScaleVal()

	if err != nil {
		fmt.Printf("Error returned by dec.GetScaleVal(). Error='%v\n", err.Error())
		return
	}

	fmt.Println("dec.GetScaleFactor() ", decScaleVal.Text(10))


}

func ExampleDecimalConvertToNumStr_02() {

	str := "123456.654321"

	dec := mathops.Decimal{}.New()

	err := dec.SetNumStr(str)

	if err != nil {
		fmt.Printf("Error returned by dec.SetNumStr(str). Error='%v\n", err.Error())
		return
	}

	scaleVal, err := dec.GetScaleVal()

	fmt.Println("dec.GetScaleFactor() ", scaleVal.Text(10) )
	fmt.Println("dec.GetNumStr() ", dec.GetNumStr())


}

func ExampleDecimalConvertToNumStr_01() {

	str := "123456.654321"
	sint := "123456654321"
	nsu := mathops.NumStrUtility{}
	dec, err := nsu.ConvertNumStrToDecimal(str)

	if err != nil {
		fmt.Printf("Error from nsu.ConvertNumStrToDecimal(str). str= '%v'. Error= %v", str, err.Error())
		return
	}

	if dec.GetNumStr() != str {
		fmt.Printf("Expected NumStrOut= '%v'. Instead, got %v", str, dec.GetNumStr())
		return
	}

	if dec.GetPrecision() != 6 {
		fmt.Printf("Expected precision = '6'. Instead, got %v", dec.GetPrecision())
		return
	}

	allDigits, _ := dec.GetSignedAllDigitsStr()

	if allDigits != sint {
		fmt.Printf("Expected nDto.SignedAllDigitsBigInt='%v'. Instead, got %v. ", sint, allDigits)
		return
	}

	_, accuracy, err := dec.GetFloat64()

	if accuracy.String() != "Exact" {
		fmt.Printf("Expected nDto.SignedFloat64Accuracy == 'Exact'. Instead, got %v.", accuracy.String())
		return
	}

	bf, _ := dec.GetBigFloat()
	s := fmt.Sprintf("%s", bf.Text('f', dec.GetPrecision()))

	if s != str {
		fmt.Printf("Expected nDto.SignedBigFloat='%v'. Instead, got %v. ", str, s)
		return

	}

	if !dec.GetIsValid() {
		fmt.Printf("Expected nDto.isValid == 'true'. Instead, got %v", dec.GetIsValid())
		return
	}

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
		fmt.Printf("Expected Relevant precision = %v. Instead, got %v\n",
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
		fmt.Printf("Expected precision='%v'. Instead, precision='%v'",
			ePrecision, nResultIntAry.GetPrecision())
		return
	}

	if eSignVal != nResultIntAry.GetSign() {
		fmt.Printf("Expected sign='%v'. Instead, sign='%v' \n",
			eSignVal, nResultIntAry.GetSign())
	}

	fmt.Println("Successful Completion!")
}