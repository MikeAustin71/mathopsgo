package examples

import (
	"../mathops"
	"fmt"
	"math/big"
)


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

	minuendStr := "-1718973642.1234567"

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

	subtrahend0:= "-28934682.721"
	subtrahend1:= "424.987654321"
	subtrahend2:= "-987"
	subtrahend3:= "62.94"
	subtrahend4:= "-999999999.99999"
	subtrahend5:= "-9638932.371"

	subtrahendAry := make([]mathops.BigIntNum, 6)
	// Confirmed Subtraction Result
	// ia Result5:    2,757,547,756.287792379
	// Array Result:  2,757,547,756.287792379

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
	fmt.Println("       bigISub0: ", bigISub0.GetNumStr())
	fmt.Println("        Result0: ", result.Result.GetNumStr())
	iaMinuend.SubtractFromThis(&iaSub0)
	fmt.Println("     ia Result0: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub1)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub1 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("       bigISub1: ", bigISub1.GetNumStr())
	fmt.Println("        Result1: ", result.Result.GetNumStr())
	fmt.Println("     ia Result1: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub2)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub2 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("       bigISub2: ", bigISub2.GetNumStr())
	fmt.Println("        Result2: ", result.Result.GetNumStr())
	fmt.Println("     ia Result2: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub3)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub3 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("       bigISub3: ", bigISub3.GetNumStr())
	fmt.Println("        Result3: ", result.Result.GetNumStr())
	fmt.Println("     ia Result3: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub4)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub4 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("       bigISub4: ", bigISub4.GetNumStr())
	fmt.Println("        Result4: ", result.Result.GetNumStr())
	fmt.Println("     ia Result4: ", iaMinuend.GetNumStr())

	iaMinuend.SubtractFromThis(&iaSub5)
	bPair = mathops.BigIntPair{}.NewBigIntNum(result.Result, bigISub5 )
	result = mathops.BigIntMathSubtract{}.SubtractPair(bPair)

	fmt.Println("       bigISub5: ", bigISub5.GetNumStr())
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


func ExampleBigIntAdd_01() {
	//n1Str := "0.000123"

	//n2Str := ".001"

	b1 := big.NewInt(123)
	b2 := big.NewInt(1)

	expectedResultStr:= "1123"
	expectedPrecision := uint(6)

	result := mathops.BigIntMathAdd{}.AddBigInts(b1, 6, b2, 3 )


	ExamplePrintBasicMathResult(expectedResultStr, expectedPrecision, result)

}

func ExampleBigIntAddNumStr_01() {
	n1Str := "0.000123"


	n2Str := "1"


	// Result = 	1.000123
	expectedResultStr := "1000123"
	expectedPrecision := uint(6)


	result, err := mathops.BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		fmt.Printf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
		return
	}

	ExamplePrintBasicMathResult(expectedResultStr, expectedPrecision, result)

	fmt.Println("Successful Completion")
}

func ExampleBigIntAddNumStr_02() {
	n1Str := "0.000123"


	n2Str := ".001"


	// Result = 	"0.001123"
	expectedResultStr := "001123"
	expectedPrecision := uint(6)


	result, err := mathops.BigIntMathAdd{}.AddNumStr(n1Str, n2Str)

	if err != nil {
		fmt.Printf("Error returned by BigIntMathAdd{}.AddNumStr(n1Str, n2Str). " +
			"Error='%v' ", err.Error())
		return
	}

	ExamplePrintBasicMathResult(expectedResultStr, expectedPrecision, result)

}

func ExamplePrintBasicMathResult(expectedResultStr string,
																	expectedPrecision uint,
																		result mathops.BigIntBasicMathResult) {

  getNumStr := result.Result.GetNumStr()

	fmt.Println("            expected result: ", expectedResultStr)
	fmt.Println("              result.BigInt: ", result.Result.BigInt.Text(10))
	fmt.Println("         result.GetNumStr(): ", getNumStr)
	fmt.Println("")
	fmt.Println("         Expected Precision: ", expectedPrecision)
	fmt.Println("           result.Precision: ", result.Result.Precision)
	fmt.Println("")
	fmt.Println("          result.Input.Big1: ", result.Input.Big1.BigInt.Text(10))
	fmt.Println("result.Input.Big1.Precision: ", result.Input.Big1.Precision)
	fmt.Println("")
	fmt.Println("          result.Input.Big2: ", result.Input.Big2.BigInt.Text(10))
	fmt.Println("result.Input.Big2.Precision: ", result.Input.Big2.Precision)


}
