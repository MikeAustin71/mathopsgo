package mathops

import (
  "fmt"
  "testing"
)

func TestDecimal_ShiftPrecisionLeft_01(t *testing.T) {

  ePrefix := "TestDecimal_ShiftPrecisionLeft_01"

  originalNumberStr := "123456789"

  shiftLeftPlacesUint := uint(3)

  expectedNumberStr := "123456.789"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-original")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-original')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)\n"+
      "shiftLeftPlacesUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, shiftLeftPlacesUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-shiftLeftPlacesUint")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-shiftLeftPlacesUint')\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_ShiftPrecisionLeft_02(t *testing.T) {

  ePrefix := "TestDecimal_ShiftPrecisionLeft_02"

  originalNumberStr := "12345"

  shiftLeftPlacesUint := uint(6)

  expectedNumberStr := "0.012345"

  expectedPrecisionUint := uint(6)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-original")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-original')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)\n"+
      "shiftLeftPlacesUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, shiftLeftPlacesUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-shiftLeftPlacesUint")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-shiftLeftPlacesUint')\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_ShiftPrecisionLeft_03(t *testing.T) {

  ePrefix := "TestDecimal_ShiftPrecisionLeft_03"

  originalNumberStr := "12345"

  shiftLeftPlacesUint := uint(9)

  expectedNumberStr := "0.000012345"

  expectedPrecisionUint := uint(9)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-original")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-original')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)\n"+
      "shiftLeftPlacesUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, shiftLeftPlacesUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-shiftLeftPlacesUint")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-shiftLeftPlacesUint')\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_ShiftPrecisionLeft_04(t *testing.T) {

  ePrefix := "TestDecimal_ShiftPrecisionLeft_04"

  originalNumberStr := "-123456789"

  shiftLeftPlacesUint := uint(3)

  expectedNumberStr := "-123456.789"

  expectedPrecisionUint := uint(3)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-original")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-original')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)\n"+
      "shiftLeftPlacesUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, shiftLeftPlacesUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-shiftLeftPlacesUint")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-shiftLeftPlacesUint')\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_ShiftPrecisionLeft_05(t *testing.T) {

  ePrefix := "TestDecimal_ShiftPrecisionLeft_05"

  originalNumberStr := "-12345"

  shiftLeftPlacesUint := uint(6)

  expectedNumberStr := "-0.012345"

  expectedPrecisionUint := uint(6)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-original")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-original')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)\n"+
      "shiftLeftPlacesUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, shiftLeftPlacesUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-shiftLeftPlacesUint")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-shiftLeftPlacesUint')\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_ShiftPrecisionLeft_06(t *testing.T) {

  ePrefix := "TestDecimal_ShiftPrecisionLeft_06"

  originalNumberStr := "-12345"

  shiftLeftPlacesUint := uint(9)

  expectedNumberStr := "-0.000012345"

  expectedPrecisionUint := uint(9)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-original")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-original')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.ShiftPrecisionLeft(shiftLeftPlacesUint)\n"+
      "shiftLeftPlacesUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, shiftLeftPlacesUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-shiftLeftPlacesUint")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-shiftLeftPlacesUint')\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to shiftLeftPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_ShiftPrecisionRight_01(t *testing.T) {

  ePrefix := "TestDecimal_ShiftPrecisionRight_01"

  originalNumberStr := "123456789"

  shiftRightPlacesUint := uint(3)

  expectedNumberStr := "123456789000"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-originalNumberStr')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.ShiftPrecisionRight(shiftRightPlacesUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.ShiftPrecisionRight(shiftRightPlacesUint)\n"+
      "shiftRightPlacesUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, shiftRightPlacesUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-shiftRightPlacesUint")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-shiftRightPlacesUint')\n"+
      "decNum set to shiftRightPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to shiftRightPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_ShiftPrecisionRight_02(t *testing.T) {

  ePrefix := "TestDecimal_ShiftPrecisionRight_02"

  originalNumberStr := "12345"

  shiftRightPlacesUint := uint(6)

  expectedNumberStr := "12345000000"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-originalNumberStr')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.ShiftPrecisionRight(shiftRightPlacesUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.ShiftPrecisionRight(shiftRightPlacesUint)\n"+
      "shiftRightPlacesUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, shiftRightPlacesUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-shiftRightPlacesUint")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-shiftRightPlacesUint')\n"+
      "decNum set to shiftRightPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to shiftRightPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_ShiftPrecisionRight_03(t *testing.T) {

  ePrefix := "TestDecimal_ShiftPrecisionRight_03"

  originalNumberStr := "-12.345"

  shiftRightPlacesUint := uint(4)

  expectedNumberStr := "-123450"

  expectedPrecisionUint := uint(0)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-originalNumberStr')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.ShiftPrecisionRight(shiftRightPlacesUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.ShiftPrecisionRight(shiftRightPlacesUint)\n"+
      "shiftRightPlacesUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, shiftRightPlacesUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-shiftRightPlacesUint")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-shiftRightPlacesUint')\n"+
      "decNum set to shiftRightPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to shiftRightPlacesUint\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SquareRoot_01(t *testing.T) {

  ePrefix := "TestDecimal_SquareRoot_01"

  originalNumberStr := "2686.5"

  maxPrecision := uint(30)

  //                                1         2         3
  //                     0.123456789012345678901234567890
  expectedNumberStr := "51.831457629512986714934518985668"

  expectedPrecisionUint := maxPrecision

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-originalNumberStr')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  decFinalNum, err := decNum.SquareRoot(maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNum, err := decNum.SquareRoot(maxPrecision)\n"+
      "decNum= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = decFinalNum.IsValid("Validating decFinalNum-SquareRoot")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decFinalNum.IsValid('Validating decFinalNum-SquareRoot')\n"+
      "decFinalNum set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decFinalNumNumberStr, err := decFinalNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNumNumberStr, err := decFinalNum.GetNumStr()\n"+
      "decFinalNum set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decFinalNumPrecisionUint, err := decFinalNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNumPrecisionUint, err :=\n"+
      "  decFinalNum.GetPrecisionUint()\n"+
      "decFinalNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decFinalNumNumberStr, err.Error())
    return
  }

  decFinalNumSignValue, err := decFinalNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNumSignValue, err := decFinalNum.GetSign()\n"+
      "decFinalNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decFinalNumNumberStr, err.Error())
    return
  }

  decFinalNumNumSeps, err := decFinalNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNumNumSeps, err := decFinalNum.GetNumericSeparatorsDto()\n"+
      "decFinalNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decFinalNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decFinalNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decFinalNumNumberStr \n"+
      "Expected decFinalNumNumberStr = '%v'\n"+
      "  Actual decFinalNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decFinalNumNumberStr)

    return
  }

  if expectedPrecisionUint != decFinalNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decFinalNumPrecisionUint\n"+
      "Expected decFinalNumPrecisionUint = '%v'\n"+
      "  Actual decFinalNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decFinalNumPrecisionUint)

    return
  }

  if expectedSignVal != decFinalNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decFinalNumSignValue\n"+
      "Expected decFinalNumSignValue = '%v'\n"+
      "  Actual decFinalNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decFinalNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decFinalNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decFinalNumNumSeps \n"+
      "Expected decFinalNumNumSeps = '%v'\n"+
      "  Actual decFinalNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decFinalNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SquareRoot_02(t *testing.T) {

  ePrefix := "TestDecimal_SquareRoot_02"

  originalNumberStr := "390626"

  maxPrecision := uint(29)

  //                                 1         2
  //                      0.12345678901234567890123456789
  expectedNumberStr := "625.00079999948800065535895142588"

  expectedPrecisionUint := maxPrecision

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-originalNumberStr')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  decFinalNum, err := decNum.SquareRoot(maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNum, err := decNum.SquareRoot(maxPrecision)\n"+
      "decNum= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = decFinalNum.IsValid("Validating decFinalNum-SquareRoot")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decFinalNum.IsValid('Validating decFinalNum-SquareRoot')\n"+
      "decFinalNum set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decFinalNumNumberStr, err := decFinalNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNumNumberStr, err := decFinalNum.GetNumStr()\n"+
      "decFinalNum set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decFinalNumPrecisionUint, err := decFinalNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNumPrecisionUint, err :=\n"+
      "  decFinalNum.GetPrecisionUint()\n"+
      "decFinalNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decFinalNumNumberStr, err.Error())
    return
  }

  decFinalNumSignValue, err := decFinalNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNumSignValue, err := decFinalNum.GetSign()\n"+
      "decFinalNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decFinalNumNumberStr, err.Error())
    return
  }

  decFinalNumNumSeps, err := decFinalNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decFinalNumNumSeps, err := decFinalNum.GetNumericSeparatorsDto()\n"+
      "decFinalNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decFinalNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decFinalNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decFinalNumNumberStr \n"+
      "Expected decFinalNumNumberStr = '%v'\n"+
      "  Actual decFinalNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decFinalNumNumberStr)

    return
  }

  if expectedPrecisionUint != decFinalNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decFinalNumPrecisionUint\n"+
      "Expected decFinalNumPrecisionUint = '%v'\n"+
      "  Actual decFinalNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decFinalNumPrecisionUint)

    return
  }

  if expectedSignVal != decFinalNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decFinalNumSignValue\n"+
      "Expected decFinalNumSignValue = '%v'\n"+
      "  Actual decFinalNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decFinalNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decFinalNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decFinalNumNumSeps \n"+
      "Expected decFinalNumNumSeps = '%v'\n"+
      "  Actual decFinalNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decFinalNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SquareRoot_03(t *testing.T) {

  ePrefix := "TestDecimal_SquareRoot_02"

  originalNumberStr := "-390626"

  maxPrecision := uint(29)

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-originalNumberStr')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  _, err = decNum.SquareRoot(maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected Error NOT returned by:\n"+
      "_, err = decNum.SquareRoot(maxPrecision)\n"+
      "decNum= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Expected an error return. However,\n"+
      "NO ERROR WAS RETURNED!\n\n",
      ePrefix,
      decNumNumberStr,
      maxPrecision)

    return
  }

}

func TestDecimal_Subtract_01(t *testing.T) {

  ePrefix := "TestDecimal_Subtract_01"

  originalMinuendNumStr := "123456"

  originalMinuendPrecisionUint := uint(0)

  originalSubtrahendNumStr := "0.5"

  originalSubtrahendPrecisionUint := uint(1)

  expectedNumberStr := "123455.5"

  expectedPrecisionUint := uint(1)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumMinuend, err := new(Decimal).NewNumStrPrecision(originalMinuendNumStr, originalMinuendPrecisionUint, true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuend, err := new(Decimal).NewNumStrPrecision(\n"+
      "  originalMinuendNumStr, originalMinuendPrecisionUint, true)\n"+
      "originalMinuendNumStr= '%v'\n"+
      "originalMinuendPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMinuendNumStr,
      originalMinuendPrecisionUint,
      err.Error())

    return
  }

  err = decNumMinuend.IsValid("Validating decNumMinuend-originalMinuendNumStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMinuend.IsValid('Validating decNumMinuend-originalMinuendNumStr')\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumSubtrahend, err := new(Decimal).NewNumStrPrecision(originalSubtrahendNumStr, originalSubtrahendPrecisionUint, false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSubtrahend, err := new(Decimal).NewNumStrPrecision(\n"+
      "  originalSubtrahendNumStr, originalSubtrahendPrecisionUint, false)\n"+
      "originalSubtrahendNumStr= '%v'\n"+
      "originalSubtrahendPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalSubtrahendNumStr,
      originalSubtrahendPrecisionUint,
      err.Error())

    return
  }

  err = decNumSubtrahend.IsValid("Validating decNumSubtrahend-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumSubtrahend.IsValid('Validating decNumSubtrahend-originalNumberStr')\n"+
      "decNumSubtrahend set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumSubtrahendNumberStr, err := decNumSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSubtrahendNumberStr, err := decNumSubtrahend.GetNumStr()\n"+
      "decNumSubtrahend set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinal, err := decNumMinuend.Subtract(decNumSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinal, err := decNumMinuend.Subtract(decNumSubtrahend)\n"+
      "decNumMinuend= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumMinuendNumberStr,
      decNumSubtrahendNumberStr,
      err.Error())

    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal-SquareRoot")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal-SquareRoot')\n"+
      "decNumFinal set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_Subtract_02(t *testing.T) {

  ePrefix := "TestDecimal_Subtract_02"

  originalMinuendNumStr := "123.222"

  originalSubtrahendNumStr := "-1.2223"

  expectedNumberStr := "124.4443"

  expectedPrecisionUint := uint(4)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumMinuend, err := new(Decimal).NewNumStr(originalMinuendNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuend, err := new(Decimal).NewNumStr(originalMinuendNumStr)\n"+
      "originalMinuendNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMinuendNumStr,
      err.Error())

    return
  }

  err = decNumMinuend.IsValid("Validating decNumMinuend-originalMinuendNumStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMinuend.IsValid('Validating decNumMinuend-originalMinuendNumStr')\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumSubtrahend, err := new(Decimal).NewNumStr(originalSubtrahendNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSubtrahend, err := new(Decimal).NewNumStr(originalSubtrahendNumStr)\n"+
      "originalSubtrahendNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalSubtrahendNumStr, err.Error())
    return
  }

  err = decNumSubtrahend.IsValid("Validating decNumSubtrahend-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumSubtrahend.IsValid('Validating decNumSubtrahend-originalNumberStr')\n"+
      "decNumSubtrahend set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumSubtrahendNumberStr, err := decNumSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSubtrahendNumberStr, err := decNumSubtrahend.GetNumStr()\n"+
      "decNumSubtrahend set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinal, err := decNumMinuend.Subtract(decNumSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinal, err := decNumMinuend.Subtract(decNumSubtrahend)\n"+
      "decNumMinuend= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumMinuendNumberStr,
      decNumSubtrahendNumberStr,
      err.Error())

    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal-SquareRoot")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal-SquareRoot')\n"+
      "decNumFinal set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_Subtract_03(t *testing.T) {

  ePrefix := "TestDecimal_Subtract_03"

  originalMinuendNumStr := "-4"

  originalSubtrahendNumStr := "-2"

  expectedNumberStr := "-2"

  expectedPrecisionUint := uint(0)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumMinuend, err := new(Decimal).NewNumStr(originalMinuendNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuend, err := new(Decimal).NewNumStr(originalMinuendNumStr)\n"+
      "originalMinuendNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMinuendNumStr,
      err.Error())

    return
  }

  err = decNumMinuend.IsValid("Validating decNumMinuend-originalMinuendNumStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMinuend.IsValid('Validating decNumMinuend-originalMinuendNumStr')\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumSubtrahend, err := new(Decimal).NewNumStr(originalSubtrahendNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSubtrahend, err := new(Decimal).NewNumStr(originalSubtrahendNumStr)\n"+
      "originalSubtrahendNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalSubtrahendNumStr, err.Error())
    return
  }

  err = decNumSubtrahend.IsValid("Validating decNumSubtrahend-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumSubtrahend.IsValid('Validating decNumSubtrahend-originalNumberStr')\n"+
      "decNumSubtrahend set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumSubtrahendNumberStr, err := decNumSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSubtrahendNumberStr, err := decNumSubtrahend.GetNumStr()\n"+
      "decNumSubtrahend set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinal, err := decNumMinuend.Subtract(decNumSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinal, err := decNumMinuend.Subtract(decNumSubtrahend)\n"+
      "decNumMinuend= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumMinuendNumberStr,
      decNumSubtrahendNumberStr,
      err.Error())

    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal-SquareRoot")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal-SquareRoot')\n"+
      "decNumFinal set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal set to SquareRoot\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SubtractTotal_01(t *testing.T) {

  ePrefix := "TestDecimal_SubtractTotal_01"

  originalMinuendNumStr := "500.00"

  expectedNumberStr := "472.75"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nStrSubtrahendAry := []string{
    "5.50",
    "6.50",
    "7.00",
    "8.25",
  }

  decNumMinuend, err := new(Decimal).NewNumStr(originalMinuendNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuend, err := new(Decimal).NewNumStr(originalMinuendNumStr)\n"+
      "originalMinuendNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMinuendNumStr,
      err.Error())

    return
  }

  err = decNumMinuend.IsValid("Validating decNumMinuend-originalMinuendNumStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMinuend.IsValid('Validating decNumMinuend-originalMinuendNumStr')\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  var decNumSubtrahendNumberStr string

  var decNumSubtrahend Decimal

  for i := 0; i < len(nStrSubtrahendAry); i++ {

    decNumSubtrahend, err = new(Decimal).NewNumStr(nStrSubtrahendAry[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decNumSubtrahend, err := new(Decimal).NewNumStr(nStrSubtrahendAry[%v])\n"+
        "nStrSubtrahendAry[%v]= '%v'\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, nStrSubtrahendAry[i], i, err.Error())
      return
    }

    err = decNumSubtrahend.IsValid(fmt.Sprintf("Validating decNumSubtrahend-nStrSubtrahendAry[%v]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decNumSubtrahend.IsValid('Validating decNumSubtrahend-nStrSubtrahendAry[%v]')\n"+
        "decNumSubtrahend set to nStrSubtrahendAry[%v]\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, i, err.Error())
      return
    }

    decNumSubtrahendNumberStr, err = decNumSubtrahend.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decNumSubtrahendNumberStr, err := decNumSubtrahend.GetNumStr()\n"+
        "decNumSubtrahend set to nStrSubtrahendAry[%v]\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, err.Error())
      return
    }

    err = decNumMinuend.SubtractFromThis(decNumSubtrahend)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decNumMinuend.SubtractFromThis(decNumSubtrahend)\n"+
        "decNumMinuend= '%v'\n"+
        "decNumSubtrahend= '%v'\n"+
        "decNumSubtrahend set to nStrSubtrahendAry[%v]\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        decNumMinuendNumberStr,
        decNumSubtrahendNumberStr,
        i, i,
        err.Error())

      return
    }

    decNumMinuendNumberStr, err = decNumMinuend.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  }

  err = decNumMinuend.IsValid("Validating decNumMinuend-Final")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMinuend.IsValid('Validating decNumMinuend-Final')\n"+
      "decNumMinuend set to Final Value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMinuendPrecisionUint, err := decNumMinuend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendPrecisionUint, err :=\n"+
      "  decNumMinuend.GetPrecisionUint()\n"+
      "decNumMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumMinuendNumberStr, err.Error())
    return
  }

  decNumMinuendSignValue, err := decNumMinuend.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendSignValue, err := decNumMinuend.GetSign()\n"+
      "decNumMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumMinuendNumberStr, err.Error())
    return
  }

  decNumMinuendNumSeps, err := decNumMinuend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendNumSeps, err := decNumMinuend.GetNumericSeparatorsDto()\n"+
      "decNumMinuend= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumMinuendNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumMinuendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumMinuendNumberStr \n"+
      "Expected decNumMinuendNumberStr Final Value = '%v'\n"+
      "  Actual decNumMinuendNumberStr Final Value = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumMinuendNumberStr)

    return
  }

  if expectedPrecisionUint != decNumMinuendPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumMinuendPrecisionUint\n"+
      "Expected decNumMinuendPrecisionUint = '%v'\n"+
      "  Actual decNumMinuendPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumMinuendPrecisionUint)

    return
  }

  if expectedSignVal != decNumMinuendSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumMinuendSignValue\n"+
      "Expected decNumMinuendSignValue = '%v'\n"+
      "  Actual decNumMinuendSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumMinuendSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumMinuendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumMinuendNumSeps \n"+
      "Expected decNumMinuendNumSeps = '%v'\n"+
      "  Actual decNumMinuendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumMinuendNumSeps.String())

    return
  }

  return
}

func TestDecimal_SubtractTotal_02(t *testing.T) {

  ePrefix := "TestDecimal_SubtractTotal_01"

  originalMinuendNumStr := "0"

  expectedNumberStr := "-27.25"

  expectedPrecisionUint := uint(2)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nStrSubtrahendAry := []string{
    "5.50",
    "6.50",
    "7.00",
    "8.25",
  }

  decNumMinuend, err := new(Decimal).NewNumStr(originalMinuendNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuend, err := new(Decimal).NewNumStr(originalMinuendNumStr)\n"+
      "originalMinuendNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMinuendNumStr,
      err.Error())

    return
  }

  err = decNumMinuend.IsValid("Validating decNumMinuend-originalMinuendNumStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMinuend.IsValid('Validating decNumMinuend-originalMinuendNumStr')\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()\n"+
      "decNumMinuend set to originalMinuendNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  var decNumSubtrahendNumberStr string

  var decNumSubtrahend Decimal

  for i := 0; i < len(nStrSubtrahendAry); i++ {

    decNumSubtrahend, err = new(Decimal).NewNumStr(nStrSubtrahendAry[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decNumSubtrahend, err := new(Decimal).NewNumStr(nStrSubtrahendAry[%v])\n"+
        "nStrSubtrahendAry[%v]= '%v'\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, nStrSubtrahendAry[i], i, err.Error())
      return
    }

    err = decNumSubtrahend.IsValid(fmt.Sprintf("Validating decNumSubtrahend-nStrSubtrahendAry[%v]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decNumSubtrahend.IsValid('Validating decNumSubtrahend-nStrSubtrahendAry[%v]')\n"+
        "decNumSubtrahend set to nStrSubtrahendAry[%v]\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, i, err.Error())
      return
    }

    decNumSubtrahendNumberStr, err = decNumSubtrahend.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decNumSubtrahendNumberStr, err := decNumSubtrahend.GetNumStr()\n"+
        "decNumSubtrahend set to nStrSubtrahendAry[%v]\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, err.Error())
      return
    }

    err = decNumMinuend.SubtractFromThis(decNumSubtrahend)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = decNumMinuend.SubtractFromThis(decNumSubtrahend)\n"+
        "decNumMinuend= '%v'\n"+
        "decNumSubtrahend= '%v'\n"+
        "decNumSubtrahend set to nStrSubtrahendAry[%v]\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        decNumMinuendNumberStr,
        decNumSubtrahendNumberStr,
        i, i,
        err.Error())

      return
    }

    decNumMinuendNumberStr, err = decNumMinuend.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decNumMinuendNumberStr, err := decNumMinuend.GetNumStr()\n"+
        "index= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  }

  err = decNumMinuend.IsValid("Validating decNumMinuend-Final")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMinuend.IsValid('Validating decNumMinuend-Final')\n"+
      "decNumMinuend set to Final Value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMinuendPrecisionUint, err := decNumMinuend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendPrecisionUint, err :=\n"+
      "  decNumMinuend.GetPrecisionUint()\n"+
      "decNumMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumMinuendNumberStr, err.Error())
    return
  }

  decNumMinuendSignValue, err := decNumMinuend.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendSignValue, err := decNumMinuend.GetSign()\n"+
      "decNumMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumMinuendNumberStr, err.Error())
    return
  }

  decNumMinuendNumSeps, err := decNumMinuend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMinuendNumSeps, err := decNumMinuend.GetNumericSeparatorsDto()\n"+
      "decNumMinuend= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumMinuendNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumMinuendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumMinuendNumberStr \n"+
      "Expected decNumMinuendNumberStr Final Value = '%v'\n"+
      "  Actual decNumMinuendNumberStr Final Value = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumMinuendNumberStr)

    return
  }

  if expectedPrecisionUint != decNumMinuendPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumMinuendPrecisionUint\n"+
      "Expected decNumMinuendPrecisionUint = '%v'\n"+
      "  Actual decNumMinuendPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumMinuendPrecisionUint)

    return
  }

  if expectedSignVal != decNumMinuendSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumMinuendSignValue\n"+
      "Expected decNumMinuendSignValue = '%v'\n"+
      "  Actual decNumMinuendSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumMinuendSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumMinuendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumMinuendNumSeps \n"+
      "Expected decNumMinuendNumSeps = '%v'\n"+
      "  Actual decNumMinuendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumMinuendNumSeps.String())

    return
  }

  return
}

func TestDecimal_SubtractFromThisArray_01(t *testing.T) {

  nStrAry := []string{
    "5.50",
    "6.50",
    "7.00",
    "8.25",
  }

  expected := "-27.25"

  decs, err := Decimal{}.NewNumStrArray(nStrAry)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStrArray(nStrAry) "+
      "Error = '%v' ",
      err.Error())
  }

  total, err := Decimal{}.NewNumStr("0.0")

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(\"0.0\") "+
      "Error = '%v' ",
      err.Error())
  }

  err = total.SubtractFromThisArray(decs)

  if err != nil {
    t.Errorf("Error returned by total.SubtractFromThisArray(decs) "+
      "Error = '%v' ",
      err.Error())
  }

  if expected != total.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, total.GetNumStr())
  }

}

func TestDecimal_SubtractFromThisMultiple_01(t *testing.T) {

  expected := "-27.25"

  dec1, err := Decimal{}.NewNumStr("5.50")

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(\"5.50\") "+
      "Error = '%v' ",
      err.Error())
  }

  dec2, err := Decimal{}.NewNumStr("6.50")

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(\"6.50\") "+
      "Error = '%v' ",
      err.Error())
  }

  dec3, err := Decimal{}.NewNumStr("7.00")

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(\"7.00\") "+
      "Error = '%v' ",
      err.Error())
  }

  dec4, err := Decimal{}.NewNumStr("8.25")

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(\"8.25\") "+
      "Error = '%v' ",
      err.Error())
  }

  total, err := Decimal{}.NewNumStr("0.0")

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(\"0.0\") "+
      "Error = '%v' ",
      err.Error())
  }

  err = total.SubtractFromThisMultiple(
    dec1,
    dec2,
    dec3,
    dec4,
  )

  if err != nil {
    t.Errorf("Error returned by total.SubtractFromThisMultiple(...) "+
      "Error = '%v' ",
      err.Error())
  }

  if expected != total.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, total.GetNumStr())
  }

}
