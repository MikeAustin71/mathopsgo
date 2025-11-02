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

  ePrefix := "TestIntAry_DecrementIntegerOne_01"

  originalNumberStr := "100.123"

  expectedNumberStr := "-100.123"

  expectedPrecisionUint := uint(3)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  cycles := 200

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  for i := 0; i < cycles; i++ {

    err = intAry.DecrementIntegerOne()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = intAry.DecrementIntegerOne()\n"+
        "Index Cycle = '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DecrementIntegerOne_02(t *testing.T) {

  ePrefix := "TestIntAry_DecrementIntegerOne_02"

  originalNumberStr := "2000"

  expectedNumberStr := "-2000"

  expectedPrecisionUint := uint(0)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  cycles := 4000

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  for i := 0; i < cycles; i++ {

    err = intAry.DecrementIntegerOne()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = intAry.DecrementIntegerOne()\n"+
        "Index Cycle = '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DecrementIntegerOne_03(t *testing.T) {

  ePrefix := "TestIntAry_DecrementIntegerOne_03"

  originalNumberStr := "2000.123"

  expectedNumberStr := "-2000.123"

  expectedPrecisionUint := uint(3)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  cycles := 4000

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  for i := 0; i < cycles; i++ {

    err = intAry.DecrementIntegerOne()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = intAry.DecrementIntegerOne()\n"+
        "Index Cycle = '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByInt64_01(t *testing.T) {

  ePrefix := "TestIntAry_DivideByInt64_01"

  originalNumberStr := "579"

  maxPrecision := 30

  divisorInt64 := int64(33)

  //                                1         2         3
  //                     0.123456789012345678901234567890
  expectedNumberStr := "17.545454545454545454545454545455"

  expectedPrecisionUint := uint(30)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByInt64(divisorInt64, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByInt64(divisorInt64, maxPrecision)\n"+
      "intAry= '%v'\n"+
      "divisorInt64= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      divisorInt64,
      maxPrecision,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByInt64_02(t *testing.T) {

  ePrefix := "TestIntAry_DivideByInt64_02"

  originalNumberStr := "579"

  maxPrecision := 0

  divisorInt64 := int64(0)

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByInt64(divisorInt64, maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Function Call:\n"+
      "   err = intAry.DivideByInt64(divisorInt64, maxPrecision)\n"+
      "Expected a Divide By Zero Error, BUT NO ERROR WAS RETURNED!\n"+
      "divisorInt64=0\n"+
      "Divide by zero should produce an error.\n\n", ePrefix)
    return
  }

  return
}

func TestIntAry_DivideByInt64_03(t *testing.T) {

  ePrefix := "TestIntAry_DivideByInt64_03"

  originalNumberStr := "4"

  maxPrecision := 0

  divisorInt64 := int64(2)

  //                                1         2         3
  //                     0.123456789012345678901234567890
  expectedNumberStr := "2"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByInt64(divisorInt64, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByInt64(divisorInt64, maxPrecision)\n"+
      "intAry= '%v'\n"+
      "divisorInt64= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      divisorInt64,
      maxPrecision,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByInt64_04(t *testing.T) {

  ePrefix := "TestIntAry_DivideByInt64_04"

  originalNumberStr := "476"

  maxPrecision := 10

  divisorInt64 := int64(-34)

  //                                 1         2         3
  //                      0.123456789012345678901234567890
  expectedNumberStr := "-14"

  expectedPrecisionUint := uint(0)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByInt64(divisorInt64, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByInt64(divisorInt64, maxPrecision)\n"+
      "intAry= '%v'\n"+
      "divisorInt64= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      divisorInt64,
      maxPrecision,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByInt64_05(t *testing.T) {

  ePrefix := "TestIntAry_DivideByInt64_05"

  originalNumberStr := "-476"

  maxPrecision := 10

  divisorInt64 := int64(34)

  expectedNumberStr := "-14"

  expectedPrecisionUint := uint(0)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByInt64(divisorInt64, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByInt64(divisorInt64, maxPrecision)\n"+
      "intAry= '%v'\n"+
      "divisorInt64= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      divisorInt64,
      maxPrecision,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByInt64_06(t *testing.T) {

  ePrefix := "TestIntAry_DivideByInt64_06"

  originalNumberStr := "-476"

  maxPrecision := 10

  divisorInt64 := int64(-34)

  expectedNumberStr := "14"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByInt64(divisorInt64, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByInt64(divisorInt64, maxPrecision)\n"+
      "intAry= '%v'\n"+
      "divisorInt64= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      divisorInt64,
      maxPrecision,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByTwo_01(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTwo_01"

  originalNumberStr := "1959"

  //                                 1         2         3
  //                      0.123456789012345678901234567890
  expectedNumberStr := "979.5"

  expectedPrecisionUint := uint(1)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTwo()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTwo()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByTwo_02(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTwo_02"

  originalNumberStr := "1"

  //                               1         2         3
  //                    0.123456789012345678901234567890
  expectedNumberStr := "0.5"

  expectedPrecisionUint := uint(1)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTwo()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTwo()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByTwo_03(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTwo_03"

  originalNumberStr := "0"

  //                               1         2         3
  //                    0.123456789012345678901234567890
  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTwo()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTwo()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByTwo_04(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTwo_04"

  originalNumberStr := "-2959"

  //                                   1         2         3
  //                        0.123456789012345678901234567890
  expectedNumberStr := "-1479.5"

  expectedPrecisionUint := uint(1)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTwo()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTwo()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideByTenToPower_01(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTenToPower_01"

  originalNumberStr := "457.3"

  power := uint(1)

  //                                1         2         3
  //                     0.123456789012345678901234567890
  expectedNumberStr := "45.73"

  expectedPrecisionUint := uint(2)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntAry := []uint8{4, 5, 7, 3}

  lenExpectedIntArray := len(expectedIntAry)

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTenToPower(power)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTenToPower(power)\n"+
      "intAry= '%v'\n"+
      "power= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      power,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intAryArrayLen := intAry.GetIntAryLength()

  if lenExpectedIntArray != intAryArrayLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedIntArray != intAryArrayLen\n"+
      "Expected intAryArrayLen = '%v'\n"+
      "  Actual intAryArrayLen = '%v'\n\n",
      ePrefix, lenExpectedIntArray, intAryArrayLen)

    return
  }

  var actualElement uint8

  for i := 0; i < lenExpectedIntArray; i++ {

    actualElement, err = intAry.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualElement, err = intAry.GetIntAryElement(%v)\n"+
        "intAry= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAryNumberStr,
        err.Error())

      return
    }

    if expectedIntAry[i] != actualElement {
      t.Errorf("%v\n"+
        "Error: Expected and Actual IntAry Elements ARE NOT EQUAL!\n"+
        "Because expectedIntAry[%v] != actualElement\n"+
        "Expected actualElement = '%v'\n"+
        "  Actual actualElement = '%v'\n\n",
        ePrefix, i, expectedIntAry[i], actualElement)

      return
    }
  }

  return
}

func TestIntAry_DivideByTenToPower_02(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTenToPower_02"

  originalNumberStr := "457.3"

  power := uint(3)

  //                               1         2         3
  //                    0.123456789012345678901234567890
  expectedNumberStr := "0.4573"

  expectedPrecisionUint := uint(4)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntAry := []uint8{0, 4, 5, 7, 3}

  lenExpectedIntArray := len(expectedIntAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTenToPower(power)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTenToPower(power)\n"+
      "intAry= '%v'\n"+
      "power= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      power,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intAryArrayLen := intAry.GetIntAryLength()

  if lenExpectedIntArray != intAryArrayLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedIntArray != intAryArrayLen\n"+
      "Expected intAryArrayLen = '%v'\n"+
      "  Actual intAryArrayLen = '%v'\n\n",
      ePrefix, lenExpectedIntArray, intAryArrayLen)

    return
  }

  var actualElement uint8

  for i := 0; i < lenExpectedIntArray; i++ {

    actualElement, err = intAry.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualElement, err = intAry.GetIntAryElement(%v)\n"+
        "intAry= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAryNumberStr,
        err.Error())

      return
    }

    if expectedIntAry[i] != actualElement {
      t.Errorf("%v\n"+
        "Error: Expected and Actual IntAry Elements ARE NOT EQUAL!\n"+
        "Because expectedIntAry[%v] != actualElement\n"+
        "Expected actualElement = '%v'\n"+
        "  Actual actualElement = '%v'\n\n",
        ePrefix, i, expectedIntAry[i], actualElement)

      return
    }
  }

  return
}

func TestIntAry_DivideByTenToPower_03(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTenToPower_03"

  originalNumberStr := "457.3"

  power := uint(7)

  //                               1         2         3
  //                    0.123456789012345678901234567890
  expectedNumberStr := "0.00004573"

  expectedPrecisionUint := uint(8)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntAry := []uint8{0, 0, 0, 0, 0, 4, 5, 7, 3}

  lenExpectedIntArray := len(expectedIntAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTenToPower(power)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTenToPower(power)\n"+
      "intAry= '%v'\n"+
      "power= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      power,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intAryArrayLen := intAry.GetIntAryLength()

  if lenExpectedIntArray != intAryArrayLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedIntArray != intAryArrayLen\n"+
      "Expected intAryArrayLen = '%v'\n"+
      "  Actual intAryArrayLen = '%v'\n\n",
      ePrefix, lenExpectedIntArray, intAryArrayLen)

    return
  }

  var actualElement uint8

  for i := 0; i < lenExpectedIntArray; i++ {

    actualElement, err = intAry.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualElement, err = intAry.GetIntAryElement(%v)\n"+
        "intAry= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAryNumberStr,
        err.Error())

      return
    }

    if expectedIntAry[i] != actualElement {
      t.Errorf("%v\n"+
        "Error: Expected and Actual IntAry Elements ARE NOT EQUAL!\n"+
        "Because expectedIntAry[%v] != actualElement\n"+
        "Expected actualElement = '%v'\n"+
        "  Actual actualElement = '%v'\n\n",
        ePrefix, i, expectedIntAry[i], actualElement)

      return
    }
  }

  return
}

func TestIntAry_DivideByTenToPower_04(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTenToPower_04"

  originalNumberStr := "-457.3"

  power := uint(7)

  //                                1         2         3
  //                     0.123456789012345678901234567890
  expectedNumberStr := "-0.00004573"

  expectedPrecisionUint := uint(8)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntAry := []uint8{0, 0, 0, 0, 0, 4, 5, 7, 3}

  lenExpectedIntArray := len(expectedIntAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTenToPower(power)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTenToPower(power)\n"+
      "intAry= '%v'\n"+
      "power= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      power,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intAryArrayLen := intAry.GetIntAryLength()

  if lenExpectedIntArray != intAryArrayLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedIntArray != intAryArrayLen\n"+
      "Expected intAryArrayLen = '%v'\n"+
      "  Actual intAryArrayLen = '%v'\n\n",
      ePrefix, lenExpectedIntArray, intAryArrayLen)

    return
  }

  var actualElement uint8

  for i := 0; i < lenExpectedIntArray; i++ {

    actualElement, err = intAry.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualElement, err = intAry.GetIntAryElement(%v)\n"+
        "intAry= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAryNumberStr,
        err.Error())

      return
    }

    if expectedIntAry[i] != actualElement {
      t.Errorf("%v\n"+
        "Error: Expected and Actual IntAry Elements ARE NOT EQUAL!\n"+
        "Because expectedIntAry[%v] != actualElement\n"+
        "Expected actualElement = '%v'\n"+
        "  Actual actualElement = '%v'\n\n",
        ePrefix, i, expectedIntAry[i], actualElement)

      return
    }
  }

  return
}

func TestIntAry_DivideByTenToPower_05(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTenToPower_05"

  originalNumberStr := "0"

  power := uint(2)

  //                               1         2         3
  //                    0.123456789012345678901234567890
  expectedNumberStr := "0.00"

  expectedPrecisionUint := uint(2)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntAry := []uint8{0, 0, 0}

  lenExpectedIntArray := len(expectedIntAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTenToPower(power)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTenToPower(power)\n"+
      "intAry= '%v'\n"+
      "power= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      power,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intAryArrayLen := intAry.GetIntAryLength()

  if lenExpectedIntArray != intAryArrayLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedIntArray != intAryArrayLen\n"+
      "Expected intAryArrayLen = '%v'\n"+
      "  Actual intAryArrayLen = '%v'\n\n",
      ePrefix, lenExpectedIntArray, intAryArrayLen)

    return
  }

  var actualElement uint8

  for i := 0; i < lenExpectedIntArray; i++ {

    actualElement, err = intAry.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualElement, err = intAry.GetIntAryElement(%v)\n"+
        "intAry= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAryNumberStr,
        err.Error())

      return
    }

    if expectedIntAry[i] != actualElement {
      t.Errorf("%v\n"+
        "Error: Expected and Actual IntAry Elements ARE NOT EQUAL!\n"+
        "Because expectedIntAry[%v] != actualElement\n"+
        "Expected actualElement = '%v'\n"+
        "  Actual actualElement = '%v'\n\n",
        ePrefix, i, expectedIntAry[i], actualElement)

      return
    }
  }

  return
}

func TestIntAry_DivideByTenToPower_06(t *testing.T) {

  ePrefix := "TestIntAry_DivideByTenToPower_06"

  originalNumberStr := "-4573"

  power := uint(1)

  //                                  1         2         3
  //                       0.123456789012345678901234567890
  expectedNumberStr := "-457.3"

  expectedPrecisionUint := uint(1)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntAry := []uint8{4, 5, 7, 3}

  lenExpectedIntArray := len(expectedIntAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.DivideByTenToPower(power)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.DivideByTenToPower(power)\n"+
      "intAry= '%v'\n"+
      "power= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      power,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validating Final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAryNumberStr set to final IntAry value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intAryArrayLen := intAry.GetIntAryLength()

  if lenExpectedIntArray != intAryArrayLen {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedIntArray != intAryArrayLen\n"+
      "Expected intAryArrayLen = '%v'\n"+
      "  Actual intAryArrayLen = '%v'\n\n",
      ePrefix, lenExpectedIntArray, intAryArrayLen)

    return
  }

  var actualElement uint8

  for i := 0; i < lenExpectedIntArray; i++ {

    actualElement, err = intAry.GetIntAryElement(i)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualElement, err = intAry.GetIntAryElement(%v)\n"+
        "intAry= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        i,
        intAryNumberStr,
        err.Error())

      return
    }

    if expectedIntAry[i] != actualElement {
      t.Errorf("%v\n"+
        "Error: Expected and Actual IntAry Elements ARE NOT EQUAL!\n"+
        "Because expectedIntAry[%v] != actualElement\n"+
        "Expected actualElement = '%v'\n"+
        "  Actual actualElement = '%v'\n\n",
        ePrefix, i, expectedIntAry[i], actualElement)

      return
    }
  }

  return
}

func TestIntAry_DivideThisBy_01(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_01"

  originalDividendNumberStr := "56234369384300"

  originalDivisorNumberStr := "24"

  maxPrecision := 21

  //                                                   1         2         3
  //                                        0.123456789012345678901234567890
  expectedQuotientNumberStr := "2343098724345.833333333333333333333"

  expectedPrecisionUint := uint(21)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend := new(IntAry).New()

  err := intAryDividend.SetIntAryWithNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryDividend.SetIntAryWithNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  err = intAryDividend.IsValid("Validating initial intAryDividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDividend.IsValid('Validating initial intAryDividend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor := new(IntAry).New()

  err = intAryDivisor.SetIntAryWithNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryDivisor.SetIntAryWithNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_02(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_02"

  originalDividendNumberStr := "48"

  originalDivisorNumberStr := "24"

  maxPrecision := 21

  //                                       1         2         3
  //                            0.123456789012345678901234567890
  expectedQuotientNumberStr := "2"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_03(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_03"

  originalDividendNumberStr := "24"

  originalDivisorNumberStr := "24"

  maxPrecision := 21

  //                                       1         2         3
  //                            0.123456789012345678901234567890
  expectedQuotientNumberStr := "1"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_04(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_04"

  originalDividendNumberStr := "0.05"

  originalDivisorNumberStr := "24"

  maxPrecision := 32

  //                                       1         2         3
  //                            0.12345678901234567890123456789012
  expectedQuotientNumberStr := "0.00208333333333333333333333333333"

  expectedPrecisionUint := uint(32)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_05(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_05"

  originalDividendNumberStr := "0"

  originalDivisorNumberStr := "24"

  maxPrecision := 7

  //                                       1         2         3
  //                            0.12345678901234567890123456789012
  expectedQuotientNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_06(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_06"

  originalDividendNumberStr := "48"

  originalDivisorNumberStr := "0"

  maxPrecision := 15

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  _, err = intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected a Divide by Zero Error, BUT NO ERROR WAS RETURNED!\n"+
      "intAryDivisor= 0\n"+
      "Division by zero should produce an error.\n\n", ePrefix)
    return
  }

  return
}

func TestIntAry_DivideThisBy_07(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_07"

  originalDividendNumberStr := "-9360"

  originalDivisorNumberStr := "24.48"

  maxPrecision := 29

  //                                          1         2         3
  //                               0.12345678901234567890123456789012
  expectedQuotientNumberStr := "-382.35294117647058823529411764706"

  expectedPrecisionUint := uint(29)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_08(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_08"

  originalDividendNumberStr := "-9360"

  originalDivisorNumberStr := "-24.48"

  maxPrecision := 29

  //                                         1         2         3
  //                              0.12345678901234567890123456789012
  expectedQuotientNumberStr := "382.35294117647058823529411764706"

  expectedPrecisionUint := uint(29)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_09(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_09"

  originalDividendNumberStr := "9360"

  originalDivisorNumberStr := "-24.48"

  maxPrecision := 29

  //                                          1         2         3
  //                               0.12345678901234567890123456789012
  expectedQuotientNumberStr := "-382.35294117647058823529411764706"

  expectedPrecisionUint := uint(29)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_10(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_10"

  originalDividendNumberStr := "-19260.549"

  originalDivisorNumberStr := "-246.483"

  maxPrecision := 30

  //                                        1         2         3
  //                             0.12345678901234567890123456789012
  expectedQuotientNumberStr := "78.141490488187826340964691276883"

  expectedPrecisionUint := uint(30)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_11(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_11"

  // Testing condition where maxPrecision is
  // less than significant result.

  originalDividendNumberStr := "1"

  originalDivisorNumberStr := "25"

  maxPrecision := 1

  // actual result is "0.04"

  //                                       1         2         3
  //                            0.12345678901234567890123456789012
  expectedQuotientNumberStr := "0.0"

  expectedPrecisionUint := uint(1)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_12(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_12"

  // Testing condition where maxPrecision is
  // greater than significant result.

  originalDividendNumberStr := "1"

  originalDivisorNumberStr := "25"

  maxPrecision := 10 // maxPrecision exceeds actual result

  // actual result is "0.04"

  //                                       1         2         3
  //                            0.12345678901234567890123456789012
  expectedQuotientNumberStr := "0.04"

  expectedPrecisionUint := uint(2)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_13(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_13"

  // Testing condition where maxPrecision is
  // set equal to -1.

  originalDividendNumberStr := "1"

  originalDivisorNumberStr := "26"

  maxPrecision := -1 // maxPrecision exceeds actual result

  expectedQuotientNumberStr := "0.0384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615385"

  expectedPrecisionUint := uint(4096)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_14(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_14"

  // Testing condition where maxPrecision is
  // invalid

  originalDividendNumberStr := "1"

  originalDivisorNumberStr := "25"

  maxPrecision := -10 // maxPrecision exceeds actual result

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  _, err = intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Function Call:\n"+
      " _, err = intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "maxPrecision= %v\n"+
      "'maxPrecision' IS INVALID!\n"+
      "Invalid 'maxPrecision' should produce an error.\n\n", ePrefix, maxPrecision)
    return
  }

  return
}

func TestIntAry_DivideThisBy_15(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_15"

  // Test Condition were actual decimal places are
  // 21+ and maxPrecision is set to zero

  originalDividendNumberStr := "56234369384300"

  originalDivisorNumberStr := "24"

  maxPrecision := 0

  // actual eQuotient := "2343098724345.833333333333333333333"

  expectedQuotientNumberStr := "2343098724346"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}

func TestIntAry_DivideThisBy_16(t *testing.T) {

  ePrefix := "TestIntAry_DivideThisBy_16"

  // Test Condition were actual decimal places are
  // 21+ and maxPrecision is set to zero

  originalDividendNumberStr := "1"

  originalDivisorNumberStr := "5132188731375616"

  maxPrecision := 1024

  expectedQuotientNumberStr := "0.0000000000000001948486410654588479586365019367735312710807776562419879994613957690782179395268844976594976574190856910462423734806284591100904041759223999872534562498706764291924041486253101558649510819197172143928398136686955522152677536556303208645361801676397426605377561320278442021917586553676450097223493264453235937316338207393368423631288419620924881566999933402357520007410670136732435322811808257677794240393475046309703702321880635871858096536620946259448966209463978714360724504546027791919182097769678551824119271867445986093109132873712847086353082760161728835291168886244778340678734742887929927691392086547229915192975721487525948779793899832544746398668414935327086017026715079825557839286764777130808629545941416744923613139031065794055208370453857117267343145382178198213289314418222458170836304178717692361810139999277887517350389701068636718076181550016169308117583049414576551610275225085122626807472053468620691915924510100661007669795710586291378217910587138791895083422555361272265231787002422046566"

  expectedPrecisionUint := uint(1024)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryDividend, err := new(IntAry).NewNumStr(originalDividendNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDividend, err := new(IntAry).\n"+
      "  NewNumStr(originalDividendNumberStr)\n"+
      "originalDividendNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDividendNumberStr, err.Error())
    return
  }

  intAryDividendNumberStr, err := intAryDividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDividendNumberStr != intAryDividendNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Dividend Number String Values ARE NOT Equal\n"+
      "Because originalDividendNumberStr != intAryDividendNumberStr\n"+
      "Expected intAryDividendNumberStr = '%v'\n"+
      "  Actual intAryDividendNumberStr = '%v'\n\n",
      ePrefix, originalDividendNumberStr, intAryDividendNumberStr)

    return
  }

  intAryDivisor, err := new(IntAry).NewNumStr(originalDivisorNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDivisor, err := new(IntAry).\n"+
      "  NewNumStr(originalDivisorNumberStr)\n"+
      "originalDivisorNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalDivisorNumberStr, err.Error())
    return
  }

  err = intAryDivisor.IsValid("Validating initial intAryDivisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDivisor.IsValid('Validating initial intAryDivisor')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalDivisorNumberStr != intAryDivisorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Divisor Number String Values ARE NOT Equal\n"+
      "Because originalDivisorNumberStr != intAryDivisorNumberStr\n"+
      "Expected intAryDivisorNumberStr = '%v'\n"+
      "  Actual intAryDivisorNumberStr = '%v'\n\n",
      ePrefix, originalDivisorNumberStr, intAryDivisorNumberStr)

    return
  }

  intAryQuotient, err := intAryDividend.DivideThisBy(&intAryDivisor, 0, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotient, err := intAryDividend.DivideThisBy(\n"+
      "  &intAryDivisor, 0, maxPrecision)\n"+
      "intAryDividend= '%v'\n"+
      "intAryDivisor= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryDividendNumberStr,
      intAryDivisorNumberStr,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryQuotient.IsValid("Validating intAryQuotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryQuotient.IsValid('Validating intAryQuotient')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientPrecisionUint, err :=\n"+
      "  intAryQuotient.GetPrecisionUint()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientSignValue, err := intAryQuotient.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
      "intAryQuotient= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
      "intAryQuotient= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
    return
  }

  if expectedQuotientNumberStr != intAryQuotientNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedQuotientNumberStr != intAryQuotientNumberStr \n"+
      "Expected intAryQuotientNumberStr = '%v'\n"+
      "  Actual intAryQuotientNumberStr = '%v'\n\n",
      ePrefix, expectedQuotientNumberStr, intAryQuotientNumberStr)

    return
  }

  if expectedPrecisionUint != intAryQuotientPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
      "Expected intAryQuotientPrecisionUint = '%v'\n"+
      "  Actual intAryQuotientPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

    return
  }

  if expectedSignValue != intAryQuotientSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryQuotientSignValue\n"+
      "Expected intAryQuotientSignValue = '%v'\n"+
      "  Actual intAryQuotientSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryQuotientSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryQuotientNumSeps \n"+
      "Expected intAryQuotientNumSeps = '%v'\n"+
      "  Actual intAryQuotientNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

    return
  }

  return
}
