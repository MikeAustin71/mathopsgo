package mathops

import "testing"

func TestBigIntNum_SetPrecision_01(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_01"

  originalNumStr := "654.123456"

  newPrecision := uint(3)

  expectedNumberStr := "654.123"

  expectedPrecisionUint := uint(3)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_02(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_02"

  originalNumStr := "654.123456"

  newPrecision := uint(4)

  expectedNumberStr := "654.1235"

  expectedPrecisionUint := uint(4)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_03(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_03"

  originalNumStr := "654.123456"

  newPrecision := uint(0)

  expectedNumberStr := "654"

  expectedPrecisionUint := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_04(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_04"

  originalNumStr := "-654.123456"

  newPrecision := uint(3)

  expectedNumberStr := "-654.123"

  expectedPrecisionUint := uint(3)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_05(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_05"

  originalNumStr := "-654.123456"

  newPrecision := uint(4)

  expectedNumberStr := "-654.1235"

  expectedPrecisionUint := uint(4)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_06(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_06"

  originalNumStr := "654"

  newPrecision := uint(3)

  expectedNumberStr := "654.000"

  expectedPrecisionUint := uint(3)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_07(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_07"

  originalNumStr := "654.123456"

  newPrecision := uint(9)

  expectedNumberStr := "654.123456000"

  expectedPrecisionUint := uint(9)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_08(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_08"

  originalNumStr := "-654"

  newPrecision := uint(9)

  expectedNumberStr := "-654.000000000"

  expectedPrecisionUint := uint(9)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_09(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_09"

  originalNumStr := "-654.123456"

  newPrecision := uint(9)

  expectedNumberStr := "-654.123456000"

  expectedPrecisionUint := uint(9)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_10(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_10"

  originalNumStr := "0"

  newPrecision := uint(4)

  expectedNumberStr := "0.0000"

  expectedPrecisionUint := uint(4)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_SetPrecision_11(t *testing.T) {

  ePrefix := "TestBigIntNum_SetPrecision_11"

  originalNumStr := "0.0000"

  newPrecision := uint(0)

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
      "expectedNumberStr='%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumberStr,
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

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumberStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumPrecisionUint, err :=\n"+
      "  expectedBigINum.GetPrecisionUint()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumberStr, err.Error())

    return
  }

  if expectedPrecisionUint != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
      "originalNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStr,
      expectedNumSeps.String(),
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

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumNumberStr)

    return
  }

  err = bINum.SetPrecision(newPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, newPrecision, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum After Precision")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum After Precision')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err = bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "After new 'bINum' Precision\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
      "Because expectedBigINumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  if expectedBigINumPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

    return
  }

  expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " expectedBigINumEqualsBINum, err :=\n"+
      "  expectedBigINum.Equal(bINum)\n"+
      "expectedBigINum= '%v'\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
    return
  }

  if !expectedBigINumEqualsBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'bINum' values ARE NOT Equal\n"+
      "Because expectedBigINumEqualsBINum = 'false' \n"+
      "Expected bINum = '%v'\n"+
      "  Actual bINum = '%v'\n\n",
      ePrefix, expectedBigINumberStr, bINumNumberStr)

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

  return
}

func TestBigIntNum_TrimTrailingFracZeros_01(t *testing.T) {

  nStr := "-123.000"
  expectedNumStr := "-123"

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TrimTrailingFracZeros()

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntNum_TrimTrailingFracZeros_02(t *testing.T) {

  nStr := "123.000"
  expectedNumStr := "123"

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TrimTrailingFracZeros()

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntNum_TrimTrailingFracZeros_03(t *testing.T) {

  nStr := "123.0090"
  expectedNumStr := "123.009"

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TrimTrailingFracZeros()

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntNum_TrimTrailingFracZeros_04(t *testing.T) {

  nStr := "-123.0090"
  expectedNumStr := "-123.009"

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TrimTrailingFracZeros()

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntNum_TrimTrailingFracZeros_05(t *testing.T) {

  nStr := "0.000"
  expectedNumStr := "0"

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TrimTrailingFracZeros()

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntNum_TruncToDecPlace_01(t *testing.T) {

  nStr := "-123.567"
  expectedNumStr := "-123.56"
  truncToDec := uint(2)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_02(t *testing.T) {

  nStr := "123.567"
  expectedNumStr := "123.56"
  truncToDec := uint(2)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_03(t *testing.T) {

  nStr := "123.567"
  expectedNumStr := "123.567"
  truncToDec := uint(3)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_04(t *testing.T) {

  nStr := "123.567"
  expectedNumStr := "123.5670"
  truncToPlace := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToPlace)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_05(t *testing.T) {

  nStr := "-123.567"
  expectedNumStr := "-123.5670"
  truncToPlace := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToPlace)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_06(t *testing.T) {

  nStr := "0.000"
  expectedNumStr := "0.00"
  truncToDec := uint(2)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_07(t *testing.T) {

  nStr := "654.123456"
  expectedNumStr := "654.123"
  truncToDec := uint(3)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_08(t *testing.T) {

  nStr := "654.123456789"
  expectedNumStr := "654.1234"
  truncToDec := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_09(t *testing.T) {

  nStr := "654.123456789"
  expectedNumStr := "654"
  truncToDec := uint(0)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_10(t *testing.T) {

  nStr := "654"
  expectedNumStr := "654.00000"
  truncToDec := uint(5)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_11(t *testing.T) {

  nStr := "654.123"
  expectedNumStr := "654.123000000"
  truncToDec := uint(9)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_12(t *testing.T) {

  nStr := "0"
  expectedNumStr := "0.000000"
  truncToDec := uint(6)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_13(t *testing.T) {

  nStr := "0.000000"
  expectedNumStr := "0"
  truncToDec := uint(0)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_TruncToDecPlace_14(t *testing.T) {

  nStr := "654.123456789015"
  expectedNumStr := "654.12345678901"
  truncToDec := uint(11)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.TruncToDecPlace(truncToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}
