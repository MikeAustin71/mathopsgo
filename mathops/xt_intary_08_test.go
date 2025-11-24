package mathops

import (
  "testing"
)

func TestIntAry_SetSignificantDigitIdxs_01(t *testing.T) {

  ePrefix := "TestIntAry_SetSignificantDigitIdxs_01"

  originalNumberStr := ".7770"

  expectedAryLen := 5

  expectedNumberStr := "0.7770"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntegerLen := 1

  expectedSignificantIntegerLen := 1

  expectedSignificantFractionLen := 3

  var expectedIsZeroValue, expectedIsIntegerZeroValue bool

  expectedIsZeroValue = false

  expectedIsIntegerZeroValue = true

  expectedFirstDigitIdx := 0

  expectedLastDigitIdx := 3

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

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
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

  intAryStats := intAry.GetIntAryStats()

  if expectedFirstDigitIdx != intAryStats.FirstDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array First Digits DO NOT MATCH!\n"+
      "Because expectedFirstDigitIdx != intAryStats.FirstDigitIdx\n"+
      "Expected intAryStats.FirstDigitIdx = '%v'\n"+
      "  Actual intAryStats.FirstDigitIdx = '%v'\n\n",
      ePrefix, expectedFirstDigitIdx, intAryStats.FirstDigitIdx)

    return
  }

  if expectedLastDigitIdx != intAryStats.LastDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array Last Digits DO NOT MATCH!\n"+
      "Because expectedLastDigitIdx != intAryStats.LastDigitIdx\n"+
      "Expected intAryStats.LastDigitIdx = '%v'\n"+
      "  Actual intAryStats.LastDigitIdx = '%v'\n\n",
      ePrefix, expectedLastDigitIdx, intAryStats.LastDigitIdx)

    return
  }

  if expectedIsZeroValue != intAryStats.IsZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsZeroValue != intAryStats.IsZeroValue\n"+
      "Expected intAryStats.IsZeroValue = '%v'\n"+
      "  Actual intAryStats.IsZeroValue = '%v'\n\n",
      ePrefix, expectedIsZeroValue, intAryStats.IsZeroValue)

    return
  }

  if expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsIntegerZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue\n"+
      "Expected intAryStats.IsIntegerZeroValue = '%v'\n"+
      "  Actual intAryStats.IsIntegerZeroValue = '%v'\n\n",
      ePrefix, expectedIsIntegerZeroValue, intAryStats.IsIntegerZeroValue)

    return
  }

  if expectedAryLen != intAryStats.IntAryLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedAryLen != intAryStats.IntAryLen\n"+
      "Expected intAryStats.IntAryLen = '%v'\n"+
      "  Actual intAryStats.IntAryLen = '%v'\n\n",
      ePrefix, expectedAryLen, intAryStats.IntAryLen)

    return
  }

  if expectedIntegerLen != intAryStats.IntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Integer Lengths ARE NOT EQUAL!\n"+
      "Because expectedIntegerLen != intAryStats.IntegerLen\n"+
      "Expected intAryStats.IntegerLen = '%v'\n"+
      "  Actual intAryStats.IntegerLen = '%v'\n\n",
      ePrefix, expectedIntegerLen, intAryStats.IntegerLen)

    return
  }

  if expectedSignificantFractionLen != intAryStats.SignificantFractionLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Fraction Lengths NOT EQUAL!\n"+
      "Because expectedSignificantFractionLen != intAryStats.SignificantFractionLen\n"+
      "Expected intAryStats.SignificantFractionLen = '%v'\n"+
      "  Actual intAryStats.SignificantFractionLen = '%v'\n\n",
      ePrefix, expectedSignificantFractionLen, intAryStats.SignificantFractionLen)

    return
  }

  if expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Integer Lengths NOT EQUAL!\n"+
      "Because expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen\n"+
      "Expected intAryStats.SignificantIntegerLen = '%v'\n"+
      "  Actual intAryStats.SignificantIntegerLen = '%v'\n\n",
      ePrefix, expectedSignificantIntegerLen, intAryStats.SignificantIntegerLen)

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

func TestIntAry_SetSignificantDigitIdxs_02(t *testing.T) {

  ePrefix := "TestIntAry_SetSignificantDigitIdxs_02"

  originalNumberStr := "000123456.123456000"

  expectedAryLen := 18

  expectedNumberStr := "000123456.123456000"

  expectedPrecisionInt := 9

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntegerLen := 9

  expectedSignificantIntegerLen := 6

  expectedSignificantFractionLen := 6

  var expectedIsZeroValue, expectedIsIntegerZeroValue bool

  expectedIsZeroValue = false

  expectedIsIntegerZeroValue = false

  expectedFirstDigitIdx := 3

  expectedLastDigitIdx := 14

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

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
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

  intAryStats := intAry.GetIntAryStats()

  if expectedFirstDigitIdx != intAryStats.FirstDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array First Digits DO NOT MATCH!\n"+
      "Because expectedFirstDigitIdx != intAryStats.FirstDigitIdx\n"+
      "Expected intAryStats.FirstDigitIdx = '%v'\n"+
      "  Actual intAryStats.FirstDigitIdx = '%v'\n\n",
      ePrefix, expectedFirstDigitIdx, intAryStats.FirstDigitIdx)

    return
  }

  if expectedLastDigitIdx != intAryStats.LastDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array Last Digits DO NOT MATCH!\n"+
      "Because expectedLastDigitIdx != intAryStats.LastDigitIdx\n"+
      "Expected intAryStats.LastDigitIdx = '%v'\n"+
      "  Actual intAryStats.LastDigitIdx = '%v'\n\n",
      ePrefix, expectedLastDigitIdx, intAryStats.LastDigitIdx)

    return
  }

  if expectedIsZeroValue != intAryStats.IsZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsZeroValue != intAryStats.IsZeroValue\n"+
      "Expected intAryStats.IsZeroValue = '%v'\n"+
      "  Actual intAryStats.IsZeroValue = '%v'\n\n",
      ePrefix, expectedIsZeroValue, intAryStats.IsZeroValue)

    return
  }

  if expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsIntegerZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue\n"+
      "Expected intAryStats.IsIntegerZeroValue = '%v'\n"+
      "  Actual intAryStats.IsIntegerZeroValue = '%v'\n\n",
      ePrefix, expectedIsIntegerZeroValue, intAryStats.IsIntegerZeroValue)

    return
  }

  if expectedAryLen != intAryStats.IntAryLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedAryLen != intAryStats.IntAryLen\n"+
      "Expected intAryStats.IntAryLen = '%v'\n"+
      "  Actual intAryStats.IntAryLen = '%v'\n\n",
      ePrefix, expectedAryLen, intAryStats.IntAryLen)

    return
  }

  if expectedIntegerLen != intAryStats.IntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Integer Lengths ARE NOT EQUAL!\n"+
      "Because expectedIntegerLen != intAryStats.IntegerLen\n"+
      "Expected intAryStats.IntegerLen = '%v'\n"+
      "  Actual intAryStats.IntegerLen = '%v'\n\n",
      ePrefix, expectedIntegerLen, intAryStats.IntegerLen)

    return
  }

  if expectedSignificantFractionLen != intAryStats.SignificantFractionLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Fraction Lengths NOT EQUAL!\n"+
      "Because expectedSignificantFractionLen != intAryStats.SignificantFractionLen\n"+
      "Expected intAryStats.SignificantFractionLen = '%v'\n"+
      "  Actual intAryStats.SignificantFractionLen = '%v'\n\n",
      ePrefix, expectedSignificantFractionLen, intAryStats.SignificantFractionLen)

    return
  }

  if expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Integer Lengths NOT EQUAL!\n"+
      "Because expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen\n"+
      "Expected intAryStats.SignificantIntegerLen = '%v'\n"+
      "  Actual intAryStats.SignificantIntegerLen = '%v'\n\n",
      ePrefix, expectedSignificantIntegerLen, intAryStats.SignificantIntegerLen)

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

func TestIntAry_SetSignificantDigitIdxs_03(t *testing.T) {

  ePrefix := "TestIntAry_SetSignificantDigitIdxs_03"

  //                                        1         2         3
  //                             0.1234567890123456789012345678901234567
  originalNumberStr := "-000123456.123456000"

  expectedAryLen := 18

  //                                        1         2         3
  //                             0.1234567890123456789012345678901234567
  expectedNumberStr := "-000123456.123456000"

  expectedPrecisionInt := 9

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntegerLen := 9

  expectedSignificantIntegerLen := 6

  expectedSignificantFractionLen := 6

  var expectedIsZeroValue, expectedIsIntegerZeroValue bool

  expectedIsZeroValue = false

  expectedIsIntegerZeroValue = false

  expectedFirstDigitIdx := 3

  expectedLastDigitIdx := 14

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

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
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

  intAryStats := intAry.GetIntAryStats()

  if expectedFirstDigitIdx != intAryStats.FirstDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array First Digits DO NOT MATCH!\n"+
      "Because expectedFirstDigitIdx != intAryStats.FirstDigitIdx\n"+
      "Expected intAryStats.FirstDigitIdx = '%v'\n"+
      "  Actual intAryStats.FirstDigitIdx = '%v'\n\n",
      ePrefix, expectedFirstDigitIdx, intAryStats.FirstDigitIdx)

    return
  }

  if expectedLastDigitIdx != intAryStats.LastDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array Last Digits DO NOT MATCH!\n"+
      "Because expectedLastDigitIdx != intAryStats.LastDigitIdx\n"+
      "Expected intAryStats.LastDigitIdx = '%v'\n"+
      "  Actual intAryStats.LastDigitIdx = '%v'\n\n",
      ePrefix, expectedLastDigitIdx, intAryStats.LastDigitIdx)

    return
  }

  if expectedIsZeroValue != intAryStats.IsZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsZeroValue != intAryStats.IsZeroValue\n"+
      "Expected intAryStats.IsZeroValue = '%v'\n"+
      "  Actual intAryStats.IsZeroValue = '%v'\n\n",
      ePrefix, expectedIsZeroValue, intAryStats.IsZeroValue)

    return
  }

  if expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsIntegerZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue\n"+
      "Expected intAryStats.IsIntegerZeroValue = '%v'\n"+
      "  Actual intAryStats.IsIntegerZeroValue = '%v'\n\n",
      ePrefix, expectedIsIntegerZeroValue, intAryStats.IsIntegerZeroValue)

    return
  }

  if expectedAryLen != intAryStats.IntAryLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedAryLen != intAryStats.IntAryLen\n"+
      "Expected intAryStats.IntAryLen = '%v'\n"+
      "  Actual intAryStats.IntAryLen = '%v'\n\n",
      ePrefix, expectedAryLen, intAryStats.IntAryLen)

    return
  }

  if expectedIntegerLen != intAryStats.IntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Integer Lengths ARE NOT EQUAL!\n"+
      "Because expectedIntegerLen != intAryStats.IntegerLen\n"+
      "Expected intAryStats.IntegerLen = '%v'\n"+
      "  Actual intAryStats.IntegerLen = '%v'\n\n",
      ePrefix, expectedIntegerLen, intAryStats.IntegerLen)

    return
  }

  if expectedSignificantFractionLen != intAryStats.SignificantFractionLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Fraction Lengths NOT EQUAL!\n"+
      "Because expectedSignificantFractionLen != intAryStats.SignificantFractionLen\n"+
      "Expected intAryStats.SignificantFractionLen = '%v'\n"+
      "  Actual intAryStats.SignificantFractionLen = '%v'\n\n",
      ePrefix, expectedSignificantFractionLen, intAryStats.SignificantFractionLen)

    return
  }

  if expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Integer Lengths NOT EQUAL!\n"+
      "Because expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen\n"+
      "Expected intAryStats.SignificantIntegerLen = '%v'\n"+
      "  Actual intAryStats.SignificantIntegerLen = '%v'\n\n",
      ePrefix, expectedSignificantIntegerLen, intAryStats.SignificantIntegerLen)

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

func TestIntAry_SetSignificantDigitIdxs_04(t *testing.T) {

  ePrefix := "TestIntAry_SetSignificantDigitIdxs_04"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "000.123456000"

  expectedAryLen := 12

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "000.123456000"

  expectedPrecisionInt := 9

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntegerLen := 3

  expectedSignificantIntegerLen := 1

  expectedSignificantFractionLen := 6

  var expectedIsZeroValue, expectedIsIntegerZeroValue bool

  expectedIsZeroValue = false

  expectedIsIntegerZeroValue = true

  expectedFirstDigitIdx := 2

  expectedLastDigitIdx := 8

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

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
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

  intAryStats := intAry.GetIntAryStats()

  if expectedFirstDigitIdx != intAryStats.FirstDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array First Digits DO NOT MATCH!\n"+
      "Because expectedFirstDigitIdx != intAryStats.FirstDigitIdx\n"+
      "Expected intAryStats.FirstDigitIdx = '%v'\n"+
      "  Actual intAryStats.FirstDigitIdx = '%v'\n\n",
      ePrefix, expectedFirstDigitIdx, intAryStats.FirstDigitIdx)

    return
  }

  if expectedLastDigitIdx != intAryStats.LastDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array Last Digits DO NOT MATCH!\n"+
      "Because expectedLastDigitIdx != intAryStats.LastDigitIdx\n"+
      "Expected intAryStats.LastDigitIdx = '%v'\n"+
      "  Actual intAryStats.LastDigitIdx = '%v'\n\n",
      ePrefix, expectedLastDigitIdx, intAryStats.LastDigitIdx)

    return
  }

  if expectedIsZeroValue != intAryStats.IsZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsZeroValue != intAryStats.IsZeroValue\n"+
      "Expected intAryStats.IsZeroValue = '%v'\n"+
      "  Actual intAryStats.IsZeroValue = '%v'\n\n",
      ePrefix, expectedIsZeroValue, intAryStats.IsZeroValue)

    return
  }

  if expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsIntegerZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue\n"+
      "Expected intAryStats.IsIntegerZeroValue = '%v'\n"+
      "  Actual intAryStats.IsIntegerZeroValue = '%v'\n\n",
      ePrefix, expectedIsIntegerZeroValue, intAryStats.IsIntegerZeroValue)

    return
  }

  if expectedAryLen != intAryStats.IntAryLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedAryLen != intAryStats.IntAryLen\n"+
      "Expected intAryStats.IntAryLen = '%v'\n"+
      "  Actual intAryStats.IntAryLen = '%v'\n\n",
      ePrefix, expectedAryLen, intAryStats.IntAryLen)

    return
  }

  if expectedIntegerLen != intAryStats.IntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Integer Lengths ARE NOT EQUAL!\n"+
      "Because expectedIntegerLen != intAryStats.IntegerLen\n"+
      "Expected intAryStats.IntegerLen = '%v'\n"+
      "  Actual intAryStats.IntegerLen = '%v'\n\n",
      ePrefix, expectedIntegerLen, intAryStats.IntegerLen)

    return
  }

  if expectedSignificantFractionLen != intAryStats.SignificantFractionLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Fraction Lengths NOT EQUAL!\n"+
      "Because expectedSignificantFractionLen != intAryStats.SignificantFractionLen\n"+
      "Expected intAryStats.SignificantFractionLen = '%v'\n"+
      "  Actual intAryStats.SignificantFractionLen = '%v'\n\n",
      ePrefix, expectedSignificantFractionLen, intAryStats.SignificantFractionLen)

    return
  }

  if expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Integer Lengths NOT EQUAL!\n"+
      "Because expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen\n"+
      "Expected intAryStats.SignificantIntegerLen = '%v'\n"+
      "  Actual intAryStats.SignificantIntegerLen = '%v'\n\n",
      ePrefix, expectedSignificantIntegerLen, intAryStats.SignificantIntegerLen)

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

func TestIntAry_SetSignificantDigitIdxs_05(t *testing.T) {

  ePrefix := "TestIntAry_SetSignificantDigitIdxs_05"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "256"

  expectedAryLen := 3

  expectedNumberStr := "256"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntegerLen := 3

  expectedSignificantIntegerLen := 3

  expectedSignificantFractionLen := 0

  var expectedIsZeroValue, expectedIsIntegerZeroValue bool

  expectedIsZeroValue = false

  expectedIsIntegerZeroValue = false

  expectedFirstDigitIdx := 0

  expectedLastDigitIdx := 2

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

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
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

  intAryStats := intAry.GetIntAryStats()

  if expectedFirstDigitIdx != intAryStats.FirstDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array First Digits DO NOT MATCH!\n"+
      "Because expectedFirstDigitIdx != intAryStats.FirstDigitIdx\n"+
      "Expected intAryStats.FirstDigitIdx = '%v'\n"+
      "  Actual intAryStats.FirstDigitIdx = '%v'\n\n",
      ePrefix, expectedFirstDigitIdx, intAryStats.FirstDigitIdx)

    return
  }

  if expectedLastDigitIdx != intAryStats.LastDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array Last Digits DO NOT MATCH!\n"+
      "Because expectedLastDigitIdx != intAryStats.LastDigitIdx\n"+
      "Expected intAryStats.LastDigitIdx = '%v'\n"+
      "  Actual intAryStats.LastDigitIdx = '%v'\n\n",
      ePrefix, expectedLastDigitIdx, intAryStats.LastDigitIdx)

    return
  }

  if expectedIsZeroValue != intAryStats.IsZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsZeroValue != intAryStats.IsZeroValue\n"+
      "Expected intAryStats.IsZeroValue = '%v'\n"+
      "  Actual intAryStats.IsZeroValue = '%v'\n\n",
      ePrefix, expectedIsZeroValue, intAryStats.IsZeroValue)

    return
  }

  if expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsIntegerZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue\n"+
      "Expected intAryStats.IsIntegerZeroValue = '%v'\n"+
      "  Actual intAryStats.IsIntegerZeroValue = '%v'\n\n",
      ePrefix, expectedIsIntegerZeroValue, intAryStats.IsIntegerZeroValue)

    return
  }

  if expectedAryLen != intAryStats.IntAryLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedAryLen != intAryStats.IntAryLen\n"+
      "Expected intAryStats.IntAryLen = '%v'\n"+
      "  Actual intAryStats.IntAryLen = '%v'\n\n",
      ePrefix, expectedAryLen, intAryStats.IntAryLen)

    return
  }

  if expectedIntegerLen != intAryStats.IntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Integer Lengths ARE NOT EQUAL!\n"+
      "Because expectedIntegerLen != intAryStats.IntegerLen\n"+
      "Expected intAryStats.IntegerLen = '%v'\n"+
      "  Actual intAryStats.IntegerLen = '%v'\n\n",
      ePrefix, expectedIntegerLen, intAryStats.IntegerLen)

    return
  }

  if expectedSignificantFractionLen != intAryStats.SignificantFractionLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Fraction Lengths NOT EQUAL!\n"+
      "Because expectedSignificantFractionLen != intAryStats.SignificantFractionLen\n"+
      "Expected intAryStats.SignificantFractionLen = '%v'\n"+
      "  Actual intAryStats.SignificantFractionLen = '%v'\n\n",
      ePrefix, expectedSignificantFractionLen, intAryStats.SignificantFractionLen)

    return
  }

  if expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Integer Lengths NOT EQUAL!\n"+
      "Because expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen\n"+
      "Expected intAryStats.SignificantIntegerLen = '%v'\n"+
      "  Actual intAryStats.SignificantIntegerLen = '%v'\n\n",
      ePrefix, expectedSignificantIntegerLen, intAryStats.SignificantIntegerLen)

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

func TestIntAry_SetSignificantDigitIdxs_06(t *testing.T) {

  ePrefix := "TestIntAry_SetSignificantDigitIdxs_06"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  originalNumberStr := "000256"

  expectedAryLen := 6

  expectedNumberStr := "000256"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedIntegerLen := 6

  expectedSignificantIntegerLen := 3

  expectedSignificantFractionLen := 0

  var expectedIsZeroValue, expectedIsIntegerZeroValue bool

  expectedIsZeroValue = false

  expectedIsIntegerZeroValue = false

  expectedFirstDigitIdx := 3

  expectedLastDigitIdx := 5

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

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
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

  intAryStats := intAry.GetIntAryStats()

  if expectedFirstDigitIdx != intAryStats.FirstDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array First Digits DO NOT MATCH!\n"+
      "Because expectedFirstDigitIdx != intAryStats.FirstDigitIdx\n"+
      "Expected intAryStats.FirstDigitIdx = '%v'\n"+
      "  Actual intAryStats.FirstDigitIdx = '%v'\n\n",
      ePrefix, expectedFirstDigitIdx, intAryStats.FirstDigitIdx)

    return
  }

  if expectedLastDigitIdx != intAryStats.LastDigitIdx {
    t.Errorf("%v\n"+
      "Error: Array Last Digits DO NOT MATCH!\n"+
      "Because expectedLastDigitIdx != intAryStats.LastDigitIdx\n"+
      "Expected intAryStats.LastDigitIdx = '%v'\n"+
      "  Actual intAryStats.LastDigitIdx = '%v'\n\n",
      ePrefix, expectedLastDigitIdx, intAryStats.LastDigitIdx)

    return
  }

  if expectedIsZeroValue != intAryStats.IsZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsZeroValue != intAryStats.IsZeroValue\n"+
      "Expected intAryStats.IsZeroValue = '%v'\n"+
      "  Actual intAryStats.IsZeroValue = '%v'\n\n",
      ePrefix, expectedIsZeroValue, intAryStats.IsZeroValue)

    return
  }

  if expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IsIntegerZeroValue's ARE NOT EQUAL!\n"+
      "Because expectedIsIntegerZeroValue != intAryStats.IsIntegerZeroValue\n"+
      "Expected intAryStats.IsIntegerZeroValue = '%v'\n"+
      "  Actual intAryStats.IsIntegerZeroValue = '%v'\n\n",
      ePrefix, expectedIsIntegerZeroValue, intAryStats.IsIntegerZeroValue)

    return
  }

  if expectedAryLen != intAryStats.IntAryLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedAryLen != intAryStats.IntAryLen\n"+
      "Expected intAryStats.IntAryLen = '%v'\n"+
      "  Actual intAryStats.IntAryLen = '%v'\n\n",
      ePrefix, expectedAryLen, intAryStats.IntAryLen)

    return
  }

  if expectedIntegerLen != intAryStats.IntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Integer Lengths ARE NOT EQUAL!\n"+
      "Because expectedIntegerLen != intAryStats.IntegerLen\n"+
      "Expected intAryStats.IntegerLen = '%v'\n"+
      "  Actual intAryStats.IntegerLen = '%v'\n\n",
      ePrefix, expectedIntegerLen, intAryStats.IntegerLen)

    return
  }

  if expectedSignificantFractionLen != intAryStats.SignificantFractionLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Fraction Lengths NOT EQUAL!\n"+
      "Because expectedSignificantFractionLen != intAryStats.SignificantFractionLen\n"+
      "Expected intAryStats.SignificantFractionLen = '%v'\n"+
      "  Actual intAryStats.SignificantFractionLen = '%v'\n\n",
      ePrefix, expectedSignificantFractionLen, intAryStats.SignificantFractionLen)

    return
  }

  if expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Significant Integer Lengths NOT EQUAL!\n"+
      "Because expectedSignificantIntegerLen != intAryStats.SignificantIntegerLen\n"+
      "Expected intAryStats.SignificantIntegerLen = '%v'\n"+
      "  Actual intAryStats.SignificantIntegerLen = '%v'\n\n",
      ePrefix, expectedSignificantIntegerLen, intAryStats.SignificantIntegerLen)

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

func TestIntAry_SetThousandsSeparator_01(t *testing.T) {

  ePrefix := "TestIntAry_SetThousandsSeparator_01"

  expectedNumberStr := "450123647,1234"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  var expectedFrenchDecSeparator rune

  expectedFrenchDecSeparator = ','

  var expectedFrenchThousandsSeparator rune

  expectedFrenchThousandsSeparator = ' '

  intAry := new(IntAry).New()

  err := intAry.SetDecimalSeparator(expectedFrenchDecSeparator)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.SetDecimalSeparator(expectedFrenchDecSeparator)\n"+
      "expectedFrenchDecSeparator= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedFrenchDecSeparator, err.Error())
    return
  }

  err = intAry.SetThousandsSeparator(expectedFrenchThousandsSeparator)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.SetThousandsSeparator(expectedFrenchThousandsSeparator)\n"+
      "expectedFrenchThousandsSeparator= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedFrenchThousandsSeparator, err.Error())
    return
  }

  err = intAry.SetIntAryWithNumStr("450 123 647,1234")

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAry.IsValid('Validating intAry')\n"+
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

  intAryDecimalSeparator := intAry.GetDecimalSeparator()

  if expectedFrenchDecSeparator != intAryDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
      "Because expectedFrenchDecSeparator != intAryDecimalSeparator\n"+
      "Expected intAryDecimalSeparator = '%v'\n"+
      "  Actual intAryDecimalSeparator = '%v'\n\n",
      ePrefix, expectedFrenchDecSeparator, intAryDecimalSeparator)

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

  return
}

func TestIntAry_ShiftPrecisionLeft_01(t *testing.T) {

  nStr1 := "900777"
  shiftPrecisionLeft := uint(3)
  expectedStr := "900.777"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionLeft(shiftPrecisionLeft)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionLeft_02(t *testing.T) {

  nStr1 := "0.900777"
  shiftPrecisionLeft := uint(3)
  expectedStr := "0.000900777"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionLeft(shiftPrecisionLeft)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionLeft_03(t *testing.T) {

  nStr1 := "0"
  shiftPrecisionLeft := uint(3)
  expectedStr := "0.000"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionLeft(shiftPrecisionLeft)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionLeft_04(t *testing.T) {

  nStr1 := "0.0"
  shiftPrecisionLeft := uint(3)
  expectedStr := "0.0000"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionLeft(shiftPrecisionLeft)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionLeft_05(t *testing.T) {

  nStr1 := "-900777"
  shiftPrecisionLeft := uint(3)
  expectedStr := "-900.777"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionLeft(shiftPrecisionLeft)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionLeft_06(t *testing.T) {

  nStr1 := "-900.777"
  shiftPrecisionLeft := uint(3)
  expectedStr := "-0.900777"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionLeft(shiftPrecisionLeft)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionRight_01(t *testing.T) {

  nStr1 := "900.777"
  shiftPrecisionRight := uint(3)
  expectedStr := "900777"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionRight(shiftPrecisionRight)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionRight_02(t *testing.T) {

  nStr1 := "900777"
  shiftPrecisionRight := uint(3)
  expectedStr := "900777000"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionRight(shiftPrecisionRight)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionRight_03(t *testing.T) {

  nStr1 := "0.900777"
  shiftPrecisionRight := uint(3)
  expectedStr := "900.777"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionRight(shiftPrecisionRight)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionRight_04(t *testing.T) {

  nStr1 := "0"
  shiftPrecisionRight := uint(3)
  expectedStr := "0000"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionRight(shiftPrecisionRight)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionRight_05(t *testing.T) {

  nStr1 := "0.000"
  shiftPrecisionRight := uint(3)
  expectedStr := "0000"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionRight(shiftPrecisionRight)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionRight_06(t *testing.T) {

  nStr1 := "-900.777"
  shiftPrecisionRight := uint(3)
  expectedStr := "-900777"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionRight(shiftPrecisionRight)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionRight_07(t *testing.T) {

  nStr1 := "-900777"
  shiftPrecisionRight := uint(3)
  expectedStr := "-900777000"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionRight(shiftPrecisionRight)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_ShiftPrecisionRight_08(t *testing.T) {

  nStr1 := "900777"
  shiftPrecisionRight := uint(3)
  expectedStr := "900777000"

  ia1, err := IntAry{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1) "+
      "nStr1='%v' Error='%v' ", nStr1, err.Error())
  }

  ia1.ShiftPrecisionRight(shiftPrecisionRight)

  if expectedStr != ia1.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedStr, ia1.GetNumStr())
  }

}

func TestIntAry_SubtractFromThis_01(t *testing.T) {
  nStr1 := "900.777"
  nStr2 := "901.000"
  eNumStr := "-0.223"
  ePrecision := 3
  eSignVal := -1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_02(t *testing.T) {
  nStr1 := "350"
  nStr2 := "122"
  eNumStr := "228"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_03(t *testing.T) {
  nStr1 := "-350"
  nStr2 := "122"
  eNumStr := "-472"
  ePrecision := 0
  eSignVal := -1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_04(t *testing.T) {
  nStr1 := "-350"
  nStr2 := "-122"
  eNumStr := "-228"
  ePrecision := 0
  eSignVal := -1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_05(t *testing.T) {
  nStr1 := "350"
  nStr2 := "-122"
  eNumStr := "472"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_06(t *testing.T) {
  nStr1 := "350"
  nStr2 := "0"
  eNumStr := "350"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_07(t *testing.T) {
  nStr1 := "-350"
  nStr2 := "0"
  eNumStr := "-350"
  ePrecision := 0
  eSignVal := -1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_08(t *testing.T) {
  nStr1 := "122"
  nStr2 := "350"
  eNumStr := "-228"
  ePrecision := 0
  eSignVal := -1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_09(t *testing.T) {
  nStr1 := "-122"
  nStr2 := "350"
  eNumStr := "-472"
  ePrecision := 0
  eSignVal := -1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_10(t *testing.T) {
  nStr1 := "-122"
  nStr2 := "-350"
  eNumStr := "228"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_11(t *testing.T) {
  nStr1 := "122"
  nStr2 := "-350"
  eNumStr := "472"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_12(t *testing.T) {
  nStr1 := "0"
  nStr2 := "350"
  eNumStr := "-350"
  ePrecision := 0
  eSignVal := -1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_13(t *testing.T) {
  nStr1 := "0"
  nStr2 := "-350"
  eNumStr := "350"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_14(t *testing.T) {
  nStr1 := "122"
  nStr2 := "122"
  eNumStr := "0"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_15(t *testing.T) {
  nStr1 := "-122"
  nStr2 := "122"
  eNumStr := "-244"
  ePrecision := 0
  eSignVal := -1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_16(t *testing.T) {
  nStr1 := "-122"
  nStr2 := "-122"
  eNumStr := "0"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_17(t *testing.T) {
  nStr1 := "122"
  nStr2 := "-122"
  eNumStr := "244"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_18(t *testing.T) {
  nStr1 := "0"
  nStr2 := "0"
  eNumStr := "0"
  ePrecision := 0
  eSignVal := 1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractFromThis_19(t *testing.T) {
  nStr1 := "1.122"
  nStr2 := "4.5"
  eNumStr := "-3.378"
  ePrecision := 3
  eSignVal := -1

  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  err := ia1.SubtractFromThis(&ia2)

  if err != nil {
    t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
  }

  if eNumStr != ia1.GetNumStr() {
    t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
  }

  if ePrecision != ia1.GetPrecision() {
    t.Errorf("Error - Expected ia1.GetPrecisionInt()= '%v' .  Instead, ia1.GetPrecisionInt()= '%v' .", ePrecision, ia1.GetPrecision())
  }

  if eSignVal != ia1.GetSign() {
    t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
  }

}

func TestIntAry_SubtractMultipleFromThis_01(t *testing.T) {
  nStrBase := "197.452"
  nStr1 := "1.122"
  nStr2 := "4.5"
  nStr3 := "32.148"
  nStr4 := "10.0"
  eNumStr := "149.682"
  ePrecision := 3
  eSignVal := 1

  iaBase := IntAry{}.New()
  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia3 := IntAry{}.New()
  ia4 := IntAry{}.New()
  iaBase.SetIntAryWithNumStr(nStrBase)
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  ia3.SetIntAryWithNumStr(nStr3)
  ia4.SetIntAryWithNumStr(nStr4)
  err := iaBase.SubtractMultipleFromThis(&ia1, &ia2, &ia3, &ia4)

  if err != nil {
    t.Errorf("Error returned from iaBase.SubtractMultipleFromThis(true, ia...). Error= %v", err)
  }

  if eNumStr != iaBase.GetNumStr() {
    t.Errorf("Error - Expected iaBase.GetNumStr()= '%v' .  Instead, iaBase.GetNumStr()= '%v' .", eNumStr, iaBase.GetNumStr())
  }

  if ePrecision != iaBase.GetPrecision() {
    t.Errorf("Error - Expected iaBase.GetPrecisionInt()= '%v' .  Instead, iaBase.GetPrecisionInt()= '%v' .", ePrecision, iaBase.GetPrecision())
  }

  if eSignVal != iaBase.GetSign() {
    t.Errorf("Error - Expected iaBase.GetSign()= '%v' .  Instead, iaBase.GetSign()= '%v' .", eSignVal, iaBase.GetSign())
  }

}

func TestIntAry_SubtractMultipleFromThis_02(t *testing.T) {
  nStrBase := "197.452"
  nStr1 := "1.122"
  nStr2 := "4.5"
  nStr3 := "-32.148"
  nStr4 := "10.0"
  eNumStr := "213.978"
  ePrecision := 3
  eSignVal := 1

  iaBase := IntAry{}.New()
  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia3 := IntAry{}.New()
  ia4 := IntAry{}.New()
  iaBase.SetIntAryWithNumStr(nStrBase)
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  ia3.SetIntAryWithNumStr(nStr3)
  ia4.SetIntAryWithNumStr(nStr4)
  err := iaBase.SubtractMultipleFromThis(&ia1, &ia2, &ia3, &ia4)

  if err != nil {
    t.Errorf("Error returned from iaBase.SubtractMultipleFromThis(true, ia...). Error= %v", err)
  }

  if eNumStr != iaBase.GetNumStr() {
    t.Errorf("Error - Expected iaBase.GetNumStr()= '%v' .  Instead, iaBase.GetNumStr()= '%v' .", eNumStr, iaBase.GetNumStr())
  }

  if ePrecision != iaBase.GetPrecision() {
    t.Errorf("Error - Expected iaBase.GetPrecisionInt()= '%v' .  Instead, iaBase.GetPrecisionInt()= '%v' .", ePrecision, iaBase.GetPrecision())
  }

  if eSignVal != iaBase.GetSign() {
    t.Errorf("Error - Expected iaBase.GetSign()= '%v' .  Instead, iaBase.GetSign()= '%v' .", eSignVal, iaBase.GetSign())
  }

}

func TestIntAry_SubtractMultipleFromThis_03(t *testing.T) {
  nStrBase := "197.452"
  nStr1 := "1.122"
  nStr2 := "4.5"
  nStr3 := "932.148"
  nStr4 := "10.0"
  eNumStr := "-750.318"
  ePrecision := 3
  eSignVal := -1

  iaBase := IntAry{}.New()
  ia1 := IntAry{}.New()
  ia2 := IntAry{}.New()
  ia3 := IntAry{}.New()
  ia4 := IntAry{}.New()
  iaBase.SetIntAryWithNumStr(nStrBase)
  ia1.SetIntAryWithNumStr(nStr1)
  ia2.SetIntAryWithNumStr(nStr2)
  ia3.SetIntAryWithNumStr(nStr3)
  ia4.SetIntAryWithNumStr(nStr4)
  err := iaBase.SubtractMultipleFromThis(&ia1, &ia2, &ia3, &ia4)

  if err != nil {
    t.Errorf("Error returned from iaBase.SubtractMultipleFromThis(true, ia...). Error= %v", err)
  }

  if eNumStr != iaBase.GetNumStr() {
    t.Errorf("Error - Expected iaBase.GetNumStr()= '%v' .  Instead, iaBase.GetNumStr()= '%v' .", eNumStr, iaBase.GetNumStr())
  }

  if ePrecision != iaBase.GetPrecision() {
    t.Errorf("Error - Expected iaBase.GetPrecisionInt()= '%v' .  Instead, iaBase.GetPrecisionInt()= '%v' .", ePrecision, iaBase.GetPrecision())
  }

  if eSignVal != iaBase.GetSign() {
    t.Errorf("Error - Expected iaBase.GetSign()= '%v' .  Instead, iaBase.GetSign()= '%v' .", eSignVal, iaBase.GetSign())
  }

}
