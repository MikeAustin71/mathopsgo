package mathops

import (
  "fmt"
  "testing"
)

func TestBigIntMathDivide_IntAryQuotientMod_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryQuotientMod_01"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555					/						 2.5			=			 5							 0.055
  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedQuoStr := "5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  dividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStr(dividendStr)\n"+
      "dividendStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedModulo.IsValid("Validating expectedModulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedModulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloNumStr, err := expectedModulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumStr, err := expectedModulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).IntAryQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryQuotientMod(d\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      dividendNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = modulo.IsValid("Validating modulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = modulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := modulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := modulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err := \n"+
      " expectedModulo.Equal(modulo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected modulo = '%v'\n"+
      "  Actual modulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  if expectedModuloNumStr != actualModuloNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloNumStr != actualModuloNumStr\n"+
      "Expected actualModuloNumStr = '%v'\n"+
      "  Actual actualModuloNumStr = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err := dividendNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := dividendNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Modulo Numeric Separators = '%v'\n"+
      "  Actual Modulo Numeric Separators = '%v'\n\n",
      ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryQuotientMod_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryQuotientMod_02"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555  	 			/ 				 	 2  			= 		 6							 0.555
  dividendStr := "12.555"
  divisorStr := "2"
  expectedQuoStr := "6"
  expectedModuloStr := "0.555"
  maxPrecision := uint(15)

  dividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStr(dividendStr)\n"+
      "dividendStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedModulo.IsValid("Validating expectedModulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedModulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloNumStr, err := expectedModulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumStr, err := expectedModulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).IntAryQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryQuotientMod(d\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      dividendNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = modulo.IsValid("Validating modulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = modulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := modulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := modulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err := \n"+
      " expectedModulo.Equal(modulo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected modulo = '%v'\n"+
      "  Actual modulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  if expectedModuloNumStr != actualModuloNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloNumStr != actualModuloNumStr\n"+
      "Expected actualModuloNumStr = '%v'\n"+
      "  Actual actualModuloNumStr = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err := dividendNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := dividendNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Modulo Numeric Separators = '%v'\n"+
      "  Actual Modulo Numeric Separators = '%v'\n\n",
      ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryQuotientMod_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryQuotientMod_03"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //	-12.555 				/ 				   2.5 			= 		-5							-0.055
  dividendStr := "-12.555"
  divisorStr := "2.5"
  expectedQuoStr := "-5"
  expectedModuloStr := "-0.055"
  maxPrecision := uint(15)

  dividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStr(dividendStr)\n"+
      "dividendStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedModulo.IsValid("Validating expectedModulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedModulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloNumStr, err := expectedModulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumStr, err := expectedModulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).IntAryQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryQuotientMod(d\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      dividendNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = modulo.IsValid("Validating modulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = modulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := modulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := modulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err := \n"+
      " expectedModulo.Equal(modulo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected modulo = '%v'\n"+
      "  Actual modulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  if expectedModuloNumStr != actualModuloNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloNumStr != actualModuloNumStr\n"+
      "Expected actualModuloNumStr = '%v'\n"+
      "  Actual actualModuloNumStr = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err := dividendNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := dividendNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Modulo Numeric Separators = '%v'\n"+
      "  Actual Modulo Numeric Separators = '%v'\n\n",
      ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryQuotientMod_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryQuotientMod_04"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //  -12.555     		/    			 	 2  			= 		-6							-0.555
  dividendStr := "-12.555"
  divisorStr := "2"
  expectedQuoStr := "-6"
  expectedModuloStr := "-0.555"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  dividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  dividendStr, expectedNumSeps)\n"+
      "dividendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendStr, expectedNumSeps.String(), err.Error())

    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedModulo.IsValid("Validating expectedModulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedModulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloNumStr, err := expectedModulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumStr, err := expectedModulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).IntAryQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryQuotientMod(d\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      expectedNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = modulo.IsValid("Validating modulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = modulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := modulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := modulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err := \n"+
      " expectedModulo.Equal(modulo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected modulo = '%v'\n"+
      "  Actual modulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  if expectedModuloNumStr != actualModuloNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloNumStr != actualModuloNumStr\n"+
      "Expected actualModuloNumStr = '%v'\n"+
      "  Actual actualModuloNumStr = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Modulo Numeric Separators = '%v'\n"+
      "  Actual Modulo Numeric Separators = '%v'\n\n",
      ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryQuotientMod_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryQuotientMod_05"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12,555  	 			/ 				 	 2  			= 		 6							 0,555
  dividendStr := "12,555"
  divisorStr := "2"
  expectedQuoStr := "6"
  expectedModuloStr := "0,555"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  dividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  dividendStr, expectedNumSeps)\n"+
      "dividendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendStr, expectedNumSeps.String(), err.Error())

    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedQuo.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedModulo.IsValid("Validating expectedModulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedModulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedModulo.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedModulo.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedModuloNumStr, err := expectedModulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumStr, err := expectedModulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).IntAryQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryQuotientMod(d\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      expectedNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = modulo.IsValid("Validating modulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = modulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := modulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := modulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err := \n"+
      " expectedModulo.Equal(modulo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected modulo = '%v'\n"+
      "  Actual modulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  if expectedModuloNumStr != actualModuloNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloNumStr != actualModuloNumStr\n"+
      "Expected actualModuloNumStr = '%v'\n"+
      "  Actual actualModuloNumStr = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Modulo Numeric Separators = '%v'\n"+
      "  Actual Modulo Numeric Separators = '%v'\n\n",
      ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryQuotientMod_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryQuotientMod_06"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555  	 			/ 				 	 2  			= 		 6							 0.555
  dividendStr := "12.555"
  divisorStr := "2"
  expectedQuoStr := "6"
  expectedModuloStr := "0.555"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  dividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  dividendStr, expectedNumSeps)\n"+
      "dividendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendStr, expectedNumSeps.String(), err.Error())

    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)\n"+
      "expectedQuoStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedQuo.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedModulo.IsValid("Validating expectedModulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedModulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloNumStr, err := expectedModulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumStr, err := expectedModulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).IntAryQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryQuotientMod(d\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      expectedNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = modulo.IsValid("Validating modulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = modulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := modulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := modulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err := \n"+
      " expectedModulo.Equal(modulo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected modulo = '%v'\n"+
      "  Actual modulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  if expectedModuloNumStr != actualModuloNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloNumStr != actualModuloNumStr\n"+
      "Expected actualModuloNumStr = '%v'\n"+
      "  Actual actualModuloNumStr = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Modulo Numeric Separators = '%v'\n"+
      "  Actual Modulo Numeric Separators = '%v'\n\n",
      ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryFracQuotient_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryFracQuotient_01"

  // Dividend		divided by		Divisor			=		Quotient
  // 	 10.5  				/ 					2 				= 	 5.25

  dividendStr := "10.5"
  divisorStr := "2"
  expectedQuoStr := "5.25"
  maxPrecision := uint(15)

  dividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStr(dividendStr)\n"+
      "dividendStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).IntAryFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryFracQuotient(d\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      dividendNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryFracQuotient_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryFracQuotient_02"

  // Dividend		divided by		Divisor			=		Quotient
  //	-12.555 			/ 					2.5 			= 		-5.022

  dividendStr := "-12.555"
  divisorStr := "2.5"
  expectedQuoStr := "-5.022"
  maxPrecision := uint(15)

  dividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStr(dividendStr)\n"+
      "dividendStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).IntAryFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryFracQuotient(d\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      dividendNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryFracQuotient_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryFracQuotient_03"

  // Dividend		divided by		Divisor			=		Quotient
  //  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

  dividendStr := "-2.5"
  divisorStr := "12.555"
  expectedQuoStr := "-0.199123855037834"
  maxPrecision := uint(15)

  dividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStr(dividendStr)\n"+
      "dividendStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).IntAryFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryFracQuotient(d\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      dividendNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := dividendNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryFracQuotient_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryFracQuotient_04"

  // Dividend		divided by		Divisor			=		Quotient
  // 	-12.555 			/ 				 -2.5 			= 	 5.022

  dividendStr := "-12.555"
  divisorStr := "-2.5"
  expectedQuoStr := "5.022"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  dividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  dividendStr, expectedNumSeps)\n"+
      "dividendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, dividendStr, expectedNumSeps.String(), err.Error())

    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedQuo.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).IntAryFracQuotient(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryFracQuotient(d\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      expectedNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryFracQuotient_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryFracQuotient_05"

  // Dividend		divided by		Divisor			=		Quotient
  // 	 10.5  				/ 					2 				= 	 5.25

  dividendStr := "10.5"
  divisorStr := "2"
  expectedQuoStr := "5.25"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  dividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
      "dividendStr,expectedNumSeps). "+
      "dividendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      dividendStr, expectedNumSeps, err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = expectedQuo.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).IntAryFracQuotient(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryFracQuotient(d\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      expectedNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryFracQuotient_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryFracQuotient_06"

  // Dividend		divided by		Divisor			=		Quotient
  // 	 10,5  				/ 					2 				= 	 5,25

  dividendStr := "10,5"
  divisorStr := "2"
  expectedQuoStr := "5,25"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  dividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
      "dividendStr,expectedNumSeps). "+
      "dividendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      dividendStr, expectedNumSeps, err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
      "expectedQuoStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedQuoStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).IntAryFracQuotient(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryFracQuotient(d\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      expectedNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryFracQuotient_07(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryFracQuotient_06"

  // Dividend		divided by		Divisor			=		Quotient
  // 	 10.5  				/ 					2 				= 	 5.25

  dividendStr := "10.5"
  divisorStr := "2"
  expectedQuoStr := "5.25"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  dividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
      "dividendStr,expectedNumSeps). "+
      "dividendStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      dividendStr, expectedNumSeps, err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).\n"+
      "  NewNumStr(divisorStr).\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, err.Error())
    return
  }

  err = divisor.IsValid(ePrefix + "\nValidating divisor")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = divisor.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
      "expectedQuoStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedQuoStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedQuo.IsValid("Validating expectedQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).IntAryFracQuotient(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryFracQuotient(d\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      expectedNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid("Validating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(quotient)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "  Actual actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "quotient= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualQuotientNumStr, err.Error())

    return
  }

  expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())

    return
  }

  if !expectedQuoNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient Numeric Separators = '%v'\n"+
      "  Actual Quotient Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_IntAryFracQuotientArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_IntAryFracQuotientArray_01"

  divisorStr := "2.5"
  maxPrecision := uint(15)

  dividendArrayStr := []string{
    "10.5",
    "10",
    "11.5",
    "2.5",
    "-12.555",
    "-2.5",
    "12.555",
    "-122.783",
    "-6847.231",
    "-2.5",
    "-10",
    "-10.5",
  }

  expectedArrayStr := []string{
    "4.2",
    "4",
    "4.6",
    "1",
    "-5.022",
    "-1",
    "5.022",
    "-49.1132",
    "-2738.8924",
    "-1",
    "-4",
    "-4.2",
  }

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  lenDividends := len(dividendArrayStr)

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisor='%v' Error='%v' ",
      divisorStr, err.Error())
  }

  dividends := make([]IntAry, lenDividends)
  expectedResults := make([]IntAry, lenDividends)

  for i := 0; i < lenDividends; i++ {

    dividends[i], err = new(IntAry).NewNumStr(dividendArrayStr[i])

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStr(dividendArrayStr[i]). "+
        "dividendArrayStr[%v]='%v' Error='%v' ",
        i, dividendArrayStr[i], err.Error())
    }

    expectedResults[i], err = new(IntAry).NewNumStr(expectedArrayStr[i])

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStr(expectedArrayStr[i]). "+
        "expectedArrayStr[%v]='%v' Error='%v' ",
        i, expectedArrayStr[i], err.Error())
    }

  }

  resultArray, err := new(BigIntMathDivide).IntAryFracQuotientArray(dividends, divisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryFracQuotientArray"+
      "(dividends, divisor, maxPrecision ). "+
      "divisor='%v' maxPrecision='%v' Error='%v' ",
      divisor.GetNumStr(), maxPrecision, err.Error())
  }

  lenResultArray := len(resultArray)

  if lenDividends != lenResultArray {
    t.Errorf("Error: Expected Results Array Length='%v'. Actual Array Length='%v'.",
      lenDividends, lenResultArray)
  }

  for k := 0; k < lenDividends; k++ {

    resultEqualsExpectedResult, err := resultArray[k].Equal(&expectedResults[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultEqualsExpectedResult, err = \n"+
        "  resultArray[k].Equal(&expectedResults[k])\n"+
        fmt.Sprintf("'k' index='%v'\n", k)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return

    }

    expectedResultsNumStr, err := expectedResults[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("expectedResultsNumStr, err := \n"+
          "  expectedResults[%d].GetNumStr()\n", k)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    resultsNumStr, err := resultArray[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("resultsNumStr, err :=\n"+
          "  resultArray[%d].GetNumStr()\n", k)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    if !resultEqualsExpectedResult {
      t.Errorf("Error: Expected Result NOT Equal to Actual Result!\n"+
        "Expected Value='%v'.\n"+
        "Actual Value='%v'\n"+
        "k='%v'\n\n",
        expectedResultsNumStr, resultsNumStr, k)
    }

    actualNumSeps, err := resultArray[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("actualNumSeps, err := resultArray[%d].GetNumericSeparatorsDto()\n", k)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
        expectedNumSeps.String(), actualNumSeps.String())
      return
    }

  }

  return
}

func TestBigIntMathDivide_IntAryFracQuotientArray_02(t *testing.T) {

  divisorStr := "2.5"
  maxPrecision := uint(15)

  dividendArrayStr := []string{
    "10.5",
    "10",
    "11.5",
    "2.5",
    "-12.555",
    "-2.5",
    "12.555",
    "-122.783",
    "-6847.231",
    "-2.5",
    "-10",
    "-10.5",
  }

  expectedArrayStr := []string{
    "4,2",
    "4",
    "4,6",
    "1",
    "-5,022",
    "-1",
    "5,022",
    "-49,1132",
    "-2738,8924",
    "-1",
    "-4",
    "-4,2",
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  lenDividends := len(dividendArrayStr)

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisor='%v' Error='%v' ",
      divisorStr, err.Error())
  }

  dividends := make([]IntAry, lenDividends)
  expectedResults := make([]IntAry, lenDividends)

  for i := 0; i < lenDividends; i++ {

    dividends[i], err = new(IntAry).NewNumStr(dividendArrayStr[i])

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStr(dividendArrayStr[i]). "+
        "dividendArrayStr[%v]='%v' Error='%v' ",
        i, dividendArrayStr[i], err.Error())
    }

    expectedResults[i], err = new(IntAry).NewNumStrWithNumSeps(expectedArrayStr[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps"+
        "(expectedArrayStr[i], expectedNumSeps). "+
        "expectedArrayStr[%v]='%v' Error='%v' ",
        i, expectedArrayStr[i], err.Error())
    }

  }

  err = dividends[0].SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by dividends[0].SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  resultArray, err := new(BigIntMathDivide).IntAryFracQuotientArray(dividends, divisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryFracQuotientArray"+
      "(dividends, divisor, maxPrecision ). "+
      "divisor='%v' maxPrecision='%v' Error='%v' ",
      divisor.GetNumStr(), maxPrecision, err.Error())
  }

  lenResultArray := len(resultArray)

  if lenDividends != lenResultArray {
    t.Errorf("Error: Expected Results Array Length='%v'. Actual Array Length='%v'.",
      lenDividends, lenResultArray)
  }

  for k := 0; k < lenDividends; k++ {

    if !resultArray[k].Equals(&expectedResults[k]) {
      t.Errorf("Error: Expected Result NOT Equal to Actual Result! "+
        "Expected Value='%v'. Actual Value='%v' k='%v'",
        expectedResults[k].GetNumStr(), resultArray[k].GetNumStr(), k)
    }

    actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
        expectedNumSeps.String(), actualNumSeps.String())
    }

  }
}

func TestBigIntMathDivide_IntAryFracQuotientArray_03(t *testing.T) {

  divisorStr := "2.5"
  maxPrecision := uint(15)

  dividendArrayStr := []string{
    "10.5",
    "10",
    "11.5",
    "2.5",
    "-12.555",
    "-2.5",
    "12.555",
    "-122.783",
    "-6847.231",
    "-2.5",
    "-10",
    "-10.5",
  }

  expectedArrayStr := []string{
    "4.2",
    "4",
    "4.6",
    "1",
    "-5.022",
    "-1",
    "5.022",
    "-49.1132",
    "-2738.8924",
    "-1",
    "-4",
    "-4.2",
  }

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  lenDividends := len(dividendArrayStr)

  divisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisor='%v' Error='%v' ",
      divisorStr, err.Error())
  }

  dividends := make([]IntAry, lenDividends)
  expectedResults := make([]IntAry, lenDividends)

  for i := 0; i < lenDividends; i++ {

    dividends[i], err = new(IntAry).NewNumStr(dividendArrayStr[i])

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStr(dividendArrayStr[i]). "+
        "dividendArrayStr[%v]='%v' Error='%v' ",
        i, dividendArrayStr[i], err.Error())
    }

    expectedResults[i], err = new(IntAry).NewNumStrWithNumSeps(expectedArrayStr[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps"+
        "(expectedArrayStr[i], expectedNumSeps). "+
        "expectedArrayStr[%v]='%v' Error='%v' ",
        i, expectedArrayStr[i], err.Error())
    }

  }

  err = dividends[0].SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by dividends[0].SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  resultArray, err := new(BigIntMathDivide).IntAryFracQuotientArray(dividends, divisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryFracQuotientArray"+
      "(dividends, divisor, maxPrecision ). "+
      "divisor='%v' maxPrecision='%v' Error='%v' ",
      divisor.GetNumStr(), maxPrecision, err.Error())
  }

  lenResultArray := len(resultArray)

  if lenDividends != lenResultArray {
    t.Errorf("Error: Expected Results Array Length='%v'. Actual Array Length='%v'.",
      lenDividends, lenResultArray)
  }

  for k := 0; k < lenDividends; k++ {

    if !resultArray[k].Equals(&expectedResults[k]) {
      t.Errorf("Error: Expected Result NOT Equal to Actual Result! "+
        "Expected Value='%v'. Actual Value='%v' k='%v'",
        expectedResults[k].GetNumStr(), resultArray[k].GetNumStr(), k)
    }

    actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
        expectedNumSeps.String(), actualNumSeps.String())
    }

  }
}

func TestBigIntMathDivide_IntAryModulo_01(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12.555					%						 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  iaDividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).IntAryModulo(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModulo(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

}

func TestBigIntMathDivide_IntAryModulo_02(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   -12.555 				% 				 - 2.5 			= 		-0.055

  dividendStr := "-12.555"
  divisorStr := "-2.5"
  expectedModuloStr := "-0.055"
  maxPrecision := uint(15)

  iaDividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).IntAryModulo(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModulo(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

}

func TestBigIntMathDivide_IntAryModulo_03(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12.555					% 				 - 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "-2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  iaDividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).IntAryModulo(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModulo(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

}

func TestBigIntMathDivide_IntAryModulo_04(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //    2.5 				  % 				 -12.555		= 		 2.5

  dividendStr := "2.5"
  divisorStr := "-12.555"
  expectedModuloStr := "2.5"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  iaDividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).IntAryModulo(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModulo(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected moduloBINum='%v'. Instead moduloBINum='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathDivide_IntAryModulo_05(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12,555					%						 2.5			=			 0,055

  dividendStr := "12,555"
  divisorStr := "2.5"
  expectedModuloStr := "0,055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  iaDividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
      "dividendStr, expectedNumSeps). "+
      "dividendStr='%v' Error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).IntAryModulo(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModulo(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathDivide_IntAryModulo_06(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12.555					%						 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  iaDividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
      "dividendStr, expectedNumSeps). "+
      "dividendStr='%v' Error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).IntAryModulo(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModulo(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathDivide_IntAryModuloToIntAry_01(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12.555					%						 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  iaDividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  iaModulo, err := new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := iaModulo.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

}

func TestBigIntMathDivide_IntAryModuloToIntAry_02(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   -12.555 				% 				 - 2.5 			= 		-0.055

  dividendStr := "-12.555"
  divisorStr := "-2.5"
  expectedModuloStr := "-0.055"
  maxPrecision := uint(15)

  iaDividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  iaModulo, err := new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := iaModulo.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

}

func TestBigIntMathDivide_IntAryModuloToIntAry_03(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12.555					% 				 - 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "-2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  iaDividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  iaModulo, err := new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := iaModulo.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

}

func TestBigIntMathDivide_IntAryModuloToIntAry_04(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //    2.5 				  % 				 -12.555		= 		 2.5

  dividendStr := "2.5"
  divisorStr := "-12.555"
  expectedModuloStr := "2.5"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  iaDividend, err := new(IntAry).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  iaModulo, err := new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := iaModulo.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected iaModulo='%v'. Instead iaModulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := iaModulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathDivide_IntAryModuloToIntAry_05(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12,555					%						 2.5			=			 0,055

  dividendStr := "12,555"
  divisorStr := "2.5"
  expectedModuloStr := "0,055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  iaDividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
      "dividendStr, expectedNumSeps). "+
      "dividendStr='%v' Error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  iaModulo, err := new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := iaModulo.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := iaModulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathDivide_IntAryModuloToIntAry_06(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12.555					%						 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  iaDividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
      "dividendStr, expectedNumSeps). "+
      "dividendStr='%v' Error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  iaModulo, err := new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).IntAryModuloToIntAry(iaDividend, "+
      "iaDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := iaModulo.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := iaModulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}
