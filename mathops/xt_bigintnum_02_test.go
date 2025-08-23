package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntNum_Floor_01(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_01"

  var err error

  originalNumStr := "5.95"

  expectedNumStr := "5"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_02(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_02"

  var err error

  originalNumStr := "5.05"

  expectedNumStr := "5"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_03(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_03"

  var err error

  originalNumStr := "5"

  expectedNumStr := "5"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_04(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_04"

  var err error

  originalNumStr := "-5.05"

  expectedNumStr := "-6"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_05(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_05"

  var err error

  originalNumStr := "2.4"

  expectedNumStr := "2"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_06(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_06"

  var err error

  originalNumStr := "2.9"

  expectedNumStr := "2"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_07(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_07"

  var err error

  originalNumStr := "-2.7"

  expectedNumStr := "-3"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_08(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_08"

  var err error

  originalNumStr := "-2"

  expectedNumStr := "-2"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_09(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_09"

  var err error

  originalNumStr := "0"

  expectedNumStr := "0"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_10(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_10"

  var err error

  originalNumStr := "18972.0000000000001"

  expectedNumStr := "18972"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_11(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_11"

  var err error

  originalNumStr := "-18972.0000000000001"

  expectedNumStr := "-18973"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_12(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_12"

  var err error

  originalNumStr := "0.0000000000001"

  expectedNumStr := "0"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_13(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_13"

  var err error

  originalNumStr := "-0.0000000000001"

  expectedNumStr := "-1"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_14(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_14"

  var err error

  originalNumStr := "-189765342891.0000000000001"

  expectedNumStr := "-189765342892"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_15(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_15"

  var err error

  originalNumStr := "189765342891.0000000000001"

  expectedNumStr := "189765342891"

  expectedPrecision := uint(0)

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
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

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

    return
  }

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
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_Floor_16(t *testing.T) {

  ePrefix := "TestBigIntNum_Floor_15"

  var err error

  originalNumStr := "189765342891,0000000000001"

  expectedNumStr := "189765342891"

  expectedPrecision := uint(0)

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

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr, &expectedNumSeps)\n"+
      "expectedNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumSeps.String(), err.Error())
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

  if expectedNumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumberStr)

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

  if expectedPrecision != expectedBigINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
      "Expected expectedBigINumPrecisionUint = '%v'\n"+
      "  Actual expectedBigINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

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
      ePrefix, originalNumStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum')\n"+
      "originalNumStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
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
      "Because originalNumStr != bINumOriginalNumStr \n"+
      "Expected bINumOriginalNumStr = '%v'\n"+
      "  Actual bINumOriginalNumStr = '%v'\n\n",
      ePrefix, originalNumStr, bINumStr)

    return
  }

  floorBINum, err := bINum.Floor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINum, err := bINum.Floor()\n"+
      "bINumOriginalNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumStr, err.Error())
    return
  }

  err = floorBINum.IsValid("Validating Final floorBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = floorBINum.IsValid('Validating floorBINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumStr, err := floorBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumStr, err := floorBINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
    return
  }

  floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, floorBINumStr, err.Error())
    return
  }

  expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsFloorBINum, err :=\n"+
      "  expectedBigINum.Equal(floorBINum)\n"+
      "expectedBigINum= '%v'\n"+
      "floorBINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
    return
  }

  if !expectedEqualsFloorBINum {
    t.Errorf("%v\n"+
      "Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
      "Because expectedEqualsBINum = 'false' \n"+
      "Expected floorBINum = '%v'\n"+
      "  Actual floorBINum = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedNumStr != floorBINumStr {
    t.Errorf("%v\n"+
      "Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
      "Because expectedNumStr != floorBINumStr\n"+
      "Expected floorBINumStr = '%v'\n"+
      "  Actual floorBINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, floorBINumStr)

    return
  }

  if expectedPrecision != floorBINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecision != floorBINumPrecisionUint\n"+
      "Expected floorBINumPrecisionUint = '%v'\n"+
      "  Actual floorBINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, floorBINumPrecisionUint)

    return
  }

  if !expectedNumSeps.Equal(floorBINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal\n"+
      "Because expectedNumSeps != floorBINumSeps \n"+
      "Expected floorBINumSeps = '%v'\n"+
      "  Actual floorBINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_FormatNumStr_01(t *testing.T) {

  originalNumStr := "-123.45"
  expectedNumStr := "-123.45"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_02(t *testing.T) {

  originalNumStr := "123.45"
  expectedNumStr := "123.45"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_03(t *testing.T) {

  originalNumStr := "-123.45"
  expectedNumStr := "(123.45)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_04(t *testing.T) {

  originalNumStr := "-1234.56"
  expectedNumStr := "-1234.56"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_05(t *testing.T) {

  originalNumStr := "1234.56"
  expectedNumStr := "1234.56"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_06(t *testing.T) {

  originalNumStr := "-1234.56"
  expectedNumStr := "(1234.56)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_07(t *testing.T) {

  originalNumStr := "0"
  expectedNumStr := "0"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_08(t *testing.T) {

  originalNumStr := "0.000"
  expectedNumStr := "0.000"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_09(t *testing.T) {

  originalNumStr := "0.000"
  expectedNumStr := "0.000"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_12(t *testing.T) {

  originalNumStr := "12345"
  expectedNumStr := "12345"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_13(t *testing.T) {

  originalNumStr := "-12345"
  expectedNumStr := "(12345)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_14(t *testing.T) {

  originalNumStr := "-12345"
  expectedNumStr := "-12345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_15(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(8)
  expectedNumStr := "0.00012345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_16(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(5)
  expectedNumStr := "0.12345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_17(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(8)
  expectedNumStr := "-0.00012345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_18(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(5)
  expectedNumStr := "-0.12345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatNumStr_19(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(8)
  expectedNumStr := "(0.00012345)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatNumStr_20(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(5)
  expectedNumStr := "(0.12345)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatNumStr_21(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(5)
  expectedNumStr := "12345"
  mode := ABSOLUTEPURENUMSTRFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatNumStr_22(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(2)
  expectedNumStr := "12345"
  mode := ABSOLUTEPURENUMSTRFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatNumStr_23(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(7)
  expectedNumStr := "0012345"
  mode := ABSOLUTEPURENUMSTRFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatNumStr_24(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(2)
  expectedNumStr := "12345"
  mode := ABSOLUTEPURENUMSTRFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatNumStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatThousandsStr_01(t *testing.T) {

  originalNumStr := "1234"
  expectedNumStr := "1,234"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_02(t *testing.T) {

  originalNumStr := "123"
  expectedNumStr := "123"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_03(t *testing.T) {

  originalNumStr := "-1234"
  expectedNumStr := "(1,234)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_04(t *testing.T) {

  originalNumStr := "-1234"
  expectedNumStr := "-1,234"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_05(t *testing.T) {

  originalNumStr := "1234.567"
  expectedNumStr := "1,234.567"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_06(t *testing.T) {

  originalNumStr := "-1234.567"
  expectedNumStr := "-1,234.567"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_07(t *testing.T) {

  originalNumStr := "-1234.567"
  expectedNumStr := "(1,234.567)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_08(t *testing.T) {

  originalNumStr := "0"
  expectedNumStr := "0"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_09(t *testing.T) {

  originalNumStr := "0.0000"
  expectedNumStr := "0.0000"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_10(t *testing.T) {

  originalNumStr := "0.0000"
  expectedNumStr := "0.0000"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_11(t *testing.T) {

  originalNumStr := "1234567890.12"
  expectedNumStr := "1,234,567,890.12"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_12(t *testing.T) {

  originalNumStr := "-1234567890.12"
  expectedNumStr := "-1,234,567,890.12"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_13(t *testing.T) {

  originalNumStr := "-1234567890.12"
  expectedNumStr := "(1,234,567,890.12)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_14(t *testing.T) {

  originalNumStr := "1234567890"
  expectedNumStr := "1,234,567,890"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_15(t *testing.T) {

  originalNumStr := "-1234567890"
  expectedNumStr := "-1,234,567,890"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_16(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(8)
  expectedNumStr := "0.00012345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_17(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(5)
  expectedNumStr := "0.12345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_18(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(8)
  expectedNumStr := "-0.00012345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_19(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(5)
  expectedNumStr := "-0.12345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatThousandsStr_20(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(8)
  expectedNumStr := "(0.00012345)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatThousandsStr_21(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(5)
  expectedNumStr := "(0.12345)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatThousandsStr_22(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(5)
  expectedNumStr := "12345"
  mode := ABSOLUTEPURENUMSTRFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatThousandsStr_23(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(0)
  expectedNumStr := "12,345"
  mode := ABSOLUTEPURENUMSTRFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatThousandsStr_24(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(0)
  expectedNumStr := "12,345"
  mode := ABSOLUTEPURENUMSTRFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatThousandsStr_25(t *testing.T) {

  originalBInt := big.NewInt(123)
  precision := uint(0)
  expectedNumStr := "123"
  mode := ABSOLUTEPURENUMSTRFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatThousandsStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatCurrencyStr_01(t *testing.T) {

  originalNumStr := "1234"
  expectedNumStr := "$1,234"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_02(t *testing.T) {

  originalNumStr := "123"
  expectedNumStr := "$123"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_03(t *testing.T) {

  originalNumStr := "-1234"
  expectedNumStr := "($1,234)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_04(t *testing.T) {

  originalNumStr := "-1234"
  expectedNumStr := "-$1,234"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_05(t *testing.T) {

  originalNumStr := "1234.567"
  expectedNumStr := "$1,234.567"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_06(t *testing.T) {

  originalNumStr := "-1234.567"
  expectedNumStr := "-$1,234.567"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_07(t *testing.T) {

  originalNumStr := "-1234.567"
  expectedNumStr := "($1,234.567)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_08(t *testing.T) {

  originalNumStr := "0"
  expectedNumStr := "$0"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_09(t *testing.T) {

  originalNumStr := "0.0000"
  expectedNumStr := "$0.0000"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_10(t *testing.T) {

  originalNumStr := "0.0000"
  expectedNumStr := "$0.0000"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_11(t *testing.T) {

  originalNumStr := "1234567890.12"
  expectedNumStr := "$1,234,567,890.12"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_12(t *testing.T) {

  originalNumStr := "-1234567890.12"
  expectedNumStr := "-$1,234,567,890.12"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_13(t *testing.T) {

  originalNumStr := "-1234567890.12"
  expectedNumStr := "($1,234,567,890.12)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_14(t *testing.T) {

  originalNumStr := "1234567890"
  expectedNumStr := "$1,234,567,890"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_15(t *testing.T) {

  originalNumStr := "-1234567890"
  expectedNumStr := "-$1,234,567,890"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_16(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(8)
  expectedNumStr := "$0.00012345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_17(t *testing.T) {

  originalBInt := big.NewInt(12345)
  precision := uint(5)
  expectedNumStr := "$0.12345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_18(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(8)
  expectedNumStr := "-$0.00012345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_19(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(5)
  expectedNumStr := "-$0.12345"
  mode := LEADMINUSNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatCurrencyStr_20(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(8)
  expectedNumStr := "($0.00012345)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_21(t *testing.T) {

  originalBInt := big.NewInt(-12345)
  precision := uint(5)
  expectedNumStr := "($0.12345)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum := new(BigIntNum).NewBigInt(originalBInt, precision)

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }
}

func TestBigIntNum_FormatCurrencyStr_22(t *testing.T) {

  originalNumStr := "-1234567890"
  expectedNumStr := "-£1,234,567,890"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  bINum.SetCurrencySymbol('\U000000a3')

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_23(t *testing.T) {

  originalNumStr := "1234567890"
  expectedNumStr := "£1,234,567,890"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  bINum.SetCurrencySymbol('\U000000a3')

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_24(t *testing.T) {

  originalNumStr := "1234567890.12"
  expectedNumStr := "£1,234,567,890.12"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  bINum.SetCurrencySymbol('\U000000a3')

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_25(t *testing.T) {

  originalNumStr := "-1234567890.12"
  expectedNumStr := "-£1,234,567,890.12"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  bINum.SetCurrencySymbol('\U000000a3')

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_26(t *testing.T) {

  originalNumStr := "-1234567890.12"
  expectedNumStr := "(£1,234,567,890.12)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  bINum.SetCurrencySymbol('\U000000a3')

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_27(t *testing.T) {

  originalNumStr := "-1234567890.12"
  expectedNumStr := "($1,234,567,890.12)"
  mode := PARENTHESESNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_28(t *testing.T) {

  originalNumStr := "-1234567890.12"
  expectedNumStr := "-$1,234,567,890.12"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}

func TestBigIntNum_FormatCurrencyStr_29(t *testing.T) {

  originalNumStr := "1234567890.12"
  expectedNumStr := "$1,234,567,890.12"
  mode := LEADMINUSNEGVALFMTMODE

  bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(originalNumStr). "+
      " num1Str= '%v' Error='%v' ",
      originalNumStr, err.Error())
  }

  outStr := bINum.FormatCurrencyStr(mode)

  if expectedNumStr != outStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, outStr)
  }

}
