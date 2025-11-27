package mathops

import "testing"

func TestIntAryMathMultiply_MultiplyInPlace_01(t *testing.T) {

  ePrefix := "TestIntAryMathMultiply_MultiplyInPlace_01"

  originalNumberStr1 := "45.762"

  originalNumberStr2 := "12.851"

  minimumPrecision := 6

  maximumPrecision := 6

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "588.087462"

  expectedPrecisionInt := 6

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryResult, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAryResult.IsValid("Validating initial intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating initial intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr1 != intAryResultNumberStr\n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAryResultNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = new(IntAryMathMultiply).MultiplyInPlace(&intAryResult, &intAry2, minimumPrecision, maximumPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathMultiply).MultiplyInPlace(\n"+
      "  &intAryResult, &intAry2, minimumPrecision, maximumPrecision)\n"+
      "intAryResult= '%v'\n"+
      "intAry2= '%v'\n"+
      "minimumPrecision= '%v'\n"+
      "maximumPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryResultNumberStr,
      intAry2NumberStr,
      minimumPrecision,
      maximumPrecision,
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

  intAryResultNumberStr, err = intAryResult.GetNumStr()

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

func TestIntAryMathMultiply_MultiplyInPlace_02(t *testing.T) {

  ePrefix := "TestIntAryMathMultiply_MultiplyInPlace_02"

  originalNumberStr1 := "-45"

  originalNumberStr2 := "35"

  minimumPrecision := -1

  maximumPrecision := -1

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr := "-1575"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryResult, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAryResult.IsValid("Validating initial intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating initial intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr1 != intAryResultNumberStr\n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAryResultNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr2 != intAry2NumberStr\n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = new(IntAryMathMultiply).MultiplyInPlace(&intAryResult, &intAry2, minimumPrecision, maximumPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathMultiply).MultiplyInPlace(\n"+
      "  &intAryResult, &intAry2, minimumPrecision, maximumPrecision)\n"+
      "intAryResult= '%v'\n"+
      "intAry2= '%v'\n"+
      "minimumPrecision= '%v'\n"+
      "maximumPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryResultNumberStr,
      intAry2NumberStr,
      minimumPrecision,
      maximumPrecision,
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

  intAryResultNumberStr, err = intAryResult.GetNumStr()

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

func TestIntAryMathMultiply_Multiply_01(t *testing.T) {

  ePrefix := "TestIntAryMathMultiply_Multiply_01"

  originalNumberStr1 := "92.135"

  originalNumberStr2 := "5.76405"

  minimumPrecision := -1

  maximumPrecision := -1

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "531.07074675"

  expectedPrecisionInt := 8

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  originalNumberStr1, expectedNumSeps)\n"+
      "originalNumberStr1= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr1,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err := intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
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

  intAry2, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  originalNumberStr2, expectedNumSeps)\n"+
      "originalNumberStr2= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr2,
      expectedNumSeps.String(),
      err.Error())

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

  intAryResult, err := new(IntAry).NewZero(0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAry).NewZero(0)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = intAryResult.IsValid("Validating initial intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating initial intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = new(IntAryMathMultiply).Multiply(&intAry1, &intAry2, &intAryResult, minimumPrecision, maximumPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathMultiply).Multiply(\n"+
      "  &intAry1, &intAry2, &intAryResult, minimumPrecision, maximumPrecision)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "intAryResult= '%v'\n"+
      "minimumPrecision= '%v'\n"+
      "maximumPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      intAryResultNumberStr,
      minimumPrecision,
      maximumPrecision,
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

  intAryResultNumberStr, err = intAryResult.GetNumStr()

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

func TestIntAryMathMultiply_MultiplyByTwoToPower_01(t *testing.T) {

  ePrefix := "TestIntAryMathMultiply_MultiplyByTwoToPower_01"

  originalNumberStr1 := "5.2"

  targetPower := uint(4)

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr := "83.2"

  expectedPrecisionInt := 8

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryResult, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  originalNumberStr1, expectedNumSeps)\n"+
      "originalNumberStr1= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr1,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating initial intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating initial intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResultNumberStr set to initial intAryResult value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAryResultNumberStr\n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAryResultNumberStr)

    return
  }

  err = new(IntAryMathMultiply).MultiplyByTwoToPower(&intAryResult, targetPower)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathMultiply).MultiplyByTwoToPower(\n"+
      "  &intAryResult, targetPower)\n"+
      "intAryResult= '%v'\n"+
      "targetPower= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryResultNumberStr,
      targetPower,
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

  intAryResultNumberStr, err = intAryResult.GetNumStr()

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

func TestIntAryMathMultiply_MultiplyByTenToPower_01(t *testing.T) {

  ePrefix := "TestIntAryMathMultiply_MultiplyByTenToPower_01"

  originalNumberStr1 := "5.2"

  targetTenPower := uint(3)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "5200"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryResult, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = intAryResult.IsValid("Validating initial intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating initial intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResultNumberStr set to initial intAryResult value.\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAryResultNumberStr\n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAryResultNumberStr)

    return
  }

  err = new(IntAryMathMultiply).MultiplyByTenToPower(&intAryResult, targetTenPower)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = new(IntAryMathMultiply).MultiplyByTenToPower(\n"+
      "  &intAryResult, targetTenPower)\n"+
      "intAryResult= '%v'\n"+
      "targetTenPower= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryResultNumberStr,
      targetTenPower,
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

  intAryResultNumberStr, err = intAryResult.GetNumStr()

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
