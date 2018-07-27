package mathops

import "testing"

func TestBigIntMathMultiply_MultiplyDecimal_01(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875.94572"

	expectedSignValue := 1

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandDecimal, err := Decimal{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err =	iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedDecimal, err := Decimal{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedNumStr) " +
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimal" +
			"(multiplierDecimal, multiplicandDecimal) " +
			"multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
			multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
	}

	if expectedDecimal.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedDecimal.GetBigInt() " +
			"Error='%v'. ",
			err.Error())
	}

	if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimal_02(t *testing.T) {
	// multiplier = 57638422123.327890123
	multiplierStr := "57638422123.327890123"

	// multiplicand = 537621943.12345
	multiplicandStr := "537621943.12345"

	// product = 30987680500513189125.14259702468435
	expectedNumStr := "30987680500513189125.14259702468435"

	expectedSignValue := 1

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandDecimal, err := Decimal{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err =	iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedDecimal, err := Decimal{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedNumStr) " +
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimal" +
			"(multiplierDecimal, multiplicandDecimal) " +
			"multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
			multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
	}

	if expectedDecimal.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedDecimal.GetBigInt() " +
			"Error='%v'. ",
			err.Error())
	}

	if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimal_03(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "57638422123.327890123"

	// multiplicand = -537621943.12345
	multiplicandStr := "-537621943.12345"

	// product = -30987680500513189125.14259702468435
	expectedNumStr := "-30987680500513189125.14259702468435"

	expectedSignValue := -1

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandDecimal, err := Decimal{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err =	iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedDecimal, err := Decimal{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedNumStr) " +
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimal" +
			"(multiplierDecimal, multiplicandDecimal) " +
			"multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
			multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
	}

	if expectedDecimal.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedDecimal.GetBigInt() " +
			"Error='%v'. ",
			err.Error())
	}

	if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimal_04(t *testing.T) {
	// multiplier = 89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = -247632
	multiplicandStr := "-247632"

	// product = 22197234145.3632
	expectedNumStr := "22197234145.3632"

	expectedSignValue := 1

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandDecimal, err := Decimal{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err =	iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedDecimal, err := Decimal{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedNumStr) " +
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimal" +
			"(multiplierDecimal, multiplicandDecimal) " +
			"multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
			multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
	}

	if expectedDecimal.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedDecimal.GetBigInt() " +
			"Error='%v'. ",
			err.Error())
	}

	if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimal_05(t *testing.T) {
	// multiplier = -89637.9876
	multiplierStr := "-89637.9876"

	// multiplicand = 0.00
	multiplicandStr := "0.00"

	// product = 0.00
	expectedNumStr := "0"

	expectedSignValue := 1

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	multiplicandDecimal, err := Decimal{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaMultiplier, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	iaMultiplicand, err := IntAry{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	iaResult := IntAry{}.New()

	err =	iaMultiplier.Multiply(
		&iaMultiplier,
		&iaMultiplicand,
		&iaResult,
		-1,
		-1)

	if err != nil {
		t.Errorf("Error returned by iaMultiplier.Multiply() " +
			"Error='%v'. ", err.Error())
	}

	expectedDecimal, err := Decimal{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedNumStr) " +
			"expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimal" +
			"(multiplierDecimal, multiplicandDecimal) " +
			"multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
			multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
	}

	if expectedDecimal.GetNumStr() != result.GetNumStr() {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by expectedDecimal.GetBigInt() " +
			"Error='%v'. ",
			err.Error())
	}

	if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedDecimal.GetNumStr(), result.GetNumStr())
	}

	if expectedSignValue != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedSignValue, result.sign)
	}

	actualBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by result.GetBigInt() " +
			"Error='%v'. ", err.Error())
	}

	iaBigInt, err := iaResult.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by iaResult.GetBigInt() " +
			"Error='%v'. ", err.Error())
	}

	if actualBigInt.Cmp(iaBigInt) != 0 {
		t.Errorf("Error: Expected actualBigInt='%v' " +
			"Instead, actualBigInt='%v'",
			iaResult.GetNumStr(), actualBigInt)
	}

}

func TestBigIntMathMultiply_MultiplyDecimal_06(t *testing.T) {
	// multiplier = 123.32
	multiplierStr := "123.32"

	// multiplicand = 23.321
	multiplicandStr := "23.321"

	// product = 2875.94572
	expectedNumStr := "2875,94572"

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = multiplierDecimal.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	multiplicandDecimal, err := Decimal{}.NewNumStr(multiplicandStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStr) " +
			"multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimal" +
			"(multiplierDecimal, multiplicandDecimal) " +
			"multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
			multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathMultiply_MultiplyDecimalArray_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalArray(" +
			"multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
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
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalArray_02(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalArray(" +
			"multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
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
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalArray_03(t *testing.T) {

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalArray(" +
			"multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	iaResult.OptimizeIntArrayLen(true)

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalArray_04(t *testing.T) {

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalArray(" +
			"multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"2",
		"4",
		"6",
		"8",
		"10",
		"12",
	}


	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalOutputToArray" +
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_02(t *testing.T) {

	var err error

	// multiplier = 8
	multiplierStr := "8"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"800.8",
		"-208",
		"31.392",
		"64",
		"42376.984",
		"-39.168",
	}


	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalOutputToArray" +
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_03(t *testing.T) {

	var err error

	// multiplier = -31.2
	multiplierStr := "-31.2"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"100.1",
		"-26",
		"3.924",
		"8",
		"5297.123",
		"-4.896",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"-3123.12",
		"811.2",
		"-122.4288",
		"-249.6",
		"-165270.2376",
		"152.7552",
	}


	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalOutputToArray" +
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_04(t *testing.T) {

	var err error

	// multiplier = 283
	multiplierStr := "283"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"0",
		"-26",
		"0",
		"8",
		"5297.123",
		"0",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"0",
		"-7358",
		"0",
		"2264",
		"1499085.809",
		"0",
	}


	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalOutputToArray" +
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_05(t *testing.T) {

	var err error

	// multiplier = 0
	multiplierStr := "0"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
		"5",
		"-26",
		"9",
		"8",
		"5297.123",
		"37",
	}

	// Expected Results Array
	expectedNumStrs := [] string {
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
	}


	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalOutputToArray" +
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j:=0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr()  {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyDecimalSeries_01(t *testing.T) {

	var err error

	// multiplier = 2
	multiplierStr := "2"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalSeries(
		multiplierDecimal,
		decimalArray[0],
		decimalArray[1],
		decimalArray[2],
		decimalArray[3],
		decimalArray[4],
		decimalArray[5],)


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
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}


func TestBigIntMathMultiply_MultiplyDecimalSeries_02(t *testing.T) {

	var err error

	// multiplier = 37.9876
	multiplierStr := "37.9876"

	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalSeries(
		multiplierDecimal,
		decimalArray[0],
		decimalArray[1],
		decimalArray[2],
		decimalArray[3],
		decimalArray[4],
		decimalArray[5],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalSeries(multiplierDecimal," +
			" ...) multiplierDecimal='%v'  Error='%v'. ", multiplierDecimal.GetNumStr(), err.Error())
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
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalSeries_03(t *testing.T) {

	var err error

	// multiplier = 10.1
	multiplierStr := "10.1"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}


	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalSeries(
		multiplierDecimal,
		decimalArray[0],
		decimalArray[1],
		decimalArray[2],
		decimalArray[3],
		decimalArray[4],
		decimalArray[5],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalSeries(multiplierDecimal," +
			" ...) multiplierDecimal='%v'  Error='%v'. ", multiplierDecimal.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	iaResult.OptimizeIntArrayLen(true)

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalSeries_04(t *testing.T) {

	var err error

	// multiplier = -5.123456
	multiplierStr := "-5.123456"
	// multiplicandStrs
	multiplicandStrs :=  [] string {
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

	multiplierDecimal, err := Decimal{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := IntAry{}.NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(multiplierStr) " +
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i:=0; i < lenArray; i++ {

		decimalArray[i], err = 	Decimal{}.NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(multiplicandStrs[i]) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAry() " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) " +
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathMultiply{}.MultiplyDecimalSeries(
		multiplierDecimal,
		decimalArray[0],
		decimalArray[1],
		decimalArray[2],
		decimalArray[3],
		decimalArray[4],
		decimalArray[5],)

	if err != nil {
		t.Errorf("Error returned by BigIntMathMultiply{}.MultiplyDecimalSeries(multiplierDecimal," +
			" ...) multiplierDecimal='%v'  Error='%v'. ", multiplierDecimal.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.CmpBigInt(result) != 0 {
		t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

	actualNumStr := result.GetNumStr()

	if iaResult.GetNumStr() != actualNumStr {
		t.Errorf("Error: Expected actualNumStr='%v' " +
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

