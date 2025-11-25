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

  ePrefix := "TestIntAry_ShiftPrecisionLeft_01"

  originalNumberStr := "900777"

  originalShiftPrecisionLeft := uint(3)

  expectedNumberStr := "900.777"

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

  err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionLeft= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionLeft,
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

func TestIntAry_ShiftPrecisionLeft_02(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionLeft_02"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  originalNumberStr := "0.900777"

  originalShiftPrecisionLeft := uint(3)

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.000900777"

  expectedPrecisionInt := 9

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

  err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionLeft= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionLeft,
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

func TestIntAry_ShiftPrecisionLeft_03(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionLeft_03"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  originalNumberStr := "0"

  originalShiftPrecisionLeft := uint(3)

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.000"

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

  err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionLeft= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionLeft,
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

func TestIntAry_ShiftPrecisionLeft_04(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionLeft_04"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  originalNumberStr := "0.0"

  originalShiftPrecisionLeft := uint(3)

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0.0000"

  expectedPrecisionInt := 4

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

  err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionLeft= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionLeft,
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

func TestIntAry_ShiftPrecisionLeft_05(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionLeft_05"

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  originalNumberStr := "-900777"

  originalShiftPrecisionLeft := uint(3)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-900.777"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionLeft= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionLeft,
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

func TestIntAry_ShiftPrecisionLeft_06(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionLeft_06"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalNumberStr := "-900.777"

  originalShiftPrecisionLeft := uint(3)

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr := "-0.900777"

  expectedPrecisionInt := 6

  expectedPrecisionUint := uint(expectedPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionLeft(originalShiftPrecisionLeft)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionLeft= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionLeft,
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

func TestIntAry_ShiftPrecisionRight_01(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionRight_01"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "900.777"

  originalShiftPrecisionRight := uint(3)

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr := "900777"

  expectedPrecisionInt := 0

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

  err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionRight= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionRight,
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

func TestIntAry_ShiftPrecisionRight_02(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionRight_02"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  originalNumberStr := "900777"

  originalShiftPrecisionRight := uint(3)

  //                                       1         2         3
  //                            0.1234567890123456789012345678901234567
  expectedNumberStr := "900777000"

  expectedPrecisionInt := 0

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

  err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionRight= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionRight,
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

func TestIntAry_ShiftPrecisionRight_03(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionRight_03"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  originalNumberStr := "0.900777"

  originalShiftPrecisionRight := uint(3)

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "900.777"

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

  err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionRight= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionRight,
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
      "intAryNumberStr, err = intAry.GetNumStr()\n"+
      "intAry Number String set to final value\n"+
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

func TestIntAry_ShiftPrecisionRight_04(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionRight_04"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  originalNumberStr := "0"

  originalShiftPrecisionRight := uint(3)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "0000"

  expectedPrecisionInt := 0

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

  err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionRight= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionRight,
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

func TestIntAry_ShiftPrecisionRight_05(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionRight_05"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  originalNumberStr := "0.000"

  originalShiftPrecisionRight := uint(3)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "0000"

  expectedPrecisionInt := 0

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

  err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionRight= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionRight,
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

func TestIntAry_ShiftPrecisionRight_06(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionRight_06"

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  originalNumberStr := "-900777"

  originalShiftPrecisionRight := uint(3)

  //                                        1         2         3
  //                             0.1234567890123456789012345678901234567
  expectedNumberStr := "-900777000"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
      "Because originalNumberStr != intAryNumberStr\n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionRight= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionRight,
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

func TestIntAry_ShiftPrecisionRight_07(t *testing.T) {

  ePrefix := "TestIntAry_ShiftPrecisionRight_07"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  originalNumberStr := "900777"

  originalShiftPrecisionRight := uint(3)

  //                                       1         2         3
  //                            0.1234567890123456789012345678901234567
  expectedNumberStr := "900777000"

  expectedPrecisionInt := 0

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

  err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.ShiftPrecisionRight(originalShiftPrecisionRight)\n"+
      "intAry= '%v'\n"+
      "originalShiftPrecisionRight= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      originalShiftPrecisionRight,
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

func TestIntAry_SubtractFromThis_01(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_01"

  originalNumberStr1 := "900.777"

  originalNumberStr2 := "901.000"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr := "-0.223"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_02(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_02"

  originalNumberStr1 := "350"

  originalNumberStr2 := "122"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "228"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_03(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_03"

  originalNumberStr1 := "-350"

  originalNumberStr2 := "122"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-472"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_04(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_04"

  originalNumberStr1 := "-350"

  originalNumberStr2 := "-122"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-228"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_05(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_05"

  originalNumberStr1 := "350"

  originalNumberStr2 := "-122"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "472"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_06(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_06"

  originalNumberStr1 := "350"

  originalNumberStr2 := "0"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "350"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_07(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_07"

  originalNumberStr1 := "-350"

  originalNumberStr2 := "0"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-350"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_08(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_08"

  originalNumberStr1 := "122"

  originalNumberStr2 := "350"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-228"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_09(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_09"

  originalNumberStr1 := "-122"

  originalNumberStr2 := "350"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-472"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_10(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_10"

  originalNumberStr1 := "-122"

  originalNumberStr2 := "-350"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "228"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_11(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_11"

  originalNumberStr1 := "122"

  originalNumberStr2 := "-350"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "472"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_12(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_12"

  originalNumberStr1 := "0"

  originalNumberStr2 := "350"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-350"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_13(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_13"

  originalNumberStr1 := "0"

  originalNumberStr2 := "-350"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "350"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_14(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_14"

  originalNumberStr1 := "122"

  originalNumberStr2 := "122"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_15(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_15"

  originalNumberStr1 := "-122"

  originalNumberStr2 := "122"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-244"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_16(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_16"

  originalNumberStr1 := "-122"

  originalNumberStr2 := "-122"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_17(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_17"

  originalNumberStr1 := "122"

  originalNumberStr2 := "-122"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "244"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_18(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_18"

  originalNumberStr1 := "0"

  originalNumberStr2 := "0"

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractFromThis_19(t *testing.T) {

  ePrefix := "TestIntAry_SubtractFromThis_19"

  originalNumberStr1 := "1.122"

  originalNumberStr2 := "4.5"

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr := "-3.378"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
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
      "Error: Unexpected Result Testing Intial IntAry Number String!\n"+
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

  err = intAry1.SubtractFromThis(&intAry2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry1.SubtractFromThis(&intAry2)\n"+
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
      "intAry1 Number String set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
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

  intAry1SignValue, err := intAry1.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1SignValue, err := intAry1.GetSign()\n"+
      "intAry1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAry1NumberStr, err.Error())
    return
  }

  intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
      "intAry1= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAry1NumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAry1NumberStr \n"+
      "Expected intAry1NumberStr = '%v'\n"+
      "  Actual intAry1NumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAry1NumberStr)

    return
  }

  if expectedPrecisionInt != intAry1PrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAry1PrecisionInt\n"+
      "Expected intAry1PrecisionInt = '%v'\n"+
      "  Actual intAry1PrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

    return
  }

  if expectedPrecisionUint != intAry1PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAry1PrecisionUint\n"+
      "Expected intAry1PrecisionUint = '%v'\n"+
      "  Actual intAry1PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

    return
  }

  if expectedSignValue != intAry1SignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAry1SignValue\n"+
      "Expected intAry1SignValue = '%v'\n"+
      "  Actual intAry1SignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAry1SignValue)

    return
  }

  if !expectedNumSeps.Equal(intAry1NumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAry1NumSeps \n"+
      "Expected intAry1NumSeps = '%v'\n"+
      "  Actual intAry1NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

    return
  }

  return
}

func TestIntAry_SubtractMultipleFromThis_01(t *testing.T) {

  ePrefix := "TestIntAry_SubtractMultipleFromThis_01"

  originalBaseNumberStr := "197.452"

  originalNumberStr1 := "1.122"

  originalNumberStr2 := "4.5"

  originalNumberStr3 := "32.148"

  originalNumberStr4 := "10.0"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "149.682"

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
      "intAryBase set to initial value\n"+
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

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
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

  intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)\n"+
      "originalNumberStr3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr3, err.Error())
    return
  }

  err = intAry3.IsValid("Validating intAry3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry3.IsValid('Validating intAry3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry3NumberStr, err := intAry3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry3NumberStr, err := intAry3.GetNumStr()\n"+
      "intAry3 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr3 != intAry3NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr3 != intAry3NumberStr\n"+
      "Expected intAry3NumberStr = '%v'\n"+
      "  Actual intAry3NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr3, intAry3NumberStr)

    return
  }

  intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)\n"+
      "originalNumberStr4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr4, err.Error())
    return
  }

  err = intAry4.IsValid("Validating intAry4")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry4.IsValid('Validating intAry4')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry4NumberStr, err := intAry4.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry4NumberStr, err := intAry4.GetNumStr()\n"+
      "intAry4 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr4 != intAry4NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr4 != intAry4NumberStr\n"+
      "Expected intAry4NumberStr = '%v'\n"+
      "  Actual intAry4NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr4, intAry4NumberStr)

    return
  }

  err = intAryBase.SubtractMultipleFromThis(&intAry1, &intAry2, &intAry3, &intAry4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.SubtractMultipleFromThis(\n"+
      "  &intAry1, &intAry2, &intAry3, &intAry4)\n"+
      "intAryBase= '%v'\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "intAry3= '%v'\n"+
      "intAry4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAry1NumberStr,
      intAry2NumberStr,
      intAry3NumberStr,
      intAry4NumberStr,
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

func TestIntAry_SubtractMultipleFromThis_02(t *testing.T) {

  ePrefix := "TestIntAry_SubtractMultipleFromThis_02"

  originalBaseNumberStr := "197.452"

  originalNumberStr1 := "1.122"

  originalNumberStr2 := "4.5"

  originalNumberStr3 := "-32.148"

  originalNumberStr4 := "10.0"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  expectedNumberStr := "213.978"

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
      "intAryBase set to initial value\n"+
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

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
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

  intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)\n"+
      "originalNumberStr3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr3, err.Error())
    return
  }

  err = intAry3.IsValid("Validating intAry3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry3.IsValid('Validating intAry3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry3NumberStr, err := intAry3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry3NumberStr, err := intAry3.GetNumStr()\n"+
      "intAry3 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr3 != intAry3NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr3 != intAry3NumberStr\n"+
      "Expected intAry3NumberStr = '%v'\n"+
      "  Actual intAry3NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr3, intAry3NumberStr)

    return
  }

  intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)\n"+
      "originalNumberStr4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr4, err.Error())
    return
  }

  err = intAry4.IsValid("Validating intAry4")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry4.IsValid('Validating intAry4')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry4NumberStr, err := intAry4.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry4NumberStr, err := intAry4.GetNumStr()\n"+
      "intAry4 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr4 != intAry4NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr4 != intAry4NumberStr\n"+
      "Expected intAry4NumberStr = '%v'\n"+
      "  Actual intAry4NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr4, intAry4NumberStr)

    return
  }

  err = intAryBase.SubtractMultipleFromThis(&intAry1, &intAry2, &intAry3, &intAry4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.SubtractMultipleFromThis(\n"+
      "  &intAry1, &intAry2, &intAry3, &intAry4)\n"+
      "intAryBase= '%v'\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "intAry3= '%v'\n"+
      "intAry4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAry1NumberStr,
      intAry2NumberStr,
      intAry3NumberStr,
      intAry4NumberStr,
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

func TestIntAry_SubtractMultipleFromThis_03(t *testing.T) {

  ePrefix := "TestIntAry_SubtractMultipleFromThis_03"

  originalBaseNumberStr := "197.452"

  originalNumberStr1 := "1.122"

  originalNumberStr2 := "4.5"

  originalNumberStr3 := "932.148"

  originalNumberStr4 := "10.0"

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedNumberStr := "-750.318"

  expectedPrecisionInt := 3

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
      "intAryBase set to initial value\n"+
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

  intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
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

  intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)\n"+
      "originalNumberStr3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr3, err.Error())
    return
  }

  err = intAry3.IsValid("Validating intAry3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry3.IsValid('Validating intAry3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry3NumberStr, err := intAry3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry3NumberStr, err := intAry3.GetNumStr()\n"+
      "intAry3 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr3 != intAry3NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr3 != intAry3NumberStr\n"+
      "Expected intAry3NumberStr = '%v'\n"+
      "  Actual intAry3NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr3, intAry3NumberStr)

    return
  }

  intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)\n"+
      "originalNumberStr4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr4, err.Error())
    return
  }

  err = intAry4.IsValid("Validating intAry4")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry4.IsValid('Validating intAry4')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry4NumberStr, err := intAry4.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry4NumberStr, err := intAry4.GetNumStr()\n"+
      "intAry4 set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr4 != intAry4NumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalNumberStr4 != intAry4NumberStr\n"+
      "Expected intAry4NumberStr = '%v'\n"+
      "  Actual intAry4NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr4, intAry4NumberStr)

    return
  }

  err = intAryBase.SubtractMultipleFromThis(&intAry1, &intAry2, &intAry3, &intAry4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryBase.SubtractMultipleFromThis(\n"+
      "  &intAry1, &intAry2, &intAry3, &intAry4)\n"+
      "intAryBase= '%v'\n"+
      "intAry1= '%v'\n"+
      "intAry2= '%v'\n"+
      "intAry3= '%v'\n"+
      "intAry4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryBaseNumberStr,
      intAry1NumberStr,
      intAry2NumberStr,
      intAry3NumberStr,
      intAry4NumberStr,
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
