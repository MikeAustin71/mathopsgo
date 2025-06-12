package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathAdd_AddINumMgr_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgr_01"
	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedFinalResult := "124443.912456"
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	ia1, err := new(IntAry).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	ia1NumStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumStr, err := ia1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dec2, err := new(Decimal).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec2, err := new(Decimal).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	dec2NumStr, err := dec2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec2NumStr, err := dec2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddINumMgr(&ia1, &dec2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddINumMgr(&ia1, &dec2)\n"+
			"ia1= '%v'\n"+
			"dec2= '%v'\n"+
			"Error='%v'\n\n",
			ia1NumStr,
			dec2NumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
	}

	return
}

func TestBigIntMathAdd_AddINumMgr_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgr_02"
	n1Str := "123456.789"
	n2Str := "-987.123456"

	expectedFinalResult := "122469.665544"
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	ia1, err := new(IntAry).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	ia1NumStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumStr, err := ia1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dec2, err := new(Decimal).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec2, err := new(Decimal).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	dec2NumStr, err := dec2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec2NumStr, err := dec2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddINumMgr(&ia1, &dec2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddINumMgr(&ia1, &dec2)\n"+
			"ia1= '%v'\n"+
			"dec2= '%v'\n"+
			"Error='%v'\n\n",
			ia1NumStr,
			dec2NumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
	}

	return
}

func TestBigIntMathAdd_AddINumMgr_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgr_03"
	n1Str := "-123456.789"
	n2Str := "987.123456"

	// Result := -122469.665544
	expectedFinalResult := "-122469.665544"
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	nDto1, err := new(NumStrDto).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto1, err := new(NumStrDto).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	nDto1NumStr, err := nDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto1NumStr, err := nDto1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dec2, err := new(Decimal).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec2, err := new(Decimal).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	dec2NumStr, err := dec2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec2NumStr, err := dec2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddINumMgr(
		nDto1.GetThisPointer(),
		dec2.GetThisPointer())

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddINumMgr(\n"+
			"nDto1.GetThisPointer(), dec2.GetThisPointer())"+
			"nDto1= '%v'\n"+
			"dec2= '%v'\n"+
			"Error='%v'\n\n",
			nDto1NumStr,
			dec2NumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
	}

	return
}

func TestBigIntMathAdd_AddINumMgr_04(t *testing.T) {

	ePrefix := "BigIntMathAdd_AddINumMgr_04"
	n1Str := "-123456.789"
	n2Str := "-987.123456"

	// Result := -124443.912456
	expectedFinalResult := "-124443.912456"
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	nDto1, err := new(NumStrDto).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto1, err := new(NumStrDto).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	nDto1NumStr, err := nDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto1NumStr, err := nDto1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	ia2NumStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumStr, err := ia2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddINumMgr(&nDto1, &ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddINumMgr(\n"+
			"&nDto1, &ia2)"+
			"nDto1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error='%v'\n\n",
			nDto1NumStr,
			ia2NumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)
	}

	return
}

func TestBigIntMathAdd_AddINumMgr_05(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgr_05"
	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedFinalResult := "124443.912456"
	expectedResultStr := "124443,912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
		return
	}

	ia1, err := new(IntAry).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n1Str, err.Error())
		return
	}

	ia1NumStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumStr, err := ia1.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = ia1.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dec2, err := new(Decimal).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec2, err := new(Decimal).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n", ePrefix, n2Str, err.Error())
		return
	}

	dec2NumStr, err := dec2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec2NumStr, err := dec2.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathAdd).AddINumMgr(&ia1, &dec2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddINumMgr(&ia1, &dec2)\n"+
			"ia1= '%v'\n"+
			"dec2= '%v'\n"+
			"Error='%v'\n\n",
			ia1NumStr,
			dec2NumStr,
			ePrefix,
			err.Error())
		return
	}

	actualResultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualResultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedFinalResult != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedFinalResult,
			actualResultNumStr)
		return
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.bigInt = '%v'\n"+
			"Instead, result.bigInt = '%v'\n\n",
			ePrefix,
			biExpectedResult.Text(10),
			result.bigInt.Text(10))
		return
	}

	if expectedPrecision != result.precision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.precision = '%v'\n"+
			"Instead, result.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			result.precision)
		return
	}

	if expectedSign != result.sign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected result.sign = '%v'\n"+
			"Instead, result.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			result.sign)

		return
	}

	actualNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualNumSeps = '%v'\n"+
			"Instead, actualNumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_AddINumMgrArray_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgrArray_01"

	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry := len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)\n"+
			"expectedTotalStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedTotalStr, err.Error())
		return
	}

	expectedResultNumStr, err := expectedBNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultNumStr, err := expectedBNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	inumMgrAry := make([]INumMgr, lenStrAry)

	var dec Decimal

	for i := 0; i < lenStrAry; i++ {

		if i < 2 {

			dec, err = new(Decimal).NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("%v\n"+
					"Error returned by:\n"+
					"dec, err = new(Decimal).NewNumStr(numStrAry[%d])\n"+
					"numStrAry[%d]= '%v'\n"+
					"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
				return
			}

			inumMgrAry[i] = &dec

		} else if i < 4 {

			nDto, err := new(NumStrDto).NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("%v\n"+
					"Error returned by:\n"+
					"nDto, err := new(NumStrDto).NewNumStr(numStrAry[%d])\n"+
					"numStrAry[%d]= '%v'\n"+
					"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
				return
			}

			inumMgrAry[i] = &nDto

		} else {
			// i must be >=4

			ia, err := new(IntAry).NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("%v\n"+
					"Error returned by:\n"+
					"ia, err := new(IntAry).NewNumStr(numStrAry[%d])\n"+
					"numStrAry[%d]= '%v'\n"+
					"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
				return
			}

			inumMgrAry[i] = &ia
		}

	}

	total, err := new(BigIntMathAdd).AddINumMgrArray(inumMgrAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddINumMgrArray(inumMgrAry)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	totalNumstr, err := total.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalNumstr, err := total.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != totalNumstr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected totalNumstr = '%v'\n"+
			"Instead, totalNumstr = '%v'\n\n",
			ePrefix,
			expectedResultNumStr,
			totalNumstr)
		return
	}

	expectedPrecision, err := expectedBNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedPrecision, err := expectedBNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	totalPrecision, err := total.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalPrecision, err := total.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedPrecision != totalPrecision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total.precision = '%v'\n"+
			"Instead, total.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			totalPrecision)
		return
	}

	expectedSign, err := expectedBNum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedSign, err := expectedBNum.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	totalSign, err := total.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalSign, err := total.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedSign != totalSign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total.sign = '%v'\n"+
			"Instead, total.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			totalSign)
	}

	return
}

func TestBigIntMathAdd_AddINumMgrArray_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgrArray_02"
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry := len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)\n"+
			"expectedTotalStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedTotalStr, err.Error())
		return
	}

	expectedResultNumStr, err := expectedBNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultNumStr, err := expectedBNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	inumMgrAry := make([]INumMgr, lenStrAry)
	var dec Decimal
	var nDto NumStrDto
	var ia IntAry

	for i := 0; i < lenStrAry; i++ {

		if i < 2 {

			dec, err = new(Decimal).NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("%v\n"+
					"Error returned by:\n"+
					"dec, err = new(Decimal).NewNumStr(numStrAry[%d])\n"+
					"numStrAry[%d]= '%v'\n"+
					"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
				return
			}

			inumMgrAry[i] = &dec

		} else if i < 4 {

			nDto, err = new(NumStrDto).NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("%v\n"+
					"Error returned by:\n"+
					"nDto, err = new(NumStrDto).NewNumStr(numStrAry[%d])\n"+
					"numStrAry[%d]= '%v'\n"+
					"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
				return
			}

			inumMgrAry[i] = &nDto

		} else {
			// 'i' must be >= 4

			ia, err = new(IntAry).NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("%v\n"+
					"Error returned by:\n"+
					"ia, err = new(IntAry).NewNumStr(numStrAry[%d])\n"+
					"numStrAry[%d]= '%v'\n"+
					"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
				return
			}

			inumMgrAry[i] = &ia
		}

	}

	total, err := new(BigIntMathAdd).AddINumMgrArray(inumMgrAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddINumMgrArray(inumMgrAry)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	totalNumStr, err := total.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalNumStr, err := total.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBNumIsEqualTotal, err := expectedBNum.Equal(total)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNumIsEqualTotal, err := expectedBNum.Equal(total)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedBNumIsEqualTotal {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedBNum.bigInt.Text(10),
			total.bigInt.Text(10))
		return
	}

	if expectedResultNumStr != totalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected totalNumStr = '%v'\n"+
			"Instead, totalNumStr = '%v'\n\n",
			ePrefix,
			expectedResultNumStr,
			totalNumStr)
		return
	}

	expectedPrecision, err := expectedBNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedPrecision, err := expectedBNum.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	totalPrecision, err := total.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalPrecision, err := total.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedPrecision != totalPrecision {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total.precision = '%v'\n"+
			"Instead, total.precision = '%v'\n\n",
			ePrefix,
			expectedPrecision,
			totalPrecision)
		return
	}

	expectedSign, err := expectedBNum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedSign, err := expectedBNum.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	totalSign, err := total.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalSign, err := total.GetSign()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedSign != totalSign {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total.sign = '%v'\n"+
			"Instead, total.sign = '%v'\n\n",
			ePrefix,
			expectedSign,
			totalSign)
	}

	return
}

func TestBigIntMathAdd_AddINumMgrArray_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgrArray_03"
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry := len(numStrAry)

	expectedTotalStr := "158,14788"

	inumMgrAry := make([]INumMgr, lenStrAry)
	var err error
	var dec Decimal
	var nDto NumStrDto
	var ia IntAry

	for i := 0; i < lenStrAry; i++ {

		if i < 2 {

			dec, err = new(Decimal).NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("%v\n"+
					"Error returned by:\n"+
					"dec, err = new(Decimal).NewNumStr(numStrAry[%d])\n"+
					"numStrAry[%d]= '%v'\n"+
					"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
				return
			}

			inumMgrAry[i] = &dec

		} else if i < 4 {

			nDto, err = new(NumStrDto).NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("%v\n"+
					"Error returned by:\n"+
					"nDto, err = new(NumStrDto).NewNumStr(numStrAry[%d])\n"+
					"numStrAry[%d]= '%v'\n"+
					"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
				return
			}

			inumMgrAry[i] = &nDto

		} else {
			// 'i' must be >= 4

			ia, err = new(IntAry).NewNumStr(numStrAry[i])

			if err != nil {
				t.Errorf("%v\n"+
					"Error returned by:\n"+
					"ia, err = new(IntAry).NewNumStr(numStrAry[%d])\n"+
					"numStrAry[%d]= '%v'\n"+
					"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
				return
			}

			inumMgrAry[i] = &ia
		}

	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = inumMgrAry[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = inumMgrAry[0].SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	total, err := new(BigIntMathAdd).AddINumMgrArray(inumMgrAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddINumMgrArray(inumMgrAry)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	totalNumStr, err := total.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"totalNumStr, err := total.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedTotalStr != totalNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected totalNumStr = '%v'\n"+
			"Instead, totalNumStr = '%v'\n\n",
			ePrefix,
			expectedTotalStr,
			totalNumStr)
		return
	}

	actualNumSeps, err := total.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumSeps, err := total.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total NumSeps = '%v'\n"+
			"Instead, total NumSeps = '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			actualNumSeps.String())
	}

	return
}

func TestBigIntMathAdd_AddINumMgrOutputToArray_01(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// numMgrStrs
	numMgrStrs := []string{
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"10",
		"15.123",
		"20",
		"258.692",
		"40",
		"60",
	}

	iNumMgrAddend, err := new(Decimal).NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(addendStr) "+
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(numMgrStrs)
	iNumMgrArray := make([]INumMgr, lenArray)

	for i := 0; i < lenArray; i++ {

		nDto, err := new(NumStrDto).NewNumStr(numMgrStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(NumStrDto).NewNumStr(numMgrStrs[i]) "+
				"i='%v'  numMgrStrs[i]='%v'  Error='%v'. ", i, numMgrStrs[i], err.Error())
		}

		iNumMgrArray[i] = &nDto
	}

	result, err := new(BigIntMathAdd).AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddINumMgrOutputToArray("+
			"&iNumMgrAddend, iNumMgrArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddINumMgrOutputToArray_02(t *testing.T) {

	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// numMgrStrs
	numMgrStrs := []string{
		"5",
		"10.123",
		"0",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"8.1",
		"13.223",
		"3.1",
		"256.792",
		"38.1",
		"58.1",
	}

	iNumMgrAddend, err := new(BigIntNum).NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(addendStr) "+
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(numMgrStrs)
	iNumMgrsArray := make([]INumMgr, lenArray)

	for i := 0; i < lenArray; i++ {

		ia, err := new(IntAry).NewNumStr(numMgrStrs[i])

		if err != nil {
			t.Errorf("Error returned by INumMgr{}.NewNumStr(numMgrStrs[i]) "+
				"i='%v'  numMgrStrs[i]='%v'  Error='%v'. ", i, numMgrStrs[i], err.Error())
		}

		iNumMgrsArray[i] = &ia

	}

	result, err := new(BigIntMathAdd).AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrsArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddINumMgrOutputToArray("+
			"iNumMgrAddend, iNumMgrsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddINumMgrOutputToArray_03(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// numMgrStrs
	numMgrStrs := []string{
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"10",
		"15,123",
		"20",
		"258,692",
		"40",
		"60",
	}

	iNumMgrAddend, err := new(Decimal).NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(addendStr) "+
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = iNumMgrAddend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iNumMgrAddend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(numMgrStrs)
	iNumMgrArray := make([]INumMgr, lenArray)

	for i := 0; i < lenArray; i++ {

		nDto, err := new(NumStrDto).NewNumStr(numMgrStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(NumStrDto).NewNumStr(numMgrStrs[i]) "+
				"i='%v'  numMgrStrs[i]='%v'  Error='%v'. ", i, numMgrStrs[i], err.Error())
		}

		iNumMgrArray[i] = &nDto
	}

	result, err := new(BigIntMathAdd).AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddINumMgrOutputToArray("+
			"&iNumMgrAddend, iNumMgrArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v' index='%v' ",
				expectedNumSeps.String(), actualNumSeps.String(), j)
		}
	}
}

func TestBigIntMathAdd_AddINumMgrSeries_01(t *testing.T) {
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedTotalStr := "158.14788"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	dec1, err := new(Decimal).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by  new(Decimal).NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'. ",
			n1Str, err.Error())
	}

	nDto2, err := new(NumStrDto).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by  new(NumStrDto).NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'. ",
			n2Str, err.Error())
	}

	ia3, err := new(IntAry).NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n3Str). "+
			"n3Str='%v' Error='%v'. ",
			n3Str, err.Error())
	}

	bigINum4, err := new(BigIntNum).NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(n4Str). "+
			"n4Str='%v' Error='%v'. ",
			n4Str, err.Error())
	}

	dec5, err := new(Decimal).NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by  new(Decimal).NewNumStr(n5Str). "+
			"n5Str='%v' Error='%v'. ",
			n5Str, err.Error())
	}

	total, err := new(BigIntMathAdd).AddINumMgrSeries(&dec1, &nDto2, &ia3, &bigINum4, &dec5)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddINumMgrSeries(&dec1, "+
			"&nDto2, &ia3, &bigINum4, &dec5). Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}
}

func TestBigIntMathAdd_AddINumMgrSeries_02(t *testing.T) {
	n1Str := "-978425.648941"
	n2Str := "33.12"
	n3Str := "-804.1"
	n4Str := "32567"
	n5Str := "-41.859"
	expectedTotalStr := "-946671.487941"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	dec1, err := new(Decimal).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by  new(Decimal).NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'. ",
			n1Str, err.Error())
	}

	ia2, err := new(IntAry).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'. ",
			n2Str, err.Error())
	}

	nDto3, err := new(NumStrDto).NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(n3Str). "+
			"n3Str='%v' Error='%v'. ",
			n3Str, err.Error())
	}

	dec4, err := new(Decimal).NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by  new(Decimal).NewNumStr(n4Str). "+
			"n4Str='%v' Error='%v'. ",
			n4Str, err.Error())
	}

	bINum5, err := new(BigIntNum).NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(n5Str). "+
			"n5Str='%v' Error='%v'. ",
			n5Str, err.Error())
	}

	total, err := new(BigIntMathAdd).AddINumMgrSeries(&dec1, &ia2, &nDto3, &dec4, &bINum5)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalSeries(&dec1, "+
			"&ia2, &nDto3, &dec4, &bINum5). Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}
}

func TestBigIntMathAdd_AddINumMgrSeries_03(t *testing.T) {
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedTotalStr := "158,14788"

	dec1, err := new(Decimal).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by  new(Decimal).NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'. ",
			n1Str, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = dec1.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by dec1.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	nDto2, err := new(NumStrDto).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by  new(NumStrDto).NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'. ",
			n2Str, err.Error())
	}

	ia3, err := new(IntAry).NewNumStr(n3Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n3Str). "+
			"n3Str='%v' Error='%v'. ",
			n3Str, err.Error())
	}

	bigINum4, err := new(BigIntNum).NewNumStr(n4Str)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(n4Str). "+
			"n4Str='%v' Error='%v'. ",
			n4Str, err.Error())
	}

	dec5, err := new(Decimal).NewNumStr(n5Str)

	if err != nil {
		t.Errorf("Error returned by  new(Decimal).NewNumStr(n5Str). "+
			"n5Str='%v' Error='%v'. ",
			n5Str, err.Error())
	}

	total, err := new(BigIntMathAdd).AddINumMgrSeries(&dec1, &nDto2, &ia3, &bigINum4, &dec5)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddINumMgrSeries(&dec1, "+
			"&nDto2, &ia3, &bigINum4, &dec5). Error='%v' ", err.Error())
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedTotalStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedTotalStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddIntAry_01(t *testing.T) {
	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedResultStr := "124443912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	ia1, err := new(IntAry).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := new(IntAry).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAry(ia1, ia2). "+
			"Error='%v' ", err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

}

func TestBigIntMathAdd_AddIntAry_02(t *testing.T) {

	n1Str := "123456.789"
	n2Str := "-987.123456"
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	ia1, err := new(IntAry).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := new(IntAry).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAry(ia1, ia2). "+
			"Error='%v' ", err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

}

func TestBigIntMathAdd_AddIntAry_03(t *testing.T) {

	n1Str := "-123456.789"
	n2Str := "987.123456"
	// Result := -122469.665544
	expectedResultStr := "-122469665544"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	ia1, err := new(IntAry).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := new(IntAry).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAry(ia1, ia2). "+
			"Error='%v' ", err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}
}

func TestBigIntMathAdd_AddIntAry_04(t *testing.T) {

	n1Str := "-123456.789"
	n2Str := "-987.123456"
	// Result := -124443.912456
	expectedResultStr := "-124443912456"
	expectedPrecision := uint(6)
	expectedSign := -1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

	if !oK {
		t.Error("Error returned by biExpectedResult, oK := " +
			"big.NewInt(0).SetString(expectedResultStr, 10)")
	}

	ia1, err := new(IntAry).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	ia2, err := new(IntAry).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAry(ia1, ia2). "+
			"Error='%v' ", err.Error())
	}

	if biExpectedResult.Cmp(result.bigInt) != 0 {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			biExpectedResult.Text(10), result.bigInt.Text(10))
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

}

func TestBigIntMathAdd_AddIntAry_05(t *testing.T) {
	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedResultStr := "124443,912456"
	expectedPrecision := uint(6)
	expectedSign := 1

	ia1, err := new(IntAry).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'. ", n1Str, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = ia1.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by ia1.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v'", err.Error())
	}

	ia2, err := new(IntAry).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'. ", n2Str, err.Error())
	}

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAry(ia1, ia2). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
			expectedResultStr, actualResultStr)
	}

	if expectedPrecision != result.precision {
		t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
			expectedPrecision, result.precision)
	}

	if expectedSign != result.sign {
		t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
			expectedSign, result.sign)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddIntAryArray_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry := len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	iaArray := make([]IntAry, lenStrAry)

	for i := 0; i < lenStrAry; i++ {

		ia, err := new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by new(IntAry).NewNumStr(numStrAry[i]) "+
					"i='%v' numStrAry[i]='%v' Error='%v' ", i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := new(BigIntMathAdd).AddIntAryArray(iaArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAryArray(iaArray). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddIntAryArray_02(t *testing.T) {
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry := len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	iaArray := make([]IntAry, lenStrAry)

	for i := 0; i < lenStrAry; i++ {

		ia, err := new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by new(IntAry).NewNumStr(numStrAry[i]) "+
					"i='%v' numStrAry[i]='%v' Error='%v' ", i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := new(BigIntMathAdd).AddIntAryArray(iaArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAryArray(iaArray). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddIntAryArray_03(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry := len(numStrAry)

	expectedTotalStr := "158,14788"

	iaArray := make([]IntAry, lenStrAry)

	for i := 0; i < lenStrAry; i++ {

		ia, err := new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by new(IntAry).NewNumStr(numStrAry[i]) "+
					"i='%v' numStrAry[i]='%v' Error='%v' ", i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := iaArray[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iaArray[0].SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v'", err.Error())
	}

	total, err := new(BigIntMathAdd).AddIntAryArray(iaArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAryArray(iaArray). "+
			"Error='%v' ", err.Error())
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedTotalStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedTotalStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathAdd_AddIntAryOutputToArray_01(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// iaNumStrs
	iaNumStrs := []string{
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"10",
		"15.123",
		"20",
		"258.692",
		"40",
		"60",
	}

	iaAddend, err := new(IntAry).NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(addendStr) "+
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(iaNumStrs)
	iaArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		iaArray[i], err = new(IntAry).NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(iaNumStrs[i]) "+
				"i='%v'  iaNumStrs[i]='%v'  Error='%v'. ", i, iaNumStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathAdd).AddIntAryOutputToArray(iaAddend, iaArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAryOutputToArray("+
			"iaAddend, iaArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddIntAryOutputToArray_02(t *testing.T) {

	var err error

	// addendStr = 3.1
	addendStr := "3.1"

	// iaNumStrs
	iaNumStrs := []string{
		"5",
		"10.123",
		"0",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"8.1",
		"13.223",
		"3.1",
		"256.792",
		"38.1",
		"58.1",
	}

	iaAddend, err := new(IntAry).NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(addendStr) "+
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	lenArray := len(iaNumStrs)
	decsArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		decsArray[i], err = new(IntAry).NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(iaNumStrs[i]) "+
				"i='%v'  iaNumStrs[i]='%v'  Error='%v'. ", i, iaNumStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathAdd).AddIntAryOutputToArray(iaAddend, decsArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAryOutputToArray("+
			"iaAddend, decsArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}
	}
}

func TestBigIntMathAdd_AddIntAryOutputToArray_03(t *testing.T) {

	var err error

	// addendStr = 5
	addendStr := "5"

	// iaNumStrs
	iaNumStrs := []string{
		"5",
		"10.123",
		"15",
		"253.692",
		"35",
		"55",
	}

	// Expected Results Array
	expectedNumStrs := []string{
		"10",
		"15,123",
		"20",
		"258,692",
		"40",
		"60",
	}

	iaAddend, err := new(IntAry).NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(addendStr) "+
			"addendStr='%v'  Error='%v'. ", addendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = iaAddend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iaAddend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenArray := len(iaNumStrs)
	iaArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		iaArray[i], err = new(IntAry).NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(iaNumStrs[i]) "+
				"i='%v'  iaNumStrs[i]='%v'  Error='%v'. ", i, iaNumStrs[i], err.Error())
		}

	}

	result, err := new(BigIntMathAdd).AddIntAryOutputToArray(iaAddend, iaArray)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntAryOutputToArray("+
			"iaAddend, iaArray) addendStr='%v'  Error='%v'. ",
			addendStr, err.Error())
	}

	for j := 0; j < lenArray; j++ {

		if expectedNumStrs[j] != result[j].GetNumStr() {
			t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
				j, expectedNumStrs[j], j, result[j].GetNumStr())
		}

		actualNumSeps := result[j].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'.",
				expectedNumSeps.String(), actualNumSeps.String())
		}
	}
}

func TestBigIntMathAdd_AddIntArySeries_01(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry := len(numStrAry)

	expectedTotalStr := "158.14788"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	iaArray := make([]IntAry, lenStrAry)

	for i := 0; i < lenStrAry; i++ {

		ia, err := new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by new(IntAry).NewNumStr(numStrAry[i]) "+
					"i='%v' numStrAry[i]='%v' Error='%v' ", i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := new(BigIntMathAdd).AddIntArySeries(
		iaArray[0],
		iaArray[1],
		iaArray[2],
		iaArray[3],
		iaArray[4])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntArySeries(...). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddIntArySeries_02(t *testing.T) {
	numStrAry := []string{
		"-978425.648941",
		"33.12",
		"-804.1",
		"32567",
		"-41.859",
	}

	lenStrAry := len(numStrAry)

	expectedTotalStr := "-946671.487941"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
			"expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

	}

	expectedResultNumStr := expectedBNum.GetNumStr()

	iaArray := make([]IntAry, lenStrAry)

	for i := 0; i < lenStrAry; i++ {

		ia, err := new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by new(IntAry).NewNumStr(numStrAry[i]) "+
					"i='%v' numStrAry[i]='%v' Error='%v' ", i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia

	}

	total, err := new(BigIntMathAdd).AddIntArySeries(
		iaArray[0],
		iaArray[1],
		iaArray[2],
		iaArray[3],
		iaArray[4])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntArySeries(...). "+
			"Error='%v' ", err.Error())
	}

	if !expectedBNum.Equal(total) {
		t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
			"total='%v'. ",
			expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedResultNumStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedResultNumStr, actualTotalNumstr)
	}

}

func TestBigIntMathAdd_AddIntArySeries_03(t *testing.T) {
	numStrAry := []string{
		"45.8",
		"1.45962",
		"58.71",
		"-37.62174",
		"89.8",
	}

	lenStrAry := len(numStrAry)

	expectedTotalStr := "158,14788"

	iaArray := make([]IntAry, lenStrAry)

	for i := 0; i < lenStrAry; i++ {

		ia, err := new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {

			if err != nil {
				t.Errorf("Error returned by new(IntAry).NewNumStr(numStrAry[i]) "+
					"i='%v' numStrAry[i]='%v' Error='%v' ", i, numStrAry[i], err.Error())
			}

		}

		iaArray[i] = ia
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := iaArray[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iaArray[0].SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	total, err := new(BigIntMathAdd).AddIntArySeries(
		iaArray[0],
		iaArray[1],
		iaArray[2],
		iaArray[3],
		iaArray[4])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathAdd).AddIntArySeries(...). "+
			"Error='%v' ", err.Error())
	}

	actualTotalNumstr := total.GetNumStr()

	if expectedTotalStr != actualTotalNumstr {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedTotalStr, actualTotalNumstr)
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}
