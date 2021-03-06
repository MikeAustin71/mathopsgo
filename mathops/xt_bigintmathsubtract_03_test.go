package mathops

import "testing"

func TestBigIntMathSubtract_SubtractIntArySeries_01(t *testing.T) {

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"
	expectedBigINumSign := 1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = IntAry{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = IntAry{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = IntAry{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = IntAry{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntArySeries("+
			"iaMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractIntArySeries_02(t *testing.T) {

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"

	subtrahend0 := "737.21"
	subtrahend1 := "9637591.879546"
	subtrahend2 := "28"
	subtrahend3 := "5284.9765"
	subtrahend4 := "-189291837.12"
	subtrahend5 := "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"
	expectedBigINumSign := 1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = IntAry{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = IntAry{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = IntAry{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = IntAry{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntArySeries("+
			"iaMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractIntArySeries_03(t *testing.T) {

	var err error

	// minuend =   1,718,973,642.1234567
	minuendStr := "1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"
	expectedBigINumSign := 1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = IntAry{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = IntAry{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = IntAry{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = IntAry{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntArySeries("+
			"iaMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractIntArySeries_04(t *testing.T) {

	var err error

	// minuend =   -1,718,973,642.1234567
	minuendStr := "-1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"
	expectedBigINumSign := -1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = IntAry{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = IntAry{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = IntAry{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = IntAry{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntArySeries("+
			"iaMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractIntArySeries_05(t *testing.T) {

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedNumStr := "6805998,231036"

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = iaMinuend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iaMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = IntAry{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = IntAry{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = IntAry{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = IntAry{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntArySeries("+
			"iaMinuend, ...). Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathSubtract_SubtractINumMgr_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	nDtoSubtrahend, err := NumStrDto{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgr(&decMinuend, &nDtoSubtrahend)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgr(&decMinuend, "+
			"&nDtoSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
			minuendStr, subtrahendStr, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgr_02(t *testing.T) {
	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"
	expectedBigINumSign := 1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	bINumSubtrahend, err := BigIntNum{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgr(&iaMinuend, &bINumSubtrahend)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgr(iaMinuend, "+
			"bINumSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
			minuendStr, subtrahendStr, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgr_03(t *testing.T) {
	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"
	expectedBigINumSign := -1

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := IntAry{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgr(&nDtoMinuend, &iaSubtrahend)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgr(nDtoMinuend, "+
			"iaSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
			minuendStr, subtrahendStr, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgr_04(t *testing.T) {
	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"
	expectedBigINumSign := 1

	bINumMinuend, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgr(&bINumMinuend, &decSubtrahend)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgr(bINumMinuend, "+
			"decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
			minuendStr, subtrahendStr, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgr_05(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedNumStr := "99,999"

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = decMinuend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by decMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	nDtoSubtrahend, err := NumStrDto{}.NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgr(&decMinuend, &nDtoSubtrahend)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgr(&decMinuend, "+
			"&nDtoSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
			minuendStr, subtrahendStr, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr= '%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathSubtract_SubtractINumMgrArray_01(t *testing.T) {

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"
	expectedBigINumSign := 1

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]INumMgr, lenSubtrahends)

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[0] = &dec0

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[1] = &bINum1

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[2] = &ia2

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[3] = &nDto3

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[4] = &dec4

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	subtrahendAry[5] = &ia5

	result, err := BigIntMathSubtract{}.SubtractINumMgrArray(&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrArray("+
			"&nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrArray_02(t *testing.T) {

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"

	subtrahend0 := "737.21"
	subtrahend1 := "9637591.879546"
	subtrahend2 := "28"
	subtrahend3 := "5284.9765"
	subtrahend4 := "-189291837.12"
	subtrahend5 := "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"
	expectedBigINumSign := 1

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]INumMgr, lenSubtrahends)

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[0] = &dec0

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[1] = &bINum1

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[2] = &ia2

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[3] = &nDto3

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[4] = &dec4

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	subtrahendAry[5] = &ia5

	result, err := BigIntMathSubtract{}.SubtractINumMgrArray(&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrArray("+
			"&nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrArray_03(t *testing.T) {

	var err error

	// minuend =   1,718,973,642.1234567
	minuendStr := "1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"
	expectedBigINumSign := 1

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]INumMgr, lenSubtrahends)

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[0] = &dec0

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[1] = &bINum1

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[2] = &ia2

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[3] = &nDto3

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[4] = &dec4

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	subtrahendAry[5] = &ia5

	result, err := BigIntMathSubtract{}.SubtractINumMgrArray(&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrArray("+
			"&nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrArray_04(t *testing.T) {

	var err error

	// minuend =   -1,718,973,642.1234567
	minuendStr := "-1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"
	expectedBigINumSign := -1

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]INumMgr, lenSubtrahends)

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[0] = &dec0

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[1] = &bINum1

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[2] = &ia2

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[3] = &nDto3

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[4] = &dec4

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	subtrahendAry[5] = &ia5

	result, err := BigIntMathSubtract{}.SubtractINumMgrArray(&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrArray("+
			"&nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrArray_05(t *testing.T) {

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedNumStr := "6805998,231036"

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]INumMgr, lenSubtrahends)

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[0] = &dec0

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[1] = &bINum1

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[2] = &ia2

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[3] = &nDto3

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[4] = &dec4

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	subtrahendAry[5] = &ia5

	result, err := BigIntMathSubtract{}.SubtractINumMgrArray(&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrArray("+
			"&nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_01(t *testing.T) {

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

	minuendINumMgr, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by INumMgr{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]INumMgr, lenSubtrahends)
	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		dec, err := Decimal{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		subtrahendAry[i] = &dec

		dec2, err := Decimal{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

		expectedResultsAry[i] = &dec2
	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractINumMgrOutputToArray(&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractINumMgrOutputToArray"+
			"(minuendINumMgr, subtrahendAry) minuendINumMgr='%v'  Error='%v'. ",
			minuendINumMgr.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k].GetNumStr() != resultArray[k].GetNumStr() {
			t.Errorf("Inequality Error: Expected ResultsAry='%v'. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_02(t *testing.T) {

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

	minuendINumMgr, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]INumMgr, lenSubtrahends)
	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		nDto1, err := NumStrDto{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		subtrahendAry[i] = &nDto1

		nDto2, err := NumStrDto{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

		expectedResultsAry[i] = &nDto2

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractINumMgrOutputToArray(&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractINumMgrOutputToArray"+
			"(minuendINumMgr, subtrahendAry) minuendINumMgr='%v'  Error='%v'. ",
			minuendINumMgr.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k].GetNumStr() != resultArray[k].GetNumStr() {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_03(t *testing.T) {

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

	minuendINumMgr, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]INumMgr, lenSubtrahends)
	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		ia1, err := IntAry{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		subtrahendAry[i] = &ia1

		ia2, err := IntAry{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

		expectedResultsAry[i] = &ia2
	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractINumMgrOutputToArray(&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractINumMgrOutputToArray"+
			"(minuendINumMgr, subtrahendAry) minuendINumMgr='%v'  Error='%v'. ",
			minuendINumMgr.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k].GetNumStr() != resultArray[k].GetNumStr() {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}
}

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_04(t *testing.T) {

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

	minuendINumMgr, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]INumMgr, lenSubtrahends)
	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		dec1, err := Decimal{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		subtrahendAry[i] = &dec1

		dec2, err := Decimal{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

		expectedResultsAry[i] = &dec2

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractINumMgrOutputToArray(&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractINumMgrOutputToArray"+
			"(minuendINumMgr, subtrahendAry) minuendINumMgr='%v'  Error='%v'. ",
			minuendINumMgr.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k].GetNumStr() != resultArray[k].GetNumStr() {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_05(t *testing.T) {

	var err error

	// minuend =   0
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

	minuendINumMgr, err := BigIntNum{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]INumMgr, lenSubtrahends)
	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		nDto1, err := NumStrDto{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		subtrahendAry[i] = &nDto1

		nDto2, err := NumStrDto{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

		expectedResultsAry[i] = &nDto2

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractINumMgrOutputToArray(&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractINumMgrOutputToArray"+
			"(minuendINumMgr, subtrahendAry) minuendINumMgr='%v'  Error='%v'. ",
			minuendINumMgr.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k].GetNumStr() != resultArray[k].GetNumStr() {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}
	}
}

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_06(t *testing.T) {

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

	minuendINumMgr, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by INumMgr{}.NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = minuendINumMgr.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by minuendINumMgr.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]INumMgr, lenSubtrahends)
	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		dec, err := Decimal{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		subtrahendAry[i] = &dec

		dec2, err := Decimal{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

		expectedResultsAry[i] = &dec2
	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractINumMgrOutputToArray(&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractINumMgrOutputToArray"+
			"(minuendINumMgr, subtrahendAry) minuendINumMgr='%v'  Error='%v'. ",
			minuendINumMgr.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedStrs[k] != resultArray[k].GetNumStr() {
			t.Errorf("Inequality Error: Expected ResultsAry='%v'. Instead, ResultsAry='%v'. ",
				expectedStrs[k], resultArray[k].GetNumStr())
		}

		actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'.",
				expectedNumSeps.String(), actualNumSeps.String())
		}

	}

}

func TestBigIntMathSubtract_SubtractINumMgrSeries_01(t *testing.T) {

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"
	expectedBigINumSign := 1

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrSeries("+
			"&nDtoMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrSeries_02(t *testing.T) {

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"

	subtrahend0 := "737.21"
	subtrahend1 := "9637591.879546"
	subtrahend2 := "28"
	subtrahend3 := "5284.9765"
	subtrahend4 := "-189291837.12"
	subtrahend5 := "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"
	expectedBigINumSign := 1

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrSeries("+
			"&nDtoMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrSeries_03(t *testing.T) {

	var err error

	// minuend =   1,718,973,642.1234567
	minuendStr := "1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"
	expectedBigINumSign := 1

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrSeries("+
			"&nDtoMinuend, ... subtrahend). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrSeries_04(t *testing.T) {

	var err error

	// minuend =   -1,718,973,642.1234567
	minuendStr := "-1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"
	expectedBigINumSign := -1

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrSeries("+
			"&nDtoMinuend, ... subtrahend). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrSeries_05(t *testing.T) {

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedNumStr := "6805998,231036"

	nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	dec0, err := Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	bINum1, err := BigIntNum{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	nDto3, err := NumStrDto{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	dec4, err := Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	ia5, err := IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractINumMgrSeries("+
			"&nDtoMinuend, ...). Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}
