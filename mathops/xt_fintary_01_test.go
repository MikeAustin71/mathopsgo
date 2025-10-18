package mathops

import (
  "testing"
)

func TestFracIntAry_GetRationalValue_01(t *testing.T) {

  ePrefix := "TestFracIntAry_GetRationalValue_01"

  numeratorNumStr := "2"

  denominatorNumStr := "3"

  //                                  1         2         3
  //                       0.12345678901234567890123456789012
  expectedResultNumStr := "0.66666666666666666666666666666667"

  originalMaxPrecisionInt := 32

  fracIntAry, err := new(FracIntAry).NewNumStrs(numeratorNumStr, denominatorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewNumStrs(\n"+
      "  numeratorNumStr, denominatorNumStr)\n"+
      "numeratorNumStr= '%v'\n"+
      "denominatorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numeratorNumStr,
      denominatorNumStr,
      err.Error())

    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMaxPrecisionInt,
      err.Error())

    return
  }

  numerartorPtrInt := rationalNum.Num()

  numerartorPtrIntNumberStr := numerartorPtrInt.String()

  if numeratorNumStr != numerartorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numerators Number Strings ARE NOT Equal\n"+
      "Because numeratorNumStr != numerartorPtrIntNumberStr \n"+
      "Expected numerartorPtrIntNumberStr = '%v'\n"+
      "  Actual numerartorPtrIntNumberStr = '%v'\n\n",
      ePrefix, numeratorNumStr, numerartorPtrIntNumberStr)

    return
  }

  denominatorPtrInt := rationalNum.Denom()

  denominatorPtrIntNumberStr := denominatorPtrInt.String()

  if denominatorNumStr != denominatorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numerators Number Strings ARE NOT Equal\n"+
      "Because denominatorNumStr != denominatorPtrIntNumberStr \n"+
      "Expected denominatorPtrIntNumberStr = '%v'\n"+
      "  Actual denominatorPtrIntNumberStr = '%v'\n\n",
      ePrefix, denominatorPtrIntNumberStr, denominatorPtrIntNumberStr)

    return
  }

  floatString := rationalNum.FloatString(originalMaxPrecisionInt)

  if expectedResultNumStr != floatString {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Floating Number String Values ARE NOT Equal\n"+
      "Because expectedResultNumStr != floatString \n"+
      "Expected floatString = '%v'\n"+
      "  Actual floatString = '%v'\n\n",
      ePrefix, expectedResultNumStr, floatString)

    return
  }

  return
}

func TestFracIntAry_GetRationalValue_02(t *testing.T) {

  ePrefix := "TestFracIntAry_GetRationalValue_02"

  numeratorNumStr := "3"

  denominatorNumStr := "4"

  expectedResultNumStr := "0.75"

  originalMaxPrecisionInt := 2

  fracIntAry, err := new(FracIntAry).NewNumStrs(numeratorNumStr, denominatorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewNumStrs(\n"+
      "  numeratorNumStr, denominatorNumStr)\n"+
      "numeratorNumStr= '%v'\n"+
      "denominatorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numeratorNumStr,
      denominatorNumStr,
      err.Error())

    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMaxPrecisionInt,
      err.Error())

    return
  }

  numerartorPtrInt := rationalNum.Num()

  numerartorPtrIntNumberStr := numerartorPtrInt.String()

  if numeratorNumStr != numerartorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numerator Number Strings ARE NOT Equal\n"+
      "Because numeratorNumStr != numerartorPtrIntNumberStr \n"+
      "Expected numerartorPtrIntNumberStr = '%v'\n"+
      "  Actual numerartorPtrIntNumberStr = '%v'\n\n",
      ePrefix, numeratorNumStr, numerartorPtrIntNumberStr)

    return
  }

  denominatorPtrInt := rationalNum.Denom()

  denominatorPtrIntNumberStr := denominatorPtrInt.String()

  if denominatorNumStr != denominatorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Denominator Number Strings ARE NOT Equal\n"+
      "Because denominatorNumStr != denominatorPtrIntNumberStr \n"+
      "Expected denominatorPtrIntNumberStr = '%v'\n"+
      "  Actual denominatorPtrIntNumberStr = '%v'\n\n",
      ePrefix, denominatorPtrIntNumberStr, denominatorPtrIntNumberStr)

    return
  }

  floatString := rationalNum.FloatString(originalMaxPrecisionInt)

  if expectedResultNumStr != floatString {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Floating Number String Values ARE NOT Equal\n"+
      "Because expectedResultNumStr != floatString \n"+
      "Expected floatString = '%v'\n"+
      "  Actual floatString = '%v'\n\n",
      ePrefix, expectedResultNumStr, floatString)

    return
  }

  return
}

func TestFracIntAry_GetRationalValue_03(t *testing.T) {

  ePrefix := "TestFracIntAry_GetRationalValue_03"

  numeratorNumStr := "1000"

  denominatorNumStr := "2000"

  expectedNumeratorNumStr := "1"

  expectedDenominatorNumStr := "2"

  expectedResultNumStr := "0.5"

  originalMaxPrecisionInt := 1

  fracIntAry, err := new(FracIntAry).NewNumStrs(numeratorNumStr, denominatorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewNumStrs(\n"+
      "  numeratorNumStr, denominatorNumStr)\n"+
      "numeratorNumStr= '%v'\n"+
      "denominatorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numeratorNumStr,
      denominatorNumStr,
      err.Error())

    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMaxPrecisionInt,
      err.Error())

    return
  }

  numerartorPtrInt := rationalNum.Num()

  numerartorPtrIntNumberStr := numerartorPtrInt.String()

  if expectedNumeratorNumStr != numerartorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numerator Number Strings ARE NOT Equal\n"+
      "Because expectedNumeratorNumStr != numerartorPtrIntNumberStr \n"+
      "Expected numerartorPtrIntNumberStr = '%v'\n"+
      "  Actual numerartorPtrIntNumberStr = '%v'\n\n",
      ePrefix, expectedNumeratorNumStr, numerartorPtrIntNumberStr)

    return
  }

  denominatorPtrInt := rationalNum.Denom()

  denominatorPtrIntNumberStr := denominatorPtrInt.String()

  if expectedDenominatorNumStr != denominatorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numerators Number Strings ARE NOT Equal\n"+
      "Because expectedDenominatorNumStr != denominatorPtrIntNumberStr \n"+
      "Expected denominatorPtrIntNumberStr = '%v'\n"+
      "  Actual denominatorPtrIntNumberStr = '%v'\n\n",
      ePrefix, expectedDenominatorNumStr, denominatorPtrIntNumberStr)

    return
  }

  floatString := rationalNum.FloatString(originalMaxPrecisionInt)

  if expectedResultNumStr != floatString {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Floating Number String Values ARE NOT Equal\n"+
      "Because expectedResultNumStr != floatString \n"+
      "Expected floatString = '%v'\n"+
      "  Actual floatString = '%v'\n\n",
      ePrefix, expectedResultNumStr, floatString)

    return
  }

  return
}

func TestFracIntAry_GetRationalValue_04(t *testing.T) {

  ePrefix := "TestFracIntAry_GetRationalValue_04"

  numeratorNumStr := "9.24"

  denominatorNumStr := "15.87"

  expectedNumerator := "291115311909262759924385633270321361058601"

  expectedDenominator := "500000000000000000000000000000000000000000"

  originalMaxPrecisionInt := 42

  //                                  1         2         3         4
  //                       0.123456789012345678901234567890123456789012
  expectedResultNumStr := "0.582230623818525519848771266540642722117202"

  fracIntAry, err := new(FracIntAry).NewNumStrs(numeratorNumStr, denominatorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewNumStrs(\n"+
      "  numeratorNumStr, denominatorNumStr)\n"+
      "numeratorNumStr= '%v'\n"+
      "denominatorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numeratorNumStr,
      denominatorNumStr,
      err.Error())

    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMaxPrecisionInt,
      err.Error())

    return
  }

  numerartorPtrInt := rationalNum.Num()

  numerartorPtrIntNumberStr := numerartorPtrInt.String()

  if expectedNumerator != numerartorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numerator Number Strings ARE NOT Equal\n"+
      "Because expectedNumerator != numerartorPtrIntNumberStr \n"+
      "Expected numerartorPtrIntNumberStr = '%v'\n"+
      "  Actual numerartorPtrIntNumberStr = '%v'\n\n",
      ePrefix, expectedNumerator, numerartorPtrIntNumberStr)

    return
  }

  denominatorPtrInt := rationalNum.Denom()

  denominatorPtrIntNumberStr := denominatorPtrInt.String()

  if expectedDenominator != denominatorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Denominator Number Strings ARE NOT Equal\n"+
      "Because expectedDenominator != denominatorPtrIntNumberStr \n"+
      "Expected denominatorPtrIntNumberStr = '%v'\n"+
      "  Actual denominatorPtrIntNumberStr = '%v'\n\n",
      ePrefix, expectedDenominator, denominatorPtrIntNumberStr)

    return
  }

  floatString := rationalNum.FloatString(originalMaxPrecisionInt)

  if expectedResultNumStr != floatString {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Floating Number String Values ARE NOT Equal\n"+
      "Because expectedResultNumStr != floatString \n"+
      "Expected floatString = '%v'\n"+
      "  Actual floatString = '%v'\n\n",
      ePrefix, expectedResultNumStr, floatString)

    return
  }

  return
}

func TestFracIntAry_GetRationalValue_05(t *testing.T) {

  ePrefix := "TestFracIntAry_GetRationalValue_05"

  numeratorNumStr := "3"

  denominatorNumStr := "4"

  expectedResultNumStr := "0.75"

  originalMaxPrecisionInt := 2

  intAryNumerator, err := new(IntAry).NewNumStr(numeratorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumerator, err := new(IntAry).NewNumStr(numeratorNumStr)\n"+
      "numeratorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numeratorNumStr, err.Error())
    return
  }

  err = intAryNumerator.IsValid("Validating intAryNumerator")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryNumerator.IsValid('Validating intAryNumerator')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumeratorNumberStr, err := intAryNumerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumeratorNumberStr, err := intAryNumerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numeratorNumStr != intAryNumeratorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Numerator Number String Values ARE NOT Equal\n"+
      "Because numeratorNumStr != intAryNumeratorNumberStr \n"+
      "Expected intAryNumeratorNumberStr = '%v'\n"+
      "  Actual intAryNumeratorNumberStr = '%v'\n\n",
      ePrefix, numeratorNumStr, intAryNumeratorNumberStr)

    return
  }

  intAryDenominator, err := new(IntAry).NewNumStr(denominatorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDenominator, err := new(IntAry).NewNumStr(denominatorNumStr)\n"+
      "denominatorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, denominatorNumStr, err.Error())
    return
  }

  err = intAryDenominator.IsValid("Validating intAryDenominator")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDenominator.IsValid('Validating intAryDenominator')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDenominatorNumberStr, err := intAryDenominator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDenominatorNumberStr, err := intAryDenominator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if denominatorNumStr != intAryDenominatorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Denominator Number String Values ARE NOT Equal\n"+
      "Because denominatorNumStr != intAryDenominatorNumberStr \n"+
      "Expected intAryDenominatorNumberStr = '%v'\n"+
      "  Actual intAryDenominatorNumberStr = '%v'\n\n",
      ePrefix, denominatorNumStr, intAryDenominatorNumberStr)

    return
  }

  fracIntAry, err := new(FracIntAry).NewIntArys(&intAryNumerator, &intAryDenominator)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewIntArys(\n"+
      "  &intAryNumerator, &intAryDenominator)\n"+
      "intAryNumerator= '%v'\n"+
      "intAryDenominator= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumeratorNumberStr,
      intAryDenominatorNumberStr,
      err.Error())

    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMaxPrecisionInt,
      err.Error())

    return
  }

  numerartorPtrInt := rationalNum.Num()

  numerartorPtrIntNumberStr := numerartorPtrInt.String()

  if numeratorNumStr != numerartorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numerator Number Strings ARE NOT Equal\n"+
      "Because numeratorNumStr != numerartorPtrIntNumberStr \n"+
      "Expected numerartorPtrIntNumberStr = '%v'\n"+
      "  Actual numerartorPtrIntNumberStr = '%v'\n\n",
      ePrefix, numeratorNumStr, numerartorPtrIntNumberStr)

    return
  }

  denominatorPtrInt := rationalNum.Denom()

  denominatorPtrIntNumberStr := denominatorPtrInt.String()

  if denominatorNumStr != denominatorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Denominator Number Strings ARE NOT Equal\n"+
      "Because denominatorNumStr != denominatorPtrIntNumberStr \n"+
      "Expected denominatorPtrIntNumberStr = '%v'\n"+
      "  Actual denominatorPtrIntNumberStr = '%v'\n\n",
      ePrefix, denominatorPtrIntNumberStr, denominatorPtrIntNumberStr)

    return
  }

  floatString := rationalNum.FloatString(originalMaxPrecisionInt)

  if expectedResultNumStr != floatString {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Floating Number String Values ARE NOT Equal\n"+
      "Because expectedResultNumStr != floatString \n"+
      "Expected floatString = '%v'\n"+
      "  Actual floatString = '%v'\n\n",
      ePrefix, expectedResultNumStr, floatString)

    return
  }

  return
}

func TestFracIntAry_GetRationalValue_06(t *testing.T) {

  ePrefix := "TestFracIntAry_GetRationalValue_06"

  numeratorNumStr := "-3"

  denominatorNumStr := "4"

  expectedResultNumStr := "-0.75"

  originalMaxPrecisionInt := 2

  intAryNumerator, err := new(IntAry).NewNumStr(numeratorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumerator, err := new(IntAry).NewNumStr(numeratorNumStr)\n"+
      "numeratorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numeratorNumStr, err.Error())
    return
  }

  err = intAryNumerator.IsValid("Validating intAryNumerator")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryNumerator.IsValid('Validating intAryNumerator')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumeratorNumberStr, err := intAryNumerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumeratorNumberStr, err := intAryNumerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numeratorNumStr != intAryNumeratorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Numerator Number String Values ARE NOT Equal\n"+
      "Because numeratorNumStr != intAryNumeratorNumberStr \n"+
      "Expected intAryNumeratorNumberStr = '%v'\n"+
      "  Actual intAryNumeratorNumberStr = '%v'\n\n",
      ePrefix, numeratorNumStr, intAryNumeratorNumberStr)

    return
  }

  intAryDenominator, err := new(IntAry).NewNumStr(denominatorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDenominator, err := new(IntAry).NewNumStr(denominatorNumStr)\n"+
      "denominatorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, denominatorNumStr, err.Error())
    return
  }

  err = intAryDenominator.IsValid("Validating intAryDenominator")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDenominator.IsValid('Validating intAryDenominator')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDenominatorNumberStr, err := intAryDenominator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDenominatorNumberStr, err := intAryDenominator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if denominatorNumStr != intAryDenominatorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Denominator Number String Values ARE NOT Equal\n"+
      "Because denominatorNumStr != intAryDenominatorNumberStr \n"+
      "Expected intAryDenominatorNumberStr = '%v'\n"+
      "  Actual intAryDenominatorNumberStr = '%v'\n\n",
      ePrefix, denominatorNumStr, intAryDenominatorNumberStr)

    return
  }

  fracIntAry, err := new(FracIntAry).NewIntArys(&intAryNumerator, &intAryDenominator)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewIntArys(\n"+
      "  &intAryNumerator, &intAryDenominator)\n"+
      "intAryNumerator= '%v'\n"+
      "intAryDenominator= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumeratorNumberStr,
      intAryDenominatorNumberStr,
      err.Error())

    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMaxPrecisionInt,
      err.Error())

    return
  }

  numerartorPtrInt := rationalNum.Num()

  numerartorPtrIntNumberStr := numerartorPtrInt.String()

  if numeratorNumStr != numerartorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numerator Number Strings ARE NOT Equal\n"+
      "Because numeratorNumStr != numerartorPtrIntNumberStr \n"+
      "Expected numerartorPtrIntNumberStr = '%v'\n"+
      "  Actual numerartorPtrIntNumberStr = '%v'\n\n",
      ePrefix, numeratorNumStr, numerartorPtrIntNumberStr)

    return
  }

  denominatorPtrInt := rationalNum.Denom()

  denominatorPtrIntNumberStr := denominatorPtrInt.String()

  if denominatorNumStr != denominatorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Denominator Number Strings ARE NOT Equal\n"+
      "Because denominatorNumStr != denominatorPtrIntNumberStr \n"+
      "Expected denominatorPtrIntNumberStr = '%v'\n"+
      "  Actual denominatorPtrIntNumberStr = '%v'\n\n",
      ePrefix, denominatorPtrIntNumberStr, denominatorPtrIntNumberStr)

    return
  }

  floatString := rationalNum.FloatString(originalMaxPrecisionInt)

  if expectedResultNumStr != floatString {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Floating Number String Values ARE NOT Equal\n"+
      "Because expectedResultNumStr != floatString \n"+
      "Expected floatString = '%v'\n"+
      "  Actual floatString = '%v'\n\n",
      ePrefix, expectedResultNumStr, floatString)

    return
  }

  return
}

func TestFracIntAry_GetRationalValue_07(t *testing.T) {

  ePrefix := "TestFracIntAry_GetRationalValue_07"

  numeratorNumStr := "3.1"

  denominatorNumStr := "4.2"

  //                                  1
  //                       0.123456789012
  expectedResultNumStr := "0.738095238095"

  originalMaxPrecisionInt := 12

  intAryNumerator, err := new(IntAry).NewNumStr(numeratorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumerator, err := new(IntAry).NewNumStr(numeratorNumStr)\n"+
      "numeratorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numeratorNumStr, err.Error())
    return
  }

  err = intAryNumerator.IsValid("Validating intAryNumerator")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryNumerator.IsValid('Validating intAryNumerator')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumeratorNumberStr, err := intAryNumerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumeratorNumberStr, err := intAryNumerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numeratorNumStr != intAryNumeratorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Numerator Number String Values ARE NOT Equal\n"+
      "Because numeratorNumStr != intAryNumeratorNumberStr \n"+
      "Expected intAryNumeratorNumberStr = '%v'\n"+
      "  Actual intAryNumeratorNumberStr = '%v'\n\n",
      ePrefix, numeratorNumStr, intAryNumeratorNumberStr)

    return
  }

  intAryDenominator, err := new(IntAry).NewNumStr(denominatorNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDenominator, err := new(IntAry).NewNumStr(denominatorNumStr)\n"+
      "denominatorNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, denominatorNumStr, err.Error())
    return
  }

  err = intAryDenominator.IsValid("Validating intAryDenominator")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryDenominator.IsValid('Validating intAryDenominator')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryDenominatorNumberStr, err := intAryDenominator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryDenominatorNumberStr, err := intAryDenominator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if denominatorNumStr != intAryDenominatorNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Denominator Number String Values ARE NOT Equal\n"+
      "Because denominatorNumStr != intAryDenominatorNumberStr \n"+
      "Expected intAryDenominatorNumberStr = '%v'\n"+
      "  Actual intAryDenominatorNumberStr = '%v'\n\n",
      ePrefix, denominatorNumStr, intAryDenominatorNumberStr)

    return
  }

  fracIntAry, err := new(FracIntAry).NewIntArys(&intAryNumerator, &intAryDenominator)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewIntArys(\n"+
      "  &intAryNumerator, &intAryDenominator)\n"+
      "intAryNumerator= '%v'\n"+
      "intAryDenominator= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumeratorNumberStr,
      intAryDenominatorNumberStr,
      err.Error())

    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "rationalNum, err := fracIntAry.GetRationalValue(originalMaxPrecisionInt)\n"+
      "originalMaxPrecisionInt= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalMaxPrecisionInt,
      err.Error())

    return
  }

  numerartorPtrInt := rationalNum.Num()

  numerartorPtrIntNumberStr := numerartorPtrInt.String()

  if numeratorNumStr != numerartorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numerator Number Strings ARE NOT Equal\n"+
      "Because numeratorNumStr != numerartorPtrIntNumberStr \n"+
      "Expected numerartorPtrIntNumberStr = '%v'\n"+
      "  Actual numerartorPtrIntNumberStr = '%v'\n\n",
      ePrefix, numeratorNumStr, numerartorPtrIntNumberStr)

    return
  }

  denominatorPtrInt := rationalNum.Denom()

  denominatorPtrIntNumberStr := denominatorPtrInt.String()

  if denominatorNumStr != denominatorPtrIntNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Denominator Number Strings ARE NOT Equal\n"+
      "Because denominatorNumStr != denominatorPtrIntNumberStr \n"+
      "Expected denominatorPtrIntNumberStr = '%v'\n"+
      "  Actual denominatorPtrIntNumberStr = '%v'\n\n",
      ePrefix, denominatorPtrIntNumberStr, denominatorPtrIntNumberStr)

    return
  }

  floatString := rationalNum.FloatString(originalMaxPrecisionInt)

  if expectedResultNumStr != floatString {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Floating Number String Values ARE NOT Equal\n"+
      "Because expectedResultNumStr != floatString \n"+
      "Expected floatString = '%v'\n"+
      "  Actual floatString = '%v'\n\n",
      ePrefix, expectedResultNumStr, floatString)

    return
  }

  return
}

func TestFracIntAry_NewFracIntAry_01(t *testing.T) {

  ePrefix := "TestFracIntAry_NewFracIntAry_01"

  originalNumberStr := "3.1"

  expectedNumeratorNumStr := "31"

  expectedDenominatorNumStr := "10"

  intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAryNumber.IsValid("Validating intAryNumber")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryNumber.IsValid('Validating intAryNumber')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberNumStr, err := intAryNumber.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberNumStr, err := intAryNumber.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberNumStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberNumStr \n"+
      "Expected intAryNumberNumStr = '%v'\n"+
      "  Actual intAryNumberNumStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberNumStr)

    return
  }

  fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)\n"+
      "intAryNumber= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberNumStr, err.Error())
    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumeratorNumStr != fracIntAryNumeratorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumeratorNumStr != fracIntAryNumeratorNumStr\n"+
      "Expected fracIntAryNumeratorNumStr = '%v'\n"+
      "  Actual fracIntAryNumeratorNumStr = '%v'\n\n",
      ePrefix, expectedNumeratorNumStr, fracIntAryNumeratorNumStr)

    return
  }

  if expectedDenominatorNumStr != fracIntAryDenominatorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedDenominatorNumStr != fracIntAryDenominatorNumStr\n"+
      "Expected fracIntAryDenominatorNumStr = '%v'\n"+
      "  Actual fracIntAryDenominatorNumStr = '%v'\n\n",
      ePrefix, expectedDenominatorNumStr, fracIntAryDenominatorNumStr)

    return
  }

  return
}

func TestFracIntAry_NewFracIntAry_02(t *testing.T) {

  ePrefix := "TestFracIntAry_NewFracIntAry_02"

  originalNumberStr := "50"

  expectedNumeratorNumStr := "50"

  expectedDenominatorNumStr := "1"

  intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAryNumber.IsValid("Validating intAryNumber")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryNumber.IsValid('Validating intAryNumber')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberNumStr, err := intAryNumber.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberNumStr, err := intAryNumber.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberNumStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberNumStr \n"+
      "Expected intAryNumberNumStr = '%v'\n"+
      "  Actual intAryNumberNumStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberNumStr)

    return
  }

  fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)\n"+
      "intAryNumber= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberNumStr, err.Error())
    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumeratorNumStr != fracIntAryNumeratorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumeratorNumStr != fracIntAryNumeratorNumStr\n"+
      "Expected fracIntAryNumeratorNumStr = '%v'\n"+
      "  Actual fracIntAryNumeratorNumStr = '%v'\n\n",
      ePrefix, expectedNumeratorNumStr, fracIntAryNumeratorNumStr)

    return
  }

  if expectedDenominatorNumStr != fracIntAryDenominatorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedDenominatorNumStr != fracIntAryDenominatorNumStr\n"+
      "Expected fracIntAryDenominatorNumStr = '%v'\n"+
      "  Actual fracIntAryDenominatorNumStr = '%v'\n\n",
      ePrefix, expectedDenominatorNumStr, fracIntAryDenominatorNumStr)

    return
  }

  return
}

func TestFracIntAry_NewFracIntAry_03(t *testing.T) {

  ePrefix := "TestFracIntAry_NewFracIntAry_03"

  originalNumberStr := "-50"

  expectedNumeratorNumStr := "-50"

  expectedDenominatorNumStr := "1"

  intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAryNumber.IsValid("Validating intAryNumber")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryNumber.IsValid('Validating intAryNumber')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberNumStr, err := intAryNumber.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberNumStr, err := intAryNumber.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberNumStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberNumStr \n"+
      "Expected intAryNumberNumStr = '%v'\n"+
      "  Actual intAryNumberNumStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberNumStr)

    return
  }

  fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)\n"+
      "intAryNumber= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberNumStr, err.Error())
    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumeratorNumStr != fracIntAryNumeratorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumeratorNumStr != fracIntAryNumeratorNumStr\n"+
      "Expected fracIntAryNumeratorNumStr = '%v'\n"+
      "  Actual fracIntAryNumeratorNumStr = '%v'\n\n",
      ePrefix, expectedNumeratorNumStr, fracIntAryNumeratorNumStr)

    return
  }

  if expectedDenominatorNumStr != fracIntAryDenominatorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedDenominatorNumStr != fracIntAryDenominatorNumStr\n"+
      "Expected fracIntAryDenominatorNumStr = '%v'\n"+
      "  Actual fracIntAryDenominatorNumStr = '%v'\n\n",
      ePrefix, expectedDenominatorNumStr, fracIntAryDenominatorNumStr)

    return
  }

  return
}

func TestFracIntAry_NewFracIntAry_04(t *testing.T) {

  ePrefix := "TestFracIntAry_NewFracIntAry_04"

  originalNumberStr := "50.72589"

  expectedNumeratorNumStr := "5072589"

  expectedDenominatorNumStr := "100000"

  intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAryNumber.IsValid("Validating intAryNumber")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryNumber.IsValid('Validating intAryNumber')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberNumStr, err := intAryNumber.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberNumStr, err := intAryNumber.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberNumStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberNumStr \n"+
      "Expected intAryNumberNumStr = '%v'\n"+
      "  Actual intAryNumberNumStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberNumStr)

    return
  }

  fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)\n"+
      "intAryNumber= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberNumStr, err.Error())
    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumeratorNumStr != fracIntAryNumeratorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumeratorNumStr != fracIntAryNumeratorNumStr\n"+
      "Expected fracIntAryNumeratorNumStr = '%v'\n"+
      "  Actual fracIntAryNumeratorNumStr = '%v'\n\n",
      ePrefix, expectedNumeratorNumStr, fracIntAryNumeratorNumStr)

    return
  }

  if expectedDenominatorNumStr != fracIntAryDenominatorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedDenominatorNumStr != fracIntAryDenominatorNumStr\n"+
      "Expected fracIntAryDenominatorNumStr = '%v'\n"+
      "  Actual fracIntAryDenominatorNumStr = '%v'\n\n",
      ePrefix, expectedDenominatorNumStr, fracIntAryDenominatorNumStr)

    return
  }

  return
}

func TestFracIntAry_NewFracIntAry_05(t *testing.T) {

  ePrefix := "TestFracIntAry_NewFracIntAry_05"

  originalNumberStr := "5072589"

  expectedNumeratorNumStr := "5072589"

  expectedDenominatorNumStr := "1"

  intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAryNumber.IsValid("Validating intAryNumber")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryNumber.IsValid('Validating intAryNumber')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberNumStr, err := intAryNumber.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberNumStr, err := intAryNumber.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberNumStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberNumStr \n"+
      "Expected intAryNumberNumStr = '%v'\n"+
      "  Actual intAryNumberNumStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberNumStr)

    return
  }

  fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)\n"+
      "intAryNumber= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberNumStr, err.Error())
    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumeratorNumStr != fracIntAryNumeratorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumeratorNumStr != fracIntAryNumeratorNumStr\n"+
      "Expected fracIntAryNumeratorNumStr = '%v'\n"+
      "  Actual fracIntAryNumeratorNumStr = '%v'\n\n",
      ePrefix, expectedNumeratorNumStr, fracIntAryNumeratorNumStr)

    return
  }

  if expectedDenominatorNumStr != fracIntAryDenominatorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedDenominatorNumStr != fracIntAryDenominatorNumStr\n"+
      "Expected fracIntAryDenominatorNumStr = '%v'\n"+
      "  Actual fracIntAryDenominatorNumStr = '%v'\n\n",
      ePrefix, expectedDenominatorNumStr, fracIntAryDenominatorNumStr)

    return
  }

  return
}

func TestFracIntAry_NewFracIntAry_06(t *testing.T) {

  ePrefix := "TestFracIntAry_NewFracIntAry_06"

  originalNumberStr := "-50.72589"

  expectedNumeratorNumStr := "-5072589"

  expectedDenominatorNumStr := "100000"

  intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumber, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAryNumber.IsValid("Validating intAryNumber")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryNumber.IsValid('Validating intAryNumber')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberNumStr, err := intAryNumber.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberNumStr, err := intAryNumber.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberNumStr {
    t.Errorf("%v\n"+
      "Error: Original and IntArray Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberNumStr \n"+
      "Expected intAryNumberNumStr = '%v'\n"+
      "  Actual intAryNumberNumStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberNumStr)

    return
  }

  fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAry, err := new(FracIntAry).NewFracIntAry(&intAryNumber)\n"+
      "intAryNumber= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberNumStr, err.Error())
    return
  }

  err = fracIntAry.IsValid("Validating fracIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fracIntAry.IsValid('Validating fracIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracIntAryDenominatorNumStr, err := fracIntAry.Numerator.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumeratorNumStr != fracIntAryNumeratorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumeratorNumStr != fracIntAryNumeratorNumStr\n"+
      "Expected fracIntAryNumeratorNumStr = '%v'\n"+
      "  Actual fracIntAryNumeratorNumStr = '%v'\n\n",
      ePrefix, expectedNumeratorNumStr, fracIntAryNumeratorNumStr)

    return
  }

  if expectedDenominatorNumStr != fracIntAryDenominatorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedDenominatorNumStr != fracIntAryDenominatorNumStr\n"+
      "Expected fracIntAryDenominatorNumStr = '%v'\n"+
      "  Actual fracIntAryDenominatorNumStr = '%v'\n\n",
      ePrefix, expectedDenominatorNumStr, fracIntAryDenominatorNumStr)

    return
  }

  return
}
