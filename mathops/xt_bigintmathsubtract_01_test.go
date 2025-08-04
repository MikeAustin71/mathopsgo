package mathops

import (
  "fmt"
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
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  // subtrahendStrs
  subtrahendStrs := []string{
    "123.894000",
    "67.1",
    "93.0",
    "-124498.67158",
    "647129.57",
    "28",
  }

  //
  //subtrahend0 := "123.894000"
  //
  //subtrahend1 := "67.1"
  //
  //subtrahend2 := "93.0"
  //
  //subtrahend3 := "-124498.67158"
  //
  //subtrahend4 := "647129.57"
  //
  //subtrahend5 := "28"

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

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]BigIntNum, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(
    minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
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

func TestBigIntMathSubtract_SubtractBigIntNumArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumArray_02"

  var err error

  // minuend = -18,973,642.1234567
  minuendStr := "-18973642.1234567"

  // subtrahendStrs
  subtrahendStrs := []string{
    "737.21",
    "9637591.879546",
    "28",
    "5284.9765",
    "-189291837.12",
    "7638932.12398765",
  }

  //subtrahend0 := "737.21"
  //subtrahend1 := "9637591.879546"
  //subtrahend2 := "28"
  //subtrahend3 := "5284.9765"
  //subtrahend4 := "-189291837.12"
  //subtrahend5 := "7638932.12398765"

  // result = 153,035,620.80650965
  expectedBigINumStr := "153035620.80650965"

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

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]BigIntNum, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  }

  result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(
    minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
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

func TestBigIntMathSubtract_SubtractBigIntNumArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumArray_03"

  var err error

  // minuend =   1,718,973,642.1234567
  minuendStr := "1718973642.1234567"

  // subtrahendStrs
  subtrahendStrs := []string{
    "-28934682.721",
    "424.987654321",
    "-987",
    "62.94",
    "-999999999.99999",
    "-9638932.371",
  }

  //subtrahend0 := "-28934682.721"
  //subtrahend1 := "424.987654321"
  //subtrahend2 := "-987"
  //subtrahend3 := "62.94"
  //subtrahend4 := "-999999999.99999"
  //subtrahend5 := "-9638932.371"

  // Result:  2,757,547,756.287792379
  expectedBigINumStr := "2757547756.287792379"

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

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]BigIntNum, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  }

  result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(
    minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
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

func TestBigIntMathSubtract_SubtractBigIntNumArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumArray_04"

  var err error

  // minuend =   -1,718,973,642.1234567
  minuendStr := "-1718973642.1234567"

  // subtrahendStrs
  subtrahendStrs := []string{
    "-28934682.721",
    "424.987654321",
    "-987",
    "62.94",
    "-999999999.99999",
    "-9638932.371",
  }

  //subtrahend0 := "-28934682.721"
  //subtrahend1 := "424.987654321"
  //subtrahend2 := "-987"
  //subtrahend3 := "62.94"
  //subtrahend4 := "-999999999.99999"
  //subtrahend5 := "-9638932.371"

  // Result:   -680,399,527.959121021
  expectedBigINumStr := "-680399527.959121021"

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

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]BigIntNum, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  }

  result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(
    minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
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

func TestBigIntMathSubtract_SubtractBigIntNumArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumArray_05"

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  // subtrahendStrs
  subtrahendStrs := []string{
    "123.894000",
    "67.1",
    "93.0",
    "-124498.67158",
    "647129.57",
    "28",
  }

  //subtrahend0 := "123.894000"
  //subtrahend1 := "67.1"
  //subtrahend2 := "93.0"
  //subtrahend3 := "-124498.67158"
  //subtrahend4 := "647129.57"
  //subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998,231036"

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

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]BigIntNum, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  }

  result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(
    minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
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

  return
}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_01"

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

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedStrs := len(expectedStrs)

  if lenSubtrahends != lenExpectedStrs {
    t.Errorf("%v\n"+
      "Error: Test is Corrupted!\n"+
      "Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
      "Expected lenExpectedStrs = '%v'\n"+
      "  Actual lenExpectedStrs = '%v'\n\n",
      ePrefix, lenSubtrahends, lenExpectedStrs)

    return
  }

  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())
      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  for k := 0; k < lenSubtrahends; k++ {

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_02"

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

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedStrs := len(expectedStrs)

  if lenSubtrahends != lenExpectedStrs {
    t.Errorf("%v\n"+
      "Error: Test is Corrupted!\n"+
      "Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
      "Expected lenExpectedStrs = '%v'\n"+
      "  Actual lenExpectedStrs = '%v'\n\n",
      ePrefix, lenSubtrahends, lenExpectedStrs)

    return
  }

  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())
      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err := new(BigIntMathSubtract).
    SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  for k := 0; k < lenSubtrahends; k++ {

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_03"

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

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedStrs := len(expectedStrs)

  if lenSubtrahends != lenExpectedStrs {
    t.Errorf("%v\n"+
      "Error: Test is Corrupted!\n"+
      "Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
      "Expected lenExpectedStrs = '%v'\n"+
      "  Actual lenExpectedStrs = '%v'\n\n",
      ePrefix, lenSubtrahends, lenExpectedStrs)

    return
  }

  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())
      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err := new(BigIntMathSubtract).
    SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  for k := 0; k < lenSubtrahends; k++ {

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_04"

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

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedStrs := len(expectedStrs)

  if lenSubtrahends != lenExpectedStrs {
    t.Errorf("%v\n"+
      "Error: Test is Corrupted!\n"+
      "Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
      "Expected lenExpectedStrs = '%v'\n"+
      "  Actual lenExpectedStrs = '%v'\n\n",
      ePrefix, lenSubtrahends, lenExpectedStrs)

    return
  }

  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())
      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err := new(BigIntMathSubtract).
    SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  for k := 0; k < lenSubtrahends; k++ {

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_05"

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

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedStrs := len(expectedStrs)

  if lenSubtrahends != lenExpectedStrs {
    t.Errorf("%v\n"+
      "Error: Test is Corrupted!\n"+
      "Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
      "Expected lenExpectedStrs = '%v'\n"+
      "  Actual lenExpectedStrs = '%v'\n\n",
      ePrefix, lenSubtrahends, lenExpectedStrs)

    return
  }

  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())
      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err := new(BigIntMathSubtract).
    SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  for k := 0; k < lenSubtrahends; k++ {

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_06(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumOutputToArray_06"

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

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedStrs := len(expectedStrs)

  if lenSubtrahends != lenExpectedStrs {
    t.Errorf("%v\n"+
      "Error: Test is Corrupted!\n"+
      "Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
      "Expected lenExpectedStrs = '%v'\n"+
      "  Actual lenExpectedStrs = '%v'\n\n",
      ePrefix, lenSubtrahends, lenExpectedStrs)

    return
  }

  subtrahendAry := make([]BigIntNum, lenSubtrahends)

  expectedResultsAry := make([]BigIntNum, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(BigIntNum).NewNumStrWithNumSeps(subtrahendStrs[i], &usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStrWithNumSeps(subtrahendStrs[%d], &usaNumSeps)\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

    expectedResultsAry[i], err = new(BigIntNum).NewNumStrWithNumSeps(expectedStrs[i], &expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err = new(BigIntNum).\n"+
        "  NewNumStrWithNumSeps(expectedStrs[%d], &expectedNumSeps)\n"+
        "expectedStrs[%d]= '%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], expectedNumSeps.String(), err.Error())
      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err := new(BigIntMathSubtract).
    SubtractBigIntNumOutputToArray(minuendBiNum, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
      "  minuendBiNum, subtrahendAry[...])\n"+
      "minuendBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendBiNumStr,
      err.Error())

    return
  }

  var resultNumStr string

  for j := 0; j < lenSubtrahends; j++ {

    err = result[j].IsValid(fmt.Sprintf("Validating result[%d]", j))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[j].IsValid('Validating result[%d]')\n"+
        "Error= '%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    err = result[j].SetNumericSeparatorsDto(expectedNumSeps)

    if err != nil {

      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].SetNumericSeparatorsDto(expectedNumSeps)\n"+
        "result[%d]= '%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, j, j, resultNumStr, expectedNumSeps.String(), err.Error())

      return
    }

  } // End of Loop

  var expectedResultEqualsResult bool

  var expectedResultNumStr string

  for k := 0; k < lenSubtrahends; k++ {

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractBigIntNumSeries_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumSeries_01"

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
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(\n"+
      " subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
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

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
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

func TestBigIntMathSubtract_SubtractBigIntNumSeries_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumSeries_02"

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
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(\n"+
      " subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
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

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
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

func TestBigIntMathSubtract_SubtractBigIntNumSeries_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumSeries_03"

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
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(\n"+
      " subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
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

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
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

func TestBigIntMathSubtract_SubtractBigIntNumSeries_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumSeries_04"

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
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(\n"+
      " subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
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

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
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

func TestBigIntMathSubtract_SubtractBigIntNumSeries_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractBigIntNumSeries_05"

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

  expectedBigINumSign := 1

  var err error

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

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(BigIntNum).NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(BigIntNum).NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(BigIntNum).NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(BigIntNum).NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(BigIntNum).NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(BigIntNum).NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(
    minuendBiNum,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractBigIntNumSeries(\n"+
      " subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "minuendBiNum= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendBiNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
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

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
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

  return
}

func TestBigIntMathSubtract_SubtractDecimal_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimal_01"

  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99.999"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  decSubtrahend, err := new(Decimal).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahend, err := new(Decimal).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = decSubtrahend.IsValid("Validating decSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSubtrahend.IsValid('Validating decSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSubtrahendNumStr, err := decSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahendNumStr, err := decSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != decSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != decSubtrahendNumStr \n"+
      "Expected decSubtrahendNumStr = '%v'\n"+
      "  Actual decSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, decSubtrahendNumStr)

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

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimals(\n"+
      "  decMinuend, decSubtrahend)\n"+
      "decMinuend= '%v'\n"+
      "decSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
      decSubtrahendNumStr,
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimal_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimal_02"

  // minuend = 949321.6712
  minuendStr := "949321.6712"

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  decSubtrahend, err := new(Decimal).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahend, err := new(Decimal).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = decSubtrahend.IsValid("Validating decSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSubtrahend.IsValid('Validating decSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSubtrahendNumStr, err := decSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahendNumStr, err := decSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != decSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != decSubtrahendNumStr \n"+
      "Expected decSubtrahendNumStr = '%v'\n"+
      "  Actual decSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, decSubtrahendNumStr)

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

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimals(\n"+
      "  decMinuend, decSubtrahend)\n"+
      "decMinuend= '%v'\n"+
      "decSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
      decSubtrahendNumStr,
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimal_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimal_03"

  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"

  expectedBigINumSign := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  decSubtrahend, err := new(Decimal).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahend, err := new(Decimal).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = decSubtrahend.IsValid("Validating decSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSubtrahend.IsValid('Validating decSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSubtrahendNumStr, err := decSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahendNumStr, err := decSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != decSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != decSubtrahendNumStr \n"+
      "Expected decSubtrahendNumStr = '%v'\n"+
      "  Actual decSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, decSubtrahendNumStr)

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

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimals(\n"+
      "  decMinuend, decSubtrahend)\n"+
      "decMinuend= '%v'\n"+
      "decSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
      decSubtrahendNumStr,
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimal_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimal_04"

  // minuend = -289.673849
  minuendStr := "-289.673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  decSubtrahend, err := new(Decimal).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahend, err := new(Decimal).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = decSubtrahend.IsValid("Validating decSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSubtrahend.IsValid('Validating decSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSubtrahendNumStr, err := decSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahendNumStr, err := decSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != decSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != decSubtrahendNumStr \n"+
      "Expected decSubtrahendNumStr = '%v'\n"+
      "  Actual decSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, decSubtrahendNumStr)

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

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimals(\n"+
      "  decMinuend, decSubtrahend)\n"+
      "decMinuend= '%v'\n"+
      "decSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
      decSubtrahendNumStr,
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimal_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimal_05"

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

  decMinuend, err := new(Decimal).NewNumStrWithNumSeps(minuendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, usaNumSeps.String(), err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  decSubtrahend, err := new(Decimal).NewNumStrWithNumSeps(subtrahendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahend, err := new(Decimal).NewNumStrWithNumSeps(subtrahendStr, usaNumSeps)\n"+
      "subtrahendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, usaNumSeps.String(), err.Error())
    return
  }

  err = decSubtrahend.IsValid("Validating decSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decSubtrahend.IsValid('Validating decSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decSubtrahendNumStr, err := decSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decSubtrahendNumStr, err := decSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != decSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != decSubtrahendNumStr \n"+
      "Expected decSubtrahendNumStr = '%v'\n"+
      "  Actual decSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, decSubtrahendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
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

  result, err := new(BigIntMathSubtract).SubtractDecimals(decMinuend, decSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimals(\n"+
      "  decMinuend, decSubtrahend)\n"+
      "decMinuend= '%v'\n"+
      "decSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
      decSubtrahendNumStr,
      err.Error())

    return
  }

  err = result.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimalArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalArray_01"

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  // subtrahendStrs
  subtrahendStrs := []string{
    "123.894000",
    "67.1",
    "93.0",
    "-124498.67158",
    "647129.57",
    "28",
  }

  //subtrahend0 := "123.894000"
  //subtrahend1 := "67.1"
  //subtrahend2 := "93.0"
  //subtrahend3 := "-124498.67158"
  //subtrahend4 := "647129.57"
  //subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998.231036"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

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

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]Decimal, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(
    decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimalArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalArray_02"

  var err error

  // minuend = -18,973,642.1234567
  minuendStr := "-18973642.1234567"

  // subtrahendStrs
  subtrahendStrs := []string{
    "737.21",
    "9637591.879546",
    "28",
    "5284.9765",
    "-189291837.12",
    "7638932.12398765",
  }

  //subtrahend0 := "737.21"
  //subtrahend1 := "9637591.879546"
  //subtrahend2 := "28"
  //subtrahend3 := "5284.9765"
  //subtrahend4 := "-189291837.12"
  //subtrahend5 := "7638932.12398765"

  // result = 153,035,620.80650965
  expectedBigINumStr := "153035620.80650965"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

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

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]Decimal, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(
    decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimalArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalArray_03"

  var err error

  // minuend =   1,718,973,642.1234567
  minuendStr := "1718973642.1234567"

  // subtrahendStrs
  subtrahendStrs := []string{
    "-28934682.721",
    "424.987654321",
    "-987",
    "62.94",
    "-999999999.99999",
    "-9638932.371",
  }

  //subtrahend0 := "-28934682.721"
  //subtrahend1 := "424.987654321"
  //subtrahend2 := "-987"
  //subtrahend3 := "62.94"
  //subtrahend4 := "-999999999.99999"
  //subtrahend5 := "-9638932.371"

  // Result:  2,757,547,756.287792379
  expectedBigINumStr := "2757547756.287792379"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

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

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]Decimal, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(
    decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimalArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalArray_04"

  var err error

  // minuend =   -1,718,973,642.1234567
  minuendStr := "-1718973642.1234567"

  // subtrahendStrs
  subtrahendStrs := []string{
    "-28934682.721",
    "424.987654321",
    "-987",
    "62.94",
    "-999999999.99999",
    "-9638932.371",
  }

  //subtrahend0 := "-28934682.721"
  //subtrahend1 := "424.987654321"
  //subtrahend2 := "-987"
  //subtrahend3 := "62.94"
  //subtrahend4 := "-999999999.99999"
  //subtrahend5 := "-9638932.371"

  // Result:   -680,399,527.959121021
  expectedBigINumStr := "-680399527.959121021"

  expectedBigINumSign := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

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

  lenSubtrahendsArray := 6

  subtrahendAry := make([]Decimal, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimalArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalArray_05"

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  // subtrahendStrs
  subtrahendStrs := []string{
    "123.894000",
    "67.1",
    "93.0",
    "-124498.67158",
    "647129.57",
    "28",
  }

  //subtrahend0 := "123.894000"
  //subtrahend1 := "67.1"
  //subtrahend2 := "93.0"
  //subtrahend3 := "-124498.67158"
  //subtrahend4 := "647129.57"
  //subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998,231036"

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

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStrWithNumSeps(minuendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, usaNumSeps.String(), err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

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

  lenSubtrahendsArray := 6

  subtrahendAry := make([]Decimal, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractDecimalArray(
    decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuendNumStr,
      err.Error())

    return
  }

  err = result.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}
