package mathops

import (
	"testing"
	"math/big"
)

func TestBigIntMathSubtract_SubtractBigIntNums_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "12332"
	minuendPrecision := uint(2)

	// subtrahend = 23.321
	subtrahendStr := "23321"
	subtrahendPrecision := uint(3)

	// result = 99.999
	expectedBigINumStr := "99999"
	expectedBigINumPrecision := uint(3)
	expectedBigINumSign := 1

	bMinuend, oK := big.NewInt(0).SetString(minuendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(minuendStr, 10) " +
			"minuendStr='%v' ", minuendStr)
	}

	bSubtrahend, oK := big.NewInt(0).SetString(subtrahendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(subtrahendStr, 10) " +
			"subtrahendStr='%v' ", subtrahendStr)
	}

	expectedBigI, oK := big.NewInt(0).SetString(expectedBigINumStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(expectedBigINumStr, 10) " +
			"expectedBigINumStr='%v' ", expectedBigINumStr)
	}

	expectedBigINum := BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum := BigIntNum{}.NewBigInt(bMinuend, minuendPrecision)

	subtrahendBiNum := BigIntNum{}.NewBigInt(bSubtrahend, subtrahendPrecision)

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

}


func TestBigIntMathSubtract_SubtractBigIntNums_02(t *testing.T) {
	// minuend = 949321.6712
	minuendStr := "9493216712"
	minuendPrecision := uint(4)

	// subtrahend = 45678.21
	subtrahendStr := "4567821"
	subtrahendPrecision := uint(2)

	// result = 903643.4612
	expectedBigINumStr := "9036434612"
	expectedBigINumPrecision := uint(4)
	expectedBigINumSign := 1

	bMinuend, oK := big.NewInt(0).SetString(minuendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(minuendStr, 10) " +
			"minuendStr='%v' ", minuendStr)
	}

	bSubtrahend, oK := big.NewInt(0).SetString(subtrahendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(subtrahendStr, 10) " +
			"subtrahendStr='%v' ", subtrahendStr)
	}

	expectedBigI, oK := big.NewInt(0).SetString(expectedBigINumStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(expectedBigINumStr, 10) " +
			"expectedBigINumStr='%v' ", expectedBigINumStr)
	}

	expectedBigINum := BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum := BigIntNum{}.NewBigInt(bMinuend, minuendPrecision)

	subtrahendBiNum := BigIntNum{}.NewBigInt(bSubtrahend, subtrahendPrecision)

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNums_03(t *testing.T) {
	// minuend = -5876458.56789012
	minuendStr := "-587645856789012"
	minuendPrecision := uint(8)

	// subtrahend = 847129.876
	subtrahendStr := "847129876"
	subtrahendPrecision := uint(3)

	// result = -6723588.44389012
	expectedBigINumStr := "-672358844389012"
	expectedBigINumPrecision := uint(8)
	expectedBigINumSign := -1

	bMinuend, oK := big.NewInt(0).SetString(minuendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(minuendStr, 10) " +
			"minuendStr='%v' ", minuendStr)
	}

	bSubtrahend, oK := big.NewInt(0).SetString(subtrahendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(subtrahendStr, 10) " +
			"subtrahendStr='%v' ", subtrahendStr)
	}

	expectedBigI, oK := big.NewInt(0).SetString(expectedBigINumStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(expectedBigINumStr, 10) " +
			"expectedBigINumStr='%v' ", expectedBigINumStr)
	}

	expectedBigINum := BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum := BigIntNum{}.NewBigInt(bMinuend, minuendPrecision)

	subtrahendBiNum := BigIntNum{}.NewBigInt(bSubtrahend, subtrahendPrecision)

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNums_04(t *testing.T) {
	// minuend = -289.673849
	minuendStr := "-289673849"
	minuendPrecision := uint(6)

	// subtrahend = -14579.012
	subtrahendStr := "-14579012"
	subtrahendPrecision := uint(3)

	// result = 14289.338151
	expectedBigINumStr := "14289338151"
	expectedBigINumPrecision := uint(6)
	expectedBigINumSign := 1

	bMinuend, oK := big.NewInt(0).SetString(minuendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(minuendStr, 10) " +
			"minuendStr='%v' ", minuendStr)
	}

	bSubtrahend, oK := big.NewInt(0).SetString(subtrahendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(subtrahendStr, 10) " +
			"subtrahendStr='%v' ", subtrahendStr)
	}

	expectedBigI, oK := big.NewInt(0).SetString(expectedBigINumStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(expectedBigINumStr, 10) " +
			"expectedBigINumStr='%v' ", expectedBigINumStr)
	}

	expectedBigINum := BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum := BigIntNum{}.NewBigInt(bMinuend, minuendPrecision)

	subtrahendBiNum := BigIntNum{}.NewBigInt(bSubtrahend, subtrahendPrecision)

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

	if !expectedBigINum.Equal(result.Result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_01(t *testing.T) {

	// minuend = 7328941.123456
	minuendStr := "7328941123456"
	minuendPrecision := uint(6)


	subtrahend0:= "123.894000"
	subtrahend1:= "67.1"
	subtrahend2:= "93.0"
	subtrahend3:= "-124498.67158"
	subtrahend4:= "647129.57"
	subtrahend5:= "28"


	// result = 6805998.231036
	expectedBigINumStr := "6805998231036"
	expectedBigINumPrecision := uint(6)
	expectedBigINumSign := 1

	bMinuend, oK := big.NewInt(0).SetString(minuendStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(minuendStr, 10) " +
			"minuendStr='%v' ", minuendStr)
	}

	lenSubtrahends := 6
	var err error
	subtrahendAry := make([]BigIntNum, lenSubtrahends)
	
	subtrahendAry[0], err = BigIntNum{}.NewNumStr(subtrahend0)
	
	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend0). " +
			"Error='%v'. ", err.Error())
	}
	
	subtrahendAry[1], err = BigIntNum{}.NewNumStr(subtrahend1)
	
	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). " +
			"Error='%v'. ", err.Error())
	}
	
	subtrahendAry[2], err = BigIntNum{}.NewNumStr(subtrahend2)
	
	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend2). " +
			"Error='%v'. ", err.Error())
	}
	
	subtrahendAry[3], err = BigIntNum{}.NewNumStr(subtrahend3)
	
	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend3). " +
			"Error='%v'. ", err.Error())
	}
	
	subtrahendAry[4], err = BigIntNum{}.NewNumStr(subtrahend4)
	
	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend4). " +
			"Error='%v'. ", err.Error())
	}
	
	subtrahendAry[5], err = BigIntNum{}.NewNumStr(subtrahend5)
	
	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend5). " +
			"Error='%v'. ", err.Error())
	}
	
	expectedBigI, oK := big.NewInt(0).SetString(expectedBigINumStr, 10)

	if !oK {
		t.Errorf("Error returned by big.NewInt(0).SetString(expectedBigINumStr, 10) " +
			"expectedBigINumStr='%v' ", expectedBigINumStr)
	}

	expectedBigINum := BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum := BigIntNum{}.NewBigInt(bMinuend, minuendPrecision)

	result := BigIntMathSubtract{}.SubtractBigIntNumArray(minuendBiNum, subtrahendAry)


	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_02(t *testing.T) {

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

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"
	expectedBigINumPrecision := uint(8)
	expectedBigINumSign := 1

	minuendNDto, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr). " +
			"ninuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	bMinuend, err := minuendNDto.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by minuendNDto.GetBigInt(). Error='%v'",
			err.Error())
	}


	lenSubtrahends := 6
	subtrahendAry := make([]BigIntNum, lenSubtrahends)

	subtrahendAry[0], err = BigIntNum{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend0). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[1], err = BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[2], err = BigIntNum{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend2). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[3], err = BigIntNum{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend3). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[4], err = BigIntNum{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend4). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[5], err = BigIntNum{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend5). " +
			"Error='%v'. ", err.Error())
	}

  expectedNDto, err := NumStrDto{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(expectedBigINumStr). " +
			"expectedBigINumStr='%v' Error='%v'. ",
			expectedBigINumStr, err.Error())
	}

	expectedBigI, err := expectedNDto.GetBigInt()

	if err!=nil {
		t.Errorf("Error returned by  expectedNDto.GetBigInt() " +
			"Error='%v' ", err.Error())
	}

	expectedBigINum := BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum := BigIntNum{}.NewBigInt(bMinuend, minuendPrecision)

	result := BigIntMathSubtract{}.SubtractBigIntNumArray(minuendBiNum, subtrahendAry)


	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_03(t *testing.T) {

	var err error

	// minuend =   1,718,973,642.1234567
	minuendStr := "1718973642.1234567"


	subtrahend0:= "-28934682.721"
	subtrahend1:= "424.987654321"
	subtrahend2:= "-987"
	subtrahend3:= "62.94"
	subtrahend4:= "-999999999.99999"
	subtrahend5:= "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"
	expectedBigINumPrecision := uint(9)
	expectedBigINumSign := 1

	expectedNDto, err := NumStrDto{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(expectedBigINumStr). " +
			"expectedBigINumStr='%v' Error='%v'. ",
			expectedBigINumStr, err.Error())
	}

	expectedBigI, err := expectedNDto.GetBigInt()

	if err!=nil {
		t.Errorf("Error returned by  expectedNDto.GetBigInt() " +
			"Error='%v' ", err.Error())
	}

	expectedBigINum := BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr). " +
			"ninuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]BigIntNum, lenSubtrahends)

	subtrahendAry[0], err = BigIntNum{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend0). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[1], err = BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[2], err = BigIntNum{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend2). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[3], err = BigIntNum{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend3). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[4], err = BigIntNum{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend4). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[5], err = BigIntNum{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend5). " +
			"Error='%v'. ", err.Error())
	}


	result := BigIntMathSubtract{}.SubtractBigIntNumArray(minuendBiNum, subtrahendAry)


	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_04(t *testing.T) {

	var err error

	// minuend =   -1,718,973,642.1234567
	minuendStr := "-1718973642.1234567"


	subtrahend0:= "-28934682.721"
	subtrahend1:= "424.987654321"
	subtrahend2:= "-987"
	subtrahend3:= "62.94"
	subtrahend4:= "-999999999.99999"
	subtrahend5:= "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"
	expectedBigINumPrecision := uint(9)
	expectedBigINumSign := -1

	expectedNDto, err := NumStrDto{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(expectedBigINumStr). " +
			"expectedBigINumStr='%v' Error='%v'. ",
			expectedBigINumStr, err.Error())
	}

	expectedBigI, err := expectedNDto.GetBigInt()

	if err!=nil {
		t.Errorf("Error returned by  expectedNDto.GetBigInt() " +
			"Error='%v' ", err.Error())
	}

	expectedBigINum := BigIntNum{}.NewBigInt(expectedBigI, expectedBigINumPrecision)

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr). " +
			"ninuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]BigIntNum, lenSubtrahends)

	subtrahendAry[0], err = BigIntNum{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend0). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[1], err = BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[2], err = BigIntNum{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend2). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[3], err = BigIntNum{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend3). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[4], err = BigIntNum{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend4). " +
			"Error='%v'. ", err.Error())
	}

	subtrahendAry[5], err = BigIntNum{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend5). " +
			"Error='%v'. ", err.Error())
	}


	result := BigIntMathSubtract{}.SubtractBigIntNumArray(minuendBiNum, subtrahendAry)


	if expectedBigINum.BigInt.Cmp(result.Result.BigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.BigInt.Text(10), result.Result.BigInt.Text(10))
	}

	if expectedBigINumSign != result.Result.Sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.Result.Sign)
	}

}
