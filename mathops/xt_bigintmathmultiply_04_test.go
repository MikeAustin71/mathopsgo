package mathops

import "testing"

func TestBigIntMathMultiply_MultiplyDecimalArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyDecimalArray_01"

	var err error

	// multiplier = 2
	originalMultiplierStr := "2"
	// multiplicandStrs
	originalMultiplicandStrs := []string{
		"2",
		"2",
		"2",
		"2",
		"2",
		"2",
	}

	// product = 128
	originalExpectedNumStr := "128"

	originalExpectedSignValue := 1

	multiplierDecimal, err := new(Decimal).
		NewNumStr(originalMultiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierDecimal, err := new(Decimal).\n"+
			"  NewNumStr(originalMultiplierStr)\n"+
			"originalMultiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalMultiplierStr, err.Error())
		return
	}

	multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierDecimalNumStr, err := \n"+
			"    multiplierDecimal.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(originalMultiplicandStrs)

	if lenArray == 0 {
		t.Errorf("%v\n"+
			"Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
			ePrefix)
		return
	}

	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).
		NewNumStr(originalMultiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)\n"+
			"originalMultiplierStr='%v'\nError='%v'\n\n",
			ePrefix, originalMultiplierStr, err.Error())
		return
	}

	var ia IntAry
	var iaNumStr string

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"decimalArray[%d], err = new(Decimal).\n"+
				"    NewNumStr(originalMultiplicandStrs[%d]))\n"+
				"multiplicandStrs[%d]='%v'\nError='%v'\n\n",
				ePrefix,
				i, i, i,
				originalMultiplicandStrs[i],
				err.Error())

			return
		}

		ia, err = decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = decimalArray[%d].GetIntAry()\n"+
				"Error='%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"ia='%v'\nError='%v'\n\n",
				ePrefix, iaNumStr, err.Error())
			return
		}

	} // End of for loop

	expectedBigINum, err := new(BigIntNum).
		NewNumStr(originalExpectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalExpectedNumStr)\n"+
			"originalExpectedNumStr='%v'\nError='%v'\n\n",
			ePrefix, originalExpectedNumStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSign, err := expectedBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyDecimalArray(\n"+
			"    multiplierDecimal, decimalArray)\n"+
			"multiplierDecimalNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierDecimalNumStr, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualResultBigINum, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualResultBigINum, err := \n"+
			"  expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumEqualResultBigINum {
		t.Errorf("%v\n"+
			"Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: *big.Int values do NOT match!\n"+
			"       Expected BigIntNum= '%s'.\n"+
			"         Result BigIntNum= '%s'.\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return

	}

	expectedBigINumCmpBigIntResultBigInt, err := expectedBigINum.CmpBigInt(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumCmpBigIntResultBigInt != 0 {
		t.Errorf("%v\n"+
			"Error: Comparision of expected vs result *big.Int Unequal!\n"+
			"Expected BigIntNum *big.Int='%s'.\n"+
			"   Actual 'result' *big.Int= '%s'. ",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	if originalExpectedNumStr != expectedBigINumStr ||
		originalExpectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: BigIntNum Number Strings are unequal!\n"+
			"         Expected Number String = '%s'.\n"+
			"           Result Number String = '%s'.\n"+
			"Original Expected Number String = '%s'",
			ePrefix,
			expectedBigINumStr,
			resultNumStr,
			originalExpectedNumStr)
		return

	}

	if originalExpectedSignValue != resultSignValue ||
		originalExpectedSignValue != expectedBigINumSign {
		t.Errorf("%v\n"+
			"Error: Expected Number Signs do NOT match!\n"+
			"originalExpectedBigINumSign sign='%v'.\n"+
			"            resultSignValue sign='%v'\n"+
			"             expectedBigINumSign='%v'\n",
			ePrefix,
			originalExpectedSignValue,
			resultSignValue,
			expectedBigINumSign)
		return
	}

	err = iaResult.OptimizeIntArrayLen(true)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.OptimizeIntArrayLen(true)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: iaResult and result number strings do NOT match!\n"+
			"iaResultNumStr='%v'\n"+
			"  resultNumStr='%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyDecimalArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyDecimalArray_02"

	var err error

	// multiplier = 37.9876
	originalMultiplierStr := "37.9876"

	// multiplicandStrs
	originalMultiplicandStrs := []string{
		"-27.9",
		"48.123456",
		"59.48721",
		"-3",
		"19.1",
		"69",
	}

	// product = 11995826664.26376575446779648
	originalExpectedNumStr := "11995826664.26376575446779648"

	originalExpectedSignValue := 1

	multiplierDecimal, err := new(Decimal).
		NewNumStr(originalMultiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierDecimal, err := new(Decimal).\n"+
			"  NewNumStr(originalMultiplierStr)\n"+
			"originalMultiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalMultiplierStr, err.Error())
		return
	}

	multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierDecimalNumStr, err := \n"+
			"    multiplierDecimal.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(originalMultiplicandStrs)

	if lenArray == 0 {
		t.Errorf("%v\n"+
			"Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
			ePrefix)
		return
	}

	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).
		NewNumStr(originalMultiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).\n"+
			"    NewNumStr(originalMultiplierStr)\n"+
			"originalMultiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalMultiplierStr, err.Error())
		return
	}

	var ia IntAry
	var iaNumStr string

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).
			NewNumStr(originalMultiplicandStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"decimalArray[%d], err = new(Decimal).\n"+
				"    NewNumStr(originalMultiplicandStrs[%d]))\n"+
				"originalMultiplicandStrs[%d]='%v'\n"+
				"Error='%v'\n\n",
				ePrefix,
				i, i, i,
				originalMultiplicandStrs[i],
				err.Error())

			return
		}

		ia, err = decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = decimalArray[%d].GetIntAry()\n"+
				"Error='%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"ia='%v'\nError='%v'\n\n",
				ePrefix, iaNumStr, err.Error())
			return
		}

	} // End of for loop

	expectedBigINum, err := new(BigIntNum).
		NewNumStr(originalExpectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"    NewNumStr(originalExpectedNumStr)\n"+
			"originalExpectedNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalExpectedNumStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSign, err := expectedBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).
		MultiplyDecimalArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyDecimalArray(\n"+
			"    multiplierDecimal, decimalArray)\n"+
			"multiplierDecimalNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierDecimalNumStr, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualResultBigINum, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualResultBigINum, err := \n"+
			"  expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumEqualResultBigINum {
		t.Errorf("%v\n"+
			"Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: *big.Int values do NOT match!\n"+
			"       Expected BigIntNum= '%s'.\n"+
			"         Result BigIntNum= '%s'.\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return

	}

	expectedBigINumCmpBigIntResultBigInt, err := expectedBigINum.CmpBigInt(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumCmpBigIntResultBigInt != 0 {
		t.Errorf("%v\n"+
			"Error: Comparision of expected vs result *big.Int Unequal!\n"+
			"Expected BigIntNum *big.Int='%s'.\n"+
			"   Actual 'result' *big.Int= '%s'. ",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	if originalExpectedNumStr != expectedBigINumStr ||
		originalExpectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: BigIntNum Number Strings are unequal!\n"+
			"         Expected Number String = '%s'.\n"+
			"           Result Number String = '%s'.\n"+
			"Original Expected Number String = '%s'",
			ePrefix,
			expectedBigINumStr,
			resultNumStr,
			originalExpectedNumStr)
		return

	}

	if originalExpectedSignValue != resultSignValue ||
		originalExpectedSignValue != expectedBigINumSign {
		t.Errorf("%v\n"+
			"Error: Expected Number Signs do NOT match!\n"+
			"originalExpectedBigINumSign sign='%v'.\n"+
			"            resultSignValue sign='%v'\n"+
			"             expectedBigINumSign='%v'\n",
			ePrefix,
			originalExpectedSignValue,
			resultSignValue,
			expectedBigINumSign)
		return
	}

	err = iaResult.OptimizeIntArrayLen(true)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.OptimizeIntArrayLen(true)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: iaResult and result number strings do NOT match!\n"+
			"iaResultNumStr='%v'\n"+
			"  resultNumStr='%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyDecimalArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathMultiply_MultiplyDecimalArray_03"

	var err error

	// multiplier = 10.1
	originalMultiplierStr := "10.1"

	// multiplicandStrs
	originalMultiplicandStrs := []string{
		"2",
		"5.8",
		"68.7",
		"3.1234567",
		"8.0",
		"11",
	}

	// product = 2212352.1767579232
	originalExpectedNumStr := "2212352.1767579232"

	originalExpectedSignValue := 1

	multiplierDecimal, err := new(Decimal).
		NewNumStr(originalMultiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierDecimal, err := new(Decimal).\n"+
			"  NewNumStr(originalMultiplierStr)\n"+
			"originalMultiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalMultiplierStr, err.Error())
		return
	}

	multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"multiplierDecimalNumStr, err := \n"+
			"    multiplierDecimal.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(originalMultiplicandStrs)

	if lenArray == 0 {
		t.Errorf("%v\n"+
			"Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
			ePrefix)
		return
	}

	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).
		NewNumStr(originalMultiplierStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResult, err := new(IntAry).\n"+
			"    NewNumStr(originalMultiplierStr)\n"+
			"originalMultiplierStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalMultiplierStr, err.Error())
		return
	}

	var ia IntAry
	var iaNumStr string

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).
			NewNumStr(originalMultiplicandStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"decimalArray[%d], err = new(Decimal).\n"+
				"    NewNumStr(originalMultiplicandStrs[%d]))\n"+
				"originalMultiplicandStrs[%d]='%v'\n"+
				"Error='%v'\n\n",
				ePrefix,
				i, i, i,
				originalMultiplicandStrs[i],
				err.Error())

			return
		}

		ia, err = decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = decimalArray[%d].GetIntAry()\n"+
				"Error='%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		iaNumStr, err = ia.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaNumStr, err = ia.GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
				"ia='%v'\nError='%v'\n\n",
				ePrefix, iaNumStr, err.Error())
			return
		}

	} // End of for loop

	expectedBigINum, err := new(BigIntNum).
		NewNumStr(originalExpectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"    NewNumStr(originalExpectedNumStr)\n"+
			"originalExpectedNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalExpectedNumStr, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSign, err := expectedBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathMultiply).
		MultiplyDecimalArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathMultiply).MultiplyDecimalArray(\n"+
			"    multiplierDecimal, decimalArray)\n"+
			"multiplierDecimalNumStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, multiplierDecimalNumStr, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumEqualResultBigINum, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumEqualResultBigINum, err := \n"+
			"  expectedBigINum.Equal(result)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBigINumEqualResultBigINum {
		t.Errorf("%v\n"+
			"Error: Expected BigIntNum='%s'.\n"+
			"Instead, BigIntNum= '%s'.\n\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: *big.Int values do NOT match!\n"+
			"       Expected BigIntNum= '%s'.\n"+
			"         Result BigIntNum= '%s'.\n",
			ePrefix,
			expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return

	}

	expectedBigINumCmpBigIntResultBigInt, err := expectedBigINum.CmpBigInt(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumCmpBigIntResultBigInt != 0 {
		t.Errorf("%v\n"+
			"Error: Comparision of expected vs result *big.Int Unequal!\n"+
			"Expected BigIntNum *big.Int='%s'.\n"+
			"   Actual 'result' *big.Int= '%s'. ",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
		return
	}

	if originalExpectedNumStr != expectedBigINumStr ||
		originalExpectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: BigIntNum Number Strings are unequal!\n"+
			"         Expected Number String = '%s'.\n"+
			"           Result Number String = '%s'.\n"+
			"Original Expected Number String = '%s'",
			ePrefix,
			expectedBigINumStr,
			resultNumStr,
			originalExpectedNumStr)
		return

	}

	if originalExpectedSignValue != resultSignValue ||
		originalExpectedSignValue != expectedBigINumSign {
		t.Errorf("%v\n"+
			"Error: Expected Number Signs do NOT match!\n"+
			"originalExpectedBigINumSign sign='%v'.\n"+
			"            resultSignValue sign='%v'\n"+
			"             expectedBigINumSign='%v'\n",
			ePrefix,
			originalExpectedSignValue,
			resultSignValue,
			expectedBigINumSign)
		return
	}

	err = iaResult.OptimizeIntArrayLen(true)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaResult.OptimizeIntArrayLen(true)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaResultNumStr, err := iaResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaResultNumStr, err := iaResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if iaResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: iaResult and result number strings do NOT match!\n"+
			"iaResultNumStr='%v'\n"+
			"  resultNumStr='%v'\n\n",
			ePrefix, iaResultNumStr, resultNumStr)
	}

	return
}

func TestBigIntMathMultiply_MultiplyDecimalArray_04(t *testing.T) {

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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
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
		t.Errorf("Error returned by new(Decimal).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalArray("+
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
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalArray_05(t *testing.T) {

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
	expectedNumStr := "-20408,5138429311978576052224"

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
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
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalArray("+
			"multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
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

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_01(t *testing.T) {

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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalOutputToArray"+
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalOutputToArray"+
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalOutputToArray"+
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalOutputToArray"+
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalOutputToArray"+
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

	}

}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_06(t *testing.T) {

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
		"1499085,809",
		"0",
	}

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
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
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalOutputToArray"+
			"(multiplierDecimal, decimalArray) multiplierDecimal='%v'  Error='%v'. ",
			multiplierDecimal.GetNumStr(), err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. Index='%v'",
				expectedNumSeps.String(), actualNumSeps.String(), j)
		}
	}
}

func TestBigIntMathMultiply_MultiplyDecimalSeries_01(t *testing.T) {

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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
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
		t.Errorf("Error returned by new(Decimal).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
		multiplierDecimal,
		decimalArray[0],
		decimalArray[1],
		decimalArray[2],
		decimalArray[3],
		decimalArray[4],
		decimalArray[5])

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

}

func TestBigIntMathMultiply_MultiplyDecimalSeries_02(t *testing.T) {

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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
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
		t.Errorf("Error returned by new(Decimal).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
		multiplierDecimal,
		decimalArray[0],
		decimalArray[1],
		decimalArray[2],
		decimalArray[3],
		decimalArray[4],
		decimalArray[5])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalSeries(multiplierDecimal,"+
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
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalSeries_03(t *testing.T) {

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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
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
		t.Errorf("Error returned by new(Decimal).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
		multiplierDecimal,
		decimalArray[0],
		decimalArray[1],
		decimalArray[2],
		decimalArray[3],
		decimalArray[4],
		decimalArray[5])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalSeries(multiplierDecimal,"+
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
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalSeries_04(t *testing.T) {

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

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
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
		t.Errorf("Error returned by new(Decimal).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
		multiplierDecimal,
		decimalArray[0],
		decimalArray[1],
		decimalArray[2],
		decimalArray[3],
		decimalArray[4],
		decimalArray[5])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalSeries(multiplierDecimal,"+
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
		t.Errorf("Error: Expected actualNumStr='%v' "+
			"Instead, actualNumStr='%v'",
			iaResult.GetNumStr(), actualNumStr)
	}

}

func TestBigIntMathMultiply_MultiplyDecimalSeries_05(t *testing.T) {

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

	// product = 2212352,1767579232
	expectedNumStr := "2212352,1767579232"

	multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
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
		t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(multiplicandStrs)
	decimalArray := make([]Decimal, lenArray)

	iaResult, err := new(IntAry).NewNumStr(multiplierStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
			"multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
	}

	for i := 0; i < lenArray; i++ {

		decimalArray[i], err = new(Decimal).NewNumStr(multiplicandStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStrs[i]) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		ia, err := decimalArray[i].GetIntAry()

		if err != nil {
			t.Errorf("Error returned by decimalArray[i].GetIntAryElements() "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

		err = iaResult.MultiplyThisBy(&ia, -1, -1)

		if err != nil {
			t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
				"i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
		multiplierDecimal,
		decimalArray[0],
		decimalArray[1],
		decimalArray[2],
		decimalArray[3],
		decimalArray[4],
		decimalArray[5])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimalSeries(multiplierDecimal,"+
			" ...) multiplierDecimal='%v'  Error='%v'. ", multiplierDecimal.GetNumStr(), err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStrp='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}
