package mathops

import (
  "math/big"
  "testing"
)

func TestDecimal_NewBigInt_01(t *testing.T) {

  ePrefix := "TestDecimal_NewBigInt_01"

  bigInt := big.NewInt(int64(123456123456))

  expectedPrecisionUint := uint(6)

  expectedNumberStr := "123456.123456"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewBigInt(bigInt, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewBigInt(\n"+
      "  bigInt, expectedPrecisionUint)\n"+
      "bigInt= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bigInt.Text(10),
      expectedPrecisionUint,
      err.Error())

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

  return
}

func TestDecimal_NewInt_01(t *testing.T) {

  ePrefix := "TestDecimal_NewInt_01"

  originalIntNum := 123456

  expectedPrecisionUint := uint(3)

  expectedNumberStr := "123.456"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)\n"+
      "originalIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt_02(t *testing.T) {

  ePrefix := "TestDecimal_NewInt_02"

  originalIntNum := 123456

  expectedPrecisionUint := uint(0)

  expectedNumberStr := "123456"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)\n"+
      "originalIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt_03(t *testing.T) {

  ePrefix := "TestDecimal_NewInt_03"

  originalIntNum := -123456

  expectedPrecisionUint := uint(3)

  expectedNumberStr := "-123.456"

  expectedSignVal := -1

  decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)\n"+
      "originalIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt_04(t *testing.T) {

  ePrefix := "TestDecimal_NewInt_04"

  originalIntNum := -123456

  expectedPrecisionUint := uint(0)

  expectedNumberStr := "-123456"

  expectedSignVal := -1

  decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)\n"+
      "originalIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt_05(t *testing.T) {

  ePrefix := "TestDecimal_NewInt_05"

  originalIntNum := 0

  expectedPrecisionUint := uint(0)

  expectedNumberStr := "0"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)\n"+
      "originalIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt_06(t *testing.T) {

  ePrefix := "TestDecimal_NewInt_06"

  originalIntNum := 0

  expectedPrecisionUint := uint(2)

  expectedNumberStr := "0.00"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)\n"+
      "originalIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt_07(t *testing.T) {

  ePrefix := "TestDecimal_NewInt_07"

  originalIntNum := 0

  expectedPrecisionUint := uint(4)

  expectedNumberStr := "0.0000"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt(originalIntNum, expectedPrecisionUint)\n"+
      "originalIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewIntExponent_01(t *testing.T) {

  ePrefix := "TestDecimal_NewIntExponent_01"

  originalIntNum := 123456

  originalExponent := 3

  expectedNumberStr := "123456.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)\n"+
      "originalIntNum= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewIntExponent_02(t *testing.T) {

  ePrefix := "TestDecimal_NewIntExponent_02"

  originalIntNum := 123456

  originalExponent := -3

  expectedNumberStr := "123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)\n"+
      "originalIntNum= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewIntExponent_03(t *testing.T) {

  ePrefix := "TestDecimal_NewIntExponent_03"

  originalIntNum := -123456

  originalExponent := -3

  expectedNumberStr := "-123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := -1

  decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)\n"+
      "originalIntNum= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewIntExponent_04(t *testing.T) {

  ePrefix := "TestDecimal_NewIntExponent_04"

  originalIntNum := -123456

  originalExponent := 3

  expectedNumberStr := "-123456.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := -1

  decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)\n"+
      "originalIntNum= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewIntExponent_05(t *testing.T) {

  ePrefix := "TestDecimal_NewIntExponent_05"

  originalIntNum := 0

  originalExponent := 0

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)\n"+
      "originalIntNum= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewIntExponent_06(t *testing.T) {

  ePrefix := "TestDecimal_NewIntExponent_06"

  originalIntNum := 0

  originalExponent := 3

  expectedNumberStr := "0.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewIntExponent(originalIntNum, originalExponent)\n"+
      "originalIntNum= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalIntNum, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewInt32_01(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32_01"

  originalInt32Num := int32(123456)

  expectedNumberStr := "123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)\n"+
      "originalInt32Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt32_02(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32_02"

  originalInt32Num := int32(123456)

  expectedNumberStr := "123456"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)\n"+
      "originalInt32Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt32_03(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32_03"

  originalInt32Num := int32(-123456)

  expectedNumberStr := "-123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)\n"+
      "originalInt32Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt32_04(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32_04"

  originalInt32Num := int32(-123456)

  expectedNumberStr := "-123456"

  expectedPrecisionUint := uint(0)

  expectedSignVal := -1

  decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)\n"+
      "originalInt32Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt32_05(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32_05"

  originalInt32Num := int32(0)

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)\n"+
      "originalInt32Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt32_06(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32_06"

  originalInt32Num := int32(0)

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)\n"+
      "originalInt32Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt32_07(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32_07"

  originalInt32Num := int32(0)

  expectedNumberStr := "0.0000"

  expectedPrecisionUint := uint(4)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32(originalInt32Num, expectedPrecisionUint)\n"+
      "originalInt32Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt32Exponent_01(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32Exponent_01"

  originalInt32Num := int32(123456)

  originalExponent := 3

  expectedNumberStr := "123456.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)\n"+
      "originalInt32Num= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewInt32Exponent_02(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32Exponent_01"

  originalInt32Num := int32(123456)

  originalExponent := -3

  expectedNumberStr := "123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)\n"+
      "originalInt32Num= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewInt32Exponent_03(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32Exponent_03"

  originalInt32Num := int32(-123456)

  originalExponent := -3

  expectedNumberStr := "-123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := -1

  decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)\n"+
      "originalInt32Num= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewInt32Exponent_04(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32Exponent_04"

  originalInt32Num := int32(-123456)

  originalExponent := 3

  expectedNumberStr := "-123456.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := -1

  decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)\n"+
      "originalInt32Num= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewInt32Exponent_05(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32Exponent_05"

  originalInt32Num := int32(0)

  originalExponent := 0

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)\n"+
      "originalInt32Num= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewInt32Exponent_06(t *testing.T) {

  ePrefix := "TestDecimal_NewInt32Exponent_06"

  originalInt32Num := int32(0)

  originalExponent := 3

  expectedNumberStr := "0.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt32Exponent(originalInt32Num, originalExponent)\n"+
      "originalInt32Num= '%v'\n"+
      "originalExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt32Num, originalExponent, err.Error())
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

  return
}

func TestDecimal_NewInt64_01(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64_01"

  originalInt64Num := int64(123456)

  expectedPrecisionUint := uint(3)

  expectedNumberStr := "123.456"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt64_02(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64_02"

  originalInt64Num := int64(123456)

  expectedPrecisionUint := uint(0)

  expectedNumberStr := "123456"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt64_03(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64_03"

  originalInt64Num := int64(-123456)

  expectedPrecisionUint := uint(3)

  expectedNumberStr := "-123.456"

  expectedSignVal := -1

  decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt64_04(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64_04"

  originalInt64Num := int64(-123456)

  expectedPrecisionUint := uint(0)

  expectedNumberStr := "-123456"

  expectedSignVal := -1

  decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt64_05(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64_05"

  originalInt64Num := int64(0)

  expectedPrecisionUint := uint(0)

  expectedNumberStr := "0"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt64_06(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64_06"

  originalInt64Num := int64(0)

  expectedPrecisionUint := uint(2)

  expectedNumberStr := "0.00"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt64_07(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64_07"

  originalInt64Num := int64(0)

  expectedPrecisionUint := uint(4)

  expectedNumberStr := "0.0000"

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewInt64(originalInt64Num, expectedPrecisionUint)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewInt64Exponent_01(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64Exponent_01"

  originalInt64Num := int64(123456)

  originalExponentInt := 3

  expectedNumberStr := "123456.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64Exponent(originalInt64Num, originalExponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).\n"+
      "  NewInt64Exponent(originalInt64Num, originalExponentInt)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, originalExponentInt, err.Error())
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

  return
}

func TestDecimal_NewInt64Exponent_02(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64Exponent_02"

  originalInt64Num := int64(123456)

  originalExponentInt := -3

  expectedNumberStr := "123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64Exponent(originalInt64Num, originalExponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).\n"+
      "  NewInt64Exponent(originalInt64Num, originalExponentInt)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, originalExponentInt, err.Error())
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

  return
}

func TestDecimal_NewInt64Exponent_03(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64Exponent_03"

  originalInt64Num := int64(-123456)

  originalExponentInt := -3

  expectedNumberStr := "-123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := -1

  decNum, err := new(Decimal).NewInt64Exponent(originalInt64Num, originalExponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).\n"+
      "  NewInt64Exponent(originalInt64Num, originalExponentInt)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, originalExponentInt, err.Error())
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

  return
}

func TestDecimal_NewInt64Exponent_04(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64Exponent_04"

  originalInt64Num := int64(-123456)

  originalExponentInt := 3

  expectedNumberStr := "-123456.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := -1

  decNum, err := new(Decimal).NewInt64Exponent(originalInt64Num, originalExponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).\n"+
      "  NewInt64Exponent(originalInt64Num, originalExponentInt)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, originalExponentInt, err.Error())
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

  return
}

func TestDecimal_NewInt64Exponent_05(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64Exponent_05"

  originalInt64Num := int64(0)

  originalExponentInt := 0

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64Exponent(originalInt64Num, originalExponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).\n"+
      "  NewInt64Exponent(originalInt64Num, originalExponentInt)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, originalExponentInt, err.Error())
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

  return
}

func TestDecimal_NewInt64Exponent_06(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64Exponent_06"

  originalInt64Num := int64(0)

  originalExponentInt := 3

  expectedNumberStr := "0.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64Exponent(originalInt64Num, originalExponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).\n"+
      "  NewInt64Exponent(originalInt64Num, originalExponentInt)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, originalExponentInt, err.Error())
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

  return
}

func TestDecimal_NewInt64Exponent_07(t *testing.T) {

  ePrefix := "TestDecimal_NewInt64Exponent_07"

  originalInt64Num := int64(0)

  originalExponentInt := -3

  expectedNumberStr := "0.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewInt64Exponent(originalInt64Num, originalExponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).\n"+
      "  NewInt64Exponent(originalInt64Num, originalExponentInt)\n"+
      "originalInt64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalInt64Num, originalExponentInt, err.Error())
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

  return
}

func TestDecimal_NewOne_01(t *testing.T) {

  ePrefix := "TestDecimal_NewOne_01"

  expectedNumberStr := "1.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewOne(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewOne(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewOne_02(t *testing.T) {

  ePrefix := "TestDecimal_NewOne_02"

  expectedNumberStr := "1"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewOne(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewOne(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewOne_03(t *testing.T) {

  ePrefix := "TestDecimal_NewOne_03"

  expectedNumberStr := "1.00000"

  expectedPrecisionUint := uint(5)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewOne(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewOne(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewTwo_01(t *testing.T) {

  ePrefix := "TestDecimal_NewTwo_01"

  expectedNumberStr := "2.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewTwo(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewTwo(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewTwo_02(t *testing.T) {

  ePrefix := "TestDecimal_NewTwo_02"

  expectedNumberStr := "2"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewTwo(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewTwo(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewTwo_03(t *testing.T) {

  ePrefix := "TestDecimal_NewTwo_03"

  expectedNumberStr := "2.00000"

  expectedPrecisionUint := uint(5)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewTwo(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewTwo(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewThree_01(t *testing.T) {

  ePrefix := "TestDecimal_NewThree_01"

  expectedNumberStr := "3.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewThree(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewThree(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewThree_02(t *testing.T) {

  ePrefix := "TestDecimal_NewThree_02"

  expectedNumberStr := "3"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewThree(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewThree(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewThree_03(t *testing.T) {

  ePrefix := "TestDecimal_NewThree_03"

  expectedNumberStr := "3.00000"

  expectedPrecisionUint := uint(5)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewThree(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewThree(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewFive_01(t *testing.T) {

  ePrefix := "TestDecimal_NewFive_01"

  expectedNumberStr := "5.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewFive(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewFive(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewFive_02(t *testing.T) {

  ePrefix := "TestDecimal_NewFive_02"

  expectedNumberStr := "5"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewFive(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewFive(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewFive_03(t *testing.T) {

  ePrefix := "TestDecimal_NewFive_03"

  expectedNumberStr := "5.00000"

  expectedPrecisionUint := uint(5)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewFive(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewFive(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewTen_01(t *testing.T) {

  ePrefix := "TestDecimal_NewTen_01"

  expectedNumberStr := "10.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewTen(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewTen(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewTen_02(t *testing.T) {

  ePrefix := "TestDecimal_NewTen_02"

  expectedNumberStr := "10"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewTen(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewTen(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewTen_03(t *testing.T) {

  ePrefix := "TestDecimal_NewTen_03"

  expectedNumberStr := "10.00000"

  expectedPrecisionUint := uint(5)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewTen(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewTen(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewNumStr_01(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStr_01"

  originalNumberStr := "123.456"

  expectedNumberStr := "123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewNumStr_02(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStr_02"

  originalNumberStr := "123456"

  expectedNumberStr := "123456"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewNumStr_03(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStr_03"

  originalNumberStr := "-123456"

  expectedNumberStr := "-123456"

  expectedPrecisionUint := uint(0)

  expectedSignVal := -1

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewNumStr_04(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStr_04"

  originalNumberStr := "-123.456"

  expectedNumberStr := "-123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := -1

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
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

  return
}

func TestDecimal_NewNumStrWithNumSeps_01(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrWithNumSeps_01"

  expectedNumberStr := "123,456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  decNum, err := new(Decimal).NewNumStrWithNumSeps(expectedNumberStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrWithNumSeps(\n"+
      "  expectedNumberStr, expectedNumSeps)\n"+
      "expectedNumberStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
      expectedNumSeps.String(),
      err.Error())

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

func TestDecimal_NewNumStrWithNumSeps_02(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrWithNumSeps_02"

  expectedNumberStr := "123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStrWithNumSeps(expectedNumberStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrWithNumSeps(\n"+
      "  expectedNumberStr, expectedNumSeps)\n"+
      "expectedNumberStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
      expectedNumSeps.String(),
      err.Error())

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

func TestDecimal_NewNumStrWithNumSeps_03(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrWithNumSeps_03"

  expectedNumberStr := "123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStrWithNumSeps(expectedNumberStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrWithNumSeps(\n"+
      "  expectedNumberStr, expectedNumSeps)\n"+
      "expectedNumberStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
      expectedNumSeps.String(),
      err.Error())

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

func TestDecimal_NewNumStrDto_01(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrDto_01"

  expectedNumberStr := "1.35"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != numStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != numStrDtoNumberStr \n"+
      "Expected numStrDtoNumberStr = '%v'\n"+
      "  Actual numStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoNumberStr)

    return
  }

  decNum, err := new(Decimal).NewNumStrDto(numStrDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrDto(numStrDto)\n"+
      "numStrDto= '%v'\n"+
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

func TestDecimal_NewNumStrDto_02(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrDto_02"

  expectedNumberStr := "-1.35"

  expectedPrecisionUint := uint(2)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != numStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != numStrDtoNumberStr \n"+
      "Expected numStrDtoNumberStr = '%v'\n"+
      "  Actual numStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoNumberStr)

    return
  }

  decNum, err := new(Decimal).NewNumStrDto(numStrDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrDto(numStrDto)\n"+
      "numStrDto= '%v'\n"+
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

func TestDecimal_NewNumStrDto_03(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrDto_03"

  expectedNumberStr := "0.00"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != numStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != numStrDtoNumberStr \n"+
      "Expected numStrDtoNumberStr = '%v'\n"+
      "  Actual numStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoNumberStr)

    return
  }

  decNum, err := new(Decimal).NewNumStrDto(numStrDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrDto(numStrDto)\n"+
      "numStrDto= '%v'\n"+
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

func TestDecimal_NewNumStrDto_04(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrDto_04"

  expectedNumberStr := "-0.00"

  expectedPrecisionUint := uint(2)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != numStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != numStrDtoNumberStr \n"+
      "Expected numStrDtoNumberStr = '%v'\n"+
      "  Actual numStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoNumberStr)

    return
  }

  decNum, err := new(Decimal).NewNumStrDto(numStrDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrDto(numStrDto)\n"+
      "numStrDto= '%v'\n"+
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

func TestDecimal_NewNumStrDto_05(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrDto_04"

  expectedNumberStr := "92"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != numStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != numStrDtoNumberStr \n"+
      "Expected numStrDtoNumberStr = '%v'\n"+
      "  Actual numStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoNumberStr)

    return
  }

  decNum, err := new(Decimal).NewNumStrDto(numStrDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrDto(numStrDto)\n"+
      "numStrDto= '%v'\n"+
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

func TestDecimal_NewNumStrPrecision_01(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrPrecision_01"

  originalNumberStr := "123456"

  expectedPrecisionUint := uint(3)

  expectedNumberStr := "123456.000"

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStrPrecision(originalNumberStr, expectedPrecisionUint, true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrPrecision(\n"+
      "  originalNumberStr, expectedPrecisionUint, true)\n"+
      "originalNumberStr= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewNumStrPrecision_02(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrPrecision_02"

  originalNumberStr := "0"

  expectedNumberStr := "0.00"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStrPrecision(originalNumberStr, expectedPrecisionUint, true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrPrecision(\n"+
      "  originalNumberStr, expectedPrecisionUint, true)\n"+
      "originalNumberStr= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewNumStrPrecision_03(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrPrecision_03"

  originalNumberStr := "125"

  expectedNumberStr := "125"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStrPrecision(originalNumberStr, expectedPrecisionUint, true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrPrecision(\n"+
      "  originalNumberStr, expectedPrecisionUint, true)\n"+
      "originalNumberStr= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewNumStrPrecision_04(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrPrecision_04"

  originalNumberStr := "-123.456"

  expectedNumberStr := "-123.4560"

  expectedPrecisionUint := uint(4)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStrPrecision(originalNumberStr, expectedPrecisionUint, true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrPrecision(\n"+
      "  originalNumberStr, expectedPrecisionUint, true)\n"+
      "originalNumberStr= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewNumStrPrecision_05(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrPrecision_05"

  originalNumberStr := "123.456"

  expectedNumberStr := "123.4560"

  expectedPrecisionUint := uint(4)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStrPrecision(originalNumberStr, expectedPrecisionUint, true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrPrecision(\n"+
      "  originalNumberStr, expectedPrecisionUint, true)\n"+
      "originalNumberStr= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewNumStrPrecision_06(t *testing.T) {

  ePrefix := "TestDecimal_NewNumStrPrecision_06"

  originalNumberStr := "123456"

  expectedNumberStr := "123456.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStrPrecision(originalNumberStr, expectedPrecisionUint, true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStrPrecision(\n"+
      "  originalNumberStr, expectedPrecisionUint, true)\n"+
      "originalNumberStr= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64_01(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64_01"

  originalUint64Num := uint64(123456)

  expectedPrecisionUint := uint(3)

  expectedNumberStr := "123.456"

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64(originalUint64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64(\n"+
      "  originalUint64Num, expectedPrecisionUint)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64_02(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64_02"

  originalUint64Num := uint64(123456)

  expectedPrecisionUint := uint(0)

  expectedNumberStr := "123456"

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64(originalUint64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64(\n"+
      "  originalUint64Num, expectedPrecisionUint)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64_03(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64_03"

  originalUint64Num := uint64(0)

  expectedPrecisionUint := uint(0)

  expectedNumberStr := "0"

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64(originalUint64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64(\n"+
      "  originalUint64Num, expectedPrecisionUint)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64_04(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64_04"

  originalUint64Num := uint64(0)

  expectedPrecisionUint := uint(2)

  expectedNumberStr := "0.00"

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64(originalUint64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64(\n"+
      "  originalUint64Num, expectedPrecisionUint)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64_05(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64_05"

  originalUint64Num := uint64(0)

  expectedPrecisionUint := uint(4)

  expectedNumberStr := "0.0000"

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64(originalUint64Num, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64(\n"+
      "  originalUint64Num, expectedPrecisionUint)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64Exponent_01(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64Exponent_01"

  originalUint64Num := uint64(123456)

  exponent := 3

  expectedNumberStr := "123456.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64Exponent(originalUint64Num, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64Exponent(\n"+
      "  originalUint64Num, exponent)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64Exponent_02(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64Exponent_02"

  originalUint64Num := uint64(123456)

  exponent := -3

  expectedNumberStr := "123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64Exponent(originalUint64Num, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64Exponent(\n"+
      "  originalUint64Num, exponent)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64Exponent_03(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64Exponent_03"

  originalUint64Num := uint64(0)

  exponent := 0

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64Exponent(originalUint64Num, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64Exponent(\n"+
      "  originalUint64Num, exponent)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64Exponent_04(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64Exponent_04"

  originalUint64Num := uint64(0)

  exponent := 3

  expectedNumberStr := "0.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64Exponent(originalUint64Num, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64Exponent(\n"+
      "  originalUint64Num, exponent)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NewUint64Exponent_05(t *testing.T) {

  ePrefix := "TestDecimal_NewUint64Exponent_05"

  originalUint64Num := uint64(0)

  exponent := -3

  expectedNumberStr := "0.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewUint64Exponent(originalUint64Num, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewUint64Exponent(\n"+
      "  originalUint64Num, exponent)\n"+
      "originalUint64Num= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Num,
      expectedPrecisionUint,
      err.Error())

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

func TestDecimal_NthRoot_01(t *testing.T) {
  numStr1 := "125"
  nthRootStr := "5"
  maxPrecision := uint(14)
  expected := "2.62652780440377"
  eSignVal := 1

  d1, err := Decimal{}.NewNumStr(numStr1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) "+
      "numStr1='%v' Error = '%v' ",
      numStr1, err.Error())
  }

  decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) "+
      "nthRootStr='%v' Error = '%v' ",
      nthRootStr, err.Error())
  }

  d2, err := d1.NthRoot(decNthRoot, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from d1.OriginalNthRoot(nthRoot, maxPrecision). Error= %v ", err)
  }

  if expected != d2.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
  }

  if eSignVal != d2.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d2.GetSign())
  }

  if int(maxPrecision) != d2.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
  }

}

func TestDecimal_NthRoot_02(t *testing.T) {
  numStr1 := "5604423"
  nthRootStr := "6"
  maxPrecision := uint(13)
  expected := "13.3276982415963"
  eSignVal := 1

  d1, err := Decimal{}.NewNumStr(numStr1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) "+
      "numStr1='%v' Error = '%v' ", numStr1, err.Error())
  }

  decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) "+
      "nthRootStr='%v' Error = '%v' ",
      nthRootStr, err.Error())
  }

  d2, err := d1.NthRoot(decNthRoot, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from d1.OriginalNthRoot(nthRoot, maxPrecision). Error= %v ", err)
  }

  if expected != d2.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
  }

  if eSignVal != d2.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d2.GetSign())
  }

  if int(maxPrecision) != d2.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
  }

}

func TestDecimal_NthRoot_03(t *testing.T) {
  numStr1 := "5604423.924"
  nthRootStr := "6"
  maxPrecision := uint(13)
  expected := "13.3276986078187"
  eSignVal := 1

  d1, err := Decimal{}.NewNumStr(numStr1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) "+
      "numStr1='%v' Error = '%v' ", numStr1, err.Error())
  }

  decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) "+
      "nthRootStr='%v' Error = '%v' ",
      nthRootStr, err.Error())
  }

  d2, err := d1.NthRoot(decNthRoot, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from d1.OriginalNthRoot(nthRoot, maxPrecision). Error= %v ", err)
  }

  if expected != d2.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
  }

  if eSignVal != d2.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d2.GetSign())
  }

  if int(maxPrecision) != d2.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
  }

}

func TestDecimal_NthRoot_04(t *testing.T) {
  numStr1 := "-27"
  nthRootStr := "3"
  maxPrecision := uint(2)
  expected := "-3.00"
  eSignVal := -1

  d1, err := Decimal{}.NewNumStr(numStr1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) "+
      "numStr1='%v' Error = '%v' ", numStr1, err.Error())
  }

  decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) "+
      "nthRootStr='%v' Error = '%v' ",
      nthRootStr, err.Error())
  }

  d2, err := d1.NthRoot(decNthRoot, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from d1.OriginalNthRoot(nthRoot, maxPrecision). Error= %v ", err)
  }

  if expected != d2.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
  }

  if eSignVal != d2.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d2.GetSign())
  }

  if int(maxPrecision) != d2.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
  }

}

func TestDecimal_NthRoot_05(t *testing.T) {
  numStr1 := "-27"
  nthRootStr := "4"
  maxPrecision := uint(2)

  d1, err := Decimal{}.NewNumStr(numStr1)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) "+
      "numStr1='%v' Error = '%v' ", numStr1, err.Error())
  }

  decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) "+
      "nthRootStr='%v' Error = '%v' ",
      nthRootStr, err.Error())
  }

  _, err = d1.NthRoot(decNthRoot, maxPrecision)

  if err == nil {
    t.Error("Expected Error from d1.OriginalNthRoot(nthRoot, maxPrecision) for negative number with even nthRoot. No Error triggered")
  }

}

func TestDecimal_NumStrToDecimal_01(t *testing.T) {
  d1 := Decimal{}.New()

  numStr1 := "123456"

  d2, err := d1.NumStrToDecimal(numStr1)

  if err != nil {
    t.Errorf("Received error from d1.NumStrToDecimal(numStr1). numStr1:= %v", numStr1)
  }

  expected := numStr1

  if expected != d2.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
  }

}

func TestDecimal_NumStrToDecimal_02(t *testing.T) {
  d1 := Decimal{}.New()

  numStr1 := "12345.6"

  d2, err := d1.NumStrToDecimal(numStr1)

  if err != nil {
    t.Errorf("Received error from d1.NumStrToDecimal(numStr1). numStr1:= %v", numStr1)
  }

  expected := numStr1

  if expected != d2.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
  }

}

func TestDecimal_NumStrToDecimal_03(t *testing.T) {
  d1 := Decimal{}.New()

  numStr1 := "-123456"

  d2, err := d1.NumStrToDecimal(numStr1)

  if err != nil {
    t.Errorf("Received error from d1.NumStrToDecimal(numStr1). numStr1:= %v", numStr1)
  }

  expected := numStr1

  if expected != d2.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
  }

}

func TestDecimal_NumStrToDecimal_04(t *testing.T) {
  d1 := Decimal{}.New()

  numStr1 := "-12345.6"

  d2, err := d1.NumStrToDecimal(numStr1)

  if err != nil {
    t.Errorf("Received error from d1.NumStrToDecimal(numStr1). numStr1:= %v", numStr1)
  }

  expected := numStr1

  if expected != d2.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
  }

}

func TestDecimal_NumStrPrecisionToDecimal_01(t *testing.T) {
  inStr := "123.456789"
  precision := uint(3)
  expected := "123.457"
  eSignVal := 1

  d := Decimal{}
  d1, err := d.NumStrPrecisionToDecimal(inStr, precision, true)

  if err != nil {
    t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
  }

  if expected != d1.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
  }

  if int(precision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
  }

}

func TestDecimal_NumStrPrecisionToDecimal_02(t *testing.T) {

  inStr := "123456789"
  expected := "123456789.000"
  eSignVal := 1
  precision := uint(3)

  d1, err := Decimal{}.NewPtr().NumStrPrecisionToDecimal(inStr, precision, true)

  if err != nil {
    t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
  }

  if expected != d1.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
  }

  if int(precision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
  }

}

func TestDecimal_NumStrPrecisionToDecimal_03(t *testing.T) {

  inStr := "123456789"
  expected := "123456789.000000000"
  eSignVal := 1

  precision := uint(9)
  d1, err := Decimal{}.NewPtr().NumStrPrecisionToDecimal(inStr, precision, false)

  if err != nil {
    t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
  }

  if expected != d1.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
  }

  if int(precision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
  }

}

func TestDecimal_NumStrPrecisionToDecimal_04(t *testing.T) {

  inStr := "123456789"
  expected := "123456789.0000000000"
  eSignVal := 1

  d := Decimal{}
  precision := uint(10)
  d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

  if err != nil {
    t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
  }

  if expected != d1.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
  }

  if int(precision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
  }

}

func TestDecimal_NumStrPrecisionToDecimal_05(t *testing.T) {

  inStr := "-123456789"
  expected := "-123456789.0000000000"
  eSignVal := -1

  d := Decimal{}
  precision := uint(10)
  d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

  if err != nil {
    t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
  }

  if expected != d1.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
  }

  if int(precision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
  }

}

func TestDecimal_NumStrPrecisionToDecimal_06(t *testing.T) {

  inStr := "-123456.789"
  expected := "-123456.789000"
  eSignVal := -1
  d := Decimal{}
  precision := uint(6)
  d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

  if err != nil {
    t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
  }

  if expected != d1.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
  }

  if int(precision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
  }

}

func TestDecimal_NumStrPrecisionToDecimal_07(t *testing.T) {

  inStr := "5"
  expected := "5.0"
  precision := uint(1)
  eSignVal := 1

  d := Decimal{}
  d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

  if err != nil {
    t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
  }

  if expected != d1.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
  }

  if int(precision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
  }
}

func TestDecimal_NumStrPrecisionToDecimal_08(t *testing.T) {

  inStr := "0.5"
  expected := "0.5"
  precision := uint(1)
  eSignVal := 1

  d := Decimal{}
  d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

  if err != nil {
    t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
  }

  if expected != d1.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
  }

  if int(precision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
  }
}

func TestDecimal_NumStrPrecisionToDecimal_09(t *testing.T) {

  inStr := "123456"
  expected := "123456.000"
  precision := uint(3)
  eSignVal := 1

  d := Decimal{}
  d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

  if err != nil {
    t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
  }

  if expected != d1.GetNumStr() {
    t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
  }

  if int(precision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
  }

}
