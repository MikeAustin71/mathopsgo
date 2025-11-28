package mathops

import "testing"

func TestIntAryMathPower_MinimumRequiredPrecision_01(t *testing.T) {

  ePrefix := "TestIntAryMathPower_MinimumRequiredPrecision_01"

  originalBaseInt := 312

  originalBasePrecisionUint := uint(2)

  originalBaseNumberStr := "3.12"

  originalExponentInt := 4

  originalExponentPrecisionUint := uint(0)

  originalExponentNumberStr := "4"

  expectedResultInt := 8

  intAryBase, err := new(IntAry).NewInt(originalBaseInt, originalBasePrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewInt(\n"+
      "  originalBaseInt, originalBasePrecisionUint)\n"+
      "originalBaseInt= '%v'\n"+
      "originalBasePrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBaseInt,
      originalBasePrecisionUint,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewInt(originalExponentInt, originalExponentPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewInt(\n"+
      "originalExponentInt, originalExponentPrecisionUint)\n"+
      "originalExponentInt= '%v'\n"+
      "originalExponentPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalExponentInt,
      originalExponentPrecisionUint,
      err.Error())

    return
  }

  err = intAryExponent.IsValid("Validating initial intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating initial intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "intAryExponentNumberStr set to initial intAryExponent value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  resultInt, err := new(IntAryMathPower).MinimumRequiredPrecision(&intAryBase, &intAryExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultInt, err := new(IntAryMathPower).MinimumRequiredPrecision(\n"+
      "  &intAryBase, &intAryExponent)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      err.Error())

    return
  }

  if expectedResultInt != resultInt {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedResultInt, resultInt)
  }

  if expectedResultInt != resultInt {
    t.Errorf("%v\n"+
      "Error: Expected and Actual ResultInt's ARE NOT EQUAL!\n"+
      "Because expectedResultInt != resultInt\n"+
      "Expected resultInt = '%v'\n"+
      "  Actual resultInt = '%v'\n\n",
      ePrefix, expectedResultInt, resultInt)

    return
  }

  return
}

func TestIntAryMathPower_MinimumRequiredPrecision_02(t *testing.T) {

  ePrefix := "TestIntAryMathPower_MinimumRequiredPrecision_02"

  originalBaseInt := 312345

  originalBasePrecisionUint := uint(5)

  originalBaseNumberStr := "3.12345"

  originalExponentInt := 18

  originalExponentPrecisionUint := uint(0)

  originalExponentNumberStr := "18"

  expectedResultInt := 90

  intAryBase, err := new(IntAry).NewInt(originalBaseInt, originalBasePrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewInt(\n"+
      "  originalBaseInt, originalBasePrecisionUint)\n"+
      "originalBaseInt= '%v'\n"+
      "originalBasePrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBaseInt,
      originalBasePrecisionUint,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewInt(originalExponentInt, originalExponentPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewInt(\n"+
      "originalExponentInt, originalExponentPrecisionUint)\n"+
      "originalExponentInt= '%v'\n"+
      "originalExponentPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalExponentInt,
      originalExponentPrecisionUint,
      err.Error())

    return
  }

  err = intAryExponent.IsValid("Validating initial intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating initial intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "intAryExponentNumberStr set to initial intAryExponent value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  resultInt, err := new(IntAryMathPower).MinimumRequiredPrecision(&intAryBase, &intAryExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultInt, err := new(IntAryMathPower).MinimumRequiredPrecision(\n"+
      "  &intAryBase, &intAryExponent)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      err.Error())

    return
  }

  if expectedResultInt != resultInt {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedResultInt, resultInt)
  }

  if expectedResultInt != resultInt {
    t.Errorf("%v\n"+
      "Error: Expected and Actual ResultInt's ARE NOT EQUAL!\n"+
      "Because expectedResultInt != resultInt\n"+
      "Expected resultInt = '%v'\n"+
      "  Actual resultInt = '%v'\n\n",
      ePrefix, expectedResultInt, resultInt)

    return
  }

  return
}

func TestIntAryMathPower_MinimumRequiredPrecision_03(t *testing.T) {

  ePrefix := "TestIntAryMathPower_MinimumRequiredPrecision_03"

  originalBaseInt := -312345

  originalBasePrecisionUint := uint(5)

  originalBaseNumberStr := "-3.12345"

  originalExponentInt := 18

  originalExponentPrecisionUint := uint(0)

  originalExponentNumberStr := "18"

  expectedResultInt := 90

  intAryBase, err := new(IntAry).NewInt(originalBaseInt, originalBasePrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewInt(\n"+
      "  originalBaseInt, originalBasePrecisionUint)\n"+
      "originalBaseInt= '%v'\n"+
      "originalBasePrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBaseInt,
      originalBasePrecisionUint,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewInt(originalExponentInt, originalExponentPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewInt(\n"+
      "originalExponentInt, originalExponentPrecisionUint)\n"+
      "originalExponentInt= '%v'\n"+
      "originalExponentPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalExponentInt,
      originalExponentPrecisionUint,
      err.Error())

    return
  }

  err = intAryExponent.IsValid("Validating initial intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating initial intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "intAryExponentNumberStr set to initial intAryExponent value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  resultInt, err := new(IntAryMathPower).MinimumRequiredPrecision(&intAryBase, &intAryExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultInt, err := new(IntAryMathPower).MinimumRequiredPrecision(\n"+
      "  &intAryBase, &intAryExponent)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      err.Error())

    return
  }

  if expectedResultInt != resultInt {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedResultInt, resultInt)
  }

  if expectedResultInt != resultInt {
    t.Errorf("%v\n"+
      "Error: Expected and Actual ResultInt's ARE NOT EQUAL!\n"+
      "Because expectedResultInt != resultInt\n"+
      "Expected resultInt = '%v'\n"+
      "  Actual resultInt = '%v'\n\n",
      ePrefix, expectedResultInt, resultInt)

    return
  }

  return
}

func TestIntAryMathPower_MinimumRequiredPrecision_04(t *testing.T) {

  ePrefix := "TestIntAryMathPower_MinimumRequiredPrecision_04"

  originalBaseInt := 312345

  originalBasePrecisionUint := uint(5)

  originalBaseNumberStr := "3.12345"

  originalExponentInt := -18

  originalExponentPrecisionUint := uint(0)

  originalExponentNumberStr := "-18"

  expectedResultInt := 90

  intAryBase, err := new(IntAry).NewInt(originalBaseInt, originalBasePrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewInt(\n"+
      "  originalBaseInt, originalBasePrecisionUint)\n"+
      "originalBaseInt= '%v'\n"+
      "originalBasePrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBaseInt,
      originalBasePrecisionUint,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewInt(originalExponentInt, originalExponentPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewInt(\n"+
      "originalExponentInt, originalExponentPrecisionUint)\n"+
      "originalExponentInt= '%v'\n"+
      "originalExponentPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalExponentInt,
      originalExponentPrecisionUint,
      err.Error())

    return
  }

  err = intAryExponent.IsValid("Validating initial intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating initial intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "intAryExponentNumberStr set to initial intAryExponent value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  resultInt, err := new(IntAryMathPower).MinimumRequiredPrecision(&intAryBase, &intAryExponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultInt, err := new(IntAryMathPower).MinimumRequiredPrecision(\n"+
      "  &intAryBase, &intAryExponent)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      err.Error())

    return
  }

  if expectedResultInt != resultInt {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedResultInt, resultInt)
  }

  if expectedResultInt != resultInt {
    t.Errorf("%v\n"+
      "Error: Expected and Actual ResultInt's ARE NOT EQUAL!\n"+
      "Because expectedResultInt != resultInt\n"+
      "Expected resultInt = '%v'\n"+
      "  Actual resultInt = '%v'\n\n",
      ePrefix, expectedResultInt, resultInt)

    return
  }

  return
}

func TestIntAryMathPower_MinimumRequiredPrecision_05(t *testing.T) {

  ePrefix := "TestIntAryMathPower_MinimumRequiredPrecision_04"

  originalBaseInt := 312345

  originalBasePrecisionUint := uint(5)

  originalBaseNumberStr := "3.12345"

  originalExponentUint64 := uint64(12345678901234567890)

  originalExponentPrecisionUint := uint(0)

  originalExponentSignValue := 1

  originalExponentNumberStr := "12345678901234567890"

  intAryBase, err := new(IntAry).NewInt(originalBaseInt, originalBasePrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewInt(\n"+
      "  originalBaseInt, originalBasePrecisionUint)\n"+
      "originalBaseInt= '%v'\n"+
      "originalBasePrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBaseInt,
      originalBasePrecisionUint,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewUint64(originalExponentUint64, originalExponentSignValue, originalExponentPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewUint64(\n"+
      "  originalExponentUint64, originalExponentSignValue, originalExponentPrecisionUint)\n"+
      "originalExponentUint64= '%v'\n"+
      "originalExponentSignValue= '%v'\n"+
      "originalExponentPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalExponentUint64,
      originalExponentSignValue,
      originalExponentPrecisionUint,
      err.Error())

    return
  }

  err = intAryExponent.IsValid("Validating initial intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating initial intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "intAryExponentNumberStr set to initial intAryExponent value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  _, err = new(IntAryMathPower).MinimumRequiredPrecision(&intAryBase, &intAryExponent)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "  _, err := new(IntAryMathPower).MinimumRequiredPrecision(\n"+
      "    &intAryBase, &intAryExponent)\n"+
      "MinimumRequiredPrecision Exceeded +2,147,483,646 and should produce and error!\n\n",
      ePrefix)
    return
  }

  return
}

func TestIntAryMathPower_Pwr_01(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_01"

  originalBaseNumberStr := "2"

  originalExponentNumberStr := "4"

  maximumPrecisionInt := 17

  minimumPrecisionInt := 0

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr := "16"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_02(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_02"

  originalBaseNumberStr := "2"

  originalExponentNumberStr := "-4"

  maximumPrecisionInt := 17

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.0625"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_03(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_03"

  originalBaseNumberStr := "37.241"

  originalExponentNumberStr := "8"

  maximumPrecisionInt := 19

  minimumPrecisionInt := 0

  //                                           1         2         3
  //                                0.1234567890123456789012345678901234567
  expectedNumberStr := "3699735472699.4101912057680101525"

  expectedPrecisionInt := 19

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_04(t *testing.T) {

  // Elapsed Time
  // 2-Milliseconds 998-Microseconds 700-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_04"

  originalBaseNumberStr := "37"

  originalExponentNumberStr := "3.25"

  maximumPrecisionInt := 26

  minimumPrecisionInt := 0

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr := "124926.79641959048051506133768818"

  expectedPrecisionInt := 26

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_05(t *testing.T) {

  // Elapsed Time
  // 3-Milliseconds 997-Microseconds 400-Nanosecond

  ePrefix := "TestIntAryMathPower_Pwr_05"

  originalBaseNumberStr := "37"

  originalExponentNumberStr := "-3.25"

  maximumPrecisionInt := 37

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.0000080046877744411952288377104402677"

  expectedPrecisionInt := 37

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_06(t *testing.T) {

  // Elapseded Time
  // 1-Milliseconds 998-Microseconds 300-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_06"

  originalBaseNumberStr := "-37"

  originalExponentNumberStr := "-3"

  maximumPrecisionInt := 36

  minimumPrecisionInt := 0

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr := "-0.000019742167295125658894833474818866"

  expectedPrecisionInt := 36

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_07(t *testing.T) {

  // Elapseded Time
  // 3-Milliseconds 997-Microseconds 500-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_07"

  originalBaseNumberStr := "32"

  originalExponentNumberStr := "-3.6"

  maximumPrecisionInt := 18

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.000003814697265625"

  expectedPrecisionInt := 18

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_08(t *testing.T) {

  // Elapseded Time
  // 2-Milliseconds 999-Microseconds 300-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_08"

  originalBaseNumberStr := "-32"

  originalExponentNumberStr := "-3.6"

  maximumPrecisionInt := 18

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.000003814697265625"

  expectedPrecisionInt := 18

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_09(t *testing.T) {

  // Elapseded Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_09"

  originalBaseNumberStr := "5"

  originalExponentNumberStr := "-3"

  maximumPrecisionInt := 3

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.008"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_10(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_10"

  originalBaseNumberStr := "-5"

  originalExponentNumberStr := "4"

  maximumPrecisionInt := 0

  minimumPrecisionInt := 0

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "625"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_11(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_11"

  originalBaseNumberStr := "-5"

  originalExponentNumberStr := "5"

  maximumPrecisionInt := 0

  minimumPrecisionInt := 0

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr := "-3125"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_12(t *testing.T) {

  // Elapsed Time
  // 1-Milliseconds 998-Microseconds 900-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_12"

  originalBaseNumberStr := "4"

  originalExponentNumberStr := "0.25"

  maximumPrecisionInt := 31

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "1.4142135623730950488016887242097"

  expectedPrecisionInt := 31

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_13(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_13"

  originalBaseNumberStr := "45"

  originalExponentNumberStr := "120"

  maximumPrecisionInt := 200

  minimumPrecisionInt := 0

  //                             1         2         3         4         5         6         7         8         9        10        11        12        13        14        15        16        17        18        19        20
  //                    12345678901234567890123456789012345675901234567890123456789012345678901234567590123456789012345678901234567890123456759012345678901234567890123456789012345675901234567890123456789012345678901234567590
  expectedNumberStr := "2429414689006507011047680668198610544614376056243160093821961151708217872022541081600097552192834563575596196518385368260399467853246465311637303360756939677768395657864175518625415861606597900390625"

  expectedPrecisionInt := 199

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_14(t *testing.T) {

  // Elapsed Time
  // 4-Milliseconds 997-Microseconds 300-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_14"

  originalBaseNumberStr := "19"

  originalExponentNumberStr := "2.3"

  maximumPrecisionInt := 29

  minimumPrecisionInt := 0

  //                                 1         2         3
  //                        123456489012345678901234567890
  expectedNumberStr := "873.23931881701910176203214553167"

  expectedPrecisionInt := 29

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_15(t *testing.T) {

  // Elapsed Time
  // 8-Milliseconds 47-Microseconds 300-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_15"

  originalBaseNumberStr := "19"

  originalExponentNumberStr := "-2.3"

  maximumPrecisionInt := 32

  minimumPrecisionInt := 0

  //                               1         2         3
  //                      12345648901234567890123456789012
  expectedNumberStr := "0.00114516144480839927662319331457"

  expectedPrecisionInt := 32

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_16(t *testing.T) {

  // Elapsed Time
  // 2-Milliseconds 998-Microseconds 200-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_16"

  originalBaseNumberStr := "45"

  originalExponentNumberStr := "-2"

  maximumPrecisionInt := 35

  minimumPrecisionInt := 0

  //                               1         2         3
  //                      12345648901234567890123456789012345
  expectedNumberStr := "0.00049382716049382716049382716049383"

  expectedPrecisionInt := 35

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_17(t *testing.T) {

  ePrefix := "TestIntAryMathPower_Pwr_17"

  originalBaseNumberStr := "0"

  originalExponentNumberStr := "5"

  maximumPrecisionInt := 35

  minimumPrecisionInt := 0

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent,\n"+
      "     minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt)

    return
  }

  return
}

func TestIntAryMathPower_Pwr_18(t *testing.T) {

  // Elapsed Time
  //  0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_18"

  originalBaseNumberStr := "45.1"

  originalExponentNumberStr := "0"

  maximumPrecisionInt := 5

  minimumPrecisionInt := 0

  //                              1         2         3
  //                    0.12345648901234567890123456789012345
  expectedNumberStr := "1"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_19(t *testing.T) {

  // Elapsed Time:
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_19"

  originalBaseNumberStr := "45.1"

  originalExponentNumberStr := "1"

  maximumPrecisionInt := 5

  minimumPrecisionInt := 0

  //                                1         2         3
  //                     0.12345648901234567890123456789012345
  expectedNumberStr := "45.1"

  expectedPrecisionInt := 1

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_20(t *testing.T) {

  // Elapsed Time:
  // 0-Milliseconds 0-Microseconds 0-Nanosecond

  ePrefix := "TestIntAryMathPower_Pwr_20"

  originalBaseNumberStr := "-45.1"

  originalExponentNumberStr := "1"

  maximumPrecisionInt := 5

  minimumPrecisionInt := 0

  //                                 1         2         3
  //                      0.12345648901234567890123456789012345
  expectedNumberStr := "-45.1"

  expectedPrecisionInt := 1

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_21(t *testing.T) {

  // Elapsed Time:
  // 6-Milliseconds 995-Microseconds 900-Nanoseconds

  ePrefix := "TestIntAryMathPower_Pwr_21"

  originalBaseNumberStr := "-45.6"

  originalExponentNumberStr := "-3.2"

  maximumPrecisionInt := 35

  minimumPrecisionInt := 0

  //                      12345678901234567890123456789012345
  //                    0.00000491261243811417457984700270545
  // 4.91261243811417457984700270545e-6

  //                               1         2         3
  //                    0.12345648901234567890123456789012345
  expectedNumberStr := "0.00000491261243811417457984700270545"

  expectedPrecisionInt := 35

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_Pwr_22(t *testing.T) {

  ePrefix := "TestIntAryMathPower_Pwr_22"

  originalBaseNumberStr := "-45.632"

  originalExponentNumberStr := "-1.01579"

  maximumPrecisionInt := 15

  minimumPrecisionInt := 0

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent,\n"+
      "     minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt)

    return
  }

  return
}

func TestIntAryMathPower_Pwr_23(t *testing.T) {

  ePrefix := "TestIntAryMathPower_Pwr_21"

  originalBaseNumberStr := "2.125"

  originalExponentNumberStr := "-5"

  maximumPrecisionInt := 32

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.12345648901234567890123456789012345
  expectedNumberStr := "0.02307838042845159759046157465153"

  expectedPrecisionInt := 32

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
      "originalBaseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalBaseNumberStr, err.Error())
    return
  }

  err = intAryBase.IsValid("Validating initial intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating initial intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err := intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBaseNumberStr set to initial intAryBase value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalBaseNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalBaseNumberStr != intAryBaseNumberStr\n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

    return
  }

  intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponent, err := new(IntAry).NewNumStr(originalExponentNumberStr)\n"+
      "originalExponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalExponentNumberStr, err.Error())
    return
  }

  err = intAryExponent.IsValid("Validating intAryExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExponent.IsValid('Validating intAryExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExponentNumberStr, err := intAryExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExponentNumberStr, err := intAryExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExponentNumberStr != intAryExponentNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExponentNumberStr != intAryExponentNumberStr\n"+
      "Expected intAryExponentNumberStr = '%v'\n"+
      "  Actual intAryExponentNumberStr = '%v'\n\n",
      ePrefix, originalExponentNumberStr, intAryExponentNumberStr)

    return
  }

  err = new(IntAryMathPower).Pwr(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathPower).Pwr(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt,
      err.Error())

    return
  }

  err = intAryBase.IsValid("Validating final intAryBase")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.IsValid('Validating final intAryBase')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBaseNumberStr, err = intAryBase.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
      "intAryBase Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryBasePrecisionInt := intAryBase.GetPrecision()

  intAryBasePrecisionUint, err := intAryBase.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBasePrecisionUint, err :=\n"+
      "  intAryBase.GetPrecisionUint()\n"+
      "intAryBaseResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseSignValue, err := intAryBase.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseSignValue, err := intAryBase.GetSign()\n"+
      "intAryBase= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryBaseNumSeps, err := intAryBase.GetNumericSeparatorsDto()\n"+
      "intAryBase= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryBaseNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryBaseNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryBaseNumberStr \n"+
      "Expected intAryBaseNumberStr = '%v'\n"+
      "  Actual intAryBaseNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryBaseNumberStr)

    return
  }

  if expectedPrecisionInt != intAryBasePrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryBasePrecisionInt\n"+
      "Expected intAryBasePrecisionInt = '%v'\n"+
      "  Actual intAryBasePrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryBasePrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryBasePrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryBase Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryBasePrecisionUint\n"+
      "Expected intAryBasePrecisionUint = '%v'\n"+
      "  Actual intAryBasePrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryBasePrecisionUint)

    return
  }

  if expectedSignValue != intAryBaseSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryBase Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryBaseSignValue\n"+
      "Expected intAryBaseSignValue = '%v'\n"+
      "  Actual intAryBaseSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryBaseSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryBaseNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryBaseNumSeps \n"+
      "Expected intAryBaseNumSeps = '%v'\n"+
      "  Actual intAryBaseNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryBaseNumSeps.String())

    return
  }

  return
}
