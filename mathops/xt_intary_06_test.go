package mathops

import (
  "math/big"
  "testing"
)

func TestIntAry_OptimizeIntArrayLen_07(t *testing.T) {

  ePrefix := "TestIntAry_OptimizeIntArrayLen_07"

  //                               1         2         3
  //                    0.123456789012345678901234567890
  originalNumberStr := "0"

  //                               1         2         3
  //                    0.123456789012345678901234567890
  expectedNumberStr := "0"

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedArrayLength := 1

  var optimizeFracDigits bool

  optimizeFracDigits = true

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

  err = intAry.IsValid("Validating intAry Before Optimize")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry Before Optimize')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = intAry.OptimizeIntArrayLen(optimizeFracDigits)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.OptimizeIntArrayLen(optimizeFracDigits)\n"+
      "optimizeFracDigits= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, optimizeFracDigits, err.Error())
    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  actualIntAryLength := intAry.GetIntAryLength()

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

  if expectedArrayLength != actualIntAryLength {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedArrayLength != actualIntAryLength\n"+
      "Expected actualIntAryLength = '%v'\n"+
      "  Actual actualIntAryLength = '%v'\n\n",
      ePrefix, expectedArrayLength, actualIntAryLength)

    return
  }

  return
}

func TestIntAry_OptimizeIntArrayLen_08(t *testing.T) {

  ePrefix := "TestIntAry_OptimizeIntArrayLen_08"

  //                                    1         2         3
  //                         0.123456789012345678901234567890
  originalNumberStr := "-00579.000000000"

  //                                  1         2         3
  //                       0.123456789012345678901234567890
  expectedNumberStr := "-579"

  expectedPrecisionInt := 0

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedArrayLength := 3

  var optimizeFracDigits bool

  optimizeFracDigits = true

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

  err = intAry.IsValid("Validating intAry Before Optimize")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry Before Optimize')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = intAry.OptimizeIntArrayLen(optimizeFracDigits)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.OptimizeIntArrayLen(optimizeFracDigits)\n"+
      "optimizeFracDigits= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, optimizeFracDigits, err.Error())
    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  actualIntAryLength := intAry.GetIntAryLength()

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

  if expectedArrayLength != actualIntAryLength {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedArrayLength != actualIntAryLength\n"+
      "Expected actualIntAryLength = '%v'\n"+
      "  Actual actualIntAryLength = '%v'\n\n",
      ePrefix, expectedArrayLength, actualIntAryLength)

    return
  }

  return
}

func TestIntAry_OptimizeIntArrayLen_09(t *testing.T) {

  ePrefix := "TestIntAry_OptimizeIntArrayLen_09"

  //                                    1         2         3
  //                         0.123456789012345678901234567890
  originalNumberStr := "-00579.123000000"

  //                                  1         2         3
  //                       0.123456789012345678901234567890
  expectedNumberStr := "-579.123"

  expectedPrecisionInt := 3

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedArrayLength := 6

  var optimizeFracDigits bool

  optimizeFracDigits = true

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

  err = intAry.IsValid("Validating intAry Before Optimize")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry Before Optimize')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = intAry.OptimizeIntArrayLen(optimizeFracDigits)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.OptimizeIntArrayLen(optimizeFracDigits)\n"+
      "optimizeFracDigits= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, optimizeFracDigits, err.Error())
    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  actualIntAryLength := intAry.GetIntAryLength()

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

  if expectedArrayLength != actualIntAryLength {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedArrayLength != actualIntAryLength\n"+
      "Expected actualIntAryLength = '%v'\n"+
      "  Actual actualIntAryLength = '%v'\n\n",
      ePrefix, expectedArrayLength, actualIntAryLength)

    return
  }

  return
}

func TestIntAry_Pow_01(t *testing.T) {

  ePrefix := "TestIntAry_Pow_01"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  originalNumberStr := "3"

  originalTargetPowerInt := 3

  originalMaxPrecisionInt := 0

  originalInternalPrecisionInt := -1

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr := "27"

  expectedPrecisionInt := originalMaxPrecisionInt

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.Pow(originalTargetPowerInt, originalMaxPrecisionInt, originalInternalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.Pow(originalTargetPowerInt, originalTargetPrecisionInt, originalInternalPrecisionInt)\n"+
      "originalTargetPowerInt= '%v'\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "originalInternalPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalTargetPowerInt,
      originalMaxPrecisionInt,
      originalInternalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

func TestIntAry_Pow_02(t *testing.T) {

  ePrefix := "TestIntAry_Pow_02"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  originalNumberStr := "-3"

  originalTargetPowerInt := 3

  originalMaxPrecisionInt := 0

  originalInternalPrecisionInt := -1

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "-27"

  expectedPrecisionInt := originalMaxPrecisionInt

  expectedSignValue := -1

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.Pow(originalTargetPowerInt, originalMaxPrecisionInt, originalInternalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.Pow(originalTargetPowerInt, originalTargetPrecisionInt, originalInternalPrecisionInt)\n"+
      "originalTargetPowerInt= '%v'\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "originalInternalPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalTargetPowerInt,
      originalMaxPrecisionInt,
      originalInternalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

func TestIntAry_Pow_03(t *testing.T) {

  ePrefix := "TestIntAry_Pow_03"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  originalNumberStr := "-3.97"

  originalTargetPowerInt := 4

  originalMaxPrecisionInt := 8

  originalInternalPrecisionInt := -1

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "248.40596881"

  expectedPrecisionInt := originalMaxPrecisionInt

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.Pow(originalTargetPowerInt, originalMaxPrecisionInt, originalInternalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.Pow(originalTargetPowerInt, originalTargetPrecisionInt, originalInternalPrecisionInt)\n"+
      "originalTargetPowerInt= '%v'\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "originalInternalPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalTargetPowerInt,
      originalMaxPrecisionInt,
      originalInternalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

func TestIntAry_Pow_04(t *testing.T) {

  ePrefix := "TestIntAry_Pow_04"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  originalNumberStr := "-3.97"

  originalTargetPowerInt := 3

  originalMaxPrecisionInt := 6

  originalInternalPrecisionInt := -1

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "-62.570773"

  expectedPrecisionInt := originalMaxPrecisionInt

  expectedSignValue := -1

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.Pow(originalTargetPowerInt, originalMaxPrecisionInt, originalInternalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.Pow(originalTargetPowerInt, originalTargetPrecisionInt, originalInternalPrecisionInt)\n"+
      "originalTargetPowerInt= '%v'\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "originalInternalPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalTargetPowerInt,
      originalMaxPrecisionInt,
      originalInternalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

func TestIntAry_Pow_05(t *testing.T) {

  ePrefix := "TestIntAry_Pow_05"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  originalNumberStr := "4"

  originalTargetPowerInt := -3

  originalMaxPrecisionInt := 10

  originalInternalPrecisionInt := -1

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.015625"

  expectedPrecisionInt := 6

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.Pow(originalTargetPowerInt, originalMaxPrecisionInt, originalInternalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.Pow(originalTargetPowerInt, originalTargetPrecisionInt, originalInternalPrecisionInt)\n"+
      "originalTargetPowerInt= '%v'\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "originalInternalPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalTargetPowerInt,
      originalMaxPrecisionInt,
      originalInternalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

func TestIntAry_Pow_06(t *testing.T) {

  ePrefix := "TestIntAry_Pow_06"

  originalNumberStr := "92"

  originalTargetPowerInt := -8

  originalMaxPrecisionInt := 350

  originalInternalPrecisionInt := -1

  expectedNumberStr := "0.00000000000000019484864106545884795863650193677353127108077765624198799946139576907821793952688449765949765741908569104624237348062845911009040417592239998725345624987067642919240414862531015586495108191971721439283981366869555221526775365563032086453618016763974266053775613202784420219175865536764500972234932644532359373163382073933684236312884196"

  expectedPrecisionInt := 350

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.Pow(originalTargetPowerInt, originalMaxPrecisionInt, originalInternalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.Pow(originalTargetPowerInt, originalTargetPrecisionInt, originalInternalPrecisionInt)\n"+
      "originalTargetPowerInt= '%v'\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "originalInternalPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalTargetPowerInt,
      originalMaxPrecisionInt,
      originalInternalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

func TestIntAry_ResetFromBackUp_01(t *testing.T) {

  ePrefix := "TestIntAry_ResetFromBackUp_01"

  originalNumberStr1 := "99.4564"

  originalNumberStr2 := "-5034.123"

  expectedSignValue := 1

  expectedPrecisionInt := 4

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
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

  if originalNumberStr1 != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr1 != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAryNumberStr)

    return
  }

  err = intAry.CopyToBackUp()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.CopyToBackUp()\n"+
      "CopyToBackUp #1\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  err = intAry.SetIntAryWithNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "intAry= '%v'\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalNumberStr2,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating originalNumberStr2 intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating originalNumberStr2 intAry')"+
      "intAry Set to 'originalNumberStr2'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to originalNumberStr2 value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: intAry NOT Equal to originalNumberStr2\n"+
      "Because originalNumberStr2 != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAryNumberStr)

    return
  }

  err = intAry.ResetFromBackUp()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ResetFromBackUp()\n"+
      "intAry= '%v'\n"+
      "Resetting to originalNumberStr1\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating originalNumberStr1 intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating originalNumberStr1 intAry')\n"+
      "Attempted Reset to original originalNumberStr1 value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to originalNumberStr2 value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result After Reset to originalNumberStr1\n"+
      "Because originalNumberStr1 != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, intAryNumberStr)

    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

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

func TestIntAry_RoundToPrecision_01(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_01"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "999.9952"

  originalTargetPrecisionInt := 2

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "1000.00"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_RoundToPrecision_02(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_02"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalNumberStr := "-999.9952"

  originalTargetPrecisionInt := 2

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr := "-1000.00"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

  expectedSignValue := -1

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_RoundToPrecision_03(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_03"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalNumberStr := "352.954"

  originalTargetPrecisionInt := 2

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "352.95"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_RoundToPrecision_04(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_04"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalNumberStr := "-352.954"

  originalTargetPrecisionInt := 2

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "-352.95"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

  expectedSignValue := -1

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_RoundToPrecision_05(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_05"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalNumberStr := "-352.954"

  originalTargetPrecisionInt := 2

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "-352.95"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

  expectedSignValue := -1

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_RoundToPrecision_06(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_06"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "999.99"

  originalTargetPrecisionInt := 2

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "999.99"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_RoundToPrecision_07(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_07"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "999.99"

  originalTargetPrecisionInt := 5

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "999.99000"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_RoundToPrecision_08(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_08"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  originalNumberStr := "0.000"

  originalTargetPrecisionInt := 2

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.00"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_RoundToPrecision_09(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_09"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "999.995"

  originalTargetPrecisionInt := 3

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "999.995"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_RoundToPrecision_10(t *testing.T) {

  ePrefix := "TestIntAry_RoundToPrecision_10"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  originalNumberStr := "51.821810080312709972402243542716917"

  originalTargetPrecisionInt := 30

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr := "51.821810080312709972402243542717"

  expectedPrecisionInt := originalTargetPrecisionInt

  expectedPrecisionUint := uint(originalTargetPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.RoundToPrecision(originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.RoundToPrecision(originalTargetPrecisionInt)\n"+
      "intAry= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalTargetPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final value 'after rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAryResult= '%v'\n"+
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

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetAbsoluteValueThis_01(t *testing.T) {

  nStr1 := "-987.652"
  eStr := "987.652"

  ia, _ := IntAry{}.NewNumStr(nStr1)

  if nStr1 != ia.GetNumStr() {
    t.Errorf("Expected intialized NumStr= %v .  Instead, NumStr= %v  .", nStr1, ia.GetNumStr())
  }

  ia.SetAbsoluteValueThis()

  if eStr != ia.GetNumStr() {
    t.Errorf("Expected Absolute Value NumStr= %v .  Instead, NumStr= %v  .", eStr, ia.GetNumStr())
  }
}

func TestIntAry_SetAbsoluteValueThis_02(t *testing.T) {

  nStr1 := "987.652"
  eStr := "987.652"

  ia, _ := IntAry{}.NewNumStr(nStr1)

  if nStr1 != ia.GetNumStr() {
    t.Errorf("Expected intialized NumStr= %v .  Instead, NumStr= %v  .", nStr1, ia.GetNumStr())
  }

  ia.SetAbsoluteValueThis()

  if eStr != ia.GetNumStr() {
    t.Errorf("Expected Absolute Value NumStr= %v .  Instead, NumStr= %v  .", eStr, ia.GetNumStr())
  }
}

func TestIntAry_SetCurrencySymbol_01(t *testing.T) {

  ia, _ := IntAry{}.NewNumStr("50.37")

  var poundSym rune

  poundSym = '\U000000a3'

  ia.SetCurrencySymbol(poundSym)

  curSymbol := ia.GetCurrencySymbol()

  if poundSym != curSymbol {
    t.Errorf("Error: Expected Currency Symbol= '%v'. Instead, received Currency Symbol= '%v'", poundSym, curSymbol)
  }

}

func TestIntAry_SetDecimalSeparator_01(t *testing.T) {

  ia := IntAry{}.New()

  var frenchDecSeparator rune

  frenchDecSeparator = ','

  var frenchThousandsSeparator rune

  frenchThousandsSeparator = ' '

  ia.SetDecimalSeparator(frenchDecSeparator)
  ia.SetThousandsSeparator(frenchThousandsSeparator)

  ia.SetIntAryWithNumStr("450 123 647,1234")

  decimalSeparator := ia.GetDecimalSeparator()

  if frenchDecSeparator != decimalSeparator {
    t.Errorf("Error: Expected Currency Symbol= '%v'. Instead, received Currency Symbol= '%v'", frenchDecSeparator, decimalSeparator)
  }

  numStr := ia.GetNumStr()

  expectedNumStr := "450123647,1234"

  if expectedNumStr != numStr {
    t.Errorf("Error: Expected French Decimal separated NumStr= '%v'. Instead received NumStr= '%v'", expectedNumStr, numStr)
  }

}

func TestIntAry_SetEqualArrayLengths_01(t *testing.T) {
  nStr1 := "3536.123456"
  eNStr1 := "3536.123456"
  ePrecision1 := 6
  eSignVal1 := 1

  nStr2 := "12.14"
  eNStr2 := "0012.140000"
  ePrecision2 := 6
  eSignVal2 := 1

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(nStr2)
  ia1.SetEqualArrayLengths(&ia2)

  if ia1.GetNumStr() != eNStr1 {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
  }

  if ia1.GetSign() != eSignVal1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
  }

  if ia1.GetPrecision() != ePrecision1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
  }

  if ia2.GetNumStr() != eNStr2 {
    t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
  }

  if ia2.GetSign() != eSignVal2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
  }

  if ia2.GetPrecision() != ePrecision2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
  }

}

func TestIntAry_SetEqualArrayLengths_02(t *testing.T) {
  nStr1 := "12.14"
  eNStr1 := "0012.140000"
  ePrecision1 := 6
  eSignVal1 := 1

  nStr2 := "3536.123456"
  eNStr2 := "3536.123456"
  ePrecision2 := 6
  eSignVal2 := 1

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(nStr2)
  ia1.SetEqualArrayLengths(&ia2)

  if ia1.GetNumStr() != eNStr1 {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
  }

  if ia1.GetSign() != eSignVal1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
  }

  if ia1.GetPrecision() != ePrecision1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
  }

  if ia2.GetNumStr() != eNStr2 {
    t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
  }

  if ia2.GetSign() != eSignVal2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
  }

  if ia2.GetPrecision() != ePrecision2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
  }

}

func TestIntAry_SetEqualArrayLengths_03(t *testing.T) {
  nStr1 := "3536.123456"
  eNStr1 := "3536.123456"
  ePrecision1 := 6
  eSignVal1 := 1

  nStr2 := "-12.14"
  eNStr2 := "-0012.140000"
  ePrecision2 := 6
  eSignVal2 := -1

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(nStr2)
  ia1.SetEqualArrayLengths(&ia2)

  if ia1.GetNumStr() != eNStr1 {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
  }

  if ia1.GetSign() != eSignVal1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
  }

  if ia1.GetPrecision() != ePrecision1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
  }

  if ia2.GetNumStr() != eNStr2 {
    t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
  }

  if ia2.GetSign() != eSignVal2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
  }

  if ia2.GetPrecision() != ePrecision2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
  }

}

func TestIntAry_SetEqualArrayLengths_04(t *testing.T) {
  nStr1 := "-12.14"
  eNStr1 := "-0012.140000"
  ePrecision1 := 6
  eSignVal1 := -1

  nStr2 := "3536.123456"
  eNStr2 := "3536.123456"
  ePrecision2 := 6
  eSignVal2 := 1

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(nStr2)
  ia1.SetEqualArrayLengths(&ia2)

  if ia1.GetNumStr() != eNStr1 {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
  }

  if ia1.GetSign() != eSignVal1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
  }

  if ia1.GetPrecision() != ePrecision1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
  }

  if ia2.GetNumStr() != eNStr2 {
    t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
  }

  if ia2.GetSign() != eSignVal2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
  }

  if ia2.GetPrecision() != ePrecision2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
  }

}

func TestIntAry_SetEqualArrayLengths_05(t *testing.T) {
  nStr1 := "-123456.143456"
  eNStr1 := "-123456.143456"
  ePrecision1 := 6
  eSignVal1 := -1

  nStr2 := "353678.123456"
  eNStr2 := "353678.123456"
  ePrecision2 := 6
  eSignVal2 := 1

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(nStr2)
  ia1.SetEqualArrayLengths(&ia2)

  if ia1.GetNumStr() != eNStr1 {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
  }

  if ia1.GetSign() != eSignVal1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
  }

  if ia1.GetPrecision() != ePrecision1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
  }

  if ia2.GetNumStr() != eNStr2 {
    t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
  }

  if ia2.GetSign() != eSignVal2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
  }

  if ia2.GetPrecision() != ePrecision2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
  }

}

func TestIntAry_SetEqualArrayLengths_06(t *testing.T) {
  nStr1 := "0.00"
  eNStr1 := "0.00"
  ePrecision1 := 2
  eSignVal1 := 1

  nStr2 := "0"
  eNStr2 := "0.00"
  ePrecision2 := 2
  eSignVal2 := 1

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(nStr2)
  ia1.SetEqualArrayLengths(&ia2)

  if ia1.GetNumStr() != eNStr1 {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
  }

  if ia1.GetSign() != eSignVal1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
  }

  if ia1.GetPrecision() != ePrecision1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
  }

  if ia2.GetNumStr() != eNStr2 {
    t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
  }

  if ia2.GetSign() != eSignVal2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
  }

  if ia2.GetPrecision() != ePrecision2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
  }

}

func TestIntAry_SetEqualArrayLengths_07(t *testing.T) {
  nStr1 := "0"
  eNStr1 := "0.00"
  ePrecision1 := 2
  eSignVal1 := 1

  nStr2 := "0.00"
  eNStr2 := "0.00"
  ePrecision2 := 2
  eSignVal2 := 1

  ia1 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2 := IntAry{}.New()
  ia2.SetIntAryWithNumStr(nStr2)
  ia1.SetEqualArrayLengths(&ia2)

  if ia1.GetNumStr() != eNStr1 {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
  }

  if ia1.GetSign() != eSignVal1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
  }

  if ia1.GetPrecision() != ePrecision1 {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
  }

  if ia2.GetNumStr() != eNStr2 {
    t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
  }

  if ia2.GetSign() != eSignVal2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
  }

  if ia2.GetPrecision() != ePrecision2 {
    t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
  }

}

func TestIntAry_SetIntAryWithFloat_01(t *testing.T) {
  num := float32(475.895)
  expectedNumStr := "475.895"
  precision := 3
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloat32(num, precision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloat32(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloat_02(t *testing.T) {
  num := float32(475.895)
  expectedNumStr := "475.8950"
  precision := 4
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloat32(num, precision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloat32(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloat_03(t *testing.T) {
  num := float32(475.895600)
  expectedNumStr := "475.8956"
  precision := 4
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloat32(num, -1)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloat32(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloat_04(t *testing.T) {
  num := float32(-475.8956)
  expectedNumStr := "-475.8956"
  precision := 4
  signVal := -1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloat32(num, -1)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloat32(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloat64_01(t *testing.T) {
  num := float64(4756625.895)
  expectedNumStr := "4756625.895"
  precision := 3
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloat64(num, precision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloat64(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloat64_02(t *testing.T) {
  num := float64(4756625.895)
  expectedNumStr := "4756625.8950"
  precision := 4
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloat64(num, precision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloat64(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloat64_03(t *testing.T) {
  num := float64(475.895600)
  expectedNumStr := "475.8956"
  precision := 4
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloat64(num, -1)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloat64(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloat64_04(t *testing.T) {
  num := float64(-475.8956)
  expectedNumStr := "-475.8956"
  precision := 4
  signVal := -1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloat64(num, -1)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloat64(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloatBig_01(t *testing.T) {
  num := big.NewFloat(float64(4756625.895))
  expectedNumStr := "4756625.895"
  precision := 3
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloatBig(num, precision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v Error= %v", num, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloatBig_02(t *testing.T) {
  num := big.NewFloat(float64(4756625.8950))
  expectedNumStr := "4756625.895"
  precision := 3
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloatBig(num, precision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v  Error= %v", num, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloatBig_03(t *testing.T) {
  num := big.NewFloat(float64(475.895600))
  expectedNumStr := "475.8956"
  precision := 4
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloatBig(num, 4)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v  Error= %v", num.String(), err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloatBig_04(t *testing.T) {

  num := big.NewFloat(float64(-475.8956))
  expectedNumStr := "-475.8956"
  precision := 4
  signVal := -1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloatBig(num, precision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v Error= %v", num.String(), err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloatBig_05(t *testing.T) {
  num := big.NewFloat(float64(475.895600))
  expectedNumStr := "475.8956"
  precision := 4
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloatBig(num, -1)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v  Error= %v", num.String(), err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithFloatBig_06(t *testing.T) {
  num := big.NewFloat(float64(4756625.8957))
  expectedNumStr := "4756625.896"
  precision := 3
  signVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithFloatBig(num, precision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v Error= %v", num, err)
  }

  if expectedNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
  }

  if precision != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' .  Instead, ia.GetPrecisionInt()= '%v' .", precision, ia.GetPrecision())
  }

  if signVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
  }
}

func TestIntAry_SetIntAryWithBigInt_01(t *testing.T) {

  num := big.NewInt(int64(123456789))

  eNumStr := "123456.789"
  ePrecision := 3
  eSignVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithBigInt(num, ePrecision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithBigInt(num, ePrecision). num= %v  ePrecision= %v  eSignVal= %v", num.String(), ePrecision, eSignVal)
  }

  if eNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
  }

  if int(ePrecision) != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' . Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
  }

}

func TestIntAry_SetIntAryWithBigInt_02(t *testing.T) {

  num := big.NewInt(int64(-123456789))

  eNumStr := "-123456.789"
  ePrecision := 3
  eSignVal := -1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithBigInt(num, ePrecision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithBigInt(num, ePrecision). num= %v  ePrecision= %v  .", num.String(), ePrecision)
  }

  if eNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
  }

  if int(ePrecision) != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' . Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
  }

}

func TestIntAry_SetIntAryWithBigInt_03(t *testing.T) {

  num := big.NewInt(int64(0))

  eNumStr := "0.000"
  ePrecision := 3
  eSignVal := 1

  ia := IntAry{}.New()

  err := ia.SetIntAryWithBigInt(num, ePrecision)

  if err != nil {
    t.Errorf("Error returned from ia.SetIntAryWithBigInt(num, ePrecision). num= %v  ePrecision= %v  .", num.String(), ePrecision)
  }

  if eNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
  }

  if int(ePrecision) != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' . Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
  }

}

func TestIntAry_SetIntAryWithInt_01(t *testing.T) {

  num := 123456789

  eNumStr := "123456.789"
  ePrecision := uint(3)
  eSignVal := 1

  ia := IntAry{}.New()

  ia.SetIntAryWithInt(num, ePrecision)

  if eNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
  }

  if ePrecision != ia.GetPrecisionUint() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' . Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
  }

}

func TestIntAry_SetIntAryWithInt_02(t *testing.T) {

  num := -123456789

  eNumStr := "-12345.6789"
  ePrecision := uint(4)
  eSignVal := -1

  ia := IntAry{}.New()

  ia.SetIntAryWithInt(num, ePrecision)

  if eNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
  }

  if ePrecision != ia.GetPrecisionUint() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' . Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
  }

}

func TestIntAry_SetIntAryWithInt_03(t *testing.T) {

  num := 0

  eNumStr := "0.0000"
  ePrecision := uint(4)
  eSignVal := 1

  ia := IntAry{}.New()

  ia.SetIntAryWithInt(num, ePrecision)

  if eNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
  }

  if int(ePrecision) != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' . Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
  }

}

func TestIntAry_SetIntAryWithInt_04(t *testing.T) {

  num := 32

  eNumStr := "0.0032"
  ePrecision := uint(4)
  eSignVal := 1

  ia := IntAry{}.New()

  ia.SetIntAryWithInt(num, ePrecision)

  if eNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
  }

  if int(ePrecision) != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' . Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
  }

}

func TestIntAry_SetIntAryWithInt_05(t *testing.T) {

  num := -32

  eNumStr := "-32"
  ePrecision := uint(0)
  eSignVal := -1

  ia := IntAry{}.New()

  ia.SetIntAryWithInt(num, ePrecision)

  if eNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
  }

  if int(ePrecision) != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' . Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
  }

}

func TestIntAry_SetIntAryWithInt_06(t *testing.T) {

  num := 32

  eNumStr := "0.32"
  ePrecision := uint(2)
  eSignVal := 1

  ia := IntAry{}.New()

  ia.SetIntAryWithInt(num, ePrecision)

  if eNumStr != ia.GetNumStr() {
    t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
  }

  if int(ePrecision) != ia.GetPrecision() {
    t.Errorf("Expected ia.GetPrecisionInt()= '%v' . Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
  }

  if eSignVal != ia.GetSign() {
    t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
  }

}
