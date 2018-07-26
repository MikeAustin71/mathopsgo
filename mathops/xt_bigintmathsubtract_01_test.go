package mathops

import (
	"testing"
	"math/big"
)

func TestBigIntMathSubtract_SubtractBigInts_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result := BigIntMathSubtract{}.SubtractBigInts(
								bIMinuend,
									minuendPrecision,
										biSubtrahend,
											subtrahendPrecision)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigInts_02(t *testing.T) {
	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"
	expectedBigINumSign := 1
	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result := BigIntMathSubtract{}.SubtractBigInts(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigInts_03(t *testing.T) {
	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"
	expectedBigINumSign := -1
	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result := BigIntMathSubtract{}.SubtractBigInts(
		bIMinuend,
		minuendPrecision,
		biSubtrahend,
		subtrahendPrecision)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}
}

func TestBigIntMathSubtract_SubtractBigInts_04(t *testing.T) {
	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"
	expectedBigINumSign := 1

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
	minuendPrecision := minuendBiNum.GetPrecisionUint()
	biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
	subtrahendPrecision := subtrahendBiNum.GetPrecisionUint()

	result := BigIntMathSubtract{}.SubtractBigInts(
					bIMinuend,
						minuendPrecision,
							biSubtrahend,
								subtrahendPrecision)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNums_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNums_02(t *testing.T) {
	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"
	expectedBigINumSign := 1
	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNums_03(t *testing.T) {
	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"
	expectedBigINumSign := -1
	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}
}

func TestBigIntMathSubtract_SubtractBigIntNums_04(t *testing.T) {
	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"
	expectedBigINumSign := 1

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNums_05(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedResultStr := "99,999"


	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	subtrahendBiNum, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
	}

	result := BigIntMathSubtract{}.SubtractBigIntNums(minuendBiNum, subtrahendBiNum)

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected BigIntNum='%v'. Instead, BigIntNum= '%v'. ",
			expectedResultStr, actualResultStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_01(t *testing.T) {

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0:= "123.894000"
	subtrahend1:= "67.1"
	subtrahend2:= "93.0"
	subtrahend3:= "-124498.67158"
	subtrahend4:= "647129.57"
	subtrahend5:= "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"
	expectedBigINumSign := 1
	var err error

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
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

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_02(t *testing.T) {

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"


	subtrahend0:= "737.21"
	subtrahend1:= "9637591.879546"
	subtrahend2:= "28"
	subtrahend3:= "5284.9765"
	subtrahend4:= "-189291837.12"
	subtrahend5:= "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"
	expectedBigINumSign := 1

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
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

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
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
	expectedBigINumSign := 1

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
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

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
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
	expectedBigINumSign := -1

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
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

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_05(t *testing.T) {

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0:= "123.894000"
	subtrahend1:= "67.1"
	subtrahend2:= "93.0"
	subtrahend3:= "-124498.67158"
	subtrahend4:= "647129.57"
	subtrahend5:= "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998,231036"
	var err error


	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
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

	actualResultStr := result.GetNumStr()

	if expectedBigINumStr != actualResultStr {
		t.Errorf("Comparison Error: Expected result='%v'. Instead, result= '%v'. ",
			expectedBigINumStr, actualResultStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_01(t *testing.T) {

	var err error

	// minuend =   100
	minuendStr := "100"

	subtrahendStrs := []string{
				"5",
				"10",
				"30",
				"60.55",
				"-100.1",
				"-5.6",
	}

	expectedStrs := []string{
				"95",
				"90",
				"70",
				"39.45",
				"200.1",
				"105.6",
	}


	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]BigIntNum, lenSubtrahends)
	expectedResultsAry := make([]BigIntNum, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = BigIntNum{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = BigIntNum{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray :=
				BigIntMathSubtract{}.SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_02(t *testing.T) {

	var err error

	// minuend =   5051
	minuendStr := "5051"

	subtrahendStrs := []string{
				"8000",
				"6051.123456",
				"-30871.25",
				"604.55",
				"9100.123",
				"-115.76",
	}

	expectedStrs := []string{
				"-2949",
				"-1000.123456",
				"35922.25",
				"4446.45",
				"-4049.123",
				"5166.76",
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]BigIntNum, lenSubtrahends)
	expectedResultsAry := make([]BigIntNum, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = BigIntNum{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = BigIntNum{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray :=
				BigIntMathSubtract{}.SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_03(t *testing.T) {

	var err error

	// minuend =   -20051.974578
	minuendStr := "-20051.974578"

	subtrahendStrs := []string{
				"476.543798",
				"6051.123456",
				"-270871.25",
				"15604.5589321",
				"987100.123",
				"-114555.76",
	}

	expectedStrs := []string{
				"-20528.518376",
				"-26103.098034",
				"250819.275422",
				"-35656.5335101",
				"-1007152.097578",
				"94503.785422",
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]BigIntNum, lenSubtrahends)
	expectedResultsAry := make([]BigIntNum, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = BigIntNum{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = BigIntNum{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray :=
				BigIntMathSubtract{}.SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_04(t *testing.T) {

	var err error

	// minuend =   0
	minuendStr := "0"

	subtrahendStrs := []string{
				"476.543798",
				"6051.123456",
				"-270871.25",
				"15604.5589321",
				"987100.123",
				"-114555.76",
	}

	expectedStrs := []string{
							"-476.543798",
							"-6051.123456",
							"270871.25",
							"-15604.5589321",
							"-987100.123",
							"114555.76",
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]BigIntNum, lenSubtrahends)
	expectedResultsAry := make([]BigIntNum, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = BigIntNum{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = BigIntNum{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray :=
				BigIntMathSubtract{}.SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}


func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_05(t *testing.T) {

	var err error

	// minuend =   98.2
	minuendStr := "98.2"

	subtrahendStrs := []string{
				"0",
				"0.000",
				"0",
				"0",
				"0.00000",
				"0.0",
	}

	expectedStrs := []string{
							"98.2",
							"98.200",
							"98.2",
							"98.2",
							"98.20000",
							"98.2",
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]BigIntNum, lenSubtrahends)
	expectedResultsAry := make([]BigIntNum, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = BigIntNum{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = BigIntNum{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray :=
				BigIntMathSubtract{}.SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}
}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_06(t *testing.T) {

	var err error

	// minuend =   100
	minuendStr := "100"

	subtrahendStrs := []string{
		"5",
		"10",
		"30",
		"60.55",
		"-100.1",
		"-5.6",
	}

	expectedStrs := []string{
		"95",
		"90",
		"70",
		"39,45",
		"200,1",
		"105,6",
	}


	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]BigIntNum, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = BigIntNum{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

	}

	resultArray :=
		BigIntMathSubtract{}.SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

	for k:=0; k < lenSubtrahends; k++ {

		actualResultStr := resultArray[k].GetNumStr()

		if expectedStrs[k] != actualResultStr {
			t.Errorf("Error: Expected result='%v'. Instead, result='%v' Index='%v'",
				expectedStrs[k], actualResultStr, k)
		}

		actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. Index='%v'",
				expectedNumSeps.String(), actualNumSeps.String(), k)
		}

	}

}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_01(t *testing.T) {

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0:= "123.894000"
	subtrahend1:= "67.1"
	subtrahend2:= "93.0"
	subtrahend3:= "-124498.67158"
	subtrahend4:= "647129.57"
	subtrahend5:= "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"
	expectedBigINumSign := 1
	var err error

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
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


	result := BigIntMathSubtract{}.SubtractBigIntNumSeries(
			minuendBiNum,
			subtrahendAry[0],
		  subtrahendAry[1],
		  subtrahendAry[2],
		  subtrahendAry[3],
			subtrahendAry[4],
		  subtrahendAry[5])


	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_02(t *testing.T) {

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"


	subtrahend0:= "737.21"
	subtrahend1:= "9637591.879546"
	subtrahend2:= "28"
	subtrahend3:= "5284.9765"
	subtrahend4:= "-189291837.12"
	subtrahend5:= "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"
	expectedBigINumSign := 1

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
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

	result := BigIntMathSubtract{}.SubtractBigIntNumSeries(
																	minuendBiNum,
																	subtrahendAry[0],
																	subtrahendAry[1],
																	subtrahendAry[2],
																	subtrahendAry[3],
																	subtrahendAry[4],
																	subtrahendAry[5])


	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_03(t *testing.T) {

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
	expectedBigINumSign := 1

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
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


	result := BigIntMathSubtract{}.SubtractBigIntNumSeries(
											minuendBiNum,
											subtrahendAry[0],
											subtrahendAry[1],
											subtrahendAry[2],
											subtrahendAry[3],
											subtrahendAry[4],
											subtrahendAry[5])


	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_04(t *testing.T) {

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
	expectedBigINumSign := -1

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	minuendBiNum, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
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

	result := BigIntMathSubtract{}.SubtractBigIntNumSeries(
									minuendBiNum,
									subtrahendAry[0],
									subtrahendAry[1],
									subtrahendAry[2],
									subtrahendAry[3],
									subtrahendAry[4],
									subtrahendAry[5])

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimal_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

	if err != nil  {
		t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimal(decMinuend, decSubtrahend)

	if err != nil  {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimal(decMinuend, "+
			"decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
			minuendStr ,subtrahendStr, err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimal_02(t *testing.T) {
	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"
	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

	if err != nil  {
		t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimal(decMinuend, decSubtrahend)

	if err != nil  {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimal(decMinuend, "+
			"decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
			minuendStr ,subtrahendStr, err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimal_03(t *testing.T) {
	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"


	// subtrahend = 847129.876
	subtrahendStr := "847129.876"


	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"
	expectedBigINumSign := -1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

	if err != nil  {
		t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimal(decMinuend, decSubtrahend)

	if err != nil  {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimal(decMinuend, "+
			"decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
			minuendStr ,subtrahendStr, err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimal_04(t *testing.T) {
	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"
	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

	if err != nil  {
		t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimal(decMinuend, decSubtrahend)

	if err != nil  {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimal(decMinuend, "+
			"decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
			minuendStr ,subtrahendStr, err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimalArray_01(t *testing.T) {

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0:= "123.894000"
	subtrahend1:= "67.1"
	subtrahend2:= "93.0"
	subtrahend3:= "-124498.67158"
	subtrahend4:= "647129.57"
	subtrahend5:= "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"
	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}
	
	subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}
	
	subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}
	
	subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}


	result, err := BigIntMathSubtract{}.SubtractDecimalArray(decMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimalArray(" +
			"decMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimalArray_02(t *testing.T) {

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"

	subtrahend0:= "737.21"
	subtrahend1:= "9637591.879546"
	subtrahend2:= "28"
	subtrahend3:= "5284.9765"
	subtrahend4:= "-189291837.12"
	subtrahend5:= "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"

	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimalArray(decMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimalArray(" +
			"decMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimalArray_03(t *testing.T) {

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
	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimalArray(decMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimalArray(" +
			"decMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimalArray_04(t *testing.T) {

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
	expectedBigINumSign := -1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimalArray(decMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimalArray(" +
			"decMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}
