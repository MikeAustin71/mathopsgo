package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntNum_BigInt_01(t *testing.T) {

  ePrefix := "TestBigIntNum_BigInt_01"

  expectedBigINumStr := "123.456"

  expectedBigINumPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nbStr := "123456"

  expectedScale := big.NewInt(1000)

  bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)\n"+
      "nbStr= '%v'\n\n",
      ePrefix, nbStr)
    return
  }

  expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

  bINum, err := new(BigIntNum).NewBigInt(bOriginal, expectedBigINumPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewBigInt(\n"+
      "  bOriginal, expectedPrecision)\n"+
      "bOriginal= '%v'\n"+
      "expectedPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bOriginal.Text(10),
      expectedBigINumPrecisionUint,
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

  bINumPrecisionUint, err := bINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  bINumBigInt, err := bINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumBigInt, err := bINum.GetBigInt()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumScaleFactor, err := bINum.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
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

  if bOriginal.Cmp(bINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values ARE NOT Equal!\n"+
      "Because bOriginal.Cmp(bINumBigInt) != 0\n"+
      "Expected bINumBigInt = '%v'\n"+
      "  Actual bINumBigInt = '%v'\n\n",
      ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

    return
  }

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  if expectedScale.Cmp(bINumScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factors ARE NOT Equal!\n"+
      "Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
      "Expected bINumScaleFactor = '%v'\n"+
      "  Actual bINumScaleFactor = '%v'\n\n",
      ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

    return
  }

  if expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Absolute BigInt Values ARE NOT Equal!\n"+
      "Because expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0\n"+
      "Expected bINumAbsBigInt = '%v'\n"+
      "  Actual bINumAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigInt.Text(10), bINumAbsBigInt.Text(10))

    return
  }

  if expectedSignVal != bINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedSignVal != bINumSignValue\n"+
      "Expected bINumSignValue = '%v'\n"+
      "  Actual bINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINumSignValue)

    return
  }

  if expectedBigINumStr != bINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != bINumStr \n"+
      "Expected bINumStr = '%v'\n"+
      "  Actual bINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumStr)

    return
  }

  if !expectedNumSeps.Equal(bINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators ARE NOT Equal!\n"+
      "Because expectedNumSeps != bINumSeps\n"+
      "Expected bINumSeps = '%v'\n"+
      "  Actual bINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_BigInt_02(t *testing.T) {

  ePrefix := "TestBigIntNum_BigInt_02"

  expectedBigINumStr := "-123.456"

  expectedBigINumPrecisionUint := uint(3)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nbStr := "-123456"

  expectedScale := big.NewInt(1000)

  bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)\n"+
      "nbStr= '%v'\n\n",
      ePrefix, nbStr)
    return
  }

  expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

  bINum, err := new(BigIntNum).NewBigInt(bOriginal, expectedBigINumPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewBigInt(\n"+
      "  bOriginal, expectedPrecision)\n"+
      "bOriginal= '%v'\n"+
      "expectedPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bOriginal.Text(10),
      expectedBigINumPrecisionUint,
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

  bINumPrecisionUint, err := bINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  bINumBigInt, err := bINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumBigInt, err := bINum.GetBigInt()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumScaleFactor, err := bINum.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
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

  if bOriginal.Cmp(bINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values ARE NOT Equal!\n"+
      "Because bOriginal.Cmp(bINumBigInt) != 0\n"+
      "Expected bINumBigInt = '%v'\n"+
      "  Actual bINumBigInt = '%v'\n\n",
      ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

    return
  }

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  if expectedScale.Cmp(bINumScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factors ARE NOT Equal!\n"+
      "Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
      "Expected bINumScaleFactor = '%v'\n"+
      "  Actual bINumScaleFactor = '%v'\n\n",
      ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

    return
  }

  if expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Absolute BigInt Values ARE NOT Equal!\n"+
      "Because expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0\n"+
      "Expected bINumAbsBigInt = '%v'\n"+
      "  Actual bINumAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigInt.Text(10), bINumAbsBigInt.Text(10))

    return
  }

  if expectedSignVal != bINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedSignVal != bINumSignValue\n"+
      "Expected bINumSignValue = '%v'\n"+
      "  Actual bINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINumSignValue)

    return
  }

  if expectedBigINumStr != bINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != bINumStr \n"+
      "Expected bINumStr = '%v'\n"+
      "  Actual bINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumStr)

    return
  }

  if !expectedNumSeps.Equal(bINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators ARE NOT Equal!\n"+
      "Because expectedNumSeps != bINumSeps\n"+
      "Expected bINumSeps = '%v'\n"+
      "  Actual bINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_BigInt_03(t *testing.T) {

  ePrefix := "TestBigIntNum_BigInt_03"

  expectedBigINumStr := "0.000123456"

  expectedBigINumPrecisionUint := uint(9)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nbI64 := int64(123456)

  expectedScale := big.NewInt(1000000000)

  bOriginal := big.NewInt(nbI64)

  expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

  bINum, err := new(BigIntNum).NewBigInt(bOriginal, expectedBigINumPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewBigInt(\n"+
      "  bOriginal, expectedPrecision)\n"+
      "bOriginal= '%v'\n"+
      "expectedPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bOriginal.Text(10),
      expectedBigINumPrecisionUint,
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

  bINumPrecisionUint, err := bINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  bINumBigInt, err := bINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumBigInt, err := bINum.GetBigInt()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumScaleFactor, err := bINum.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
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

  if bOriginal.Cmp(bINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values ARE NOT Equal!\n"+
      "Because bOriginal.Cmp(bINumBigInt) != 0\n"+
      "Expected bINumBigInt = '%v'\n"+
      "  Actual bINumBigInt = '%v'\n\n",
      ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

    return
  }

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  if expectedScale.Cmp(bINumScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factors ARE NOT Equal!\n"+
      "Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
      "Expected bINumScaleFactor = '%v'\n"+
      "  Actual bINumScaleFactor = '%v'\n\n",
      ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

    return
  }

  if expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Absolute BigInt Values ARE NOT Equal!\n"+
      "Because expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0\n"+
      "Expected bINumAbsBigInt = '%v'\n"+
      "  Actual bINumAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigInt.Text(10), bINumAbsBigInt.Text(10))

    return
  }

  if expectedSignVal != bINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedSignVal != bINumSignValue\n"+
      "Expected bINumSignValue = '%v'\n"+
      "  Actual bINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINumSignValue)

    return
  }

  if expectedBigINumStr != bINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != bINumStr \n"+
      "Expected bINumStr = '%v'\n"+
      "  Actual bINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumStr)

    return
  }

  if !expectedNumSeps.Equal(bINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators ARE NOT Equal!\n"+
      "Because expectedNumSeps != bINumSeps\n"+
      "Expected bINumSeps = '%v'\n"+
      "  Actual bINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_BigInt_04(t *testing.T) {

  ePrefix := "TestBigIntNum_BigInt_04"

  expectedBigINumStr := "-0.000123456"

  expectedBigINumPrecisionUint := uint(9)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nbI64 := int64(-123456)

  expectedScale := big.NewInt(1000000000)

  bOriginal := big.NewInt(nbI64)

  expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

  bINum, err := new(BigIntNum).NewBigInt(bOriginal, expectedBigINumPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewBigInt(\n"+
      "  bOriginal, expectedPrecision)\n"+
      "bOriginal= '%v'\n"+
      "expectedPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bOriginal.Text(10),
      expectedBigINumPrecisionUint,
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

  bINumPrecisionUint, err := bINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  bINumBigInt, err := bINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumBigInt, err := bINum.GetBigInt()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumScaleFactor, err := bINum.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAbsBigInt, err := bINum.GetAbsoluteBigIntValue()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
    return
  }

  bINumSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumStr, err.Error())
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

  if bOriginal.Cmp(bINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values ARE NOT Equal!\n"+
      "Because bOriginal.Cmp(bINumBigInt) != 0\n"+
      "Expected bINumBigInt = '%v'\n"+
      "  Actual bINumBigInt = '%v'\n\n",
      ePrefix, bOriginal.Text(10), bINumBigInt.Text(10))

    return
  }

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  if expectedScale.Cmp(bINumScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factors ARE NOT Equal!\n"+
      "Because expectedScale.Cmp(bINumScaleFactor) != 0\n"+
      "Expected bINumScaleFactor = '%v'\n"+
      "  Actual bINumScaleFactor = '%v'\n\n",
      ePrefix, expectedScale.Text(10), bINumScaleFactor.Text(10))

    return
  }

  if expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Absolute BigInt Values ARE NOT Equal!\n"+
      "Because expectedAbsBigInt.Cmp(bINumAbsBigInt) != 0\n"+
      "Expected bINumAbsBigInt = '%v'\n"+
      "  Actual bINumAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigInt.Text(10), bINumAbsBigInt.Text(10))

    return
  }

  if expectedSignVal != bINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedSignVal != bINumSignValue\n"+
      "Expected bINumSignValue = '%v'\n"+
      "  Actual bINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINumSignValue)

    return
  }

  if expectedBigINumStr != bINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != bINumStr \n"+
      "Expected bINumStr = '%v'\n"+
      "  Actual bINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumStr)

    return
  }

  if !expectedNumSeps.Equal(bINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators ARE NOT Equal!\n"+
      "Because expectedNumSeps != bINumSeps\n"+
      "Expected bINumSeps = '%v'\n"+
      "  Actual bINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_01(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_01"

  originalNumStr := "5.95"

  expectedBigINumStr := "6"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_02(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_02"

  originalNumStr := "5.05"

  expectedBigINumStr := "6"

  expectedBigINumPrecisionUint := uint(0)

  expectedBigINumSign := 1

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_03(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_03"

  originalNumStr := "5"

  expectedBigINumStr := "5"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_04(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_04"

  originalNumStr := "-5.05"

  expectedBigINumStr := "-5"

  expectedBigINumSign := -1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_05(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_05"

  originalNumStr := "2.4"

  expectedBigINumStr := "3"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_06(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_06"

  originalNumStr := "2.9"

  expectedBigINumStr := "3"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_07(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_07"

  originalNumStr := "-2.7"

  expectedBigINumStr := "-2"

  expectedBigINumSign := -1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_08(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_08"

  originalNumStr := "-2"

  expectedBigINumStr := "-2"

  expectedBigINumSign := -1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_09(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_09"

  originalNumStr := "0"

  expectedBigINumStr := "0"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_10(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_10"

  originalNumStr := "0.00000"

  expectedBigINumStr := "0"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_11(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_11"

  originalNumStr := "159876231.9999999999"

  expectedBigINumStr := "159876232"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_12(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_12"

  originalNumStr := "-159876231.9999999999"

  expectedBigINumStr := "-159876231"

  expectedBigINumSign := -1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_13(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_13"

  originalNumStr := "159876231.0000000000000001"

  expectedBigINumStr := "159876232"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_14(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_14"

  originalNumStr := "-159876231.0000000000000001"

  expectedBigINumStr := "-159876231"

  expectedBigINumSign := -1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_15(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_15"

  originalNumStr := "-0.0000000000000001"

  expectedBigINumStr := "0"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_16(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_16"

  originalNumStr := "0.0000000000000001"

  expectedBigINumStr := "1"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Ceil_17(t *testing.T) {

  ePrefix := "TestBigIntNum_Ceil_17"

  originalNumStr := "-0.0000000000000001"

  expectedBigINumStr := "0"

  expectedBigINumSign := 1

  expectedBigINumPrecisionUint := uint(0)

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bINumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeiling, err := bINum.Ceiling()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeiling, err := bINum.Ceiling()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = bINumCeiling.IsValid("Validating bINumCeiling")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINumCeiling')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumCeilingNumStr, err := bINumCeiling.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumStr, err := bINumCeiling.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumCeilingPrecisionUint, err := bINumCeiling.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingPrecisionUint, err :=\n"+
      "  bINumCeiling.GetPrecisionUint()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingSignValue, err := bINumCeiling.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingSignValue, err := bINumCeiling.GetSign()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  bINumCeilingNumSeps, err := bINumCeiling.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumCeilingNumSeps, err := \n"+
      "  bINumCeiling.GetNumericSeparatorsDto()\n"+
      "bINumCeiling= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumCeilingNumStr, err.Error())
    return
  }

  if expectedBigINumStr != bINumCeilingNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumCeilingNumStr\n"+
      "Expected bINumCeilingNumStr = '%v'\n"+
      "  Actual bINumCeilingNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumCeilingNumStr)

    return
  }

  if expectedBigINumPrecisionUint != bINumCeilingPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected bINumCeilingPrecisionUint = '%v'\n"+
      "  Actual bINumCeilingPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumCeilingPrecisionUint)

    return
  }

  if expectedBigINumSign != bINumCeilingSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumCeilingSignValue\n"+
      "Expected bINumCeilingSignValue = '%v'\n"+
      "  Actual bINumCeilingSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumCeilingSignValue)

    return
  }

  if !expectedNumSeps.Equal(bINumCeilingNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected bINumCeilingNumSeps = '%v'\n"+
      "  Actual bINumCeilingNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumCeilingNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_ChangeSign_01(t *testing.T) {

  ePrefix := "TestBigIntNum_ChangeSign_01"

  originalNumStr := "123.456"

  expectedBigINumStr := "-123.456"

  originalNumSign := 1

  expectedBigINumSign := -1

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bIOriginalNumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bIOriginalNumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumOriginalSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumOriginalSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bIOriginalNumStr, err.Error())
    return
  }

  err = bINum.ChangeSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bIOriginalNumStr, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Sign Change")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Sign Change')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNewSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNewSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bIOriginalNumStr, err.Error())
    return
  }

  bINumNewNumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNewNumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumNewNumSeps, err := bINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNewNumSeps, err :=\n"+
      "  bINum.GetNumericSeparatorsDto()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumNewNumStr, err.Error())
    return
  }

  if originalNumSign != bINumOriginalSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumSign != bINumOriginalSignValue\n"+
      "Expected bINumOriginalSignValue = '%v'\n"+
      "  Actual bINumOriginalSignValue = '%v'\n"+
      "Original bINumStr= '%v'\n\n",
      ePrefix, originalNumSign, bINumOriginalSignValue, bIOriginalNumStr)

    return
  }

  if expectedBigINumSign != bINumNewSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumNewSignValue\n"+
      "Expected bINumNewSignValue = '%v'\n"+
      "  Actual bINumNewSignValue = '%v'\n"+
      "New/Final bINumStr= '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumNewSignValue, bINumNewNumStr)

    return
  }

  if originalNumStr != bIOriginalNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumStr != bIOriginalNumStr \n"+
      "Expected bIOriginalNumStr = '%v'\n"+
      "  Actual bIOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bIOriginalNumStr)

    return
  }

  if expectedBigINumStr != bINumNewNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumNewNumStr\n"+
      "Expected bINumNewNumStr = '%v'\n"+
      "  Actual bINumNewNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumNewNumStr)

    return
  }

  if !expectedNumSeps.Equal(bINumNewNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != bINumNewNumSeps \n"+
      "Expected bINumNewNumSeps = '%v'\n"+
      "  Actual bINumNewNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumNewNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_ChangeSign_02(t *testing.T) {

  ePrefix := "TestBigIntNum_ChangeSign_02"

  originalNumStr := "-123.456"

  expectedBigINumStr := "123.456"

  originalNumSign := -1

  expectedBigINumSign := 1

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bIOriginalNumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bIOriginalNumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumOriginalSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumOriginalSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bIOriginalNumStr, err.Error())
    return
  }

  err = bINum.ChangeSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bIOriginalNumStr, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Sign Change")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Sign Change')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNewSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNewSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bIOriginalNumStr, err.Error())
    return
  }

  bINumNewNumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNewNumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumNewNumSeps, err := bINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNewNumSeps, err :=\n"+
      "  bINum.GetNumericSeparatorsDto()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumNewNumStr, err.Error())
    return
  }

  if originalNumSign != bINumOriginalSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumSign != bINumOriginalSignValue\n"+
      "Expected bINumOriginalSignValue = '%v'\n"+
      "  Actual bINumOriginalSignValue = '%v'\n"+
      "Original bINumStr= '%v'\n\n",
      ePrefix, originalNumSign, bINumOriginalSignValue, bIOriginalNumStr)

    return
  }

  if expectedBigINumSign != bINumNewSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumNewSignValue\n"+
      "Expected bINumNewSignValue = '%v'\n"+
      "  Actual bINumNewSignValue = '%v'\n"+
      "New/Final bINumStr= '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumNewSignValue, bINumNewNumStr)

    return
  }

  if originalNumStr != bIOriginalNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumStr != bIOriginalNumStr \n"+
      "Expected bIOriginalNumStr = '%v'\n"+
      "  Actual bIOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bIOriginalNumStr)

    return
  }

  if expectedBigINumStr != bINumNewNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumNewNumStr\n"+
      "Expected bINumNewNumStr = '%v'\n"+
      "  Actual bINumNewNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumNewNumStr)

    return
  }

  if !expectedNumSeps.Equal(bINumNewNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != bINumNewNumSeps \n"+
      "Expected bINumNewNumSeps = '%v'\n"+
      "  Actual bINumNewNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumNewNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_ChangeSign_03(t *testing.T) {

  ePrefix := "TestBigIntNum_ChangeSign_03"

  originalNumStr := "0.00"

  expectedBigINumStr := "0.00"

  originalNumSign := 1

  expectedBigINumSign := 1

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

  bIOriginalNumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bIOriginalNumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumOriginalSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumOriginalSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bIOriginalNumStr, err.Error())
    return
  }

  err = bINum.ChangeSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bIOriginalNumStr, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Sign Change")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Sign Change')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNewSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNewSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bIOriginalNumStr, err.Error())
    return
  }

  bINumNewNumStr, err := bINum.GetNumStr()

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNewNumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  bINumNewNumSeps, err := bINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNewNumSeps, err :=\n"+
      "  bINum.GetNumericSeparatorsDto()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, bINumNewNumStr, err.Error())
    return
  }

  if originalNumSign != bINumOriginalSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumSign != bINumOriginalSignValue\n"+
      "Expected bINumOriginalSignValue = '%v'\n"+
      "  Actual bINumOriginalSignValue = '%v'\n"+
      "Original bINumStr= '%v'\n\n",
      ePrefix, originalNumSign, bINumOriginalSignValue, bIOriginalNumStr)

    return
  }

  if expectedBigINumSign != bINumNewSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != bINumNewSignValue\n"+
      "Expected bINumNewSignValue = '%v'\n"+
      "  Actual bINumNewSignValue = '%v'\n"+
      "New/Final bINumStr= '%v'\n\n",
      ePrefix, expectedBigINumSign, bINumNewSignValue, bINumNewNumStr)

    return
  }

  if originalNumStr != bIOriginalNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumStr != bIOriginalNumStr \n"+
      "Expected bIOriginalNumStr = '%v'\n"+
      "  Actual bIOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bIOriginalNumStr)

    return
  }

  if expectedBigINumStr != bINumNewNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumStr != bINumNewNumStr\n"+
      "Expected bINumNewNumStr = '%v'\n"+
      "  Actual bINumNewNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, bINumNewNumStr)

    return
  }

  if !expectedNumSeps.Equal(bINumNewNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != bINumNewNumSeps \n"+
      "Expected bINumNewNumSeps = '%v'\n"+
      "  Actual bINumNewNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumNewNumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Cmp_01(t *testing.T) {

  n1Str := "123.456"
  n2Str := "123.455"
  expectedCmpResult := 1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_02(t *testing.T) {

  n1Str := "123.456"
  n2Str := "123.457"
  expectedCmpResult := -1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_03(t *testing.T) {

  n1Str := "123.456"
  n2Str := "123.456"
  expectedCmpResult := 0

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_04(t *testing.T) {

  n1Str := "-123.456"
  n2Str := "-123.457"
  expectedCmpResult := 1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_05(t *testing.T) {

  n1Str := "-123.456"
  n2Str := "-123.455"
  expectedCmpResult := -1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_06(t *testing.T) {

  n1Str := "-123.456"
  n2Str := "-123.456"
  expectedCmpResult := 0

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_07(t *testing.T) {

  n1Str := "123456000643218"
  n2Str := "123456000643217"
  expectedCmpResult := 1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_08(t *testing.T) {

  n1Str := "123456000643218"
  n2Str := "123456000643219"
  expectedCmpResult := -1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_09(t *testing.T) {

  n1Str := "123456000643218"
  n2Str := "123456000643218"
  expectedCmpResult := 0

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_10(t *testing.T) {

  n1Str := "-123456000643218"
  n2Str := "-123456000643217"
  expectedCmpResult := -1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_11(t *testing.T) {

  n1Str := "-123456000643218"
  n2Str := "-123456000643219"
  expectedCmpResult := 1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_12(t *testing.T) {

  n1Str := "-123456000643218"
  n2Str := "-123456000643218"
  expectedCmpResult := 0

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_13(t *testing.T) {

  n1Str := "1.7"
  n2Str := "-12345.6000643218"
  expectedCmpResult := 1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_14(t *testing.T) {

  n1Str := "-1.7"
  n2Str := "0.01"
  expectedCmpResult := -1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_15(t *testing.T) {

  n1Str := "17"
  n2Str := "-1"
  expectedCmpResult := 1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_16(t *testing.T) {

  n1Str := "-17"
  n2Str := "1"
  expectedCmpResult := -1

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Cmp_17(t *testing.T) {

  n1Str := "0"
  n2Str := "0.000"
  expectedCmpResult := 0

  bINum1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'", n1Str, err.Error())
  }

  bINum2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'", n2Str, err.Error())
  }

  cmpResult := bINum1.Cmp(bINum2)

  if expectedCmpResult != cmpResult {
    t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
      expectedCmpResult, cmpResult)
  }

}

func TestBigIntNum_Decimal_01(t *testing.T) {

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

  bINum, err := new(BigIntNum).NewDecimal(dec)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewDecimal(dec) "+
      "Error='%v' ", err.Error())
  }

  nDto, err := new(NumStrDto).NewBigInt(bINum.bigInt, bINum.precision)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewBigInt(bINum.bigInt, bINum.precision) "+
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

func TestBigIntNum_Decimal_02(t *testing.T) {

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

  dec, err := Decimal{}.NewNumStr(nStr)

  bINum, err := new(BigIntNum).NewDecimal(dec)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewIntAry(ia) "+
      "Error='%v' ", err.Error())
  }

  nDto, err := new(NumStrDto).NewBigInt(bINum.bigInt, bINum.precision)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewBigInt(bINum.bigInt, bINum.precision) "+
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

func TestBigIntNum_Decrement_01(t *testing.T) {

  numStr := "8"
  expectedNumStr := "7"
  expectedNumSeps := new(NumericSeparatorDto).New()

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
      "numStr='%v' Error='%v' \n", numStr, err.Error())
  }

  bINum.Decrement()

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_Decrement_02(t *testing.T) {

  numStr := "8.2"
  expectedNumStr := "7.2"
  expectedNumSeps := new(NumericSeparatorDto).New()

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
      "numStr='%v' Error='%v' \n", numStr, err.Error())
  }

  bINum.Decrement()

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_Decrement_03(t *testing.T) {

  numStr := "-8.2"
  expectedNumStr := "-9.2"
  expectedNumSeps := new(NumericSeparatorDto).New()

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
      "numStr='%v' Error='%v' \n", numStr, err.Error())
  }

  bINum.Decrement()

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := bINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntNum_Decrement_04(t *testing.T) {

  numStr := "8"
  expectedNumStr := "7"

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
      "numStr='%v' Error='%v' \n", numStr, err.Error())
  }

  bINum.Decrement()

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := bINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntNum_DivideByTwo_01(t *testing.T) {

  numStr := "658.78562347"
  // expectedNumStr := "329.392811735"
  expectedQuoStr := "329"
  expectedModStr := "0.78562347"

  maxPrecision := uint(10)

  base, err := new(BigIntNum).NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)

  if err != nil {
    t.Errorf("Error returned by base.DivideByTwoQuoMod(maxPrecision). Error='%v' ",
      err.Error())
  }

  actualNumStr := quotient.GetNumStr()

  if expectedQuoStr != actualNumStr {
    t.Errorf("Error: Expected quotient NumStr='%v'. Instead, quotient NumStr='%v'. ",
      expectedQuoStr, actualNumStr)
  }

  actualNumStr = modulo.GetNumStr()

  if expectedModStr != actualNumStr {
    t.Errorf("Error: Expected modulo NumStr='%v'. Instead, modulo NumStr='%v'. ",
      expectedModStr, actualNumStr)
  }
}

func TestBigIntNum_DivideByTwo_02(t *testing.T) {

  numStr := "8"

  expectedQuoStr := "4"
  expectedModStr := "0"

  maxPrecision := uint(10)

  base, err := new(BigIntNum).NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)

  if err != nil {
    t.Errorf("Error returned by base.DivideByTwoQuoMod(maxPrecision). Error='%v' ",
      err.Error())
  }

  actualNumStr := quotient.GetNumStr()

  if expectedQuoStr != actualNumStr {
    t.Errorf("Error: Expected quotient NumStr='%v'. Instead, quotient NumStr='%v'. ",
      expectedQuoStr, actualNumStr)
  }

  actualNumStr = modulo.GetNumStr()

  if expectedModStr != actualNumStr {
    t.Errorf("Error: Expected modulo NumStr='%v'. Instead, modulo NumStr='%v'. ",
      expectedModStr, actualNumStr)
  }
}

func TestBigIntNum_DivideByTenToPower_01(t *testing.T) {
  num1Str := "654.123"
  expectedStr := "0.654123"
  exponent := uint(3)

  bINum, err := new(BigIntNum).NewNumStr(num1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  bINum.DivideByTenToPower(exponent)

  actualStr := bINum.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedStr, actualStr)
  }

}

func TestBigIntNum_DivideByTenToPower_02(t *testing.T) {
  num1Str := "-654.123"
  expectedStr := "-0.654123"
  exponent := uint(3)

  bINum, err := new(BigIntNum).NewNumStr(num1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  bINum.DivideByTenToPower(exponent)

  actualStr := bINum.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedStr, actualStr)
  }

}

func TestBigIntNum_DivideByTenToPower_03(t *testing.T) {
  num1Str := "654123"
  expectedStr := "0.000654123"
  exponent := uint(9)

  bINum, err := new(BigIntNum).NewNumStr(num1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  bINum.DivideByTenToPower(exponent)

  actualStr := bINum.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedStr, actualStr)
  }

}

func TestBigIntNum_DivideByTenToPower_04(t *testing.T) {
  num1Str := "654123"
  expectedStr := "654123"
  exponent := uint(0)

  bINum, err := new(BigIntNum).NewNumStr(num1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  bINum.DivideByTenToPower(exponent)

  actualStr := bINum.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedStr, actualStr)
  }

}

func TestBigIntNum_DivideByTenToPower_05(t *testing.T) {
  num1Str := "-654123"
  expectedStr := "-654123"
  exponent := uint(0)

  bINum, err := new(BigIntNum).NewNumStr(num1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  bINum.DivideByTenToPower(exponent)

  actualStr := bINum.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedStr, actualStr)
  }

}

func TestBigIntNum_ExtendPrecision_01(t *testing.T) {

  num1Str := "654.123"
  expectedNumStr := "654.12300"

  bNum1, err := new(BigIntNum).NewNumStr(num1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  expectedNum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedNumStr). Error='%v' ",
      err.Error())
  }

  bNum1.ExtendPrecision(2)

  if !expectedNum.Equal(bNum1) {
    t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
      expectedNum.GetNumStr(), bNum1.GetNumStr())
  }

}

func TestBigIntNum_ExtendPrecision_02(t *testing.T) {

  num1Str := "-654.123"
  expectedNumStr := "-654.12300"

  bNum1, err := new(BigIntNum).NewNumStr(num1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  expectedNum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedNumStr). Error='%v' ",
      err.Error())
  }

  bNum1.ExtendPrecision(2)

  if !expectedNum.Equal(bNum1) {
    t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
      expectedNum.GetNumStr(), bNum1.GetNumStr())
  }

}

func TestBigIntNum_ExtendPrecision_03(t *testing.T) {

  num1Str := "7"
  expectedNumStr := "7.00"

  bNum1, err := new(BigIntNum).NewNumStr(num1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  expectedNum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedNumStr). Error='%v' ",
      err.Error())
  }

  bNum1.ExtendPrecision(2)

  if !expectedNum.Equal(bNum1) {
    t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
      expectedNum.GetNumStr(), bNum1.GetNumStr())
  }

}

func TestBigIntNum_ExtendPrecision_04(t *testing.T) {

  num1Str := "0"
  expectedNumStr := "0.00"

  bNum1, err := new(BigIntNum).NewNumStr(num1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(num1Str). Error='%v' ",
      err.Error())
  }

  expectedNum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedNumStr). Error='%v' ",
      err.Error())
  }

  bNum1.ExtendPrecision(2)

  if !expectedNum.Equal(bNum1) {
    t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
      expectedNum.GetNumStr(), bNum1.GetNumStr())
  }

}
