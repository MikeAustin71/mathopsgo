package mathops

import "testing"

func TestBigIntMathSubtract_SubtractNumStr_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStr_01"

  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99.999"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "subtrahendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      subtrahendStr,
      usaNumSeps.String(),
      err.Error())

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

func TestBigIntMathSubtract_SubtractNumStr_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStr_02"

  var err error

  // minuend = 949321.6712
  minuendStr := "949321.6712"

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "subtrahendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      subtrahendStr,
      usaNumSeps.String(),
      err.Error())

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

func TestBigIntMathSubtract_SubtractNumStr_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStr_03"

  var err error

  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"

  expectedBigINumSign := -1

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

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

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "subtrahendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      subtrahendStr,
      usaNumSeps.String(),
      err.Error())

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

func TestBigIntMathSubtract_SubtractNumStr_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStr_04"

  var err error

  // minuend = -289.673849
  minuendStr := "-289.673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "subtrahendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      subtrahendStr,
      usaNumSeps.String(),
      err.Error())

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

func TestBigIntMathSubtract_SubtractNumStr_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStr_05"

  var err error

  // minuend = -289,673849
  minuendStr := "-289,673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579,012"

  // result = 14289.338151
  expectedBigINumStr := "14289,338151"

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

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStr(minuendStr, subtrahendStr, expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "subtrahendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      subtrahendStr,
      expectedNumSeps.String(),
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

func TestBigIntMathSubtract_SubtractNumStrArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_01"

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

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  lenSubtrahends := 6

  subtrahendAry := make([]string, lenSubtrahends)

  subtrahendAry[0] = subtrahend0
  subtrahendAry[1] = subtrahend1
  subtrahendAry[2] = subtrahend2
  subtrahendAry[3] = subtrahend3
  subtrahendAry[4] = subtrahend4
  subtrahendAry[5] = subtrahend5

  result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
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

func TestBigIntMathSubtract_SubtractNumStrArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_02"

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

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  lenSubtrahends := 6

  subtrahendAry := make([]string, lenSubtrahends)

  subtrahendAry[0] = subtrahend0
  subtrahendAry[1] = subtrahend1
  subtrahendAry[2] = subtrahend2
  subtrahendAry[3] = subtrahend3
  subtrahendAry[4] = subtrahend4
  subtrahendAry[5] = subtrahend5

  result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
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

func TestBigIntMathSubtract_SubtractNumStrArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_03"

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

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  lenSubtrahends := 6

  subtrahendAry := make([]string, lenSubtrahends)

  subtrahendAry[0] = subtrahend0
  subtrahendAry[1] = subtrahend1
  subtrahendAry[2] = subtrahend2
  subtrahendAry[3] = subtrahend3
  subtrahendAry[4] = subtrahend4
  subtrahendAry[5] = subtrahend5

  result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
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

func TestBigIntMathSubtract_SubtractNumStrArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_04"

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

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  lenSubtrahends := 6
  subtrahendAry := make([]string, lenSubtrahends)

  subtrahendAry[0] = subtrahend0
  subtrahendAry[1] = subtrahend1
  subtrahendAry[2] = subtrahend2
  subtrahendAry[3] = subtrahend3
  subtrahendAry[4] = subtrahend4
  subtrahendAry[5] = subtrahend5

  result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
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

func TestBigIntMathSubtract_SubtractNumStrArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_05"

  var err error

  // minuend =   -1 718 973 642,1234567
  minuendStr := "-1718973642,1234567"

  subtrahend0 := "-28934682,721"
  subtrahend1 := "424,987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62,94"
  subtrahend4 := "-999999999,99999"
  subtrahend5 := "-9638932,371"

  // Result:   -680 399 527,959121021
  expectedBigINumStr := "-680399527,959121021"

  expectedBigINumSign := -1

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

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumStr,
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

  lenSubtrahends := 6

  subtrahendAry := make([]string, lenSubtrahends)

  subtrahendAry[0] = subtrahend0
  subtrahendAry[1] = subtrahend1
  subtrahendAry[2] = subtrahend2
  subtrahendAry[3] = subtrahend3
  subtrahendAry[4] = subtrahend4
  subtrahendAry[5] = subtrahend5

  result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
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

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_01"

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

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  subtrahendAry := make([]string, lenSubtrahends)

  expectedResultsAry := make([]string, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i] = subtrahendStrs[i]

    expectedResultsAry[i] = expectedStrs[i]

  }

  result, err :=
    new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrOutputToArray(minuendNumStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  for k := 0; k < lenSubtrahends; k++ {

    if expectedResultsAry[k] != result[k] {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because result[%d] != expectedResultsAry[%d]\n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

      return
    }

  }

  return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_02"

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

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  subtrahendAry := make([]string, lenSubtrahends)

  expectedResultsAry := make([]string, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i] = subtrahendStrs[i]

    expectedResultsAry[i] = expectedStrs[i]

  }

  result, err :=
    new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  for k := 0; k < lenSubtrahends; k++ {

    if expectedResultsAry[k] != result[k] {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because result[%d] != expectedResultsAry[%d]\n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

      return
    }

  }

  return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_03"

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

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  subtrahendAry := make([]string, lenSubtrahends)

  expectedResultsAry := make([]string, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i] = subtrahendStrs[i]

    expectedResultsAry[i] = expectedStrs[i]

  }

  result, err :=
    new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  for k := 0; k < lenSubtrahends; k++ {

    if expectedResultsAry[k] != result[k] {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because result[%d] != expectedResultsAry[%d]\n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

      return
    }

  }

  return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_04"

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

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  subtrahendAry := make([]string, lenSubtrahends)

  expectedResultsAry := make([]string, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i] = subtrahendStrs[i]

    expectedResultsAry[i] = expectedStrs[i]

  }

  result, err :=
    new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  for k := 0; k < lenSubtrahends; k++ {

    if expectedResultsAry[k] != result[k] {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because result[%d] != expectedResultsAry[%d]\n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

      return
    }

  }

  return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_05"

  var err error

  // minuend =   0
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

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  subtrahendAry := make([]string, lenSubtrahends)

  expectedResultsAry := make([]string, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i] = subtrahendStrs[i]

    expectedResultsAry[i] = expectedStrs[i]

  }

  result, err :=
    new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  for k := 0; k < lenSubtrahends; k++ {

    if expectedResultsAry[k] != result[k] {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because result[%d] != expectedResultsAry[%d]\n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

      return
    }

  }

  return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_06(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_06"

  var err error

  // minuend =   5051
  minuendStr := "5051"

  subtrahendStrs := []string{
    "8000",
    "6051,123456",
    "-30871,25",
    "604,55",
    "9100,123",
    "-115,76",
  }

  expectedStrs := []string{
    "-2949",
    "-1000,123456",
    "35922,25",
    "4446,45",
    "-4049,123",
    "5166,76",
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

  lenSubtrahends := len(subtrahendStrs)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  subtrahendAry := make([]string, lenSubtrahends)

  expectedResultsAry := make([]string, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i] = subtrahendStrs[i]

    expectedResultsAry[i] = expectedStrs[i]

  }

  result, err :=
    new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).\n"+
      "  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  for k := 0; k < lenSubtrahends; k++ {

    if expectedResultsAry[k] != result[k] {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because result[%d] != expectedResultsAry[%d]\n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

      return
    }

  }

  return
}

func TestBigIntMathSubtract_SubtractNumStrSeries_01(t *testing.T) {

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
  expectedNumStr := "6805998.231036"
  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).New()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
    expectedNumSeps,
    minuendStr,
    subtrahend0,
    subtrahend1,
    subtrahend2,
    subtrahend3,
    subtrahend4,
    subtrahend5)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrSeries("+
      "minuendStr, ...). Error='%v' ", err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)

  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathSubtract_SubtractNumStrSeries_02(t *testing.T) {

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
  expectedNumStr := "153035620.80650965"
  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).New()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' Error='%v' ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
    expectedNumSeps,
    minuendStr,
    subtrahend0,
    subtrahend1,
    subtrahend2,
    subtrahend3,
    subtrahend4,
    subtrahend5)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrSeries("+
      "minuendStr, ...). Error='%v' ", err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)

  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathSubtract_SubtractNumStrSeries_03(t *testing.T) {

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
  expectedNumStr := "2757547756.287792379"
  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).New()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' Error='%v' ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
    expectedNumSeps,
    minuendStr,
    subtrahend0,
    subtrahend1,
    subtrahend2,
    subtrahend3,
    subtrahend4,
    subtrahend5)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrSeries("+
      "minuendStr, ...). Error='%v' ", err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)

  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathSubtract_SubtractNumStrSeries_04(t *testing.T) {

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
  expectedNumStr := "-680399527.959121021"
  expectedBigINumSign := -1

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' Error='%v' ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
    expectedNumSeps,
    minuendStr,
    subtrahend0,
    subtrahend1,
    subtrahend2,
    subtrahend3,
    subtrahend4,
    subtrahend5)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrSeries("+
      "minuendStr, ...). Error='%v' ", err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)

  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathSubtract_SubtractNumStrSeries_05(t *testing.T) {

  var err error

  // minuend = 7328941,123456
  minuendStr := "7328941,123456"

  subtrahend0 := "123,894000"
  subtrahend1 := "67,1"
  subtrahend2 := "93,0"
  subtrahend3 := "-124498,67158"
  subtrahend4 := "647129,57"
  subtrahend5 := "28"

  // result = 6805998,231036
  expectedNumStr := "6805998,231036"
  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
    expectedNumSeps,
    minuendStr,
    subtrahend0,
    subtrahend1,
    subtrahend2,
    subtrahend3,
    subtrahend4,
    subtrahend5)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrSeries("+
      "minuendStr, ...). Error='%v' ", err.Error())
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)

  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathSubtract_SubtractNumStrDto_01(t *testing.T) {
  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99.999"

  expectedBigINumSign := 1

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  nDtoSubtrahend, err := NumStrDto{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, "+
      "nDtoSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
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

func TestBigIntMathSubtract_SubtractNumStrDto_02(t *testing.T) {
  // minuend = 949321.6712
  minuendStr := "949321.6712"

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"
  expectedBigINumSign := 1

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  nDtoSubtrahend, err := NumStrDto{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, "+
      "nDtoSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
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

func TestBigIntMathSubtract_SubtractNumStrDto_03(t *testing.T) {
  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"
  expectedBigINumSign := -1

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  nDtoSubtrahend, err := NumStrDto{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, "+
      "nDtoSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
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

func TestBigIntMathSubtract_SubtractNumStrDto_04(t *testing.T) {
  // minuend = -289.673849
  minuendStr := "-289.673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"
  expectedBigINumSign := 1

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  nDtoSubtrahend, err := NumStrDto{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, "+
      "nDtoSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
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

func TestBigIntMathSubtract_SubtractNumStrDto_05(t *testing.T) {
  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedNumStr := "99,999"

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v'", minuendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  nDtoSubtrahend, err := NumStrDto{}.NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStr) "+
      "subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, "+
      "nDtoSubtrahend) minuendStr='%v' subtrahendStr='%v' Error='%v' ",
      minuendStr, subtrahendStr, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathSubtract_SubtractNumStrDtoArray_01(t *testing.T) {

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

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalArray("+
      "nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractNumStrDtoArray_02(t *testing.T) {

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

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDtoArray("+
      "nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractNumStrDtoArray_03(t *testing.T) {

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

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDtoArray("+
      "nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractNumStrDtoArray_04(t *testing.T) {

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

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDtoArray("+
      "nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractNumStrDtoArray_05(t *testing.T) {

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"
  subtrahend1 := "67.1"
  subtrahend2 := "93.0"
  subtrahend3 := "-124498.67158"
  subtrahend4 := "647129.57"
  subtrahend5 := "28"

  // result = 6805998,231036
  expectedNumStr := "6805998,231036"

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalArray("+
      "nDtoMinuend, subtrahendAry). Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr= '%v'. ",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_01(t *testing.T) {

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

  minuendNumStrDto, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]NumStrDto, lenSubtrahends)
  expectedResultsAry := make([]NumStrDto, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = NumStrDto{}.NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = NumStrDto{}.NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }
  }

  resultArray, err :=
    new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractNumStrDtoOutputToArray"+
      "(minuendNumStrDto, subtrahendAry) minuendNumStrDto='%v'  Error='%v'. ",
      minuendNumStrDto.GetNumStr(), err.Error())
  }

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Inequality Error: Expected ResultsAry='%v'. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }
  }
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_02(t *testing.T) {

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

  minuendNumStrDto, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]NumStrDto, lenSubtrahends)
  expectedResultsAry := make([]NumStrDto, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = NumStrDto{}.NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = NumStrDto{}.NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }

  }

  resultArray, err :=
    new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractNumStrDtoOutputToArray"+
      "(minuendNumStrDto, subtrahendAry) minuendNumStrDto='%v'  Error='%v'. ",
      minuendNumStrDto.GetNumStr(), err.Error())
  }

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }

  }

}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_03(t *testing.T) {

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

  minuendNumStrDto, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]NumStrDto, lenSubtrahends)
  expectedResultsAry := make([]NumStrDto, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = NumStrDto{}.NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = NumStrDto{}.NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }

  }

  resultArray, err :=
    new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractNumStrDtoOutputToArray"+
      "(minuendNumStrDto, subtrahendAry) minuendNumStrDto='%v'  Error='%v'. ",
      minuendNumStrDto.GetNumStr(), err.Error())
  }

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }

  }
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_04(t *testing.T) {

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

  minuendNumStrDto, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]NumStrDto, lenSubtrahends)
  expectedResultsAry := make([]NumStrDto, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = NumStrDto{}.NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = NumStrDto{}.NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }

  }

  resultArray, err :=
    new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractNumStrDtoOutputToArray"+
      "(minuendNumStrDto, subtrahendAry) minuendNumStrDto='%v'  Error='%v'. ",
      minuendNumStrDto.GetNumStr(), err.Error())
  }

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }
  }
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_05(t *testing.T) {

  var err error

  // minuend =   0
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

  minuendNumStrDto, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]NumStrDto, lenSubtrahends)
  expectedResultsAry := make([]NumStrDto, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = NumStrDto{}.NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = NumStrDto{}.NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }

  }

  resultArray, err :=
    new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractNumStrDtoOutputToArray"+
      "(minuendNumStrDto, subtrahendAry) minuendNumStrDto='%v'  Error='%v'. ",
      minuendNumStrDto.GetNumStr(), err.Error())
  }

  for k := 0; k < lenSubtrahends; k++ {

    if !resultArray[k].Equal(expectedResultsAry[k]) {
      t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
        expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
    }

  }
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_06(t *testing.T) {

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

  minuendNumStrDto, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = minuendNumStrDto.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenSubtrahends := len(subtrahendStrs)
  subtrahendAry := make([]NumStrDto, lenSubtrahends)
  expectedResultsAry := make([]NumStrDto, lenSubtrahends)

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = NumStrDto{}.NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(subtrahendStrs[i]) "+
        "subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
    }

    expectedResultsAry[i], err = NumStrDto{}.NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("Error returned by NumStrDto{}.NewNumStr(expectedStrs[i]) "+
        "expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
    }
  }

  resultArray, err :=
    new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

  if err != nil {
    t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractNumStrDtoOutputToArray"+
      "(minuendNumStrDto, subtrahendAry) minuendNumStrDto='%v'  Error='%v'. ",
      minuendNumStrDto.GetNumStr(), err.Error())
  }

  for k := 0; k < lenSubtrahends; k++ {

    actualNumStr := resultArray[k].GetNumStr()

    if expectedStrs[k] != actualNumStr {
      t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. Index='%v'",
        expectedStrs[k], actualNumStr, k)
    }

    actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. Index='%v'",
        expectedNumSeps.String(), actualNumSeps.String(), k)
    }
  }
}

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_01(t *testing.T) {

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

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
    nDtoMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDtoSeries("+
      "nDtoMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_02(t *testing.T) {

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

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
    nDtoMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDtoSeries("+
      "nDtoMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_03(t *testing.T) {

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

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
    nDtoMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDtoSeries("+
      "nDtoMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_04(t *testing.T) {

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

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
    nDtoMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDtoSeries("+
      "nDtoMinuend, ...). Error='%v' ", err.Error())
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

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_05(t *testing.T) {

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"
  subtrahend1 := "67.1"
  subtrahend2 := "93.0"
  subtrahend3 := "-124498.67158"
  subtrahend4 := "647129.57"
  subtrahend5 := "28"

  // result = 6805998,231036
  expectedNumStr := "6805998,231036"

  nDtoMinuend, err := NumStrDto{}.NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by nDtoMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenSubtrahends := 6
  subtrahendAry := make([]NumStrDto, lenSubtrahends)

  subtrahendAry[0], err = NumStrDto{}.NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
  }

  subtrahendAry[1], err = NumStrDto{}.NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
  }

  subtrahendAry[2], err = NumStrDto{}.NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
  }

  subtrahendAry[3], err = NumStrDto{}.NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
  }

  subtrahendAry[4], err = NumStrDto{}.NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
  }

  subtrahendAry[5], err = NumStrDto{}.NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
  }

  result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
    nDtoMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathSubtract).SubtractNumStrDtoSeries("+
      "nDtoMinuend, ...). Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}
