package mathops

import (
  "testing"
)

func TestNumStrDto_CompareAbsoluteVals_01(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareAbsoluteVals_01"

  originalNumberStr1 := "-12567.218956"

  originalNumberStr2 := "-9211.40"

  expectedCompareInt := 1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareAbsoluteVals_02(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareAbsoluteVals_02"

  originalNumberStr1 := "-12567.218956"

  originalNumberStr2 := "9211.40"

  expectedCompareInt := 1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareAbsoluteVals_03(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareAbsoluteVals_03"

  originalNumberStr1 := "-12567.218956"

  originalNumberStr2 := "12567.218956"

  expectedCompareInt := 0

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareAbsoluteVals_04(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareAbsoluteVals_04"

  originalNumberStr1 := "567.21"

  originalNumberStr2 := "12567.218956"

  expectedCompareInt := -1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareAbsoluteVals_05(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareAbsoluteVals_05"

  originalNumberStr1 := "567.21"

  originalNumberStr2 := "-12567.218956"

  expectedCompareInt := -1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareAbsoluteVals_06(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareAbsoluteVals_06"

  originalNumberStr1 := "567.21"

  originalNumberStr2 := "-567.21"

  expectedCompareInt := 0

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareAbsoluteVals_07(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareAbsoluteVals_07"

  originalNumberStr1 := "567.21"

  originalNumberStr2 := "567.21"

  expectedCompareInt := 0

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareAbsoluteVals_08(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareAbsoluteVals_08"

  originalNumberStr1 := "567.21"

  originalNumberStr2 := "1567.21"

  expectedCompareInt := -1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareSignedVals_01(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareSignedVals_01"

  originalNumberStr1 := "-12567.218956"

  originalNumberStr2 := "-9211.40"

  expectedCompareInt := -1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareSignedValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if actualCompareInt != expectedCompareInt {
    t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareSignedVals_02(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareSignedVals_02"

  originalNumberStr1 := "12567.218956"

  originalNumberStr2 := "9211.40"

  expectedCompareInt := 1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareSignedValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if actualCompareInt != expectedCompareInt {
    t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareSignedVals_03(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareSignedVals_03"

  originalNumberStr1 := "-12567.218956"

  originalNumberStr2 := "9211.40"

  expectedCompareInt := -1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareSignedValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if actualCompareInt != expectedCompareInt {
    t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareSignedVals_04(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareSignedVals_04"

  originalNumberStr1 := "12567.218956"

  originalNumberStr2 := "-9211.40"

  expectedCompareInt := 1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareSignedValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if actualCompareInt != expectedCompareInt {
    t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareSignedVals_05(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareSignedVals_05"

  originalNumberStr1 := "12567.218956"

  originalNumberStr2 := "-12567.218956"

  expectedCompareInt := 1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareSignedValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if actualCompareInt != expectedCompareInt {
    t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareSignedVals_06(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareSignedVals_06"

  originalNumberStr1 := "-12567.218956"

  originalNumberStr2 := "12567.218956"

  expectedCompareInt := -1

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareSignedValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if actualCompareInt != expectedCompareInt {
    t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareSignedVals_07(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareSignedVals_07"

  originalNumberStr1 := "-12567.218956"

  originalNumberStr2 := "-12567.218956"

  expectedCompareInt := 0

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareSignedValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if actualCompareInt != expectedCompareInt {
    t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CompareSignedVals_08(t *testing.T) {

  ePrefix := "TestNumStrDto_CompareSignedVals_08"

  originalNumberStr1 := "12567.218956"

  originalNumberStr2 := "12567.218956"

  expectedCompareInt := 0

  nDto := new(NumStrDto).New()

  numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = numStrDto1.IsValid("Validating numStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto1.IsValid('Validating numStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto1NumberStr, err := numStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
      "numStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != numStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != numStrDto1NumberStr\n"+
      "Expected numStrDto1NumberStr = '%v'\n"+
      "  Actual numStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, numStrDto1NumberStr)

    return
  }

  numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
      "  originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = numStrDto2.IsValid("Validating numStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDto2.IsValid('Validating numStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDto2NumberStr, err := numStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
      "numStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != numStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: numStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != numStrDto2NumberStr\n"+
      "Expected numStrDto2NumberStr = '%v'\n"+
      "  Actual numStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, numStrDto2NumberStr)

    return
  }

  actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareInt, err := nDto.CompareSignedValues(\n"+
      "  &numStrDto1, &numStrDto2)\n"+
      "numStrDto1= '%v'\n"+
      "numStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDto1NumberStr,
      numStrDto2NumberStr,
      err.Error())

    return
  }

  if actualCompareInt != expectedCompareInt {
    t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
  }

  if expectedCompareInt != actualCompareInt {
    t.Errorf("%v\n"+
      "Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
      "Because expectedCompareInt != actualCompareInt\n"+
      "Expected actualCompareInt = '%v'\n"+
      "  Actual actualCompareInt = '%v'\n\n",
      ePrefix, expectedCompareInt, actualCompareInt)

    return
  }

  return
}

func TestNumStrDto_CopyIn_01(t *testing.T) {

  nStr := "123.456"
  iStr := "123"
  fracStr := "456"
  signVal := 1
  precision := uint(3)

  n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

  if err != nil {
    t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  nDto := NumStrDto{}.New()

  nDto.CopyIn(n1)

  s := nDto.GetNumStr()

  if s != nStr {
    t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
  }

  s = string(nDto.GetAbsIntRunes())

  if iStr != s {
    t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

  }

  s = string(nDto.GetAbsFracRunes())

  if fracStr != s {
    t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
  }

  if nDto.GetSign() != signVal {
    t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
  }

  if !nDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
  }

  if !nDto.IsFractionalValue() {
    t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
  }

  if precision != nDto.GetPrecisionUint() {
    t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecisionUint())

  }

  err = nDto.IsValid("Test 'nDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
  }

}

func TestNumStrDto_CopyIn_02(t *testing.T) {

  nStr := "-123.456"
  iStr := "123"
  fracStr := "456"
  signVal := -1
  precision := uint(3)

  n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

  if err != nil {
    t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  nDto := NumStrDto{}.New()

  nDto.CopyIn(n1)

  s := nDto.GetNumStr()

  if s != nStr {
    t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
  }

  s = string(nDto.GetAbsIntRunes())

  if iStr != s {
    t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

  }

  s = string(nDto.GetAbsFracRunes())

  if fracStr != s {
    t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
  }

  if nDto.GetSign() != signVal {
    t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
  }

  if !nDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
  }

  if !nDto.IsFractionalValue() {
    t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
  }

  if precision != nDto.GetPrecisionUint() {
    t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecisionUint())

  }

  err = nDto.IsValid("Test 'nDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
  }

}

func TestNumStrDto_CopyOut_01(t *testing.T) {

  nStr := "123.456"
  iStr := "123"
  fracStr := "456"
  signVal := 1
  precision := uint(3)

  n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

  if err != nil {
    t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  nDto := n1.CopyOut()

  s := nDto.GetNumStr()

  if s != nStr {
    t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
  }

  s = string(nDto.GetAbsIntRunes())

  if iStr != s {
    t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

  }

  s = string(nDto.GetAbsFracRunes())

  if fracStr != s {
    t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
  }

  if nDto.GetSign() != signVal {
    t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
  }

  if !nDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
  }

  if !nDto.IsFractionalValue() {
    t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
  }

  if precision != nDto.GetPrecisionUint() {
    t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecisionUint())

  }

  err = nDto.IsValid("Test 'nDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
  }

}

func TestNumStrDto_CopyOut_02(t *testing.T) {

  nStr := "-123.456"
  iStr := "123"
  fracStr := "456"
  signVal := -1
  precision := uint(3)

  n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

  if err != nil {
    t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  nDto := n1.CopyOut()

  s := nDto.GetNumStr()

  if s != nStr {
    t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
  }

  s = string(nDto.GetAbsIntRunes())

  if iStr != s {
    t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

  }

  s = string(nDto.GetAbsFracRunes())

  if fracStr != s {
    t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
  }

  if nDto.GetSign() != signVal {
    t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
  }

  if !nDto.HasNumericDigits() {
    t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
  }

  if !nDto.IsFractionalValue() {
    t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
  }

  if precision != nDto.GetPrecisionUint() {
    t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecisionUint())

  }

  err = nDto.IsValid("Test 'nDto' is INVALID! ")

  if err != nil {
    t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
  }

}

func TestNumStrDto_Divide_01(t *testing.T) {

  nStr1 := "12"
  nStr2 := "3"
  expectedStr := "4.0"

  n1Dto, err := NumStrDto{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(nStr1). "+
      "nStr1='%v' Error= %v",
      nStr1, err.Error())
  }

  n2Dto, err := NumStrDto{}.NewNumStr(nStr2)

  if err != nil {
    t.Errorf("Error returned from NumStrDto{}.NewNumStr(nStr2). "+
      "nStr2='%v' Error= %v",
      nStr2, err.Error())
  }

  err = n1Dto.Divide(n2Dto, 1, 5)

  if err != nil {
    t.Errorf("Error returned from n1Dto.Divide(n2Dto, 1, 5). "+
      " Error= %v", err.Error())
  }

  numStr := n1Dto.GetNumStr()

  if expectedStr != numStr {
    t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
      expectedStr, numStr)
  }

}

func TestNumStrDto_FormatForMathOps_01(t *testing.T) {
  nStr1 := "-12567.218956"
  nStr2 := "-9211.40"
  nStr3 := "-12567.218956"
  nStr4 := "-09211.400000"
  expectedCompare := 1

  nDto := NumStrDto{}.New()

  n1, _ := nDto.ParseNumStr(nStr1)
  n2, _ := nDto.ParseNumStr(nStr2)
  nOut1, _ := nDto.ParseNumStr(nStr3)
  nOut2, _ := nDto.ParseNumStr(nStr4)

  n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

  if err != nil {
    t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
  }

  if false != isReversed {
    t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
  }

  if compare != expectedCompare {
    t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
  }

  if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
    t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
  }

  if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
    t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
  }

  if nOut1.GetSign() != n1OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
  }

  if nOut2.GetSign() != n2OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
  }

  if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
    t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
  }

  if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
    t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
  }

}

func TestNumStrDto_FormatForMathOps_02(t *testing.T) {
  nStr1 := "-9211.40"
  nStr2 := "-12567.218956"
  nStr3 := "-12567.218956"
  nStr4 := "-09211.400000"
  expectedCompare := 1

  nDto := NumStrDto{}.New()

  n1, _ := nDto.ParseNumStr(nStr1)
  n2, _ := nDto.ParseNumStr(nStr2)
  nOut1, _ := nDto.ParseNumStr(nStr3)
  nOut2, _ := nDto.ParseNumStr(nStr4)

  n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

  if err != nil {
    t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
  }

  if true != isReversed {
    t.Errorf("Expected isReverse = '%v'. Instead got '%v'", true, isReversed)
  }

  if compare != expectedCompare {
    t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
  }

  if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
    t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
  }

  if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
    t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
  }

  if nOut1.GetSign() != n1OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
  }

  if nOut2.GetSign() != n2OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
  }

  if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
    t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
  }

  if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
    t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
  }

}

func TestNumStrDto_FormatForMathOps_03(t *testing.T) {
  nStr1 := "-6"
  nStr2 := "67.521"
  nStr3 := "67.521"
  nStr4 := "-06.000"
  expectedCompare := 1

  nDto := NumStrDto{}.New()

  n1, _ := nDto.ParseNumStr(nStr1)
  n2, _ := nDto.ParseNumStr(nStr2)
  nOut1, _ := nDto.ParseNumStr(nStr3)
  nOut2, _ := nDto.ParseNumStr(nStr4)

  n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

  if err != nil {
    t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
  }

  if true != isReversed {
    t.Errorf("Expected isReverse = '%v'. Instead got '%v'", true, isReversed)
  }

  if compare != expectedCompare {
    t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
  }

  if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
    t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
  }

  if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
    t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
  }

  if nOut1.GetSign() != n1OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
  }

  if nOut2.GetSign() != n2OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
  }

  if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
    t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
  }

  if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
    t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
  }

}

func TestNumStrDto_FormatForMathOps_04(t *testing.T) {
  nStr1 := "67.521"
  nStr2 := "-6"
  nStr3 := "67.521"
  nStr4 := "-06.000"
  expectedCompare := 1

  nDto := NumStrDto{}.New()

  n1, _ := nDto.ParseNumStr(nStr1)
  n2, _ := nDto.ParseNumStr(nStr2)
  nOut1, _ := nDto.ParseNumStr(nStr3)
  nOut2, _ := nDto.ParseNumStr(nStr4)

  n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

  if err != nil {
    t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
  }

  if false != isReversed {
    t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
  }

  if compare != expectedCompare {
    t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
  }

  if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
    t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
  }

  if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
    t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
  }

  if nOut1.GetSign() != n1OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
  }

  if nOut2.GetSign() != n2OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
  }

  if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
    t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
  }

  if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
    t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
  }

}

func TestNumStrDto_FormatForMathOps_05(t *testing.T) {
  nStr1 := "-67.521"
  nStr2 := "6"
  nStr3 := "-67.521"
  nStr4 := "06.000"
  expectedCompare := 1
  nDto := NumStrDto{}.New()

  n1, _ := nDto.ParseNumStr(nStr1)
  n2, _ := nDto.ParseNumStr(nStr2)
  nOut1, _ := nDto.ParseNumStr(nStr3)
  nOut2, _ := nDto.ParseNumStr(nStr4)

  n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

  if err != nil {
    t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
  }

  if false != isReversed {
    t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
  }

  if compare != expectedCompare {
    t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
  }

  if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
    t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
  }

  if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
    t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
  }

  if nOut1.GetSign() != n1OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
  }

  if nOut2.GetSign() != n2OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
  }

  if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
    t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
  }

  if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
    t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
  }

}

func TestNumStrDto_FormatForMathOps_06(t *testing.T) {
  nStr1 := "-67.521"
  nStr2 := "67.521"
  nStr3 := "-67.521"
  nStr4 := "67.521"
  expectedCompare := 0
  nDto := NumStrDto{}.New()

  n1, _ := nDto.ParseNumStr(nStr1)
  n2, _ := nDto.ParseNumStr(nStr2)
  nOut1, _ := nDto.ParseNumStr(nStr3)
  nOut2, _ := nDto.ParseNumStr(nStr4)

  n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

  if err != nil {
    t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
  }

  if false != isReversed {
    t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
  }

  if compare != expectedCompare {
    t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
  }

  if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
    t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
  }

  if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
    t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
  }

  if nOut1.GetSign() != n1OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
  }

  if nOut2.GetSign() != n2OutDto.GetSign() {
    t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
  }

  if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
    t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
  }

  if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
    t.Errorf("Expected n2OutDto.GetPrecisionInt()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
  }

}
