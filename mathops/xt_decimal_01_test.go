package mathops

import (
  "fmt"
  "math/big"
  "testing"
)

/*
	These tests are associated with the library routines contained in source code file
	decimal.go. The source code repository for these tests is located at:

				https://github.com/MikeAustin71/decimalnum.git
*/

func TestNumStrUtility_ConvertNumStrToDecimal_01(t *testing.T) {

  ePrefix := "TestNumStrUtility_ConvertNumStrToDecimal_01"

  originalNumStr := "123456.654321"

  expectedPrecisionUint := uint(6)

  expectedPrecisionInt := int(expectedPrecisionUint)

  expectedBigIntNumStr := "123456654321"

  expectedAbsBigIntNumStr := "123456654321"

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nsu := new(NumStrUtility)

  err := nsu.SetNumSeps(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := nsu.SetNumSeps(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumSeps.String(), err.Error())

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

  dec, err := nsu.ConvertNumStrToDecimal(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec, err := nsu.\n"+
      "  ConvertNumStrToDecimal(originalNumStr)\n"+
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumStr, err.Error())
    return
  }

  err = dec.IsValid("Validating dec")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec.IsValid('Validating dec')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumberStr, err := dec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumberStr, err := dec.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != decNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumStr != decNumberStr \n"+
      "Expected decNumberStr = '%v'\n"+
      "  Actual decNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, decNumberStr)

    return
  }

  decBigInt, err := dec.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decBigInt, err := dec.GetBigInt()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(decBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/dec Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(decBigInt) != 0\n"+
      "Expected decBigInt = '%v'\n"+
      "  Actual decBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), decBigInt.Text(10))

    return
  }

  decAbsValue, err := dec.GetAbsoluteValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAbsValue, err := dec.GetAbsoluteValue()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  decAbsBigInt, err := decAbsValue.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAbsBigInt, err := decAbsValue.GetBigInt()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(decAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & dec Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(decAbsBigInt) != 0\n"+
      "Expected decAbsBigInt = '%v'\n"+
      "  Actual decAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), decAbsBigInt.Text(10))

    return
  }

  decPrecisionUint, err := dec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decPrecisionUint, err := dec.GetPrecisionUint()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != decPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decPrecisionUint\n"+
      "Expected decPrecisionUint = '%v'\n"+
      "  Actual decPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decPrecisionUint)

    return
  }

  decBigIntNumStr, err := dec.GetSignedAllDigitsStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decBigIntNumStr, err := dec.GetSignedAllDigitsStr()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNumStr != decBigIntNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Big Int Number Strings ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNumStr != decBigIntNumStr\n"+
      "Expected decBigIntNumStr = '%v'\n"+
      "  Actual decBigIntNumStr = '%v'\n\n",
      ePrefix, expectedAbsBigIntNumStr, decBigIntNumStr)

    return
  }

  decSignValue, err := dec.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSignValue, err := dec.GetSign()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedSignVal != decSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != decSignValue\n"+
      "Expected decSignValue = '%v'\n"+
      "  Actual decSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decSignValue)

    return
  }

  decSeps, err := dec.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSeps, err := dec.GetNumericSeparatorsDto()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(decSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != decSeps \n"+
      "Expected decSeps = '%v'\n"+
      "  Actual decSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decSeps.String())

    return
  }

  _, accuracy, err := dec.GetFloat64()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, accuracy, err := dec.GetFloat64()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  accuracyStr := accuracy.String()

  if accuracyStr != "Exact" {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because accuracyStr != 'Exact'\n"+
      "Expected accuracyStr = 'Exact'\n"+
      "  Actual accuracyStr = '%v'\n\n",
      ePrefix, accuracyStr)

    return
  }

  bf, err := dec.GetBigFloat()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bf, err := dec.GetBigFloat()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  bigFloatNumberStr := fmt.Sprintf("%s", bf.Text('f', expectedPrecisionInt))

  if originalNumStr != bigFloatNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumStr != bigFloatNumberStr\n"+
      "Expected bigFloatNumberStr = '%v'\n"+
      "  Actual bigFloatNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bigFloatNumberStr)

    return
  }

  return
}

func TestNumStrUtility_ConvertNumStrToDecimal_02(t *testing.T) {

  ePrefix := "TestNumStrUtility_ConvertNumStrToDecimal_02"

  originalNumStr := "-123456.654321"

  expectedPrecisionUint := uint(6)

  expectedPrecisionInt := int(expectedPrecisionUint)

  expectedBigIntNumStr := "-123456654321"

  expectedAbsBigIntNumStr := "123456654321"

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nsu := new(NumStrUtility)

  err := nsu.SetNumSeps(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := nsu.SetNumSeps(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumSeps.String(), err.Error())

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

  dec, err := nsu.ConvertNumStrToDecimal(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec, err := nsu.\n"+
      "  ConvertNumStrToDecimal(originalNumStr)\n"+
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumStr, err.Error())
    return
  }

  err = dec.IsValid("Validating dec")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec.IsValid('Validating dec')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumberStr, err := dec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumberStr, err := dec.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != decNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumStr != decNumberStr \n"+
      "Expected decNumberStr = '%v'\n"+
      "  Actual decNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, decNumberStr)

    return
  }

  decBigInt, err := dec.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decBigInt, err := dec.GetBigInt()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(decBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/dec Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(decBigInt) != 0\n"+
      "Expected decBigInt = '%v'\n"+
      "  Actual decBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), decBigInt.Text(10))

    return
  }

  decAbsValue, err := dec.GetAbsoluteValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAbsValue, err := dec.GetAbsoluteValue()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  decAbsBigInt, err := decAbsValue.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAbsBigInt, err := decAbsValue.GetBigInt()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(decAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & dec Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(decAbsBigInt) != 0\n"+
      "Expected decAbsBigInt = '%v'\n"+
      "  Actual decAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), decAbsBigInt.Text(10))

    return
  }

  decPrecisionUint, err := dec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decPrecisionUint, err := dec.GetPrecisionUint()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != decPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decPrecisionUint\n"+
      "Expected decPrecisionUint = '%v'\n"+
      "  Actual decPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decPrecisionUint)

    return
  }

  decBigIntNumStr, err := dec.GetSignedAllDigitsStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decBigIntNumStr, err := dec.GetSignedAllDigitsStr()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNumStr != decBigIntNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Big Int Number Strings ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNumStr != decBigIntNumStr\n"+
      "Expected decBigIntNumStr = '%v'\n"+
      "  Actual decBigIntNumStr = '%v'\n\n",
      ePrefix, expectedAbsBigIntNumStr, decBigIntNumStr)

    return
  }

  decSignValue, err := dec.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSignValue, err := dec.GetSign()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedSignVal != decSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != decSignValue\n"+
      "Expected decSignValue = '%v'\n"+
      "  Actual decSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decSignValue)

    return
  }

  decSeps, err := dec.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSeps, err := dec.GetNumericSeparatorsDto()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(decSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != decSeps \n"+
      "Expected decSeps = '%v'\n"+
      "  Actual decSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decSeps.String())

    return
  }

  _, accuracy, err := dec.GetFloat64()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, accuracy, err := dec.GetFloat64()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  accuracyStr := accuracy.String()

  if accuracyStr != "Exact" {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because accuracyStr != 'Exact'\n"+
      "Expected accuracyStr = 'Exact'\n"+
      "  Actual accuracyStr = '%v'\n\n",
      ePrefix, accuracyStr)

    return
  }

  bf, err := dec.GetBigFloat()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bf, err := dec.GetBigFloat()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  bigFloatNumberStr := fmt.Sprintf("%s", bf.Text('f', expectedPrecisionInt))

  if originalNumStr != bigFloatNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumStr != bigFloatNumberStr\n"+
      "Expected bigFloatNumberStr = '%v'\n"+
      "  Actual bigFloatNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bigFloatNumberStr)

    return
  }

  return
}

func TestNumStrUtility_ConvertNumStrToDecimal_03(t *testing.T) {

  ePrefix := "TestNumStrUtility_ConvertNumStrToDecimal_03"

  rawNumStr := "zyx -123456.654321 xyx"

  originalNumStr := "-123456.654321"

  expectedPrecisionUint := uint(6)

  expectedPrecisionInt := int(expectedPrecisionUint)

  expectedBigIntNumStr := "-123456654321"

  expectedAbsBigIntNumStr := "123456654321"

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nsu := new(NumStrUtility).New()

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

  dec, err := nsu.ConvertNumStrToDecimal(rawNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec, err := nsu.\n"+
      "  ConvertNumStrToDecimal(rawNumStr)\n"+
      "rawNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumStr, err.Error())
    return
  }

  err = dec.IsValid("Validating dec")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec.IsValid('Validating dec')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumberStr, err := dec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumberStr, err := dec.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != decNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumStr != decNumberStr \n"+
      "Expected decNumberStr = '%v'\n"+
      "  Actual decNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, decNumberStr)

    return
  }

  decBigInt, err := dec.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decBigInt, err := dec.GetBigInt()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(decBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/dec Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(decBigInt) != 0\n"+
      "Expected decBigInt = '%v'\n"+
      "  Actual decBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), decBigInt.Text(10))

    return
  }

  decAbsValue, err := dec.GetAbsoluteValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAbsValue, err := dec.GetAbsoluteValue()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  decAbsBigInt, err := decAbsValue.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAbsBigInt, err := decAbsValue.GetBigInt()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(decAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & dec Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(decAbsBigInt) != 0\n"+
      "Expected decAbsBigInt = '%v'\n"+
      "  Actual decAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), decAbsBigInt.Text(10))

    return
  }

  decPrecisionUint, err := dec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decPrecisionUint, err := dec.GetPrecisionUint()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != decPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decPrecisionUint\n"+
      "Expected decPrecisionUint = '%v'\n"+
      "  Actual decPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decPrecisionUint)

    return
  }

  decBigIntNumStr, err := dec.GetSignedAllDigitsStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decBigIntNumStr, err := dec.GetSignedAllDigitsStr()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNumStr != decBigIntNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Big Int Number Strings ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNumStr != decBigIntNumStr\n"+
      "Expected decBigIntNumStr = '%v'\n"+
      "  Actual decBigIntNumStr = '%v'\n\n",
      ePrefix, expectedAbsBigIntNumStr, decBigIntNumStr)

    return
  }

  decSignValue, err := dec.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSignValue, err := dec.GetSign()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedSignVal != decSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != decSignValue\n"+
      "Expected decSignValue = '%v'\n"+
      "  Actual decSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decSignValue)

    return
  }

  decSeps, err := dec.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSeps, err := dec.GetNumericSeparatorsDto()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(decSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != decSeps \n"+
      "Expected decSeps = '%v'\n"+
      "  Actual decSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decSeps.String())

    return
  }

  _, accuracy, err := dec.GetFloat64()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, accuracy, err := dec.GetFloat64()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  accuracyStr := accuracy.String()

  if accuracyStr != "Exact" {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because accuracyStr != 'Exact'\n"+
      "Expected accuracyStr = 'Exact'\n"+
      "  Actual accuracyStr = '%v'\n\n",
      ePrefix, accuracyStr)

    return
  }

  bf, err := dec.GetBigFloat()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bf, err := dec.GetBigFloat()\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
    return
  }

  bigFloatNumberStr := fmt.Sprintf("%s", bf.Text('f', expectedPrecisionInt))

  if originalNumStr != bigFloatNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumStr != bigFloatNumberStr\n"+
      "Expected bigFloatNumberStr = '%v'\n"+
      "  Actual bigFloatNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bigFloatNumberStr)

    return
  }

  return
}

func TestNumStrUtility_ConvertNumStrToDecimal_04(t *testing.T) {

  ePrefix := "TestNumStrUtility_ConvertNumStrToDecimal_04"

  rawStr := "Nothing"

  nsu := NumStrUtility{}

  _, err := nsu.ConvertNumStrToDecimal(rawStr)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected Error to be generated by INVALID Number String. "+
      "Instead, NO ERROR WAS RETURNED!\n"+
      "_, err := nsu.ConvertNumStrToDecimal(rawStr)\n"+
      "rawStr='%v'\n\n",
      ePrefix, rawStr)
  }

}

func TestDecimal_AllDigitsNumStr_01(t *testing.T) {

  ePrefix := "TestDecimal_AllDigitsNumStr_01"

  rawNumStr := "x 123.456 x"

  expectedNumberStr := "123456"

  dec := new(Decimal).New()

  actualAbsNumberStr, err := dec.AllDigitsNumStr(rawNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualAbsNumberStr, err := dec.AllDigitsNumStr(rawNumStr)\n"+
      "rawNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, rawNumStr, err.Error())
    return
  }

  if expectedNumberStr != actualAbsNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != actualAbsNumberStr\n"+
      "Expected actualAbsNumberStr = '%v'\n"+
      "  Actual actualAbsNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, actualAbsNumberStr)

    return
  }

  return
}

func TestDecimal_Add_01(t *testing.T) {

  ePrefix := "TestDecimal_Add_01"

  addend1 := "35.50"

  addend2 := "35.51"

  expectedSumNumberStr := "71.01"

  nu := new(NumStrUtility).New()

  decAddend1, err := nu.ConvertNumStrToDecimal(addend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend1, err := nu.ConvertNumStrToDecimal(addend1)\n"+
      "addend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, addend1, err.Error())
    return
  }

  err = decAddend1.IsValid("Validating decAddend1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend1.IsValid('Validating decAddend1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend1NumberStr, err := decAddend1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend1NumberStr, err := decAddend1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend2, err := nu.ConvertNumStrToDecimal(addend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend2, err := nu.ConvertNumStrToDecimal(addend2)\n"+
      "addend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, addend2, err.Error())
    return
  }

  err = decAddend2.IsValid("Validating decAddend2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend2.IsValid('Validating decAddend2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend2NumberStr, err := decAddend2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend2NumberStr, err := decAddend2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSum3, err := decAddend1.Add(decAddend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSum3, err := decAddend1.Add(decAddend2)\n"+
      "decAddend1= '%v'\n"+
      "decAddend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decAddend1NumberStr,
      decAddend2NumberStr,
      err.Error())

    return
  }

  err = decSum3.IsValid("Validating decSum3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSum3.IsValid('Validating decSum3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSum3NumberStr, err := decSum3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSum3NumberStr, err := decSum3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedSumNumberStr != decSum3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedSumNumberStr != decSum3NumberStr\n"+
      "Expected decSum3NumberStr = '%v'\n"+
      "  Actual decSum3NumberStr = '%v'\n\n",
      ePrefix, expectedSumNumberStr, decSum3NumberStr)

    return
  }

  return
}

func TestDecimal_Add_02(t *testing.T) {

  ePrefix := "TestDecimal_Add_02"

  addend1 := "-35.50"

  addend2 := "35.51"

  expectedSumNumberStr := "0.01"

  decAddend1, err := new(Decimal).NewNumStr(addend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend1, err := new(Decimal).NewNumStr(addend1)\n"+
      "addend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, addend1, err.Error())
    return
  }

  err = decAddend1.IsValid("Validating decAddend1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend1.IsValid('Validating decAddend1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend1NumberStr, err := decAddend1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend1NumberStr, err := decAddend1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend2, err := new(Decimal).NewNumStr(addend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend2, err := new(Decimal).NewNumStr(addend2)\n"+
      "addend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, addend2, err.Error())
    return
  }

  err = decAddend2.IsValid("Validating decAddend2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend2.IsValid('Validating decAddend2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend2NumberStr, err := decAddend2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend2NumberStr, err := decAddend2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSum3, err := decAddend1.Add(decAddend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSum3, err := decAddend1.Add(decAddend2)\n"+
      "decAddend1= '%v'\n"+
      "decAddend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decAddend1NumberStr,
      decAddend2NumberStr,
      err.Error())

    return
  }

  err = decSum3.IsValid("Validating decSum3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSum3.IsValid('Validating decSum3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSum3NumberStr, err := decSum3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSum3NumberStr, err := decSum3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedSumNumberStr != decSum3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedSumNumberStr != decSum3NumberStr\n"+
      "Expected decSum3NumberStr = '%v'\n"+
      "  Actual decSum3NumberStr = '%v'\n\n",
      ePrefix, expectedSumNumberStr, decSum3NumberStr)

    return
  }

  return
}

func TestDecimal_Add_03(t *testing.T) {

  ePrefix := "TestDecimal_Add_03"

  numStr1 := "35.50"

  numStr2 := "35.51"

  sub1 := "71.01"

  numStr3 := "0.5"

  sub2 := "71.51"

  numStr4 := "1.00"

  sub3 := "72.51"

  expectedNumberStr := ""

  dec1, err := new(Decimal).NewNumStr(numStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec1, err := new(Decimal).NewNumStr(numStr1)\n"+
      "numStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr1, err.Error())
    return
  }

  err = dec1.IsValid("Validating dec1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec1.IsValid('Validating dec1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec1NumberStr, err := dec1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec1NumberStr, err := dec1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec2, err := new(Decimal).NewNumStr(numStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2, err := new(Decimal).NewNumStr(numStr2)\n"+
      "numStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr1, err.Error())
    return
  }

  err = dec2.IsValid("Validating dec2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.IsValid('Validating dec2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec2NumberStr, err := dec2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2NumberStr, err := dec2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3, err := dec1.Add(dec2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3, err := dec1.Add(dec2)\n"+
      "dec1= '%v'\n"+
      "dec2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dec1NumberStr, dec2NumberStr, err.Error())
    return
  }

  err = dec3.IsValid("Validating dec3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec3.IsValid('Validating dec3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3NumberStr, err := dec3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3NumberStr, err := dec3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumberStr = sub1 // 71.01

  if expectedNumberStr != dec3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != dec3NumberStr\n"+
      "Expected dec3NumberStr = '%v'\n"+
      "  Actual dec3NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, dec3NumberStr)

    return
  }

  decy, err := new(Decimal).NewNumStr(numStr3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decy, err := new(Decimal).NewNumStr(numStr3)\n"+
      "numStr3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr3, err.Error())
    return
  }

  err = decy.IsValid("Validating decy")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decy.IsValid('Validating decy')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decyNumberStr, err := decy.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decyNumberStr, err := decy.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumberStr = "0.5"

  if expectedNumberStr != decyNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != decyNumberStr\n"+
      "Expected decyNumberStr = '%v'\n"+
      "  Actual decyNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decyNumberStr)

    return
  }

  dec4, err := dec3.Add(decy)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec4, err := dec3.Add(decy)\n"+
      "decy= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decyNumberStr, err.Error())
    return
  }

  err = dec4.IsValid("Validating dec4")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec4.IsValid('Validating dec4')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec4NumberStr, err := dec4.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec4NumberStr, err := dec4.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumberStr = sub2 // 71.51

  if expectedNumberStr != dec4NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != dec4NumberStr\n"+
      "Expected dec4NumberStr = '%v'\n"+
      "  Actual dec4NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, dec4NumberStr)

    return
  }

  err = dec2.SetNumStr(numStr4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.SetNumStr(numStr4)\n"+
      "numStr4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr1, err.Error())
    return
  }

  err = dec2.IsValid("Validating dec2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.IsValid('Validating dec2')\n"+
      "dec2 is set to 'numStr4'\n"+
      "numStr4= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, numStr4, err.Error())
    return
  }

  dec2NumberStr, err = dec2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2NumberStr, err = dec2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumberStr = "1.00"

  if expectedNumberStr != dec2NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != dec3NumberStr\n"+
      "Expected dec2NumberStr = '%v'\n"+
      "  Actual dec2NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, dec2NumberStr)

    return
  }

  // Re-use dec3
  dec3, err = dec4.Add(dec2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3, err = dec4.Add(dec2)\n"+
      "dec2= '%v'\n"+
      "dec4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dec2NumberStr, dec4NumberStr, err.Error())
    return
  }

  err = dec3.IsValid("Validating dec3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec3.IsValid('Validating dec3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3NumberStr, err = dec3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3NumberStr, err = dec3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumberStr = sub3 // 72.56

  if expectedNumberStr != dec3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != dec3NumberStr\n"+
      "Expected dec3NumberStr = '%v'\n"+
      "  Actual dec3NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, dec3NumberStr)

    return
  }

  return
}

func TestDecimal_Add_04(t *testing.T) {

  ePrefix := "TestDecimal_Add_04"

  numStr1 := "35.50"

  numStr2 := ".5"

  expectedNumberStr := "36.00"

  dec1 := Decimal{}

  err := dec1.SetNumStr(numStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := dec1.SetNumStr(numStr1)\n"+
      "numStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr1, err.Error())
    return
  }

  err = dec1.IsValid("Validating dec1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec1.IsValid('Validating dec1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec1NumberStr, err := dec1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec1NumberStr, err := dec1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec2, err := new(Decimal).NewNumStr(numStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2, err := new(Decimal).NewNumStr(numStr2)\n"+
      "numStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr2, err.Error())
    return
  }

  err = dec2.IsValid("Validating dec2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.IsValid('Validating dec2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec2NumberStr, err := dec2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2NumberStr, err := dec2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3, err := dec2.Add(dec1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3, err := dec2.Add(dec1)\n"+
      "dec1= '%v'\n"+
      "dec2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dec1NumberStr, dec2NumberStr, err.Error())
    return
  }

  err = dec3.IsValid("Validating dec3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec3.IsValid('Validating dec3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3NumberStr, err := dec3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3NumberStr, err := dec3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != dec3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != dec3NumberStr\n"+
      "Expected dec3NumberStr = '%v'\n"+
      "  Actual dec3NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, dec3NumberStr)

    return
  }

  return
}

func TestDecimal_Add_05(t *testing.T) {

  ePrefix := "TestDecimal_Add_05"

  numStr1 := "35.50"

  numStr2 := ".5"

  numStr3 := "1.00"

  numStr4 := "9.32"

  numStr5 := "101.912"

  expectedNumberStr := "148.232"

  dec1 := Decimal{}

  err := dec1.SetNumStr(numStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := dec1.SetNumStr(numStr1)\n"+
      "numStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr1, err.Error())
    return
  }

  err = dec1.IsValid("Validating dec1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec1.IsValid('Validating dec1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec1NumberStr, err := dec1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec1NumberStr, err := dec1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec2, err := new(Decimal).NewNumStr(numStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2, err := new(Decimal).NewNumStr(numStr2)\n"+
      "numStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr2, err.Error())
    return
  }

  err = dec2.IsValid("Validating dec2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.IsValid('Validating dec2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec2NumberStr, err := dec2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2NumberStr, err := dec2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3, err := dec2.Add(dec1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3, err := dec2.Add(dec1)\n"+
      "dec1= '%v'\n"+
      "dec2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dec1NumberStr, dec2NumberStr, err.Error())
    return
  }

  err = dec3.IsValid("Validating dec3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec3.IsValid('Validating dec3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3NumberStr, err := dec3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3NumberStr, err := dec3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dec2.SetNumStr(numStr3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.SetNumStr(numStr3)\n"+
      "numStr3= '%v'\n"+
      "dec2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr3, dec2NumberStr, err.Error())
    return
  }

  err = dec2.IsValid("Validating dec2-numStr3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.IsValid('Validating -numStr3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3, err = dec3.Add(dec2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3, err = dec3.Add(dec2)\n"+
      "dec2= '%v'\n"+
      "original - dec3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      dec2NumberStr,
      dec3NumberStr,
      err.Error())

    return
  }

  err = dec3.IsValid("Validating dec3++")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec3.IsValid('Validating dec3++')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3NumberStr, err = dec3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3NumberStr, err = dec3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dec2.SetNumStr(numStr4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.SetNumStr(numStr4)\n"+
      "numStr4= '%v'\n"+
      "original - dec2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStr4,
      dec2NumberStr,
      err.Error())

    return
  }

  err = dec2.IsValid("Validating dec2-numStr4")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.IsValid('Validating dec2-numStr4')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec2NumberStr, err = dec2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2NumberStr, err = dec2.GetNumStr()\n"+
      "dec2 = numStr4\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3, err = dec3.Add(dec2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3, err = dec3.Add(dec2)\n"+
      "dec2= '%v'\n"+
      "original - dec3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      dec2NumberStr,
      dec3NumberStr,
      err.Error())

    return
  }

  err = dec3.IsValid("Validating dec3- add dec2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec3.IsValid('Validating dec3- add dec2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3NumberStr, err = dec3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3NumberStr, err = dec3.GetNumStr()\n"+
      "dec3 - add dec2"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dec2.SetNumStr(numStr5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec2.SetNumStr(numStr5)\n"+
      "numStr5= '%v'\n"+
      "original - dec2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStr5,
      dec2NumberStr,
      err.Error())

    return
  }

  dec3, err = dec3.Add(dec2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3, err = dec3.Add(dec2)\n"+
      "dec2= '%v'\n"+
      "original - dec3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      dec2NumberStr,
      dec3NumberStr,
      err.Error())

    return
  }

  err = dec3.IsValid("Validating dec3- add dec2#2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dec3.IsValid('Validating dec3- add dec2#2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec3NumberStr, err = dec3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec3NumberStr, err = dec3.GetNumStr()\n"+
      "dec3 - add dec2"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != dec3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != dec3NumberStr\n"+
      "Expected dec3NumberStr = '%v'\n"+
      "  Actual dec3NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, dec3NumberStr)

    return
  }

  return
}

func TestDecimal_AddToThis_01(t *testing.T) {

  ePrefix := "TestDecimal_AddToThis_01"

  nStrAry := []string{
    "35.50",
    "36.50",
    "5.5",
    "92.75",
  }

  originalNumerStr := "0"

  expectedNumberStr := "170.25"

  decSum, err := new(Decimal).NewNumStr(originalNumerStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSum, err := new(Decimal).NewNumStr(originalNumerStr)\n"+
      "originalNumerStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumerStr, err.Error())
    return
  }

  err = decSum.IsValid("Validating decSum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSum.IsValid('Validating decSum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSumNumberStr, err := decSum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSumNumberStr, err := decSum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend := Decimal{}

  var decAddendNumberStr string

  for i := 0; i < len(nStrAry); i++ {

    err = decAddend.SetNumStr(nStrAry[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decAddend.SetNumStr(nStrAry[%v])\n"+
        "nStrAry[%v]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, nStrAry[i], err.Error())
      return
    }

    err = decAddend.IsValid(fmt.Sprintf("Validating decAddend[%v]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decAddend.IsValid('Validating decAddend[%v]')\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    decAddendNumberStr, err = decAddend.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decAddendNumberStr, err := decAddend.GetNumStr()\n"+
        "index i= '%v'"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    err = decSum.AddToThis(decAddend)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decSum.AddToThis(decAddend)\n"+
        "decAddend= '%v'\n"+
        "orignal-decSum= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        decAddendNumberStr,
        decSumNumberStr,
        err.Error())

      return
    }

    err = decSum.IsValid(fmt.Sprintf("Validating decSum[%v]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decSum.IsValid('Validating decSum[%v]')\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    decSumNumberStr, err = decSum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decSumNumberStr, err = decSum[%v].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }
  }

  if expectedNumberStr != decSumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != decSumNumberStr\n"+
      "Expected decSumNumberStr = '%v'\n"+
      "  Actual decSumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decSumNumberStr)

    return
  }

  return
}

func TestDecimal_AddToThis_02(t *testing.T) {

  ePrefix := "TestDecimal_AddToThis_02"

  nStrAry := []string{
    "35.50",
    "36.50",
    "5.5",
    "92.75",
  }

  originalNumerStr := "0"

  expectedNumberStr := "170.25"

  decSum, err := new(Decimal).NewNumStr(originalNumerStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSum, err := new(Decimal).NewNumStr(originalNumerStr)\n"+
      "originalNumerStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumerStr, err.Error())
    return
  }

  err = decSum.IsValid("Validating decSum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSum.IsValid('Validating decSum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSumNumberStr, err := decSum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSumNumberStr, err := decSum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend := Decimal{}

  var decAddendNumberStr string

  for i := 0; i < len(nStrAry); i++ {

    err = decAddend.SetNumStr(nStrAry[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decAddend.SetNumStr(nStrAry[%v])\n"+
        "nStrAry[%v]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, nStrAry[i], err.Error())
      return
    }

    err = decAddend.IsValid(fmt.Sprintf("Validating decAddend[%v]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decAddend.IsValid('Validating decAddend[%v]')\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    decAddendNumberStr, err = decAddend.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decAddendNumberStr, err := decAddend.GetNumStr()\n"+
        "index i= '%v'"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    err = decSum.AddToThis(decAddend)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decSum.AddToThis(decAddend)\n"+
        "decAddend= '%v'\n"+
        "orignal-decSum= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        decAddendNumberStr,
        decSumNumberStr,
        err.Error())

      return
    }

    err = decSum.IsValid(fmt.Sprintf("Validating decSum[%v]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decSum.IsValid('Validating decSum[%v]')\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    decSumNumberStr, err = decSum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decSumNumberStr, err = decSum[%v].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }
  }

  if expectedNumberStr != decSumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != decSumNumberStr\n"+
      "Expected decSumNumberStr = '%v'\n"+
      "  Actual decSumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decSumNumberStr)

    return
  }

  return
}

func TestDecimal_AddToThis_03(t *testing.T) {

  ePrefix := "TestDecimal_AddToThis_03"

  nStrAry := []string{
    "35.50",
    "36.50",
    "5.5",
    "92.75",
  }

  originalNumerStr := "150"

  expectedNumberStr := "320.25"

  decSum, err := new(Decimal).NewNumStr(originalNumerStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSum, err := new(Decimal).NewNumStr(originalNumerStr)\n"+
      "originalNumerStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumerStr, err.Error())
    return
  }

  err = decSum.IsValid("Validating decSum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSum.IsValid('Validating decSum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSumNumberStr, err := decSum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSumNumberStr, err := decSum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend := Decimal{}

  var decAddendNumberStr string

  for i := 0; i < len(nStrAry); i++ {

    err = decAddend.SetNumStr(nStrAry[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decAddend.SetNumStr(nStrAry[%v])\n"+
        "nStrAry[%v]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, nStrAry[i], err.Error())
      return
    }

    err = decAddend.IsValid(fmt.Sprintf("Validating decAddend[%v]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decAddend.IsValid('Validating decAddend[%v]')\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    decAddendNumberStr, err = decAddend.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decAddendNumberStr, err := decAddend.GetNumStr()\n"+
        "index i= '%v'"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    err = decSum.AddToThis(decAddend)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decSum.AddToThis(decAddend)\n"+
        "decAddend= '%v'\n"+
        "orignal-decSum= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        decAddendNumberStr,
        decSumNumberStr,
        err.Error())

      return
    }

    err = decSum.IsValid(fmt.Sprintf("Validating decSum[%v]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decSum.IsValid('Validating decSum[%v]')\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    decSumNumberStr, err = decSum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decSumNumberStr, err = decSum[%v].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }
  }

  if expectedNumberStr != decSumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != decSumNumberStr\n"+
      "Expected decSumNumberStr = '%v'\n"+
      "  Actual decSumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decSumNumberStr)

    return
  }

  return
}

func TestDecimal_AddToThisArray_01(t *testing.T) {

  ePrefix := "TestDecimal_AddToThisArray_01"

  originalNumerStr := "25.72"

  expectedPrecisionInt := 4

  expectedSignValue := 1

  expectedNumberStr := "68139.6265"

  decNumBase, err := new(Decimal).NewNumStr(originalNumerStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBase, err := new(Decimal).NewNumStr(originalNumerStr)\n"+
      "originalNumerStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumerStr, err.Error())
    return
  }

  err = decNumBase.IsValid("Validating decNumBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.IsValid('Validating decNumBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBaseNumberStr, err := decNumBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseNumberStr, err := decNumBase.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumArray, err := new(Decimal).NewNumStrsMultiple(
    "351.7",
    "6224.894",
    "34.8",
    "150",
    "150726.9",
    "-89421.6175",
    "47.23",
  )

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumArray, err := new(Decimal).NewNumStrsMultiple(... array strings)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = decNumBase.AddToThisArray(decNumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.AddToThisArray(decNumArray)\n"+
      "original - decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  err = decNumBase.IsValid("Validating decNumBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.IsValid('Validating decNumBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBaseNumberStr, err = decNumBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseNumberStr, err = decNumBase.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != decNumBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != decNumBaseNumberStr\n"+
      "Expected decNumBaseNumberStr = '%v'\n"+
      "  Actual decNumBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumBaseNumberStr)

    return
  }

  decNumBasePrecisionInt, err := decNumBase.GetPrecision()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBasePrecisionInt, err := decNumBase.GetPrecision()\n"+
      "decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  if expectedPrecisionInt != decNumBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != decNumBasePrecisionInt\n"+
      "Expected decNumBasePrecisionInt = '%v'\n"+
      "  Actual decNumBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, decNumBasePrecisionInt)

    return
  }

  decNumBaseSignValue, err := decNumBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseSignValue, err := decNumBase.GetSign()\n"+
      "decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  if expectedSignValue != decNumBaseSignValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != decNumBaseSignValue\n"+
      "Expected decNumBaseSignValue = '%v'\n"+
      "  Actual decNumBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, decNumBaseSignValue)

    return
  }

  return
}

func TestDecimal_AddToThisArray_02(t *testing.T) {

  ePrefix := "TestDecimal_AddToThisArray_02"

  originalNumerStr := "25.72"

  expectedPrecisionInt := 4

  expectedSignValue := -1

  expectedNumberStr := "-233314.1735"

  decNumBase, err := new(Decimal).NewNumStr(originalNumerStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBase, err := new(Decimal).NewNumStr(originalNumerStr)\n"+
      "originalNumerStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumerStr, err.Error())
    return
  }

  err = decNumBase.IsValid("Validating decNumBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.IsValid('Validating decNumBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBaseNumberStr, err := decNumBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseNumberStr, err := decNumBase.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumArray, err := new(Decimal).NewNumStrsMultiple(
    "351.7",
    "6224.894",
    "34.8",
    "150",
    "-150726.9",
    "-89421.6175",
    "47.23",
  )

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumArray, err := new(Decimal).NewNumStrsMultiple(... array strings)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = decNumBase.AddToThisArray(decNumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.AddToThisArray(decNumArray)\n"+
      "original - decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  err = decNumBase.IsValid("Validating decNumBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.IsValid('Validating decNumBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBaseNumberStr, err = decNumBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseNumberStr, err = decNumBase.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != decNumBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != decNumBaseNumberStr\n"+
      "Expected decNumBaseNumberStr = '%v'\n"+
      "  Actual decNumBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumBaseNumberStr)

    return
  }

  decNumBasePrecisionInt, err := decNumBase.GetPrecision()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBasePrecisionInt, err := decNumBase.GetPrecision()\n"+
      "decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  if expectedPrecisionInt != decNumBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != decNumBasePrecisionInt\n"+
      "Expected decNumBasePrecisionInt = '%v'\n"+
      "  Actual decNumBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, decNumBasePrecisionInt)

    return
  }

  decNumBaseSignValue, err := decNumBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseSignValue, err := decNumBase.GetSign()\n"+
      "decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  if expectedSignValue != decNumBaseSignValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != decNumBaseSignValue\n"+
      "Expected decNumBaseSignValue = '%v'\n"+
      "  Actual decNumBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, decNumBaseSignValue)

    return
  }

  return
}

func TestDecimal_AddToThisSeries_01(t *testing.T) {

  ePrefix := "TestDecimal_AddToThisSeries_01"

  originalNumerStr := "25.72"

  decAddend1NumberStr := "150"

  decAddend2NumberStr := "650.25"

  decAddend3NumberStr := "20.625"

  decAddend4NumberStr := "-16.59"

  expectedNumberStr := "830.005"

  expectedPrecisionInt := 3

  expectedSignValue := 1

  decNumBase, err := new(Decimal).NewNumStr(originalNumerStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBase, err := new(Decimal).NewNumStr(originalNumerStr)\n"+
      "originalNumerStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumerStr, err.Error())
    return
  }

  err = decNumBase.IsValid("Validating decNumBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.IsValid('Validating decNumBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBaseNumberStr, err := decNumBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseNumberStr, err := decNumBase.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend1, err := new(Decimal).NewNumStr(decAddend1NumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend1, err := new(Decimal).NewNumStr(decAddend1NumberStr)\n"+
      "decAddend1NumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decAddend1NumberStr, err.Error())
    return
  }

  err = decAddend1.IsValid("Validating decAddend1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend1.IsValid('Validating decAddend1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend1FinalNumStr, err := decAddend1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend1FinalNumStr, err := decAddend1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend2, err := new(Decimal).NewNumStr(decAddend2NumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend2, err := new(Decimal).NewNumStr(decAddend2NumberStr)\n"+
      "decAddend2NumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decAddend2NumberStr, err.Error())
    return
  }

  err = decAddend2.IsValid("Validating decAddend2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend2.IsValid('Validating decAddend2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend2FinalNumStr, err := decAddend2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend2FinalNumStr, err := decAddend2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend3, err := new(Decimal).NewNumStr(decAddend3NumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend3, err := new(Decimal).NewNumStr(decAddend3NumberStr)\n"+
      "decAddend3NumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decAddend3NumberStr, err.Error())
    return
  }

  err = decAddend3.IsValid("Validating decAddend3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend3.IsValid('Validating decAddend3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend3FinalNumStr, err := decAddend3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend3FinalNumStr, err := decAddend3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend4, err := new(Decimal).NewNumStr(decAddend4NumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend4, err := new(Decimal).NewNumStr(decAddend4NumberStr)\n"+
      "decAddend4NumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decAddend4NumberStr, err.Error())
    return
  }

  err = decAddend4.IsValid("Validating decAddend4")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend4.IsValid('Validating decAddend4')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend4FinalNumStr, err := decAddend4.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend4FinalNumStr, err := decAddend4.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = decNumBase.AddToThisSeries(decAddend1, decAddend2, decAddend3, decAddend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.AddToThisSeries(\n"+
      "  decAddend1, decAddend2, decAddend3, decAddend4)\n"+
      "decAddend1= '%v'\n"+
      "decAddend2= '%v'\n"+
      "decAddend3= '%v'\n"+
      "decAddend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decAddend1FinalNumStr,
      decAddend2FinalNumStr,
      decAddend3FinalNumStr,
      decAddend4FinalNumStr,
      err.Error())

    return
  }

  err = decNumBase.IsValid("Validating Final decNumBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.IsValid('Validating Final decNumBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBaseNumberStr, err = decNumBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseNumberStr, err = decNumBase.GetNumStr()\n"+
      "Final decNumBaseNumberStr Value"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBasePrecisionInt, err := decNumBase.GetPrecision()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBasePrecisionInt, err := decNumBase.GetPrecision()\n"+
      "decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  decNumBaseSignValue, err := decNumBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseSignValue, err := decNumBase.GetSign()\n"+
      "decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != decNumBaseNumberStr\n"+
      "Expected decNumBaseNumberStr = '%v'\n"+
      "  Actual decNumBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumBaseNumberStr)

    return
  }

  if expectedPrecisionInt != decNumBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != decNumBasePrecisionInt\n"+
      "Expected decNumBasePrecisionInt = '%v'\n"+
      "  Actual decNumBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, decNumBasePrecisionInt)

    return
  }

  if expectedSignValue != decNumBaseSignValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != decNumBaseSignValue\n"+
      "Expected decNumBaseSignValue = '%v'\n"+
      "  Actual decNumBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, decNumBaseSignValue)

    return
  }

  return
}

func TestDecimal_AddToThisSeries_02(t *testing.T) {

  ePrefix := "TestDecimal_AddToThisSeries_02"

  originalNumerStr := "25.72"

  decAddend1NumberStr := "150"

  decAddend2NumberStr := "-6050.25"

  decAddend3NumberStr := "20.625"

  decAddend4NumberStr := "-16.59"

  expectedNumberStr := "-5870.495"

  expectedPrecisionInt := 3

  expectedSignValue := -1

  decNumBase, err := new(Decimal).NewNumStr(originalNumerStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBase, err := new(Decimal).NewNumStr(originalNumerStr)\n"+
      "originalNumerStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumerStr, err.Error())
    return
  }

  err = decNumBase.IsValid("Validating decNumBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.IsValid('Validating decNumBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBaseNumberStr, err := decNumBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseNumberStr, err := decNumBase.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend1, err := new(Decimal).NewNumStr(decAddend1NumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend1, err := new(Decimal).NewNumStr(decAddend1NumberStr)\n"+
      "decAddend1NumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decAddend1NumberStr, err.Error())
    return
  }

  err = decAddend1.IsValid("Validating decAddend1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend1.IsValid('Validating decAddend1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend1FinalNumStr, err := decAddend1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend1FinalNumStr, err := decAddend1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend2, err := new(Decimal).NewNumStr(decAddend2NumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend2, err := new(Decimal).NewNumStr(decAddend2NumberStr)\n"+
      "decAddend2NumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decAddend2NumberStr, err.Error())
    return
  }

  err = decAddend2.IsValid("Validating decAddend2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend2.IsValid('Validating decAddend2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend2FinalNumStr, err := decAddend2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend2FinalNumStr, err := decAddend2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend3, err := new(Decimal).NewNumStr(decAddend3NumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend3, err := new(Decimal).NewNumStr(decAddend3NumberStr)\n"+
      "decAddend3NumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decAddend3NumberStr, err.Error())
    return
  }

  err = decAddend3.IsValid("Validating decAddend3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend3.IsValid('Validating decAddend3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend3FinalNumStr, err := decAddend3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend3FinalNumStr, err := decAddend3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend4, err := new(Decimal).NewNumStr(decAddend4NumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend4, err := new(Decimal).NewNumStr(decAddend4NumberStr)\n"+
      "decAddend4NumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decAddend4NumberStr, err.Error())
    return
  }

  err = decAddend4.IsValid("Validating decAddend4")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decAddend4.IsValid('Validating decAddend4')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decAddend4FinalNumStr, err := decAddend4.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decAddend4FinalNumStr, err := decAddend4.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = decNumBase.AddToThisSeries(decAddend1, decAddend2, decAddend3, decAddend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.AddToThisSeries(\n"+
      "  decAddend1, decAddend2, decAddend3, decAddend4)\n"+
      "decAddend1= '%v'\n"+
      "decAddend2= '%v'\n"+
      "decAddend3= '%v'\n"+
      "decAddend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decAddend1FinalNumStr,
      decAddend2FinalNumStr,
      decAddend3FinalNumStr,
      decAddend4FinalNumStr,
      err.Error())

    return
  }

  err = decNumBase.IsValid("Validating Final decNumBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumBase.IsValid('Validating Final decNumBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBaseNumberStr, err = decNumBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseNumberStr, err = decNumBase.GetNumStr()\n"+
      "Final decNumBaseNumberStr Value"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumBasePrecisionInt, err := decNumBase.GetPrecision()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBasePrecisionInt, err := decNumBase.GetPrecision()\n"+
      "decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  decNumBaseSignValue, err := decNumBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumBaseSignValue, err := decNumBase.GetSign()\n"+
      "decNumBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != decNumBaseNumberStr\n"+
      "Expected decNumBaseNumberStr = '%v'\n"+
      "  Actual decNumBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumBaseNumberStr)

    return
  }

  if expectedPrecisionInt != decNumBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != decNumBasePrecisionInt\n"+
      "Expected decNumBasePrecisionInt = '%v'\n"+
      "  Actual decNumBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, decNumBasePrecisionInt)

    return
  }

  if expectedSignValue != decNumBaseSignValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Number Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != decNumBaseSignValue\n"+
      "Expected decNumBaseSignValue = '%v'\n"+
      "  Actual decNumBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, decNumBaseSignValue)

    return
  }

  return
}

func TestDecimal_Cmp_01(t *testing.T) {

  ePrefix := "TestDecimal_Cmp_01"

  n1Str := "123.456"

  n2Str := "123.455"

  expectedCmpResult := 1

  decNum1, err := new(Decimal).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1, err := new(Decimal).NewNumStr(n1Str)\n"+
      "n1Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n1Str, err.Error())
    return
  }

  err = decNum1.IsValid("Validating decNum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum1.IsValid('Validating decNum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum1NumberStr, err := decNum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1NumberStr, err := decNum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2, err := new(Decimal).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2, err := new(Decimal).NewNumStr(n2Str)\n"+
      "n2Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n2Str, err.Error())
    return
  }

  err = decNum2.IsValid("Validating decNum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum2.IsValid('Validating decNum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2NumberStr, err := decNum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2NumberStr, err := decNum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  cmpResult, err := decNum1.Cmp(decNum2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "cmpResult, err := decNum1.Cmp(decNum2)\n"+
      "decNum1= '%v'\n"+
      "decNum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNum1NumberStr,
      decNum2NumberStr,
      err.Error())

    return
  }

  if expectedCmpResult != cmpResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCmpResult != cmpResult\n"+
      "Expected cmpResult = '%v'\n"+
      "  Actual cmpResult = '%v'\n\n",
      ePrefix, expectedCmpResult, cmpResult)

    return
  }

  return
}

func TestDecimal_Cmp_02(t *testing.T) {

  ePrefix := "TestDecimal_Cmp_02"

  n1Str := "123.456"

  n2Str := "123.457"

  expectedCmpResult := -1

  decNum1, err := new(Decimal).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1, err := new(Decimal).NewNumStr(n1Str)\n"+
      "n1Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n1Str, err.Error())
    return
  }

  err = decNum1.IsValid("Validating decNum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum1.IsValid('Validating decNum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum1NumberStr, err := decNum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1NumberStr, err := decNum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2, err := new(Decimal).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2, err := new(Decimal).NewNumStr(n2Str)\n"+
      "n2Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n2Str, err.Error())
    return
  }

  err = decNum2.IsValid("Validating decNum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum2.IsValid('Validating decNum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2NumberStr, err := decNum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2NumberStr, err := decNum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  cmpResult, err := decNum1.Cmp(decNum2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "cmpResult, err := decNum1.Cmp(decNum2)\n"+
      "decNum1= '%v'\n"+
      "decNum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNum1NumberStr,
      decNum2NumberStr,
      err.Error())

    return
  }

  if expectedCmpResult != cmpResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCmpResult != cmpResult\n"+
      "Expected cmpResult = '%v'\n"+
      "  Actual cmpResult = '%v'\n\n",
      ePrefix, expectedCmpResult, cmpResult)

    return
  }

  return
}

func TestDecimal_Cmp_03(t *testing.T) {

  ePrefix := "TestDecimal_Cmp_03"

  n1Str := "123.456"

  n2Str := "123.456"

  expectedCmpResult := 0

  decNum1, err := new(Decimal).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1, err := new(Decimal).NewNumStr(n1Str)\n"+
      "n1Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n1Str, err.Error())
    return
  }

  err = decNum1.IsValid("Validating decNum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum1.IsValid('Validating decNum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum1NumberStr, err := decNum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1NumberStr, err := decNum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2, err := new(Decimal).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2, err := new(Decimal).NewNumStr(n2Str)\n"+
      "n2Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n2Str, err.Error())
    return
  }

  err = decNum2.IsValid("Validating decNum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum2.IsValid('Validating decNum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2NumberStr, err := decNum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2NumberStr, err := decNum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  cmpResult, err := decNum1.Cmp(decNum2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "cmpResult, err := decNum1.Cmp(decNum2)\n"+
      "decNum1= '%v'\n"+
      "decNum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNum1NumberStr,
      decNum2NumberStr,
      err.Error())

    return
  }

  if expectedCmpResult != cmpResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCmpResult != cmpResult\n"+
      "Expected cmpResult = '%v'\n"+
      "  Actual cmpResult = '%v'\n\n",
      ePrefix, expectedCmpResult, cmpResult)

    return
  }

  return
}

func TestDecimal_Cmp_04(t *testing.T) {

  ePrefix := "TestDecimal_Cmp_04"

  n1Str := "-123.456"

  n2Str := "-123.457"

  expectedCmpResult := 1

  decNum1, err := new(Decimal).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1, err := new(Decimal).NewNumStr(n1Str)\n"+
      "n1Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n1Str, err.Error())
    return
  }

  err = decNum1.IsValid("Validating decNum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum1.IsValid('Validating decNum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum1NumberStr, err := decNum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1NumberStr, err := decNum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2, err := new(Decimal).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2, err := new(Decimal).NewNumStr(n2Str)\n"+
      "n2Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n2Str, err.Error())
    return
  }

  err = decNum2.IsValid("Validating decNum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum2.IsValid('Validating decNum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2NumberStr, err := decNum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2NumberStr, err := decNum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  cmpResult, err := decNum1.Cmp(decNum2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "cmpResult, err := decNum1.Cmp(decNum2)\n"+
      "decNum1= '%v'\n"+
      "decNum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNum1NumberStr,
      decNum2NumberStr,
      err.Error())

    return
  }

  if expectedCmpResult != cmpResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCmpResult != cmpResult\n"+
      "Expected cmpResult = '%v'\n"+
      "  Actual cmpResult = '%v'\n\n",
      ePrefix, expectedCmpResult, cmpResult)

    return
  }

  return
}

func TestDecimal_Cmp_05(t *testing.T) {

  ePrefix := "TestDecimal_Cmp_05"

  n1Str := "-123.456"

  n2Str := "-123.455"

  expectedCmpResult := -1

  decNum1, err := new(Decimal).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1, err := new(Decimal).NewNumStr(n1Str)\n"+
      "n1Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n1Str, err.Error())
    return
  }

  err = decNum1.IsValid("Validating decNum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum1.IsValid('Validating decNum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum1NumberStr, err := decNum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1NumberStr, err := decNum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2, err := new(Decimal).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2, err := new(Decimal).NewNumStr(n2Str)\n"+
      "n2Str= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, n2Str, err.Error())
    return
  }

  err = decNum2.IsValid("Validating decNum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum2.IsValid('Validating decNum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2NumberStr, err := decNum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2NumberStr, err := decNum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  cmpResult, err := decNum1.Cmp(decNum2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "cmpResult, err := decNum1.Cmp(decNum2)\n"+
      "decNum1= '%v'\n"+
      "decNum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNum1NumberStr,
      decNum2NumberStr,
      err.Error())

    return
  }

  if expectedCmpResult != cmpResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCmpResult != cmpResult\n"+
      "Expected cmpResult = '%v'\n"+
      "  Actual cmpResult = '%v'\n\n",
      ePrefix, expectedCmpResult, cmpResult)

    return
  }

  return
}

func TestDecimal_CubeRoot_01(t *testing.T) {

  ePrefix := "TestDecimal_CubeRoot_01"

  numStr1 := "2686.5"

  maxPrecisionUint := uint(30)

  expectedCubeRootNumberStr := "13.901519768959674425418091364468"

  expectedSignVal := 1

  decNum1, err := new(Decimal).NewNumStr(numStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1, err := new(Decimal).NewNumStr(numStr1)\n"+
      "numStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr1, err.Error())
    return
  }

  err = decNum1.IsValid("Validating decNum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum1.IsValid('Validating decNum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum1NumberStr, err := decNum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1NumberStr, err := decNum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decCubeRoot, err := decNum1.CubeRoot(maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRoot, err := decNum1.CubeRoot(maxPrecision)\n"+
      "decNum1= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNum1NumberStr,
      maxPrecisionUint,
      err.Error())

    return
  }

  err = decCubeRoot.IsValid("Validating decCubeRoot")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decCubeRoot.IsValid('Validating decCubeRoot')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decCubeRootNumberStr, err := decCubeRoot.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRootNumberStr, err := decCubeRoot.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedCubeRootNumberStr != decCubeRootNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCubeRoot != decCubeRootNumberStr\n"+
      "Expected decCubeRootNumberStr = '%v'\n"+
      "  Actual decCubeRootNumberStr = '%v'\n\n",
      ePrefix, expectedCubeRootNumberStr, decCubeRootNumberStr)

    return
  }

  decCubeRootSignVal, err := decCubeRoot.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRootSignVal, err := decCubeRoot.GetSign()\n"+
      "decCubeRoot= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decCubeRootNumberStr, err.Error())
    return
  }

  if expectedSignVal != decCubeRootSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != decCubeRootSignVal\n"+
      "Expected decCubeRootSignVal = '%v'\n"+
      "  Actual decCubeRootSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, decCubeRootSignVal)

    return
  }

  decCubeRootPrecisionInt, err := decCubeRoot.GetPrecision()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRootPrecisionInt, err := decCubeRoot.GetPrecision()\n"+
      "decCubeRoot= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decCubeRootNumberStr, err.Error())
    return
  }

  maxPrecisionInt := int(maxPrecisionUint)

  if maxPrecisionInt != decCubeRootPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecisionInt != decCubeRootPrecisionInt\n"+
      "Expected decCubeRootPrecisionInt = '%v'\n"+
      "  Actual decCubeRootPrecisionInt = '%v'\n\n",
      ePrefix, maxPrecisionInt, decCubeRootPrecisionInt)

    return
  }

  return
}

func TestDecimal_CubeRoot_02(t *testing.T) {

  ePrefix := "TestDecimal_CubeRoot_02"

  numStr1 := "390626"

  maxPrecisionUint := uint(29)

  expectedCubeRootNumberStr := "73.10050583431350346938081010498"

  expectedSignVal := 1

  decNum1, err := new(Decimal).NewNumStr(numStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1, err := new(Decimal).NewNumStr(numStr1)\n"+
      "numStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr1, err.Error())
    return
  }

  err = decNum1.IsValid("Validating decNum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum1.IsValid('Validating decNum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum1NumberStr, err := decNum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1NumberStr, err := decNum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decCubeRoot, err := decNum1.CubeRoot(maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRoot, err := decNum1.CubeRoot(maxPrecision)\n"+
      "decNum1= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNum1NumberStr,
      maxPrecisionUint,
      err.Error())

    return
  }

  err = decCubeRoot.IsValid("Validating decCubeRoot")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decCubeRoot.IsValid('Validating decCubeRoot')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decCubeRootNumberStr, err := decCubeRoot.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRootNumberStr, err := decCubeRoot.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedCubeRootNumberStr != decCubeRootNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCubeRoot != decCubeRootNumberStr\n"+
      "Expected decCubeRootNumberStr = '%v'\n"+
      "  Actual decCubeRootNumberStr = '%v'\n\n",
      ePrefix, expectedCubeRootNumberStr, decCubeRootNumberStr)

    return
  }

  decCubeRootSignVal, err := decCubeRoot.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRootSignVal, err := decCubeRoot.GetSign()\n"+
      "decCubeRoot= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decCubeRootNumberStr, err.Error())
    return
  }

  if expectedSignVal != decCubeRootSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != decCubeRootSignVal\n"+
      "Expected decCubeRootSignVal = '%v'\n"+
      "  Actual decCubeRootSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, decCubeRootSignVal)

    return
  }

  decCubeRootPrecisionInt, err := decCubeRoot.GetPrecision()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRootPrecisionInt, err := decCubeRoot.GetPrecision()\n"+
      "decCubeRoot= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decCubeRootNumberStr, err.Error())
    return
  }

  maxPrecisionInt := int(maxPrecisionUint)

  if maxPrecisionInt != decCubeRootPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecisionInt != decCubeRootPrecisionInt\n"+
      "Expected decCubeRootPrecisionInt = '%v'\n"+
      "  Actual decCubeRootPrecisionInt = '%v'\n\n",
      ePrefix, maxPrecisionInt, decCubeRootPrecisionInt)

    return
  }

  return
}

func TestDecimal_CubeRoot_03(t *testing.T) {

  ePrefix := "TestDecimal_CubeRoot_03"

  numStr1 := "-390626"

  maxPrecisionUint := uint(29)

  expectedCubeRootNumberStr := "-73.10050583431350346938081010498"

  expectedSignVal := -1

  decNum1, err := new(Decimal).NewNumStr(numStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1, err := new(Decimal).NewNumStr(numStr1)\n"+
      "numStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStr1, err.Error())
    return
  }

  err = decNum1.IsValid("Validating decNum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum1.IsValid('Validating decNum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum1NumberStr, err := decNum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum1NumberStr, err := decNum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decCubeRoot, err := decNum1.CubeRoot(maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRoot, err := decNum1.CubeRoot(maxPrecision)\n"+
      "decNum1= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNum1NumberStr,
      maxPrecisionUint,
      err.Error())

    return
  }

  err = decCubeRoot.IsValid("Validating decCubeRoot")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decCubeRoot.IsValid('Validating decCubeRoot')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decCubeRootNumberStr, err := decCubeRoot.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRootNumberStr, err := decCubeRoot.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedCubeRootNumberStr != decCubeRootNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCubeRoot != decCubeRootNumberStr\n"+
      "Expected decCubeRootNumberStr = '%v'\n"+
      "  Actual decCubeRootNumberStr = '%v'\n\n",
      ePrefix, expectedCubeRootNumberStr, decCubeRootNumberStr)

    return
  }

  decCubeRootSignVal, err := decCubeRoot.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRootSignVal, err := decCubeRoot.GetSign()\n"+
      "decCubeRoot= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decCubeRootNumberStr, err.Error())
    return
  }

  if expectedSignVal != decCubeRootSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != decCubeRootSignVal\n"+
      "Expected decCubeRootSignVal = '%v'\n"+
      "  Actual decCubeRootSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, decCubeRootSignVal)

    return
  }

  decCubeRootPrecisionInt, err := decCubeRoot.GetPrecision()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decCubeRootPrecisionInt, err := decCubeRoot.GetPrecision()\n"+
      "decCubeRoot= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decCubeRootNumberStr, err.Error())
    return
  }

  maxPrecisionInt := int(maxPrecisionUint)

  if maxPrecisionInt != decCubeRootPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecisionInt != decCubeRootPrecisionInt\n"+
      "Expected decCubeRootPrecisionInt = '%v'\n"+
      "  Actual decCubeRootPrecisionInt = '%v'\n\n",
      ePrefix, maxPrecisionInt, decCubeRootPrecisionInt)

    return
  }

  return
}

func TestDecimal_Divide_01(t *testing.T) {

  // quotient = dividend / divisor

  ePrefix := "TestDecimal_Divide_01"

  dividendNumberStr := "575.63"

  divisorNumberStr := "2014.123"

  expectedPrecisionUint := uint(20)

  expectedQuotientNumberStr := "0.28579684557497233287"

  expectedSignVal := 1

  decNumDividend := new(Decimal).New()

  err := decNumDividend.SetNumStr(dividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumDividend.SetNumStr(dividendNumberStr)\n"+
      "dividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendNumberStr, err.Error())
    return
  }

  err = decNumDividend.IsValid("Validating decNumDividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDividend.IsValid('Validating decNumDividend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDividendNumberStr, err := decNumDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumDividendNumberStr, err := decNumDividend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDivisor := new(Decimal).New()

  err = decNumDivisor.SetNumStr(divisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDivisor.SetNumStr(divisorNumberStr)\n"+
      "divisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, divisorNumberStr, err.Error())
    return
  }

  err = decNumDivisor.IsValid("Validating decNumDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDivisor.IsValid('Validating decNumDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDivisorNumberStr, err := decNumDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumDivisorNumberStr, err := decNumDivisor.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotient, err := decNumDividend.Divide(decNumDivisor, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotient, err := decNumDividend.\n"+
      "  Divide(decNumDivisor, ePrecisionUint)\n"+
      "decNumDividend= '%v'\n"+
      "decNumDivisor= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumDividendNumberStr,
      decNumDivisorNumberStr,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNumQuotient.IsValid("Validating decNumQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumQuotient.IsValid('Validating decNumQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotientNumberStr, err := decNumQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientNumberStr, err := decNumQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotientPrecisionUint, err := decNumQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientPrecisionUint, err := decNumQuotient.GetPrecisionUint()\n"+
      "decNumQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumQuotientNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != decNumQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != decNumQuotientPrecisionUint\n"+
      "Expected decNumQuotientPrecisionUint = '%v'\n"+
      "  Actual decNumQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumQuotientPrecisionUint)

    return
  }

  decNumQuotientSignVal, err := decNumQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientSignVal, err := decNumQuotient.GetSign()\n"+
      "decNumQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumQuotientNumberStr, err.Error())
    return
  }

  if expectedSignVal != decNumQuotientSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != decNumQuotientSignVal\n"+
      "Expected decNumQuotientSignVal = '%v'\n"+
      "  Actual decNumQuotientSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, decNumQuotientSignVal)

    return
  }

  decNumQuotientFloatStr, err := decNumQuotient.GetBigFloatString(decNumQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientFloatStr, err := decNumQuotient.\n"+
      "  GetBigFloatString(decNumQuotientPrecisionUint)\n"+
      "decNumQuotient= '%v'\n"+
      "decNumQuotientPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumQuotientNumberStr,
      decNumQuotientPrecisionUint,
      err.Error())

    return
  }

  if expectedQuotientNumberStr != decNumQuotientFloatStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != decNumQuotientFloatStr\n"+
      "Expected decNumQuotientFloatStr = '%v'\n"+
      "  Actual decNumQuotientFloatStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, decNumQuotientFloatStr)

    return
  }

  if expectedQuotientNumberStr != decNumQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != decNumQuotientNumberStr\n"+
      "Expected decNumQuotientNumberStr = '%v'\n"+
      "  Actual decNumQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, decNumQuotientNumberStr)

    return
  }

  iaDividend, err := new(IntAry).NewNumStr(dividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDividend, err := new(IntAry).NewNumStr(dividendNumberStr)\n"+
      "dividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendNumberStr, err.Error())
    return
  }

  err = iaDividend.IsValid("Validating iaDividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaDividend.IsValid('Validating iaDividend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDividendNumberStr, err := iaDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDividendNumberStr, err := iaDividend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisor, err := new(IntAry).NewNumStr(divisorNumberStr)\n"+
      "divisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, divisorNumberStr, err.Error())
    return
  }

  err = iaDivisor.IsValid("Validating iaDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaDivisor.IsValid('Validating iaDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDivisorNumberStr, err := iaDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisorNumberStr, err := iaDivisor.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaQuotient, err := iaDividend.DivideThisBy(&iaDivisor, 0, 20)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "iaDividend= '%v'\n"+
      "iaDivisor= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaDividendNumberStr,
      iaDivisorNumberStr,
      err.Error())

    return
  }

  err = iaQuotient.IsValid("Validating iaQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaQuotient.IsValid('Validating iaQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaQuotientNumberStr, err := iaQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientNumberStr, err := iaQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedQuotientNumberStr != iaQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != iaQuotientNumberStr\n"+
      "Expected iaQuotientNumberStr = '%v'\n"+
      "  Actual iaQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, iaQuotientNumberStr)

    return
  }

  iaQuotientPrecisionUint, err := iaQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientPrecisionUint, err := iaQuotient.GetPrecisionUint()\n"+
      "iaQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaQuotientNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != iaQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != iaQuotientPrecisionUint\n"+
      "Expected iaQuotientPrecisionUint = '%v'\n"+
      "  Actual iaQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, iaQuotientPrecisionUint)

    return
  }

  iaQuotientSignVal, err := iaQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientSignVal, err := iaQuotient.GetSign()\n"+
      "iaQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaQuotientNumberStr, err.Error())
    return
  }

  if expectedSignVal != iaQuotientSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != iaQuotientSignVal\n"+
      "Expected iaQuotientSignVal = '%v'\n"+
      "  Actual iaQuotientSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, iaQuotientSignVal)

    return
  }

  return
}

func TestDecimal_Divide_02(t *testing.T) {

  // quotient = dividend / divisor

  ePrefix := "TestDecimal_Divide_02"

  dividendNumberStr := "-8076.63"

  divisorNumberStr := "12014.123"

  expectedQuotientNumberStr := "−0.67226130446641839775"

  expectedPrecisionUint := uint(20)

  expectedSignVal := -1

  decNumDividend := new(Decimal).New()

  err := decNumDividend.SetNumStr(dividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumDividend.SetNumStr(dividendNumberStr)\n"+
      "dividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendNumberStr, err.Error())
    return
  }

  err = decNumDividend.IsValid("Validating decNumDividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDividend.IsValid('Validating decNumDividend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDividendNumberStr, err := decNumDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumDividendNumberStr, err := decNumDividend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDivisor := new(Decimal).New()

  err = decNumDivisor.SetNumStr(divisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDivisor.SetNumStr(divisorNumberStr)\n"+
      "divisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, divisorNumberStr, err.Error())
    return
  }

  err = decNumDivisor.IsValid("Validating decNumDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDivisor.IsValid('Validating decNumDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDivisorNumberStr, err := decNumDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumDivisorNumberStr, err := decNumDivisor.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotient, err := decNumDividend.Divide(decNumDivisor, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotient, err := decNumDividend.\n"+
      "  Divide(decNumDivisor, ePrecisionUint)\n"+
      "decNumDividend= '%v'\n"+
      "decNumDivisor= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumDividendNumberStr,
      decNumDivisorNumberStr,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNumQuotient.IsValid("Validating decNumQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumQuotient.IsValid('Validating decNumQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotientNumberStr, err := decNumQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientNumberStr, err := decNumQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotientPrecisionUint, err := decNumQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientPrecisionUint, err := decNumQuotient.GetPrecisionUint()\n"+
      "decNumQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumQuotientNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != decNumQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != decNumQuotientPrecisionUint\n"+
      "Expected decNumQuotientPrecisionUint = '%v'\n"+
      "  Actual decNumQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumQuotientPrecisionUint)

    return
  }

  decNumQuotientSignVal, err := decNumQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientSignVal, err := decNumQuotient.GetSign()\n"+
      "decNumQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumQuotientNumberStr, err.Error())
    return
  }

  if expectedSignVal != decNumQuotientSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != decNumQuotientSignVal\n"+
      "Expected decNumQuotientSignVal = '%v'\n"+
      "  Actual decNumQuotientSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, decNumQuotientSignVal)

    return
  }

  decNumQuotientFloatStr, err := decNumQuotient.GetBigFloatString(decNumQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientFloatStr, err := decNumQuotient.\n"+
      "  GetBigFloatString(decNumQuotientPrecisionUint)\n"+
      "decNumQuotient= '%v'\n"+
      "decNumQuotientPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumQuotientNumberStr,
      decNumQuotientPrecisionUint,
      err.Error())

    return
  }

  if expectedQuotientNumberStr != decNumQuotientFloatStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != decNumQuotientFloatStr\n"+
      "Expected decNumQuotientFloatStr = '%v'\n"+
      "  Actual decNumQuotientFloatStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, decNumQuotientFloatStr)

    return
  }

  if expectedQuotientNumberStr != decNumQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != decNumQuotientNumberStr\n"+
      "Expected decNumQuotientNumberStr = '%v'\n"+
      "  Actual decNumQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, decNumQuotientNumberStr)

    return
  }

  iaDividend, err := new(IntAry).NewNumStr(dividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDividend, err := new(IntAry).NewNumStr(dividendNumberStr)\n"+
      "dividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendNumberStr, err.Error())
    return
  }

  err = iaDividend.IsValid("Validating iaDividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaDividend.IsValid('Validating iaDividend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDividendNumberStr, err := iaDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDividendNumberStr, err := iaDividend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisor, err := new(IntAry).NewNumStr(divisorNumberStr)\n"+
      "divisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, divisorNumberStr, err.Error())
    return
  }

  err = iaDivisor.IsValid("Validating iaDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaDivisor.IsValid('Validating iaDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDivisorNumberStr, err := iaDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisorNumberStr, err := iaDivisor.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaQuotient, err := iaDividend.DivideThisBy(&iaDivisor, 0, 20)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "iaDividend= '%v'\n"+
      "iaDivisor= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaDividendNumberStr,
      iaDivisorNumberStr,
      err.Error())

    return
  }

  err = iaQuotient.IsValid("Validating iaQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaQuotient.IsValid('Validating iaQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaQuotientNumberStr, err := iaQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientNumberStr, err := iaQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedQuotientNumberStr != iaQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != iaQuotientNumberStr\n"+
      "Expected iaQuotientNumberStr = '%v'\n"+
      "  Actual iaQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, iaQuotientNumberStr)

    return
  }

  iaQuotientPrecisionUint, err := iaQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientPrecisionUint, err := iaQuotient.GetPrecisionUint()\n"+
      "iaQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaQuotientNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != iaQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != iaQuotientPrecisionUint\n"+
      "Expected iaQuotientPrecisionUint = '%v'\n"+
      "  Actual iaQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, iaQuotientPrecisionUint)

    return
  }

  iaQuotientSignVal, err := iaQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientSignVal, err := iaQuotient.GetSign()\n"+
      "iaQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaQuotientNumberStr, err.Error())
    return
  }

  if expectedSignVal != iaQuotientSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != iaQuotientSignVal\n"+
      "Expected iaQuotientSignVal = '%v'\n"+
      "  Actual iaQuotientSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, iaQuotientSignVal)

    return
  }

  return
}

func TestDecimal_Divide_03(t *testing.T) {

  // quotient = dividend / divisor

  ePrefix := "TestDecimal_Divide_03"

  dividendNumberStr := "100"

  divisorNumberStr := "33"

  expectedQuotientNumberStr := "3.030 303 030 303 030 303 03"

  expectedPrecisionUint := uint(20)

  expectedSignVal := 1

  decNumDividend := new(Decimal).New()

  err := decNumDividend.SetNumStr(dividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumDividend.SetNumStr(dividendNumberStr)\n"+
      "dividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendNumberStr, err.Error())
    return
  }

  err = decNumDividend.IsValid("Validating decNumDividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDividend.IsValid('Validating decNumDividend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDividendNumberStr, err := decNumDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumDividendNumberStr, err := decNumDividend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDivisor := new(Decimal).New()

  err = decNumDivisor.SetNumStr(divisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDivisor.SetNumStr(divisorNumberStr)\n"+
      "divisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, divisorNumberStr, err.Error())
    return
  }

  err = decNumDivisor.IsValid("Validating decNumDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDivisor.IsValid('Validating decNumDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDivisorNumberStr, err := decNumDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumDivisorNumberStr, err := decNumDivisor.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotient, err := decNumDividend.Divide(decNumDivisor, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotient, err := decNumDividend.\n"+
      "  Divide(decNumDivisor, ePrecisionUint)\n"+
      "decNumDividend= '%v'\n"+
      "decNumDivisor= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumDividendNumberStr,
      decNumDivisorNumberStr,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNumQuotient.IsValid("Validating decNumQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumQuotient.IsValid('Validating decNumQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotientNumberStr, err := decNumQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientNumberStr, err := decNumQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotientPrecisionUint, err := decNumQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientPrecisionUint, err := decNumQuotient.GetPrecisionUint()\n"+
      "decNumQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumQuotientNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != decNumQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != decNumQuotientPrecisionUint\n"+
      "Expected decNumQuotientPrecisionUint = '%v'\n"+
      "  Actual decNumQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumQuotientPrecisionUint)

    return
  }

  decNumQuotientSignVal, err := decNumQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientSignVal, err := decNumQuotient.GetSign()\n"+
      "decNumQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumQuotientNumberStr, err.Error())
    return
  }

  if expectedSignVal != decNumQuotientSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != decNumQuotientSignVal\n"+
      "Expected decNumQuotientSignVal = '%v'\n"+
      "  Actual decNumQuotientSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, decNumQuotientSignVal)

    return
  }

  decNumQuotientFloatStr, err := decNumQuotient.GetBigFloatString(decNumQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientFloatStr, err := decNumQuotient.\n"+
      "  GetBigFloatString(decNumQuotientPrecisionUint)\n"+
      "decNumQuotient= '%v'\n"+
      "decNumQuotientPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumQuotientNumberStr,
      decNumQuotientPrecisionUint,
      err.Error())

    return
  }

  if expectedQuotientNumberStr != decNumQuotientFloatStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != decNumQuotientFloatStr\n"+
      "Expected decNumQuotientFloatStr = '%v'\n"+
      "  Actual decNumQuotientFloatStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, decNumQuotientFloatStr)

    return
  }

  if expectedQuotientNumberStr != decNumQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != decNumQuotientNumberStr\n"+
      "Expected decNumQuotientNumberStr = '%v'\n"+
      "  Actual decNumQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, decNumQuotientNumberStr)

    return
  }

  iaDividend, err := new(IntAry).NewNumStr(dividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDividend, err := new(IntAry).NewNumStr(dividendNumberStr)\n"+
      "dividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendNumberStr, err.Error())
    return
  }

  err = iaDividend.IsValid("Validating iaDividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaDividend.IsValid('Validating iaDividend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDividendNumberStr, err := iaDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDividendNumberStr, err := iaDividend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisor, err := new(IntAry).NewNumStr(divisorNumberStr)\n"+
      "divisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, divisorNumberStr, err.Error())
    return
  }

  err = iaDivisor.IsValid("Validating iaDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaDivisor.IsValid('Validating iaDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDivisorNumberStr, err := iaDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisorNumberStr, err := iaDivisor.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaQuotient, err := iaDividend.DivideThisBy(&iaDivisor, 0, 20)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "iaDividend= '%v'\n"+
      "iaDivisor= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaDividendNumberStr,
      iaDivisorNumberStr,
      err.Error())

    return
  }

  err = iaQuotient.IsValid("Validating iaQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaQuotient.IsValid('Validating iaQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaQuotientNumberStr, err := iaQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientNumberStr, err := iaQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedQuotientNumberStr != iaQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != iaQuotientNumberStr\n"+
      "Expected iaQuotientNumberStr = '%v'\n"+
      "  Actual iaQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, iaQuotientNumberStr)

    return
  }

  iaQuotientPrecisionUint, err := iaQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientPrecisionUint, err := iaQuotient.GetPrecisionUint()\n"+
      "iaQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaQuotientNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != iaQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != iaQuotientPrecisionUint\n"+
      "Expected iaQuotientPrecisionUint = '%v'\n"+
      "  Actual iaQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, iaQuotientPrecisionUint)

    return
  }

  iaQuotientSignVal, err := iaQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientSignVal, err := iaQuotient.GetSign()\n"+
      "iaQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaQuotientNumberStr, err.Error())
    return
  }

  if expectedSignVal != iaQuotientSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != iaQuotientSignVal\n"+
      "Expected iaQuotientSignVal = '%v'\n"+
      "  Actual iaQuotientSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, iaQuotientSignVal)

    return
  }

  return
}

func TestDecimal_Divide_04(t *testing.T) {

  // quotient = dividend / divisor

  ePrefix := "TestDecimal_Divide_04"

  dividendNumberStr := "975.69"

  divisorNumberStr := "589.7654321"

  expectedQuotientNumberStr := "1.654369597"

  expectedPrecisionUint := uint(9)

  expectedSignVal := 1

  decNumDividend := new(Decimal).New()

  err := decNumDividend.SetNumStr(dividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumDividend.SetNumStr(dividendNumberStr)\n"+
      "dividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendNumberStr, err.Error())
    return
  }

  err = decNumDividend.IsValid("Validating decNumDividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDividend.IsValid('Validating decNumDividend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDividendNumberStr, err := decNumDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumDividendNumberStr, err := decNumDividend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDivisor := new(Decimal).New()

  err = decNumDivisor.SetNumStr(divisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDivisor.SetNumStr(divisorNumberStr)\n"+
      "divisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, divisorNumberStr, err.Error())
    return
  }

  err = decNumDivisor.IsValid("Validating decNumDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumDivisor.IsValid('Validating decNumDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumDivisorNumberStr, err := decNumDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumDivisorNumberStr, err := decNumDivisor.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotient, err := decNumDividend.Divide(decNumDivisor, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotient, err := decNumDividend.\n"+
      "  Divide(decNumDivisor, ePrecisionUint)\n"+
      "decNumDividend= '%v'\n"+
      "decNumDivisor= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumDividendNumberStr,
      decNumDivisorNumberStr,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNumQuotient.IsValid("Validating decNumQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumQuotient.IsValid('Validating decNumQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotientNumberStr, err := decNumQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientNumberStr, err := decNumQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumQuotientPrecisionUint, err := decNumQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientPrecisionUint, err := decNumQuotient.GetPrecisionUint()\n"+
      "decNumQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumQuotientNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != decNumQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != decNumQuotientPrecisionUint\n"+
      "Expected decNumQuotientPrecisionUint = '%v'\n"+
      "  Actual decNumQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumQuotientPrecisionUint)

    return
  }

  decNumQuotientSignVal, err := decNumQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientSignVal, err := decNumQuotient.GetSign()\n"+
      "decNumQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumQuotientNumberStr, err.Error())
    return
  }

  if expectedSignVal != decNumQuotientSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != decNumQuotientSignVal\n"+
      "Expected decNumQuotientSignVal = '%v'\n"+
      "  Actual decNumQuotientSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, decNumQuotientSignVal)

    return
  }

  decNumQuotientFloatStr, err := decNumQuotient.GetBigFloatString(decNumQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumQuotientFloatStr, err := decNumQuotient.\n"+
      "  GetBigFloatString(decNumQuotientPrecisionUint)\n"+
      "decNumQuotient= '%v'\n"+
      "decNumQuotientPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumQuotientNumberStr,
      decNumQuotientPrecisionUint,
      err.Error())

    return
  }

  if expectedQuotientNumberStr != decNumQuotientFloatStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != decNumQuotientFloatStr\n"+
      "Expected decNumQuotientFloatStr = '%v'\n"+
      "  Actual decNumQuotientFloatStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, decNumQuotientFloatStr)

    return
  }

  if expectedQuotientNumberStr != decNumQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != decNumQuotientNumberStr\n"+
      "Expected decNumQuotientNumberStr = '%v'\n"+
      "  Actual decNumQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, decNumQuotientNumberStr)

    return
  }

  iaDividend, err := new(IntAry).NewNumStr(dividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDividend, err := new(IntAry).NewNumStr(dividendNumberStr)\n"+
      "dividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendNumberStr, err.Error())
    return
  }

  err = iaDividend.IsValid("Validating iaDividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaDividend.IsValid('Validating iaDividend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDividendNumberStr, err := iaDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDividendNumberStr, err := iaDividend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisor, err := new(IntAry).NewNumStr(divisorNumberStr)\n"+
      "divisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, divisorNumberStr, err.Error())
    return
  }

  err = iaDivisor.IsValid("Validating iaDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaDivisor.IsValid('Validating iaDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaDivisorNumberStr, err := iaDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisorNumberStr, err := iaDivisor.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaQuotient, err := iaDividend.DivideThisBy(&iaDivisor, 0, 20)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "iaDividend= '%v'\n"+
      "iaDivisor= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaDividendNumberStr,
      iaDivisorNumberStr,
      err.Error())

    return
  }

  err = iaQuotient.IsValid("Validating iaQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaQuotient.IsValid('Validating iaQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaQuotientNumberStr, err := iaQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientNumberStr, err := iaQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedQuotientNumberStr != iaQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuotientNumberStr != iaQuotientNumberStr\n"+
      "Expected iaQuotientNumberStr = '%v'\n"+
      "  Actual iaQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, iaQuotientNumberStr)

    return
  }

  iaQuotientPrecisionUint, err := iaQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientPrecisionUint, err := iaQuotient.GetPrecisionUint()\n"+
      "iaQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaQuotientNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != iaQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != iaQuotientPrecisionUint\n"+
      "Expected iaQuotientPrecisionUint = '%v'\n"+
      "  Actual iaQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, iaQuotientPrecisionUint)

    return
  }

  iaQuotientSignVal, err := iaQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaQuotientSignVal, err := iaQuotient.GetSign()\n"+
      "iaQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaQuotientNumberStr, err.Error())
    return
  }

  if expectedSignVal != iaQuotientSignVal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignVal != iaQuotientSignVal\n"+
      "Expected iaQuotientSignVal = '%v'\n"+
      "  Actual iaQuotientSignVal = '%v'\n\n",
      ePrefix, expectedSignVal, iaQuotientSignVal)

    return
  }

  return
}

func TestDecimal_GetNumStr_01(t *testing.T) {

  ePrefix := "TestDecimal_GetNumStr_01"

  expectedNumberStr := "575.63"

  decNum, err := new(Decimal).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumberStr != decNumNumberStr\n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  return
}

func TestDecimal_GetNumStr_02(t *testing.T) {

  ePrefix := "TestDecimal_GetNumStr_02"

  expectedNumberStr := "-575.63"

  decNum, err := new(Decimal).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumberStr != decNumNumberStr\n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  return
}

func TestDecimal_GetNumParen_01(t *testing.T) {

  str1 := "575.63"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  actualStr := d1.GetNumParen()

  if str1 != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      str1, actualStr)
  }

}

func TestDecimal_GetNumParen_02(t *testing.T) {

  str1 := "-575.63"
  expectedStr := "(575.63)"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  actualStr := d1.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetThouStr_01(t *testing.T) {

  str1 := "2567894.63"

  expectedStr := "2,567,894.63"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  actualStr := d1.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetThouStr_02(t *testing.T) {

  str1 := "-2567894.63"

  expectedStr := "-2,567,894.63"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  actualStr := d1.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetThouParen_01(t *testing.T) {

  str1 := "2567894.63"

  expectedStr := "2,567,894.63"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  actualStr := d1.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetThouParen_02(t *testing.T) {

  str1 := "-2567894.63"

  expectedStr := "(2,567,894.63)"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  actualStr := d1.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetCurrencyStr_01(t *testing.T) {

  str1 := "2567894.63"

  expectedStr := "$2,567,894.63"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  actualStr := d1.GetCurrencyStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetCurrencyStr_02(t *testing.T) {

  str1 := "2567894.63"

  expectedStr := "€2,567,894.63"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  // '\U000020ac', // Euro €  													 7
  d1.SetCurrencySymbol('\U000020ac')

  actualStr := d1.GetCurrencyStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetCurrencyStr_03(t *testing.T) {

  str1 := "-2567894.63"

  expectedStr := "-$2,567,894.63"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  actualStr := d1.GetCurrencyStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetCurrencyParen_01(t *testing.T) {

  str1 := "2567894.63"

  expectedStr := "$2,567,894.63"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  actualStr := d1.GetCurrencyParen()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetCurrencyParen_02(t *testing.T) {

  str1 := "2567894.63"

  expectedStr := "€2,567,894.63"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  // '\U000020ac', // Euro €  													 7
  d1.SetCurrencySymbol('\U000020ac')

  actualStr := d1.GetCurrencyParen()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetCurrencyParen_03(t *testing.T) {

  str1 := "-2567894.63"

  expectedStr := "(£2,567,894.63)"

  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
      "Error='%v' ", err.Error())
  }

  // '\U000000a3', // United Kingdom Pound (£)					29
  d1.SetCurrencySymbol('\U000000a3')

  actualStr := d1.GetCurrencyParen()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
      expectedStr, actualStr)
  }

}

func TestDecimal_GetAbsoluteAllDigitsStr_01(t *testing.T) {
  str1 := "-123456.123456"
  expected := "123456123456"
  d := Decimal{}.New()
  d.SetNumStr(str1)

  absStr, err := d.GetAbsoluteAllDigitsStr()

  if err != nil {
    t.Errorf("d.GetAbsoluteAllDigitsStr() returned an Error. Error= %v", err)
  }

  if absStr != expected {
    t.Errorf("Expected %v. Instead, got %v", expected, absStr)
  }

}

func TestDecimal_GetAbsoluteAllDigitsStr_02(t *testing.T) {
  str1 := "123456.123456"
  expected := "123456123456"
  d := Decimal{}.New()
  d.SetNumStr(str1)

  absStr, err := d.GetAbsoluteAllDigitsStr()

  if err != nil {
    t.Errorf("d.GetAbsoluteAllDigitsStr() returned an Error. Error= %v", err)
  }

  if absStr != expected {
    t.Errorf("Expected %v. Instead, got %v", expected, absStr)
  }

}

func TestDecimal_GetAbsoluteValue_01(t *testing.T) {
  str1 := "-123456.123456"
  expected := "123456.123456"

  d := Decimal{}.New()

  err := d.SetNumStr(str1)

  if err != nil {
    t.Errorf("Error returned from NewNumStr(str1). str1='%v' Error= %v", str1, err)
  }

  dAbs := d.GetAbsoluteValue()

  if expected != dAbs.GetNumStr() {
    t.Errorf("Expected absolute value = '%v'. Instead, got '%v'.", expected, dAbs.GetNumStr())
  }

}

func TestDecimal_GetAbsoluteValue_02(t *testing.T) {
  str1 := "123456.123456"
  expected := "123456.123456"

  d := Decimal{}.New()

  err := d.SetNumStr(str1)

  if err != nil {
    t.Errorf("Error returned from NewNumStr(str1). str1='%v' Error= %v", str1, err)
  }

  dAbs := d.GetAbsoluteValue()

  if expected != dAbs.GetNumStr() {
    t.Errorf("Expected absolute value = '%v'. Instead, got '%v'.", expected, dAbs.GetNumStr())
  }

}

func TestDecimal_GetBigIntNum_01(t *testing.T) {

  bigI := big.NewInt(int64(123456123456))
  precision := uint(6)

  exStr := "123456.123456"

  expectedBigIntNum := BigIntNum{}.NewBigInt(bigI, precision)

  dec := Decimal{}.NewBigInt(bigI, precision)

  bigINum, err := dec.GetBigIntNum()

  if err != nil {
    t.Errorf("Error returned by dec.GetBigIntNum(). Error='%v'",
      err.Error())
  }

  if !expectedBigIntNum.Equal(bigINum) {
    t.Errorf("Error: Expected BigIntNum NOT Equal to Actual BigIntNum! "+
      "expectedBi='%v', expectedPrecision='%v'. actualBi='%v' actualPrecision='%v'",
      expectedBigIntNum.bigInt.Text(10), expectedBigIntNum.precision,
      bigINum.bigInt.Text(10), bigINum.precision)
  }

  actualNumStr := bigINum.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
      "Error='%v' ", err.Error())
  }

  if exStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'",
      exStr, actualNumStr)
  }
}

func TestDecimal_GetIntAry_01(t *testing.T) {
  bigI := big.NewInt(int64(123456123456))
  precision := uint(6)
  exStr := "123456.123456"
  d := Decimal{}.NewBigInt(bigI, precision)

  signVal := 1

  ia, err := d.GetIntAry()

  if err != nil {
    t.Errorf("Error returned from d.GetIntAryElements(). Error= %v ", err)
  }

  if exStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()== %v .  Instead ia.GetNumStr() == %v ", exStr, ia.GetNumStr())
  }

  if int(precision) != ia.GetPrecision() {
    t.Errorf("Expected ia.Precsion== %v .   Instead, ia.precision== %v", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.SignVal== %v .   Instead, ia.SignVal== %v", signVal, ia.GetSign())
  }

}

func TestDecimal_GetIntAry_02(t *testing.T) {
  bigI := big.NewInt(int64(-123456123456))
  precision := uint(6)
  exStr := "-123456.123456"
  d := Decimal{}.NewBigInt(bigI, precision)

  signVal := -1

  ia, err := d.GetIntAry()

  if err != nil {
    t.Errorf("Error returned from d.GetIntAryElements(). Error= %v ", err)
  }

  if exStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()== %v .  Instead ia.GetNumStr() == %v ", exStr, ia.GetNumStr())
  }

  if int(precision) != ia.GetPrecision() {
    t.Errorf("Expected ia.Precsion== %v .   Instead, ia.precision== %v", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.SignVal== %v .   Instead, ia.SignVal== %v", signVal, ia.GetSign())
  }

}

func TestDecimal_GetRelevantPrecision(t *testing.T) {
  str1 := "-2.0105000"
  expected := uint(4)
  d1, err := Decimal{}.NewNumStr(str1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(str1) "+
      "str1='%v' Error = '%v' ", str1, err.Error())
  }

  rP := d1.GetRelevantPrecision()

  if rP != expected {
    t.Errorf("Expected Relevant precision = %v. Instead, got %v", expected, rP)
  }

}

func TestDecimal_Inverse_01(t *testing.T) {

  numStr := "25"
  precision := uint(0)
  expected := "0.04"

  d1, err := Decimal{}.NewNumStrPrecision(numStr, precision, false)

  if err != nil {
    t.Errorf("Error Returned from Decimal d1.NewNumStrPrecision(numStrDto, precision, false). numStrDto= '%v' precision= '%v' Error= %v", numStr, precision, err)
  }

  d2, err := d1.Inverse(2)

  if err != nil {
    t.Errorf("Received error from d1.Inverse(). Error= %v", err)
  }

  if expected != d2.GetNumStr() {
    t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d2.GetNumStr())
  }

}
