package mathops

import (
  "math/big"
  "strconv"
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

  ePrefix := "TestIntAry_SetAbsoluteValueThis_01"

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  originalNumberStr := "-987.652"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "987.652"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

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

  err = intAry.SetAbsoluteValueThis()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.SetAbsoluteValueThis()n"+
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
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final 'absolute' value\n"+
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

func TestIntAry_SetAbsoluteValueThis_02(t *testing.T) {

  ePrefix := "TestIntAry_SetAbsoluteValueThis_02"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "987.652"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "987.652"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

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

  err = intAry.SetAbsoluteValueThis()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.SetAbsoluteValueThis()n"+
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
      "Validation after 'rounding'\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry set to final 'absolute' value\n"+
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

func TestIntAry_SetCurrencySymbol_01(t *testing.T) {

  ePrefix := "TestIntAry_SetAbsoluteValueThis_02"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  originalNumberStr := "50.37"

  poundCurrencySymbol := '\U000000a3'

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

  err = intAry.SetCurrencySymbol(poundCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.SetCurrencySymbol(poundCurrencySymbol)\n"+
      "intAry= '%v'\n"+
      "poundCurrencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      string(poundCurrencySymbol),
      err.Error())

    return
  }

  actualCurrencySymbol := intAry.GetCurrencySymbol()

  if poundCurrencySymbol != actualCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because poundCurrencySymbol != actualCurrencySymbol\n"+
      "Expected actualCurrencySymbol = '%v'\n"+
      "  Actual actualCurrencySymbol = '%v'\n\n",
      ePrefix, string(poundCurrencySymbol), string(actualCurrencySymbol))

    return
  }

  return
}

func TestIntAry_SetDecimalSeparator_01(t *testing.T) {

  ePrefix := "TestIntAry_SetDecimalSeparator_01"

  originalNumberStr := "450 123 647,1234"

  expectedNumberStr := "450123647,1234"

  var frenchDecSeparator, frenchThousandsSeparator,
    frenchCurrencySymbol rune

  frenchDecSeparator = ','

  frenchThousandsSeparator = ' '

  frenchCurrencySymbol = '€'

  intAry := new(IntAry).New()

  err := intAry.SetDecimalSeparator(frenchDecSeparator)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetDecimalSeparator(frenchDecSeparator)\n"+
      "frenchDecSeparator= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, string(frenchDecSeparator), err.Error())
    return
  }

  err = intAry.SetThousandsSeparator(frenchThousandsSeparator)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.SetThousandsSeparator(frenchThousandsSeparator)\n"+
      "frenchThousandsSeparator= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, string(frenchThousandsSeparator), err.Error())
    return
  }

  err = intAry.SetCurrencySymbol(frenchCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.SetCurrencySymbol(frenchCurrencySymbol)\n"+
      "frenchCurrencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, string(frenchCurrencySymbol), err.Error())
    return
  }

  err = intAry.SetIntAryWithNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
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

  actualDecimalSeparator := intAry.GetDecimalSeparator()

  if frenchDecSeparator != actualDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Decimal Separators DO NOT MATCH!\n"+
      "Because frenchDecSeparator != actualDecimalSeparator\n"+
      "Expected actualDecimalSeparator = '%v'\n"+
      "  Actual actualDecimalSeparator = '%v'\n\n",
      ePrefix, string(frenchDecSeparator), string(actualDecimalSeparator))

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

  return
}

func TestIntAry_SetEqualArrayLengths_01(t *testing.T) {

  ePrefix := "TestIntAry_SetEqualArrayLengths_01"

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  originalNumberStr1 := "3536.123456"

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr1 := "3536.123456"

  expectedPrecisionInt1 := 6

  expectedPrecisionUint1 := uint(expectedPrecisionInt1)

  expectedSignValue1 := 1

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr2 := "12.14"

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr2 := "0012.140000"

  expectedPrecisionInt2 := 6

  expectedPrecisionUint2 := uint(expectedPrecisionInt2)

  expectedSignValue2 := 1

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr1)\n"+
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

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr2)\n"+
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

  intAry2PrecisionInt := intAry2.GetPrecision()

  intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2PrecisionUint, err :=\n"+
      "  intAry2.GetPrecisionUint()\n"+
      "intAry2Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  intAry2SignValue, err := intAry2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2SignValue, err := intAry2.GetSign()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  if expectedNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr2, intAry2NumberStr)

    return
  }

  if expectedPrecisionInt2 != intAry2PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt2 != intAry2PrecisionInt\n"+
      "Expected intAry2PrecisionInt = '%v'\n"+
      "  Actual intAry2PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt2, intAry2PrecisionInt)

    return
  }

  if expectedPrecisionUint2 != intAry2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint2 != intAry2PrecisionUint\n"+
      "Expected intAry2PrecisionUint = '%v'\n"+
      "  Actual intAry2PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint2, intAry2PrecisionUint)

    return
  }

  if expectedSignValue2 != intAry2SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue2 != intAry2SignValue\n"+
      "Expected intAry2SignValue = '%v'\n"+
      "  Actual intAry2SignValue = '%v'\n\n",
      ePrefix, expectedSignValue2, intAry2SignValue)

    return
  }

  err = intAry1.SetEqualArrayLengths(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SetEqualArrayLengths(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating final intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err = intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intArySignValue1, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue1, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1PrecisionInt := intAry1.GetPrecision()

  intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1PrecisionUint, err :=\n"+
      "  intAry1.GetPrecisionUint()\n"+
      "intAry1Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr1 != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr1, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt1 != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt1 != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt1, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint1 != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint1 != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint1, intAry1PrecisionUint)

    return
  }

  if expectedSignValue1 != intArySignValue1 {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue1 != intArySignValue1\n"+
      "Expected intArySignValue1 = '%v'\n"+
      "  Actual intArySignValue1 = '%v'\n\n",
      ePrefix, expectedSignValue1, intArySignValue1)

    return
  }

  return
}

func TestIntAry_SetEqualArrayLengths_02(t *testing.T) {

  ePrefix := "TestIntAry_SetEqualArrayLengths_02"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr1 := "12.14"

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr1 := "0012.140000"

  expectedPrecisionInt1 := 6

  expectedPrecisionUint1 := uint(expectedPrecisionInt1)

  expectedSignValue1 := 1

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  originalNumberStr2 := "3536.123456"

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr2 := "3536.123456"

  expectedPrecisionInt2 := 6

  expectedPrecisionUint2 := uint(expectedPrecisionInt2)

  expectedSignValue2 := 1

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr1)\n"+
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

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr2)\n"+
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

  intAry2PrecisionInt := intAry2.GetPrecision()

  intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2PrecisionUint, err :=\n"+
      "  intAry2.GetPrecisionUint()\n"+
      "intAry2Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  intAry2SignValue, err := intAry2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2SignValue, err := intAry2.GetSign()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  if expectedNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr2, intAry2NumberStr)

    return
  }

  if expectedPrecisionInt2 != intAry2PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt2 != intAry2PrecisionInt\n"+
      "Expected intAry2PrecisionInt = '%v'\n"+
      "  Actual intAry2PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt2, intAry2PrecisionInt)

    return
  }

  if expectedPrecisionUint2 != intAry2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint2 != intAry2PrecisionUint\n"+
      "Expected intAry2PrecisionUint = '%v'\n"+
      "  Actual intAry2PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint2, intAry2PrecisionUint)

    return
  }

  if expectedSignValue2 != intAry2SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue2 != intAry2SignValue\n"+
      "Expected intAry2SignValue = '%v'\n"+
      "  Actual intAry2SignValue = '%v'\n\n",
      ePrefix, expectedSignValue2, intAry2SignValue)

    return
  }

  err = intAry1.SetEqualArrayLengths(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SetEqualArrayLengths(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating final intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err = intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intArySignValue1, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue1, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1PrecisionInt := intAry1.GetPrecision()

  intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1PrecisionUint, err :=\n"+
      "  intAry1.GetPrecisionUint()\n"+
      "intAry1Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr1 != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr1, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt1 != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt1 != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt1, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint1 != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint1 != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint1, intAry1PrecisionUint)

    return
  }

  if expectedSignValue1 != intArySignValue1 {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue1 != intArySignValue1\n"+
      "Expected intArySignValue1 = '%v'\n"+
      "  Actual intArySignValue1 = '%v'\n\n",
      ePrefix, expectedSignValue1, intArySignValue1)

    return
  }

  return
}

func TestIntAry_SetEqualArrayLengths_03(t *testing.T) {

  ePrefix := "TestIntAry_SetEqualArrayLengths_03"

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  originalNumberStr1 := "3536.123456"

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr1 := "3536.123456"

  expectedPrecisionInt1 := 6

  expectedPrecisionUint1 := uint(expectedPrecisionInt1)

  expectedSignValue1 := 1

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalNumberStr2 := "-12.14"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr2 := "-0012.140000"

  expectedPrecisionInt2 := 6

  expectedPrecisionUint2 := uint(expectedPrecisionInt2)

  expectedSignValue2 := -1

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr1)\n"+
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

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr2)\n"+
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

  intAry2PrecisionInt := intAry2.GetPrecision()

  intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2PrecisionUint, err :=\n"+
      "  intAry2.GetPrecisionUint()\n"+
      "intAry2Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  intAry2SignValue, err := intAry2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2SignValue, err := intAry2.GetSign()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  if expectedNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr2, intAry2NumberStr)

    return
  }

  if expectedPrecisionInt2 != intAry2PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt2 != intAry2PrecisionInt\n"+
      "Expected intAry2PrecisionInt = '%v'\n"+
      "  Actual intAry2PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt2, intAry2PrecisionInt)

    return
  }

  if expectedPrecisionUint2 != intAry2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint2 != intAry2PrecisionUint\n"+
      "Expected intAry2PrecisionUint = '%v'\n"+
      "  Actual intAry2PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint2, intAry2PrecisionUint)

    return
  }

  if expectedSignValue2 != intAry2SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue2 != intAry2SignValue\n"+
      "Expected intAry2SignValue = '%v'\n"+
      "  Actual intAry2SignValue = '%v'\n\n",
      ePrefix, expectedSignValue2, intAry2SignValue)

    return
  }

  err = intAry1.SetEqualArrayLengths(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SetEqualArrayLengths(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating final intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err = intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intArySignValue1, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue1, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1PrecisionInt := intAry1.GetPrecision()

  intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1PrecisionUint, err :=\n"+
      "  intAry1.GetPrecisionUint()\n"+
      "intAry1Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr1 != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr1, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt1 != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt1 != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt1, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint1 != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint1 != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint1, intAry1PrecisionUint)

    return
  }

  if expectedSignValue1 != intArySignValue1 {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue1 != intArySignValue1\n"+
      "Expected intArySignValue1 = '%v'\n"+
      "  Actual intArySignValue1 = '%v'\n\n",
      ePrefix, expectedSignValue1, intArySignValue1)

    return
  }

  return
}

func TestIntAry_SetEqualArrayLengths_04(t *testing.T) {

  ePrefix := "TestIntAry_SetEqualArrayLengths_04"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalNumberStr1 := "-12.14"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr1 := "-0012.140000"

  expectedPrecisionInt1 := 6

  expectedPrecisionUint1 := uint(expectedPrecisionInt1)

  expectedSignValue1 := -1

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  originalNumberStr2 := "3536.123456"

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr2 := "3536.123456"

  expectedPrecisionInt2 := 6

  expectedPrecisionUint2 := uint(expectedPrecisionInt2)

  expectedSignValue2 := 1

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr1)\n"+
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

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr2)\n"+
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

  intAry2PrecisionInt := intAry2.GetPrecision()

  intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2PrecisionUint, err :=\n"+
      "  intAry2.GetPrecisionUint()\n"+
      "intAry2Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  intAry2SignValue, err := intAry2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2SignValue, err := intAry2.GetSign()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  if expectedNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr2, intAry2NumberStr)

    return
  }

  if expectedPrecisionInt2 != intAry2PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt2 != intAry2PrecisionInt\n"+
      "Expected intAry2PrecisionInt = '%v'\n"+
      "  Actual intAry2PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt2, intAry2PrecisionInt)

    return
  }

  if expectedPrecisionUint2 != intAry2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint2 != intAry2PrecisionUint\n"+
      "Expected intAry2PrecisionUint = '%v'\n"+
      "  Actual intAry2PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint2, intAry2PrecisionUint)

    return
  }

  if expectedSignValue2 != intAry2SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue2 != intAry2SignValue\n"+
      "Expected intAry2SignValue = '%v'\n"+
      "  Actual intAry2SignValue = '%v'\n\n",
      ePrefix, expectedSignValue2, intAry2SignValue)

    return
  }

  err = intAry1.SetEqualArrayLengths(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SetEqualArrayLengths(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating final intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err = intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intArySignValue1, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue1, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1PrecisionInt := intAry1.GetPrecision()

  intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1PrecisionUint, err :=\n"+
      "  intAry1.GetPrecisionUint()\n"+
      "intAry1Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr1 != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr1, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt1 != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt1 != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt1, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint1 != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint1 != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint1, intAry1PrecisionUint)

    return
  }

  if expectedSignValue1 != intArySignValue1 {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue1 != intArySignValue1\n"+
      "Expected intArySignValue1 = '%v'\n"+
      "  Actual intArySignValue1 = '%v'\n\n",
      ePrefix, expectedSignValue1, intArySignValue1)

    return
  }

  return
}

func TestIntAry_SetEqualArrayLengths_05(t *testing.T) {

  ePrefix := "TestIntAry_SetEqualArrayLengths_05"

  //                                      1         2         3
  //                           0.1234567890123456789012345678901234567
  originalNumberStr1 := "-123456.143456"

  //                                      1         2         3
  //                           0.1234567890123456789012345678901234567
  expectedNumberStr1 := "-123456.143456"

  expectedPrecisionInt1 := 6

  expectedPrecisionUint1 := uint(expectedPrecisionInt1)

  expectedSignValue1 := -1

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  originalNumberStr2 := "353678.123456"

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  expectedNumberStr2 := "353678.123456"

  expectedPrecisionInt2 := 6

  expectedPrecisionUint2 := uint(expectedPrecisionInt2)

  expectedSignValue2 := 1

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr1)\n"+
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

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr2)\n"+
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

  intAry2PrecisionInt := intAry2.GetPrecision()

  intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2PrecisionUint, err :=\n"+
      "  intAry2.GetPrecisionUint()\n"+
      "intAry2Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  intAry2SignValue, err := intAry2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2SignValue, err := intAry2.GetSign()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  if expectedNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr2, intAry2NumberStr)

    return
  }

  if expectedPrecisionInt2 != intAry2PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt2 != intAry2PrecisionInt\n"+
      "Expected intAry2PrecisionInt = '%v'\n"+
      "  Actual intAry2PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt2, intAry2PrecisionInt)

    return
  }

  if expectedPrecisionUint2 != intAry2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint2 != intAry2PrecisionUint\n"+
      "Expected intAry2PrecisionUint = '%v'\n"+
      "  Actual intAry2PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint2, intAry2PrecisionUint)

    return
  }

  if expectedSignValue2 != intAry2SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue2 != intAry2SignValue\n"+
      "Expected intAry2SignValue = '%v'\n"+
      "  Actual intAry2SignValue = '%v'\n\n",
      ePrefix, expectedSignValue2, intAry2SignValue)

    return
  }

  err = intAry1.SetEqualArrayLengths(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SetEqualArrayLengths(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating final intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err = intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intArySignValue1, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue1, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1PrecisionInt := intAry1.GetPrecision()

  intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1PrecisionUint, err :=\n"+
      "  intAry1.GetPrecisionUint()\n"+
      "intAry1Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr1 != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr1, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt1 != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt1 != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt1, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint1 != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint1 != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint1, intAry1PrecisionUint)

    return
  }

  if expectedSignValue1 != intArySignValue1 {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue1 != intArySignValue1\n"+
      "Expected intArySignValue1 = '%v'\n"+
      "  Actual intArySignValue1 = '%v'\n\n",
      ePrefix, expectedSignValue1, intArySignValue1)

    return
  }

  return
}

func TestIntAry_SetEqualArrayLengths_06(t *testing.T) {

  ePrefix := "TestIntAry_SetEqualArrayLengths_06"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  originalNumberStr1 := "0.00"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr1 := "0.00"

  expectedPrecisionInt1 := 2

  expectedPrecisionUint1 := uint(expectedPrecisionInt1)

  expectedSignValue1 := 1

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  originalNumberStr2 := "0"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr2 := "0.00"

  expectedPrecisionInt2 := 2

  expectedPrecisionUint2 := uint(expectedPrecisionInt2)

  expectedSignValue2 := 1

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr1)\n"+
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

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr2)\n"+
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

  intAry2PrecisionInt := intAry2.GetPrecision()

  intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2PrecisionUint, err :=\n"+
      "  intAry2.GetPrecisionUint()\n"+
      "intAry2Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  intAry2SignValue, err := intAry2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2SignValue, err := intAry2.GetSign()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  if expectedNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr2, intAry2NumberStr)

    return
  }

  if expectedPrecisionInt2 != intAry2PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt2 != intAry2PrecisionInt\n"+
      "Expected intAry2PrecisionInt = '%v'\n"+
      "  Actual intAry2PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt2, intAry2PrecisionInt)

    return
  }

  if expectedPrecisionUint2 != intAry2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint2 != intAry2PrecisionUint\n"+
      "Expected intAry2PrecisionUint = '%v'\n"+
      "  Actual intAry2PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint2, intAry2PrecisionUint)

    return
  }

  if expectedSignValue2 != intAry2SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue2 != intAry2SignValue\n"+
      "Expected intAry2SignValue = '%v'\n"+
      "  Actual intAry2SignValue = '%v'\n\n",
      ePrefix, expectedSignValue2, intAry2SignValue)

    return
  }

  err = intAry1.SetEqualArrayLengths(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SetEqualArrayLengths(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating final intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err = intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intArySignValue1, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue1, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1PrecisionInt := intAry1.GetPrecision()

  intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1PrecisionUint, err :=\n"+
      "  intAry1.GetPrecisionUint()\n"+
      "intAry1Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr1 != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr1, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt1 != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt1 != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt1, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint1 != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint1 != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint1, intAry1PrecisionUint)

    return
  }

  if expectedSignValue1 != intArySignValue1 {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue1 != intArySignValue1\n"+
      "Expected intArySignValue1 = '%v'\n"+
      "  Actual intArySignValue1 = '%v'\n\n",
      ePrefix, expectedSignValue1, intArySignValue1)

    return
  }

  return
}

func TestIntAry_SetEqualArrayLengths_07(t *testing.T) {

  ePrefix := "TestIntAry_SetEqualArrayLengths_06"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  originalNumberStr1 := "0"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr1 := "0.00"

  expectedPrecisionInt1 := 2

  expectedPrecisionUint1 := uint(expectedPrecisionInt1)

  expectedSignValue1 := 1

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  originalNumberStr2 := "0.00"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr2 := "0.00"

  expectedPrecisionInt2 := 2

  expectedPrecisionUint2 := uint(expectedPrecisionInt2)

  expectedSignValue2 := 1

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr1)\n"+
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

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(\n"+
      "  originalNumberStr2)\n"+
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

  intAry2PrecisionInt := intAry2.GetPrecision()

  intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2PrecisionUint, err :=\n"+
      "  intAry2.GetPrecisionUint()\n"+
      "intAry2Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  intAry2SignValue, err := intAry2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2SignValue, err := intAry2.GetSign()\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry2NumberStr, err.Error())
    return
  }

  if expectedNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr2, intAry2NumberStr)

    return
  }

  if expectedPrecisionInt2 != intAry2PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt2 != intAry2PrecisionInt\n"+
      "Expected intAry2PrecisionInt = '%v'\n"+
      "  Actual intAry2PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt2, intAry2PrecisionInt)

    return
  }

  if expectedPrecisionUint2 != intAry2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint2 != intAry2PrecisionUint\n"+
      "Expected intAry2PrecisionUint = '%v'\n"+
      "  Actual intAry2PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint2, intAry2PrecisionUint)

    return
  }

  if expectedSignValue2 != intAry2SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue2 != intAry2SignValue\n"+
      "Expected intAry2SignValue = '%v'\n"+
      "  Actual intAry2SignValue = '%v'\n\n",
      ePrefix, expectedSignValue2, intAry2SignValue)

    return
  }

  err = intAry1.SetEqualArrayLengths(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SetEqualArrayLengths(&intAry2)\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAry1NumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAry1.IsValid("Validating final intAry1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.IsValid('Validating final intAry1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry1NumberStr, err = intAry1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumberStr, err := intAry1.GetNumStr()\n"+
      "intAry1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intArySignValue1, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue1, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1PrecisionInt := intAry1.GetPrecision()

  intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1PrecisionUint, err :=\n"+
      "  intAry1.GetPrecisionUint()\n"+
      "intAry1Result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr1 != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr1 != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr1, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt1 != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt1 != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt1, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint1 != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint1 != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint1, intAry1PrecisionUint)

    return
  }

  if expectedSignValue1 != intArySignValue1 {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue1 != intArySignValue1\n"+
      "Expected intArySignValue1 = '%v'\n"+
      "  Actual intArySignValue1 = '%v'\n\n",
      ePrefix, expectedSignValue1, intArySignValue1)

    return
  }

  return
}

func TestIntAry_SetIntAryWithFloat32_01(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloat32_01"

  //                                         1         2         3
  //                              0.1234567890123456789012345678901234567
  originalFloat32Num := float32(475.895)

  originalRawPrecisionInt := 3

  originalTargetPrecisionInt := 3

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "475.895"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloat32(originalFloat32Num, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry.SetIntAryWithFloat32(\n"+
      "  originalFloat32Num, originalTargetPrecisionInt)\n"+
      "originalFloat32Num= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(float64(originalFloat32Num), 'f', originalRawPrecisionInt, 32),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloat32_02(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloat32_02"

  //                                         1         2         3
  //                              0.1234567890123456789012345678901234567
  originalFloat32Num := float32(475.895)

  originalRawPrecisionInt := 3

  originalTargetPrecisionInt := 4

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "475.8950"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloat32(originalFloat32Num, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry.SetIntAryWithFloat32(\n"+
      "  originalFloat32Num, originalTargetPrecisionInt)\n"+
      "originalFloat32Num= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(float64(originalFloat32Num), 'f', originalRawPrecisionInt, 32),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloat32_03(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloat32_03"

  //                                         1         2         3
  //                              0.1234567890123456789012345678901234567
  originalFloat32Num := float32(475.895600)

  originalRawPrecisionInt := 6

  originalTargetPrecisionInt := 4

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "475.8956"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloat32(originalFloat32Num, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry.SetIntAryWithFloat32(\n"+
      "  originalFloat32Num, originalTargetPrecisionInt)\n"+
      "originalFloat32Num= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(float64(originalFloat32Num), 'f', originalRawPrecisionInt, 32),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloat32_04(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloat32_04"

  //                                          1         2         3
  //                               0.1234567890123456789012345678901234567
  originalFloat32Num := float32(-475.8956)

  originalRawPrecisionInt := 4

  originalTargetPrecisionInt := -1

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-475.8956"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloat32(originalFloat32Num, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry.SetIntAryWithFloat32(\n"+
      "  originalFloat32Num, originalTargetPrecisionInt)\n"+
      "originalFloat32Num= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(float64(originalFloat32Num), 'f', originalRawPrecisionInt, 32),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloat64_01(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloat64_01"

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  originalFloat64Num := 4756625.895

  originalRawPrecisionInt := 3

  originalTargetPrecisionInt := 3

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  expectedNumberStr := "4756625.895"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloat64(originalFloat64Num, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloat64(\n"+
      "  originalFloat64Num, originalTargetPrecisionInt)\n"+
      "originalFloat64Num= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(originalFloat64Num, 'f', originalRawPrecisionInt, 64),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloat64_02(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloat64_02"

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  originalFloat64Num := 4756625.895

  originalRawPrecisionInt := 3

  originalTargetPrecisionInt := 4

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  expectedNumberStr := "4756625.8950"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloat64(originalFloat64Num, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloat64(\n"+
      "  originalFloat64Num, originalTargetPrecisionInt)\n"+
      "originalFloat64Num= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(originalFloat64Num, 'f', originalRawPrecisionInt, 64),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloat64_03(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloat64_03"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalFloat64Num := 475.895600

  originalRawPrecisionInt := 6

  originalTargetPrecisionInt := -1

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "475.8956"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloat64(originalFloat64Num, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloat64(\n"+
      "  originalFloat64Num, originalTargetPrecisionInt)\n"+
      "originalFloat64Num= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(originalFloat64Num, 'f', originalRawPrecisionInt, 64),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloat64_04(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloat64_04"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalFloat64Num := -475.8956

  originalRawPrecisionInt := 4

  originalTargetPrecisionInt := -1

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-475.8956"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloat64(originalFloat64Num, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloat64(\n"+
      "  originalFloat64Num, originalTargetPrecisionInt)\n"+
      "originalFloat64Num= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(originalFloat64Num, 'f', originalRawPrecisionInt, 64),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloatBig_01(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloatBig_01"

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  originalFloat64Num := 4756625.895

  originalTargetPrecisionInt := 3

  originalBigFloatNum := big.NewFloat(originalFloat64Num)

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  expectedNumberStr := "4756625.895"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)\n"+
      "originalBigFloatNum= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBigFloatNum.Text('f', -1),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloatBig_02(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloatBig_02"

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  originalFloat64Num := 4756625.8950

  originalTargetPrecisionInt := 3

  originalBigFloatNum := big.NewFloat(originalFloat64Num)

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  expectedNumberStr := "4756625.895"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)\n"+
      "originalBigFloatNum= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBigFloatNum.Text('f', -1),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloatBig_03(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloatBig_03"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalFloat64Num := 475.895600

  originalTargetPrecisionInt := 4

  originalBigFloatNum := big.NewFloat(originalFloat64Num)

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "475.8956"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)\n"+
      "originalBigFloatNum= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBigFloatNum.Text('f', -1),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloatBig_04(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloatBig_04"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalFloat64Num := -475.8956

  originalTargetPrecisionInt := 4

  originalBigFloatNum := big.NewFloat(originalFloat64Num)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-475.8956"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)\n"+
      "originalBigFloatNum= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBigFloatNum.Text('f', -1),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloatBig_05(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloatBig_05"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalFloat64Num := 475.895600

  originalTargetPrecisionInt := 4

  originalBigFloatNum := big.NewFloat(originalFloat64Num)

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "475.8956"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)\n"+
      "originalBigFloatNum= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBigFloatNum.Text('f', -1),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithFloatBig_06(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithFloatBig_06"

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  originalFloat64Num := 4756625.8957

  originalBigFloatNum := big.NewFloat(originalFloat64Num)

  originalTargetPrecisionInt := 3

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  expectedNumberStr := "4756625.896"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithFloatBig(originalBigFloatNum, originalTargetPrecisionInt)\n"+
      "originalBigFloatNum= '%v'\n"+
      "originalTargetPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBigFloatNum.Text('f', -1),
      originalTargetPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithBigInt_01(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithBigInt_01"

  originalInt64Num := int64(123456789)

  originalBigIntNum := big.NewInt(originalInt64Num)

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr := "123456.789"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithBigInt(originalBigIntNum, expectedPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithBigInt(originalBigIntNum, expectedPrecisionInt)\n"+
      "originalBigIntNum= '%v'\n"+
      "expectedPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBigIntNum.Text(10),
      expectedPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithBigInt_02(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithBigInt_02"

  originalInt64Num := int64(-123456789)

  originalBigIntNum := big.NewInt(originalInt64Num)

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  expectedNumberStr := "-123456.789"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithBigInt(originalBigIntNum, expectedPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithBigInt(originalBigIntNum, expectedPrecisionInt)\n"+
      "originalBigIntNum= '%v'\n"+
      "expectedPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBigIntNum.Text(10),
      expectedPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithBigInt_03(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithBigInt_03"

  originalInt64Num := int64(0)

  originalBigIntNum := big.NewInt(originalInt64Num)

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.000"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithBigInt(originalBigIntNum, expectedPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithBigInt(originalBigIntNum, expectedPrecisionInt)\n"+
      "originalBigIntNum= '%v'\n"+
      "expectedPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalBigIntNum.Text(10),
      expectedPrecisionInt,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithInt_01(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithInt_01"

  originalNumInt := 123456789

  originalTargetPrecisionUint := uint(3)

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr := "123456.789"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)\n"+
      "originalNumInt= '%v'\n"+
      "originalTargetPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumInt,
      originalTargetPrecisionUint,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithInt_02(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithInt_02"

  originalNumInt := -123456789

  originalTargetPrecisionUint := uint(4)

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr := "-12345.6789"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)\n"+
      "originalNumInt= '%v'\n"+
      "originalTargetPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumInt,
      originalTargetPrecisionUint,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithInt_03(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithInt_03"

  originalNumInt := 0

  originalTargetPrecisionUint := uint(4)

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr := "0.0000"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)\n"+
      "originalNumInt= '%v'\n"+
      "originalTargetPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumInt,
      originalTargetPrecisionUint,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithInt_04(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithInt_04"

  originalNumInt := 32

  originalTargetPrecisionUint := uint(4)

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.0032"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)\n"+
      "originalNumInt= '%v'\n"+
      "originalTargetPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumInt,
      originalTargetPrecisionUint,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithInt_05(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithInt_05"

  originalNumInt := -32

  originalTargetPrecisionUint := uint(0)

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "-32"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)\n"+
      "originalNumInt= '%v'\n"+
      "originalTargetPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumInt,
      originalTargetPrecisionUint,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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

func TestIntAry_SetIntAryWithInt_06(t *testing.T) {

  ePrefix := "TestIntAry_SetIntAryWithInt_06"

  originalNumInt := 32

  originalTargetPrecisionUint := uint(2)

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.32"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

  err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithInt(originalNumInt, originalTargetPrecisionUint)\n"+
      "originalNumInt= '%v'\n"+
      "originalTargetPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumInt,
      originalTargetPrecisionUint,
      err.Error())

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

  intAryPrecisionInt := intAry.GetPrecision()

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err := intAry.GetPrecisionUint()\n"+
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

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
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
