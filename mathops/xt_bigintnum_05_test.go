package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntNum_NewNumStr_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_01"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumberStr := "123.456"

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignValue := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(expectedNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumberStr)\n"+
			"expectedNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr, err.Error())
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignValue != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignValue != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_02"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumberStr := "-123.456"

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "-123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := -1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(expectedNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumberStr)\n"+
			"expectedNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr, err.Error())
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_03"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumberStr := "0.123456789012"

	expectedPrecisionUint := uint(12)

	expectedBigIntNumStr := "123456789012"

	expectedAbsBigIntNumStr := "123456789012"

	expectedScaleFactor := big.NewInt(1000000000000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(expectedNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(expectedNumberStr)\n"+
			"expectedNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr, err.Error())
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_04(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_04"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := ".123456789012"

	expectedNumberStr := "0.123456789012"

	expectedPrecisionUint := uint(12)

	expectedBigIntNumStr := "123456789012"

	expectedAbsBigIntNumStr := "123456789012"

	expectedScaleFactor := big.NewInt(1000000000000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_05(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_05"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-.123456789012"

	expectedNumberStr := "-0.123456789012"

	expectedPrecisionUint := uint(12)

	expectedBigIntNumStr := "-123456789012"

	expectedAbsBigIntNumStr := "123456789012"

	expectedScaleFactor := big.NewInt(1000000000000)

	expectedSignVal := -1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_06(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_06"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "10"

	expectedNumberStr := "10"

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "10"

	expectedAbsBigIntNumStr := "10"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_07(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_07"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-52"

	expectedNumberStr := "-52"

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "-52"

	expectedAbsBigIntNumStr := "52"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := -1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_08(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_08"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-00052.1234"

	expectedNumberStr := "-52.1234"

	expectedPrecisionUint := uint(4)

	expectedScaleFactor := big.NewInt(10000)

	expectedSignVal := -1

	expectedBigIntNumStr := "-521234"

	expectedAbsBigIntNumStr := "521234"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_09(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_09"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "(00052.1234)"

	expectedNumberStr := "-52.1234"

	expectedPrecisionUint := uint(4)

	expectedScaleFactor := big.NewInt(10000)

	expectedSignVal := -1

	expectedBigIntNumStr := "-521234"

	expectedAbsBigIntNumStr := "521234"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_10(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_10"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "(00052.1234"

	expectedNumberStr := "52.1234"

	expectedPrecisionUint := uint(4)

	expectedScaleFactor := big.NewInt(10000)

	expectedSignVal := 1

	expectedBigIntNumStr := "521234"

	expectedAbsBigIntNumStr := "521234"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_11(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_11"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "+00052.1234"

	expectedNumberStr := "52.1234"

	expectedPrecisionUint := uint(4)

	expectedScaleFactor := big.NewInt(10000)

	expectedSignVal := 1

	expectedBigIntNumStr := "521234"

	expectedAbsBigIntNumStr := "521234"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_12(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_12"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "-00052.1234567890123456"

	expectedNumberStr := "-52.1234567890123456"

	expectedPrecisionUint := uint(16)

	expectedScaleFactor := big.NewInt(10000000000000000)

	expectedSignVal := -1

	expectedBigIntNumStr := "-521234567890123456"

	expectedAbsBigIntNumStr := "521234567890123456"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_13(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_13"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "52 . 123 4567 8901 23456"

	expectedNumberStr := "52.1234567890123456"

	expectedPrecisionUint := uint(16)

	expectedScaleFactor := big.NewInt(10000000000000000)

	expectedSignVal := 1

	expectedBigIntNumStr := "521234567890123456"

	expectedAbsBigIntNumStr := "521234567890123456"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_14(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_14"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "5 2"

	expectedNumberStr := "52"

	expectedPrecisionUint := uint(0)

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNumStr := "52"

	expectedAbsBigIntNumStr := "52"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_15(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_15"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "    (52)     "

	expectedNumberStr := "-52"

	expectedPrecisionUint := uint(0)

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := -1

	expectedBigIntNumStr := "-52"

	expectedAbsBigIntNumStr := "52"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_16(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_16"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	originalNumStr := "    (52)    1234567 "

	expectedNumberStr := "-52"

	expectedPrecisionUint := uint(0)

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := -1

	expectedBigIntNumStr := "-52"

	expectedAbsBigIntNumStr := "52"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStr_17(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_17"

	expectedPrecisionUint := uint(1024)

	bigIFxDecNum, err := new(BigIntFixedDecimal).NewNumStr(EulersNum50kStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIFxDecNum, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStr(EulersNum50kStr, '.')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bigIFxDecNum.RoundToDecPlace(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bigIFxDecNum.RoundToDecPlace(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedPrecisionUint, err.Error())
		return
	}

	bigIFxDecNumberStr, err := bigIFxDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIFxDecNumberStr, err := bigIFxDecNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIFxDecNumPrecisionUint, err := bigIFxDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIFxDecNumPrecisionUint, err := bigIFxDecNum.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fdEulers1kFxDecNum := GetEulersNum1k()

	fdEulers1kNumStr, err := fdEulers1kFxDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fdEulers1kNumStr, err := fdEulers1kFxDecNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	fdEulers1kFxDecNumPrecisionUint, err := fdEulers1kFxDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fdEulers1kFxDecNumPrecisionUint, err :=\n"+
			"  fdEulers1kFxDecNum.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedPrecisionUint != fdEulers1kFxDecNumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision for fdEulers1kFxDecNum is WRONG!\n"+
			"Because expectedPrecisionUint != fdEulers1kFxDecNumPrecisionUint\n"+
			"Expected fdEulers1kFxDecNumPrecisionUint = '%v'\n"+
			"  Actual fdEulers1kFxDecNumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, fdEulers1kFxDecNumPrecisionUint)

		return
	}

	if expectedPrecisionUint != bigIFxDecNumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: bigIFxDecNum is WRONG!\n"+
			"Because expectedPrecisionUint != bigIFxDecNumPrecisionUint\n"+
			"Expected bigIFxDecNumPrecisionUint = '%v'\n"+
			"  Actual bigIFxDecNumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bigIFxDecNumPrecisionUint)

		return
	}

	fdEulers1kNumStrLength := len(fdEulers1kNumStr)

	bigIFxDecNumberStrLen := len(bigIFxDecNumberStr)

	if fdEulers1kNumStrLength != bigIFxDecNumberStrLen {
		t.Errorf("%v\n"+
			"Error: Expected String Lengths DO NOT MATCH!\n"+
			"Because fdEulers1kNumStrLength != bigIFxDecNumberStrLen\n"+
			"Expected bigIFxDecNumberStrLen = '%v'\n"+
			"  Actual bigIFxDecNumberStrLen = '%v'\n\n",
			ePrefix, fdEulers1kNumStrLength, bigIFxDecNumberStrLen)

		return
	}

	if fdEulers1kNumStr != bigIFxDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"HOWEVER, Precision Values DO MATCH!\n"+
			"Because fdEulers1kNumStr != bigIFxDecNumberStr\n"+
			"  Length of fdEulers1kNumStr = '%v'\n"+
			"Length of bigIFxDecNumberStr = '%v'\n\n",
			ePrefix, fdEulers1kNumStrLength, bigIFxDecNumberStrLen)

		return
	}

	return
}

func TestBigIntNum_NewNumStr_18(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStr_18"

	numStr := "abcdefghijklmnop"

	_, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error, BUT NO ERROR WAS RETURNED!\n"+
			"_, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr= '%v'\n\n",
			ePrefix, numStr)

		return
	}
}

func TestBigIntNum_NewNumStrWithNumSeps_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStrWithNumSeps_01"

	expectedNumberStr := "123,456"

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewNumStrWithNumSeps_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewNumStrWithNumSeps_02"

	expectedNumberStr := "123.456"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewBigIntExponent_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewBigIntExponent_01"

	originalNumI64 := int64(123456)

	bOriginalBigIntNum := big.NewInt(originalNumI64)

	originalTensExponentInt := 3

	expectedNumberStr := "123456.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "123456000"

	expectedAbsBigIntNumStr := "123456000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewBigIntExponent(bOriginalBigIntNum, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigIntExponent(bOriginalBigIntNum, originalTensExponentInt)\n"+
			"bOriginalBigIntNum= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bOriginalBigIntNum.Text(10),
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewBigIntExponent_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewBigIntExponent_02"

	originalNumI64 := int64(123456)

	bOriginalBigIntNum := big.NewInt(originalNumI64)

	originalTensExponentInt := -3

	expectedNumberStr := "123.456"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewBigIntExponent(bOriginalBigIntNum, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigIntExponent(bOriginalBigIntNum, originalTensExponentInt)\n"+
			"bOriginalBigIntNum= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bOriginalBigIntNum.Text(10),
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewBigIntExponent_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewBigIntExponent_03"

	originalNumI64 := int64(-123456)

	bOriginalBigIntNum := big.NewInt(originalNumI64)

	originalTensExponentInt := 3

	expectedNumberStr := "-123456.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "-123456000"

	expectedAbsBigIntNumStr := "123456000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := -1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewBigIntExponent(bOriginalBigIntNum, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigIntExponent(bOriginalBigIntNum, originalTensExponentInt)\n"+
			"bOriginalBigIntNum= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bOriginalBigIntNum.Text(10),
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewBigIntExponent_04(t *testing.T) {

	ePrefix := "TestBigIntNum_NewBigIntExponent_04"

	originalNumI64 := int64(-123456)

	bOriginalBigIntNum := big.NewInt(originalNumI64)

	originalTensExponentInt := -3

	expectedNumberStr := "-123.456"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "-123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := -1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewBigIntExponent(bOriginalBigIntNum, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigIntExponent(bOriginalBigIntNum, originalTensExponentInt)\n"+
			"bOriginalBigIntNum= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bOriginalBigIntNum.Text(10),
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewOne_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewOne_01"

	expectedNumberStr := "1.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "1000"

	expectedAbsBigIntNumStr := "1000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewOne(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewOne(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewOne_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewOne_02"

	expectedNumberStr := "1"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "1"

	expectedAbsBigIntNumStr := "1"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewOne(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewOne(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewOne_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewOne_03"

	expectedNumberStr := "1.00000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(5)

	expectedBigIntNumStr := "100000"

	expectedAbsBigIntNumStr := "100000"

	expectedScaleFactor := big.NewInt(100000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewOne(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewOne(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewTwo_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewTwo_01"

	expectedNumberStr := "2.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "2000"

	expectedAbsBigIntNumStr := "2000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewTwo(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewTwo(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewTwo_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewTwo_02"

	expectedNumberStr := "2"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "2"

	expectedAbsBigIntNumStr := "2"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewTwo(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewTwo(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewTwo_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewTwo_03"

	expectedNumberStr := "2.00000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(5)

	expectedBigIntNumStr := "200000"

	expectedAbsBigIntNumStr := "200000"

	expectedScaleFactor := big.NewInt(100000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewTwo(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewTwo(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewThree_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewThree_01"

	expectedNumberStr := "3.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "3000"

	expectedAbsBigIntNumStr := "3000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewThree(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewThree(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewThree_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewThree_02"

	expectedNumberStr := "3"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "3"

	expectedAbsBigIntNumStr := "3"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewThree(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewThree(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewThree_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewThree_03"

	expectedNumberStr := "3.00000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(5)

	expectedBigIntNumStr := "300000"

	expectedAbsBigIntNumStr := "300000"

	expectedScaleFactor := big.NewInt(100000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewThree(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewThree(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewFive_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewFive_01"

	expectedNumberStr := "5.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "5000"

	expectedAbsBigIntNumStr := "5000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewFive(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewFive(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewFive_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewFive_02"

	expectedNumberStr := "5"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "5"

	expectedAbsBigIntNumStr := "5"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewFive(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewFive(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewFive_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewFive_03"

	expectedNumberStr := "5.00000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(5)

	expectedBigIntNumStr := "500000"

	expectedAbsBigIntNumStr := "500000"

	expectedScaleFactor := big.NewInt(100000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewFive(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewFive(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewTen_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewTen_01"

	expectedNumberStr := "10.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "10000"

	expectedAbsBigIntNumStr := "10000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewTen(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewTen(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewTen_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewTen_02"

	expectedNumberStr := "10"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "10"

	expectedAbsBigIntNumStr := "10"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewTen(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewTen(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewTen_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewTen_03"

	expectedNumberStr := "10.00000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(5)

	expectedBigIntNumStr := "1000000"

	expectedAbsBigIntNumStr := "1000000"

	expectedScaleFactor := big.NewInt(100000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewTen(expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewTen(expectedPrecisionUint)\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint_01"

	originalNumUint := uint(1234)

	expectedPrecisionUint := uint(3)

	expectedNumberStr := "1.234"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNumStr := "1234"

	expectedAbsBigIntNumStr := "1234"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint(originalNumUint, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint(originalNumUint, expectedPrecisionUint)\n"+
			"originalNumUint= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint_02"

	originalNumUint := uint(1234)

	expectedPrecisionUint := uint(0)

	expectedNumberStr := "1234"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNumStr := "1234"

	expectedAbsBigIntNumStr := "1234"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint(originalNumUint, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint(originalNumUint, expectedPrecisionUint)\n"+
			"originalNumUint= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint_03"

	originalNumUint := uint(0)

	expectedPrecisionUint := uint(0)

	expectedNumberStr := "0"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNumStr := "0"

	expectedAbsBigIntNumStr := "0"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint(originalNumUint, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint(originalNumUint, expectedPrecisionUint)\n"+
			"originalNumUint= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint_04(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint_04"

	originalNumUint := uint(0)

	expectedPrecisionUint := uint(3)

	expectedNumberStr := "0.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNumStr := "0000"

	expectedAbsBigIntNumStr := "0000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint(originalNumUint, expectedPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint(originalNumUint, expectedPrecisionUint)\n"+
			"originalNumUint= '%v'\n"+
			"expectedPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			expectedPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUintExponent_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUintExponent_01"

	originalNumUint := uint(1234)

	originalTensExponentInt := 3

	expectedNumberStr := "1234.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "1234000"

	expectedAbsBigIntNumStr := "1234000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUintExponent(originalNumUint, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUintExponent(originalNumUint, originalTensExponentInt)\n"+
			"originalNumUint= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUintExponent_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUintExponent_02"

	originalNumUint := uint(123456)

	originalTensExponentInt := -2

	expectedNumberStr := "1234.56"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(2)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(100)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUintExponent(originalNumUint, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUintExponent(originalNumUint, originalTensExponentInt)\n"+
			"originalNumUint= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUintExponent_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUintExponent_03"

	originalNumUint := uint(123456)

	originalTensExponentInt := 0

	expectedNumberStr := "123456"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUintExponent(originalNumUint, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUintExponent(originalNumUint, originalTensExponentInt)\n"+
			"originalNumUint= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUintExponent_04(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUintExponent_04"

	originalNumUint := uint(0)

	originalTensExponentInt := 0

	expectedNumberStr := "0"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "0"

	expectedAbsBigIntNumStr := "0"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUintExponent(originalNumUint, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUintExponent(originalNumUint, originalTensExponentInt)\n"+
			"originalNumUint= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUintExponent_05(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUintExponent_05"

	originalNumUint := uint(0)

	originalTensExponentInt := 3

	expectedNumberStr := "0.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "0000"

	expectedAbsBigIntNumStr := "0000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUintExponent(originalNumUint, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUintExponent(originalNumUint, originalTensExponentInt)\n"+
			"originalNumUint= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUintExponent_06(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUintExponent_06"

	originalNumUint := uint(0)

	originalTensExponentInt := -3

	expectedNumberStr := "0.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "0000"

	expectedAbsBigIntNumStr := "0000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUintExponent(originalNumUint, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUintExponent(originalNumUint, originalTensExponentInt)\n"+
			"originalNumUint= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32_01"

	originalNumUint32 := uint32(1234)

	originalPrecisionUint := uint(3)

	expectedNumberStr := "1.234"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "1234"

	expectedAbsBigIntNumStr := "1234"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32(originalNumUint32, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32(originalNumUint32, originalPrecisionUint)\n"+
			"originalNumUint32= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint32,
			originalPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32_02"

	originalNumUint32 := uint32(1234)

	originalPrecisionUint := uint(0)

	expectedNumberStr := "1234"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "1234"

	expectedAbsBigIntNumStr := "1234"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32(originalNumUint32, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32(originalNumUint32, originalPrecisionUint)\n"+
			"originalNumUint32= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint32,
			originalPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32_03"

	originalNumUint32 := uint32(0)

	originalPrecisionUint := uint(0)

	expectedNumberStr := "0"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "0"

	expectedAbsBigIntNumStr := "0"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32(originalNumUint32, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32(originalNumUint32, originalPrecisionUint)\n"+
			"originalPrecisionUint= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint32,
			originalPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32_04(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32_04"

	originalNumUint32 := uint32(0)

	originalPrecisionUint := uint(3)

	expectedNumberStr := "0.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "0000"

	expectedAbsBigIntNumStr := "0000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32(originalNumUint32, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32(originalNumUint32, originalPrecisionUint)\n"+
			"originalNumUint32= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumUint32,
			originalPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32Exponent_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32Exponent_01"

	originalBaseNumUint32 := uint32(1234)

	originalTensExponentInt := 3

	expectedNumberStr := "1234.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "1234000"

	expectedAbsBigIntNumStr := "1234000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint32,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32Exponent_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32Exponent_02"

	originalBaseNumUint32 := uint32(123456)

	originalTensExponentInt := -2

	expectedNumberStr := "1234.56"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(2)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(100)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint32,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32Exponent_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32Exponent_03"

	originalBaseNumUint32 := uint32(123456)

	originalTensExponentInt := 0

	expectedNumberStr := "123456"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint32,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32Exponent_04(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32Exponent_04"

	originalBaseNumUint32 := uint32(0)

	originalTensExponentInt := 0

	expectedNumberStr := "0"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "0"

	expectedAbsBigIntNumStr := "0"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint32,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32Exponent_05(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32Exponent_05"

	originalBaseNumUint32 := uint32(0)

	originalTensExponentInt := 3

	expectedNumberStr := "0.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "0000"

	expectedAbsBigIntNumStr := "0000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint32,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint32Exponent_06(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint32Exponent_06"

	originalBaseNumUint32 := uint32(0)

	originalTensExponentInt := -3

	expectedNumberStr := "0.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "0000"

	expectedAbsBigIntNumStr := "0000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint32Exponent(originalBaseNumUint32, originalTensExponentInt)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint32,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64_01"

	originalBaseNumUint64 := uint64(1234)

	originalPrecisionUint := uint(3)

	expectedNumberStr := "1.234"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "1234"

	expectedAbsBigIntNumStr := "1234"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64(originalBaseNumUint64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64(originalBaseNumUint64, originalPrecisionUint)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64_02"

	originalBaseNumUint64 := uint64(1234)

	originalPrecisionUint := uint(0)

	expectedNumberStr := "1234"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "1234"

	expectedAbsBigIntNumStr := "1234"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64(originalBaseNumUint64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64(originalBaseNumUint64, originalPrecisionUint)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64_03"

	originalBaseNumUint64 := uint64(0)

	originalPrecisionUint := uint(0)

	expectedNumberStr := "0"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "0"

	expectedAbsBigIntNumStr := "0"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64(originalBaseNumUint64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64(originalBaseNumUint64, originalPrecisionUint)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64_04(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64_04"

	originalBaseNumUint64 := uint64(0)

	originalPrecisionUint := uint(3)

	expectedNumberStr := "0.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := originalPrecisionUint

	expectedBigIntNumStr := "0000"

	expectedAbsBigIntNumStr := "0000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64(originalBaseNumUint64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64(originalBaseNumUint64, originalPrecisionUint)\n"+
			"originalBaseNumUint32= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalPrecisionUint,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64Exponent_01(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64Exponent_01"

	originalBaseNumUint64 := uint64(1234)

	originalTensExponentInt := 3

	expectedNumberStr := "1234.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "1234000"

	expectedAbsBigIntNumStr := "1234000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)\n"+
			"originalBaseNumUint64= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64Exponent_02(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64Exponent_02"

	originalBaseNumUint64 := uint64(123456)

	originalTensExponentInt := -2

	expectedNumberStr := "1234.56"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(2)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(100)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)\n"+
			"originalBaseNumUint64= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64Exponent_03(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64Exponent_03"

	originalBaseNumUint64 := uint64(123456)

	originalTensExponentInt := 0

	expectedNumberStr := "123456"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "123456"

	expectedAbsBigIntNumStr := "123456"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)\n"+
			"originalBaseNumUint64= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64Exponent_04(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64Exponent_04"

	originalBaseNumUint64 := uint64(0)

	originalTensExponentInt := 0

	expectedNumberStr := "0"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "0"

	expectedAbsBigIntNumStr := "0"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)\n"+
			"originalBaseNumUint64= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64Exponent_05(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64Exponent_05"

	originalBaseNumUint64 := uint64(0)

	originalTensExponentInt := 3

	expectedNumberStr := "0.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "0000"

	expectedAbsBigIntNumStr := "0000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)\n"+
			"originalBaseNumUint64= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_NewUint64Exponent_06(t *testing.T) {

	ePrefix := "TestBigIntNum_NewUint64Exponent_06"

	originalBaseNumUint64 := uint64(0)

	originalTensExponentInt := -3

	expectedNumberStr := "0.000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "0000"

	expectedAbsBigIntNumStr := "0000"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewUint64Exponent(originalBaseNumUint64, originalTensExponentInt)\n"+
			"originalBaseNumUint64= '%v'\n"+
			"originalTensExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBaseNumUint64,
			originalTensExponentInt,
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

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}
