package main

import (

	"fmt"
	"../mathops"
)

func main() {

	ExampleSubtraction_02()
	ExampleSubtraction_01()

}

func ExampleSubtraction_01() {

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"
	minuendPrecision := uint(7)


	subtrahend0:= "737.21"
	subtrahend1:= "9637591.879546"
	subtrahend2:= "28"
	subtrahend3:= "5284.9765"
	subtrahend4:= "-189291837.12"
	subtrahend5:= "7638932.12398765"

	// result = 15,3035,620.80650965
	expectedBigINumStr := "153035620.80650965"
	expectedBigINumPrecision := uint(8)
	expectedBigINumSign := 1

	minuendNDto, err := mathops.NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		fmt.Printf("Error returned by NumStrDto{}.NewNumStr(minuendStr). " +
			"ninuendStr='%v' Error='%v' ", minuendStr, err.Error())
		return
	}

	fmt.Println("Minuend NumStrDto: ", minuendNDto.GetNumStr())

	bMinuend, err := minuendNDto.GetBigInt()

	if err != nil {
		fmt.Printf("Error returned by minuendNDto.GetBigInt(). Error='%v'",
			err.Error())
		return
	}

	fmt.Println(" Minuend BigIntNum: ", bMinuend.Text(10))


	lenSubtrahends := 6
	subtrahendAry := make([]mathops.BigIntNum, lenSubtrahends)

	subtrahendAry[0], err = mathops.BigIntNum{}.NewNumStr(subtrahend0)

	if err != nil {
		fmt.Printf("Error returned from BigIntNum{}.NewNumStr(subtrahend0). " +
			"Error='%v'. ", err.Error())
		return
	}

	subtrahendAry[1], err = mathops.BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		fmt.Printf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). " +
			"Error='%v'. ", err.Error())
		return
	}

	subtrahendAry[2], err = mathops.BigIntNum{}.NewNumStr(subtrahend2)

	if err != nil {
		fmt.Printf("Error returned from BigIntNum{}.NewNumStr(subtrahend2). " +
			"Error='%v'. ", err.Error())
		return
	}

	subtrahendAry[3], err = mathops.BigIntNum{}.NewNumStr(subtrahend3)

	if err != nil {
		fmt.Printf("Error returned from BigIntNum{}.NewNumStr(subtrahend3). " +
			"Error='%v'. ", err.Error())
		return
	}

	subtrahendAry[4], err = mathops.BigIntNum{}.NewNumStr(subtrahend4)

	if err != nil {
		fmt.Printf("Error returned from BigIntNum{}.NewNumStr(subtrahend4). " +
			"Error='%v'. ", err.Error())
		return
	}

	subtrahendAry[5], err = mathops.BigIntNum{}.NewNumStr(subtrahend5)

	if err != nil {
		fmt.Printf("Error returned from BigIntNum{}.NewNumStr(subtrahend5). " +
			"Error='%v'. ", err.Error())
		return
	}

	expectedNDto, err := mathops.NumStrDto{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		fmt.Printf("Error returned from NumStrDto{}.NewNumStr(expectedBigINumStr). " +
			"expectedBigINumStr='%v' Error='%v'. ",
			expectedBigINumStr, err.Error())
		return
	}

	expectedBigI, err := expectedNDto.GetBigInt()

	if err!=nil {
		fmt.Printf("Error returned by  expectedNDto.GetBigInt() " +
			"Error='%v' ", err.Error())
		return
	}

	expectedBigINum := mathops.BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum := mathops.BigIntNum{}.NewBigInt(bMinuend, minuendPrecision)

	result := mathops.BigIntMathSubtract{}.SubtractBigIntNumArray(minuendBiNum, subtrahendAry)

	fmt.Println("Minuend: ", minuendBiNum.GetNumStr())

	for i:=0; i < lenSubtrahends; i++ {
		fmt.Printf("Subtrahend[%v]='%v' \n", i, subtrahendAry[i].GetNumStr())
	}

	fmt.Println("Result: ", result.Result.GetNumStr())

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		fmt.Printf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
		return
	}

	if expectedBigINumSign != result.Result.Sign {
		fmt.Printf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
		return
	}


}

func ExampleSubtraction_02() {

	minuendStr := "-18973642.1234567"
	//minuendPrecision := uint(6)

	iaMinuend, err := mathops.IntAry{}.NewNumStr(minuendStr)

	if err!=nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(minuendStr). \n" +
			"Error='%v' \n", err.Error())
		return
	}

	bigIMinuend, err := mathops.BigIntNum{}.NewNumStr(minuendStr)

	if err!=nil {
		fmt.Printf("Error returned by BigIntNum{}.NewNumStr(minuendStr). \n" +
			"Error='%v' \n", err.Error())
		return
	}

	subtrahend0:= "737.21"
	subtrahend1:= "9637591.879546"
	subtrahend2:= "28"
	subtrahend3:= "5284.9765"
	subtrahend4:= "-189291837.12"
	subtrahend5:= "7638932.12398765"

	subtrahendAry := make([]mathops.BigIntNum, 6)
	// Confirmed Subtraction Result
	//        Result5:  153035620.80650965
	//     ia Result5:  153035620.80650965

	iaSub0, err := mathops.IntAry{}.NewNumStr(subtrahend0)

	if err!=nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend0). \n" +
			"Error='%v' \n", err.Error())
		return
	}

	iaSub1, err := mathops.IntAry{}.NewNumStr(subtrahend1)

	if err!=nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend1). \n" +
			"Error='%v' \n", err.Error())
		return
	}
	
	iaSub2, err := mathops.IntAry{}.NewNumStr(subtrahend2)

	if err!=nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend2). \n" +
			"Error='%v' \n", err.Error())
		return
	}
	
	iaSub3, err := mathops.IntAry{}.NewNumStr(subtrahend3)

	if err!=nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend3). \n" +
			"Error='%v' \n", err.Error())
		return
	}
	
	iaSub4, err := mathops.IntAry{}.NewNumStr(subtrahend4)

	if err!=nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend4). \n" +
			"Error='%v' \n", err.Error())
		return
	}
	
	iaSub5, err := mathops.IntAry{}.NewNumStr(subtrahend5)

	if err!=nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend5). \n" +
			"Error='%v' \n", err.Error())
		return
	}
	
	
	bigISub0, err := mathops.BigIntNum{}.NewNumStr(subtrahend0)
	
	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend0). \n" +
			"subtrahend0='%v' Error='%v' \n", subtrahend0, err.Error())
		return
	}

	subtrahendAry[0] = bigISub0


	bigISub1, err := mathops.BigIntNum{}.NewNumStr(subtrahend1)
	
	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend1). \n" +
			"subtrahend1='%v' Error='%v' \n", subtrahend1, err.Error())
		return
	}

	subtrahendAry[1] = bigISub1


	bigISub2, err := mathops.BigIntNum{}.NewNumStr(subtrahend2)
	
	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend2). \n" +
			"subtrahend2='%v' Error='%v' \n", subtrahend2, err.Error())
		return
	}

	subtrahendAry[2] = bigISub2

	bigISub3, err := mathops.BigIntNum{}.NewNumStr(subtrahend3)
	
	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend3). \n" +
			"subtrahend3='%v' Error='%v' \n", subtrahend3, err.Error())
		return
	}

	subtrahendAry[3] = bigISub3

	bigISub4, err := mathops.BigIntNum{}.NewNumStr(subtrahend4)
	
	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend4). \n" +
			"subtrahend4='%v' Error='%v' \n", subtrahend4, err.Error())
		return
	}

	subtrahendAry[4] = bigISub4

	bigISub5, err := mathops.BigIntNum{}.NewNumStr(subtrahend5)
	
	if err != nil {
		fmt.Printf("Error returned by IntAry{}.NewNumStr(subtrahend5). \n" +
			"subtrahend5='%v' Error='%v' \n", subtrahend5, err.Error())
		return
	}

	subtrahendAry[5] = bigISub5

	bPair := mathops.BigIntPair{}.NewBigIntNum(bigIMinuend, bigISub0)

	result := mathops.BigIntMathSubtract{}.SubtractPair(bPair)
	
	fmt.Println("Original Minued: ", iaMinuend.GetNumStr())
	fmt.Println("        Result0: ", result.Result.GetNumStr())
	iaMinuend.SubtractFromThis(&iaSub0)
	fmt.Println("     ia Result0: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub1)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub1 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("        Result1: ", result.Result.GetNumStr())
	fmt.Println("     ia Result1: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub2)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub2 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("        Result2: ", result.Result.GetNumStr())
	fmt.Println("     ia Result2: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub3)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub3 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("        Result3: ", result.Result.GetNumStr())
	fmt.Println("     ia Result3: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub4)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub4 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("        Result4: ", result.Result.GetNumStr())
	fmt.Println("     ia Result4: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub5)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub5 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("        Result5: ", result.Result.GetNumStr())
	fmt.Println("     ia Result5: ", iaMinuend.GetNumStr())

	result = mathops.BigIntMathSubtract{}.SubtractBigIntNumArray(bigIMinuend, subtrahendAry)
	
	fmt.Println("   Array Result: ", result.Result.GetNumStr())

}

func ExampleBigIntRounding_01() {
	nStr := "0.000"
	expectedNumStr := "0.00"
	roundToDec := uint(2)

	bINum1, _ := mathops.BigIntNum{}.NewNumStr(nStr)

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	fmt.Println("Expected NumStr: ", expectedNumStr)
	fmt.Println("  Actual NumStr: ", actualNumStr)


}

func ExampleDecimalDivide_01() {
	// str1 / str2
	/*
	str1 := "575.63"
	str2 := "2014.123"
	ePrecision := 20
	expected := "0.28579684557497233287"
*/

	str1 := "975.69"
	str2 := "589.7654321"
	expected := "1.654369597"
	ePrecision := 9

	d1, err := mathops.Decimal{}.NewNumStr(str1)

	if err != nil {
		fmt.Printf("Error returned from Decimal{}.NewNumStr(str1). "+
			"str1='%v' Error='%v \n", str1, err.Error())
		return
	}

	d2, err := mathops.Decimal{}.NewNumStr(str2)

	if err != nil {
		fmt.Printf("Error returned from Decimal{}.NewNumStr(str2). "+
			"str2='%v' Error='%v \n", str2, err.Error())
		return
	}

	d3, err := d1.Divide(d2, ePrecision)

	if err != nil {
		fmt.Printf("Error returned from d1.Divide(d2, ePrecision). "+
			" Error='%v \n",  err.Error())
		return
	}

	actualResult := d3.GetNumStr()

	ia1, err := mathops.IntAry{}.NewNumStr(str1)

	if err != nil {
		fmt.Printf("Error returned from IntAry{}.NewNumStr(str1). "+
			"str1='%v' Error='%v \n", str1, err.Error())
		return
	}

	ia2, err := mathops.IntAry{}.NewNumStr(str2)

	if err != nil {
		fmt.Printf("Error returned from IntAry{}.NewNumStr(str2). "+
			"str2='%v' Error='%v \n", str2, err.Error())
		return
	}


	ia3, err := ia1.DivideThisBy(&ia2, 0, ePrecision )

	if err != nil {
		fmt.Printf("Error returned from ia1.DivideThisBy(&ia2, 0, 20 ). " +
			"Error='%v \n", err.Error())
		return
	}

	chkResult := ia3.GetNumStr()

	fmt.Println("			    Dividend: ", str1)
	fmt.Println(" 			   Divisor: ", str2)
	fmt.Println("	   Actual Result: ", actualResult)
	fmt.Println("	 Expected Result: ", expected)
	fmt.Println("	    Check Result: ", chkResult)

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