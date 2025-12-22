package mathops

import (
  "math/big"
  "testing"
)

func TestNumStrDto_SetPrecision_01(t *testing.T) {

  ePrefix := "TestNumStrDto_SetPrecision_01"

  inputNumStr := "123456789"

  inputPrecisionUint := uint(7)

  var inputRoundResult bool

  inputRoundResult = false

  expectedNumStr := "123456789.0000000"

  expectedBigInt := big.NewInt(1234567890000000)

  expectedPrecisionInt := 7

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

  expectedNumHasNumericDigits = true

  expectedNumIsFractionalValue = true

  expectedNumAbsIntStr := "123456789"

  expectedNumAbsFracStr := "0000000"

  expectedAbsAllRunesNumStr := "1234567890000000"

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  numStrDtoResult, err := new(NumStrDto).NewPtr().SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
      "  SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)\n"+
      "inputNumStr= '%v'\n"+
      "inputPrecisionUint= '%v'\n"+
      "inputRoundResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr,
      inputPrecisionUint,
      inputRoundResult,
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

  numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

  numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsIntRunes, err := \n"+
      "  numStrDtoResult.GetAbsIntRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

  numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsFracRunes, err :=\n"+
      "  numStrDtoResult.GetAbsFracRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

  numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

  numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINumNumberStr, err :=\n"+
      "  numStrDtoResultBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
      "expectedBigINum= '%v'\n"+
      "numStrDtoResultBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResultBigINum is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected numStrDtoResultBigINum = '%v'\n"+
      "  Actual numStrDtoResultBigINum = '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
    t.Errorf("%v\n"+
      "Error: Numeric Digits Flag is Invalid!\n"+
      "Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
      "Expected numStrDtoResultHasNumericDigits = '%v'\n"+
      "  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
      ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

    return
  }

  if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
    t.Errorf("%v\n"+
      "Error: IsFractionalValue Flag Invalid!\n"+
      "Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
      "Expected numStrDtoResultIsFractionalValue = '%v'\n"+
      "  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
      ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

    return
  }

  if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
      "Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
      "Expected numStrDtoResultAbsIntStr = '%v'\n"+
      "  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
      ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

    return
  }

  if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected numStrDtoResultAbsFracStr = '%v'\n"+
      "  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
      ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

    return
  }

  if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
      "Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
      "Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
      "  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
      ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factor INVALID!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
      "Expected numStrDtoBigInt = '%v'\n"+
      "  Actual numStrDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

    return
  }

  return
}

func TestNumStrDto_SetPrecision_02(t *testing.T) {

  ePrefix := "TestNumStrDto_SetPrecision_02"

  inputNumStr := "123456789"

  inputPrecisionUint := uint(7)

  var inputRoundResult bool

  inputRoundResult = true

  expectedNumStr := "123456789.0000000"

  expectedBigInt := big.NewInt(1234567890000000)

  expectedPrecisionInt := 7

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

  expectedNumHasNumericDigits = true

  expectedNumIsFractionalValue = true

  expectedNumAbsIntStr := "123456789"

  expectedNumAbsFracStr := "0000000"

  expectedAbsAllRunesNumStr := "1234567890000000"

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  numStrDtoResult, err := new(NumStrDto).NewPtr().SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
      "  SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)\n"+
      "inputNumStr= '%v'\n"+
      "inputPrecisionUint= '%v'\n"+
      "inputRoundResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr,
      inputPrecisionUint,
      inputRoundResult,
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

  numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

  numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsIntRunes, err := \n"+
      "  numStrDtoResult.GetAbsIntRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

  numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsFracRunes, err :=\n"+
      "  numStrDtoResult.GetAbsFracRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

  numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

  numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINumNumberStr, err :=\n"+
      "  numStrDtoResultBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
      "expectedBigINum= '%v'\n"+
      "numStrDtoResultBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResultBigINum is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected numStrDtoResultBigINum = '%v'\n"+
      "  Actual numStrDtoResultBigINum = '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
    t.Errorf("%v\n"+
      "Error: Numeric Digits Flag is Invalid!\n"+
      "Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
      "Expected numStrDtoResultHasNumericDigits = '%v'\n"+
      "  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
      ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

    return
  }

  if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
    t.Errorf("%v\n"+
      "Error: IsFractionalValue Flag Invalid!\n"+
      "Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
      "Expected numStrDtoResultIsFractionalValue = '%v'\n"+
      "  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
      ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

    return
  }

  if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
      "Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
      "Expected numStrDtoResultAbsIntStr = '%v'\n"+
      "  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
      ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

    return
  }

  if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected numStrDtoResultAbsFracStr = '%v'\n"+
      "  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
      ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

    return
  }

  if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
      "Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
      "Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
      "  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
      ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factor INVALID!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
      "Expected numStrDtoBigInt = '%v'\n"+
      "  Actual numStrDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

    return
  }

  return
}

func TestNumStrDto_SetPrecision_03(t *testing.T) {

  ePrefix := "TestNumStrDto_SetPrecision_03"

  inputNumStr := "-123456789"

  inputPrecisionUint := uint(7)

  var inputRoundResult bool

  inputRoundResult = false

  expectedNumStr := "-123456789.0000000"

  expectedBigInt := big.NewInt(-1234567890000000)

  expectedPrecisionInt := 7

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

  expectedNumHasNumericDigits = true

  expectedNumIsFractionalValue = true

  expectedNumAbsIntStr := "123456789"

  expectedNumAbsFracStr := "0000000"

  expectedAbsAllRunesNumStr := "1234567890000000"

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  numStrDtoResult, err := new(NumStrDto).NewPtr().SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewPtr().\n"+
      "  SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)\n"+
      "inputNumStr= '%v'\n"+
      "inputPrecisionUint= '%v'\n"+
      "inputRoundResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr,
      inputPrecisionUint,
      inputRoundResult,
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

  numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

  numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsIntRunes, err := \n"+
      "  numStrDtoResult.GetAbsIntRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

  numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsFracRunes, err :=\n"+
      "  numStrDtoResult.GetAbsFracRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

  numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

  numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINumNumberStr, err :=\n"+
      "  numStrDtoResultBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
      "expectedBigINum= '%v'\n"+
      "numStrDtoResultBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResultBigINum is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected numStrDtoResultBigINum = '%v'\n"+
      "  Actual numStrDtoResultBigINum = '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
    t.Errorf("%v\n"+
      "Error: Numeric Digits Flag is Invalid!\n"+
      "Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
      "Expected numStrDtoResultHasNumericDigits = '%v'\n"+
      "  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
      ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

    return
  }

  if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
    t.Errorf("%v\n"+
      "Error: IsFractionalValue Flag Invalid!\n"+
      "Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
      "Expected numStrDtoResultIsFractionalValue = '%v'\n"+
      "  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
      ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

    return
  }

  if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
      "Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
      "Expected numStrDtoResultAbsIntStr = '%v'\n"+
      "  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
      ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

    return
  }

  if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected numStrDtoResultAbsFracStr = '%v'\n"+
      "  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
      ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

    return
  }

  if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
      "Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
      "Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
      "  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
      ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factor INVALID!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
      "Expected numStrDtoBigInt = '%v'\n"+
      "  Actual numStrDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

    return
  }

  return
}

func TestNumStrDto_SetPrecision_04(t *testing.T) {

  ePrefix := "TestNumStrDto_SetPrecision_04"

  inputNumStr := "-123456789"

  inputPrecisionUint := uint(7)

  var inputRoundResult bool

  inputRoundResult = true

  expectedNumStr := "-123456789.0000000"

  expectedBigInt := big.NewInt(-1234567890000000)

  expectedPrecisionInt := 7

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

  expectedNumHasNumericDigits = true

  expectedNumIsFractionalValue = true

  expectedNumAbsIntStr := "123456789"

  expectedNumAbsFracStr := "0000000"

  expectedAbsAllRunesNumStr := "1234567890000000"

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  numStrDtoResult, err := new(NumStrDto).SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)\n"+
      "inputNumStr= '%v'\n"+
      "inputPrecisionUint= '%v'\n"+
      "inputRoundResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr,
      inputPrecisionUint,
      inputRoundResult,
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

  numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

  numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsIntRunes, err := \n"+
      "  numStrDtoResult.GetAbsIntRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

  numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsFracRunes, err :=\n"+
      "  numStrDtoResult.GetAbsFracRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

  numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

  numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINumNumberStr, err :=\n"+
      "  numStrDtoResultBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
      "expectedBigINum= '%v'\n"+
      "numStrDtoResultBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResultBigINum is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected numStrDtoResultBigINum = '%v'\n"+
      "  Actual numStrDtoResultBigINum = '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
    t.Errorf("%v\n"+
      "Error: Numeric Digits Flag is Invalid!\n"+
      "Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
      "Expected numStrDtoResultHasNumericDigits = '%v'\n"+
      "  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
      ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

    return
  }

  if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
    t.Errorf("%v\n"+
      "Error: IsFractionalValue Flag Invalid!\n"+
      "Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
      "Expected numStrDtoResultIsFractionalValue = '%v'\n"+
      "  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
      ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

    return
  }

  if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
      "Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
      "Expected numStrDtoResultAbsIntStr = '%v'\n"+
      "  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
      ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

    return
  }

  if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected numStrDtoResultAbsFracStr = '%v'\n"+
      "  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
      ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

    return
  }

  if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
      "Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
      "Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
      "  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
      ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factor INVALID!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
      "Expected numStrDtoBigInt = '%v'\n"+
      "  Actual numStrDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

    return
  }

  return
}

func TestNumStrDto_SetPrecision_05(t *testing.T) {

  ePrefix := "TestNumStrDto_SetPrecision_05"

  inputNumStr := "123456.789"

  inputPrecisionUint := uint(2)

  var inputRoundResult bool

  inputRoundResult = true

  expectedNumStr := "123456.79"

  expectedBigInt := big.NewInt(12345679)

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

  expectedNumHasNumericDigits = true

  expectedNumIsFractionalValue = true

  expectedNumAbsIntStr := "123456"

  expectedNumAbsFracStr := "79"

  expectedAbsAllRunesNumStr := "12345679"

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  numStrDtoResult, err := new(NumStrDto).SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)\n"+
      "inputNumStr= '%v'\n"+
      "inputPrecisionUint= '%v'\n"+
      "inputRoundResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr,
      inputPrecisionUint,
      inputRoundResult,
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

  numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

  numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsIntRunes, err := \n"+
      "  numStrDtoResult.GetAbsIntRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

  numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsFracRunes, err :=\n"+
      "  numStrDtoResult.GetAbsFracRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

  numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

  numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINumNumberStr, err :=\n"+
      "  numStrDtoResultBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
      "expectedBigINum= '%v'\n"+
      "numStrDtoResultBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResultBigINum is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected numStrDtoResultBigINum = '%v'\n"+
      "  Actual numStrDtoResultBigINum = '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
    t.Errorf("%v\n"+
      "Error: Numeric Digits Flag is Invalid!\n"+
      "Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
      "Expected numStrDtoResultHasNumericDigits = '%v'\n"+
      "  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
      ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

    return
  }

  if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
    t.Errorf("%v\n"+
      "Error: IsFractionalValue Flag Invalid!\n"+
      "Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
      "Expected numStrDtoResultIsFractionalValue = '%v'\n"+
      "  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
      ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

    return
  }

  if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
      "Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
      "Expected numStrDtoResultAbsIntStr = '%v'\n"+
      "  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
      ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

    return
  }

  if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected numStrDtoResultAbsFracStr = '%v'\n"+
      "  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
      ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

    return
  }

  if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
      "Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
      "Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
      "  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
      ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factor INVALID!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
      "Expected numStrDtoBigInt = '%v'\n"+
      "  Actual numStrDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

    return
  }

  return
}

func TestNumStrDto_SetPrecision_06(t *testing.T) {

  ePrefix := "TestNumStrDto_SetPrecision_06"

  inputNumStr := "123456.789"

  inputPrecisionUint := uint(2)

  var inputRoundResult bool

  inputRoundResult = false

  expectedNumStr := "123456.78"

  expectedBigInt := big.NewInt(12345678)

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

  expectedNumHasNumericDigits = true

  expectedNumIsFractionalValue = true

  expectedNumAbsIntStr := "123456"

  expectedNumAbsFracStr := "78"

  expectedAbsAllRunesNumStr := "12345678"

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  numStrDtoResult, err := new(NumStrDto).SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  SetPrecision(inputNumStr, inputPrecisionUint, inputRoundResult)\n"+
      "inputNumStr= '%v'\n"+
      "inputPrecisionUint= '%v'\n"+
      "inputRoundResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      inputNumStr,
      inputPrecisionUint,
      inputRoundResult,
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

  numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

  numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsIntRunes, err := \n"+
      "  numStrDtoResult.GetAbsIntRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

  numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsFracRunes, err :=\n"+
      "  numStrDtoResult.GetAbsFracRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

  numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

  numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINumNumberStr, err :=\n"+
      "  numStrDtoResultBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
      "expectedBigINum= '%v'\n"+
      "numStrDtoResultBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResultBigINum is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected numStrDtoResultBigINum = '%v'\n"+
      "  Actual numStrDtoResultBigINum = '%v'\n\n",
      ePrefix, expectedBigINumStr, numStrDtoResultBigINumNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
    t.Errorf("%v\n"+
      "Error: Numeric Digits Flag is Invalid!\n"+
      "Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
      "Expected numStrDtoResultHasNumericDigits = '%v'\n"+
      "  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
      ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

    return
  }

  if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
    t.Errorf("%v\n"+
      "Error: IsFractionalValue Flag Invalid!\n"+
      "Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
      "Expected numStrDtoResultIsFractionalValue = '%v'\n"+
      "  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
      ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

    return
  }

  if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
      "Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
      "Expected numStrDtoResultAbsIntStr = '%v'\n"+
      "  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
      ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

    return
  }

  if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected numStrDtoResultAbsFracStr = '%v'\n"+
      "  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
      ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

    return
  }

  if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
      "Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
      "Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
      "  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
      ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Scale Factor INVALID!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Values ARE NOT EQUAL!\n"+
      "Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
      "Expected numStrDtoBigInt = '%v'\n"+
      "  Actual numStrDtoBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

    return
  }

  return
}

func TestNumStrDto_SetPrecision_07(t *testing.T) {

  nStr := "123456.789"
  precision := uint(5)
  roundResult := false
  outPrecision := uint(5)
  expected := "123456.78900"
  signVal := 1
  absIntRuneStr := "123456"
  absFracRuneStr := "78900"

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_08(t *testing.T) {

  nStr := "123.456789"
  precision := uint(1)
  roundResult := false
  outPrecision := uint(1)
  expected := "123.4"
  signVal := 1
  absIntRuneStr := "123"
  absFracRuneStr := "4"

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_09(t *testing.T) {

  nStr := "123.456789"
  precision := uint(1)
  roundResult := true
  outPrecision := uint(1)
  expected := "123.5"
  signVal := 1
  absIntRuneStr := "123"
  absFracRuneStr := "5"

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_10(t *testing.T) {

  nStr := "-123.456789"
  precision := uint(1)
  roundResult := false
  outPrecision := uint(1)
  expected := "-123.4"
  signVal := -1
  absIntRuneStr := "123"
  absFracRuneStr := "4"

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_11(t *testing.T) {

  nStr := "-123.456789"
  precision := uint(1)
  roundResult := true
  outPrecision := uint(1)
  expected := "-123.5"
  signVal := -1
  absIntRuneStr := "123"
  absFracRuneStr := "5"

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_12(t *testing.T) {

  nStr := "123456.789"
  precision := uint(0)
  roundResult := true
  outPrecision := uint(0)
  expected := "123457"
  signVal := 1
  absIntRuneStr := "123457"
  absFracRuneStr := ""

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_13(t *testing.T) {

  nStr := "-123456.789"
  precision := uint(0)
  roundResult := true
  outPrecision := uint(0)
  expected := "-123457"
  signVal := -1
  absIntRuneStr := "123457"
  absFracRuneStr := ""

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_14(t *testing.T) {

  nStr := "123456.789"
  precision := uint(0)
  roundResult := false
  outPrecision := uint(0)
  expected := "123456"
  signVal := 1
  absIntRuneStr := "123456"
  absFracRuneStr := ""

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_15(t *testing.T) {

  nStr := "-123456.789"
  precision := uint(0)
  roundResult := false
  outPrecision := uint(0)
  expected := "-123456"
  signVal := -1
  absIntRuneStr := "123456"
  absFracRuneStr := ""

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_16(t *testing.T) {

  nStr := "123457"
  precision := uint(1)
  roundResult := false
  outPrecision := uint(1)
  expected := "123457.0"
  signVal := 1
  absIntRuneStr := "123457"
  absFracRuneStr := "0"

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_17(t *testing.T) {

  nStr := "123457"
  precision := uint(1)
  roundResult := true
  outPrecision := uint(1)
  expected := "123457.0"
  signVal := 1
  absIntRuneStr := "123457"
  absFracRuneStr := "0"

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_18(t *testing.T) {

  nStr := "-123457"
  precision := uint(1)
  roundResult := false
  outPrecision := uint(1)
  expected := "-123457.0"
  signVal := -1
  absIntRuneStr := "123457"
  absFracRuneStr := "0"

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetPrecision_19(t *testing.T) {

  nStr := "-123457"
  precision := uint(1)
  roundResult := true
  outPrecision := uint(1)
  expected := "-123457.0"
  signVal := -1
  absIntRuneStr := "123457"
  absFracRuneStr := "0"

  nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetThisPrecision_01(t *testing.T) {

  nStr := "123456789"
  initialPrecision := uint(0)
  precision := uint(3)
  roundResult := true
  outPrecision := uint(3)
  expected := "123456789.000"
  signVal := 1
  absIntRuneStr := "123456789"
  absFracRuneStr := "000"

  nsDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). "+
      "Err='%v' ", err.Error())
  }

  if initialPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Error: Expected Initial precision='%v'. Actual Initial precision='%v' ",
      initialPrecision, nsDto.GetPrecisionUint())
  }

  nsDto.SetThisPrecision(precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetThisPrecision(precision, roundResult). "+
      "precision= '%v'. Error= %v", precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_SetThisPrecision_02(t *testing.T) {

  nStr := "123456.789"
  initialPrecision := uint(3)
  precision := uint(2)
  roundResult := true
  outPrecision := uint(2)
  expected := "123456.79"
  signVal := 1
  absIntRuneStr := "123456"
  absFracRuneStr := "79"

  nsDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). "+
      "Err='%v' ", err.Error())
  }

  if initialPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Error: Expected Initial precision='%v'. Actual Initial precision='%v' ",
      initialPrecision, nsDto.GetPrecisionUint())
  }

  nsDto.SetThisPrecision(precision, roundResult)

  if err != nil {
    t.Errorf("Received error from nsu.SetThisPrecision(precision, roundResult). "+
      "precision= '%v'. Error= %v", precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_01(t *testing.T) {

  nStr := "123456.789"
  precision := uint(3)
  outPrecision := uint(6)
  expected := "123.456789"
  signVal := 1
  absIntRuneStr := "123"
  absFracRuneStr := "456789"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_02(t *testing.T) {

  nStr := "123456.789"
  precision := uint(2)
  outPrecision := uint(5)
  expected := "1234.56789"
  signVal := 1
  absIntRuneStr := "1234"
  absFracRuneStr := "56789"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_03(t *testing.T) {

  nStr := "123456.789"
  precision := uint(6)
  outPrecision := uint(9)
  expected := "0.123456789"
  signVal := 1
  absIntRuneStr := "0"
  absFracRuneStr := "123456789"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_04(t *testing.T) {

  nStr := "123456789"
  precision := uint(6)
  outPrecision := uint(6)
  expected := "123.456789"
  signVal := 1
  absIntRuneStr := "123"
  absFracRuneStr := "456789"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_05(t *testing.T) {

  nStr := "123"
  precision := uint(5)
  outPrecision := uint(5)
  expected := "0.00123"
  signVal := 1
  absIntRuneStr := "0"
  absFracRuneStr := "00123"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_06(t *testing.T) {

  nStr := "0"
  precision := uint(3)
  outPrecision := uint(3)
  expected := "0.000"
  signVal := 1
  absIntRuneStr := "0"
  absFracRuneStr := "000"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_07(t *testing.T) {

  nStr := "0.000"
  precision := uint(2)
  outPrecision := uint(5)
  expected := "0.00000"
  signVal := 1
  absIntRuneStr := "0"
  absFracRuneStr := "00000"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_08(t *testing.T) {

  nStr := "123456.789"
  precision := uint(0)
  outPrecision := uint(3)
  expected := "123456.789"
  signVal := 1
  absIntRuneStr := "123456"
  absFracRuneStr := "789"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_09(t *testing.T) {

  nStr := "-123456.789"
  precision := uint(0)
  outPrecision := uint(3)
  expected := "-123456.789"
  signVal := -1
  absIntRuneStr := "123456"
  absFracRuneStr := "789"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_10(t *testing.T) {

  nStr := "-123456.789"
  precision := uint(3)
  outPrecision := uint(6)
  expected := "-123.456789"
  signVal := -1
  absIntRuneStr := "123"
  absFracRuneStr := "456789"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}

func TestNumStrDto_ShiftPrecisionLeft_11(t *testing.T) {

  nStr := "-123456789"
  precision := uint(6)
  outPrecision := uint(6)
  expected := "-123.456789"
  signVal := -1
  absIntRuneStr := "123"
  absFracRuneStr := "456789"

  nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

  if err != nil {
    t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
  }

  if nsDto.GetNumStr() != expected {
    t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
  }

  if outPrecision != nsDto.GetPrecisionUint() {
    t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecisionUint())
  }

  if signVal != nsDto.GetSign() {
    t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
  }

  err = nsDto.IsValid("Test 'nsDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nsDto.IsValid() Error='%v'", err.Error())
  }

  if !nsDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits())
  }

  s := string(nsDto.GetAbsIntRunes())

  if s != absIntRuneStr {
    t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
  }

  s = string(nsDto.GetAbsFracRunes())

  if s != absFracRuneStr {
    t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
  }

}
