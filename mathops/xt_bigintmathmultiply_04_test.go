package mathops

import "testing"

func TestBigIntMathMultiply_MultiplyNumStrs_01(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875.94572"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}.New()

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by .NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStr"+
			"(multiplierStr, multiplicandStr, expectedNumSeps) "+
			"multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrs_02(t *testing.T) {
	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedNumStr := "30987680500513189125.14259702468435"

	expectedNumSeps := NumericSeparatorDto{}

	expectedBigINumSign := 1

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStr"+
			"(multiplierStr, multiplicandStr, expectedNumSeps) "+
			"multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
	}

	expectedNumSeps.SetDefaultsIfEmpty()

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

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrs_03(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedNumStr := "-30987680500513189125.14259702468435"

	expectedNumSeps := NumericSeparatorDto{}.New()

	expectedBigINumSign := -1

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStr"+
			"(multiplierStr, multiplicandStr, expectedNumSeps) "+
			"multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrs_04(t *testing.T) {
	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedNumStr := "22197234145.3632"

	expectedNumSeps := NumericSeparatorDto{}.New()

	expectedBigINumSign := 1

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStr"+
			"(multiplierStr, multiplicandStr, expectedNumSeps) "+
			"multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrs_05(t *testing.T) {
	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0
	expectedNumStr := "0"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStr"+
			"(multiplierStr, multiplicandStr, expectedNumSeps) "+
			"multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrs_06(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875.94572"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	expectedBigINumSign := 1

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStr"+
			"(multiplierStr, multiplicandStr, expectedNumSeps) "+
			"multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathMultiply_MultiplyNumStrs_07(t *testing.T) {
	// multiplier = 123,32
	multiplierStr := "123,32"

	// multiplicand = 23,321
	multiplicandStr := "23,321"

	// product = 2875,94572
	expectedNumStr := "2875,94572"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	expectedBigINumSign := 1

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStr"+
			"(multiplierStr, multiplicandStr, expectedNumSeps) "+
			"multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathMultiply_MultiplyNumStrArray_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"

	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
	}

	// product = 128
	expectedNumStr := "128"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}.New()

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrArray("+
			"multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrArray_02(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664.26376575446779648
	expectedNumStr := "11995826664.26376575446779648"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}.New()

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v'  Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrArray("+
			"multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrArray_03(t *testing.T) {

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"
	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"5.8",
		"68.7",
		"3.1234567",
		"8.0",
		"11",
	}

	// product = 2212352.1767579232
	expectedNumStr := "2212352.1767579232"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}.New()

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v'  Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrArray("+
			"multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrArray_04(t *testing.T) {

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"
	// multiplicandStrs
	multiplicandStrs := []string{
		"1.879",
		"3.824",
		"21.756",
		"2.1234567",
		"6",
		"2",
	}

	// product = -20408.5138429311978576052224
	expectedNumStr := "-20408.5138429311978576052224"

	expectedNumSeps := NumericSeparatorDto{}.New()

	expectedBigINumSign := -1

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v'  Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrArray("+
			"multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrArray_05(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664.26376575446779648
	expectedNumStr := "11995826664.26376575446779648"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrArray("+
			"multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathMultiply_MultiplyNumStrArray_06(t *testing.T) {

	var err error

	// multiplier = 37,9876
	multiplierStr := "37,9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27,9",
		"48,123456",
		"59,48721",
		"-3",
		"19,1",
		"69",
	}

	// product = 11995826664,26376575446779648
	expectedNumStr := "11995826664,26376575446779648"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrArray("+
			"multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected expectedBigINum bigInt ='%s'. Instead, result bigInt = '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error:  Expected iaResultNumStr='%v' "+
			"Instead, actual NumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"

	// multiplicandStrs
	multiplicandStrs := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"2",
		"4",
		"6",
		"8",
		"10",
		"12",
	}

	expectedNumSeps := NumericSeparatorDto{}.New()

	result, err :=
		BigIntMathMultiply{}.MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrOutputToArray"+
			"(multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	lenArray := len(multiplicandStrs)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j] {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j])
		}
	}
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_02(t *testing.T) {

	var err error

	// multiplier = 8
	multiplierStr := "8"
	// multiplicandStrs
	multiplicandStrs := []string{
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"800.8",
		"-208",
		"31.392",
		"64",
		"42376.984",
		"-39.168",
	}

	expectedNumSeps := NumericSeparatorDto{}.New()

	result, err :=
		BigIntMathMultiply{}.MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrOutputToArray"+
			"(multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	lenArray := len(multiplicandStrs)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j] {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j])
		}

	}
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_03(t *testing.T) {

	var err error

	// multiplier = -31.2
	multiplierStr := "-31.2"
	// multiplicandStrs
	multiplicandStrs := []string{
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"-3123.12",
		"811.2",
		"-122.4288",
		"-249.6",
		"-165270.2376",
		"152.7552",
	}

	expectedNumSeps := NumericSeparatorDto{}.New()

	result, err :=
		BigIntMathMultiply{}.MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrOutputToArray"+
			"(multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	lenArray := len(multiplicandStrs)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j] {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j])
		}

	}
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_04(t *testing.T) {

	var err error

	// multiplier = 283
	multiplierStr := "283"
	// multiplicandStrs
	multiplicandStrs := []string{
		"0",
		"-26",
		"0",
		"8",
		"5297.123",
		"0",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"0",
		"-7358",
		"0",
		"2264",
		"1499085.809",
		"0",
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	result, err :=
		BigIntMathMultiply{}.MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrOutputToArray"+
			"(multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	lenArray := len(multiplicandStrs)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j] {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j])
		}

	}
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_05(t *testing.T) {

	var err error

	// multiplier = 0
	multiplierStr := "0"
	// multiplicandStrs
	multiplicandStrs := []string{
		"5",
		"-26",
		"9",
		"8",
		"5297.123",
		"37",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	result, err :=
		BigIntMathMultiply{}.MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrOutputToArray"+
			"(multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	lenArray := len(multiplicandStrs)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j] {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j])
		}

	}
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_06(t *testing.T) {

	var err error

	// multiplier = -31,2
	multiplierStr := "-31,2"
	// multiplicandStrs
	multiplicandStrs := []string{
		"100,1",
		"-26",
		"3,924",
		"8",
		"5297,123",
		"-4,896",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"-3123,12",
		"811,2",
		"-122,4288",
		"-249,6",
		"-165270,2376",
		"152,7552",
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'
	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	result, err :=
		BigIntMathMultiply{}.MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyNumStrOutputToArray"+
			"(multiplierStr, multiplicandStrs, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	lenArray := len(multiplicandStrs)

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j] {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j])
		}

	}
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"

	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
	}

	// product = 128
	expectedNumStr := "128"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}.New()

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrSeries(
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_02(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664.26376575446779648
	expectedNumStr := "11995826664.26376575446779648"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}.New()

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrSeries(
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}
	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_03(t *testing.T) {

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"
	// multiplicandStrs
	multiplicandStrs := []string{
		"2",
		"5.8",
		"68.7",
		"3.1234567",
		"8.0",
		"11",
	}

	// product = 2212352.1767579232
	expectedNumStr := "2212352.1767579232"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrSeries(
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	expectedNumSeps.SetDefaultsIfEmpty()

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_04(t *testing.T) {

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"
	// multiplicandStrs
	multiplicandStrs := []string{
		"1.879",
		"3.824",
		"21.756",
		"2.1234567",
		"6",
		"2",
	}

	// product = -20408.5138429311978576052224
	expectedNumStr := "-20408.5138429311978576052224"

	expectedBigINumSign := -1

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrSeries(
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathMultiply_MultiplyNumStrSeries_05(t *testing.T) {

	var err error

	// multiplier = -5,123456
	multiplierStr := "-5,123456"
	// multiplicandStrs
	multiplicandStrs := []string{
		"1,879",
		"3,824",
		"21,756",
		"2,1234567",
		"6",
		"2",
	}

	// product = -20408,5138429311978576052224
	expectedNumStr := "-20408,5138429311978576052224"

	expectedBigINumSign := -1

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'
	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	lenArray := len(multiplicandStrs)

	iaResult, err := IntAry{}.NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := IntAry{}.NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyNumStrSeries(
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}
