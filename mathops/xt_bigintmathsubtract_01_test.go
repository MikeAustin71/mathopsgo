package mathops

import (
  "testing"
)

func TestBigIntMathSubtract_SubtractBigInts_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigInts_01"

  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99.999"

  expectedBigINumSign := 1

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

  bIMinuendBigInt, err := minuendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bIMinuendBigInt, err := minuendBiNum.GetBigInt()\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, err.Error())
    return
  }

  minuendPrecision, err := minuendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendPrecision, err := minuendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  biSubtrahendBigInt, err := subtrahendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biSubtrahendBigInt, err = subtrahendBiNum.GetBigInt()\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendBiNumStr, err.Error())
    return
  }

  subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  result, err := new(BigIntMathSubtract).SubtractBigInts(
    bIMinuendBigInt,
    minuendPrecision,
    biSubtrahendBigInt,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigInts(\n"+
      "  bIMinuendBigInt, minuendPrecision, biSubtrahendBigInt, subtrahendPrecision)\n"+
      "bIMinuendBigInt= '%v'\n"+
      "minuendPrecision= '%v'\n"+
      "biSubtrahendBigInt= '%v'\n"+
      "subtrahendPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bIMinuendBigInt.Text(10),
      minuendPrecision,
      biSubtrahendBigInt.Text(10),
      subtrahendPrecision,
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

  return
}

func TestBigIntMathSubtract_SubtractBigInts_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigInts_02"

  // minuend = 949321.6712
  minuendStr := "949321.6712"

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"

  expectedBigINumSign := 1

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

  bIMinuendBigInt, err := minuendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bIMinuendBigInt, err := minuendBiNum.GetBigInt()\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, err.Error())
    return
  }

  minuendPrecision, err := minuendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendPrecision, err := minuendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  biSubtrahendBigInt, err := subtrahendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biSubtrahendBigInt, err = subtrahendBiNum.GetBigInt()\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendBiNumStr, err.Error())
    return
  }

  subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  result, err := new(BigIntMathSubtract).SubtractBigInts(
    bIMinuendBigInt,
    minuendPrecision,
    biSubtrahendBigInt,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigInts(\n"+
      "  bIMinuendBigInt, minuendPrecision, biSubtrahendBigInt, subtrahendPrecision)\n"+
      "bIMinuendBigInt= '%v'\n"+
      "minuendPrecision= '%v'\n"+
      "biSubtrahendBigInt= '%v'\n"+
      "subtrahendPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bIMinuendBigInt.Text(10),
      minuendPrecision,
      biSubtrahendBigInt.Text(10),
      subtrahendPrecision,
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

  return
}

func TestBigIntMathSubtract_SubtractBigInts_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigInts_03"

  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"

  expectedBigINumSign := -1

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

  bIMinuendBigInt, err := minuendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bIMinuendBigInt, err := minuendBiNum.GetBigInt()\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, err.Error())
    return
  }

  minuendPrecision, err := minuendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendPrecision, err := minuendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  biSubtrahendBigInt, err := subtrahendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biSubtrahendBigInt, err = subtrahendBiNum.GetBigInt()\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendBiNumStr, err.Error())
    return
  }

  subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  result, err := new(BigIntMathSubtract).SubtractBigInts(
    bIMinuendBigInt,
    minuendPrecision,
    biSubtrahendBigInt,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigInts(\n"+
      "  bIMinuendBigInt, minuendPrecision, biSubtrahendBigInt, subtrahendPrecision)\n"+
      "bIMinuendBigInt= '%v'\n"+
      "minuendPrecision= '%v'\n"+
      "biSubtrahendBigInt= '%v'\n"+
      "subtrahendPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bIMinuendBigInt.Text(10),
      minuendPrecision,
      biSubtrahendBigInt.Text(10),
      subtrahendPrecision,
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

  return
}

func TestBigIntMathSubtract_SubtractBigInts_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigInts_04"

  // minuend = -289.673849
  minuendStr := "-289.673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"

  expectedBigINumSign := 1

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

  bIMinuendBigInt, err := minuendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bIMinuendBigInt, err := minuendBiNum.GetBigInt()\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, err.Error())
    return
  }

  minuendPrecision, err := minuendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendPrecision, err := minuendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  biSubtrahendBigInt, err := subtrahendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biSubtrahendBigInt, err = subtrahendBiNum.GetBigInt()\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendBiNumStr, err.Error())
    return
  }

  subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  result, err := new(BigIntMathSubtract).SubtractBigInts(
    bIMinuendBigInt,
    minuendPrecision,
    biSubtrahendBigInt,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigInts(\n"+
      "  bIMinuendBigInt, minuendPrecision, biSubtrahendBigInt, subtrahendPrecision)\n"+
      "bIMinuendBigInt= '%v'\n"+
      "minuendPrecision= '%v'\n"+
      "biSubtrahendBigInt= '%v'\n"+
      "subtrahendPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bIMinuendBigInt.Text(10),
      minuendPrecision,
      biSubtrahendBigInt.Text(10),
      subtrahendPrecision,
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

  return
}

func TestBigIntMathSubtract_SubtractBigInts_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigInts_05"

  // minuend = 25
  minuendStr := "25"

  // subtrahend = 1.00
  subtrahendStr := "1.00"

  // result = 24
  expectedBigINumStr := "24"

  expectedBigINumSign := 1

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

  bIMinuendBigInt, err := minuendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bIMinuendBigInt, err := minuendBiNum.GetBigInt()\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, err.Error())
    return
  }

  minuendPrecision, err := minuendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendPrecision, err := minuendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  biSubtrahendBigInt, err := subtrahendBiNum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biSubtrahendBigInt, err = subtrahendBiNum.GetBigInt()\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendBiNumStr, err.Error())
    return
  }

  subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendPrecision, err := subtrahendBiNum.GetPrecisionUint()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  result, err := new(BigIntMathSubtract).SubtractBigInts(
    bIMinuendBigInt,
    minuendPrecision,
    biSubtrahendBigInt,
    subtrahendPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigInts(\n"+
      "  bIMinuendBigInt, minuendPrecision, biSubtrahendBigInt, subtrahendPrecision)\n"+
      "bIMinuendBigInt= '%v'\n"+
      "minuendPrecision= '%v'\n"+
      "biSubtrahendBigInt= '%v'\n"+
      "subtrahendPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bIMinuendBigInt.Text(10),
      minuendPrecision,
      biSubtrahendBigInt.Text(10),
      subtrahendPrecision,
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

  return
}

func TestBigIntMathSubtract_SubtractBigIntNums_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNums_01"

  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99.999"

  expectedBigINumSign := 1

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

  result, err := new(BigIntMathSubtract).SubtractBigIntNums(
    minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNums(\n"+
      "  minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
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

  return
}

func TestBigIntMathSubtract_SubtractBigIntNums_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNums_02"

  // minuend = 949321.6712
  minuendStr := "949321.6712"

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"

  expectedBigINumSign := 1

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

  result, err := new(BigIntMathSubtract).SubtractBigIntNums(
    minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNums(\n"+
      "  minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
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

  return
}

func TestBigIntMathSubtract_SubtractBigIntNums_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNums_03"

  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"

  expectedBigINumSign := -1

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

  result, err := new(BigIntMathSubtract).SubtractBigIntNums(
    minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNums(\n"+
      "  minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
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

  return
}

func TestBigIntMathSubtract_SubtractBigIntNums_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNums_04"

  // minuend = -289.673849
  minuendStr := "-289.673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"

  expectedBigINumSign := 1

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

  result, err := new(BigIntMathSubtract).SubtractBigIntNums(
    minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNums(\n"+
      "  minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
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

  return
}

func TestBigIntMathSubtract_SubtractBigIntNums_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNums_05"

  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

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

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  minuendBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(minuendStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(minuendStr, &usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, usaNumSeps.String(), err.Error())
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

  subtrahendBiNum, err := new(BigIntNum).NewNumStrWithNumSeps(subtrahendStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(subtrahendStr, &usaNumSeps)\n"+
      "subtrahendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, usaNumSeps.String(), err.Error())
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

  result, err := new(BigIntMathSubtract).SubtractBigIntNums(
    minuendBiNum, subtrahendBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNums(\n"+
      "  minuendBiNum, subtrahendBiNum)\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      subtrahendBiNumStr,
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

  err = result.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumSeps.String(), err.Error())

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

  return
}

func TestBigIntMathSubtract_SubtractBigIntNumArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumArray_01"

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"

  subtrahend1 := "67.1"

  subtrahend2 := "93.0"

  subtrahend3 := "-124498.67158"

  subtrahend4 := "647129.57"

  subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998.231036"

  expectedBigINumSign := 1

  var err error

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

  lenSubtrahends := 6

  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumArray(minuendBiNum, subtrahendAry)

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_02(t *testing.T) {

  var err error

  // minuend = -18,973,642.1234567
  minuendStr := "-18973642.1234567"

  subtrahend0 := "737.21"
  subtrahend1 := "9637591.879546"
  subtrahend2 := "28"
  subtrahend3 := "5284.9765"
  subtrahend4 := "-189291837.12"
  subtrahend5 := "7638932.12398765"

  // result = 153,035,620.80650965
  expectedBigINumStr := "153035620.80650965"
  expectedBigINumSign := 1

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumArray(minuendBiNum, subtrahendAry)

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_03(t *testing.T) {

  var err error
  // minuend =   1,718,973,642.1234567
  minuendStr := "1718973642.1234567"

  subtrahend0 := "-28934682.721"
  subtrahend1 := "424.987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62.94"
  subtrahend4 := "-999999999.99999"
  subtrahend5 := "-9638932.371"

  // Result:  2,757,547,756.287792379
  expectedBigINumStr := "2757547756.287792379"
  expectedBigINumSign := 1

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumArray(minuendBiNum, subtrahendAry)

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_04(t *testing.T) {

  var err error

  // minuend =   -1,718,973,642.1234567
  minuendStr := "-1718973642.1234567"

  subtrahend0 := "-28934682.721"
  subtrahend1 := "424.987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62.94"
  subtrahend4 := "-999999999.99999"
  subtrahend5 := "-9638932.371"

  // Result:   -680,399,527.959121021
  expectedBigINumStr := "-680399527.959121021"
  expectedBigINumSign := -1

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumArray(minuendBiNum, subtrahendAry)

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractBigIntNumArray_05(t *testing.T) {

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"
  subtrahend1 := "67.1"
  subtrahend2 := "93.0"
  subtrahend3 := "-124498.67158"
  subtrahend4 := "647129.57"
  subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998,231036"
  var err error

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumArray(minuendBiNum, subtrahendAry)

  actualResultStr := result.GetNumStr()

  if expectedBigINumStr != actualResultStr {
    t.Errorf("Comparison Error: Expected result='%v'. Instead, result= '%v'. ",
      expectedBigINumStr, actualResultStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_01(t *testing.T) {

  var err error

  // minuend =   100
  minuendStr := "100"

  subtrahendStrs := []string{
    "5",
    "10",
    "30",
    "60.55",
    "-100.1",
    "-5.6",
  }

  expectedStrs := []string{
    "95",
    "90",
    "70",
    "39.45",
    "200.1",
    "105.6",
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]BigIntNum, lenSubtrahends)
  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }

  }

  resultArray :=
    new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }

  }

}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_02(t *testing.T) {

  var err error

  // minuend =   5051
  minuendStr := "5051"

  subtrahendStrs := []string{
    "8000",
    "6051.123456",
    "-30871.25",
    "604.55",
    "9100.123",
    "-115.76",
  }

  expectedStrs := []string{
    "-2949",
    "-1000.123456",
    "35922.25",
    "4446.45",
    "-4049.123",
    "5166.76",
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]BigIntNum, lenSubtrahends)
  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }

  }

  resultArray :=
    new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }

  }

}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_03(t *testing.T) {

  var err error

  // minuend =   -20051.974578
  minuendStr := "-20051.974578"

  subtrahendStrs := []string{
    "476.543798",
    "6051.123456",
    "-270871.25",
    "15604.5589321",
    "987100.123",
    "-114555.76",
  }

  expectedStrs := []string{
    "-20528.518376",
    "-26103.098034",
    "250819.275422",
    "-35656.5335101",
    "-1007152.097578",
    "94503.785422",
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]BigIntNum, lenSubtrahends)
  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }

  }

  resultArray :=
    new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }

  }

}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_04(t *testing.T) {

  var err error

  // minuend =   0
  minuendStr := "0"

  subtrahendStrs := []string{
    "476.543798",
    "6051.123456",
    "-270871.25",
    "15604.5589321",
    "987100.123",
    "-114555.76",
  }

  expectedStrs := []string{
    "-476.543798",
    "-6051.123456",
    "270871.25",
    "-15604.5589321",
    "-987100.123",
    "114555.76",
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]BigIntNum, lenSubtrahends)
  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }

  }

  resultArray :=
    new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }

  }

}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_05(t *testing.T) {

  var err error

  // minuend =   98.2
  minuendStr := "98.2"

  subtrahendStrs := []string{
    "0",
    "0.000",
    "0",
    "0",
    "0.00000",
    "0.0",
  }

  expectedStrs := []string{
    "98.2",
    "98.200",
    "98.2",
    "98.2",
    "98.20000",
    "98.2",
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]BigIntNum, lenSubtrahends)
  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }

  }

  resultArray :=
    new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }

  }
}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_06(t *testing.T) {

  var err error

  // minuend =   100
  minuendStr := "100"

  subtrahendStrs := []string{
    "5",
    "10",
    "30",
    "60.55",
    "-100.1",
    "-5.6",
  }

  expectedStrs := []string{
    "95",
    "90",
    "70",
    "39,45",
    "200,1",
    "105,6",
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

  }

  resultArray :=
    new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  for k := 0; k < lenSubtrahends; k++ {

    actualResultStr := resultArray[k].GetNumStr()

    if expectedStrs[k] != actualResultStr {
      t.Errorf("Error: Expected result='%v'. Instead, result='%v' Index='%v'",
        expectedStrs[k], actualResultStr, k)
    }

    actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. Index='%v'",
        expectedNumSeps.String(), actualNumSeps.String(), k)
    }

  }

}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_01(t *testing.T) {

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"
  subtrahend1 := "67.1"
  subtrahend2 := "93.0"
  subtrahend3 := "-124498.67158"
  subtrahend4 := "647129.57"
  subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998.231036"
  expectedBigINumSign := 1
  var err error

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := 6

  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_02(t *testing.T) {

  var err error

  // minuend = -18,973,642.1234567
  minuendStr := "-18973642.1234567"

  subtrahend0 := "737.21"
  subtrahend1 := "9637591.879546"
  subtrahend2 := "28"
  subtrahend3 := "5284.9765"
  subtrahend4 := "-189291837.12"
  subtrahend5 := "7638932.12398765"

  // result = 153,035,620.80650965
  expectedBigINumStr := "153035620.80650965"
  expectedBigINumSign := 1

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_03(t *testing.T) {

  var err error

  // minuend =   1,718,973,642.1234567
  minuendStr := "1718973642.1234567"

  subtrahend0 := "-28934682.721"
  subtrahend1 := "424.987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62.94"
  subtrahend4 := "-999999999.99999"
  subtrahend5 := "-9638932.371"

  // Result:  2,757,547,756.287792379
  expectedBigINumStr := "2757547756.287792379"
  expectedBigINumSign := 1

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_04(t *testing.T) {

  var err error

  // minuend =   -1,718,973,642.1234567
  minuendStr := "-1718973642.1234567"

  subtrahend0 := "-28934682.721"
  subtrahend1 := "424.987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62.94"
  subtrahend4 := "-999999999.99999"
  subtrahend5 := "-9638932.371"

  // Result:   -680,399,527.959121021
  expectedBigINumStr := "-680399527.959121021"
  expectedBigINumSign := -1

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_05(t *testing.T) {

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"
  subtrahend1 := "67.1"
  subtrahend2 := "93.0"
  subtrahend3 := "-124498.67158"
  subtrahend4 := "647129.57"
  subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998,231036"

  var err error

  minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by minuendBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenSubtrahends := 6

  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend0). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend1). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend2). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend3). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend4). "+
      "Error='%v'. ", err.Error())
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewNumStr(subtrahend5). "+
      "Error='%v'. ", err.Error())
  }

  result := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  actualResultStr := result.GetNumStr()

  if expectedBigINumStr != actualResultStr {
    t.Errorf("Comparison Error: Expected result='%v'. Instead, result='%v'. ",
      expectedBigINumStr, actualResultStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathSubtract_SubtractDecimal_01(t *testing.T) {
  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99.999"

  expectedBigINumSign := 1

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimals(decMinuend, "+
      "decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
      minuendStr, subtrahendStr, err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractDecimal_02(t *testing.T) {
  // minuend = 949321.6712
  minuendStr := "949321.6712"

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"
  expectedBigINumSign := 1

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimals(decMinuend, "+
      "decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
      minuendStr, subtrahendStr, err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractDecimal_03(t *testing.T) {
  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"
  expectedBigINumSign := -1

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimals(decMinuend, "+
      "decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
      minuendStr, subtrahendStr, err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractDecimal_04(t *testing.T) {
  // minuend = -289.673849
  minuendStr := "-289.673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"
  expectedBigINumSign := 1

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimals(decMinuend, "+
      "decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
      minuendStr, subtrahendStr, err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractDecimal_05(t *testing.T) {
  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedNumStr := "99,999"

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  decSubtrahend, err := Decimal{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = decMinuend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by decMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimals(decMinuend, "+
      "decSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
      minuendStr, subtrahendStr, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathSubtract_SubtractDecimalArray_01(t *testing.T) {

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"
  subtrahend1 := "67.1"
  subtrahend2 := "93.0"
  subtrahend3 := "-124498.67158"
  subtrahend4 := "647129.57"
  subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998.231036"
  expectedBigINumSign := 1

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalArray("+
      "decMinuend, subtrahendAry). Error='%v' ", err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractDecimalArray_02(t *testing.T) {

  var err error

  // minuend = -18,973,642.1234567
  minuendStr := "-18973642.1234567"

  subtrahend0 := "737.21"
  subtrahend1 := "9637591.879546"
  subtrahend2 := "28"
  subtrahend3 := "5284.9765"
  subtrahend4 := "-189291837.12"
  subtrahend5 := "7638932.12398765"

  // result = 153,035,620.80650965
  expectedBigINumStr := "153035620.80650965"

  expectedBigINumSign := 1

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalArray("+
      "decMinuend, subtrahendAry). Error='%v' ", err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractDecimalArray_03(t *testing.T) {

  var err error

  // minuend =   1,718,973,642.1234567
  minuendStr := "1718973642.1234567"

  subtrahend0 := "-28934682.721"
  subtrahend1 := "424.987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62.94"
  subtrahend4 := "-999999999.99999"
  subtrahend5 := "-9638932.371"

  // Result:  2,757,547,756.287792379
  expectedBigINumStr := "2757547756.287792379"
  expectedBigINumSign := 1

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalArray("+
      "decMinuend, subtrahendAry). Error='%v' ", err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractDecimalArray_04(t *testing.T) {

  var err error

  // minuend =   -1,718,973,642.1234567
  minuendStr := "-1718973642.1234567"

  subtrahend0 := "-28934682.721"
  subtrahend1 := "424.987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62.94"
  subtrahend4 := "-999999999.99999"
  subtrahend5 := "-9638932.371"

  // Result:   -680,399,527.959121021
  expectedBigINumStr := "-680399527.959121021"
  expectedBigINumSign := -1

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalArray("+
      "decMinuend, subtrahendAry). Error='%v' ", err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathSubtract_SubtractDecimalArray_05(t *testing.T) {

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"
  subtrahend1 := "67.1"
  subtrahend2 := "93.0"
  subtrahend3 := "-124498.67158"
  subtrahend4 := "647129.57"
  subtrahend5 := "28"

  // result = 6805998.231036
  expectedNumStr := "6805998,231036"

  decMinuend, err := Decimal{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = decMinuend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by decMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalArray("+
      "decMinuend, subtrahendAry). Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}
