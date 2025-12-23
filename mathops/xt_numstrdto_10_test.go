package mathops

import (
	"math/big"
	"testing"
)

func TestNumStrDto_ShiftPrecisionRight_01(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_01"

	inputNumStr := "123456.789"

	inputPrecisionUint := uint(3)

	expectedNumStr := "123456789"

	expectedBigInt := big.NewInt(123456789)

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

	expectedNumAbsIntStr := "123456789"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "123456789"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_02(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_02"

	inputNumStr := "123456.789"

	inputPrecisionUint := uint(2)

	expectedNumStr := "12345678.9"

	expectedBigInt := big.NewInt(123456789)

	expectedPrecisionInt := 1

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = true

	expectedNumAbsIntStr := "12345678"

	expectedNumAbsFracStr := "9"

	expectedAbsAllRunesNumStr := "123456789"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_03(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_03"

	inputNumStr := "123456.789"

	inputPrecisionUint := uint(6)

	expectedNumStr := "123456789000"

	expectedBigInt := big.NewInt(123456789000)

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

	expectedNumAbsIntStr := "123456789000"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "123456789000"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_04(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_04"

	inputNumStr := "123456789"

	inputPrecisionUint := uint(6)

	expectedNumStr := "123456789000000"

	expectedBigInt := big.NewInt(123456789000000)

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

	expectedNumAbsIntStr := "123456789000000"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "123456789000000"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_05(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_05"

	inputNumStr := "123"

	inputPrecisionUint := uint(5)

	expectedNumStr := "12300000"

	expectedBigInt := big.NewInt(12300000)

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

	expectedNumAbsIntStr := "12300000"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "12300000"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_06(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_06"

	inputNumStr := "0"

	inputPrecisionUint := uint(3)

	expectedNumStr := "0"

	expectedBigInt := big.NewInt(0)

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

	expectedNumAbsIntStr := "0"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "0"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_07(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_07"

	inputNumStr := "123456.789"

	inputPrecisionUint := uint(0)

	expectedNumStr := "123456.789"

	expectedBigInt := big.NewInt(123456789)

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

	expectedNumAbsIntStr := "123456"

	expectedNumAbsFracStr := "789"

	expectedAbsAllRunesNumStr := "123456789"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_08(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_08"

	inputNumStr := "-123456.789"

	inputPrecisionUint := uint(0)

	expectedNumStr := "-123456.789"

	expectedBigInt := big.NewInt(-123456789)

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

	expectedNumAbsIntStr := "123456"

	expectedNumAbsFracStr := "789"

	expectedAbsAllRunesNumStr := "123456789"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_09(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_09"

	inputNumStr := "-123456.789"

	inputPrecisionUint := uint(3)

	expectedNumStr := "-123456789"

	expectedBigInt := big.NewInt(-123456789)

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

	expectedNumAbsIntStr := "123456789"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "123456789"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_10(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_10"

	inputNumStr := "-123456789"

	inputPrecisionUint := uint(6)

	expectedNumStr := "-123456789000000"

	expectedBigInt := big.NewInt(-123456789000000)

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

	expectedNumAbsIntStr := "123456789000000"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "123456789000000"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_ShiftPrecisionRight_11(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_11"

	inputNumStr := "0.123456789"

	inputPrecisionUint := uint(9)

	expectedNumStr := "123456789"

	expectedBigInt := big.NewInt(123456789)

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

	expectedNumAbsIntStr := "123456789"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "123456789"

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

	numStrDtoResult, err := new(NumStrDto).NewPtr().ShiftPrecisionRight(inputNumStr, inputPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
			"  ShiftPrecisionRight(inputNumStr, inputPrecisionUint)\n"+
			"inputNumStr= '%v'\n"+
			"inputPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			inputNumStr,
			inputPrecisionUint,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_SubtractNumStrs_01(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_01"

	inputNumStr01 := "67.521"

	inputNumStr02 := "-6"

	expectedNumStr := "73.521"

	expectedBigInt := big.NewInt(73521)

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

	expectedNumAbsIntStr := "73"

	expectedNumAbsFracStr := "521"

	expectedAbsAllRunesNumStr := "73521"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDto01, err := new(NumStrDto).ParseNumStr(inputNumStr01)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr01)\n"+
			"inputNumStr01= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr01, err.Error())
		return
	}

	err = numStrDto01.IsValid("Validating numStrDto01")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto01.IsValid('Validating numStrDto01')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto01NumStr, err := numStrDto01.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01NumStr, err := numStrDto01.GetNumStr()\n"+
			"numStrDto01 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr01 != numStrDto01NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto01 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr01 != numStrDto01NumStr\n"+
			"Expected numStrDto01NumStr = '%v'\n"+
			"  Actual numStrDto01NumStr = '%v'\n\n",
			ePrefix, inputNumStr01, numStrDto01NumStr)

		return
	}

	numStrDto02, err := new(NumStrDto).ParseNumStr(inputNumStr02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr02)\n"+
			"inputNumStr02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr02, err.Error())
		return
	}

	err = numStrDto02.IsValid("Validating numStrDto02")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto02.IsValid('Validating numStrDto02')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto02NumStr, err := numStrDto02.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02NumStr, err := numStrDto02.GetNumStr()\n"+
			"numStrDto02 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr02 != numStrDto02NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto02 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr02 != numStrDto02NumStr\n"+
			"Expected numStrDto02NumStr = '%v'\n"+
			"  Actual numStrDto02NumStr = '%v'\n\n",
			ePrefix, inputNumStr02, numStrDto02NumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).SubtractNumStrs(numStrDto01, numStrDto02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  SubtractNumStrs(numStrDto01, numStrDto02)\n"+
			"numStrDto01= '%v'\n"+
			"numStrDto02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto01NumStr,
			numStrDto02NumStr,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
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

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_SubtractNumStrs_02(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_02"

	inputNumStr01 := "-67.521"

	inputNumStr02 := "6"

	expectedNumStr := "-73.521"

	expectedBigInt := big.NewInt(-73521)

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

	expectedNumAbsIntStr := "73"

	expectedNumAbsFracStr := "521"

	expectedAbsAllRunesNumStr := "73521"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDto01, err := new(NumStrDto).ParseNumStr(inputNumStr01)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr01)\n"+
			"inputNumStr01= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr01, err.Error())
		return
	}

	err = numStrDto01.IsValid("Validating numStrDto01")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto01.IsValid('Validating numStrDto01')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto01NumStr, err := numStrDto01.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01NumStr, err := numStrDto01.GetNumStr()\n"+
			"numStrDto01 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr01 != numStrDto01NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto01 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr01 != numStrDto01NumStr\n"+
			"Expected numStrDto01NumStr = '%v'\n"+
			"  Actual numStrDto01NumStr = '%v'\n\n",
			ePrefix, inputNumStr01, numStrDto01NumStr)

		return
	}

	numStrDto02, err := new(NumStrDto).ParseNumStr(inputNumStr02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr02)\n"+
			"inputNumStr02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr02, err.Error())
		return
	}

	err = numStrDto02.IsValid("Validating numStrDto02")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto02.IsValid('Validating numStrDto02')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto02NumStr, err := numStrDto02.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02NumStr, err := numStrDto02.GetNumStr()\n"+
			"numStrDto02 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr02 != numStrDto02NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto02 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr02 != numStrDto02NumStr\n"+
			"Expected numStrDto02NumStr = '%v'\n"+
			"  Actual numStrDto02NumStr = '%v'\n\n",
			ePrefix, inputNumStr02, numStrDto02NumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).SubtractNumStrs(numStrDto01, numStrDto02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  SubtractNumStrs(numStrDto01, numStrDto02)\n"+
			"numStrDto01= '%v'\n"+
			"numStrDto02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto01NumStr,
			numStrDto02NumStr,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
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

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_SubtractNumStrs_03(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_03"

	inputNumStr01 := "67.521"

	inputNumStr02 := "691.1"

	expectedNumStr := "-623.579"

	expectedBigInt := big.NewInt(-623579)

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

	expectedNumAbsIntStr := "623"

	expectedNumAbsFracStr := "579"

	expectedAbsAllRunesNumStr := "623579"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDto01, err := new(NumStrDto).ParseNumStr(inputNumStr01)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr01)\n"+
			"inputNumStr01= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr01, err.Error())
		return
	}

	err = numStrDto01.IsValid("Validating numStrDto01")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto01.IsValid('Validating numStrDto01')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto01NumStr, err := numStrDto01.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01NumStr, err := numStrDto01.GetNumStr()\n"+
			"numStrDto01 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr01 != numStrDto01NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto01 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr01 != numStrDto01NumStr\n"+
			"Expected numStrDto01NumStr = '%v'\n"+
			"  Actual numStrDto01NumStr = '%v'\n\n",
			ePrefix, inputNumStr01, numStrDto01NumStr)

		return
	}

	numStrDto02, err := new(NumStrDto).ParseNumStr(inputNumStr02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr02)\n"+
			"inputNumStr02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr02, err.Error())
		return
	}

	err = numStrDto02.IsValid("Validating numStrDto02")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto02.IsValid('Validating numStrDto02')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto02NumStr, err := numStrDto02.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02NumStr, err := numStrDto02.GetNumStr()\n"+
			"numStrDto02 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr02 != numStrDto02NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto02 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr02 != numStrDto02NumStr\n"+
			"Expected numStrDto02NumStr = '%v'\n"+
			"  Actual numStrDto02NumStr = '%v'\n\n",
			ePrefix, inputNumStr02, numStrDto02NumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).SubtractNumStrs(numStrDto01, numStrDto02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  SubtractNumStrs(numStrDto01, numStrDto02)\n"+
			"numStrDto01= '%v'\n"+
			"numStrDto02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto01NumStr,
			numStrDto02NumStr,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
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

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_SubtractNumStrs_04(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_04"

	inputNumStr01 := "691.1"

	inputNumStr02 := "67.521"

	expectedNumStr := "623.579"

	expectedBigInt := big.NewInt(623579)

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

	expectedNumAbsIntStr := "623"

	expectedNumAbsFracStr := "579"

	expectedAbsAllRunesNumStr := "623579"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDto01, err := new(NumStrDto).ParseNumStr(inputNumStr01)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr01)\n"+
			"inputNumStr01= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr01, err.Error())
		return
	}

	err = numStrDto01.IsValid("Validating numStrDto01")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto01.IsValid('Validating numStrDto01')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto01NumStr, err := numStrDto01.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01NumStr, err := numStrDto01.GetNumStr()\n"+
			"numStrDto01 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr01 != numStrDto01NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto01 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr01 != numStrDto01NumStr\n"+
			"Expected numStrDto01NumStr = '%v'\n"+
			"  Actual numStrDto01NumStr = '%v'\n\n",
			ePrefix, inputNumStr01, numStrDto01NumStr)

		return
	}

	numStrDto02, err := new(NumStrDto).ParseNumStr(inputNumStr02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr02)\n"+
			"inputNumStr02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr02, err.Error())
		return
	}

	err = numStrDto02.IsValid("Validating numStrDto02")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto02.IsValid('Validating numStrDto02')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto02NumStr, err := numStrDto02.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02NumStr, err := numStrDto02.GetNumStr()\n"+
			"numStrDto02 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr02 != numStrDto02NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto02 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr02 != numStrDto02NumStr\n"+
			"Expected numStrDto02NumStr = '%v'\n"+
			"  Actual numStrDto02NumStr = '%v'\n\n",
			ePrefix, inputNumStr02, numStrDto02NumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).SubtractNumStrs(numStrDto01, numStrDto02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  SubtractNumStrs(numStrDto01, numStrDto02)\n"+
			"numStrDto01= '%v'\n"+
			"numStrDto02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto01NumStr,
			numStrDto02NumStr,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
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

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_SubtractNumStrs_05(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_05"

	inputNumStr01 := "691.1"

	inputNumStr02 := "0"

	expectedNumStr := "691.1"

	expectedBigInt := big.NewInt(6911)

	expectedPrecisionInt := 1

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = true

	expectedNumAbsIntStr := "691"

	expectedNumAbsFracStr := "1"

	expectedAbsAllRunesNumStr := "6911"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDto01, err := new(NumStrDto).ParseNumStr(inputNumStr01)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr01)\n"+
			"inputNumStr01= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr01, err.Error())
		return
	}

	err = numStrDto01.IsValid("Validating numStrDto01")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto01.IsValid('Validating numStrDto01')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto01NumStr, err := numStrDto01.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01NumStr, err := numStrDto01.GetNumStr()\n"+
			"numStrDto01 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr01 != numStrDto01NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto01 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr01 != numStrDto01NumStr\n"+
			"Expected numStrDto01NumStr = '%v'\n"+
			"  Actual numStrDto01NumStr = '%v'\n\n",
			ePrefix, inputNumStr01, numStrDto01NumStr)

		return
	}

	numStrDto02, err := new(NumStrDto).ParseNumStr(inputNumStr02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr02)\n"+
			"inputNumStr02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr02, err.Error())
		return
	}

	err = numStrDto02.IsValid("Validating numStrDto02")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto02.IsValid('Validating numStrDto02')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto02NumStr, err := numStrDto02.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02NumStr, err := numStrDto02.GetNumStr()\n"+
			"numStrDto02 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr02 != numStrDto02NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto02 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr02 != numStrDto02NumStr\n"+
			"Expected numStrDto02NumStr = '%v'\n"+
			"  Actual numStrDto02NumStr = '%v'\n\n",
			ePrefix, inputNumStr02, numStrDto02NumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).SubtractNumStrs(numStrDto01, numStrDto02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  SubtractNumStrs(numStrDto01, numStrDto02)\n"+
			"numStrDto01= '%v'\n"+
			"numStrDto02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto01NumStr,
			numStrDto02NumStr,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
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

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_SubtractNumStrs_06(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_06"

	inputNumStr01 := "0"

	inputNumStr02 := "691.1"

	expectedNumStr := "-691.1"

	expectedBigInt := big.NewInt(-6911)

	expectedPrecisionInt := 1

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = true

	expectedNumAbsIntStr := "691"

	expectedNumAbsFracStr := "1"

	expectedAbsAllRunesNumStr := "6911"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDto01, err := new(NumStrDto).ParseNumStr(inputNumStr01)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr01)\n"+
			"inputNumStr01= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr01, err.Error())
		return
	}

	err = numStrDto01.IsValid("Validating numStrDto01")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto01.IsValid('Validating numStrDto01')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto01NumStr, err := numStrDto01.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01NumStr, err := numStrDto01.GetNumStr()\n"+
			"numStrDto01 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr01 != numStrDto01NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto01 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr01 != numStrDto01NumStr\n"+
			"Expected numStrDto01NumStr = '%v'\n"+
			"  Actual numStrDto01NumStr = '%v'\n\n",
			ePrefix, inputNumStr01, numStrDto01NumStr)

		return
	}

	numStrDto02, err := new(NumStrDto).ParseNumStr(inputNumStr02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr02)\n"+
			"inputNumStr02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr02, err.Error())
		return
	}

	err = numStrDto02.IsValid("Validating numStrDto02")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto02.IsValid('Validating numStrDto02')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto02NumStr, err := numStrDto02.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02NumStr, err := numStrDto02.GetNumStr()\n"+
			"numStrDto02 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr02 != numStrDto02NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto02 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr02 != numStrDto02NumStr\n"+
			"Expected numStrDto02NumStr = '%v'\n"+
			"  Actual numStrDto02NumStr = '%v'\n\n",
			ePrefix, inputNumStr02, numStrDto02NumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).SubtractNumStrs(numStrDto01, numStrDto02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  SubtractNumStrs(numStrDto01, numStrDto02)\n"+
			"numStrDto01= '%v'\n"+
			"numStrDto02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto01NumStr,
			numStrDto02NumStr,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
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

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_SubtractNumStrs_07(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_07"

	inputNumStr01 := "-691.1"

	inputNumStr02 := "0"

	expectedNumStr := "-691.1"

	expectedBigInt := big.NewInt(-6911)

	expectedPrecisionInt := 1

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = true

	expectedNumAbsIntStr := "691"

	expectedNumAbsFracStr := "1"

	expectedAbsAllRunesNumStr := "6911"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDto01, err := new(NumStrDto).ParseNumStr(inputNumStr01)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr01)\n"+
			"inputNumStr01= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr01, err.Error())
		return
	}

	err = numStrDto01.IsValid("Validating numStrDto01")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto01.IsValid('Validating numStrDto01')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto01NumStr, err := numStrDto01.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01NumStr, err := numStrDto01.GetNumStr()\n"+
			"numStrDto01 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr01 != numStrDto01NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto01 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr01 != numStrDto01NumStr\n"+
			"Expected numStrDto01NumStr = '%v'\n"+
			"  Actual numStrDto01NumStr = '%v'\n\n",
			ePrefix, inputNumStr01, numStrDto01NumStr)

		return
	}

	numStrDto02, err := new(NumStrDto).ParseNumStr(inputNumStr02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr02)\n"+
			"inputNumStr02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr02, err.Error())
		return
	}

	err = numStrDto02.IsValid("Validating numStrDto02")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto02.IsValid('Validating numStrDto02')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto02NumStr, err := numStrDto02.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02NumStr, err := numStrDto02.GetNumStr()\n"+
			"numStrDto02 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr02 != numStrDto02NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto02 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr02 != numStrDto02NumStr\n"+
			"Expected numStrDto02NumStr = '%v'\n"+
			"  Actual numStrDto02NumStr = '%v'\n\n",
			ePrefix, inputNumStr02, numStrDto02NumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).SubtractNumStrs(numStrDto01, numStrDto02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  SubtractNumStrs(numStrDto01, numStrDto02)\n"+
			"numStrDto01= '%v'\n"+
			"numStrDto02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto01NumStr,
			numStrDto02NumStr,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
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

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrDto_SubtractNumStrs_08(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_08"

	inputNumStr01 := "0"

	inputNumStr02 := "0"

	expectedNumStr := "0"

	expectedBigInt := big.NewInt(0)

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

	expectedNumAbsIntStr := "0"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "0"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDto01, err := new(NumStrDto).ParseNumStr(inputNumStr01)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr01)\n"+
			"inputNumStr01= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr01, err.Error())
		return
	}

	err = numStrDto01.IsValid("Validating numStrDto01")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto01.IsValid('Validating numStrDto01')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto01NumStr, err := numStrDto01.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto01NumStr, err := numStrDto01.GetNumStr()\n"+
			"numStrDto01 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr01 != numStrDto01NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto01 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr01 != numStrDto01NumStr\n"+
			"Expected numStrDto01NumStr = '%v'\n"+
			"  Actual numStrDto01NumStr = '%v'\n\n",
			ePrefix, inputNumStr01, numStrDto01NumStr)

		return
	}

	numStrDto02, err := new(NumStrDto).ParseNumStr(inputNumStr02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02, err := new(NumStrDto).\n"+
			"  ParseNumStr(inputNumStr02)\n"+
			"inputNumStr02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr02, err.Error())
		return
	}

	err = numStrDto02.IsValid("Validating numStrDto02")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDto02.IsValid('Validating numStrDto02')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto02NumStr, err := numStrDto02.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto02NumStr, err := numStrDto02.GetNumStr()\n"+
			"numStrDto02 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumStr02 != numStrDto02NumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto02 Number Strings DON'T MATCH!\n"+
			"Because inputNumStr02 != numStrDto02NumStr\n"+
			"Expected numStrDto02NumStr = '%v'\n"+
			"  Actual numStrDto02NumStr = '%v'\n\n",
			ePrefix, inputNumStr02, numStrDto02NumStr)

		return
	}

	numStrDtoResult, err := new(NumStrDto).SubtractNumStrs(numStrDto01, numStrDto02)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  SubtractNumStrs(numStrDto01, numStrDto02)\n"+
			"numStrDto01= '%v'\n"+
			"numStrDto02= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto01NumStr,
			numStrDto02NumStr,
			err.Error())

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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
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

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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
