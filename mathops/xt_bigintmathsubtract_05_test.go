package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntMathSubtract_SubtractPair_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_01"

  // minuend = 123.32
  minuendStr := "123.32"

  minuendPrecision := uint(2)

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  subtrahendPrecision := uint(3)

  // result = 99.999
  expectedBigINumStr := "99.999"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_SubtractPair_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_02"

  // minuend = 949321.6712
  minuendStr := "949321.6712"

  minuendPrecision := uint(4)

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  subtrahendPrecision := uint(2)

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"

  expectedBigINumSign := 1

  var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_SubtractPair_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_03"

  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  minuendPrecision := uint(8)

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  subtrahendPrecision := uint(3)

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"

  expectedBigINumSign := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_SubtractPair_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_04"

  // minuend = -289.673849
  minuendStr := "-289.673849"

  minuendPrecision := uint(6)

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  subtrahendPrecision := uint(3)

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_SubtractPair_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_05"

  // minuend = 0
  minuendStr := "0"

  minuendPrecision := uint(0)

  // subtrahend = 0
  subtrahendStr := "0"

  subtrahendPrecision := uint(0)

  // result = 0
  expectedBigINumStr := "0"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_SubtractPair_06(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_06"

  // minuend = 270.1
  minuendStr := "270.1"

  minuendPrecision := uint(1)

  // subtrahend = 0
  subtrahendStr := "0"

  subtrahendPrecision := uint(0)

  // result = 270.1
  expectedBigINumStr := "270.1"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_SubtractPair_07(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_07"

  // minuend = 0
  minuendStr := "0"

  minuendPrecision := uint(0)

  // subtrahend = 270.1
  subtrahendStr := "270.1"

  subtrahendPrecision := uint(1)

  // result = -270.1
  expectedBigINumStr := "-270.1"

  expectedBigINumSign := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_SubtractPair_08(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_08"

  // minuend = 0
  minuendStr := "0"

  minuendPrecision := uint(0)

  // subtrahend = -270.1
  subtrahendStr := "-270.1"

  subtrahendPrecision := uint(1)

  // result = 270.1
  expectedBigINumStr := "270.1"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_SubtractPair_09(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_09"

  // minuend = 2.5
  minuendStr := "2.5"

  minuendPrecision := uint(1)

  // subtrahend = 2.5
  subtrahendStr := "2.5"

  subtrahendPrecision := uint(1)

  // result = 0.0
  expectedBigINumStr := "0.0"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_SubtractPair_10(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractPair_10"

  var err error

  // minuend = 123.32
  minuendStr := "123.32"

  minuendPrecision := uint(2)

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  subtrahendPrecision := uint(3)

  // result = 99.999
  expectedBigINumStr := "99,999"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  maxPrecision := minuendPrecision

  if subtrahendPrecision > maxPrecision {
    maxPrecision = subtrahendPrecision
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendBiNum.IsValid("Validating minuendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendBiNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Are Not Equal\n"+
      "Because minuendStr != minuendBiNumStr\n"+
      "Expected minuendBiNumStr = '%v'\n"+
      "  Actual minuendBiNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendBiNumStr)

    return
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != subtrahendBiNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != subtrahendBiNumStr \n"+
      "Expected subtrahendBiNumStr = '%v'\n"+
      "  Actual subtrahendBiNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, subtrahendBiNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

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

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
      err.Error())

    return
  }

  err = bPair.IsValid("Validating bPair")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.IsValid('Validating bPair')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair1NumStr, err := bPair.Big1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair2NumStr, err := bPair.Big2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig1PrecisionUint\n"+
      "Expected bPairBig1PrecisionUint = '%v'\n"+
      "  Actual bPairBig1PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig1PrecisionUint)

    return
  }

  bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if maxPrecision != bPairBig2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because maxPrecision != bPairBig2PrecisionUint\n"+
      "Expected bPairBig2PrecisionUint = '%v'\n"+
      "  Actual bPairBig2PrecisionUint = '%v'\n\n",
      ePrefix, maxPrecision, bPairBig2PrecisionUint)

    return
  }

  result, err := new(BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
      "bPair.Big1= '%v'\n"+
      "bPair.Big2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bPair1NumStr,
      bPair2NumStr,
      err.Error())

    return
  }

  err = result.SetNumericSeparatorsDto(expectedNumSeps)

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

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

  return
}

func TestBigIntMathSubtract_BigIntSubtract_01(t *testing.T) {
  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99.999"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_BigIntSubtract_02(t *testing.T) {
  // minuend = 949321.6712
  minuendStr := "949321.6712"

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_BigIntSubtract_03(t *testing.T) {
  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_BigIntSubtract_04(t *testing.T) {
  // minuend = -289.673849
  minuendStr := "-289.673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_BigIntSubtract_05(t *testing.T) {

  minuendStr := "5"

  subtrahendStr := "5"

  expectedBigINumStr := "0"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_BigIntSubtract_06(t *testing.T) {

  minuendStr := "-5"
  subtrahendStr := "5"
  expectedBigINumStr := "-10"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_BigIntSubtract_07(t *testing.T) {

  minuendStr := "0"
  subtrahendStr := "0"
  expectedBigINumStr := "0"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_BigIntSubtract_08(t *testing.T) {

  minuendStr := "50.0"
  subtrahendStr := "2.60134"
  expectedBigINumStr := "47.39866"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_BigIntSubtract_09(t *testing.T) {

  minuendStr := "50.0"
  subtrahendStr := "50.0000"
  expectedBigINumStr := "0"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_BigIntSubtract_10(t *testing.T) {

  minuendStr := "0.0"
  subtrahendStr := "0.0000"
  expectedBigINumStr := "0"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend := big.NewInt(0).Set(minuendBiNum.bigInt)
  minuendPrecision := minuendBiNum.GetPrecisionBigInt()
  biSubtrahend := big.NewInt(0).Set(subtrahendBiNum.bigInt)
  subtrahendPrecision := subtrahendBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(
    bIMinuend,
    minuendPrecision,
    biSubtrahend,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).BigIntSubtract(...) "+
      "Error='%v' ", err.Error())
  }

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathSubtract_FixedDecimalSubtract_01(t *testing.T) {
  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99.999"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend :=
    BigIntFixedDecimal{}.New(
      minuendBiNum.GetIntegerValue(),
      minuendBiNum.GetPrecisionUint())

  biSubtrahend :=
    BigIntFixedDecimal{}.New(
      subtrahendBiNum.GetIntegerValue(),
      subtrahendBiNum.GetPrecisionUint())

  result := new(BigIntMathSubtract).FixedDecimalSubtract(
    bIMinuend,
    biSubtrahend)

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result.GetIntegerValue()) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.GetIntegerValue().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecisionUint() {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), result.GetPrecisionUint())
  }
}

func TestBigIntMathSubtract_FixedDecimalSubtract_02(t *testing.T) {
  // minuend = 949321.6712
  minuendStr := "949321.6712"

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend :=
    BigIntFixedDecimal{}.New(
      minuendBiNum.GetIntegerValue(),
      minuendBiNum.GetPrecisionUint())

  biSubtrahend :=
    BigIntFixedDecimal{}.New(
      subtrahendBiNum.GetIntegerValue(),
      subtrahendBiNum.GetPrecisionUint())

  result := new(BigIntMathSubtract).FixedDecimalSubtract(
    bIMinuend,
    biSubtrahend)

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result.GetIntegerValue()) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.GetIntegerValue().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecisionUint() {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), result.GetPrecisionUint())
  }
}

func TestBigIntMathSubtract_FixedDecimalSubtract_03(t *testing.T) {
  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend :=
    BigIntFixedDecimal{}.New(
      minuendBiNum.GetIntegerValue(),
      minuendBiNum.GetPrecisionUint())

  biSubtrahend :=
    BigIntFixedDecimal{}.New(
      subtrahendBiNum.GetIntegerValue(),
      subtrahendBiNum.GetPrecisionUint())

  result := new(BigIntMathSubtract).FixedDecimalSubtract(
    bIMinuend,
    biSubtrahend)

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result.GetIntegerValue()) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.GetIntegerValue().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecisionUint() {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), result.GetPrecisionUint())
  }
}

func TestBigIntMathSubtract_FixedDecimalSubtract_04(t *testing.T) {
  // minuend = -289.673849
  minuendStr := "-289.673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend :=
    BigIntFixedDecimal{}.New(
      minuendBiNum.GetIntegerValue(),
      minuendBiNum.GetPrecisionUint())

  biSubtrahend :=
    BigIntFixedDecimal{}.New(
      subtrahendBiNum.GetIntegerValue(),
      subtrahendBiNum.GetPrecisionUint())

  result := new(BigIntMathSubtract).FixedDecimalSubtract(
    bIMinuend,
    biSubtrahend)

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result.GetIntegerValue()) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.GetIntegerValue().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecisionUint() {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), result.GetPrecisionUint())
  }
}

func TestBigIntMathSubtract_FixedDecimalSubtract_05(t *testing.T) {

  minuendStr := "5"

  subtrahendStr := "5"

  expectedBigINumStr := "0"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend :=
    BigIntFixedDecimal{}.New(
      minuendBiNum.GetIntegerValue(),
      minuendBiNum.GetPrecisionUint())

  biSubtrahend :=
    BigIntFixedDecimal{}.New(
      subtrahendBiNum.GetIntegerValue(),
      subtrahendBiNum.GetPrecisionUint())

  result := new(BigIntMathSubtract).FixedDecimalSubtract(
    bIMinuend,
    biSubtrahend)

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result.GetIntegerValue()) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.GetIntegerValue().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecisionUint() {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), result.GetPrecisionUint())
  }
}

func TestBigIntMathSubtract_FixedDecimalSubtract_06(t *testing.T) {

  minuendStr := "-5"
  subtrahendStr := "5"
  expectedBigINumStr := "-10"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend :=
    BigIntFixedDecimal{}.New(
      minuendBiNum.GetIntegerValue(),
      minuendBiNum.GetPrecisionUint())

  biSubtrahend :=
    BigIntFixedDecimal{}.New(
      subtrahendBiNum.GetIntegerValue(),
      subtrahendBiNum.GetPrecisionUint())

  result := new(BigIntMathSubtract).FixedDecimalSubtract(
    bIMinuend,
    biSubtrahend)

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result.GetIntegerValue()) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.GetIntegerValue().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecisionUint() {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), result.GetPrecisionUint())
  }
}

func TestBigIntMathSubtract_FixedDecimalSubtract_07(t *testing.T) {

  minuendStr := "0"
  subtrahendStr := "0"
  expectedBigINumStr := "0"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend :=
    BigIntFixedDecimal{}.New(
      minuendBiNum.GetIntegerValue(),
      minuendBiNum.GetPrecisionUint())

  biSubtrahend :=
    BigIntFixedDecimal{}.New(
      subtrahendBiNum.GetIntegerValue(),
      subtrahendBiNum.GetPrecisionUint())

  result := new(BigIntMathSubtract).FixedDecimalSubtract(
    bIMinuend,
    biSubtrahend)

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result.GetIntegerValue()) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.GetIntegerValue().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecisionUint() {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), result.GetPrecisionUint())
  }
}

func TestBigIntMathSubtract_FixedDecimalSubtract_08(t *testing.T) {

  minuendStr := "50.0"
  subtrahendStr := "2.60134"
  expectedBigINumStr := "47.39866"

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v'  Error='%v'. ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMinuend :=
    BigIntFixedDecimal{}.New(
      minuendBiNum.GetIntegerValue(),
      minuendBiNum.GetPrecisionUint())

  biSubtrahend :=
    BigIntFixedDecimal{}.New(
      subtrahendBiNum.GetIntegerValue(),
      subtrahendBiNum.GetPrecisionUint())

  result := new(BigIntMathSubtract).FixedDecimalSubtract(
    bIMinuend,
    biSubtrahend)

  expectedBI := expectedBigINum.GetIntegerValue()

  if expectedBI.Cmp(result.GetIntegerValue()) != 0 {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBI.Text(10), result.GetIntegerValue().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecisionUint() {
    t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'. ",
      expectedBigINum.GetPrecisionUint(), result.GetPrecisionUint())
  }
}
