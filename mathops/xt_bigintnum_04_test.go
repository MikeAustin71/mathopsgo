package mathops

import (
  "fmt"
  "math/big"
  "strconv"
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

  ePrefix := "TestBigIntNum_NewBigFloat_01"

  expectedNumStr := "32.123"

  maxPrecisionUint := uint(3)

  bfloat, isOk := big.NewFloat(0).SetString(expectedNumStr)

  if !isOk {
    t.Errorf("%v\n"+
      "Error: big.NewFloat(0).SetString(expectedNumStr) Failed!\n"+
      "Because isOk == false\n"+
      "expectedNumStr = '%v'\n\n",
      ePrefix, expectedNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewBigFloat(bfloat, maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewBigFloat(bfloat, maxPrecisionUint)\n"+
      "bfloat= '%v'\n"+
      "maxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bfloat.Text('f', int(maxPrecisionUint)),
      maxPrecisionUint,
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

  return
}

func TestBigIntNum_NewBigFloat_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewBigFloat_02"

  expectedNumStr := "32.13"

  maxPrecisionUint := uint(2)

  bfloat, isOk := big.NewFloat(0).SetString(expectedNumStr)

  if !isOk {
    t.Errorf("%v\n"+
      "Error: big.NewFloat(0).SetString(expectedNumStr) Failed!\n"+
      "Because isOk == false\n"+
      "expectedNumStr = '%v'\n\n",
      ePrefix, expectedNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewBigFloat(bfloat, maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewBigFloat(bfloat, maxPrecisionUint)\n"+
      "bfloat= '%v'\n"+
      "maxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bfloat.Text('f', int(maxPrecisionUint)),
      maxPrecisionUint,
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

  return
}

func TestBigIntNum_NewBigFloat_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewBigFloat_03"

  expectedNumStr := "-32.13"

  maxPrecisionUint := uint(2)

  bfloat, isOk := big.NewFloat(0).SetString(expectedNumStr)

  if !isOk {
    t.Errorf("%v\n"+
      "Error: big.NewFloat(0).SetString(expectedNumStr) Failed!\n"+
      "Because isOk == false\n"+
      "expectedNumStr = '%v'\n\n",
      ePrefix, expectedNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewBigFloat(bfloat, maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewBigFloat(bfloat, maxPrecisionUint)\n"+
      "bfloat= '%v'\n"+
      "maxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bfloat.Text('f', int(maxPrecisionUint)),
      maxPrecisionUint,
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

  return
}

func TestBigIntNum_NewFloat32_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NewFloat32_01"

  numf32 := float32(32.123)

  expectedNumStr := "32.123"

  maxPrecisionUint := uint(4)

  numf32Str := fmt.Sprintf("%f4", numf32)

  expectedBigINum, err := new(BigIntNum).NewFloat32(numf32, maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewFloat32(numf32, maxPrecision)\n"+
      "numf32='%v'\n"+
      "maxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numf32Str,
      maxPrecisionUint,
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

  return
}

func TestBigIntNum_NewFloat32_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewFloat32_02"

  numf32 := float32(32.129)

  expectedNumStr := "32.13"

  maxPrecisionUint := uint(2)

  numf32Str := fmt.Sprintf("%f2", numf32)

  expectedBigINum, err := new(BigIntNum).NewFloat32(numf32, maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewFloat32(numf32, maxPrecision)\n"+
      "numf32='%v'\n"+
      "maxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numf32Str,
      maxPrecisionUint,
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

  return
}

func TestBigIntNum_NewFloat32_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewFloat32_03"

  numf32 := float32(-32.129)

  expectedNumStr := "-32.13"

  maxPrecisionUint := uint(2)

  numf32Str := fmt.Sprintf("%f2", numf32)

  expectedBigINum, err := new(BigIntNum).NewFloat32(numf32, maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewFloat32(numf32, maxPrecision)\n"+
      "numf32='%v'\n"+
      "maxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numf32Str,
      maxPrecisionUint,
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

  return
}

func TestBigIntNum_NewFloat64_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NewFloat64_01"

  var numf64 float64

  numf64 = 32.123

  numf64Str := strconv.FormatFloat(numf64, 'f', 3, 64)

  expectedNumStr := "32.123"

  maxPrecisionUint := uint(3)

  expectedBigINum, err := new(BigIntNum).NewFloat64(numf64, maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewFloat64(numf64, maxPrecision)\n"+
      "numf64='%v'\n"+
      "maxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numf64Str,
      maxPrecisionUint,
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

  return
}

func TestBigIntNum_NewFloat64_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewFloat64_02"

  var numf64 float64

  numf64 = 32.129

  numf64Str := strconv.FormatFloat(numf64, 'f', 3, 64)

  expectedNumStr := "32.13"

  maxPrecisionUint := uint(2)

  expectedBigINum, err := new(BigIntNum).NewFloat64(numf64, maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewFloat64(numf64, maxPrecision)\n"+
      "numf64='%v'\n"+
      "maxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numf64Str,
      maxPrecisionUint,
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

  return
}

func TestBigIntNum_NewFloat64_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewFloat64_03"

  var numf64 float64

  numf64 = -32.129

  numf64Str := strconv.FormatFloat(numf64, 'f', 3, 64)

  expectedNumStr := "-32.13"

  maxPrecisionUint := uint(2)

  expectedBigINum, err := new(BigIntNum).NewFloat64(numf64, maxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewFloat64(numf64, maxPrecision)\n"+
      "numf64='%v'\n"+
      "maxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numf64Str,
      maxPrecisionUint,
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

  return
}

func TestBigIntNum_NewInt_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt_01"

  var numInt int

  numInt = 1234

  precisionUint := uint(3)

  expectedNumStr := "1.234"

  expectedBigINum, err := new(BigIntNum).NewInt(numInt, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt(numInt, precisionUint)\n"+
      "numInt= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt_02"

  var numInt int

  numInt = 1234

  precisionUint := uint(0)

  expectedNumStr := "1234"

  expectedBigINum, err := new(BigIntNum).NewInt(numInt, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt(numInt, precisionUint)\n"+
      "numInt= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt_03"

  var numInt int

  numInt = -1234

  precisionUint := uint(3)

  expectedNumStr := "-1.234"

  expectedBigINum, err := new(BigIntNum).NewInt(numInt, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt(numInt, precisionUint)\n"+
      "numInt= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt_04(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt_04"

  var numInt int

  numInt = -1234

  precisionUint := uint(0)

  expectedNumStr := "-1234"

  expectedBigINum, err := new(BigIntNum).NewInt(numInt, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt(numInt, precisionUint)\n"+
      "numInt= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt_05(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt_05"

  var numInt int

  numInt = 0

  precisionUint := uint(0)

  expectedNumStr := "0"

  expectedBigINum, err := new(BigIntNum).NewInt(numInt, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt(numInt, precisionUint)\n"+
      "numInt= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt_06(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt_06"

  var numInt int

  numInt = 0

  precisionUint := uint(3)

  expectedNumStr := "0.000"

  expectedBigINum, err := new(BigIntNum).NewInt(numInt, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt(numInt, precisionUint)\n"+
      "numInt= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      precisionUint,
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

  return
}

func TestBigIntNum_NewIntExponent_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NewIntExponent_01"

  var numInt, exponentInt int

  numInt = 1234

  exponentInt = 3

  expectedNumStr := "1234.000"

  expectedBigINum, err := new(BigIntNum).NewIntExponent(numInt, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewIntExponent(numInt, exponentInt)\n"+
      "numInt= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      exponentInt,
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

  return
}

func TestBigIntNum_NewIntExponent_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewIntExponent_02"

  var numInt, exponentInt int

  numInt = 123456

  exponentInt = -2

  expectedNumStr := "1234.56"

  expectedBigINum, err := new(BigIntNum).NewIntExponent(numInt, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewIntExponent(numInt, exponentInt)\n"+
      "numInt= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      exponentInt,
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

  return
}

func TestBigIntNum_NewIntExponent_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewIntExponent_03"

  var numInt, exponentInt int

  numInt = 123456

  exponentInt = 0

  expectedNumStr := "123456"

  expectedBigINum, err := new(BigIntNum).NewIntExponent(numInt, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewIntExponent(numInt, exponentInt)\n"+
      "numInt= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      exponentInt,
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

  return
}

func TestBigIntNum_NewIntExponent_04(t *testing.T) {

  ePrefix := "TestBigIntNum_NewIntExponent_04"

  var numInt, exponentInt int

  numInt = 0

  exponentInt = 0

  expectedNumStr := "0"

  expectedBigINum, err := new(BigIntNum).NewIntExponent(numInt, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewIntExponent(numInt, exponentInt)\n"+
      "numInt= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      exponentInt,
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

  return
}

func TestBigIntNum_NewIntExponent_05(t *testing.T) {

  ePrefix := "TestBigIntNum_NewIntExponent_05"

  var numInt, exponentInt int

  numInt = 0

  exponentInt = 3

  expectedNumStr := "0.000"

  expectedBigINum, err := new(BigIntNum).NewIntExponent(numInt, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewIntExponent(numInt, exponentInt)\n"+
      "numInt= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      exponentInt,
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

  return
}

func TestBigIntNum_NewIntExponent_06(t *testing.T) {

  ePrefix := "TestBigIntNum_NewIntExponent_05"

  var numInt, exponentInt int

  numInt = 0

  exponentInt = -3

  expectedNumStr := "0.000"

  expectedBigINum, err := new(BigIntNum).NewIntExponent(numInt, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewIntExponent(numInt, exponentInt)\n"+
      "numInt= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt32_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32_01"

  var numInt32 int32

  var precisionUint uint

  numInt32 = 1234

  precisionUint = 3

  expectedNumStr := "1.234"

  expectedBigINum, err := new(BigIntNum).NewInt32(numInt32, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32(numInt32, precisionUint)\n"+
      "numInt32= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt32_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32_02"

  var numInt32 int32

  var precisionUint uint

  numInt32 = int32(1234)

  precisionUint = uint(0)

  expectedNumStr := "1234"

  expectedBigINum, err := new(BigIntNum).NewInt32(numInt32, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32(numInt32, precisionUint)\n"+
      "numInt32= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt32_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32_03"

  var numInt32 int32

  var precisionUint uint

  numInt32 = -1234

  precisionUint = 3

  expectedNumStr := "-1.234"

  expectedBigINum, err := new(BigIntNum).NewInt32(numInt32, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32(numInt32, precisionUint)\n"+
      "numInt32= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt32_04(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32_04"

  var numInt32 int32

  var precisionUint uint

  numInt32 = -1234

  precisionUint = 0

  expectedNumStr := "-1234"

  expectedBigINum, err := new(BigIntNum).NewInt32(numInt32, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32(numInt32, precisionUint)\n"+
      "numInt32= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt32_05(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32_05"

  var numInt32 int32

  var precisionUint uint

  numInt32 = 0

  precisionUint = 0

  expectedNumStr := "0"

  expectedBigINum, err := new(BigIntNum).NewInt32(numInt32, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32(numInt32, precisionUint)\n"+
      "numInt32= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt32_06(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32_06"

  var numInt32 int32

  var precisionUint uint

  numInt32 = 0

  precisionUint = 3

  expectedNumStr := "0.000"

  expectedBigINum, err := new(BigIntNum).NewInt32(numInt32, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32(numInt32, precisionUint)\n"+
      "numInt32= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt32Exponent_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32Exponent_01"

  var numInt32 int32

  var exponentInt int

  numInt32 = 1234

  exponentInt = 3

  expectedNumStr := "1234.000"

  expectedBigINum, err := new(BigIntNum).NewInt32Exponent(numInt32, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32Exponent(numInt32, exponentInt)\n"+
      "numInt32= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt32Exponent_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32Exponent_02"

  var numInt32 int32

  var exponentInt int

  numInt32 = 123456

  exponentInt = -2

  expectedNumStr := "1234.56"

  expectedBigINum, err := new(BigIntNum).NewInt32Exponent(numInt32, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32Exponent(numInt32, exponentInt)\n"+
      "numInt32= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt32Exponent_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32Exponent_03"

  var numInt32 int32

  var exponentInt int

  numInt32 = 123456

  exponentInt = 0

  expectedNumStr := "123456"

  expectedBigINum, err := new(BigIntNum).NewInt32Exponent(numInt32, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32Exponent(numInt32, exponentInt)\n"+
      "numInt32= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt32Exponent_04(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32Exponent_04"

  var numInt32 int32

  var exponentInt int

  numInt32 = 0

  exponentInt = 0

  expectedNumStr := "0"

  expectedBigINum, err := new(BigIntNum).NewInt32Exponent(numInt32, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32Exponent(numInt32, exponentInt)\n"+
      "numInt32= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt32Exponent_05(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32Exponent_05"

  var numInt32 int32

  var exponentInt int

  numInt32 = 0

  exponentInt = 3

  expectedNumStr := "0.000"

  expectedBigINum, err := new(BigIntNum).NewInt32Exponent(numInt32, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32Exponent(numInt32, exponentInt)\n"+
      "numInt32= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt32Exponent_06(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt32Exponent_06"

  var numInt32 int32

  var exponentInt int

  numInt32 = 0

  exponentInt = -3

  expectedNumStr := "0.000"

  expectedBigINum, err := new(BigIntNum).NewInt32Exponent(numInt32, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt32Exponent(numInt32, exponentInt)\n"+
      "numInt32= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt32,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt64_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64_01"

  var numInt64 int64

  var precisionUint uint

  numInt64 = 1234

  precisionUint = 3

  expectedNumStr := "1.234"

  expectedBigINum, err := new(BigIntNum).NewInt64(numInt64, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64(numInt64, precisionUint)\n"+
      "numInt64= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt64_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64_02"

  var numInt64 int64

  var precisionUint uint

  numInt64 = 1234

  precisionUint = 0

  expectedNumStr := "1234"

  expectedBigINum, err := new(BigIntNum).NewInt64(numInt64, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64(numInt64, precisionUint)\n"+
      "numInt64= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt64_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64_03"

  var numInt64 int64

  var precisionUint uint

  numInt64 = -1234

  precisionUint = 3

  expectedNumStr := "-1.234"

  expectedBigINum, err := new(BigIntNum).NewInt64(numInt64, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64(numInt64, precisionUint)\n"+
      "numInt64= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt64_04(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64_04"

  var numInt64 int64

  var precisionUint uint

  numInt64 = -1234

  precisionUint = 0

  expectedNumStr := "-1234"

  expectedBigINum, err := new(BigIntNum).NewInt64(numInt64, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64(numInt64, precisionUint)\n"+
      "numInt64= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt64_05(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64_05"

  var numInt64 int64

  var precisionUint uint

  numInt64 = 0

  precisionUint = 0

  expectedNumStr := "0"

  expectedBigINum, err := new(BigIntNum).NewInt64(numInt64, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64(numInt64, precisionUint)\n"+
      "numInt64= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt64_06(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64_06"

  var numInt64 int64

  var precisionUint uint

  numInt64 = 0

  precisionUint = 3

  expectedNumStr := "0.000"

  expectedBigINum, err := new(BigIntNum).NewInt64(numInt64, precisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64(numInt64, precisionUint)\n"+
      "numInt64= '%v'\n"+
      "precisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      precisionUint,
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

  return
}

func TestBigIntNum_NewInt64Exponent_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64Exponent_01"

  var numInt64 int64

  var exponentInt int

  numInt64 = 1234

  exponentInt = 3

  expectedNumStr := "1234.000"

  expectedBigINum, err := new(BigIntNum).NewInt64Exponent(numInt64, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64Exponent(numInt64, exponentInt)\n"+
      "numInt64= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt64Exponent_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64Exponent_02"

  var numInt64 int64

  var exponentInt int

  numInt64 = 123456

  exponentInt = -2

  expectedNumStr := "1234.56"

  expectedBigINum, err := new(BigIntNum).NewInt64Exponent(numInt64, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64Exponent(numInt64, exponentInt)\n"+
      "numInt64= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt64Exponent_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64Exponent_03"

  var numInt64 int64

  var exponentInt int

  numInt64 = 123456

  exponentInt = 0

  expectedNumStr := "123456"

  expectedBigINum, err := new(BigIntNum).NewInt64Exponent(numInt64, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64Exponent(numInt64, exponentInt)\n"+
      "numInt64= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt64Exponent_04(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64Exponent_04"

  var numInt64 int64

  var exponentInt int

  numInt64 = -123456

  exponentInt = 0

  expectedNumStr := "-123456"

  expectedBigINum, err := new(BigIntNum).NewInt64Exponent(numInt64, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64Exponent(numInt64, exponentInt)\n"+
      "numInt64= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt64Exponent_05(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64Exponent_05"

  var numInt64 int64

  var exponentInt int

  numInt64 = -123456

  exponentInt = -3

  expectedNumStr := "-123.456"

  expectedBigINum, err := new(BigIntNum).NewInt64Exponent(numInt64, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64Exponent(numInt64, exponentInt)\n"+
      "numInt64= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt64Exponent_06(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64Exponent_06"

  var numInt64 int64

  var exponentInt int

  numInt64 = -123456

  exponentInt = 3

  expectedNumStr := "-123456.000"

  expectedBigINum, err := new(BigIntNum).NewInt64Exponent(numInt64, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64Exponent(numInt64, exponentInt)\n"+
      "numInt64= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt64Exponent_07(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64Exponent_07"

  var numInt64 int64

  var exponentInt int

  numInt64 = 0

  exponentInt = 0

  expectedNumStr := "0"

  expectedBigINum, err := new(BigIntNum).NewInt64Exponent(numInt64, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64Exponent(numInt64, exponentInt)\n"+
      "numInt64= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt64Exponent_08(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64Exponent_08"

  var numInt64 int64

  var exponentInt int

  numInt64 = 0

  exponentInt = 3

  expectedNumStr := "0.000"

  expectedBigINum, err := new(BigIntNum).NewInt64Exponent(numInt64, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64Exponent(numInt64, exponentInt)\n"+
      "numInt64= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      exponentInt,
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

  return
}

func TestBigIntNum_NewInt64Exponent_09(t *testing.T) {

  ePrefix := "TestBigIntNum_NewInt64Exponent_09"

  var numInt64 int64

  var exponentInt int

  numInt64 = 0

  exponentInt = -3

  expectedNumStr := "0.000"

  expectedBigINum, err := new(BigIntNum).NewInt64Exponent(numInt64, exponentInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewInt64Exponent(numInt64, exponentInt)\n"+
      "numInt64= '%v'\n"+
      "exponentInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numInt64,
      exponentInt,
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

  return
}

func TestBigIntNum_NewINumMgr_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NewINumMgr_01"

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

  dec, err := new(Decimal).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec, err := new(Decimal).\n"+
      "  NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != decNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumberStr \n"+
      "Expected decNumberStr = '%v'\n"+
      "  Actual decNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumberStr)

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

  decScaleFactor, err := dec.GetScaleVal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decScaleFactor, err := dec.GetScaleVal()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(decScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & dec Scale Factors ARE NOT EQUAL!\n"+
      "Because expectedScaleFactor.Cmp(decScaleFactor) != 0\n"+
      "Expected decScaleFactor = '%v'\n"+
      "  Actual decScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), decScaleFactor.Text(10))

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

  bINum, err := new(BigIntNum).NewINumMgr(&dec)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewINumMgr(&dec)\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
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

  nDto, err := new(NumStrDto).NewBigInt(expectedBigIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto, err := new(NumStrDto).\n"+
      "  NewBigInt(expectedBigIntNum, expectedPrecisionUint)\n"+
      "expectedBigIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigIntNum.Text(10),
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = nDto.IsValid("Validating nDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = nDto.IsValid('Validating nDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nDtoNumberStr, err := nDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoNumberStr, err := nDto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != nDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != nDtoNumberStr \n"+
      "Expected nDtoNumberStr = '%v'\n"+
      "  Actual nDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, nDtoNumberStr)

    return
  }

  nDtoPrecisionUint, err := nDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoPrecisionUint, err := nDto.GetPrecisionUint()\n"+
      "nDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != nDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/nDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != nDtoPrecisionUint\n"+
      "Expected nDtoPrecisionUint = '%v'\n"+
      "  Actual nDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, nDtoPrecisionUint)

    return
  }

  nDtoBigInt, err := nDto.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoBigInt, err := nDto.GetBigInt()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(nDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/nDto Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(nDtoBigInt) != 0\n"+
      "Expected nDtoBigInt = '%v'\n"+
      "  Actual nDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), nDtoBigInt.Text(10))

    return
  }

  nDtoAbsBigIntNum, _, err := nDto.GetAbsoluteBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoAbsBigIntNum, _, err := nDto.GetAbsoluteBigInt()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(nDtoAbsBigIntNum) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(nDtoAbsBigIntNum) != 0\n"+
      "Expected nDtoAbsBigIntNum = '%v'\n"+
      "  Actual nDtoAbsBigIntNum = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), nDtoAbsBigIntNum.Text(10))

    return
  }

  nDtoScaleFactor, err := nDto.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoScaleFactor, err := nDto.GetScaleVal()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(nDtoScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto Scale Factors ARE NOT EQUAL!\n"+
      "Because  expectedScaleFactor.Cmp(nDtoScaleFactor) != 0\n"+
      "Expected nDtoScaleFactor = '%v'\n"+
      "  Actual nDtoScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), nDtoScaleFactor.Text(10))

    return
  }

  nDtoSignValue, err := nDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoSignValue, err := nDto.GetSign()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedSignVal != nDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & nDto Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != nDtoSignValue\n"+
      "Expected nDtoSignValue = '%v'\n"+
      "  Actual nDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, nDtoSignValue)

    return
  }

  nDtoSeps, err := nDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoSeps, err := nDto.GetNumericSeparatorsDto()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(nDtoSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != nDtoSeps \n"+
      "Expected nDtoSeps = '%v'\n"+
      "  Actual nDtoSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), nDtoSeps.String())

    return
  }

  return
}

func TestBigIntNum_NewINumMgr_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NewINumMgr_02"

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

  ia, err := new(IntAry).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "ia, err := new(IntAry).\n"+
      "  NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
    return
  }

  err = ia.IsValid("Validating ia")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = ia.IsValid('Validating ia')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaNumberStr, err := ia.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaNumberStr, err := ia.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != iaNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != iaNumberStr \n"+
      "Expected iaNumberStr = '%v'\n"+
      "  Actual iaNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, iaNumberStr)

    return
  }

  iaPrecisionUint, err := ia.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaPrecisionUint, err := ia.GetPrecisionUint()\n"+
      "ia= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != iaPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ia Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != iaPrecisionUint\n"+
      "Expected iaPrecisionUint = '%v'\n"+
      "  Actual iaPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, iaPrecisionUint)

    return
  }

  iaBigInt, err := ia.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaBigInt, err := ia.GetBigInt()\n"+
      "ia= '%v\n"+
      "Error= '%v'\n\n", ePrefix, iaNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(iaBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/ia Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(iaBigInt) != 0\n"+
      "Expected iaBigInt = '%v'\n"+
      "  Actual iaBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), iaBigInt.Text(10))

    return
  }

  iaAbsValue, err := ia.GetAbsoluteValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaAbsValue, err := ia.GetAbsoluteValue()\n"+
      "ia= '%v\n"+
      "Error= '%v'\n\n", ePrefix, iaNumberStr, err.Error())
    return
  }

  iaAbsBigInt, err := iaAbsValue.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaAbsBigInt, err := iaAbsValue.GetBigInt()\n"+
      "ia= '%v\n"+
      "Error= '%v'\n\n", ePrefix, iaNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(iaAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & ia Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(iaAbsBigInt) != 0\n"+
      "Expected iaAbsBigInt = '%v'\n"+
      "  Actual iaAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), iaAbsBigInt.Text(10))

    return
  }

  iaScaleFactor, err := ia.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaScaleFactor, err := ia.GetScaleFactor()\n"+
      "ia= '%v\n"+
      "Error= '%v'\n\n", ePrefix, iaNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(iaScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & ia Scale Factors ARE NOT EQUAL!\n"+
      "Because expectedScaleFactor.Cmp(iaScaleFactor) != 0\n"+
      "Expected iaScaleFactor = '%v'\n"+
      "  Actual iaScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), iaScaleFactor.Text(10))

    return
  }

  iaSignValue, err := ia.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSignValue, err := ia.GetSign()\n"+
      "ia= '%v\n"+
      "Error= '%v'\n\n", ePrefix, iaNumberStr, err.Error())
    return
  }

  if expectedSignVal != iaSignValue {
    t.Errorf("%v\n"+
      "Error: expected & ia Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != iaSignValue\n"+
      "Expected iaSignValue = '%v'\n"+
      "  Actual iaSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, iaSignValue)

    return
  }

  iaNumSeps, err := ia.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaNumSeps, err := ia.GetNumericSeparatorsDto()\n"+
      "ia= '%v\n"+
      "Error= '%v'\n\n", ePrefix, iaNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(iaNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != iaNumSeps \n"+
      "Expected iaNumSeps = '%v'\n"+
      "  Actual iaNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), iaNumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewINumMgr(&ia)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewINumMgr(&ia)\n"+
      "ia= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, iaNumberStr, err.Error())
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

  nDto, err := new(NumStrDto).NewBigInt(expectedBigIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto, err := new(NumStrDto).\n"+
      "  NewBigInt(expectedBigIntNum, expectedPrecisionUint)\n"+
      "expectedBigIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigIntNum.Text(10),
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = nDto.IsValid("Validating nDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = nDto.IsValid('Validating nDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nDtoNumberStr, err := nDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoNumberStr, err := nDto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != nDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != nDtoNumberStr \n"+
      "Expected nDtoNumberStr = '%v'\n"+
      "  Actual nDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, nDtoNumberStr)

    return
  }

  nDtoPrecisionUint, err := nDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoPrecisionUint, err := nDto.GetPrecisionUint()\n"+
      "nDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != nDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/nDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != nDtoPrecisionUint\n"+
      "Expected nDtoPrecisionUint = '%v'\n"+
      "  Actual nDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, nDtoPrecisionUint)

    return
  }

  nDtoBigInt, err := nDto.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoBigInt, err := nDto.GetBigInt()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(nDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/nDto Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(nDtoBigInt) != 0\n"+
      "Expected nDtoBigInt = '%v'\n"+
      "  Actual nDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), nDtoBigInt.Text(10))

    return
  }

  nDtoAbsBigIntNum, _, err := nDto.GetAbsoluteBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoAbsBigIntNum, _, err := nDto.GetAbsoluteBigInt()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(nDtoAbsBigIntNum) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(nDtoAbsBigIntNum) != 0\n"+
      "Expected nDtoAbsBigIntNum = '%v'\n"+
      "  Actual nDtoAbsBigIntNum = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), nDtoAbsBigIntNum.Text(10))

    return
  }

  nDtoScaleFactor, err := nDto.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoScaleFactor, err := nDto.GetScaleVal()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(nDtoScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto Scale Factors ARE NOT EQUAL!\n"+
      "Because  expectedScaleFactor.Cmp(nDtoScaleFactor) != 0\n"+
      "Expected nDtoScaleFactor = '%v'\n"+
      "  Actual nDtoScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), nDtoScaleFactor.Text(10))

    return
  }

  nDtoSignValue, err := nDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoSignValue, err := nDto.GetSign()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedSignVal != nDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & nDto Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != nDtoSignValue\n"+
      "Expected nDtoSignValue = '%v'\n"+
      "  Actual nDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, nDtoSignValue)

    return
  }

  nDtoSeps, err := nDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoSeps, err := nDto.GetNumericSeparatorsDto()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(nDtoSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != nDtoSeps \n"+
      "Expected nDtoSeps = '%v'\n"+
      "  Actual nDtoSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), nDtoSeps.String())

    return
  }

  return
}

func TestBigIntNum_NewINumMgr_03(t *testing.T) {

  ePrefix := "TestBigIntNum_NewINumMgr_03"

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

  nDto0, err := new(NumStrDto).NewBigInt(expectedBigIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto0, err := new(NumStrDto).\n"+
      "  NewBigInt(expectedBigIntNum, expectedPrecisionUint)\n"+
      "expectedBigIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigIntNum.Text(10),
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = nDto0.IsValid("Validating nDto0")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = nDto0.IsValid('Validating nDto0')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nDto0NumberStr, err := nDto0.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto0NumberStr, err := nDto0.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != nDto0NumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != nDto0NumberStr \n"+
      "Expected nDto0NumberStr = '%v'\n"+
      "  Actual nDto0NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, nDto0NumberStr)

    return
  }

  nDto0PrecisionUint, err := nDto0.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto0PrecisionUint, err := nDto0.GetPrecisionUint()\n"+
      "nDto0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, nDto0NumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != nDto0PrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/nDto0 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != nDto0PrecisionUint\n"+
      "Expected nDto0PrecisionUint = '%v'\n"+
      "  Actual nDto0PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, nDto0PrecisionUint)

    return
  }

  nDto0BigInt, err := nDto0.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto0BigInt, err := nDto0.GetBigInt()\n"+
      "nDto0= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto0NumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(nDto0BigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/nDto0 Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(nDto0BigInt) != 0\n"+
      "Expected nDto0BigInt = '%v'\n"+
      "  Actual nDto0BigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), nDto0BigInt.Text(10))

    return
  }

  nDto0AbsBigIntNum, _, err := nDto0.GetAbsoluteBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto0AbsBigIntNum, _, err := nDto0.GetAbsoluteBigInt()\n"+
      "nDto0= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto0NumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(nDto0AbsBigIntNum) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto0 Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(nDto0AbsBigIntNum) != 0\n"+
      "Expected nDto0AbsBigIntNum = '%v'\n"+
      "  Actual nDto0AbsBigIntNum = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), nDto0AbsBigIntNum.Text(10))

    return
  }

  nDto0ScaleFactor, err := nDto0.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto0ScaleFactor, err := nDto0.GetScaleVal()\n"+
      "nDto0= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto0NumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(nDto0ScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto0 Scale Factors ARE NOT EQUAL!\n"+
      "Because  expectedScaleFactor.Cmp(nDto0ScaleFactor) != 0\n"+
      "Expected nDto0ScaleFactor = '%v'\n"+
      "  Actual nDto0ScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), nDto0ScaleFactor.Text(10))

    return
  }

  nDto0SignValue, err := nDto0.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto0SignValue, err := nDto0.GetSign()\n"+
      "nDto0= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto0NumberStr, err.Error())
    return
  }

  if expectedSignVal != nDto0SignValue {
    t.Errorf("%v\n"+
      "Error: expected & nDto0 Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != nDto0SignValue\n"+
      "Expected nDto0SignValue = '%v'\n"+
      "  Actual nDto0SignValue = '%v'\n\n",
      ePrefix, expectedSignVal, nDto0SignValue)

    return
  }

  nDto0Seps, err := nDto0.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto0Seps, err := nDto0.GetNumericSeparatorsDto()\n"+
      "nDto0= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto0NumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(nDto0Seps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != nDto0Seps \n"+
      "Expected nDto0Seps = '%v'\n"+
      "  Actual nDto0Seps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), nDto0Seps.String())

    return
  }

  bINum, err := new(BigIntNum).NewINumMgr(&nDto0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewINumMgr(&nDto0)\n"+
      "nDto0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, nDto0NumberStr, err.Error())
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

  nDto1, err := new(NumStrDto).NewBigInt(expectedBigIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto1, err := new(NumStrDto).\n"+
      "  NewBigInt(expectedBigIntNum, expectedPrecisionUint)\n"+
      "expectedBigIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigIntNum.Text(10),
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = nDto1.IsValid("Validating nDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = nDto1.IsValid('Validating nDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nDto1NumberStr, err := nDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto1NumberStr, err := nDto1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != nDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != nDto1NumberStr \n"+
      "Expected nDto1NumberStr = '%v'\n"+
      "  Actual nDto1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, nDto1NumberStr)

    return
  }

  nDto1PrecisionUint, err := nDto1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto1PrecisionUint, err := nDto1.GetPrecisionUint()\n"+
      "nDto1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, nDto1NumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != nDto1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/nDto1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != nDto1PrecisionUint\n"+
      "Expected nDto1PrecisionUint = '%v'\n"+
      "  Actual nDto1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, nDto1PrecisionUint)

    return
  }

  nDto1BigInt, err := nDto1.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto1BigInt, err := nDto1.GetBigInt()\n"+
      "nDto1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto1NumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(nDto1BigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/nDto1 Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(nDto1BigInt) != 0\n"+
      "Expected nDto1BigInt = '%v'\n"+
      "  Actual nDto1BigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), nDto1BigInt.Text(10))

    return
  }

  nDto1AbsBigIntNum, _, err := nDto1.GetAbsoluteBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto1AbsBigIntNum, _, err := nDto1.GetAbsoluteBigInt()\n"+
      "nDto1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto1NumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(nDto1AbsBigIntNum) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto1 Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(nDto1AbsBigIntNum) != 0\n"+
      "Expected nDto1AbsBigIntNum = '%v'\n"+
      "  Actual nDto1AbsBigIntNum = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), nDto1AbsBigIntNum.Text(10))

    return
  }

  nDto1ScaleFactor, err := nDto1.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto1ScaleFactor, err := nDto1.GetScaleVal()\n"+
      "nDto1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto1NumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(nDto1ScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto1 Scale Factors ARE NOT EQUAL!\n"+
      "Because  expectedScaleFactor.Cmp(nDto1ScaleFactor) != 0\n"+
      "Expected nDto1ScaleFactor = '%v'\n"+
      "  Actual nDto1ScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), nDto1ScaleFactor.Text(10))

    return
  }

  nDto1SignValue, err := nDto1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto1SignValue, err := nDto1.GetSign()\n"+
      "nDto1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto1NumberStr, err.Error())
    return
  }

  if expectedSignVal != nDto1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & nDto1 Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != nDto1SignValue\n"+
      "Expected nDto1SignValue = '%v'\n"+
      "  Actual nDto1SignValue = '%v'\n\n",
      ePrefix, expectedSignVal, nDto1SignValue)

    return
  }

  nDto1Seps, err := nDto1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto1Seps, err := nDto1.GetNumericSeparatorsDto()\n"+
      "nDto1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDto1NumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(nDto1Seps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != nDto1Seps \n"+
      "Expected nDto1Seps = '%v'\n"+
      "  Actual nDto1Seps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), nDto1Seps.String())

    return
  }

  return
}

func TestBigIntNum_NewINumMgr_04(t *testing.T) {

  ePrefix := "TestBigIntNum_NewINumMgr_04"

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

  dec, err := new(Decimal).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec, err := new(Decimal).\n"+
      "  NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != decNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumberStr \n"+
      "Expected decNumberStr = '%v'\n"+
      "  Actual decNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumberStr)

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

  decScaleFactor, err := dec.GetScaleVal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decScaleFactor, err := dec.GetScaleVal()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(decScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & dec Scale Factors ARE NOT EQUAL!\n"+
      "Because expectedScaleFactor.Cmp(decScaleFactor) != 0\n"+
      "Expected decScaleFactor = '%v'\n"+
      "  Actual decScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), decScaleFactor.Text(10))

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

  bINum, err := new(BigIntNum).NewINumMgr(&dec)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewINumMgr(&dec)\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
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

  nDto, err := new(NumStrDto).NewBigInt(expectedBigIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto, err := new(NumStrDto).\n"+
      "  NewBigInt(expectedBigIntNum, expectedPrecisionUint)\n"+
      "expectedBigIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigIntNum.Text(10),
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = nDto.IsValid("Validating nDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = nDto.IsValid('Validating nDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nDtoNumberStr, err := nDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoNumberStr, err := nDto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != nDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != nDtoNumberStr \n"+
      "Expected nDtoNumberStr = '%v'\n"+
      "  Actual nDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, nDtoNumberStr)

    return
  }

  nDtoPrecisionUint, err := nDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoPrecisionUint, err := nDto.GetPrecisionUint()\n"+
      "nDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != nDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/nDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != nDtoPrecisionUint\n"+
      "Expected nDtoPrecisionUint = '%v'\n"+
      "  Actual nDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, nDtoPrecisionUint)

    return
  }

  nDtoBigInt, err := nDto.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoBigInt, err := nDto.GetBigInt()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(nDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/nDto Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(nDtoBigInt) != 0\n"+
      "Expected nDtoBigInt = '%v'\n"+
      "  Actual nDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), nDtoBigInt.Text(10))

    return
  }

  nDtoAbsBigIntNum, _, err := nDto.GetAbsoluteBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoAbsBigIntNum, _, err := nDto.GetAbsoluteBigInt()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(nDtoAbsBigIntNum) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(nDtoAbsBigIntNum) != 0\n"+
      "Expected nDtoAbsBigIntNum = '%v'\n"+
      "  Actual nDtoAbsBigIntNum = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), nDtoAbsBigIntNum.Text(10))

    return
  }

  nDtoScaleFactor, err := nDto.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoScaleFactor, err := nDto.GetScaleVal()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(nDtoScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto Scale Factors ARE NOT EQUAL!\n"+
      "Because  expectedScaleFactor.Cmp(nDtoScaleFactor) != 0\n"+
      "Expected nDtoScaleFactor = '%v'\n"+
      "  Actual nDtoScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), nDtoScaleFactor.Text(10))

    return
  }

  nDtoSignValue, err := nDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoSignValue, err := nDto.GetSign()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedSignVal != nDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & nDto Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != nDtoSignValue\n"+
      "Expected nDtoSignValue = '%v'\n"+
      "  Actual nDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, nDtoSignValue)

    return
  }

  nDtoSeps, err := nDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoSeps, err := nDto.GetNumericSeparatorsDto()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(nDtoSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != nDtoSeps \n"+
      "Expected nDtoSeps = '%v'\n"+
      "  Actual nDtoSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), nDtoSeps.String())

    return
  }

  return
}

func TestBigIntNum_NewINumMgr_05(t *testing.T) {

  ePrefix := "TestBigIntNum_NewINumMgr_01"

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

  dec, err := new(Decimal).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec, err := new(Decimal).\n"+
      "  NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != decNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumberStr \n"+
      "Expected decNumberStr = '%v'\n"+
      "  Actual decNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumberStr)

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

  decScaleFactor, err := dec.GetScaleVal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decScaleFactor, err := dec.GetScaleVal()\n"+
      "dec= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(decScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & dec Scale Factors ARE NOT EQUAL!\n"+
      "Because expectedScaleFactor.Cmp(decScaleFactor) != 0\n"+
      "Expected decScaleFactor = '%v'\n"+
      "  Actual decScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), decScaleFactor.Text(10))

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

  bINum, err := new(BigIntNum).NewINumMgr(&dec)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewINumMgr(&dec)\n"+
      "dec= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumberStr, err.Error())
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

  nDto, err := new(NumStrDto).NewBigInt(expectedBigIntNum, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto, err := new(NumStrDto).\n"+
      "  NewBigInt(expectedBigIntNum, expectedPrecisionUint)\n"+
      "expectedBigIntNum= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigIntNum.Text(10),
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = nDto.IsValid("Validating nDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = nDto.IsValid('Validating nDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nDtoNumberStr, err := nDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoNumberStr, err := nDto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != nDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != nDtoNumberStr \n"+
      "Expected nDtoNumberStr = '%v'\n"+
      "  Actual nDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, nDtoNumberStr)

    return
  }

  nDtoPrecisionUint, err := nDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoPrecisionUint, err := nDto.GetPrecisionUint()\n"+
      "nDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != nDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/nDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != nDtoPrecisionUint\n"+
      "Expected nDtoPrecisionUint = '%v'\n"+
      "  Actual nDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, nDtoPrecisionUint)

    return
  }

  nDtoBigInt, err := nDto.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoBigInt, err := nDto.GetBigInt()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(nDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/nDto Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(nDtoBigInt) != 0\n"+
      "Expected nDtoBigInt = '%v'\n"+
      "  Actual nDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), nDtoBigInt.Text(10))

    return
  }

  nDtoAbsBigIntNum, _, err := nDto.GetAbsoluteBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoAbsBigIntNum, _, err := nDto.GetAbsoluteBigInt()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(nDtoAbsBigIntNum) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(nDtoAbsBigIntNum) != 0\n"+
      "Expected nDtoAbsBigIntNum = '%v'\n"+
      "  Actual nDtoAbsBigIntNum = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), nDtoAbsBigIntNum.Text(10))

    return
  }

  nDtoScaleFactor, err := nDto.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoScaleFactor, err := nDto.GetScaleVal()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(nDtoScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & nDto Scale Factors ARE NOT EQUAL!\n"+
      "Because  expectedScaleFactor.Cmp(nDtoScaleFactor) != 0\n"+
      "Expected nDtoScaleFactor = '%v'\n"+
      "  Actual nDtoScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), nDtoScaleFactor.Text(10))

    return
  }

  nDtoSignValue, err := nDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoSignValue, err := nDto.GetSign()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if expectedSignVal != nDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & nDto Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != nDtoSignValue\n"+
      "Expected nDtoSignValue = '%v'\n"+
      "  Actual nDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, nDtoSignValue)

    return
  }

  nDtoSeps, err := nDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoSeps, err := nDto.GetNumericSeparatorsDto()\n"+
      "nDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, nDtoNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(nDtoSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != nDtoSeps \n"+
      "Expected nDtoSeps = '%v'\n"+
      "  Actual nDtoSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), nDtoSeps.String())

    return
  }

  return
}
