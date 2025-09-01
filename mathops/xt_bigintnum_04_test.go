package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntNum_Multiply_01(t *testing.T) {

  ePrefix := "TestBigIntNum_Multiply_01"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  str1 := "575.63"

  str2 := "2014.123"

  expectedNumStr := "1159389.62249"

  expectedNumIntStr := "115938962249"

  expectedSignVal := 1

  expectedPrecisionUint := uint(5)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  bINum1 := new(BigIntNum).New()

  err = bINum1.SetNumStr(str1, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum1.SetNumStr(str1, expectedNumSeps)\n"+
      "str1= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str1,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum1.IsValid("Validating bINum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum1.IsValid('Validating bINum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum1NumberStr, err := bINum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum1NumberStr, err := bINum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum2 := new(BigIntNum).New()

  err = bINum2.SetNumStr(str2, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum2.SetNumStr(str2, expectedNumSeps)\n"+
      "str2= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str2,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum2.IsValid("Validating bINum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum2.IsValid('Validating bINum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum2NumberStr, err := bINum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum2NumberStr, err := bINum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum3, err := bINum1.Multiply(bINum2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3, err := bINum1.Multiply(bINum2)\n"+
      "bINum1= '%v'\n"+
      "bINum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bINum1NumberStr,
      bINum2NumberStr,
      err.Error())

    return
  }

  err = bINum3.IsValid("Validating bINum3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum3.IsValid('Validating bINum3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum3NumberStr, err := bINum3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumberStr, err := bINum3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != bINum3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != bINum3NumberStr \n"+
      "Expected bINum3NumberStr = '%v'\n"+
      "  Actual bINum3NumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, bINum3NumberStr)

    return
  }

  bINum3BigInt, err := bINum3.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3BigInt, err := bINum3.GetBigInt()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(bINum3BigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & bINum3 Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(bINum3BigInt) != 0\n"+
      "Expected bINum3BigInt = '%v'\n"+
      "  Actual bINum3BigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bINum3BigInt.Text(10))

    return
  }

  bINum3PrecisionUint, err := bINum3.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3PrecisionUint, err := bINum3.GetPrecisionUint()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != bINum3PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != bINum3PrecisionUint \n"+
      "Expected bINum3PrecisionUint = '%v'\n"+
      "  Actual bINum3PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bINum3PrecisionUint)

    return
  }

  bINum3SignValue, err := bINum3.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3SignValue, err := bINum3.GetSign()()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedSignVal != bINum3SignValue {
    t.Errorf("%v\n"+
      "Error: expected & bINum3SignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != bINum3SignValue\n"+
      "Expected bINum3SignValue = '%v'\n"+
      "  Actual bINum3SignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINum3SignValue)

    return
  }

  bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(bINum3NumSeps) {
    t.Errorf("%v\n"+
      "Error: bINum3 Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != bINum3NumSeps \n"+
      "Expected bINum3NumSeps = '%v'\n"+
      "  Actual bINum3NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINum3NumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Multiply_02(t *testing.T) {

  ePrefix := "TestBigIntNum_Multiply_02"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  str1 := "-575.63"

  str2 := "2014.123"

  expectedNumStr := "-1159389.62249"

  expectedNumIntStr := "-115938962249"

  expectedSignVal := -1

  expectedPrecisionUint := uint(5)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  bINum1 := new(BigIntNum).New()

  err = bINum1.SetNumStr(str1, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum1.SetNumStr(str1, expectedNumSeps)\n"+
      "str1= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str1,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum1.IsValid("Validating bINum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum1.IsValid('Validating bINum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum1NumberStr, err := bINum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum1NumberStr, err := bINum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum2 := new(BigIntNum).New()

  err = bINum2.SetNumStr(str2, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum2.SetNumStr(str2, expectedNumSeps)\n"+
      "str2= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str2,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum2.IsValid("Validating bINum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum2.IsValid('Validating bINum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum2NumberStr, err := bINum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum2NumberStr, err := bINum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum3, err := bINum1.Multiply(bINum2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3, err := bINum1.Multiply(bINum2)\n"+
      "bINum1= '%v'\n"+
      "bINum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bINum1NumberStr,
      bINum2NumberStr,
      err.Error())

    return
  }

  err = bINum3.IsValid("Validating bINum3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum3.IsValid('Validating bINum3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum3NumberStr, err := bINum3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumberStr, err := bINum3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != bINum3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != bINum3NumberStr \n"+
      "Expected bINum3NumberStr = '%v'\n"+
      "  Actual bINum3NumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, bINum3NumberStr)

    return
  }

  bINum3BigInt, err := bINum3.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3BigInt, err := bINum3.GetBigInt()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(bINum3BigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & bINum3 Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(bINum3BigInt) != 0\n"+
      "Expected bINum3BigInt = '%v'\n"+
      "  Actual bINum3BigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bINum3BigInt.Text(10))

    return
  }

  bINum3PrecisionUint, err := bINum3.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3PrecisionUint, err := bINum3.GetPrecisionUint()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != bINum3PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != bINum3PrecisionUint \n"+
      "Expected bINum3PrecisionUint = '%v'\n"+
      "  Actual bINum3PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bINum3PrecisionUint)

    return
  }

  bINum3SignValue, err := bINum3.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3SignValue, err := bINum3.GetSign()()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedSignVal != bINum3SignValue {
    t.Errorf("%v\n"+
      "Error: expected & bINum3SignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != bINum3SignValue\n"+
      "Expected bINum3SignValue = '%v'\n"+
      "  Actual bINum3SignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINum3SignValue)

    return
  }

  bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(bINum3NumSeps) {
    t.Errorf("%v\n"+
      "Error: bINum3 Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != bINum3NumSeps \n"+
      "Expected bINum3NumSeps = '%v'\n"+
      "  Actual bINum3NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINum3NumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Multiply_03(t *testing.T) {

  ePrefix := "TestBigIntNum_Multiply_03"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  str1 := "-575.63"

  str2 := "-2014.123"

  expectedNumStr := "1159389.62249"

  expectedNumIntStr := "115938962249"

  expectedSignVal := 1

  expectedPrecisionUint := uint(5)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  bINum1 := new(BigIntNum).New()

  err = bINum1.SetNumStr(str1, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum1.SetNumStr(str1, expectedNumSeps)\n"+
      "str1= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str1,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum1.IsValid("Validating bINum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum1.IsValid('Validating bINum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum1NumberStr, err := bINum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum1NumberStr, err := bINum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum2 := new(BigIntNum).New()

  err = bINum2.SetNumStr(str2, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum2.SetNumStr(str2, expectedNumSeps)\n"+
      "str2= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str2,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum2.IsValid("Validating bINum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum2.IsValid('Validating bINum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum2NumberStr, err := bINum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum2NumberStr, err := bINum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum3, err := bINum1.Multiply(bINum2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3, err := bINum1.Multiply(bINum2)\n"+
      "bINum1= '%v'\n"+
      "bINum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bINum1NumberStr,
      bINum2NumberStr,
      err.Error())

    return
  }

  err = bINum3.IsValid("Validating bINum3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum3.IsValid('Validating bINum3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum3NumberStr, err := bINum3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumberStr, err := bINum3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != bINum3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != bINum3NumberStr \n"+
      "Expected bINum3NumberStr = '%v'\n"+
      "  Actual bINum3NumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, bINum3NumberStr)

    return
  }

  bINum3BigInt, err := bINum3.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3BigInt, err := bINum3.GetBigInt()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(bINum3BigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & bINum3 Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(bINum3BigInt) != 0\n"+
      "Expected bINum3BigInt = '%v'\n"+
      "  Actual bINum3BigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bINum3BigInt.Text(10))

    return
  }

  bINum3PrecisionUint, err := bINum3.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3PrecisionUint, err := bINum3.GetPrecisionUint()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != bINum3PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != bINum3PrecisionUint \n"+
      "Expected bINum3PrecisionUint = '%v'\n"+
      "  Actual bINum3PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bINum3PrecisionUint)

    return
  }

  bINum3SignValue, err := bINum3.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3SignValue, err := bINum3.GetSign()()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedSignVal != bINum3SignValue {
    t.Errorf("%v\n"+
      "Error: expected & bINum3SignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != bINum3SignValue\n"+
      "Expected bINum3SignValue = '%v'\n"+
      "  Actual bINum3SignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINum3SignValue)

    return
  }

  bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(bINum3NumSeps) {
    t.Errorf("%v\n"+
      "Error: bINum3 Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != bINum3NumSeps \n"+
      "Expected bINum3NumSeps = '%v'\n"+
      "  Actual bINum3NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINum3NumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Multiply_04(t *testing.T) {

  ePrefix := "TestBigIntNum_Multiply_04"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  str1 := "0"

  str2 := "-2014.123"

  expectedNumStr := "0"

  expectedNumIntStr := "0"

  expectedSignVal := 1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  bINum1 := new(BigIntNum).New()

  err = bINum1.SetNumStr(str1, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum1.SetNumStr(str1, expectedNumSeps)\n"+
      "str1= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str1,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum1.IsValid("Validating bINum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum1.IsValid('Validating bINum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum1NumberStr, err := bINum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum1NumberStr, err := bINum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum2 := new(BigIntNum).New()

  err = bINum2.SetNumStr(str2, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum2.SetNumStr(str2, expectedNumSeps)\n"+
      "str2= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str2,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum2.IsValid("Validating bINum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum2.IsValid('Validating bINum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum2NumberStr, err := bINum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum2NumberStr, err := bINum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum3, err := bINum1.Multiply(bINum2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3, err := bINum1.Multiply(bINum2)\n"+
      "bINum1= '%v'\n"+
      "bINum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bINum1NumberStr,
      bINum2NumberStr,
      err.Error())

    return
  }

  err = bINum3.IsValid("Validating bINum3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum3.IsValid('Validating bINum3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum3NumberStr, err := bINum3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumberStr, err := bINum3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != bINum3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != bINum3NumberStr \n"+
      "Expected bINum3NumberStr = '%v'\n"+
      "  Actual bINum3NumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, bINum3NumberStr)

    return
  }

  bINum3BigInt, err := bINum3.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3BigInt, err := bINum3.GetBigInt()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(bINum3BigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & bINum3 Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(bINum3BigInt) != 0\n"+
      "Expected bINum3BigInt = '%v'\n"+
      "  Actual bINum3BigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bINum3BigInt.Text(10))

    return
  }

  bINum3PrecisionUint, err := bINum3.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3PrecisionUint, err := bINum3.GetPrecisionUint()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != bINum3PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != bINum3PrecisionUint \n"+
      "Expected bINum3PrecisionUint = '%v'\n"+
      "  Actual bINum3PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bINum3PrecisionUint)

    return
  }

  bINum3SignValue, err := bINum3.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3SignValue, err := bINum3.GetSign()()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedSignVal != bINum3SignValue {
    t.Errorf("%v\n"+
      "Error: expected & bINum3SignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != bINum3SignValue\n"+
      "Expected bINum3SignValue = '%v'\n"+
      "  Actual bINum3SignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINum3SignValue)

    return
  }

  bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(bINum3NumSeps) {
    t.Errorf("%v\n"+
      "Error: bINum3 Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != bINum3NumSeps \n"+
      "Expected bINum3NumSeps = '%v'\n"+
      "  Actual bINum3NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINum3NumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Multiply_05(t *testing.T) {

  ePrefix := "TestBigIntNum_Multiply_05"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  // numStr := "3"
  str1 := "3"

  str2 := "1"

  multiplyCycles := 4

  expectedNumStr := "81"

  expectedNumIntStr := "0"

  expectedSignVal := 1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  bINum1, err := new(BigIntNum).NewNumStrWithNumSeps(str1, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum1, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(str1, &expectedNumSeps)\n"+
      "str1= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str1,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum1.IsValid("Validating bINum1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum1.IsValid('Validating bINum1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum1NumberStr, err := bINum1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum1NumberStr, err := bINum1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if str1 != bINum1NumberStr {
    t.Errorf("%v\n"+
      "Error: bINum1 expected number string #1 error!\n"+
      "Because str1 != bINum1NumberStr\n"+
      "Expected bINum1NumberStr = '%v'\n"+
      "  Actual bINum1NumberStr = '%v'\n\n",
      ePrefix, str1, bINum1NumberStr)

    return
  }

  bINum2, err := new(BigIntNum).
    NewNumStrWithNumSeps(str2, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum2, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(str2, &expectedNumSeps)\n"+
      "str2= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      str2,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINum2.IsValid("Validating bINum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum2.IsValid('Validating bINum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum2NumberStr, err := bINum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum2NumberStr, err := bINum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if str2 != bINum2NumberStr {
    t.Errorf("%v\n"+
      "Error: bINum1 expected number string #2 error!\n"+
      "Because str2 != bINum2NumberStr\n"+
      "Expected bINum2NumberStr = '%v'\n"+
      "  Actual bINum2NumberStr = '%v'\n\n",
      ePrefix, str2, bINum2NumberStr)

    return
  }

  for i := 0; i < multiplyCycles; i++ {

    bINum1, err = bINum1.Multiply(bINum2)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINum1, err = bINum1.Multiply(bINum2)\n"+
        "Pre Multiply bINum1= '%v'\n"+
        "bINum2= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        bINum1NumberStr,
        bINum2NumberStr,
        err.Error())

      return
    }

    bINum1NumberStr, err = bINum1.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINum1NumberStr, err = bINum1.GetNumStr()\n"+
        "Zero Based Multiply Cycle No.= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  }

  bINum3, err := bINum1.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3, err := bINum1.GetBigIntNum()\n"+
      "bINum1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINum1NumberStr, err.Error())
    return
  }

  err = bINum3.IsValid("Validating bINum3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum3.IsValid('Validating bINum3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum3NumberStr, err := bINum3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumberStr, err := bINum3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != bINum3NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != bINum3NumberStr \n"+
      "Expected bINum3NumberStr = '%v'\n"+
      "  Actual bINum3NumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, bINum3NumberStr)

    return
  }

  bINum3BigInt, err := bINum3.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3BigInt, err := bINum3.GetBigInt()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(bINum3BigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & bINum3 Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(bINum3BigInt) != 0\n"+
      "Expected bINum3BigInt = '%v'\n"+
      "  Actual bINum3BigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bINum3BigInt.Text(10))

    return
  }

  bINum3PrecisionUint, err := bINum3.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3PrecisionUint, err := bINum3.GetPrecisionUint()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != bINum3PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != bINum3PrecisionUint \n"+
      "Expected bINum3PrecisionUint = '%v'\n"+
      "  Actual bINum3PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bINum3PrecisionUint)

    return
  }

  bINum3SignValue, err := bINum3.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3SignValue, err := bINum3.GetSign()()\n"+
      "bINum3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if expectedSignVal != bINum3SignValue {
    t.Errorf("%v\n"+
      "Error: expected & bINum3SignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != bINum3SignValue\n"+
      "Expected bINum3SignValue = '%v'\n"+
      "  Actual bINum3SignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINum3SignValue)

    return
  }

  bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum3NumSeps, err := bINum3.GetNumericSeparatorsDto()\n"+
      "bINum3= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINum3NumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(bINum3NumSeps) {
    t.Errorf("%v\n"+
      "Error: bINum3 Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != bINum3NumSeps \n"+
      "Expected bINum3NumSeps = '%v'\n"+
      "  Actual bINum3NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINum3NumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPower_01(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPower_01"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseNumStr := "85.621"

  tenExponent := uint(3)

  expectedNumStr := "85621"

  expectedNumIntStr := "85621"

  expectedSignVal := 1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPower(tenExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.MultiplyByTenToPower(tenExponent)\n"+
      "tenExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPower_02(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPower_02"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseNumStr := "85.621"

  tenExponent := uint(5)

  expectedNumStr := "8562100"

  expectedNumIntStr := "8562100"

  expectedSignVal := 1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPower(tenExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.MultiplyByTenToPower(tenExponent)\n"+
      "tenExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPower_03(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPower_03"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseNumStr := "-85.621"

  tenExponent := uint(3)

  expectedNumStr := "-85621"

  expectedNumIntStr := "-85621"

  expectedSignVal := -1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPower(tenExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.MultiplyByTenToPower(tenExponent)\n"+
      "tenExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPower_04(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPower_04"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseNumStr := "-85.621"

  tenExponent := uint(5)

  expectedNumStr := "-8562100"

  expectedNumIntStr := "-8562100"

  expectedSignVal := -1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPower(tenExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.MultiplyByTenToPower(tenExponent)\n"+
      "tenExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPower_05(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPower_05"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseNumStr := "-85.621"

  tenExponent := uint(0)

  expectedNumStr := "-85.621"

  expectedNumIntStr := "-85621"

  expectedSignVal := -1

  expectedPrecisionUint := uint(3)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPower(tenExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.MultiplyByTenToPower(tenExponent)\n"+
      "tenExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPower_06(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPower_06"

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

  baseNumStr := "-85,621"

  tenExponent := uint(0)

  expectedNumStr := "-85,621"

  expectedNumIntStr := "-85621"

  expectedSignVal := -1

  expectedPrecisionUint := uint(3)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPower(tenExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.MultiplyByTenToPower(tenExponent)\n"+
      "tenExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPowerAdd_01(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPowerAdd_01"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseNumStr := "85621"

  tenExponent := uint(1)

  addendStr := "9"

  expectedNumStr := "856219"

  expectedNumIntStr := "856219"

  expectedSignVal := 1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  bINumAddend, err := new(BigIntNum).NewNumStrWithNumSeps(addendStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAddend, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(addendStr, &expectedNumSeps)\n"+
      "addendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      addendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINumAddend.IsValid("Validating bINumAddend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINumAddend.IsValid('Validating bINumAddend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumAddendNumberStr, err := bINumAddend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAddendNumberStr, err := bINumAddend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if addendStr != bINumAddendNumberStr {
    t.Errorf("%v\n"+
      "Error: bINumAddend Number Strings ARE NOT EQUAL!\n"+
      "Because addendStr != bINumAddendNumberStr\n"+
      "Expected bINumAddendNumberStr = '%v'\n"+
      "  Actual bINumAddendNumberStr = '%v'\n\n",
      ePrefix, addendStr, bINumAddendNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPowerAdd(tenExponent, bINumAddend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.\n"+
      "  MultiplyByTenToPowerAdd(tenExponent, bINumAddend)\n"+
      "tenExponent= '%v'\n"+
      "bINumAddend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, bINumAddendNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPowerAdd_02(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPowerAdd_02"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseNumStr := "8562.1"

  tenExponent := uint(2)

  addendStr := "9"

  expectedNumStr := "856219"

  expectedNumIntStr := "856219"

  expectedSignVal := 1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  bINumAddend, err := new(BigIntNum).NewNumStrWithNumSeps(addendStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAddend, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(addendStr, &expectedNumSeps)\n"+
      "addendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      addendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINumAddend.IsValid("Validating bINumAddend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINumAddend.IsValid('Validating bINumAddend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumAddendNumberStr, err := bINumAddend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAddendNumberStr, err := bINumAddend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if addendStr != bINumAddendNumberStr {
    t.Errorf("%v\n"+
      "Error: bINumAddend Number Strings ARE NOT EQUAL!\n"+
      "Because addendStr != bINumAddendNumberStr\n"+
      "Expected bINumAddendNumberStr = '%v'\n"+
      "  Actual bINumAddendNumberStr = '%v'\n\n",
      ePrefix, addendStr, bINumAddendNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPowerAdd(tenExponent, bINumAddend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.\n"+
      "  MultiplyByTenToPowerAdd(tenExponent, bINumAddend)\n"+
      "tenExponent= '%v'\n"+
      "bINumAddend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, bINumAddendNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPowerAdd_03(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPowerAdd_03"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseNumStr := "8562.123"

  tenExponent := uint(4)

  addendStr := "9"

  expectedNumStr := "85621239"

  expectedNumIntStr := "85621239"

  expectedSignVal := 1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  bINumAddend, err := new(BigIntNum).NewNumStrWithNumSeps(addendStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAddend, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(addendStr, &expectedNumSeps)\n"+
      "addendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      addendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINumAddend.IsValid("Validating bINumAddend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINumAddend.IsValid('Validating bINumAddend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumAddendNumberStr, err := bINumAddend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAddendNumberStr, err := bINumAddend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if addendStr != bINumAddendNumberStr {
    t.Errorf("%v\n"+
      "Error: bINumAddend Number Strings ARE NOT EQUAL!\n"+
      "Because addendStr != bINumAddendNumberStr\n"+
      "Expected bINumAddendNumberStr = '%v'\n"+
      "  Actual bINumAddendNumberStr = '%v'\n\n",
      ePrefix, addendStr, bINumAddendNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPowerAdd(tenExponent, bINumAddend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.\n"+
      "  MultiplyByTenToPowerAdd(tenExponent, bINumAddend)\n"+
      "tenExponent= '%v'\n"+
      "bINumAddend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, bINumAddendNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_MultiplyByTenToPowerAdd_04(t *testing.T) {

  ePrefix := "TestBigIntNum_MultiplyByTenToPowerAdd_04"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseNumStr := "8562.123"

  tenExponent := uint(4)

  addendStr := "-10"

  expectedNumStr := "85621220"

  expectedNumIntStr := "85621220"

  expectedSignVal := 1

  expectedPrecisionUint := uint(0)

  expectedBigInt, isOk := big.NewInt(0).
    SetString(expectedNumIntStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).\n"+
      "   SetString(expectedNumIntStr, 10)\n"+
      "expectedNumIntStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedNumIntStr)
    return
  }

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
      "Error: Expected Number String Values NOT Equal\n"+
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
      "expectedBigINumBigInt, err :=\n"+
      "  expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(expectedBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(expectedBigINumBigInt) != 0\n"+
      "Expected expectedBigINumBigInt = '%v'\n"+
      "  Actual expectedBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), expectedBigINumBigInt.Text(10))

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

  expectedBigINumSignValue, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSignValue, err := expectedBigINum.GetSign()\n"+
      "expectedBigINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if expectedSignVal != expectedBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & expectedBigINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != expectedBigINum\n"+
      "Expected expectedBigINum = '%v'\n"+
      "  Actual expectedBigINum = '%v'\n\n",
      ePrefix, expectedSignVal, expectedBigINum)

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

  baseBINum, err := new(BigIntNum).NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(baseNumStr, &expectedNumSeps)\n"+
      "baseNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      baseNumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = baseBINum.IsValid("Validating baseBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseBINum.IsValid('Validating baseBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseBINumNumberStr, err := baseBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseBINumNumberStr, err := baseBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumStr != baseBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseBINum Number Strings ARE NOT EQUAL!\n"+
      "Because baseNumStr != baseBINumNumberStr\n"+
      "Expected baseBINumNumberStr = '%v'\n"+
      "  Actual baseBINumNumberStr = '%v'\n\n",
      ePrefix, baseNumStr, baseBINumNumberStr)

    return
  }

  bINumAddend, err := new(BigIntNum).NewNumStrWithNumSeps(addendStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAddend, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(addendStr, &expectedNumSeps)\n"+
      "addendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      addendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = bINumAddend.IsValid("Validating bINumAddend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINumAddend.IsValid('Validating bINumAddend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumAddendNumberStr, err := bINumAddend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAddendNumberStr, err := bINumAddend.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if addendStr != bINumAddendNumberStr {
    t.Errorf("%v\n"+
      "Error: bINumAddend Number Strings ARE NOT EQUAL!\n"+
      "Because addendStr != bINumAddendNumberStr\n"+
      "Expected bINumAddendNumberStr = '%v'\n"+
      "  Actual bINumAddendNumberStr = '%v'\n\n",
      ePrefix, addendStr, bINumAddendNumberStr)

    return
  }

  resultBINum, err := baseBINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINum, err := baseBINum.GetBigIntNum()\n"+
      "baseBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseBINumNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating base resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating base resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBINum.MultiplyByTenToPowerAdd(tenExponent, bINumAddend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.\n"+
      "  MultiplyByTenToPowerAdd(tenExponent, bINumAddend)\n"+
      "tenExponent= '%v'\n"+
      "bINumAddend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, tenExponent, bINumAddendNumberStr, err.Error())
    return
  }

  err = resultBINum.IsValid("Validating Final resultBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBINum.IsValid('Validating Final resultBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBINumNumberStr, err := resultBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumberStr, err := resultBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != resultBINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != resultBINumNumberStr \n"+
      "Expected resultBINumNumberStr = '%v'\n"+
      "  Actual resultBINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, resultBINumNumberStr)

    return
  }

  expectedBINumEqualResultBINum, err := expectedBigINum.Equal(resultBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBINumEqualResultBINum, err :=\n"+
      "  expectedBigINum.Equal(resultBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      resultBINumNumberStr,
      err.Error())

    return
  }

  if !expectedBINumEqualResultBINum {
    t.Errorf("%v\n"+
      "Error: expectedBigINum & resultBINum ARE NOT EQUAL bigIntNums!\n"+
      "Because expectedBINumEqualResultBINum == false\n"+
      "Expected expectedBINumEqualResultBINum = '%t'\n"+
      "  Actual expectedBINumEqualResultBINum = '%t'\n"+
      "expectedBigINum= '%v'\n"+
      "    resultBINum= '%v'",
      ePrefix, true, false, expectedBigINumberStr, resultBINumNumberStr)

    return
  }

  resultBINumBigInt, err := resultBINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumBigInt, err := resultBINum.GetBigInt()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & resultBINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(resultBINumBigInt) != 0\n"+
      "Expected resultBINumBigInt = '%v'\n"+
      "  Actual resultBINumBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), resultBINumBigInt.Text(10))

    return
  }

  resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumPrecisionUint, err := resultBINum.GetPrecisionUint()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != resultBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != resultBINumPrecisionUint \n"+
      "Expected resultBINumPrecisionUint = '%v'\n"+
      "  Actual resultBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBINumPrecisionUint)

    return
  }

  resultBINumSignValue, err := resultBINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumSignValue, err := resultBINum.GetSign()()\n"+
      "resultBINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != resultBINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultBINumSignValue Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != resultBINumSignValue\n"+
      "Expected resultBINumSignValue = '%v'\n"+
      "  Actual resultBINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, resultBINumSignValue)

    return
  }

  resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBINumNumSeps, err := resultBINum.GetNumericSeparatorsDto()\n"+
      "resultBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultBINumNumSeps) {
    t.Errorf("%v\n"+
      "Error: resultBINum Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultBINumNumSeps \n"+
      "Expected resultBINumNumSeps = '%v'\n"+
      "  Actual resultBINumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultBINumNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_NewBigFloat_01(t *testing.T) {

  bfloat := big.NewFloat(32.123)

  expectedNumStr := "32.123"

  bINum, err := new(BigIntNum).NewBigFloat(bfloat, 3)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewBigFloat(bfloat, 4) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewBigFloat_02(t *testing.T) {

  bFloat := big.NewFloat(float64(32.129))

  expectedNumStr := "32.13"

  bINum, err := new(BigIntNum).NewBigFloat(bFloat, 2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewBigFloat(bFloat, 4) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewBigFloat_03(t *testing.T) {

  bFloat := big.NewFloat(-32.129)

  expectedNumStr := "-32.13"

  bINum, err := new(BigIntNum).NewBigFloat(bFloat, 2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewBigFloat(bFloat, 4) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewFloat32_01(t *testing.T) {

  numf32 := float32(32.123)

  expectedNumStr := "32.123"

  maxPrecision := uint(4)

  bINum, err := new(BigIntNum).NewFloat32(numf32, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewFloat32(numf32, 4) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewFloat32_02(t *testing.T) {

  numf32 := float32(32.129)

  expectedNumStr := "32.13"

  bINum, err := new(BigIntNum).NewFloat32(numf32, 2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewFloat32(numf32, 4) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewFloat32_03(t *testing.T) {

  numf32 := float32(-32.129)

  expectedNumStr := "-32.13"

  bINum, err := new(BigIntNum).NewFloat32(numf32, 2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewFloat32(numf32, 4) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewFloat64_01(t *testing.T) {

  numf32 := float64(32.123)

  expectedNumStr := "32.123"

  maxPrecision := uint(3)

  bINum, err := new(BigIntNum).NewFloat64(numf32, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewFloat64(numf32, 4) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewFloat64_02(t *testing.T) {

  numf32 := float64(32.129)

  expectedNumStr := "32.13"

  bINum, err := new(BigIntNum).NewFloat64(numf32, 2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewFloat64(numf32, 4) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewFloat64_03(t *testing.T) {

  numf32 := float64(-32.129)

  expectedNumStr := "-32.13"

  bINum, err := new(BigIntNum).NewFloat64(numf32, 2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewFloat64(numf32, 4) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt_01(t *testing.T) {

  numInt := int(1234)
  precision := uint(3)
  expectedNumStr := "1.234"

  bINum := new(BigIntNum).NewInt(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt_02(t *testing.T) {

  numInt := int(1234)
  precision := uint(0)
  expectedNumStr := "1234"

  bINum := new(BigIntNum).NewInt(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt_03(t *testing.T) {

  numInt := int(-1234)
  precision := uint(3)
  expectedNumStr := "-1.234"

  bINum := new(BigIntNum).NewInt(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt_04(t *testing.T) {

  numInt := int(-1234)
  precision := uint(0)
  expectedNumStr := "-1234"

  bINum := new(BigIntNum).NewInt(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt_05(t *testing.T) {

  numInt := int(0)
  precision := uint(0)
  expectedNumStr := "0"

  bINum := new(BigIntNum).NewInt(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt_06(t *testing.T) {

  numInt := int(0)
  precision := uint(3)
  expectedNumStr := "0.000"

  bINum := new(BigIntNum).NewInt(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewIntExponent_01(t *testing.T) {

  numInt := 1234

  expectedNumStr := "1234.000"

  bINum := new(BigIntNum).NewIntExponent(numInt, 3)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewIntExponent_02(t *testing.T) {

  numInt := 123456

  expectedNumStr := "1234.56"

  bINum := new(BigIntNum).NewIntExponent(numInt, -2)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewIntExponent_03(t *testing.T) {

  numInt := 123456
  exponent := 0
  expectedNumStr := "123456"

  bINum := new(BigIntNum).NewIntExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewIntExponent_04(t *testing.T) {

  numInt := 0
  exponent := 0
  expectedNumStr := "0"

  bINum := new(BigIntNum).NewIntExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewIntExponent_05(t *testing.T) {

  numInt := 0
  exponent := 3
  expectedNumStr := "0.000"

  bINum := new(BigIntNum).NewIntExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewIntExponent_06(t *testing.T) {

  numInt := 0
  exponent := -3
  expectedNumStr := "0.000"

  bINum := new(BigIntNum).NewIntExponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32_01(t *testing.T) {

  numInt := int32(1234)
  precision := uint(3)
  expectedNumStr := "1.234"

  bINum := new(BigIntNum).NewInt32(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32_02(t *testing.T) {

  numInt := int32(1234)
  precision := uint(0)
  expectedNumStr := "1234"

  bINum := new(BigIntNum).NewInt32(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32_03(t *testing.T) {

  numInt := int32(-1234)
  precision := uint(3)
  expectedNumStr := "-1.234"

  bINum := new(BigIntNum).NewInt32(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32_04(t *testing.T) {

  numInt := int32(-1234)
  precision := uint(0)
  expectedNumStr := "-1234"

  bINum := new(BigIntNum).NewInt32(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32_05(t *testing.T) {

  numInt := int32(0)
  precision := uint(0)
  expectedNumStr := "0"

  bINum := new(BigIntNum).NewInt32(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32_06(t *testing.T) {

  numInt := int32(0)
  precision := uint(3)
  expectedNumStr := "0.000"

  bINum := new(BigIntNum).NewInt32(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32Exponent_01(t *testing.T) {

  numInt := int32(1234)

  expectedNumStr := "1234.000"

  bINum := new(BigIntNum).NewInt32Exponent(numInt, 3)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32Exponent_02(t *testing.T) {

  numInt := int32(123456)

  expectedNumStr := "1234.56"

  bINum := new(BigIntNum).NewInt32Exponent(numInt, -2)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32Exponent_03(t *testing.T) {

  numInt := int32(123456)

  expectedNumStr := "123456"

  bINum := new(BigIntNum).NewInt32Exponent(numInt, 0)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32Exponent_04(t *testing.T) {

  numInt := int32(0)
  exponent := 0
  expectedNumStr := "0"

  bINum := new(BigIntNum).NewInt32Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32Exponent_05(t *testing.T) {

  numInt := int32(0)
  exponent := 3
  expectedNumStr := "0.000"

  bINum := new(BigIntNum).NewInt32Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt32Exponent_06(t *testing.T) {

  numInt := int32(0)
  exponent := -3
  expectedNumStr := "0.000"

  bINum := new(BigIntNum).NewInt32Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64_01(t *testing.T) {

  numInt := int64(1234)
  precision := uint(3)
  expectedNumStr := "1.234"

  bINum := new(BigIntNum).NewInt64(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64_02(t *testing.T) {

  numInt := int64(1234)
  precision := uint(0)
  expectedNumStr := "1234"

  bINum := new(BigIntNum).NewInt64(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64_03(t *testing.T) {

  numInt := int64(-1234)
  precision := uint(3)
  expectedNumStr := "-1.234"

  bINum := new(BigIntNum).NewInt64(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64_04(t *testing.T) {

  numInt := int64(-1234)
  precision := uint(0)
  expectedNumStr := "-1234"

  bINum := new(BigIntNum).NewInt64(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64_05(t *testing.T) {

  numInt := int64(0)
  precision := uint(0)
  expectedNumStr := "0"

  bINum := new(BigIntNum).NewInt64(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64_06(t *testing.T) {

  numInt := int64(0)
  precision := uint(3)
  expectedNumStr := "0.000"

  bINum := new(BigIntNum).NewInt64(numInt, precision)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64Exponent_01(t *testing.T) {

  numInt := int64(1234)

  expectedNumStr := "1234.000"

  bINum := new(BigIntNum).NewInt64Exponent(numInt, 3)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64Exponent_02(t *testing.T) {

  numInt := int64(123456)

  expectedNumStr := "1234.56"

  bINum := new(BigIntNum).NewInt64Exponent(numInt, -2)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64Exponent_03(t *testing.T) {

  numInt := int64(123456)

  expectedNumStr := "123456"

  bINum := new(BigIntNum).NewInt64Exponent(numInt, 0)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64Exponent_04(t *testing.T) {

  numInt := int64(-123456)

  expectedNumStr := "-123456"

  exponent := 0

  bINum := new(BigIntNum).NewInt64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64Exponent_05(t *testing.T) {

  numInt := int64(-123456)

  expectedNumStr := "-123.456"

  exponent := -3

  bINum := new(BigIntNum).NewInt64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64Exponent_06(t *testing.T) {

  numInt := int64(-123456)

  expectedNumStr := "-123456.000"

  exponent := 3

  bINum := new(BigIntNum).NewInt64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64Exponent_07(t *testing.T) {

  numInt := int64(0)

  expectedNumStr := "0"

  exponent := 0

  bINum := new(BigIntNum).NewInt64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64Exponent_08(t *testing.T) {

  numInt := int64(0)

  expectedNumStr := "0.000"

  exponent := 3

  bINum := new(BigIntNum).NewInt64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewInt64Exponent_09(t *testing.T) {

  numInt := int64(0)

  exponent := -3

  expectedNumStr := "0.000"

  bINum := new(BigIntNum).NewInt64Exponent(numInt, exponent)

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_NewINumMgr_01(t *testing.T) {

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

  dec, err := Decimal{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nStr). Error='%v' ",
      err.Error())
  }

  bINum, err := new(BigIntNum).NewINumMgr(&dec)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewINumMgr(&dec) "+
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

func TestBigIntNum_NewINumMgr_02(t *testing.T) {

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

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr). Error='%v' ",
      err.Error())
  }

  bINum, err := new(BigIntNum).NewINumMgr(&ia)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewINumMgr(&ia) "+
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

func TestBigIntNum_NewINumMgr_03(t *testing.T) {

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

  nDto0, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). Error='%v' ",
      err.Error())
  }

  bINum, err := new(BigIntNum).NewINumMgr(&nDto0)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewINumMgr(&nDto0) "+
      "Error='%v' ", err.Error())
  }

  nDto1, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

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

  if nStr != nDto1.GetNumStr() {
    t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
      nStr, nDto1.GetNumStr())
  }

}

func TestBigIntNum_NewINumMgr_04(t *testing.T) {

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

  dec, err := Decimal{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nStr). Error='%v' ",
      err.Error())
  }

  bINum, err := new(BigIntNum).NewINumMgr(dec.GetThisPointer())

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewINumMgr(&dec) "+
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

func TestBigIntNum_NewINumMgr_05(t *testing.T) {

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

  dec := Decimal{}.NewPtr()
  err := dec.SetNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nStr). Error='%v' ",
      err.Error())
  }

  bINum, err := new(BigIntNum).NewINumMgr(dec)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewINumMgr(&dec) "+
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
