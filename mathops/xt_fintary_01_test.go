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
  numStrNum := "9.24"
  eNum := "291115311909262759924385633270321361058601"
  numStrDenom := "15.87"
  eDenom := "500000000000000000000000000000000000000000"
  eFrac := "0.582230623818525519848771266540642722117202"
  precision := 42
  fIary, _ := FracIntAry{}.NewNumStrs(numStrNum, numStrDenom)

  ratNum, err := fIary.GetRationalValue(precision)

  if err != nil {
    t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
  }

  num := ratNum.Num()

  if eNum != num.String() {
    t.Errorf("Expected Numerator= %v . Instead, Numerator= %v  .", eNum, num.String())
  }

  denom := ratNum.Denom()

  if eDenom != denom.String() {
    t.Errorf("Expected Denominator= %v . Instead, Denominator= %v  .", numStrDenom, denom.String())
  }

  fStr := ratNum.FloatString(precision)

  if eFrac != fStr {
    t.Errorf("Expected %v-digit decimal string = %v  .   Instead, decimal string = %v .", precision, eFrac, fStr)
  }

}

func TestFracIntAry_GetRationalValue_05(t *testing.T) {
  numStrNum := "3"
  numStrDenom := "4"
  eFrac := "0.75"
  precision := 2
  iaNum, _ := IntAry{}.NewNumStr(numStrNum)

  iaDenom, _ := IntAry{}.NewNumStr(numStrDenom)

  fIary := FracIntAry{}.NewIntArys(&iaNum, &iaDenom)

  ratNum, err := fIary.GetRationalValue(precision)

  if err != nil {
    t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
  }

  num := ratNum.Num()

  if numStrNum != num.String() {
    t.Errorf("Expected Numerator= %v . Instead, Numerator= %v  .", numStrNum, num.String())
  }

  denom := ratNum.Denom()

  if numStrDenom != denom.String() {
    t.Errorf("Expected Denominator= %v . Instead, Denominator= %v  .", numStrDenom, denom.String())
  }

  fStr := ratNum.FloatString(precision)

  if eFrac != fStr {
    t.Errorf("Expected 32-digit decimal string = %v  .   Instead, decimal string = %v .", eFrac, fStr)
  }

}

func TestFracIntAry_GetRationalValue_06(t *testing.T) {
  numStrNum := "-3"
  numStrDenom := "4"
  eFrac := "-0.75"
  precision := 2
  iaNum, _ := IntAry{}.NewNumStr(numStrNum)

  iaDenom, _ := IntAry{}.NewNumStr(numStrDenom)

  fIary := FracIntAry{}.NewIntArys(&iaNum, &iaDenom)

  ratNum, err := fIary.GetRationalValue(precision)

  if err != nil {
    t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
  }

  num := ratNum.Num()

  if numStrNum != num.String() {
    t.Errorf("Expected Numerator= %v . Instead, Numerator= %v  .", numStrNum, num.String())
  }

  denom := ratNum.Denom()

  if numStrDenom != denom.String() {
    t.Errorf("Expected Denominator= %v . Instead, Denominator= %v  .", numStrDenom, denom.String())
  }

  fStr := ratNum.FloatString(precision)

  if eFrac != fStr {
    t.Errorf("Expected 32-digit decimal string = %v  .   Instead, decimal string = %v .", eFrac, fStr)
  }

}

func TestFracIntAry_GetRationalValue_07(t *testing.T) {
  numStrNum := "3.1"
  numStrDenom := "4.2"
  eFrac := "0.738095238095"
  precision := 12
  iaNum, _ := IntAry{}.NewNumStr(numStrNum)

  iaDenom, _ := IntAry{}.NewNumStr(numStrDenom)

  fIary := FracIntAry{}.NewIntArys(&iaNum, &iaDenom)

  ratNum, err := fIary.GetRationalValue(-1)

  if err != nil {
    t.Errorf("Error returned from fIary.GetRationalValue() - Error= %v", err)
  }

  fStr := ratNum.FloatString(precision)

  if eFrac != fStr {
    t.Errorf("Expected 32-digit decimal string = %v  .   Instead, decimal string = %v .", eFrac, fStr)
  }

}

func TestFracIntAry_NewFracIntAry_01(t *testing.T) {

  numStr := "3.1"
  expectedNumerator := "31"
  expectedDenominator := "10"

  iaNum, err := IntAry{}.NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(numStr). "+
      "numStr='%v' Error='%v' ", numStr, err.Error())
  }

  fracIa := FracIntAry{}.NewFracIntAry(&iaNum)

  if expectedNumerator != fracIa.Numerator.GetNumStr() {
    t.Errorf("Error: Expected Numerator='%v'. Instead, Numerator='%v'. ",
      expectedNumerator, fracIa.Numerator.GetNumStr())
  }

  if expectedDenominator != fracIa.Denominator.GetNumStr() {
    t.Errorf("Error: Expected Denominator='%v'. Instead, Denominator='%v'. ",
      expectedDenominator, fracIa.Denominator.GetNumStr())
  }

}

func TestFracIntAry_NewFracIntAry_02(t *testing.T) {

  numStr := "50"
  expectedNumerator := "50"
  expectedDenominator := "1"

  iaNum, err := IntAry{}.NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(numStr). "+
      "numStr='%v' Error='%v' ", numStr, err.Error())
  }

  fracIa := FracIntAry{}.NewFracIntAry(&iaNum)

  if expectedNumerator != fracIa.Numerator.GetNumStr() {
    t.Errorf("Error: Expected Numerator='%v'. Instead, Numerator='%v'. ",
      expectedNumerator, fracIa.Numerator.GetNumStr())
  }

  if expectedDenominator != fracIa.Denominator.GetNumStr() {
    t.Errorf("Error: Expected Denominator='%v'. Instead, Denominator='%v'. ",
      expectedDenominator, fracIa.Denominator.GetNumStr())
  }

}

func TestFracIntAry_NewFracIntAry_03(t *testing.T) {

  numStr := "-50"
  expectedNumerator := "-50"
  expectedDenominator := "1"

  iaNum, err := IntAry{}.NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(numStr). "+
      "numStr='%v' Error='%v' ", numStr, err.Error())
  }

  fracIa := FracIntAry{}.NewFracIntAry(&iaNum)

  if expectedNumerator != fracIa.Numerator.GetNumStr() {
    t.Errorf("Error: Expected Numerator='%v'. Instead, Numerator='%v'. ",
      expectedNumerator, fracIa.Numerator.GetNumStr())
  }

  if expectedDenominator != fracIa.Denominator.GetNumStr() {
    t.Errorf("Error: Expected Denominator='%v'. Instead, Denominator='%v'. ",
      expectedDenominator, fracIa.Denominator.GetNumStr())
  }

}

func TestFracIntAry_NewFracIntAry_04(t *testing.T) {

  numStr := "50.72589"
  expectedNumerator := "5072589"
  expectedDenominator := "100000"

  iaNum, err := IntAry{}.NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(numStr). "+
      "numStr='%v' Error='%v' ", numStr, err.Error())
  }

  fracIa := FracIntAry{}.NewFracIntAry(&iaNum)

  if expectedNumerator != fracIa.Numerator.GetNumStr() {
    t.Errorf("Error: Expected Numerator='%v'. Instead, Numerator='%v'. ",
      expectedNumerator, fracIa.Numerator.GetNumStr())
  }

  if expectedDenominator != fracIa.Denominator.GetNumStr() {
    t.Errorf("Error: Expected Denominator='%v'. Instead, Denominator='%v'. ",
      expectedDenominator, fracIa.Denominator.GetNumStr())
  }

}

func TestFracIntAry_NewFracIntAry_05(t *testing.T) {

  numStr := "5072589"
  expectedNumerator := "5072589"
  expectedDenominator := "1"

  iaNum, err := IntAry{}.NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(numStr). "+
      "numStr='%v' Error='%v' ", numStr, err.Error())
  }

  fracIa := FracIntAry{}.NewFracIntAry(&iaNum)

  if expectedNumerator != fracIa.Numerator.GetNumStr() {
    t.Errorf("Error: Expected Numerator='%v'. Instead, Numerator='%v'. ",
      expectedNumerator, fracIa.Numerator.GetNumStr())
  }

  if expectedDenominator != fracIa.Denominator.GetNumStr() {
    t.Errorf("Error: Expected Denominator='%v'. Instead, Denominator='%v'. ",
      expectedDenominator, fracIa.Denominator.GetNumStr())
  }

}

func TestFracIntAry_NewFracIntAry_06(t *testing.T) {

  numStr := "-50.72589"
  expectedNumerator := "-5072589"
  expectedDenominator := "100000"

  iaNum, err := IntAry{}.NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by IntAry{}.NewNumStr(numStr). "+
      "numStr='%v' Error='%v' ", numStr, err.Error())
  }

  fracIa := FracIntAry{}.NewFracIntAry(&iaNum)

  if expectedNumerator != fracIa.Numerator.GetNumStr() {
    t.Errorf("Error: Expected Numerator='%v'. Instead, Numerator='%v'. ",
      expectedNumerator, fracIa.Numerator.GetNumStr())
  }

  if expectedDenominator != fracIa.Denominator.GetNumStr() {
    t.Errorf("Error: Expected Denominator='%v'. Instead, Denominator='%v'. ",
      expectedDenominator, fracIa.Denominator.GetNumStr())
  }

}
