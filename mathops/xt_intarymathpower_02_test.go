package mathops

import (
  "testing"
)

func TestIntAryMathPower_PwrByMultiplication_01(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_01"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_02(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_02"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_03(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 999-Microseconds 500-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_03"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_04(t *testing.T) {

  // Elapsed Time
  // 3-Milliseconds 998-Microseconds 600-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_04"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_05(t *testing.T) {

  // Elapsed Time
  // 5-Milliseconds 996-Microseconds 500-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_05"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_06(t *testing.T) {

  // Elapsed Time
  // 1-Milliseconds 998-Microseconds 300-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_06"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_07(t *testing.T) {

  // Elapsed Time
  // 3-Milliseconds 997-Microseconds 400-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_07"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_08(t *testing.T) {

  // Elapsed Time
  // 3-Milliseconds 997-Microseconds 700-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_08"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_09(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_09"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_10(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_10"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_11(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_11"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_12(t *testing.T) {

  // Elapsed Time
  // 1-Milliseconds 999-Microseconds 200-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_12"

  originalBaseNumberStr := "4"

  originalExponentNumberStr := "0.25"

  maximumPrecisionInt := 31

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "1.4142135623730950488016887242097"

  expectedPrecisionInt := 31

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_13(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_13"

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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_14(t *testing.T) {

  // Elapsed Time
  // 6-Milliseconds 995-Microseconds 800-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_14"

  originalBaseNumberStr := "19"

  originalExponentNumberStr := "2.3"

  maximumPrecisionInt := 29

  minimumPrecisionInt := 0

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_15(t *testing.T) {

  // Elapsed Time
  // 8-Milliseconds 994-Microseconds 100-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_15"

  originalBaseNumberStr := "19"

  originalExponentNumberStr := "-2.3"

  maximumPrecisionInt := 32

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_16(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 999-Microseconds 600-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_16"

  originalBaseNumberStr := "45"

  originalExponentNumberStr := "-2"

  maximumPrecisionInt := 35

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_17(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 999-Microseconds 600-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_17"

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

  _, err = new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent,\n"+
      "  minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "An Error should have been produced!\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt)
    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_18(t *testing.T) {

  // Elapsed Time
  //  0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_18"

  originalBaseNumberStr := "45.1"

  originalExponentNumberStr := "0"

  maximumPrecisionInt := 5

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_19(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_19"

  originalBaseNumberStr := "45.1"

  originalExponentNumberStr := "1"

  maximumPrecisionInt := 5

  minimumPrecisionInt := 0

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_20(t *testing.T) {

  // Elapsed Time
  // 0-Milliseconds 0-Microseconds 0-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_20"

  originalBaseNumberStr := "-45.1"

  originalExponentNumberStr := "1"

  maximumPrecisionInt := 5

  minimumPrecisionInt := 0

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_21(t *testing.T) {

  // Elapsed Time
  // 11-Milliseconds 993-Microseconds 200-Nanoseconds

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_21"

  originalBaseNumberStr := "-45.6"

  originalExponentNumberStr := "-3.2"

  maximumPrecisionInt := 35

  minimumPrecisionInt := 0

  //                      12345678901234567890123456789012345
  //                    0.00000491261243811417457984700270545
  // 4.91261243811417457984700270545e-6

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_22(t *testing.T) {

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_22"

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

  _, err = new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent,\n"+
      "  minimumPrecisionInt, maximumPrecisionInt)\n"+
      "intAryBase= '%v'\n"+
      "intAryExponent= '%v'\n"+
      "minimumPrecisionInt= '%v'\n"+
      "maximumPrecisionInt= '%v'\n"+
      "An Error should have been produced!\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAryExponentNumberStr,
      minimumPrecisionInt,
      maximumPrecisionInt)
    return
  }

  return
}

func TestIntAryMathPower_PwrByMultiplication_23(t *testing.T) {

  ePrefix := "TestIntAryMathPower_PwrByMultiplication_23"

  originalBaseNumberStr := "2.125"

  originalExponentNumberStr := "-5"

  maximumPrecisionInt := 32

  minimumPrecisionInt := 0

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
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

  intAryResult, err := new(IntAryMathPower).PwrByMultiplication(&intAryBase, &intAryExponent, minimumPrecisionInt, maximumPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAryMathPower).PwrByMultiplication(\n"+
      "  &intAryBase, &intAryExponent, minimumPrecisionInt,\n"+
      "  maximumPrecisionInt)\n"+
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

  err = intAryResult.IsValid("Validating final intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating final intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryResultNumberStr \n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}
