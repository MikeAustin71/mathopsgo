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

  expectedBigIntNumStr := "-123456"

  expectedAbsBigIntNumStr := "123456"

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

  expectedScaleFactor := big.NewInt(0)

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

  expectedScaleFactor := big.NewInt(0)

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

  expectedScaleFactor := big.NewInt(0)

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

  expectedScaleFactor := big.NewInt(0)

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

  expectedScaleFactor := big.NewInt(0)

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
  expectedNumStr := "3.000"
  expectedPrecision := uint(3)

  bINum := BigIntNum{}.NewThree(expectedPrecision)

  if expectedNumStr != bINum.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, bINum.GetNumStr())
  }

  if expectedPrecision != bINum.GetPrecisionUint() {
    t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
      expectedPrecision, bINum.GetPrecisionUint())
  }

}

func TestBigIntNum_NewThree_02(t *testing.T) {
  expectedNumStr := "3"
  expectedPrecision := uint(0)

  bINum := BigIntNum{}.NewThree(expectedPrecision)

  if expectedNumStr != bINum.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, bINum.GetNumStr())
  }

  if expectedPrecision != bINum.GetPrecisionUint() {
    t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
      expectedPrecision, bINum.GetPrecisionUint())
  }

}

func TestBigIntNum_NewThree_03(t *testing.T) {
  expectedNumStr := "3.00000"
  expectedPrecision := uint(5)

  bINum := BigIntNum{}.NewThree(expectedPrecision)

  if expectedNumStr != bINum.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, bINum.GetNumStr())
  }

  if expectedPrecision != bINum.GetPrecisionUint() {
    t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
      expectedPrecision, bINum.GetPrecisionUint())
  }

}

func TestBigIntNum_NewFive_01(t *testing.T) {
  expectedNumStr := "5.000"
  expectedPrecision := uint(3)

  bINum := BigIntNum{}.NewFive(expectedPrecision)

  if expectedNumStr != bINum.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, bINum.GetNumStr())
  }

  if expectedPrecision != bINum.GetPrecisionUint() {
    t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
      expectedPrecision, bINum.GetPrecisionUint())
  }

}

func TestBigIntNum_NewFive_02(t *testing.T) {
  expectedNumStr := "5"
  expectedPrecision := uint(0)

  bINum := BigIntNum{}.NewFive(expectedPrecision)

  if expectedNumStr != bINum.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, bINum.GetNumStr())
  }

  if expectedPrecision != bINum.GetPrecisionUint() {
    t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
      expectedPrecision, bINum.GetPrecisionUint())
  }

}

func TestBigIntNum_NewFive_03(t *testing.T) {
  expectedNumStr := "5.00000"
  expectedPrecision := uint(5)

  bINum := BigIntNum{}.NewFive(expectedPrecision)

  if expectedNumStr != bINum.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, bINum.GetNumStr())
  }

  if expectedPrecision != bINum.GetPrecisionUint() {
    t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
      expectedPrecision, bINum.GetPrecisionUint())
  }
}

func TestBigIntNum_NewTen_01(t *testing.T) {
  expectedNumStr := "10.000"
  expectedPrecision := uint(3)

  bINum := BigIntNum{}.NewTen(expectedPrecision)

  if expectedNumStr != bINum.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, bINum.GetNumStr())
  }

  if expectedPrecision != bINum.GetPrecisionUint() {
    t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
      expectedPrecision, bINum.GetPrecisionUint())
  }

}

func TestBigIntNum_NewTen_02(t *testing.T) {
  expectedNumStr := "10"
  expectedPrecision := uint(0)

  bINum := BigIntNum{}.NewTen(expectedPrecision)

  if expectedNumStr != bINum.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, bINum.GetNumStr())
  }

  if expectedPrecision != bINum.GetPrecisionUint() {
    t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
      expectedPrecision, bINum.GetPrecisionUint())
  }

}

func TestBigIntNum_NewTen_03(t *testing.T) {
  expectedNumStr := "10.00000"
  expectedPrecision := uint(5)

  bINum := BigIntNum{}.NewTen(expectedPrecision)

  if expectedNumStr != bINum.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, bINum.GetNumStr())
  }

  if expectedPrecision != bINum.GetPrecisionUint() {
    t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
      expectedPrecision, bINum.GetPrecisionUint())
  }

}

func TestBigIntNum_NewUint_01(t *testing.T) {

  numInt := uint(1234)
  precision := uint(3)
  expectedNumStr := "1.234"

  bINum := BigIntNum{}.NewUint(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint_02(t *testing.T) {

  numInt := uint(1234)
  precision := uint(0)
  expectedNumStr := "1234"

  bINum := BigIntNum{}.NewUint(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint_03(t *testing.T) {

  numInt := uint(0)
  precision := uint(0)
  expectedNumStr := "0"

  bINum := BigIntNum{}.NewUint(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint_04(t *testing.T) {

  numInt := uint(0)
  precision := uint(3)
  expectedNumStr := "0.000"

  bINum := BigIntNum{}.NewUint(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUintExponent_01(t *testing.T) {

  numInt := uint(1234)
  exponent := 3
  expectedNumStr := "1234.000"

  bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUintExponent_02(t *testing.T) {

  numInt := uint(123456)
  exponent := -2
  expectedNumStr := "1234.56"

  bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUintExponent_03(t *testing.T) {

  numInt := uint(123456)
  exponent := 0
  expectedNumStr := "123456"

  bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUintExponent_04(t *testing.T) {

  numInt := uint(0)
  exponent := 0
  expectedNumStr := "0"

  bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUintExponent_05(t *testing.T) {

  numInt := uint(0)
  exponent := 3
  expectedNumStr := "0.000"

  bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUintExponent_06(t *testing.T) {

  numInt := uint(0)
  exponent := -3
  expectedNumStr := "0.000"

  bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32_01(t *testing.T) {

  num32Uint := uint32(1234)
  precision := uint(3)
  expectedNumStr := "1.234"

  bINum := BigIntNum{}.NewUint32(num32Uint, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32_02(t *testing.T) {

  num32Uint := uint32(1234)
  precision := uint(0)
  expectedNumStr := "1234"

  bINum := BigIntNum{}.NewUint32(num32Uint, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32_03(t *testing.T) {

  num32Uint := uint32(0)
  precision := uint(0)
  expectedNumStr := "0"

  bINum := BigIntNum{}.NewUint32(num32Uint, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32_04(t *testing.T) {

  num32Uint := uint32(0)
  precision := uint(3)
  expectedNumStr := "0.000"

  bINum := BigIntNum{}.NewUint32(num32Uint, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32Exponent_01(t *testing.T) {

  numInt := uint32(1234)
  exponent := 3
  expectedNumStr := "1234.000"

  bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32Exponent_02(t *testing.T) {

  numInt := uint32(123456)
  exponent := -2
  expectedNumStr := "1234.56"

  bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32Exponent_03(t *testing.T) {

  numInt := uint32(123456)
  exponent := 0
  expectedNumStr := "123456"

  bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32Exponent_04(t *testing.T) {

  numInt := uint32(0)
  exponent := 0
  expectedNumStr := "0"

  bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32Exponent_05(t *testing.T) {

  numInt := uint32(0)
  exponent := 3
  expectedNumStr := "0.000"

  bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint32Exponent_06(t *testing.T) {

  numInt := uint32(0)
  exponent := -3
  expectedNumStr := "0.000"

  bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64_01(t *testing.T) {

  num64Uint := uint64(1234)
  precision := uint(3)
  expectedNumStr := "1.234"

  bINum := BigIntNum{}.NewUint64(num64Uint, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64_02(t *testing.T) {

  num64Uint := uint64(1234)
  precision := uint(0)
  expectedNumStr := "1234"

  bINum := BigIntNum{}.NewUint64(num64Uint, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64_03(t *testing.T) {

  num64Uint := uint64(0)
  precision := uint(0)
  expectedNumStr := "0"

  bINum := BigIntNum{}.NewUint64(num64Uint, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64_04(t *testing.T) {

  num64Uint := uint64(0)
  precision := uint(3)
  expectedNumStr := "0.000"

  bINum := BigIntNum{}.NewUint64(num64Uint, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64Exponent_01(t *testing.T) {

  numInt := uint64(1234)
  exponent := 3
  expectedNumStr := "1234.000"

  bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64Exponent_02(t *testing.T) {

  numInt := uint64(123456)
  exponent := -2
  expectedNumStr := "1234.56"

  bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64Exponent_03(t *testing.T) {

  numInt := uint64(123456)
  exponent := 0
  expectedNumStr := "123456"

  bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64Exponent_04(t *testing.T) {

  numInt := uint64(0)
  exponent := 0
  expectedNumStr := "0"

  bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64Exponent_05(t *testing.T) {

  numInt := uint64(0)
  exponent := 3
  expectedNumStr := "0.000"

  bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewUint64Exponent_06(t *testing.T) {

  numInt := uint64(0)
  exponent := -3
  expectedNumStr := "0.000"

  bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}
