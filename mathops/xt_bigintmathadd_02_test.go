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
			"nDto1.GetThisPointer(), dec2.GetThisPointer())\n"+
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
	//expectedFinalResult := "124443.912456"
	expectedResultStr := "124443,912456"

	expectedBigIntResultStr := "124443912456"

	expectedPrecision := uint(6)

	expectedSign := 1

	biExpectedResult, oK := big.NewInt(0).SetString(expectedBigIntResultStr, 10)

	if !oK {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"biExpectedResult, oK := big.NewInt(0).SetString(expectedBigIntResultStr, 10)\n"+
			"expectedResultStr='%v'\n\n", ePrefix, expectedBigIntResultStr)
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

	if expectedResultStr != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedResultStr,
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

	err = expectedBNum.IsValid("Validating expectedBNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBNum.IsValid(\"Validating expectedBNum\")\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedTotalStr != expectedResultNumStr {
		t.Errorf("%v\n"+
			"Error: expectedBNum Initialization FAILED!\n"+
			"Because expectedTotalStr != expectedResultNumStr\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"  Actual expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	inumMgrAry := make([]INumMgr, lenStrAry)

	var iNumMgrNumStr string

	for i := 0; i < lenStrAry; i++ {

		dec, err := new(Decimal).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"dec, err = new(Decimal).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
			return
		}

		inumMgrAry[i] = &dec

		iNumMgrNumStr, err = inumMgrAry[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iNumMgrNumStr, err = inumMgrAry[i].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, err.Error())
			return
		}

		if iNumMgrNumStr != numStrAry[i] {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because iNumMgrNumStr != numStrAry[i]\n"+
				"Expected numStrAry[i] = '%v'\n"+
				"  Actual numStrAry[i] = '%v'\n\n",
				ePrefix, iNumMgrNumStr, numStrAry[i])

			return

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

	err = total.IsValid("Validating total")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = total.IsValid(\"Validating total\")\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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
	var nDto NumStrDto
	var ia IntAry

	for i := 0; i < lenStrAry; i++ {

		if i < 2 {

			dec, err := new(Decimal).NewNumStr(numStrAry[i])

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
			"Expected BNum Int = '%v'\n"+
			"Instead, BNum Int = '%v'\n\n",
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
	ePrefix := "TestBigIntMathAdd_AddINumMgrOutputToArray_01"
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
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iNumMgrAddend, err := new(Decimal).NewNumStr(addendStr)\n"+
			"addendStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, addendStr, err.Error())
		return
	}

	iNumMgrAddendNumStr, err := iNumMgrAddend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iNumMgrAddendNumStr, err := iNumMgrAddend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(numMgrStrs)
	iNumMgrArray := make([]INumMgr, lenArray)
	var nDto NumStrDto

	for i := 0; i < lenArray; i++ {

		nDto, err = new(NumStrDto).NewNumStr(numMgrStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDto, err = new(NumStrDto).NewNumStr(numMgrStrs[%d])\n"+
				"numMgrStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numMgrStrs[i], err.Error())
			return
		}

		iNumMgrArray[i] = &nDto
	}

	result, err := new(BigIntMathAdd).AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrArray)\n"+
			"iNumMgrAddendNum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, iNumMgrAddendNumStr, err.Error())
		return
	}

	var resultNumStr string

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected result[%d] = '%v'\n"+
				"Instead, result[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				resultNumStr)

			return
		}
	}

	return
}

func TestBigIntMathAdd_AddINumMgrOutputToArray_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgrOutputToArray_02"
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
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iNumMgrAddend, err := new(BigIntNum).NewNumStr(addendStr)\n"+
			"addendStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, addendStr, err.Error())
		return
	}

	iNumMgrAddendNumStr, err := iNumMgrAddend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iNumMgrAddendNumStr, err := iNumMgrAddend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(numMgrStrs)
	iNumMgrsArray := make([]INumMgr, lenArray)
	var ia IntAry

	for i := 0; i < lenArray; i++ {

		ia, err = new(IntAry).NewNumStr(numMgrStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStr(numMgrStrs[%d])\n"+
				"numMgrStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numMgrStrs[i], err.Error())
			return
		}

		iNumMgrsArray[i] = &ia

	}

	result, err := new(BigIntMathAdd).AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrsArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrsArray)\n"+
			"iNumMgrAddend= '%v'\n"+
			"Error='%v'\n\n", ePrefix, iNumMgrAddendNumStr, err.Error())
		return
	}

	var resultNumStr string

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected result[%d] = '%v'\n"+
				"Instead, result[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				resultNumStr)

			return
		}
	}
	return
}

func TestBigIntMathAdd_AddINumMgrOutputToArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddINumMgrOutputToArray_03"

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
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iNumMgrAddend, err := new(Decimal).NewNumStr(addendStr)\n"+
			"addendStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, addendStr, err.Error())
		return
	}

	iNumMgrAddendNumStr, err := iNumMgrAddend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iNumMgrAddendNumStr, err := iNumMgrAddend.GetNumStr()\n"+
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

	err = iNumMgrAddend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iNumMgrAddend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	lenArray := len(numMgrStrs)
	iNumMgrArray := make([]INumMgr, lenArray)
	var nDto NumStrDto

	for i := 0; i < lenArray; i++ {

		nDto, err = new(NumStrDto).NewNumStr(numMgrStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDto, err = new(NumStrDto).NewNumStr(numMgrStrs[%d])\n"+
				"numMgrStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numMgrStrs[i], err.Error())
			return
		}

		iNumMgrArray[i] = &nDto
	}

	result, err := new(BigIntMathAdd).AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddINumMgrOutputToArray(&iNumMgrAddend, iNumMgrArray)\n"+
			"iNumMgrAddend= '%v'\n"+
			"Error='%v'\n\n", ePrefix, iNumMgrAddendNumStr, err.Error())
		return
	}

	var resultNumStr string
	var actualNumSeps NumericSeparatorDto

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected result[%d] = '%v'\n"+
				"Instead, result[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				resultNumStr)

			return
		}

		actualNumSeps, err = result[j].GetNumericSeparatorsDto()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"actualNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected actualNumSeps = '%v'\n"+
				"Instead, actualNumSeps = '%v'\n"+
				"Index Cycle= '%v'\n\n",
				ePrefix,
				expectedNumSeps.String(),
				actualNumSeps.String(),
				j)

			return
		}
	}
	return
}

func TestBigIntMathAdd_AddINumMgrSeries_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgrSeries_01"
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedTotalStr := "158.14788"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)\n"+
			"expectedTotalStr= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			err.Error())
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

	if expectedTotalStr != expectedResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, expectedResultNumStr = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	dec1, err := new(Decimal).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec1, err := new(Decimal).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n1Str,
			err.Error())
		return
	}

	nDto2, err := new(NumStrDto).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto2, err := new(NumStrDto).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n2Str,
			err.Error())
		return
	}

	ia3, err := new(IntAry).NewNumStr(n3Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia3, err := new(IntAry).NewNumStr(n3Str)\n"+
			"n3Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n3Str,
			err.Error())
		return
	}

	bigINum4, err := new(BigIntNum).NewNumStr(n4Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINum4, err := new(BigIntNum).NewNumStr(n4Str)\n"+
			"n4Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n4Str,
			err.Error())
		return
	}

	dec5, err := new(Decimal).NewNumStr(n5Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec5, err := new(Decimal).NewNumStr(n5Str)\n"+
			"n5Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n5Str,
			err.Error())
		return
	}

	total, err := new(BigIntMathAdd).AddINumMgrSeries(&dec1, &nDto2, &ia3, &bigINum4, &dec5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddINumMgrSeries(&dec1, &nDto2, &ia3, &bigINum4, &dec5)\n"+
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
			"expectedBNum != total"+
			"Expected total = '%v'\n"+
			"Instead, total = '%v'\n\n",
			ePrefix,
			expectedBNum.bigInt.Text(10),
			total.bigInt.Text(10))
		return
	}

	if expectedResultNumStr != totalNumstr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total NumStr = '%v'\n"+
			"Instead, total NumStr = '%v'\n\n",
			ePrefix,
			expectedResultNumStr,
			totalNumstr)
	}

	return
}

func TestBigIntMathAdd_AddINumMgrSeries_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgrSeries_02"
	n1Str := "-978425.648941"
	n2Str := "33.12"
	n3Str := "-804.1"
	n4Str := "32567"
	n5Str := "-41.859"
	expectedTotalStr := "-946671.487941"

	expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)\n"+
			"expectedTotalStr= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			expectedTotalStr,
			err.Error())
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

	if expectedTotalStr != expectedResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"'expectedBNum' does NOT match 'expectedTotalStr'\n"+
			"Expected expectedResultNumStr = '%v'\n"+
			"Instead, fixedDecNumStr2 = '%v'\n\n",
			ePrefix, expectedTotalStr, expectedResultNumStr)

		return
	}

	dec1, err := new(Decimal).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec1, err := new(Decimal).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n1Str,
			err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n2Str,
			err.Error())
		return
	}

	nDto3, err := new(NumStrDto).NewNumStr(n3Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto3, err := new(NumStrDto).NewNumStr(n3Str)\n"+
			"n3Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n3Str,
			err.Error())
		return
	}

	dec4, err := new(Decimal).NewNumStr(n4Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec4, err := new(Decimal).NewNumStr(n4Str)\n"+
			"n4Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n4Str,
			err.Error())
		return
	}

	bINum5, err := new(BigIntNum).NewNumStr(n5Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum5, err := new(BigIntNum).NewNumStr(n5Str)\n"+
			"n5Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n4Str,
			err.Error())
		return
	}

	total, err := new(BigIntMathAdd).AddINumMgrSeries(&dec1, &ia2, &nDto3, &dec4, &bINum5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddINumMgrSeries(&dec1, &ia2, &nDto3, &dec4, &bINum5)\n"+
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
			"expectedBNum != total"+
			"Expected total = '%v'\n"+
			"Instead, total = '%v'\n\n",
			ePrefix,
			expectedBNum.bigInt.Text(10),
			total.bigInt.Text(10))
		return
	}

	if expectedResultNumStr != totalNumstr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total NumStr = '%v'\n"+
			"Instead, total NumStr = '%v'\n\n",
			ePrefix,
			expectedResultNumStr,
			totalNumstr)
	}

	return
}

func TestBigIntMathAdd_AddINumMgrSeries_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddINumMgrSeries_03"
	n1Str := "45.8"
	n2Str := "1.45962"
	n3Str := "58.71"
	n4Str := "-37.62174"
	n5Str := "89.8"
	expectedTotalStr := "158,14788"

	dec1, err := new(Decimal).NewNumStr(n1Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec1, err := new(Decimal).NewNumStr(n1Str)\n"+
			"n1Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n1Str,
			err.Error())
		return
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
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dec1.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	nDto2, err := new(NumStrDto).NewNumStr(n2Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto2, err := new(NumStrDto).NewNumStr(n2Str)\n"+
			"n2Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n2Str,
			err.Error())
		return
	}

	ia3, err := new(IntAry).NewNumStr(n3Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia3, err := new(IntAry).NewNumStr(n3Str)\n"+
			"n3Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n3Str,
			err.Error())
		return
	}

	bigINum4, err := new(BigIntNum).NewNumStr(n4Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINum4, err := new(BigIntNum).NewNumStr(n4Str)\n"+
			"n4Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n4Str,
			err.Error())
		return
	}

	dec5, err := new(Decimal).NewNumStr(n5Str)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec5, err := new(Decimal).NewNumStr(n5Str)\n"+
			"n5Str= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			n5Str,
			err.Error())
		return
	}

	total, err := new(BigIntMathAdd).AddINumMgrSeries(&dec1, &nDto2, &ia3, &bigINum4, &dec5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddINumMgrSeries(&dec1, &nDto2, &ia3, &bigINum4, &dec5)\n"+
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

	if expectedTotalStr != totalNumstr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total NumStr = '%v'\n"+
			"Instead, total NumStr = '%v'\n\n",
			ePrefix,
			expectedTotalStr,
			totalNumstr)
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

func TestBigIntMathAdd_AddIntAry_01(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddIntAry_01"
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

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error='%v'\n\n",
			ia1NumStr,
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

func TestBigIntMathAdd_AddIntAry_02(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddIntAry_02"
	n1Str := "123456.789"
	n2Str := "-987.123456"
	expectedResultStr := "122469665544"
	expectedPrecision := uint(6)
	expectedFinalResult := "122469.665544"

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

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddINumMgr(&ia1, &dec2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error='%v'\n\n",
			ia1NumStr,
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

func TestBigIntMathAdd_AddIntAry_03(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddIntAry_03"
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

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error='%v'\n\n",
			ia1NumStr,
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

func TestBigIntMathAdd_AddIntAry_04(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddIntAry_04"
	n1Str := "-123456.789"
	n2Str := "-987.123456"
	// Result := -124443.912456
	expectedFinalResult := "-124443.912456"
	expectedResultStr := "-124443912456"
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

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error='%v'\n\n",
			ia1NumStr,
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

func TestBigIntMathAdd_AddIntAry_05(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddIntAry_05"
	n1Str := "123456.789"
	n2Str := "987.123456"

	// Result = 	124443.912456
	expectedResultStr := "124443,912456"
	expectedPrecision := uint(6)
	expectedSign := 1

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

	result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddIntAry(ia1, ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error='%v'\n\n",
			ia1NumStr,
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

	if expectedResultStr != actualResultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected actualResultNumStr = '%v'\n"+
			"Instead, actualResultNumStr = '%v'\n\n",
			ePrefix,
			expectedResultStr,
			actualResultNumStr)
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

func TestBigIntMathAdd_AddIntAryArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddIntAryArray_01"
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

	iaArray := make([]IntAry, lenStrAry)

	var ia IntAry

	for i := 0; i < lenStrAry; i++ {

		ia, err = new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
			return
		}

		iaArray[i] = ia
	}

	total, err := new(BigIntMathAdd).AddIntAryArray(iaArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddIntAryArray(iaArray)\n"+
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
			"Expected BNum Int = '%v'\n"+
			"Instead, BNum Int = '%v'\n\n",
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

func TestBigIntMathAdd_AddIntAryArray_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddIntAryArray_02"
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

	iaArray := make([]IntAry, lenStrAry)
	var ia IntAry

	for i := 0; i < lenStrAry; i++ {

		ia, err = new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
			return
		}

		iaArray[i] = ia

	}

	total, err := new(BigIntMathAdd).AddIntAryArray(iaArray)

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
			"Expected BNum Int = '%v'\n"+
			"Instead, BNum Int = '%v'\n\n",
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

func TestBigIntMathAdd_AddIntAryArray_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddIntAryArray_03"
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
	var ia IntAry
	var err error

	for i := 0; i < lenStrAry; i++ {

		ia, err = new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
			return
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

	err = iaArray[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = inumMgrAry[0].SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	total, err := new(BigIntMathAdd).AddIntAryArray(iaArray)

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

func TestBigIntMathAdd_AddIntAryOutputToArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddIntAryOutputToArray_01"
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
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaAddend, err := new(IntAry).NewNumStr(addendStr)\n"+
			"addendStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, addendStr, err.Error())
		return
	}

	iaAddendNumStr, err := iaAddend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaAddendNumStr, err := iaAddend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(iaNumStrs)
	iaArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		iaArray[i], err = new(IntAry).NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaArray[%d], err = new(IntAry).NewNumStr(iaNumStrs[%d])\n"+
				"iaNumStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, i, iaNumStrs[i], err.Error())
			return
		}

	}

	result, err := new(BigIntMathAdd).AddIntAryOutputToArray(iaAddend, iaArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddIntAryOutputToArray(iaAddend, iaArray)\n"+
			"iaAddend= '%v'\n"+
			"Error='%v'\n\n", ePrefix, iaAddendNumStr, err.Error())
		return
	}

	var resultNumStr string

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected result[%d] = '%v'\n"+
				"Instead, result[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				resultNumStr)

			return
		}
	}
	return
}

func TestBigIntMathAdd_AddIntAryOutputToArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddIntAryOutputToArray_02"

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
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaAddend, err := new(IntAry).NewNumStr(addendStr)\n"+
			"addendStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, addendStr, err.Error())
		return
	}

	iaAddendNumStr, err := iaAddend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaAddendNumStr, err := iaAddend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenArray := len(iaNumStrs)
	decsArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		decsArray[i], err = new(IntAry).NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"decsArray[%d], err = new(IntAry).NewNumStr(iaNumStrs[%d])\n"+
				"iaNumStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, i, iaNumStrs[i], err.Error())
			return
		}
	}

	result, err := new(BigIntMathAdd).AddIntAryOutputToArray(iaAddend, decsArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddIntAryOutputToArray(iaAddend, decsArray)\n"+
			"iaAddend= '%v'\n"+
			"Error='%v'\n\n", ePrefix, iaAddendNumStr, err.Error())
		return
	}

	var resultNumStr string

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected result[%d] = '%v'\n"+
				"Instead, result[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				resultNumStr)

			return
		}
	}

	return
}

func TestBigIntMathAdd_AddIntAryOutputToArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddIntAryOutputToArray_03"

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
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaAddend, err := new(IntAry).NewNumStr(addendStr)\n"+
			"addendStr= '%v'\n"+
			"Error='%v'\n\n", ePrefix, addendStr, err.Error())
		return
	}

	iaAddendNumStr, err := iaAddend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaAddendNumStr, err := iaAddend.GetNumStr()\n"+
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

	err = iaAddend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaAddend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	lenArray := len(iaNumStrs)
	iaArray := make([]IntAry, lenArray)

	for i := 0; i < lenArray; i++ {

		iaArray[i], err = new(IntAry).NewNumStr(iaNumStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"iaArray[%d], err = new(IntAry).NewNumStr(iaNumStrs[%d])\n"+
				"iaNumStrs[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, i, iaNumStrs[i], err.Error())
			return
		}
	}

	result, err := new(BigIntMathAdd).AddIntAryOutputToArray(iaAddend, iaArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathAdd).AddIntAryOutputToArray(iaAddend, iaArray)\n"+
			"iaAddend= '%v'\n"+
			"Error='%v'\n\n", ePrefix, iaAddendNumStr, err.Error())
		return
	}

	var resultNumStr string
	var actualNumSeps NumericSeparatorDto

	for j := 0; j < lenArray; j++ {

		resultNumStr, err = result[j].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if expectedNumStrs[j] != resultNumStr {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected result[%d] = '%v'\n"+
				"Instead, result[%d] = '%v'\n\n",
				ePrefix,
				j,
				expectedNumStrs[j],
				j,
				resultNumStr)

			return
		}

		actualNumSeps, err = result[j].GetNumericSeparatorsDto()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"actualNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
				"Error='%v'\n\n", ePrefix, j, err.Error())
			return
		}

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Expected actualNumSeps = '%v'\n"+
				"Instead, actualNumSeps = '%v'\n"+
				"Index Cycle= '%v'\n\n",
				ePrefix,
				expectedNumSeps.String(),
				actualNumSeps.String(),
				j)

			return
		}
	}

	return
}

func TestBigIntMathAdd_AddIntArySeries_01(t *testing.T) {

	ePrefix := "TestBigIntMathAdd_AddIntArySeries_01"
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

	iaArray := make([]IntAry, lenStrAry)

	var ia IntAry

	for i := 0; i < lenStrAry; i++ {

		ia, err = new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
			return
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
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddIntArySeries(\n"+
			"iaArray[0], iaArray[1], iaArray[2], iaArray[3], iaArray[4])\n"+
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
			"expectedBNum != total"+
			"Expected total = '%v'\n"+
			"Instead, total = '%v'\n\n",
			ePrefix,
			expectedBNum.bigInt.Text(10),
			total.bigInt.Text(10))
		return
	}

	if expectedResultNumStr != totalNumstr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total NumStr = '%v'\n"+
			"Instead, total NumStr = '%v'\n\n",
			ePrefix,
			expectedResultNumStr,
			totalNumstr)
	}

	return
}

func TestBigIntMathAdd_AddIntArySeries_02(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddIntArySeries_02"
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

	iaArray := make([]IntAry, lenStrAry)
	var ia IntAry

	for i := 0; i < lenStrAry; i++ {

		ia, err = new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err = new(IntAry).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
			return
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
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddIntArySeries(\n"+
			"iaArray[0], iaArray[1], iaArray[2], iaArray[3], iaArray[4])\n"+
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
			"expectedBNum != total"+
			"Expected total = '%v'\n"+
			"Instead, total = '%v'\n\n",
			ePrefix,
			expectedBNum.bigInt.Text(10),
			total.bigInt.Text(10))
		return
	}

	if expectedResultNumStr != totalNumstr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total NumStr = '%v'\n"+
			"Instead, total NumStr = '%v'\n\n",
			ePrefix,
			expectedResultNumStr,
			totalNumstr)
	}

	return
}

func TestBigIntMathAdd_AddIntArySeries_03(t *testing.T) {
	ePrefix := "TestBigIntMathAdd_AddIntArySeries_03"
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

	var ia IntAry
	var err error

	for i := 0; i < lenStrAry; i++ {

		ia, err = new(IntAry).NewNumStr(numStrAry[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia, err := new(IntAry).NewNumStr(numStrAry[%d])\n"+
				"numStrAry[%d]= '%v'\n"+
				"Error='%v'\n\n", ePrefix, i, i, numStrAry[i], err.Error())
			return
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

	err = iaArray[0].SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaArray[0].SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	total, err := new(BigIntMathAdd).AddIntArySeries(
		iaArray[0],
		iaArray[1],
		iaArray[2],
		iaArray[3],
		iaArray[4])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(BigIntMathAdd).AddIntArySeries(\n"+
			"iaArray[0], iaArray[1], iaArray[2], iaArray[3], iaArray[4])\n"+
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

	if expectedTotalStr != totalNumstr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected total NumStr = '%v'\n"+
			"Instead, total NumStr = '%v'\n\n",
			ePrefix,
			expectedTotalStr,
			totalNumstr)
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
