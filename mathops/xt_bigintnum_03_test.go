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

  nStr := "123.456"
  expectedNumStr := "0.456"
  expectedPrecision := uint(3)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  fractionalPart := bINum1.GetFractionalPart()

  actualNumStr := fractionalPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Fractional Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != fractionalPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Fractional Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, fractionalPart.GetPrecisionUint())
  }
}

func TestBigIntNum_GetFractionalPart_02(t *testing.T) {

  nStr := "-123.456"
  expectedNumStr := "-0.456"
  expectedPrecision := uint(3)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  fractionalPart := bINum1.GetFractionalPart()

  actualNumStr := fractionalPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Fractional Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != fractionalPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Fractional Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, fractionalPart.GetPrecisionUint())
  }
}

func TestBigIntNum_GetFractionalPart_03(t *testing.T) {

  nStr := "123"
  expectedNumStr := "0"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  fractionalPart := bINum1.GetFractionalPart()

  actualNumStr := fractionalPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Fractional Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != fractionalPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Fractional Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, fractionalPart.GetPrecisionUint())
  }
}

func TestBigIntNum_GetFractionalPart_04(t *testing.T) {

  nStr := "-123"
  expectedNumStr := "0"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  fractionalPart := bINum1.GetFractionalPart()

  actualNumStr := fractionalPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Fractional Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != fractionalPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Fractional Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, fractionalPart.GetPrecisionUint())
  }
}

func TestBigIntNum_GetFractionalPart_05(t *testing.T) {

  nStr := "0.000"
  expectedNumStr := "0"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  fractionalPart := bINum1.GetFractionalPart()

  actualNumStr := fractionalPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Fractional Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != fractionalPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Fractional Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, fractionalPart.GetPrecisionUint())
  }
}

func TestBigIntNum_GetFractionalPart_06(t *testing.T) {

  nStr := "0"
  expectedNumStr := "0"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  fractionalPart := bINum1.GetFractionalPart()

  actualNumStr := fractionalPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Fractional Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != fractionalPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Fractional Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, fractionalPart.GetPrecisionUint())
  }
}

func TestBigIntNum_GetDecimal_01(t *testing.T) {
  expectedStr := "-847921684.347"
  ePrecision := uint(3)

  bINum, err := new(BigIntNum).NewNumStr(expectedStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedStr). Error='%v'", err.Error())
  }

  decActual, err := bINum.GetDecimal()

  if expectedStr != decActual.GetNumStr() {
    t.Errorf("Error: Expected decActual.GetNumStr()='%v'. "+
      "Instead, decActual.GetNumStr()='%v'.",
      expectedStr, decActual.GetNumStr())
  }

  if ePrecision != decActual.GetPrecisionUint() {
    t.Errorf("Error: Expected decActual.GetPrecisionUint()='%v'. "+
      "Instead, decActual.GetPrecisionUint()='%v'.",
      ePrecision, decActual.GetPrecisionUint())
  }

}

func TestBigIntNum_GetNumStrDto_01(t *testing.T) {
  expectedStr := "-847921684.347"
  ePrecision := uint(3)

  bINum, err := new(BigIntNum).NewNumStr(expectedStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedStr). Error='%v'", err.Error())
  }

  numStrDto, err := bINum.GetNumStrDto()

  if expectedStr != numStrDto.GetNumStr() {
    t.Errorf("Error: Expected numStrDto.GetNumStr()='%v'. "+
      "Instead, numStrDto.GetNumStr()='%v'.",
      expectedStr, numStrDto.GetNumStr())
  }

  if ePrecision != numStrDto.GetPrecisionUint() {
    t.Errorf("Error: Expected numStrDto.GetPrecisionUint()='%v'. "+
      "Instead, numStrDto.GetPrecisionUint()='%v'.",
      ePrecision, numStrDto.GetPrecisionUint())
  }

}

func TestBigIntNum_GetIntAry_01(t *testing.T) {
  expectedStr := "-847921684.347"
  ePrecision := uint(3)

  bINum, err := new(BigIntNum).NewNumStr(expectedStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedStr). Error='%v'", err.Error())
  }

  intAry, err := bINum.GetIntAry()

  if err != nil {
    t.Errorf("Error returned by bINum.GetIntAry(). Error='%v'", err.Error())
  }

  if expectedStr != intAry.GetNumStr() {
    t.Errorf("Error: Expected intAry.GetNumStr()='%v'. "+
      "Instead, intAry.GetNumStr()='%v'.",
      expectedStr, intAry.GetNumStr())
  }

  if ePrecision != intAry.GetPrecisionUint() {
    t.Errorf("Error: Expected intAry.GetPrecisionUint()='%v'. "+
      "Instead, intAry.GetPrecisionUint()='%v'.",
      ePrecision, intAry.GetPrecisionUint())
  }

}

func TestBigIntNum_GetIntegerPart_01(t *testing.T) {

  nStr := "123.456"
  expectedNumStr := "123"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  integerPart := bINum1.GetIntegerPart()

  actualNumStr := integerPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Integer Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != integerPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Integer Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, integerPart.GetPrecisionUint())
  }

}

func TestBigIntNum_GetIntegerPart_02(t *testing.T) {

  nStr := "-123.456"
  expectedNumStr := "-123"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  integerPart := bINum1.GetIntegerPart()

  actualNumStr := integerPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Integer Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != integerPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Integer Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, integerPart.GetPrecisionUint())
  }

}

func TestBigIntNum_GetIntegerPart_03(t *testing.T) {

  nStr := "123"
  expectedNumStr := "123"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  integerPart := bINum1.GetIntegerPart()

  actualNumStr := integerPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Integer Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != integerPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Integer Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, integerPart.GetPrecisionUint())
  }

}

func TestBigIntNum_GetIntegerPart_04(t *testing.T) {

  nStr := "-123"
  expectedNumStr := "-123"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  integerPart := bINum1.GetIntegerPart()

  actualNumStr := integerPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Integer Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != integerPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Integer Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, integerPart.GetPrecisionUint())
  }

}

func TestBigIntNum_GetIntegerPart_05(t *testing.T) {

  nStr := "0.000"
  expectedNumStr := "0"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  integerPart := bINum1.GetIntegerPart()

  actualNumStr := integerPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Integer Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != integerPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Integer Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, integerPart.GetPrecisionUint())
  }

}

func TestBigIntNum_GetIntegerPart_06(t *testing.T) {

  nStr := "0"
  expectedNumStr := "0"
  expectedPrecision := uint(0)

  bINum1, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  integerPart := bINum1.GetIntegerPart()

  actualNumStr := integerPart.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected Integer Part NumStr='%v'. "+
      "Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != integerPart.GetPrecisionUint() {
    t.Errorf("Error: Expected Integer Part precision='%v' "+
      "Instead, precision='%v'",
      expectedPrecision, integerPart.GetPrecisionUint())
  }

}

func TestBigIntNum_GetNumberOfDigits_01(t *testing.T) {

  nStr := "123.45"
  expectedDigitCnt := 5
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_02(t *testing.T) {

  nStr := "1,234,567"
  expectedDigitCnt := 7
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_03(t *testing.T) {

  nStr := "-1,234,567.8"
  expectedDigitCnt := 8
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_04(t *testing.T) {

  nStr := "0"
  expectedDigitCnt := 1
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_05(t *testing.T) {

  //nStr := "0.00"
  expectedDigitCnt := 1
  bINum := new(BigIntNum).NewZero(2)

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_06(t *testing.T) {

  nStr := "012.34"
  expectedDigitCnt := 4
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_07(t *testing.T) {

  nStr := "0.1234"
  expectedDigitCnt := 4
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_08(t *testing.T) {

  nStr := "0.123400"
  expectedDigitCnt := 6
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_09(t *testing.T) {

  nStr := "0.0123400"
  expectedDigitCnt := 6
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_10(t *testing.T) {

  nStr := "1,234,567.800"
  expectedDigitCnt := 10
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_GetNumberOfDigits_11(t *testing.T) {

  nStr := "5"
  expectedDigitCnt := 1
  bINum, err := new(BigIntNum).NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(nStr). "+
      "Error='%v' ", err.Error())
  }

  actualDigitCnt := bINum.GetNumberOfDigits()

  if expectedDigitCnt != actualDigitCnt {
    t.Errorf("Error: Expected Digit Count='%v'. Instead, Digit Count='%v' ",
      expectedDigitCnt, actualDigitCnt)
  }

}

func TestBigIntNum_Increment_01(t *testing.T) {

  numStr := "8"
  expectedNumStr := "9"
  expectedNumSeps := new(NumericSeparatorDto).New()

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
      "numStr='%v' Error='%v' \n", numStr, err.Error())
  }

  bINum.Increment()

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_Increment_02(t *testing.T) {

  numStr := "8.2"
  expectedNumStr := "9.2"
  expectedNumSeps := new(NumericSeparatorDto).New()

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
      "numStr='%v' Error='%v' \n", numStr, err.Error())
  }

  bINum.Increment()

  actualNumStr := bINum.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_Increment_03(t *testing.T) {

  numStr := "-8.2"
  expectedNumStr := "-7.2"
  expectedNumSeps := new(NumericSeparatorDto).New()

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
      "numStr='%v' Error='%v' \n", numStr, err.Error())
  }

  bINum.Increment()

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

func TestBigIntNum_Increment_04(t *testing.T) {

  numStr := "8"
  expectedNumStr := "9"

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

  bINum.Increment()

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

func TestBigIntNum_Inverse_01(t *testing.T) {
  testNumStr := "4"
  expectedNumStr := "0.25"
  maxPrecision := uint(2)

  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' \n", testNumStr, err.Error())
  }

  actualInverse, err := bINum.Inverse(maxPrecision)

  if err != nil {
    t.Errorf("Error returned by bINum.Inverse(maxPrecision). "+
      "maxPrecision='%v' Error='%v' \n", maxPrecision, err.Error())
  }

  if expectedNumStr != actualInverse.GetNumStr() {
    t.Errorf("Error: Expected Inverse='%v'. Instead, Inverse='%v' \n",
      expectedNumStr, actualInverse.GetNumStr())
  }
}

func TestBigIntNum_Inverse_02(t *testing.T) {
  testNumStr := "9.357"
  expectedNumStr := "0.10687186063909372662178048519825"
  maxPrecision := uint(32)

  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' \n", testNumStr, err.Error())
  }

  actualInverse, err := bINum.Inverse(maxPrecision)

  if err != nil {
    t.Errorf("Error returned by bINum.Inverse(maxPrecision). "+
      "maxPrecision='%v' Error='%v' \n", maxPrecision, err.Error())
  }

  if expectedNumStr != actualInverse.GetNumStr() {
    t.Errorf("Error: Expected Inverse='%v'. Instead, Inverse='%v' \n",
      expectedNumStr, actualInverse.GetNumStr())
  }
}

func TestBigIntNum_Inverse_03(t *testing.T) {
  testNumStr := "-9.357"
  expectedNumStr := "-0.10687186063909372662178048519825"
  maxPrecision := uint(32)

  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' \n", testNumStr, err.Error())
  }

  actualInverse, err := bINum.Inverse(maxPrecision)

  if err != nil {
    t.Errorf("Error returned by bINum.Inverse(maxPrecision). "+
      "maxPrecision='%v' Error='%v' \n", maxPrecision, err.Error())
  }

  if expectedNumStr != actualInverse.GetNumStr() {
    t.Errorf("Error: Expected Inverse='%v'. Instead, Inverse='%v' \n",
      expectedNumStr, actualInverse.GetNumStr())
  }
}

func TestBigIntNum_Inverse_04(t *testing.T) {
  testNumStr := "0"
  expectedNumStr := "0"
  maxPrecision := uint(0)

  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' \n", testNumStr, err.Error())
  }

  actualInverse, err := bINum.Inverse(maxPrecision)

  if err != nil {
    t.Errorf("Error returned by bINum.Inverse(maxPrecision). "+
      "maxPrecision='%v' Error='%v' \n", maxPrecision, err.Error())
  }

  if expectedNumStr != actualInverse.GetNumStr() {
    t.Errorf("Error: Expected Inverse='%v'. Instead, Inverse='%v' \n",
      expectedNumStr, actualInverse.GetNumStr())
  }
}

func TestBigIntNum_Inverse_09(t *testing.T) {
  testNumStr := "9"
  expectedNumStr := "0.11111111111111111111111111111111"
  maxPrecision := uint(32)

  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' \n", testNumStr, err.Error())
  }

  actualInverse, err := bINum.Inverse(maxPrecision)

  if err != nil {
    t.Errorf("Error returned by bINum.Inverse(maxPrecision). "+
      "maxPrecision='%v' Error='%v' \n", maxPrecision, err.Error())
  }

  if expectedNumStr != actualInverse.GetNumStr() {
    t.Errorf("Error: Expected Inverse='%v'. Instead, Inverse='%v' \n",
      expectedNumStr, actualInverse.GetNumStr())
  }
}

func TestBigIntNum_Inverse_10(t *testing.T) {
  testNumStr := "-9"
  expectedNumStr := "-0.11111111111111111111111111111111"
  maxPrecision := uint(32)

  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' \n", testNumStr, err.Error())
  }

  actualInverse, err := bINum.Inverse(maxPrecision)

  if err != nil {
    t.Errorf("Error returned by bINum.Inverse(maxPrecision). "+
      "maxPrecision='%v' Error='%v' \n", maxPrecision, err.Error())
  }

  if expectedNumStr != actualInverse.GetNumStr() {
    t.Errorf("Error: Expected Inverse='%v'. Instead, Inverse='%v' \n",
      expectedNumStr, actualInverse.GetNumStr())
  }
}

func TestBigIntNum_IsEvenNumber_01(t *testing.T) {

  testNumStr := "4"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if true != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      true, isEven)
  }

}

func TestBigIntNum_IsEvenNumber_02(t *testing.T) {

  testNumStr := "5"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if false != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      false, isEven)
  }

}

func TestBigIntNum_IsEvenNumber_03(t *testing.T) {

  testNumStr := "4.4"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if false != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      false, isEven)
  }

}

func TestBigIntNum_IsEvenNumber_04(t *testing.T) {

  testNumStr := "9793442794"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if true != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      true, isEven)
  }

}

func TestBigIntNum_IsEvenNumber_05(t *testing.T) {

  testNumStr := "9793442795"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if false != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      false, isEven)
  }

}

func TestBigIntNum_IsEvenNumber_06(t *testing.T) {

  testNumStr := "-9793442794"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if true != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      true, isEven)
  }

}

func TestBigIntNum_IsEvenNumber_07(t *testing.T) {

  testNumStr := "757849035736836546"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if true != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      true, isEven)
  }

}

func TestBigIntNum_IsEvenNumber_08(t *testing.T) {

  testNumStr := "-757849035736836546"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if true != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      true, isEven)
  }

}

func TestBigIntNum_IsEvenNumber_09(t *testing.T) {

  testNumStr := "757849035736836547"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if false != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      false, isEven)
  }

}

func TestBigIntNum_IsEvenNumber_10(t *testing.T) {

  testNumStr := "-757849035736836547"
  bINum, err := new(BigIntNum).NewNumStr(testNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(testNumStr). "+
      "testNumStr='%v' Error='%v' ", testNumStr, err.Error())
  }

  isEven, err := bINum.IsEvenNumber()

  if err != nil {
    t.Errorf("Error returned by bINum.IsEvenNumber(). "+
      "bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())

  }

  if false != isEven {
    t.Errorf("Expected TestNumber to be IsEven='%v'. Instead TestNumber IsEven='%v'",
      false, isEven)
  }

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
