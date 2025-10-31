package mathops

import (
  "testing"
)

func TestIntAry_CompareSignedValues_01(t *testing.T) {

  ePrefix := "TestIntAry_AddIntToThis_01"

  originalNumberStr1 := "451.3"

  originalNumberStr2 := "451.2"

  expectedCompareResult := 1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_02(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_02"

  originalNumberStr1 := "45.13"

  originalNumberStr2 := "451.2"

  expectedCompareResult := -1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_03(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_03"

  originalNumberStr1 := "45.13975"

  originalNumberStr2 := "-451.21"

  expectedCompareResult := 1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_04(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_04"

  originalNumberStr1 := "0"

  originalNumberStr2 := "0.00"

  expectedCompareResult := 0

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_05(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_05"

  originalNumberStr1 := "-625.414"

  originalNumberStr2 := "-625.413"

  expectedCompareResult := -1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_06(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_06"

  originalNumberStr1 := "625.414"

  originalNumberStr2 := "625.413000"

  expectedCompareResult := 1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_07(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_07"

  originalNumberStr1 := "625.413"

  originalNumberStr2 := "625.413001"

  expectedCompareResult := -1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_08(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_08"

  originalNumberStr1 := "625.413"

  originalNumberStr2 := "625.413000"

  expectedCompareResult := 0

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_09(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_09"

  originalNumberStr1 := "625.413"

  originalNumberStr2 := "00625.413000"

  expectedCompareResult := 0

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_10(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_10"

  originalNumberStr1 := "625.413"

  originalNumberStr2 := "-00625.413000"

  expectedCompareResult := 1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_11(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_11"

  originalNumberStr1 := "-625.413"

  originalNumberStr2 := "-00625.413000"

  expectedCompareResult := 0

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareSignedValues_12(t *testing.T) {

  ePrefix := "TestIntAry_CompareSignedValues_12"

  originalNumberStr1 := "-625.413"

  originalNumberStr2 := "00625.413000"

  expectedCompareResult := -1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareSignedValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareSignedValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareAbsoluteValues_01(t *testing.T) {

  ePrefix := "TestIntAry_CompareAbsoluteValues_01"

  originalNumberStr1 := "45.13"

  originalNumberStr2 := "451.2"

  expectedCompareResult := -1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareAbsoluteValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareAbsoluteValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareAbsoluteValues_02(t *testing.T) {

  ePrefix := "TestIntAry_CompareAbsoluteValues_02"

  originalNumberStr1 := "45.13"

  originalNumberStr2 := "-451.2"

  expectedCompareResult := -1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareAbsoluteValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareAbsoluteValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareAbsoluteValues_03(t *testing.T) {

  ePrefix := "TestIntAry_CompareAbsoluteValues_03"

  originalNumberStr1 := "-45.13"

  originalNumberStr2 := "45.13"

  expectedCompareResult := 0

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareAbsoluteValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareAbsoluteValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareAbsoluteValues_04(t *testing.T) {

  ePrefix := "TestIntAry_CompareAbsoluteValues_04"

  originalNumberStr1 := "-45.13000"

  originalNumberStr2 := "45.13"

  expectedCompareResult := 0

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareAbsoluteValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareAbsoluteValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareAbsoluteValues_05(t *testing.T) {

  ePrefix := "TestIntAry_CompareAbsoluteValues_05"

  originalNumberStr1 := "-45.13001"

  originalNumberStr2 := "45.13"

  expectedCompareResult := 1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareAbsoluteValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareAbsoluteValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CompareAbsoluteValues_06(t *testing.T) {

  ePrefix := "TestIntAry_CompareAbsoluteValues_06"

  originalNumberStr1 := "-45.1300"

  originalNumberStr2 := "452"

  expectedCompareResult := -1

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  actualCompareResult, err := intAry1.CompareAbsoluteValues(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualCompareResult, err :=\n"+
      "  intAry1.CompareAbsoluteValues(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  if expectedCompareResult != actualCompareResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedCompareResult != actualCompareResult\n"+
      "Expected actualCompareResult = '%v'\n"+
      "  Actual actualCompareResult = '%v'\n\n",
      ePrefix, expectedCompareResult, actualCompareResult)

    return
  }

  return
}

func TestIntAry_CopyIn_01(t *testing.T) {

  ePrefix := "TestIntAry_CopyIn_01"

  originalNumberStr1 := "00100.1230"

  originalNumberStr2 := "-0000052.795813000"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = ia.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating initial intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating initial intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "intAry2 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  intAry2SignValue, err := intAry2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2SignValue, err := intAry2.GetSign()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry2NumberStr,
      err.Error())

    return
  }

  intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2PrecisionUint, err := intAry2.GetPrecisionUint()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry2NumberStr,
      err.Error())

    return
  }

  intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
      "intAry2= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(intAry2NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry2NumSeps \n"+
      "Expected intAry2NumSeps = '%v'\n"+
      "  Actual intAry2NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

    return
  }

  intAry2AryLength := intAry2.GetIntAryLength()

  err = intAry1.CopyIn(&intAry2, false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.CopyIn(&intAry2, false)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating final intAry1 value")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1 value')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1FinalNumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1FinalNumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1Final set to intAry1 after copy operation.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1FinalSignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1FinalSignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1FinalNumberStr,
      err.Error())

    return
  }

  intAry1FinalPrecisionUint, err := intAry1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1FinalPrecisionUint, err := intAry1.GetPrecisionUint()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry2NumberStr,
      err.Error())

    return
  }

  intAry1FinalNumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1FinalNumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1FinalNumberStr, err.Error())
    return
  }

  intAry1FinalAryLength := intAry1.GetIntAryLength()

  if intAry2NumberStr != intAry1FinalNumberStr {
    t.Errorf("%v\n"+
      "Error: IntAry Number String Values ARE NOT Equal\n"+
      "Because intAry2NumberStr != intAry1FinalNumberStr \n"+
      "Expected intAry1FinalNumberStr = '%v'\n"+
      "  Actual intAry1FinalNumberStr = '%v'\n\n",
      ePrefix, intAry2NumberStr, intAry1FinalNumberStr)

    return
  }

  if intAry2SignValue != intAry1FinalSignValue {
    t.Errorf("%v\n"+
      "Error: intAry Sign Values ARE NOT EQUAL!\n"+
      "Because intAry2SignValue != intAry1FinalSignValue\n"+
      "Expected intAry1FinalSignValue = '%v'\n"+
      "  Actual intAry1FinalSignValue = '%v'\n\n",
      ePrefix, intAry2SignValue, intAry1FinalSignValue)

    return
  }

  if intAry2PrecisionUint != intAry1FinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected Precision Values ARE NOT EQUAL!\n"+
      "Because intAry2PrecisionUint != intAry1FinalPrecisionUint\n"+
      "Expected intAry1FinalPrecisionUint = '%v'\n"+
      "  Actual intAry1FinalPrecisionUint = '%v'\n\n",
      ePrefix, intAry2PrecisionUint, intAry1FinalPrecisionUint)

    return
  }

  if intAry2AryLength != intAry1FinalAryLength {
    t.Errorf("%v\n"+
      "Error: Expected Array Lengths ARE NOT EQUAL!\n"+
      "Because intAry2AryLength != intAry1FinalAryLength\n"+
      "Expected intAry1FinalAryLength = '%v'\n"+
      "  Actual intAry1FinalAryLength = '%v'\n\n",
      ePrefix, intAry2AryLength, intAry1FinalAryLength)

    return
  }

  if !expectedNumSeps.Equal(intAry1FinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1FinalNumSeps \n"+
      "Expected intAry1FinalNumSeps = '%v'\n"+
      "  Actual intAry1FinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1FinalNumSeps.String())

    return
  }

  intAry2Stats := intAry2.GetIntAryStats()

  intAry1FinalStats := intAry1.GetIntAryStats()

  if intAry2Stats.IntegerLen != intAry1FinalStats.IntegerLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.IntegerLen != intAry1FinalStats.IntegerLen\n"+
      "Expected intAry1FinalStats.IntegerLen = '%v'\n"+
      "  Actual intAry1FinalStats.IntegerLen = '%v'\n\n",
      ePrefix, intAry2Stats.IntegerLen, intAry1FinalStats.IntegerLen)

    return
  }

  if intAry2Stats.SignificantIntegerLen != intAry1FinalStats.SignificantIntegerLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because  intAry2Stats.SignificantIntegerLen != intAry1FinalStats.SignificantIntegerLen\n"+
      "Expected intAry1FinalStats.SignificantIntegerLen = '%v'\n"+
      "  Actual intAry1FinalStats.SignificantIntegerLen = '%v'\n\n",
      ePrefix, intAry2Stats.SignificantIntegerLen, intAry1FinalStats.SignificantIntegerLen)

    return
  }

  if intAry2Stats.SignificantFractionLen != intAry1FinalStats.SignificantFractionLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.SignificantFractionLen != intAry1FinalStats.SignificantFractionLen\n"+
      "Expected intAry1FinalStats.SignificantFractionLen = '%v'\n"+
      "  Actual intAry1FinalStats.SignificantFractionLen = '%v'\n\n",
      ePrefix, intAry2Stats.SignificantFractionLen, intAry1FinalStats.SignificantFractionLen)

    return
  }

  if intAry2Stats.FirstDigitIdx != intAry1FinalStats.FirstDigitIdx {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.FirstDigitIdx != intAry1FinalStats.FirstDigitIdx\n"+
      "Expected intAry1FinalStats.FirstDigitIdx = '%v'\n"+
      "  Actual intAry1FinalStats.FirstDigitIdx = '%v'\n\n",
      ePrefix, intAry2Stats.FirstDigitIdx, intAry1FinalStats.FirstDigitIdx)

    return
  }

  if intAry2Stats.LastDigitIdx != intAry1FinalStats.LastDigitIdx {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.LastDigitIdx != intAry1FinalStats.LastDigitIdx\n"+
      "Expected intAry1FinalStats.LastDigitIdx = '%v'\n"+
      "  Actual intAry1FinalStats.LastDigitIdx = '%v'\n\n",
      ePrefix, intAry2Stats.LastDigitIdx, intAry1FinalStats.LastDigitIdx)

    return
  }

  if intAry2Stats.IsZeroValue != intAry1FinalStats.IsZeroValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.IsZeroValue != intAry1FinalStats.IsZeroValue\n"+
      "Expected intAry1FinalStats.IsZeroValue = '%v'\n"+
      "  Actual intAry1FinalStats.IsZeroValue = '%v'\n\n",
      ePrefix, intAry2Stats.IsZeroValue, intAry1FinalStats.IsZeroValue)

    return
  }

  if intAry2Stats.IsIntegerZeroValue != intAry1FinalStats.IsIntegerZeroValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.IsIntegerZeroValue != intAry1FinalStats.IsIntegerZeroValue\n"+
      "Expected intAry1FinalStats.IsIntegerZeroValue = '%v'\n"+
      "  Actual intAry1FinalStats.IsIntegerZeroValue = '%v'\n\n",
      ePrefix, intAry2Stats.IsIntegerZeroValue, intAry1FinalStats.IsIntegerZeroValue)

    return
  }

  var intAry2Element, intAry1FinalElement uint8

  for i := 0; i < intAry1.GetIntAryLength(); i++ {

    intAry2Element, err = intAry2.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "intAry2Element, err = intAry2.GetIntAryElement(%v)\n"+
        "intAry2= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAry2NumberStr,
        err.Error())

      return
    }

    intAry1FinalElement, err = intAry1.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "intAry1FinalElement, err = intAry1.GetIntAryElement(%v)\n"+
        "intAry1= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAry1FinalNumberStr,
        err.Error())

      return
    }

    if intAry2Element != intAry1FinalElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because intAry2Element != intAry1FinalElement\n"+
        "Expected intAry1FinalElement = '%v'\n"+
        "  Actual intAry1FinalElement = '%v'\n"+
        "Element Array Index (i) = %v\n\n",
        ePrefix, intAry2Element, intAry1FinalElement, i)

      return
    }

  }

  return
}

func TestIntAry_CopyIn_02(t *testing.T) {

  ePrefix := "TestIntAry_CopyIn_02"

  originalNumberStr1 := "00100.1230"

  originalNumberStr2 := "3594176.123459"

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  err = intAry1.CopyToBackUp()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.CopyToBackUp()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry2 := new(IntAry).New()

  err = intAry2.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = ia.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating initial intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating initial intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "intAry2 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  intAry2SignValue, err := intAry2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2SignValue, err := intAry2.GetSign()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry2NumberStr,
      err.Error())

    return
  }

  intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2PrecisionUint, err := intAry2.GetPrecisionUint()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry2NumberStr,
      err.Error())

    return
  }

  intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
      "intAry2= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(intAry2NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry2NumSeps \n"+
      "Expected intAry2NumSeps = '%v'\n"+
      "  Actual intAry2NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

    return
  }

  intAry2AryLength := intAry2.GetIntAryLength()

  err = intAry2.CopyToBackUp()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.CopyToBackUp()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  err = intAry1.CopyIn(&intAry2, true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.CopyIn(&intAry2, false)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating final intAry1 value")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1 value')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1FinalNumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1FinalNumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1Final set to intAry1 after copy operation.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1FinalSignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1FinalSignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1FinalNumberStr,
      err.Error())

    return
  }

  intAry1FinalPrecisionUint, err := intAry1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1FinalPrecisionUint, err := intAry1.GetPrecisionUint()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry2NumberStr,
      err.Error())

    return
  }

  intAry1FinalNumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1FinalNumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1FinalNumberStr, err.Error())
    return
  }

  intAry1FinalAryLength := intAry1.GetIntAryLength()

  if intAry2NumberStr != intAry1FinalNumberStr {
    t.Errorf("%v\n"+
      "Error: IntAry Number String Values ARE NOT Equal\n"+
      "Because intAry2NumberStr != intAry1FinalNumberStr \n"+
      "Expected intAry1FinalNumberStr = '%v'\n"+
      "  Actual intAry1FinalNumberStr = '%v'\n\n",
      ePrefix, intAry2NumberStr, intAry1FinalNumberStr)

    return
  }

  if intAry2SignValue != intAry1FinalSignValue {
    t.Errorf("%v\n"+
      "Error: intAry Sign Values ARE NOT EQUAL!\n"+
      "Because intAry2SignValue != intAry1FinalSignValue\n"+
      "Expected intAry1FinalSignValue = '%v'\n"+
      "  Actual intAry1FinalSignValue = '%v'\n\n",
      ePrefix, intAry2SignValue, intAry1FinalSignValue)

    return
  }

  if intAry2PrecisionUint != intAry1FinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected Precision Values ARE NOT EQUAL!\n"+
      "Because intAry2PrecisionUint != intAry1FinalPrecisionUint\n"+
      "Expected intAry1FinalPrecisionUint = '%v'\n"+
      "  Actual intAry1FinalPrecisionUint = '%v'\n\n",
      ePrefix, intAry2PrecisionUint, intAry1FinalPrecisionUint)

    return
  }

  if intAry2AryLength != intAry1FinalAryLength {
    t.Errorf("%v\n"+
      "Error: Expected Array Lengths ARE NOT EQUAL!\n"+
      "Because intAry2AryLength != intAry1FinalAryLength\n"+
      "Expected intAry1FinalAryLength = '%v'\n"+
      "  Actual intAry1FinalAryLength = '%v'\n\n",
      ePrefix, intAry2AryLength, intAry1FinalAryLength)

    return
  }

  if !expectedNumSeps.Equal(intAry1FinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1FinalNumSeps \n"+
      "Expected intAry1FinalNumSeps = '%v'\n"+
      "  Actual intAry1FinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1FinalNumSeps.String())

    return
  }

  intAry2Stats := intAry2.GetIntAryStats()

  intAry1FinalStats := intAry1.GetIntAryStats()

  if intAry2Stats.IntegerLen != intAry1FinalStats.IntegerLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.IntegerLen != intAry1FinalStats.IntegerLen\n"+
      "Expected intAry1FinalStats.IntegerLen = '%v'\n"+
      "  Actual intAry1FinalStats.IntegerLen = '%v'\n\n",
      ePrefix, intAry2Stats.IntegerLen, intAry1FinalStats.IntegerLen)

    return
  }

  if intAry2Stats.SignificantIntegerLen != intAry1FinalStats.SignificantIntegerLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because  intAry2Stats.SignificantIntegerLen != intAry1FinalStats.SignificantIntegerLen\n"+
      "Expected intAry1FinalStats.SignificantIntegerLen = '%v'\n"+
      "  Actual intAry1FinalStats.SignificantIntegerLen = '%v'\n\n",
      ePrefix, intAry2Stats.SignificantIntegerLen, intAry1FinalStats.SignificantIntegerLen)

    return
  }

  if intAry2Stats.SignificantFractionLen != intAry1FinalStats.SignificantFractionLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.SignificantFractionLen != intAry1FinalStats.SignificantFractionLen\n"+
      "Expected intAry1FinalStats.SignificantFractionLen = '%v'\n"+
      "  Actual intAry1FinalStats.SignificantFractionLen = '%v'\n\n",
      ePrefix, intAry2Stats.SignificantFractionLen, intAry1FinalStats.SignificantFractionLen)

    return
  }

  if intAry2Stats.FirstDigitIdx != intAry1FinalStats.FirstDigitIdx {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.FirstDigitIdx != intAry1FinalStats.FirstDigitIdx\n"+
      "Expected intAry1FinalStats.FirstDigitIdx = '%v'\n"+
      "  Actual intAry1FinalStats.FirstDigitIdx = '%v'\n\n",
      ePrefix, intAry2Stats.FirstDigitIdx, intAry1FinalStats.FirstDigitIdx)

    return
  }

  if intAry2Stats.LastDigitIdx != intAry1FinalStats.LastDigitIdx {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.LastDigitIdx != intAry1FinalStats.LastDigitIdx\n"+
      "Expected intAry1FinalStats.LastDigitIdx = '%v'\n"+
      "  Actual intAry1FinalStats.LastDigitIdx = '%v'\n\n",
      ePrefix, intAry2Stats.LastDigitIdx, intAry1FinalStats.LastDigitIdx)

    return
  }

  if intAry2Stats.IsZeroValue != intAry1FinalStats.IsZeroValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.IsZeroValue != intAry1FinalStats.IsZeroValue\n"+
      "Expected intAry1FinalStats.IsZeroValue = '%v'\n"+
      "  Actual intAry1FinalStats.IsZeroValue = '%v'\n\n",
      ePrefix, intAry2Stats.IsZeroValue, intAry1FinalStats.IsZeroValue)

    return
  }

  if intAry2Stats.IsIntegerZeroValue != intAry1FinalStats.IsIntegerZeroValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because intAry2Stats.IsIntegerZeroValue != intAry1FinalStats.IsIntegerZeroValue\n"+
      "Expected intAry1FinalStats.IsIntegerZeroValue = '%v'\n"+
      "  Actual intAry1FinalStats.IsIntegerZeroValue = '%v'\n\n",
      ePrefix, intAry2Stats.IsIntegerZeroValue, intAry1FinalStats.IsIntegerZeroValue)

    return
  }

  var intAry2Element, intAry1FinalElement uint8

  for i := 0; i < intAry1.GetIntAryLength(); i++ {

    intAry2Element, err = intAry2.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "intAry2Element, err = intAry2.GetIntAryElement(%v)\n"+
        "intAry2= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAry2NumberStr,
        err.Error())

      return
    }

    intAry1FinalElement, err = intAry1.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "intAry1FinalElement, err = intAry1.GetIntAryElement(%v)\n"+
        "intAry1= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAry1FinalNumberStr,
        err.Error())

      return
    }

    if intAry2Element != intAry1FinalElement {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because intAry2Element != intAry1FinalElement\n"+
        "Expected intAry1FinalElement = '%v'\n"+
        "  Actual intAry1FinalElement = '%v'\n"+
        "Element Array Index (i) = %v\n\n",
        ePrefix, intAry2Element, intAry1FinalElement, i)

      return
    }

  }

  if !intAry1.BackUp.Equals(&intAry2.BackUp) {
    t.Errorf("%v\n"+
      "Error: intAry1 & intAry2 Backups ARE NOT EQUAL!\n"+
      "Because !intAry1.BackUp.Equals(&intAry2.BackUp)\n\n",
      ePrefix)

    return
  }

  return
}

func TestIntAry_CopyToBackUp_01(t *testing.T) {

  ePrefix := "TestIntAry_CopyToBackUp_01"

  originalNumberStr1 := "100.123"

  originalNumberStr2 := "-52.795813"

  expectedNumberStr := "100.123"

  intAry1 := new(IntAry).New()

  err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAry1.IsValid("Validating initial intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating initial intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAry1NumberStr\n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAry1NumberStr)

    return
  }

  err = intAry1.CopyToBackUp()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.CopyToBackUp()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  err = intAry1.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Setting Final IntAry1 Value!"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry1.IsValid("Validating final intAry1 value")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1 value')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1FinalNumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1FinalNumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1Final set to intAry1 after copy operation.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry1FinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Number Strings ARE NOT EQUAL!\n"+
      "Because originalNumberStr2 != intAry1FinalNumberStr\n"+
      "Expected intAry1FinalNumberStr = '%v'\n"+
      "  Actual intAry1FinalNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry1FinalNumberStr)

    return
  }

  intAry1BackupNumberStr := intAry1.BackUp.GetNumStr()

  if expectedNumberStr != intAry1BackupNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != intAry1BackupNumberStr\n"+
      "Expected intAry1BackupNumberStr = '%v'\n"+
      "  Actual intAry1BackupNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1BackupNumberStr)

    return
  }

  return
}

func TestIntAry_DecrementIntegerOne_01(t *testing.T) {
  nStr1 := "100.123"
  expected := "-100.123"
  cycles := 200
  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)

  for i := 0; i < cycles; i++ {
    ia.DecrementIntegerOne()
  }

  if expected != ia.GetNumStr() {
    t.Errorf("Error - Expected numStrDto= '%v'. Instead, numStrDto= '%v'", expected, ia.GetNumStr())
  }
}

func TestIntAry_DecrementIntegerOne_02(t *testing.T) {
  nStr1 := "2000"
  expected := "-2000"
  cycles := 4000
  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)

  for i := 0; i < cycles; i++ {
    ia.DecrementIntegerOne()
  }

  if expected != ia.GetNumStr() {
    t.Errorf("Error - Expected numStrDto= '%v'. Instead, numStrDto= '%v'", expected, ia.GetNumStr())
  }

}

func TestIntAry_DecrementIntegerOne_03(t *testing.T) {
  nStr1 := "2000.123"
  expected := "-2000.123"
  cycles := 4000
  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)

  for i := 0; i < cycles; i++ {
    ia.DecrementIntegerOne()
  }

  if expected != ia.GetNumStr() {
    t.Errorf("Error - Expected numStrDto= '%v'. Instead, numStrDto= '%v'", expected, ia.GetNumStr())
  }

}

func TestIntAry_DivideByInt64_01(t *testing.T) {
  nStr1 := "579"
  expected := "17.545454545454545454545454545455"
  maxPrecision := 30
  divisor := int64(33)

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  err := ia.DivideByInt64(divisor, maxPrecision)

  if err != nil {
    t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
  }

  if expected != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
  }

  if ia.GetPrecision() != int(maxPrecision) {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'.", maxPrecision, ia.GetPrecision())
  }

}

func TestIntAry_DivideByInt64_02(t *testing.T) {
  nStr1 := "579"
  maxPrecision := 0
  divisor := int64(0)

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  err := ia.DivideByInt64(divisor, maxPrecision)

  if err == nil {
    t.Errorf("Expected Divide By zero Error. No Error Presented. Error = %v", err)
  }

}

func TestIntAry_DivideByInt64_03(t *testing.T) {
  nStr1 := "4"
  expected := "2"
  maxPrecision := 0
  divisor := int64(2)

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  err := ia.DivideByInt64(divisor, maxPrecision)

  if err != nil {
    t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
  }

  if expected != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
  }

  if ia.GetPrecision() != int(maxPrecision) {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'.", maxPrecision, ia.GetPrecision())
  }

}

func TestIntAry_DivideByInt64_04(t *testing.T) {
  nStr1 := "476"
  expected := "-14"
  maxPrecision := 10
  ePrecision := 0
  eSignVal := -1
  divisor := int64(-34)

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  err := ia.DivideByInt64(divisor, maxPrecision)

  if err != nil {
    t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
  }

  if expected != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
  }

  if ia.GetPrecision() != ePrecision {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'.", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v'. Instead, ia.GetPrecisionInt()= '%v'.", eSignVal, ia.GetSign())
  }

}

func TestIntAry_DivideByInt64_05(t *testing.T) {
  nStr1 := "-476"
  expected := "-14"
  maxPrecision := 10
  ePrecision := 0
  eSignVal := -1
  divisor := int64(34)

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  err := ia.DivideByInt64(divisor, maxPrecision)

  if err != nil {
    t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
  }

  if expected != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
  }

  if ia.GetPrecision() != ePrecision {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'.", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v'. Instead, ia.GetPrecisionInt()= '%v'.", eSignVal, ia.GetSign())
  }

}

func TestIntAry_DivideByInt64_06(t *testing.T) {
  nStr1 := "-476"
  expected := "14"
  maxPrecision := 10
  ePrecision := 0
  eSignVal := 1
  divisor := int64(-34)

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  err := ia.DivideByInt64(divisor, maxPrecision)

  if err != nil {
    t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
  }

  if expected != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
  }

  if ia.GetPrecision() != ePrecision {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'.", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v'. Instead, ia.GetPrecisionInt()= '%v'.", eSignVal, ia.GetSign())
  }

}

func TestIntAry_DivideByTwo_01(t *testing.T) {

  nStr1 := "1959"
  expected := "979.5"
  precision := 1

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  ia.DivideByTwo()

  if expected != ia.GetNumStr() {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", precision, ia.GetPrecision())
  }

}

func TestIntAry_DivideByTwo_02(t *testing.T) {

  nStr1 := "1"
  expected := "0.5"
  precision := 1

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  ia.DivideByTwo()

  if expected != ia.GetNumStr() {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", precision, ia.GetPrecision())
  }

}

func TestIntAry_DivideByTwo_03(t *testing.T) {

  nStr1 := "0"
  expected := "0"
  precision := 0

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  ia.DivideByTwo()

  if expected != ia.GetNumStr() {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", precision, ia.GetPrecision())
  }

}

func TestIntAry_DivideByTwo_04(t *testing.T) {

  nStr1 := "-2959"
  expected := "-1479.5"
  precision := 1
  signVal := -1

  ia := IntAry{}.New()
  ia.SetIntAryWithNumStr(nStr1)
  ia.DivideByTwo()

  if expected != ia.GetNumStr() {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", expected, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Error: Expected sign Value= '%v'. Instead, sign Value= '%v'. ", signVal, ia.GetSign())
  }

}

func TestIntAry_DivideByTenToPower_01(t *testing.T) {

  nStr := "457.3"
  power := uint(1)
  eNumStr := "45.73"
  eIAry := []uint8{4, 5, 7, 3}
  lEArray := len(eIAry)
  ePrecision := 2
  eSignVal := 1

  ia := IntAry{}.New()
  err := ia.SetIntAryWithNumStr(nStr)

  if err != nil {
    t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  ia.DivideByTenToPower(power)

  if ia.GetNumStr() != eNumStr {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
  }

  if ia.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
  }

  if lEArray != ia.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
  }

  for i := 0; i < lEArray; i++ {

    if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

      t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
      return

    }
  }
}

func TestIntAry_DivideByTenToPower_02(t *testing.T) {

  nStr := "457.3"
  power := uint(3)
  eNumStr := "0.4573"
  eIAry := []uint8{0, 4, 5, 7, 3}
  lEArray := len(eIAry)
  ePrecision := 4
  eSignVal := 1

  ia := IntAry{}.New()
  err := ia.SetIntAryWithNumStr(nStr)

  if err != nil {
    t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  ia.DivideByTenToPower(power)

  if ia.GetNumStr() != eNumStr {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
  }

  if ia.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
  }

  if lEArray != ia.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
  }

  for i := 0; i < lEArray; i++ {
    if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

      t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
      return

    }
  }
}

func TestIntAry_DivideByTenToPower_03(t *testing.T) {

  nStr := "457.3"
  power := uint(7)
  eNumStr := "0.00004573"
  eIAry := []uint8{0, 0, 0, 0, 0, 4, 5, 7, 3}
  lEArray := len(eIAry)
  ePrecision := 8
  eSignVal := 1

  ia := IntAry{}.New()
  err := ia.SetIntAryWithNumStr(nStr)

  if err != nil {
    t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  ia.DivideByTenToPower(power)

  if ia.GetNumStr() != eNumStr {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
  }

  if ia.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
  }

  if lEArray != ia.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
  }

  for i := 0; i < lEArray; i++ {
    if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

      t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
      return

    }
  }
}

func TestIntAry_DivideByTenToPower_04(t *testing.T) {

  nStr := "-457.3"
  power := uint(7)
  eNumStr := "-0.00004573"
  eIAry := []uint8{0, 0, 0, 0, 0, 4, 5, 7, 3}
  lEArray := len(eIAry)
  ePrecision := 8
  eSignVal := -1

  ia := IntAry{}.New()
  err := ia.SetIntAryWithNumStr(nStr)

  if err != nil {
    t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  ia.DivideByTenToPower(power)

  if ia.GetNumStr() != eNumStr {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
  }

  if ia.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
  }

  if lEArray != ia.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
  }

  for i := 0; i < lEArray; i++ {
    if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

      t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
      return

    }
  }
}

func TestIntAry_DivideByTenToPower_05(t *testing.T) {

  nStr := "0"
  power := uint(2)
  eNumStr := "0.00"
  eIAry := []uint8{0, 0, 0}
  lEArray := len(eIAry)
  ePrecision := 2
  eSignVal := 1

  ia := IntAry{}.New()
  err := ia.SetIntAryWithNumStr(nStr)

  if err != nil {
    t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  ia.DivideByTenToPower(power)

  if ia.GetNumStr() != eNumStr {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
  }

  if ia.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
  }

  if lEArray != ia.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
  }

  for i := 0; i < lEArray; i++ {
    if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

      t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
      return

    }
  }
}

func TestIntAry_DivideByTenToPower_06(t *testing.T) {

  nStr := "-4573"
  power := uint(1)
  eNumStr := "-457.3"
  eIAry := []uint8{4, 5, 7, 3}
  lEArray := len(eIAry)
  ePrecision := 1
  eSignVal := -1

  ia := IntAry{}.New()
  err := ia.SetIntAryWithNumStr(nStr)

  if err != nil {
    t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
  }

  ia.DivideByTenToPower(power)

  if ia.GetNumStr() != eNumStr {
    t.Errorf("Error: Expected numStrDto= '%v'. Instead received numStrDto= '%v'", eNumStr, ia.GetNumStr())
  }

  if ia.GetPrecision() != ePrecision {
    t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
  }

  if lEArray != ia.GetIntAryLength() {
    t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
  }

  for i := 0; i < lEArray; i++ {
    if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

      t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
      return

    }
  }
}

func TestIntAry_DivideThisBy_01(t *testing.T) {
  dividend := "56234369384300"
  divisor := "24"
  eQuotient := "2343098724345.833333333333333333333"
  eSignVal := 1
  maxPrecision := 21
  ePrecision := 21

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}

func TestIntAry_DivideThisBy_02(t *testing.T) {
  dividend := "48"
  divisor := "24"
  eSignVal := 1
  eQuotient := "2"
  maxPrecision := 21
  ePrecision := 0

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}

func TestIntAry_DivideThisBy_03(t *testing.T) {
  dividend := "24"
  divisor := "24"
  eQuotient := "1"
  eSignVal := 1
  maxPrecision := 21
  ePrecision := 0

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}

func TestIntAry_DivideThisBy_04(t *testing.T) {
  dividend := "0.05"
  divisor := "24"
  eQuotient := "0.00208333333333333333333333333333"
  eSignVal := 1
  maxPrecision := 32
  ePrecision := 32

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}

func TestIntAry_DivideThisBy_05(t *testing.T) {
  dividend := "0"
  divisor := "24"
  eQuotient := "0"
  eSignVal := 1
  maxPrecision := 7
  ePrecision := 0

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}

func TestIntAry_DivideThisBy_06(t *testing.T) {
  dividend := "48"
  divisor := "0"

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  _, err := ia1.DivideThisBy(&ia2, 0, 15)

  if err == nil {
    t.Error("Expected an error from Divideby zero. No Error Received!")
  }

}

func TestIntAry_DivideThisBy_07(t *testing.T) {
  dividend := "-9360"
  divisor := "24.48"
  eQuotient := "-382.35294117647058823529411764706"
  eSignVal := -1
  maxPrecision := 29
  ePrecision := 29

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }
}

func TestIntAry_DivideThisBy_08(t *testing.T) {
  dividend := "-9360"
  divisor := "-24.48"
  eQuotient := "382.35294117647058823529411764706"
  eSignVal := 1
  maxPrecision := 29
  ePrecision := 29

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }
}

func TestIntAry_DivideThisBy_09(t *testing.T) {
  dividend := "9360"
  divisor := "-24.48"
  eQuotient := "-382.35294117647058823529411764706"
  eSignVal := -1
  maxPrecision := 29
  ePrecision := 29

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }
}

func TestIntAry_DivideThisBy_10(t *testing.T) {
  dividend := "-19260.549"
  divisor := "-246.483"
  eQuotient := "78.141490488187826340964691276883"
  eSignVal := 1
  maxPrecision := 30
  ePrecision := 30

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }
}

func TestIntAry_DivideThisBy_11(t *testing.T) {
  // Testing condition where maxPrecision is
  // less than significant result.
  dividend := "1"
  divisor := "25"
  // actual result is "0.04"
  eQuotient := "0.0"
  eSignVal := 1
  maxPrecision := 1
  ePrecision := 1

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}

func TestIntAry_DivideThisBy_12(t *testing.T) {
  // Testing condition where maxPrecision is
  // greater than than significant result.
  dividend := "1"
  divisor := "25"
  // actual result is "0.04"
  eQuotient := "0.04"
  eSignVal := 1
  maxPrecision := 10 // maxPrecision exceeds actual result
  ePrecision := 2

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}

func TestIntAry_DivideThisBy_13(t *testing.T) {
  // Testing condition where maxPrecision is
  // set equal to -1.
  dividend := "1"
  divisor := "26"
  eQuotient := "0.0384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615385"
  eSignVal := 1
  maxPrecision := -1 // maxPrecision exceeds actual result
  ePrecision := 4096

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}

func TestIntAry_DivideThisBy_14(t *testing.T) {
  // Testing condition where maxPrecision is
  // invalid
  dividend := "1"
  divisor := "25"
  maxPrecision := -10 // maxPrecision exceeds actual result

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  _, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err == nil {
    t.Error("Expected error to result from invalid 'maxPrecision' value of -10. Instead, no such error was triggered")
  }

}

func TestIntAry_DivideThisBy_15(t *testing.T) {
  // Test Condition were actual decimal places are
  // 21+ and maxPrecision is set to zero
  dividend := "56234369384300"
  divisor := "24"
  // actual eQuotient := "2343098724345.833333333333333333333"
  eQuotient := "2343098724346"
  eSignVal := 1
  maxPrecision := 0
  ePrecision := 0

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}

func TestIntAry_DivideThisBy_16(t *testing.T) {
  // Test Condition were actual decimal places are
  // 21+ and maxPrecision is set to zero
  dividend := "1"
  divisor := "5132188731375616"
  eQuotient := "0.0000000000000001948486410654588479586365019367735312710807776562419879994613957690782179395268844976594976574190856910462423734806284591100904041759223999872534562498706764291924041486253101558649510819197172143928398136686955522152677536556303208645361801676397426605377561320278442021917586553676450097223493264453235937316338207393368423631288419620924881566999933402357520007410670136732435322811808257677794240393475046309703702321880635871858096536620946259448966209463978714360724504546027791919182097769678551824119271867445986093109132873712847086353082760161728835291168886244778340678734742887929927691392086547229915192975721487525948779793899832544746398668414935327086017026715079825557839286764777130808629545941416744923613139031065794055208370453857117267343145382178198213289314418222458170836304178717692361810139999277887517350389701068636718076181550016169308117583049414576551610275225085122626807472053468620691915924510100661007669795710586291378217910587138791895083422555361272265231787002422046566"
  eSignVal := 1
  maxPrecision := 1024
  ePrecision := 1024

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(dividend)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(divisor)

  quotient, err := ia1.DivideThisBy(&ia2, 0, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
  }

  if eQuotient != quotient.GetNumStr() {
    t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
  }

  if ePrecision != quotient.GetPrecision() {
    t.Errorf("Expected quotient.GetPrecisionInt()= '%v' .  Instead, quotient.GetPrecisionInt()= '%v'  .", ePrecision, quotient.GetPrecision())
  }

  if eSignVal != quotient.GetSign() {
    t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
  }

}
