package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntNum_GetActualNumberOfDigits_01(t *testing.T) {

	ePrefix := "TestBigIntNum_GetActualNumberOfDigits_01"

	originalNumStr := "123.456"

	expectedDigits := 6

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOfDigitsInt64, isZeroVal, err := bINum.GetActualNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOfDigitsInt64, isZeroVal, err := \n"+
			"  bINum.GetActualNumberOfDigits()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	bINumOfDigitsInt := int(bINumOfDigitsInt64.Int64())

	if isZeroVal != false {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because isZeroVal != false\n"+
			"Expected isZeroVal = '%t'\n"+
			"  Actual isZeroVal = '%t'\n\n",
			ePrefix, false, true)

		return
	}

	if expectedDigits != bINumOfDigitsInt {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedDigits != bINumOfDigitsInt\n"+
			"Expected bINumOfDigitsInt = '%v'\n"+
			"  Actual bINumOfDigitsInt = '%v'\n\n",
			ePrefix, expectedDigits, bINumOfDigitsInt)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetActualNumberOfDigits_02(t *testing.T) {

	ePrefix := "TestBigIntNum_GetActualNumberOfDigits_02"

	originalNumStr := "0.000"

	expectedDigits := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumOfDigitsInt64, isZeroVal, err := bINum.GetActualNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumOfDigitsInt64, isZeroVal, err := \n"+
			"  bINum.GetActualNumberOfDigits()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	bINumOfDigitsInt := int(bINumOfDigitsInt64.Int64())

	if isZeroVal != true {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because isZeroVal != true\n"+
			"Expected isZeroVal = '%t'\n"+
			"  Actual isZeroVal = '%t'\n\n",
			ePrefix, true, false)

		return
	}

	if expectedDigits != bINumOfDigitsInt {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedDigits != bINumOfDigitsInt\n"+
			"Expected bINumOfDigitsInt = '%v'\n"+
			"  Actual bINumOfDigitsInt = '%v'\n\n",
			ePrefix, expectedDigits, bINumOfDigitsInt)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetBigIntFixedDecimal_01(t *testing.T) {

	ePrefix := "TestBigIntNum_GetBigIntFixedDecimal_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "123.456"

	expectedBigINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because originalNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, originalNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	fixedDec, err := expectedBigINum.GetBigIntFixedDecimal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := expectedBigINum.GetBigIntFixedDecimal()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	err = fixedDec.IsValid("Validating fixedDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating fixedDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fixedDecNumberStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumberStr, err := fixedDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fixedDecBigInt, err := fixedDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecBigInt, err := fixedDec.GetBigInt()\n"+
			"fixedDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fixedDecNumberStr, err.Error())
		return
	}

	fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecPrecisionUint, err := \n"+
			"  fixedDec.GetPrecisionUint()\n"+
			"fixedDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fixedDecNumberStr, err.Error())
		return
	}

	fixedDecNumSeps, err := fixedDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumSeps, err := fixedDec.GetNumericSeparatorsDto()\n"+
			"fixedDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fixedDecNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(fixedDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(fixedDecBigInt) != 0 \n"+
			"Expected fixedDecBigInt = '%v'\n"+
			"  Actual fixedDecBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), fixedDecBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != fixedDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != fixedDecPrecisionUint \n"+
			"Expected fixedDecPrecisionUint = '%v'\n"+
			"  Actual fixedDecPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, fixedDecPrecisionUint)

		return
	}

	if expectedBigINumberStr != fixedDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumberStr != fixedDecNumberStr \n"+
			"Expected fixedDecNumberStr = '%v'\n"+
			"  Actual fixedDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, fixedDecNumberStr)

		return
	}

	if !expectedNumSeps.Equal(fixedDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != fixedDecNumSeps \n"+
			"Expected fixedDecNumSeps = '%v'\n"+
			"  Actual fixedDecNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), fixedDecNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetFractionalPart_01(t *testing.T) {

	ePrefix := "TestBigIntNum_GetFractionalPart_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "123.456"

	expectedNumStr := "0.456"

	expectedPrecisionUint := uint(3)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because originalNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, originalNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint \n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			originalNumStr, err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	fractionalPartBINum, err := bINum.GetFractionalPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINum, err := bINum.GetFractionalPart()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = fractionalPartBINum.IsValid("Validating fractionalPartBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fractionalPartBINum.IsValid('Validating fractionalPartBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumPrecisionUint, err := fractionalPartBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumPrecisionUint, err := \n"+
			"   fractionalPartBINum.GetPrecisionUint()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	fractionalPartBINumSeps, err := fractionalPartBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumSeps, err :=\n"+
			"  fractionalPartBINum.GetNumericSeparatorsDto()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	expectedEqualsFractionalPart, err := expectedBigINum.Equal(fractionalPartBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedEqualsFractionalPart, err :=\n"+
			"  expectedBigINum.Equal(fractionalPartBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, fractionalPartBINumStr, err.Error())
		return
	}

	if !expectedEqualsFractionalPart {
		t.Errorf("%v\n"+
			"Error: Expected and 'fractionalPartBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsFractionalPart = 'false' \n"+
			"Expected fractionalPartBINum = '%v'\n"+
			"  Actual fractionalPartBINum = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedNumStr != fractionalPartBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and fractionalPartBINumStr number strings NOT EQUAL!\n"+
			"Because expectedNumStr != fractionalPartBINumStr\n"+
			"Expected fractionalPartBINumStr = '%v'\n"+
			"  Actual fractionalPartBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedPrecisionUint != fractionalPartBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected and Fractional Part Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != fractionalPartBINumPrecisionUint\n"+
			"Expected fractionalPartBINumPrecisionUint = '%v'\n"+
			"  Actual fractionalPartBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, fractionalPartBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(fractionalPartBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != fractionalPartBINumSeps \n"+
			"Expected fractionalPartBINumSeps = '%v'\n"+
			"  Actual fractionalPartBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), fractionalPartBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetFractionalPart_02(t *testing.T) {

	ePrefix := "TestBigIntNum_GetFractionalPart_02"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-123.456"

	expectedNumStr := "-0.456"

	expectedPrecisionUint := uint(3)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because originalNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, originalNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint \n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			originalNumStr, err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	fractionalPartBINum, err := bINum.GetFractionalPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINum, err := bINum.GetFractionalPart()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = fractionalPartBINum.IsValid("Validating fractionalPartBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fractionalPartBINum.IsValid('Validating fractionalPartBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumPrecisionUint, err := fractionalPartBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumPrecisionUint, err := \n"+
			"   fractionalPartBINum.GetPrecisionUint()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	fractionalPartBINumSeps, err := fractionalPartBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumSeps, err :=\n"+
			"  fractionalPartBINum.GetNumericSeparatorsDto()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	expectedEqualsFractionalPart, err := expectedBigINum.Equal(fractionalPartBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedEqualsFractionalPart, err :=\n"+
			"  expectedBigINum.Equal(fractionalPartBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, fractionalPartBINumStr, err.Error())
		return
	}

	if !expectedEqualsFractionalPart {
		t.Errorf("%v\n"+
			"Error: Expected and 'fractionalPartBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsFractionalPart = 'false' \n"+
			"Expected fractionalPartBINum = '%v'\n"+
			"  Actual fractionalPartBINum = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedNumStr != fractionalPartBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and fractionalPartBINumStr number strings NOT EQUAL!\n"+
			"Because expectedNumStr != fractionalPartBINumStr\n"+
			"Expected fractionalPartBINumStr = '%v'\n"+
			"  Actual fractionalPartBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedPrecisionUint != fractionalPartBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected and Fractional Part Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != fractionalPartBINumPrecisionUint\n"+
			"Expected fractionalPartBINumPrecisionUint = '%v'\n"+
			"  Actual fractionalPartBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, fractionalPartBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(fractionalPartBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != fractionalPartBINumSeps \n"+
			"Expected fractionalPartBINumSeps = '%v'\n"+
			"  Actual fractionalPartBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), fractionalPartBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetFractionalPart_03(t *testing.T) {

	ePrefix := "TestBigIntNum_GetFractionalPart_03"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "123"

	expectedNumStr := "0"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because originalNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, originalNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint \n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			originalNumStr, err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	fractionalPartBINum, err := bINum.GetFractionalPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINum, err := bINum.GetFractionalPart()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = fractionalPartBINum.IsValid("Validating fractionalPartBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fractionalPartBINum.IsValid('Validating fractionalPartBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumPrecisionUint, err := fractionalPartBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumPrecisionUint, err := \n"+
			"   fractionalPartBINum.GetPrecisionUint()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	fractionalPartBINumSeps, err := fractionalPartBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumSeps, err :=\n"+
			"  fractionalPartBINum.GetNumericSeparatorsDto()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	expectedEqualsFractionalPart, err := expectedBigINum.Equal(fractionalPartBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedEqualsFractionalPart, err :=\n"+
			"  expectedBigINum.Equal(fractionalPartBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, fractionalPartBINumStr, err.Error())
		return
	}

	if !expectedEqualsFractionalPart {
		t.Errorf("%v\n"+
			"Error: Expected and 'fractionalPartBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsFractionalPart = 'false' \n"+
			"Expected fractionalPartBINum = '%v'\n"+
			"  Actual fractionalPartBINum = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedNumStr != fractionalPartBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and fractionalPartBINumStr number strings NOT EQUAL!\n"+
			"Because expectedNumStr != fractionalPartBINumStr\n"+
			"Expected fractionalPartBINumStr = '%v'\n"+
			"  Actual fractionalPartBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedPrecisionUint != fractionalPartBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected and Fractional Part Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != fractionalPartBINumPrecisionUint\n"+
			"Expected fractionalPartBINumPrecisionUint = '%v'\n"+
			"  Actual fractionalPartBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, fractionalPartBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(fractionalPartBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != fractionalPartBINumSeps \n"+
			"Expected fractionalPartBINumSeps = '%v'\n"+
			"  Actual fractionalPartBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), fractionalPartBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetFractionalPart_04(t *testing.T) {

	ePrefix := "TestBigIntNum_GetFractionalPart_04"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-123"

	expectedNumStr := "0"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because originalNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, originalNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint \n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			originalNumStr, err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	fractionalPartBINum, err := bINum.GetFractionalPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINum, err := bINum.GetFractionalPart()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = fractionalPartBINum.IsValid("Validating fractionalPartBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fractionalPartBINum.IsValid('Validating fractionalPartBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumPrecisionUint, err := fractionalPartBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumPrecisionUint, err := \n"+
			"   fractionalPartBINum.GetPrecisionUint()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	fractionalPartBINumSeps, err := fractionalPartBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumSeps, err :=\n"+
			"  fractionalPartBINum.GetNumericSeparatorsDto()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	expectedEqualsFractionalPart, err := expectedBigINum.Equal(fractionalPartBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedEqualsFractionalPart, err :=\n"+
			"  expectedBigINum.Equal(fractionalPartBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, fractionalPartBINumStr, err.Error())
		return
	}

	if !expectedEqualsFractionalPart {
		t.Errorf("%v\n"+
			"Error: Expected and 'fractionalPartBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsFractionalPart = 'false' \n"+
			"Expected fractionalPartBINum = '%v'\n"+
			"  Actual fractionalPartBINum = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedNumStr != fractionalPartBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and fractionalPartBINumStr number strings NOT EQUAL!\n"+
			"Because expectedNumStr != fractionalPartBINumStr\n"+
			"Expected fractionalPartBINumStr = '%v'\n"+
			"  Actual fractionalPartBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedPrecisionUint != fractionalPartBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected and Fractional Part Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != fractionalPartBINumPrecisionUint\n"+
			"Expected fractionalPartBINumPrecisionUint = '%v'\n"+
			"  Actual fractionalPartBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, fractionalPartBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(fractionalPartBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != fractionalPartBINumSeps \n"+
			"Expected fractionalPartBINumSeps = '%v'\n"+
			"  Actual fractionalPartBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), fractionalPartBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetFractionalPart_05(t *testing.T) {

	ePrefix := "TestBigIntNum_GetFractionalPart_05"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "0.000"

	expectedNumStr := "0"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because originalNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, originalNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint \n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			originalNumStr, err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	fractionalPartBINum, err := bINum.GetFractionalPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINum, err := bINum.GetFractionalPart()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = fractionalPartBINum.IsValid("Validating fractionalPartBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fractionalPartBINum.IsValid('Validating fractionalPartBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumPrecisionUint, err := fractionalPartBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumPrecisionUint, err := \n"+
			"   fractionalPartBINum.GetPrecisionUint()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	fractionalPartBINumSeps, err := fractionalPartBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumSeps, err :=\n"+
			"  fractionalPartBINum.GetNumericSeparatorsDto()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	expectedEqualsFractionalPart, err := expectedBigINum.Equal(fractionalPartBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedEqualsFractionalPart, err :=\n"+
			"  expectedBigINum.Equal(fractionalPartBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, fractionalPartBINumStr, err.Error())
		return
	}

	if !expectedEqualsFractionalPart {
		t.Errorf("%v\n"+
			"Error: Expected and 'fractionalPartBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsFractionalPart = 'false' \n"+
			"Expected fractionalPartBINum = '%v'\n"+
			"  Actual fractionalPartBINum = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedNumStr != fractionalPartBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and fractionalPartBINumStr number strings NOT EQUAL!\n"+
			"Because expectedNumStr != fractionalPartBINumStr\n"+
			"Expected fractionalPartBINumStr = '%v'\n"+
			"  Actual fractionalPartBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedPrecisionUint != fractionalPartBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected and Fractional Part Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != fractionalPartBINumPrecisionUint\n"+
			"Expected fractionalPartBINumPrecisionUint = '%v'\n"+
			"  Actual fractionalPartBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, fractionalPartBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(fractionalPartBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != fractionalPartBINumSeps \n"+
			"Expected fractionalPartBINumSeps = '%v'\n"+
			"  Actual fractionalPartBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), fractionalPartBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetFractionalPart_06(t *testing.T) {

	ePrefix := "TestBigIntNum_GetFractionalPart_06"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "0"

	expectedNumStr := "0"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because originalNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, originalNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint \n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			originalNumStr, err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	fractionalPartBINum, err := bINum.GetFractionalPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINum, err := bINum.GetFractionalPart()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = fractionalPartBINum.IsValid("Validating fractionalPartBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fractionalPartBINum.IsValid('Validating fractionalPartBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumStr, err := fractionalPartBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fractionalPartBINumPrecisionUint, err := fractionalPartBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumPrecisionUint, err := \n"+
			"   fractionalPartBINum.GetPrecisionUint()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	fractionalPartBINumSeps, err := fractionalPartBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fractionalPartBINumSeps, err :=\n"+
			"  fractionalPartBINum.GetNumericSeparatorsDto()\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, fractionalPartBINumStr, err.Error())
		return
	}

	expectedEqualsFractionalPart, err := expectedBigINum.Equal(fractionalPartBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedEqualsFractionalPart, err :=\n"+
			"  expectedBigINum.Equal(fractionalPartBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"fractionalPartBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, fractionalPartBINumStr, err.Error())
		return
	}

	if !expectedEqualsFractionalPart {
		t.Errorf("%v\n"+
			"Error: Expected and 'fractionalPartBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsFractionalPart = 'false' \n"+
			"Expected fractionalPartBINum = '%v'\n"+
			"  Actual fractionalPartBINum = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedNumStr != fractionalPartBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and fractionalPartBINumStr number strings NOT EQUAL!\n"+
			"Because expectedNumStr != fractionalPartBINumStr\n"+
			"Expected fractionalPartBINumStr = '%v'\n"+
			"  Actual fractionalPartBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, fractionalPartBINumStr)

		return
	}

	if expectedPrecisionUint != fractionalPartBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected and Fractional Part Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != fractionalPartBINumPrecisionUint\n"+
			"Expected fractionalPartBINumPrecisionUint = '%v'\n"+
			"  Actual fractionalPartBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, fractionalPartBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(fractionalPartBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != fractionalPartBINumSeps \n"+
			"Expected fractionalPartBINumSeps = '%v'\n"+
			"  Actual fractionalPartBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), fractionalPartBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetDecimal_01(t *testing.T) {

	ePrefix := "TestBigIntNum_GetDecimal_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "-847921684.347"

	expectedPrecisionUint := uint(3)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint \n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err :=\n"+
			"  expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	decimalNum, err := expectedBigINum.GetDecimal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNum, err := expectedBigINum.GetDecimal()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	err = decimalNum.IsValid("Validating decimalNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decimalNum.IsValid('Validating decimalNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalNumNumberStr, err := decimalNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumNumberStr, err := decimalNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalNumPrecisionUint, err := decimalNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumPrecisionUint, err :=\n"+
			"  decimalNum.GetPrecisionUint()\n"+
			"decimalNum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decimalNumNumberStr, err.Error())
		return
	}

	decimalNumBigInt, err := decimalNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumBigInt, err := decimalNum.GetBigInt()\n"+
			"decimalNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalNumNumberStr, err.Error())
		return
	}

	decimalNumSeps, err := decimalNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumSeps, err := decimalNum.GetNumericSeparatorsDto()\n"+
			"decimalNum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decimalNumNumberStr, err.Error())
		return
	}

	if expectedNumStr != decimalNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String != Decimal Number String\n"+
			"Because expectedNumStr != decimalNumNumberStr\n"+
			"Expected decimalNumNumberStr = '%v'\n"+
			"  Actual decimalNumNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalNumNumberStr)

		return
	}

	if expectedPrecisionUint != decimalNumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecisionUint != decimalNumPrecisionUint\n"+
			"Expected decimalNumPrecisionUint = '%v'\n"+
			"  Actual decimalNumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalNumPrecisionUint)

		return
	}

	if expectedBigINumBigInt.Cmp(decimalNumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Absolute BigInt Values ARE NOT Equal!\n"+
			"Because expectedBigINumBigInt.Cmp(decimalNumBigInt) != 0 \n"+
			"Expected decimalNumBigInt = '%v'\n"+
			"  Actual decimalNumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), decimalNumBigInt.Text(10))

		return
	}

	if !expectedNumSeps.Equal(decimalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected decimalNumSeps = '%v'\n"+
			"  Actual decimalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalNumSeps.String())

		return
	}
}

func TestBigIntNum_GetNumStrDto_01(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumStrDto_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "-847921684.347"

	expectedPrecisionUint := uint(3)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	numStrDto, err := new(NumStrDto).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto, err :=new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumSeps, err.Error())
		return
	}

	err = numStrDto.IsValid("Validating numStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto.IsValid('Validating numStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoNumberStr, err := numStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumberStr, err := numStrDto.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoBigInt, err := numStrDto.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" numStrDtoBigInt, err :=\n"+
			"   numStrDto.GetBigInt()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoNumberStr, err.Error())
		return
	}

	numStrDtoPrecisionUint, err := numStrDto.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" numStrDtoPrecisionUint, err :=\n"+
			"   numStrDto.GetPrecisionUint()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoNumberStr, err.Error())
		return
	}

	numStrDtoNumSeps, err := numStrDto.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumSeps, err := numStrDto.GetNumericSeparatorsDto()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	if expectedPrecisionUint != numStrDtoPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != numStrDtoPrecisionUint \n"+
			"Expected numStrDtoPrecisionUint = '%v'\n"+
			"  Actual numStrDtoPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, numStrDtoPrecisionUint)

		return
	}

	if expectedNumStr != numStrDtoNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != numStrDtoNumberStr \n"+
			"Expected numStrDtoNumberStr = '%v'\n"+
			"  Actual numStrDtoNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoNumberStr)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != numStrDtoNumSeps \n"+
			"Expected numStrDtoNumSeps = '%v'\n"+
			"  Actual numStrDtoNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetIntAry_01(t *testing.T) {

	ePrefix := "TestBigIntNum_GetIntAry_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "-847921684.347"

	expectedPrecisionUint := uint(3)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	intAry, err := new(IntAry).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err :=new(IntAry).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumSeps, err.Error())
		return
	}

	err = intAry.IsValid("Validating intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryBigInt, err := intAry.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryBigInt, err := intAry.GetBigInt()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumSeps, err :=intAry.GetNumericSeparatorsDto()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(intAryBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(intAryBigInt) != 0\n"+
			"Expected intAryBigInt = '%v'\n"+
			"  Actual intAryBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), intAryBigInt.Text(10))

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint \n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedNumStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, intAryNumberStr)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetIntegerPart_01(t *testing.T) {

	ePrefix := "TestBigIntNum_GetIntAry_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "123.456"

	expectedNumStr := "123"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedBigINumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint \n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	integerPart, err := bINum.GetIntegerPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = integerPart.IsValid("Validating integerPart")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = integerPart.IsValid('Validating integerPart')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartNumberStr, err := integerPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartNumberStr, err := integerPart.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartBigInt, err := integerPart.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartBigInt, err := integerPart.GetBigInt()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartPrecisionUint, err := integerPart.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartPrecisionUint, err := \n"+
			"  integerPart.GetPrecisionUint()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartNumSeps, err := integerPart.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumSeps, err :=integerPart.GetNumericSeparatorsDto()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(integerPartBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(integerPartBigInt) != 0\n"+
			"Expected integerPartBigInt = '%v'\n"+
			"  Actual integerPartBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), integerPartBigInt.Text(10))

		return
	}

	if expectedPrecisionUint != integerPartPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != integerPartPrecisionUint \n"+
			"Expected integerPartPrecisionUint = '%v'\n"+
			"  Actual integerPartPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, integerPartPrecisionUint)

		return
	}

	if expectedNumStr != integerPartNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != integerPartNumberStr \n"+
			"Expected integerPartNumberStr = '%v'\n"+
			"  Actual integerPartNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, integerPartNumberStr)

		return
	}

	if !expectedNumSeps.Equal(integerPartNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != integerPartNumSeps \n"+
			"Expected integerPartNumSeps = '%v'\n"+
			"  Actual integerPartNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), integerPartNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetIntegerPart_02(t *testing.T) {

	ePrefix := "TestBigIntNum_GetIntegerPart_02"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-123.456"

	expectedNumStr := "-123"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedBigINumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint \n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	integerPart, err := bINum.GetIntegerPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = integerPart.IsValid("Validating integerPart")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = integerPart.IsValid('Validating integerPart')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartNumberStr, err := integerPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartNumberStr, err := integerPart.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartBigInt, err := integerPart.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartBigInt, err := integerPart.GetBigInt()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartPrecisionUint, err := integerPart.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartPrecisionUint, err := \n"+
			"  integerPart.GetPrecisionUint()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartNumSeps, err := integerPart.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumSeps, err :=integerPart.GetNumericSeparatorsDto()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(integerPartBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(integerPartBigInt) != 0\n"+
			"Expected integerPartBigInt = '%v'\n"+
			"  Actual integerPartBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), integerPartBigInt.Text(10))

		return
	}

	if expectedPrecisionUint != integerPartPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != integerPartPrecisionUint \n"+
			"Expected integerPartPrecisionUint = '%v'\n"+
			"  Actual integerPartPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, integerPartPrecisionUint)

		return
	}

	if expectedNumStr != integerPartNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != integerPartNumberStr \n"+
			"Expected integerPartNumberStr = '%v'\n"+
			"  Actual integerPartNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, integerPartNumberStr)

		return
	}

	if !expectedNumSeps.Equal(integerPartNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != integerPartNumSeps \n"+
			"Expected integerPartNumSeps = '%v'\n"+
			"  Actual integerPartNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), integerPartNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetIntegerPart_03(t *testing.T) {

	ePrefix := "TestBigIntNum_GetIntegerPart_03"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "123"

	expectedNumStr := "123"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedBigINumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint \n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	integerPart, err := bINum.GetIntegerPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = integerPart.IsValid("Validating integerPart")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = integerPart.IsValid('Validating integerPart')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartNumberStr, err := integerPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartNumberStr, err := integerPart.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartBigInt, err := integerPart.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartBigInt, err := integerPart.GetBigInt()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartPrecisionUint, err := integerPart.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartPrecisionUint, err := \n"+
			"  integerPart.GetPrecisionUint()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartNumSeps, err := integerPart.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumSeps, err :=integerPart.GetNumericSeparatorsDto()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(integerPartBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(integerPartBigInt) != 0\n"+
			"Expected integerPartBigInt = '%v'\n"+
			"  Actual integerPartBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), integerPartBigInt.Text(10))

		return
	}

	if expectedPrecisionUint != integerPartPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != integerPartPrecisionUint \n"+
			"Expected integerPartPrecisionUint = '%v'\n"+
			"  Actual integerPartPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, integerPartPrecisionUint)

		return
	}

	if expectedNumStr != integerPartNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != integerPartNumberStr \n"+
			"Expected integerPartNumberStr = '%v'\n"+
			"  Actual integerPartNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, integerPartNumberStr)

		return
	}

	if !expectedNumSeps.Equal(integerPartNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != integerPartNumSeps \n"+
			"Expected integerPartNumSeps = '%v'\n"+
			"  Actual integerPartNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), integerPartNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetIntegerPart_04(t *testing.T) {

	ePrefix := "TestBigIntNum_GetIntegerPart_04"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-123"

	expectedNumStr := "-123"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedBigINumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint \n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	integerPart, err := bINum.GetIntegerPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = integerPart.IsValid("Validating integerPart")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = integerPart.IsValid('Validating integerPart')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartNumberStr, err := integerPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartNumberStr, err := integerPart.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartBigInt, err := integerPart.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartBigInt, err := integerPart.GetBigInt()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartPrecisionUint, err := integerPart.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartPrecisionUint, err := \n"+
			"  integerPart.GetPrecisionUint()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartNumSeps, err := integerPart.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumSeps, err :=integerPart.GetNumericSeparatorsDto()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(integerPartBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(integerPartBigInt) != 0\n"+
			"Expected integerPartBigInt = '%v'\n"+
			"  Actual integerPartBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), integerPartBigInt.Text(10))

		return
	}

	if expectedPrecisionUint != integerPartPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != integerPartPrecisionUint \n"+
			"Expected integerPartPrecisionUint = '%v'\n"+
			"  Actual integerPartPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, integerPartPrecisionUint)

		return
	}

	if expectedNumStr != integerPartNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != integerPartNumberStr \n"+
			"Expected integerPartNumberStr = '%v'\n"+
			"  Actual integerPartNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, integerPartNumberStr)

		return
	}

	if !expectedNumSeps.Equal(integerPartNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != integerPartNumSeps \n"+
			"Expected integerPartNumSeps = '%v'\n"+
			"  Actual integerPartNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), integerPartNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetIntegerPart_05(t *testing.T) {

	ePrefix := "TestBigIntNum_GetIntegerPart_05"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "0.000"

	expectedNumStr := "0"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedBigINumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint \n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	integerPart, err := bINum.GetIntegerPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = integerPart.IsValid("Validating integerPart")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = integerPart.IsValid('Validating integerPart')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartNumberStr, err := integerPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartNumberStr, err := integerPart.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartBigInt, err := integerPart.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartBigInt, err := integerPart.GetBigInt()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartPrecisionUint, err := integerPart.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartPrecisionUint, err := \n"+
			"  integerPart.GetPrecisionUint()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartNumSeps, err := integerPart.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumSeps, err :=integerPart.GetNumericSeparatorsDto()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(integerPartBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(integerPartBigInt) != 0\n"+
			"Expected integerPartBigInt = '%v'\n"+
			"  Actual integerPartBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), integerPartBigInt.Text(10))

		return
	}

	if expectedPrecisionUint != integerPartPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != integerPartPrecisionUint \n"+
			"Expected integerPartPrecisionUint = '%v'\n"+
			"  Actual integerPartPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, integerPartPrecisionUint)

		return
	}

	if expectedNumStr != integerPartNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != integerPartNumberStr \n"+
			"Expected integerPartNumberStr = '%v'\n"+
			"  Actual integerPartNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, integerPartNumberStr)

		return
	}

	if !expectedNumSeps.Equal(integerPartNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != integerPartNumSeps \n"+
			"Expected integerPartNumSeps = '%v'\n"+
			"  Actual integerPartNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), integerPartNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetIntegerPart_06(t *testing.T) {

	ePrefix := "TestBigIntNum_GetIntegerPart_05"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "0"

	expectedNumStr := "0"

	expectedPrecisionUint := uint(0)

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedBigINumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint \n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	integerPart, err := bINum.GetIntegerPart()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = integerPart.IsValid("Validating integerPart")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = integerPart.IsValid('Validating integerPart')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartNumberStr, err := integerPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartNumberStr, err := integerPart.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	integerPartBigInt, err := integerPart.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartBigInt, err := integerPart.GetBigInt()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartPrecisionUint, err := integerPart.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"integerPartPrecisionUint, err := \n"+
			"  integerPart.GetPrecisionUint()\n"+
			"integerPart= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	integerPartNumSeps, err := integerPart.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumSeps, err :=integerPart.GetNumericSeparatorsDto()\n"+
			"numStrDto= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, integerPartNumberStr, err.Error())
		return
	}

	if expectedBigINumBigInt.Cmp(integerPartBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(integerPartBigInt) != 0\n"+
			"Expected integerPartBigInt = '%v'\n"+
			"  Actual integerPartBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), integerPartBigInt.Text(10))

		return
	}

	if expectedPrecisionUint != integerPartPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != integerPartPrecisionUint \n"+
			"Expected integerPartPrecisionUint = '%v'\n"+
			"  Actual integerPartPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, integerPartPrecisionUint)

		return
	}

	if expectedNumStr != integerPartNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != integerPartNumberStr \n"+
			"Expected integerPartNumberStr = '%v'\n"+
			"  Actual integerPartNumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, integerPartNumberStr)

		return
	}

	if !expectedNumSeps.Equal(integerPartNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != integerPartNumSeps \n"+
			"Expected integerPartNumSeps = '%v'\n"+
			"  Actual integerPartNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), integerPartNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_01(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "123.45"

	expectedDigitCnt := 5

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_02(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_02"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "1,234,567"

	expectedDigitCnt := 7

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_03(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_03"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "-1,234,567.8"

	expectedDigitCnt := 8

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_04(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_04"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "0"

	expectedDigitCnt := 1

	expectedBigINum, err := new(BigIntNum).NewZero(2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_05(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_05"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "0.00"

	expectedDigitCnt := 1

	expectedBigINum, err := new(BigIntNum).NewZero(2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewZero(2)\n"+
			"Error= '%v'\n\n",
			ePrefix, err.Error())
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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_06(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_06"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "012.34"

	expectedDigitCnt := 4

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_07(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_07"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "0.1234"

	expectedDigitCnt := 4

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_08(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_08"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "0.123400"

	expectedDigitCnt := 6

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_09(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_09"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "0.0123400"

	expectedDigitCnt := 6

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_10(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_10"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "1,234,567.800"

	expectedDigitCnt := 10

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_GetNumberOfDigits_11(t *testing.T) {

	ePrefix := "TestBigIntNum_GetNumberOfDigits_11"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStr := "5"

	expectedDigitCnt := 1

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualDigitCnt, err := expectedBigINum.GetNumberOfDigits()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedDigitCnt != actualDigitCnt {
		t.Errorf("%v\n"+
			"Error: Digit Counts ARE NOT EQUAL!\n"+
			"Because expectedDigitCnt != actualDigitCnt\n"+
			"Expected actualDigitCnt = '%v'\n"+
			"  Actual actualDigitCnt = '%v'\n\n",
			ePrefix, expectedDigitCnt, actualDigitCnt)

		return
	}

	return
}

func TestBigIntNum_Increment_01(t *testing.T) {

	ePrefix := "TestBigIntNum_Increment_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "8"

	expectedNumStr := "9"

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	err = bINum.Increment()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.Increment()\n"+
			"Original bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Increment")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Increment')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumIncrementNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumIncrementNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementNumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINumIncrementNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	bINumIncrementBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	bINumIncrementPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	if expectedNumStr != bINumIncrementNumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	if expectedBigINumBigInt.Cmp(bINumIncrementBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumIncrementBigInt) != 0\n"+
			"Expected bINumIncrementBigInt = '%v'\n"+
			"  Actual bINumIncrementBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumIncrementBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != bINumIncrementPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumIncrementPrecisionUint \n"+
			"Expected bINumIncrementPrecisionUint = '%v'\n"+
			"  Actual bINumIncrementPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumIncrementPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(bINumIncrementNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumIncrementNumSeps \n"+
			"Expected bINumIncrementNumSeps = '%v'\n"+
			"  Actual bINumIncrementNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumIncrementNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Increment_02(t *testing.T) {

	ePrefix := "TestBigIntNum_Increment_02"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "8.2"

	expectedNumStr := "9.2"

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	err = bINum.Increment()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.Increment()\n"+
			"Original bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Increment")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Increment')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumIncrementNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumIncrementNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementNumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINumIncrementNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	bINumIncrementBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	bINumIncrementPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	if expectedNumStr != bINumIncrementNumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	if expectedBigINumBigInt.Cmp(bINumIncrementBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumIncrementBigInt) != 0\n"+
			"Expected bINumIncrementBigInt = '%v'\n"+
			"  Actual bINumIncrementBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumIncrementBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != bINumIncrementPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumIncrementPrecisionUint \n"+
			"Expected bINumIncrementPrecisionUint = '%v'\n"+
			"  Actual bINumIncrementPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumIncrementPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(bINumIncrementNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumIncrementNumSeps \n"+
			"Expected bINumIncrementNumSeps = '%v'\n"+
			"  Actual bINumIncrementNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumIncrementNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Increment_03(t *testing.T) {

	ePrefix := "TestBigIntNum_Increment_03"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-8.2"

	expectedNumStr := "-7.2"

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	err = bINum.Increment()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.Increment()\n"+
			"Original bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Increment")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Increment')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumIncrementNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumIncrementNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementNumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINumIncrementNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	bINumIncrementBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	bINumIncrementPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	if expectedNumStr != bINumIncrementNumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	if expectedBigINumBigInt.Cmp(bINumIncrementBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumIncrementBigInt) != 0\n"+
			"Expected bINumIncrementBigInt = '%v'\n"+
			"  Actual bINumIncrementBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumIncrementBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != bINumIncrementPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumIncrementPrecisionUint \n"+
			"Expected bINumIncrementPrecisionUint = '%v'\n"+
			"  Actual bINumIncrementPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumIncrementPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(bINumIncrementNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bINumIncrementNumSeps \n"+
			"Expected bINumIncrementNumSeps = '%v'\n"+
			"  Actual bINumIncrementNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumIncrementNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Increment_04(t *testing.T) {

	ePrefix := "TestBigIntNum_Increment_03"

	numSepsPair := NumericSeparatorPairDto{}

	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	numSepsPair.OutputSeparators.DecimalSeparator = frenchDecSeparator

	numSepsPair.OutputSeparators.ThousandsSeparator = frenchThousandsSeparator

	numSepsPair.OutputSeparators.CurrencySymbol = frenchCurrencySymbol

	err := numSepsPair.OutputSeparators.IsValid("Validating numSepsPair.OutputSeparators")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numSepsPair.OutputSeparators.\n"+
			"  IsValid('Validating  numSepsPair.OutputSeparators')\n"+
			"numSepsPair.OutputSeparators= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, numSepsPair.OutputSeparators.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	numSepsPair.InputSeparators.DecimalSeparator = usaNumSeps.DecimalSeparator

	numSepsPair.InputSeparators.ThousandsSeparator = usaNumSeps.ThousandsSeparator

	numSepsPair.InputSeparators.CurrencySymbol = usaNumSeps.CurrencySymbol

	err = numSepsPair.InputSeparators.IsValid("Validating numSepsPair.InputSeparators")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numSepsPair.InputSeparators.\n"+
			"  IsValid('Validating  numSepsPair.InputSeparators')\n"+
			"numSepsPair.InputSeparators= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, numSepsPair.InputSeparators.String(), err.Error())
		return
	}

	originalNumStr := "8"

	expectedNumStr := "9"

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &numSepsPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &numSepsPair)\n"+
			"expectedNumStr= '%v'\n"+
			"numSepsPair.InputSeparators= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			numSepsPair.InputSeparators.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !numSepsPair.OutputSeparators.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because numSepsPair.OutputSeparators != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, numSepsPair.OutputSeparators.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &numSepsPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &numSepsPair)\n"+
			"originalNumStr= '%v'\n"+
			"numSepsPair.InputSeparators= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			numSepsPair.InputSeparators.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
		return
	}

	if !numSepsPair.OutputSeparators.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because numSepsPair.OutputSeparators != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, numSepsPair.OutputSeparators.String(), bINumSeps.String())

		return
	}

	err = bINum.Increment()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.Increment()\n"+
			"Original bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Increment")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Increment')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumIncrementNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementNumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumIncrementNumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementNumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINumIncrementNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	bINumIncrementBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	bINumIncrementPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementPrecisionUint, err := \n"+
			"  bINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bINumIncrementNumStr, err.Error())
		return
	}

	if expectedNumStr != bINumIncrementNumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	if expectedBigINumBigInt.Cmp(bINumIncrementBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(bINumIncrementBigInt) != 0\n"+
			"Expected bINumIncrementBigInt = '%v'\n"+
			"  Actual bINumIncrementBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), bINumIncrementBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != bINumIncrementPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumIncrementPrecisionUint \n"+
			"Expected bINumIncrementPrecisionUint = '%v'\n"+
			"  Actual bINumIncrementPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumIncrementPrecisionUint)

		return
	}

	if !numSepsPair.OutputSeparators.Equal(bINumIncrementNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because  numSepsPair.OutputSeparators != bINumIncrementNumSeps \n"+
			"Expected bINumIncrementNumSeps = '%v'\n"+
			"  Actual bINumIncrementNumSeps = '%v'\n\n",
			ePrefix, numSepsPair.OutputSeparators.String(), bINumIncrementNumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Inverse_01(t *testing.T) {

	ePrefix := "TestBigIntNum_Inverse_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "4"

	expectedNumStr := "0.25"

	expectedPrecisionUint := uint(2)

	maxPrecisionUint := expectedPrecisionUint

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	inverseBINum, err := bINum.Inverse(maxPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINum, err :=\n"+
			"  bINum.Inverse(maxPrecisionUint)\n"+
			"Original bINum= '%v'\n"+
			"maxPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, maxPrecisionUint, err.Error())
		return
	}

	err = inverseBINum.IsValid("Validating inverseBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = inverseBINum.IsValid('Validating inverseBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumStr, err := inverseBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumStr, err := actualInverseBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumBigInt, err := inverseBINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumPrecisionUint, err := inverseBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumPrecisionUint, err :=\n"+
			"  inverseBINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	if expectedNumStr != inverseBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values ARE NOT Equal\n"+
			"Because expectedNumStr != inverseBINumStr \n"+
			"Expected inverseBINumStr = '%v'\n"+
			"  Actual inverseBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, inverseBINumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0\n"+
			"Expected inverseBINumBigInt = '%v'\n"+
			"  Actual inverseBINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), inverseBINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != inverseBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != inverseBINumPrecisionUint\n"+
			"Expected inverseBINumPrecisionUint = '%v'\n"+
			"  Actual inverseBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, inverseBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(inverseBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != inverseBINumSeps\n"+
			"Expected inverseBINumSeps = '%v'\n"+
			"  Actual inverseBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), inverseBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Inverse_02(t *testing.T) {

	ePrefix := "TestBigIntNum_Inverse_02"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "9.357"

	expectedNumStr := "0.10687186063909372662178048519825"

	expectedPrecisionUint := uint(32)

	maxPrecisionUint := expectedPrecisionUint

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	inverseBINum, err := bINum.Inverse(maxPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINum, err :=\n"+
			"  bINum.Inverse(maxPrecisionUint)\n"+
			"Original bINum= '%v'\n"+
			"maxPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, maxPrecisionUint, err.Error())
		return
	}

	err = inverseBINum.IsValid("Validating inverseBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = inverseBINum.IsValid('Validating inverseBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumStr, err := inverseBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumStr, err := actualInverseBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumBigInt, err := inverseBINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumPrecisionUint, err := inverseBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumPrecisionUint, err :=\n"+
			"  inverseBINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	if expectedNumStr != inverseBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values ARE NOT Equal\n"+
			"Because expectedNumStr != inverseBINumStr \n"+
			"Expected inverseBINumStr = '%v'\n"+
			"  Actual inverseBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, inverseBINumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0\n"+
			"Expected inverseBINumBigInt = '%v'\n"+
			"  Actual inverseBINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), inverseBINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != inverseBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != inverseBINumPrecisionUint\n"+
			"Expected inverseBINumPrecisionUint = '%v'\n"+
			"  Actual inverseBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, inverseBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(inverseBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != inverseBINumSeps\n"+
			"Expected inverseBINumSeps = '%v'\n"+
			"  Actual inverseBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), inverseBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Inverse_03(t *testing.T) {

	ePrefix := "TestBigIntNum_Inverse_03"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-9.357"

	expectedNumStr := "-0.10687186063909372662178048519825"

	expectedPrecisionUint := uint(32)

	maxPrecisionUint := expectedPrecisionUint

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	inverseBINum, err := bINum.Inverse(maxPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINum, err :=\n"+
			"  bINum.Inverse(maxPrecisionUint)\n"+
			"Original bINum= '%v'\n"+
			"maxPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, maxPrecisionUint, err.Error())
		return
	}

	err = inverseBINum.IsValid("Validating inverseBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = inverseBINum.IsValid('Validating inverseBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumStr, err := inverseBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumStr, err := actualInverseBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumBigInt, err := inverseBINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumPrecisionUint, err := inverseBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumPrecisionUint, err :=\n"+
			"  inverseBINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	if expectedNumStr != inverseBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values ARE NOT Equal\n"+
			"Because expectedNumStr != inverseBINumStr \n"+
			"Expected inverseBINumStr = '%v'\n"+
			"  Actual inverseBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, inverseBINumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0\n"+
			"Expected inverseBINumBigInt = '%v'\n"+
			"  Actual inverseBINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), inverseBINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != inverseBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != inverseBINumPrecisionUint\n"+
			"Expected inverseBINumPrecisionUint = '%v'\n"+
			"  Actual inverseBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, inverseBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(inverseBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != inverseBINumSeps\n"+
			"Expected inverseBINumSeps = '%v'\n"+
			"  Actual inverseBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), inverseBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Inverse_04(t *testing.T) {

	ePrefix := "TestBigIntNum_Inverse_04"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "0"

	expectedNumStr := "0"

	expectedPrecisionUint := uint(0)

	maxPrecisionUint := expectedPrecisionUint

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	inverseBINum, err := bINum.Inverse(maxPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINum, err :=\n"+
			"  bINum.Inverse(maxPrecisionUint)\n"+
			"Original bINum= '%v'\n"+
			"maxPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, maxPrecisionUint, err.Error())
		return
	}

	err = inverseBINum.IsValid("Validating inverseBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = inverseBINum.IsValid('Validating inverseBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumStr, err := inverseBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumStr, err := actualInverseBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumBigInt, err := inverseBINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumPrecisionUint, err := inverseBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumPrecisionUint, err :=\n"+
			"  inverseBINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	if expectedNumStr != inverseBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values ARE NOT Equal\n"+
			"Because expectedNumStr != inverseBINumStr \n"+
			"Expected inverseBINumStr = '%v'\n"+
			"  Actual inverseBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, inverseBINumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0\n"+
			"Expected inverseBINumBigInt = '%v'\n"+
			"  Actual inverseBINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), inverseBINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != inverseBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != inverseBINumPrecisionUint\n"+
			"Expected inverseBINumPrecisionUint = '%v'\n"+
			"  Actual inverseBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, inverseBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(inverseBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != inverseBINumSeps\n"+
			"Expected inverseBINumSeps = '%v'\n"+
			"  Actual inverseBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), inverseBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Inverse_09(t *testing.T) {

	ePrefix := "TestBigIntNum_Inverse_09"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "9"

	expectedNumStr := "0.11111111111111111111111111111111"

	expectedPrecisionUint := uint(32)

	maxPrecisionUint := expectedPrecisionUint

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	inverseBINum, err := bINum.Inverse(maxPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINum, err :=\n"+
			"  bINum.Inverse(maxPrecisionUint)\n"+
			"Original bINum= '%v'\n"+
			"maxPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, maxPrecisionUint, err.Error())
		return
	}

	err = inverseBINum.IsValid("Validating inverseBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = inverseBINum.IsValid('Validating inverseBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumStr, err := inverseBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumStr, err := actualInverseBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumBigInt, err := inverseBINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumPrecisionUint, err := inverseBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumPrecisionUint, err :=\n"+
			"  inverseBINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	if expectedNumStr != inverseBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values ARE NOT Equal\n"+
			"Because expectedNumStr != inverseBINumStr \n"+
			"Expected inverseBINumStr = '%v'\n"+
			"  Actual inverseBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, inverseBINumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0\n"+
			"Expected inverseBINumBigInt = '%v'\n"+
			"  Actual inverseBINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), inverseBINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != inverseBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != inverseBINumPrecisionUint\n"+
			"Expected inverseBINumPrecisionUint = '%v'\n"+
			"  Actual inverseBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, inverseBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(inverseBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != inverseBINumSeps\n"+
			"Expected inverseBINumSeps = '%v'\n"+
			"  Actual inverseBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), inverseBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Inverse_10(t *testing.T) {

	ePrefix := "TestBigIntNum_Inverse_10"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-9"

	expectedNumStr := "-0.11111111111111111111111111111111"

	expectedPrecisionUint := uint(32)

	maxPrecisionUint := expectedPrecisionUint

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumStr,
			expectedNumSeps.String(),
			err.Error())

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

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err := \n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumNumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	inverseBINum, err := bINum.Inverse(maxPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINum, err :=\n"+
			"  bINum.Inverse(maxPrecisionUint)\n"+
			"Original bINum= '%v'\n"+
			"maxPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, maxPrecisionUint, err.Error())
		return
	}

	err = inverseBINum.IsValid("Validating inverseBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = inverseBINum.IsValid('Validating inverseBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumStr, err := inverseBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumStr, err := actualInverseBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumSeps, err := inverseBINum.GetNumericSeparatorsDto()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumBigInt, err := inverseBINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumIncrementBigInt, err := bINum.GetBigInt()\n"+
			"inverseBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	inverseBINumPrecisionUint, err := inverseBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseBINumPrecisionUint, err :=\n"+
			"  inverseBINum.GetPrecisionUint()\n"+
			"bINumIncrement= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, inverseBINumStr, err.Error())
		return
	}

	if expectedNumStr != inverseBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected bI Number String Values ARE NOT Equal\n"+
			"Because expectedNumStr != inverseBINumStr \n"+
			"Expected inverseBINumStr = '%v'\n"+
			"  Actual inverseBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, inverseBINumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumBigInt.Cmp(inverseBINumBigInt) != 0\n"+
			"Expected inverseBINumBigInt = '%v'\n"+
			"  Actual inverseBINumBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), inverseBINumBigInt.Text(10))

		return
	}

	if expectedBigINumPrecisionUint != inverseBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precison Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != inverseBINumPrecisionUint\n"+
			"Expected inverseBINumPrecisionUint = '%v'\n"+
			"  Actual inverseBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, inverseBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(inverseBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != inverseBINumSeps\n"+
			"Expected inverseBINumSeps = '%v'\n"+
			"  Actual inverseBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), inverseBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_01(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_01"

	originalNumStr := "4"

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because true != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, true, isEven)

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_02(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_02"

	originalNumStr := "5"
	// isEven = false

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because false != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, false, isEven)

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_03(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_03"

	originalNumStr := "4.4"
	// isEven = false

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because false != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, false, isEven)

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_04(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_04"

	originalNumStr := "9793442794"
	// isEven = true

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because true != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, true, isEven)

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_05(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_05"

	originalNumStr := "9793442795"
	// isEven = false

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because false != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, false, isEven)

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_06(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_06"

	originalNumStr := "-9793442794"
	// isEven = true

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because true != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, true, isEven)

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_07(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_07"

	originalNumStr := "757849035736836546"
	// isEven = true

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because true != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, true, isEven)

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_08(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_08"

	originalNumStr := "-757849035736836546"
	// isEven = true

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because true != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, true, isEven)

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_09(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_09"

	originalNumStr := "757849035736836547"
	// isEven = false

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because false != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, false, isEven)

		return
	}

	return
}

func TestBigIntNum_IsEvenNumber_10(t *testing.T) {

	ePrefix := "TestBigIntNum_IsEvenNumber_10"

	originalNumStr := "-757849035736836547"
	// isEven = false

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	isEven, err := bINum.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"isEven, err := bINum.IsEvenNumber()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumberStr, err.Error())
		return
	}

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Even Number Classification FAIILED!\n"+
			"Because false != isEven \n"+
			"Expected isEven = '%v'\n"+
			"  Actual isEven = '%v'\n\n",
			ePrefix, false, isEven)

		return
	}

	return
}

func TestBigIntNum_IsZero_01(t *testing.T) {

	testNumStr := "0.000"
	bINum, err := new(BigIntNum).NewNumStr(testNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
			"testNumStr='%v' Error='%v' ", testNumStr, err.Error())
	}

	isZero := bINum.IsZero()

	if true != isZero {
		t.Errorf("Expected TestNumber to be IsZero='%v'. Instead TestNumber IsZero='%v'",
			true, isZero)
	}
}

func TestBigIntNum_IsZero_02(t *testing.T) {

	testNumStr := "0"
	bINum, err := new(BigIntNum).NewNumStr(testNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
			"testNumStr='%v' Error='%v' ", testNumStr, err.Error())
	}

	isZero := bINum.IsZero()

	if true != isZero {
		t.Errorf("Expected TestNumber to be IsZero='%v'. Instead TestNumber IsZero='%v'",
			true, isZero)
	}
}

func TestBigIntNum_IsZero_03(t *testing.T) {

	testNumStr := "0.0000000000001"
	bINum, err := new(BigIntNum).NewNumStr(testNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
			"testNumStr='%v' Error='%v' ", testNumStr, err.Error())
	}

	isZero := bINum.IsZero()

	if false != isZero {
		t.Errorf("Expected TestNumber to be IsZero='%v'. Instead TestNumber IsZero='%v'",
			false, isZero)
	}
}

func TestBigIntNum_IsZero_04(t *testing.T) {

	testNumStr := "100000000000000000000"
	bINum, err := new(BigIntNum).NewNumStr(testNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
			"testNumStr='%v' Error='%v' ", testNumStr, err.Error())
	}

	isZero := bINum.IsZero()

	if false != isZero {
		t.Errorf("Expected TestNumber to be IsZero='%v'. Instead TestNumber IsZero='%v'",
			false, isZero)
	}
}

func TestBigIntNum_IsZero_05(t *testing.T) {

	testNumStr := "0.000000000000000000000000001"
	bINum, err := new(BigIntNum).NewNumStr(testNumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
			"testNumStr='%v' Error='%v' ", testNumStr, err.Error())
	}

	isZero := bINum.IsZero()

	if false != isZero {
		t.Errorf("Expected TestNumber to be IsZero='%v'. Instead TestNumber IsZero='%v'",
			false, isZero)
	}
}

func TestBigIntNum_IntAry_01(t *testing.T) {

	nStr := "123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	ia, err := IntAry{}.NewNumStr(nStr)

	bINum, err := new(BigIntNum).NewIntAry(ia)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewIntAry(ia) "+
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) "+
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.bigInt) != 0 {
		t.Errorf("Expected bigInt='%v'  Instead, bigInt='%v'. ",
			bOriginal.Text(10), bINum.bigInt.Text(10))
	}

	if expectedPrecision != bINum.precision {
		t.Errorf("Expected precision='%v' Instead, precision='%v' ",
			expectedPrecision, bINum.precision)
	}

	if bINum.scaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.scaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.absBigInt) != 0 {
		t.Errorf("Expected absBigInt='%v'  Instead, absBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.absBigInt.Text(10))
	}

	if expectedSignVal != bINum.sign {
		t.Errorf("Expected sign Value='%v'. Instead, sign Value='%v'. ",
			expectedSignVal, bINum.sign)
	}

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_IntAry_02(t *testing.T) {

	nStr := "-123.456"
	expectedPrecision := uint(3)
	nbStr := "-123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := -1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	ia, err := IntAry{}.NewNumStr(nStr)

	bINum, err := new(BigIntNum).NewIntAry(ia)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewIntAry(ia) "+
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) "+
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.bigInt) != 0 {
		t.Errorf("Expected bigInt='%v'  Instead, bigInt='%v'. ",
			bOriginal.Text(10), bINum.bigInt.Text(10))
	}

	if expectedPrecision != bINum.precision {
		t.Errorf("Expected precision='%v' Instead, precision='%v' ",
			expectedPrecision, bINum.precision)
	}

	if bINum.scaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.scaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.absBigInt) != 0 {
		t.Errorf("Expected absBigInt='%v'  Instead, absBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.absBigInt.Text(10))
	}

	if expectedSignVal != bINum.sign {
		t.Errorf("Expected sign Value='%v'. Instead, sign Value='%v'. ",
			expectedSignVal, bINum.sign)
	}

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}
