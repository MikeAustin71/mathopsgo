package mathops

import (
	"math/big"
	"testing"
)

func TestNumStrDto_NewNumStr_01(t *testing.T) {

	ePrefix := "TestNumStrDto_NewNumStr_01"

	expectedNumStr := "123.456"

	expectedBigInt := big.NewInt(123456)

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = true

	expectedNumAbsIntStr := "123"

	expectedNumAbsFracStr := "456"

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumStr \n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"  Actual expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINumNumberStr, err :=\n"+
			"  numStrDtoResultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultBigINumsAreEqual, err :=\n"+
			"  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"numStrDtoResultBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
		return
	}

	if expectedNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedAndResultBigINumsAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultBigINum is INVALID!\n"+
			"Because expectedAndResultBigINumsAreEqual == false\n"+
			"Expected numStrDtoResultBigINum = '%v'\n"+
			"  Actual numStrDtoResultBigINum = '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_NewNumStr_02(t *testing.T) {

	ePrefix := "TestNumStrDto_NewNumStr_02"

	expectedNumStr := "123456"

	expectedBigInt := big.NewInt(123456)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = false

	expectedNumAbsIntStr := "123456"

	expectedNumAbsFracStr := ""

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumStr \n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"  Actual expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINumNumberStr, err :=\n"+
			"  numStrDtoResultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultBigINumsAreEqual, err :=\n"+
			"  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"numStrDtoResultBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
		return
	}

	if expectedNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedAndResultBigINumsAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultBigINum is INVALID!\n"+
			"Because expectedAndResultBigINumsAreEqual == false\n"+
			"Expected numStrDtoResultBigINum = '%v'\n"+
			"  Actual numStrDtoResultBigINum = '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_NewNumStr_03(t *testing.T) {

	ePrefix := "TestNumStrDto_NewNumStr_03"

	expectedNumStr := "-123456"

	expectedBigInt := big.NewInt(-123456)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = false

	expectedNumAbsIntStr := "123456"

	expectedNumAbsFracStr := ""

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumStr \n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"  Actual expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINumNumberStr, err :=\n"+
			"  numStrDtoResultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultBigINumsAreEqual, err :=\n"+
			"  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"numStrDtoResultBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
		return
	}

	if expectedNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedAndResultBigINumsAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultBigINum is INVALID!\n"+
			"Because expectedAndResultBigINumsAreEqual == false\n"+
			"Expected numStrDtoResultBigINum = '%v'\n"+
			"  Actual numStrDtoResultBigINum = '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_NewNumStr_04(t *testing.T) {

	ePrefix := "TestNumStrDto_NewNumStr_04"

	expectedNumStr := "-123.456"

	expectedBigInt := big.NewInt(-123456)

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = true

	expectedNumAbsIntStr := "123"

	expectedNumAbsFracStr := "456"

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumStr \n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"  Actual expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINumNumberStr, err :=\n"+
			"  numStrDtoResultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultBigINumsAreEqual, err :=\n"+
			"  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"numStrDtoResultBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
		return
	}

	if expectedNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedAndResultBigINumsAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultBigINum is INVALID!\n"+
			"Because expectedAndResultBigINumsAreEqual == false\n"+
			"Expected numStrDtoResultBigINum = '%v'\n"+
			"  Actual numStrDtoResultBigINum = '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_NewNumStr_05(t *testing.T) {

	ePrefix := "TestNumStrDto_NewNumStr_05"

	expectedNumStr := "0.000"

	expectedBigInt := big.NewInt(0)

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = true

	expectedNumAbsIntStr := "0"

	expectedNumAbsFracStr := "000"

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumStr \n"+
			"Expected expectedBigINumStr = '%v'\n"+
			"  Actual expectedBigINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigINumNumberStr, err :=\n"+
			"  numStrDtoResultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultBigINumsAreEqual, err :=\n"+
			"  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"numStrDtoResultBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
		return
	}

	if expectedNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedAndResultBigINumsAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultBigINum is INVALID!\n"+
			"Because expectedAndResultBigINumsAreEqual == false\n"+
			"Expected numStrDtoResultBigINum = '%v'\n"+
			"  Actual numStrDtoResultBigINum = '%v'\n\n",
			ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_NewNumStrWithNumSeps_01(t *testing.T) {
	nStr := "123,456"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	nDto, err := NumStrDto{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStrWithNumSeps("+
			"nStr, expectedNumSeps) Error='%v'", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}

	actualNumSeps := nDto.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestNumStrDto_NewNumStrWithNumSeps_02(t *testing.T) {
	nStr := "123.456"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.SetDefaultsIfEmpty()

	nDto, err := NumStrDto{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStrWithNumSeps("+
			"nStr, expectedNumSeps) Error='%v'", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}

	actualNumSeps := nDto.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestNumStrDto_NewNumStrWithNumSeps_03(t *testing.T) {
	nStr := "123.456"

	expectedNumSeps := NumericSeparatorDto{}

	nDto, err := NumStrDto{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStrWithNumSeps("+
			"nStr, expectedNumSeps) Error='%v'", err.Error())
	}

	expectedNumSeps.SetDefaultsIfEmpty()

	actualNumStr := nDto.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}

	actualNumSeps := nDto.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestNumStrDto_NewUint_01(t *testing.T) {

	intNum := uint(7)
	precision := uint(0)

	expectedStr := "7"

	nDto := NumStrDto{}.NewUint(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint_02(t *testing.T) {

	intNum := uint(7)
	precision := uint(1)

	expectedStr := "7.0"

	nDto := NumStrDto{}.NewUint(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint_03(t *testing.T) {

	intNum := uint(7)
	precision := uint(3)

	expectedStr := "7.000"

	nDto := NumStrDto{}.NewUint(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint_04(t *testing.T) {

	intNum := uint(792)
	precision := uint(3)

	expectedStr := "792.000"

	nDto := NumStrDto{}.NewUint(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint_05(t *testing.T) {

	intNum := uint(792)
	precision := uint(0)

	expectedStr := "792"

	nDto := NumStrDto{}.NewUint(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUintExponent_01(t *testing.T) {
	intNum := uint(7)
	exponent := 3

	expectedStr := "7.000"

	nDto := NumStrDto{}.NewUintExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUintExponent_02(t *testing.T) {
	intNum := uint(7123)
	exponent := -3

	expectedStr := "7.123"

	nDto := NumStrDto{}.NewUintExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUintExponent_03(t *testing.T) {
	intNum := uint(872)
	exponent := 3

	expectedStr := "872.000"

	nDto := NumStrDto{}.NewUintExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUintExponent_04(t *testing.T) {
	intNum := uint(872123)
	exponent := -3

	expectedStr := "872.123"

	nDto := NumStrDto{}.NewUintExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUintExponent_05(t *testing.T) {
	intNum := uint(72)
	exponent := 0

	expectedStr := "72"

	nDto := NumStrDto{}.NewUintExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUintExponent_06(t *testing.T) {
	intNum := uint(472)
	exponent := 0

	expectedStr := "472"

	nDto := NumStrDto{}.NewUintExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32_01(t *testing.T) {

	intNum := uint32(7)
	precision := uint(0)

	expectedStr := "7"

	nDto := NumStrDto{}.NewUint32(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32_02(t *testing.T) {

	intNum := uint32(7)
	precision := uint(1)

	expectedStr := "7.0"

	nDto := NumStrDto{}.NewUint32(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32_03(t *testing.T) {

	intNum := uint32(7)
	precision := uint(3)

	expectedStr := "7.000"

	nDto := NumStrDto{}.NewUint32(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32_04(t *testing.T) {

	intNum := uint32(792)
	precision := uint(3)

	expectedStr := "792.000"

	nDto := NumStrDto{}.NewUint32(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32_05(t *testing.T) {

	intNum := uint32(792)
	precision := uint(0)

	expectedStr := "792"

	nDto := NumStrDto{}.NewUint32(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32Exponent_01(t *testing.T) {
	intNum := uint32(7)
	exponent := 3

	expectedStr := "7.000"

	nDto := NumStrDto{}.NewUint32Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32Exponent_02(t *testing.T) {
	intNum := uint32(7123)
	exponent := -3

	expectedStr := "7.123"

	nDto := NumStrDto{}.NewUint32Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32Exponent_03(t *testing.T) {
	intNum := uint32(872)
	exponent := 3

	expectedStr := "872.000"

	nDto := NumStrDto{}.NewUint32Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32Exponent_04(t *testing.T) {
	intNum := uint32(872123)
	exponent := -3

	expectedStr := "872.123"

	nDto := NumStrDto{}.NewUint32Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32Exponent_05(t *testing.T) {
	intNum := uint32(72)
	exponent := 0

	expectedStr := "72"

	nDto := NumStrDto{}.NewUint32Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint32Exponent_06(t *testing.T) {
	intNum := uint32(472)
	exponent := 0

	expectedStr := "472"

	nDto := NumStrDto{}.NewUint32Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64_01(t *testing.T) {

	intNum := uint64(7)
	precision := uint(0)

	expectedStr := "7"

	nDto := NumStrDto{}.NewUint64(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64_02(t *testing.T) {

	intNum := uint64(7)
	precision := uint(1)

	expectedStr := "7.0"

	nDto := NumStrDto{}.NewUint64(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64_03(t *testing.T) {

	intNum := uint64(7)
	precision := uint(3)

	expectedStr := "7.000"

	nDto := NumStrDto{}.NewUint64(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64_04(t *testing.T) {

	intNum := uint64(792)
	precision := uint(3)

	expectedStr := "792.000"

	nDto := NumStrDto{}.NewUint64(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64_05(t *testing.T) {

	intNum := uint64(792)
	precision := uint(0)

	expectedStr := "792"

	nDto := NumStrDto{}.NewUint64(intNum, precision)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64Exponent_01(t *testing.T) {
	intNum := uint64(7)
	exponent := 3

	expectedStr := "7.000"

	nDto := NumStrDto{}.NewUint64Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64Exponent_02(t *testing.T) {
	intNum := uint64(7123)
	exponent := -3

	expectedStr := "7.123"

	nDto := NumStrDto{}.NewUint64Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64Exponent_03(t *testing.T) {
	intNum := uint64(872)
	exponent := 3

	expectedStr := "872.000"

	nDto := NumStrDto{}.NewUint64Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64Exponent_04(t *testing.T) {
	intNum := uint64(872123)
	exponent := -3

	expectedStr := "872.123"

	nDto := NumStrDto{}.NewUint64Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64Exponent_05(t *testing.T) {
	intNum := uint64(72)
	exponent := 0

	expectedStr := "72"

	nDto := NumStrDto{}.NewUint64Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewUint64Exponent_06(t *testing.T) {
	intNum := uint64(472)
	exponent := 0

	expectedStr := "472"

	nDto := NumStrDto{}.NewUint64Exponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewZero_01(t *testing.T) {

	expectedStr := "0"

	nDto := NumStrDto{}.NewZero(0)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}
}

func TestNumStrDto_NewZero_02(t *testing.T) {

	expectedStr := "0.00"

	nDto := NumStrDto{}.NewZero(2)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}
}

func TestNumStrDto_NewZero_03(t *testing.T) {

	expectedStr := "0.0000"

	nDto := NumStrDto{}.NewZero(4)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}
}

func TestNumStrDto_ParseNumStr_01(t *testing.T) {
	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). "+
			"nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
			precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_02(t *testing.T) {
	nStr := "123456"
	iStr := "123456"
	fracStr := ""
	signVal := 1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
			precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_03(t *testing.T) {
	nStr := "-123456"
	iStr := "123456"
	fracStr := ""
	signVal := -1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
			precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_04(t *testing.T) {
	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
			precision, nDto.GetPrecisionUint())
	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_05(t *testing.T) {
	nStr := "-000.000"
	nStrOut := "0.000"
	iStr := "0"
	fracStr := "000"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
			precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_06(t *testing.T) {
	nStr := "-123.4567#"
	nStrOut := "-123.4567"
	iStr := "123"
	fracStr := "4567"
	signVal := -1
	precision := uint(4)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
			precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_07(t *testing.T) {
	nStr := "-123.4#-567#"
	nStrOut := "-123.4567"
	iStr := "123"
	fracStr := "4567"
	signVal := -1
	precision := uint(4)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
			precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseSignedBigInt_01(t *testing.T) {

	signedAbsNumStr := "-123456789"
	absAllNumStr := "123456789"
	nStr := "-123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := -1
	sBigInt, isOk := big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("bigInt.SetString(signedAbsNumStr,10) conversion failed! nStr= '%v' ", signedAbsNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseSignedBigInt(sBigInt, precision)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseSignedBigInt(sBigInt, precision). sBigInt= '%v' Error= %v", sBigInt.String(), err)
	}

	nDto := n1.CopyOut()

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsAllNumRunes())

	if absAllNumStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", absAllNumStr, s)

	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
			precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseSignedBigInt_02(t *testing.T) {

	signedAbsNumStr := "123456789"
	absAllNumStr := "123456789"
	nStr := "123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := 1
	sBigInt, isOk := big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("bigInt.SetString(signedAbsNumStr,10) conversion failed! nStr= '%v' ", signedAbsNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseSignedBigInt(sBigInt, precision)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseSignedBigInt(sBigInt, precision). sBigInt= '%v' Error= %v", sBigInt.String(), err)
	}

	nDto := n1.CopyOut()

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsAllNumRunes())

	if absAllNumStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", absAllNumStr, s)

	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
			precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}
