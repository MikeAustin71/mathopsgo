package mathops

import "testing"

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

	expectedNumSeps := new(NumericSeparatorDto).New()

	lenArray := len(multiplicandStrs)

	iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())
	}

	for i := 0; i < lenArray; i++ {

		ia, err := new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps) "+
			"expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
			expectedNumStr, expectedNumSeps.String(), err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	isEqualToResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"isEqualToResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'", err.Error())

		return
	}

	if !isEqualToResult {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))

		return
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))

		return
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)

		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'", err.Error())

		return
	}

	expectedNumStr, err = iaResult.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedNumStr, err = iaResult.GetNumStr()\n"+
			"Error='%v'", err.Error())

		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedNumStr, err = iaResult.GetNumStr()\n"+
			"Error='%v'", err.Error())

		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

	return
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

	expectedNumSeps := new(NumericSeparatorDto).New()

	lenArray := len(multiplicandStrs)

	iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
			"multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
			multiplierStr, expectedNumSeps.String(), err.Error())

		return
	}

	for i := 0; i < lenArray; i++ {

		ia, err := new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())

			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by\n"+
				"iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"i='%v'; multiplicandStrs[i]='%v'\nError='%v'\n",
				i, multiplicandStrs[i], err.Error())
			return
		}

	}

	expectedBigINum, err := new(BigIntNum).
		NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
			"expectedNumStr, expectedNumSeps, expectedNumSeps)\n"+
			"expectedNumStr='%v' expectedNumSeps='%v'\nError='%v'.\n",
			expectedNumStr, expectedNumSeps.String(), err.Error())

		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if err != nil {
		t.Errorf("Error returned by new(IntAry).MultiplyNumStrSeries(...)\n"+
			"Error='%v'", err.Error())
		return
	}

	expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'", err.Error())
		return
	}

	if !expectedBigINumEqualResult {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
		return
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
		return
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'", err.Error())
		return
	}

	expectedNumStr, err = iaResult.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedNumStr, err = iaResult.GetNumStr()\n"+
			"Error='%v'", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

	return
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

	iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
			"multiplierStr='%v'; expectedNumSeps='%v'\nError='%v'\n",
			multiplierStr, expectedNumSeps.String(), err.Error())
		return
	}

	var ia IntAry

	for i := 0; i < lenArray; i++ {

		ia, err = new(IntAry).
			NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
			return
		}

	}

	expectedBigINum, err := new(BigIntNum).
		NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"new(BigIntNum).NewNumStrWithNumSeps(\n"+
			"expectedNumStr, expectedNumSeps, expectedNumSeps)\n"+
			"expectedNumStr='%v'; expectedNumSeps='%v'\nError='%v'\n",
			expectedNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedBigINumEqualResult {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'\n. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
		return
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	expectedNumStr, err = iaResult.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedNumStr, err = iaResult.GetNumStr()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			expectedNumStr, actualNumStr)
		return
	}

	expectedNumSeps.SetDefaultsIfEmpty()

	actualNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

	return
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

	iaResult, err := new(IntAry).
		NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"new(IntAry).NewNumStrWithNumSeps("+
			"multiplierStr, expectedNumSeps)\n"+
			"multiplierStr='%v'; expectedNumSeps='%v'\nError='%v'\n",
			multiplierStr, expectedNumSeps.String(), err.Error())
		return
	}

	var ia IntAry

	for i := 0; i < lenArray; i++ {

		ia, err = new(IntAry).
			NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by:\n"+
				"new(IntAry).NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps)\n"+
				"multiplicandStrs[%v]='%v'; expectedNumSeps='%v'\nError='%v'\n",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by:\n"+
				"iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"i='%v'; multiplicandStrs[i]='%v'\nError='%v'\n",
				i, multiplicandStrs[i], err.Error())
			return
		}

	}

	expectedBigINum, err := new(BigIntNum).
		NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"new(BigIntNum).NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps)\n"+
			"expectedNumStr='%v'; expectedNumSeps='%v'\nError='%v'\n",
			expectedNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyNumStrSeries()\n"+
			"multiplierStr='%v'; expectedNumSeps='%v'\nError='%v'\n",
			multiplierStr, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumEqualResult, err :=\n"+
			"expectedBigINum.Equal(result)\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedBigINumEqualResult {
		t.Errorf("Error: Expected Decimal='%s'.\nInstead, Decimal= '%s'\n. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
		return
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'.\nInstead, Decimal= '%s'\n",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
		return
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	expectedNumStr, err = iaResult.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.\n"+
			"Instead, NumSeps='%v'\n",
			expectedNumSeps.String(), actualNumSeps.String())
	}

	return
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

	iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by\n"+
			"new(IntAry).NewNumStrWithNumSeps("+
			"multiplierStr, expectedNumSeps)\n"+
			"multiplierStr='%v' expectedNumSeps='%v'\nError='%v'\n",
			multiplierStr, expectedNumSeps.String(), err.Error())
		return
	}

	for i := 0; i < lenArray; i++ {

		ia, err := new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
				"multiplicandStrs[i], expectedNumSeps) "+
				"multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
				i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ",
				i, multiplicandStrs[i], err.Error())
			return
		}
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(
		expectedNumStr, expectedNumSeps, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"new(BigIntNum).NewNumStrWithNumSeps("+
			"expectedNumStr, expectedNumSeps)\n"+
			"expectedNumStr='%v'; expectedNumSeps='%v'\nError='%v'\n",
			expectedNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
		expectedNumSeps,
		expectedNumSeps,
		multiplierStr,
		multiplicandStrs[0],
		multiplicandStrs[1],
		multiplicandStrs[2],
		multiplicandStrs[3],
		multiplicandStrs[4],
		multiplicandStrs[5])

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(...)\n"+
			"expectedNumStr='%v'; expectedNumSeps='%v'\nError='%v'\n",
			expectedNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedBigINumEqualResult, err := \n"+
			"expectedBigINum.Equal(result)\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedBigINumEqualResult {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
		return
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
		return
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
		return
	}

	actualNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"actualNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	expectedNumStr, err = iaResult.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedNumStr, err = iaResult.GetNumStr()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("Error returned by:\n"+
			"expectedNumStr, err = iaResult.GetNumStr()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

	return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_01(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875.94572"

	expectedSignValue := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedNumStr) "+
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDto"+
			"(multiplierNumStrDto, multiplicandNumStrDto) "+
			"multiplierNumStrDto='%v' multiplicandNumStrDto='%v' Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), multiplicandNumStrDto.GetNumStr(), err.Error())
	}

	if expectedNumStrDto.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedNumStrDto.GetBigInt() "+
			"Error='%v'. ",
			err.Error())
	}

	if expectedNumStrDtoBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDto_02(t *testing.T) {
	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedNumStr := "30987680500513189125.14259702468435"

	expectedSignValue := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedNumStr) "+
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDto"+
			"(multiplierNumStrDto, multiplicandNumStrDto) "+
			"multiplierNumStrDto='%v' multiplicandNumStrDto='%v' Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), multiplicandNumStrDto.GetNumStr(), err.Error())
	}

	if expectedNumStrDto.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedNumStrDto.GetBigInt() "+
			"Error='%v'. ",
			err.Error())
	}

	if expectedNumStrDtoBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDto_03(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedNumStr := "-30987680500513189125.14259702468435"

	expectedSignValue := -1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedNumStr) "+
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDto"+
			"(multiplierNumStrDto, multiplicandNumStrDto) "+
			"multiplierNumStrDto='%v' multiplicandNumStrDto='%v' Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), multiplicandNumStrDto.GetNumStr(), err.Error())
	}

	if expectedNumStrDto.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedNumStrDto.GetBigInt() "+
			"Error='%v'. ",
			err.Error())
	}

	if expectedNumStrDtoBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDto_04(t *testing.T) {
	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedNumStr := "22197234145.3632"

	expectedSignValue := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedNumStr) "+
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDto"+
			"(multiplierNumStrDto, multiplicandNumStrDto) "+
			"multiplierNumStrDto='%v' multiplicandNumStrDto='%v' Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), multiplicandNumStrDto.GetNumStr(), err.Error())
	}

	if expectedNumStrDto.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedNumStrDto.GetBigInt() "+
			"Error='%v'. ",
			err.Error())
	}

	if expectedNumStrDtoBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDto_05(t *testing.T) {
	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0
	expectedNumStr := "0"

	expectedSignValue := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := new(IntAry).New()

	err = iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() "+
			"Error='%v'. ", err.Error())
	}

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedNumStr) "+
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDto"+
			"(multiplierNumStrDto, multiplicandNumStrDto) "+
			"multiplierNumStrDto='%v' multiplicandNumStrDto='%v' Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), multiplicandNumStrDto.GetNumStr(), err.Error())
	}

	if expectedNumStrDto.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedNumStrDto.GetBigInt() "+
			"Error='%v'. ",
			err.Error())
	}

	if expectedNumStrDtoBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedNumStrDto.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by result.GetBigInt() "+
			"Error='%v'. ", err.Error())
	}

	iaBigInt, err := iaResult.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by iaResult.GetBigInt() "+
			"Error='%v'. ", err.Error())
	}

	if actualBigInt.Cmp(iaBigInt) != 0 {
		t.Errorf("Error: Expected actualBigInt='%v' "+
			"Instead, actualBigInt='%v'",
			iaResult.GetNumStr(), actualBigInt)
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDto_06(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875,94572
	expectedNumStr := "2875,94572"

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr) "+
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDto"+
			"(multiplierNumStrDto, multiplicandNumStrDto) "+
			"multiplierNumStrDto='%v' multiplicandNumStrDto='%v' Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), multiplicandNumStrDto.GetNumStr(), err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_01(t *testing.T) {

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
	expectedBigINumStr := "128"

	expectedBigINumSign := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := new(IntAry).NewNumStrDto(nDtoArray[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStrDto(nDtoArray[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ",
				i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDtoArray("+
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
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

}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_02(t *testing.T) {

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
	expectedBigINumStr := "11995826664.26376575446779648"

	expectedBigINumSign := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDtoArray("+
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
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

}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_03(t *testing.T) {

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
	expectedBigINumStr := "2212352.1767579232"

	expectedBigINumSign := 1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDtoArray("+
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
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

}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_04(t *testing.T) {

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
	expectedBigINumStr := "-20408.5138429311978576052224"

	expectedBigINumSign := -1

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := nDtoArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by nDtoArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDtoArray("+
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
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

}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_05(t *testing.T) {

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

	// product = 11995826664,26376575446779648
	expectedNumStr := "11995826664,26376575446779648"

	multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	nDtoArray := make([]NumStrDto, lenArray)

	for i := 0; i < lenArray; i++ {

		nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}
	}

	result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDtoArray("+
			"multiplierNumStrDto, nDtoArray) multiplierNumStrDto='%v'  Error='%v'. ",
			multiplierNumStrDto.GetNumStr(), err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v' "+
			"Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}
