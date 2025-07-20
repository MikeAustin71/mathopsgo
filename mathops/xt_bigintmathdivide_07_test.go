package mathops

import "testing"

func TestBigIntMathDivide_INumMgrQuotientMod_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrQuotientMod_01"

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

  divisor, err := new(Decimal).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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
    new(BigIntMathDivide).INumMgrQuotientMod(&dividend, &divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryQuotientMod(\n"+
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

func TestBigIntMathDivide_INumMgrQuotientMod_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrQuotientMod_02"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555  	 			/ 				 	 2  			= 		 6							 0.555
  dividendStr := "12.555"
  divisorStr := "2"
  expectedQuoStr := "6"
  expectedModuloStr := "0.555"
  maxPrecision := uint(15)

  dividend, err := new(BigIntNum).NewNumStr(dividendStr)

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

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "divisor, err := new(NumStrDto).NewNumStr(divisorStr)\n"+
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
    new(BigIntMathDivide).INumMgrQuotientMod(&dividend, &divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryQuotientMod(\n"+
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

func TestBigIntMathDivide_INumMgrQuotientMod_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrQuotientMod_03"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
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

  divisor, err := new(Decimal).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "divisor, err := new(Decimal).NewNumStr(divisorStr).\n"+
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
    new(BigIntMathDivide).INumMgrQuotientMod(&dividend, &divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).IntAryQuotientMod(\n"+
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

func TestBigIntMathDivide_INumMgrQuotientMod_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrQuotientMod_04"

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

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
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

  divisor, err := new(BigIntNum).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(BigIntNum).NewNumStrWithNumSeps(divisorStr, usaNumSeps)\n"+
      "divisorStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, usaNumSeps.String(), err.Error())
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
      "  expectedQuoStr, expectedNumSeps)\n"+
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

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloStr, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())

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
    new(BigIntMathDivide).INumMgrQuotientMod(&dividend, &divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).INumMgrQuotientMod(\n"+
      "  &dividend, &divisor, expectedNumSeps, maxPrecision)\n"+
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

func TestBigIntMathDivide_INumMgrQuotientMod_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrQuotientMod_05"

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

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

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

  divisor, err := new(BigIntNum).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(BigIntNum).NewNumStrWithNumSeps(divisorStr, usaNumSeps)\n"+
      "divisorStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, usaNumSeps.String(), err.Error())
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
      "  expectedQuoStr, expectedNumSeps)\n"+
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

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloStr, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())

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
    new(BigIntMathDivide).INumMgrQuotientMod(&dividend, &divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).INumMgrQuotientMod(\n"+
      "  &dividend, &divisor, expectedNumSeps, maxPrecision)\n"+
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

func TestBigIntMathDivide_INumMgrQuotientMod_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrQuotientMod_06"

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

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

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

  divisor, err := new(BigIntNum).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(BigIntNum).NewNumStrWithNumSeps(divisorStr, usaNumSeps)\n"+
      "divisorStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, usaNumSeps.String(), err.Error())
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
      "  expectedQuoStr, expectedNumSeps)\n"+
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

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloStr, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())

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
    new(BigIntMathDivide).INumMgrQuotientMod(&dividend, &divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).INumMgrQuotientMod(\n"+
      "  &dividend, &divisor, expectedNumSeps, maxPrecision)\n"+
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

func TestBigIntMathDivide_INumMgrFracQuotient_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrFracQuotient_01"

  // Dividend		divided by		Divisor			=		Quotient
  // 	 10.5  				/ 					2 				= 	 5.25

  dividendStr := "10.5"
  divisorStr := "2"
  expectedQuoStr := "5.25"
  maxPrecision := uint(15)

  dividend, err := new(Decimal).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
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

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "divisor, err := new(NumStrDto).NewNumStr(divisorStr)\n"+
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
    new(BigIntMathDivide).INumMgrFracQuotient(&dividend, &divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).INumMgrFracQuotient(\n"+
      "  &dividend, &divisor, dividendNumSeps, maxPrecision)\n"+
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
func TestBigIntMathDivide_INumMgrFracQuotient_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrFracQuotient_02"

  // Dividend		divided by		Divisor			=		Quotient
  //	-12.555 			/ 					2.5 			= 		-5.022

  dividendStr := "-12.555"
  divisorStr := "2.5"
  expectedQuoStr := "-5.022"
  maxPrecision := uint(15)

  dividend, err := new(NumStrDto).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).NewNumStr(dividendStr)\n"+
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
    t.Errorf("Error returned by:\n"+
      "divisor, err := new(IntAry).NewNumStr(divisorStr)\n"+
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
    new(BigIntMathDivide).INumMgrFracQuotient(&dividend, &divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).INumMgrFracQuotient(\n"+
      "  &dividend, &divisor, dividendNumSeps, maxPrecision)\n"+
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

func TestBigIntMathDivide_INumMgrFracQuotient_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrFracQuotient_03"

  // Dividend		divided by		Divisor			=		Quotient
  //  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

  dividendStr := "-2.5"
  divisorStr := "12.555"
  expectedQuoStr := "-0.199123855037834"
  maxPrecision := uint(15)

  dividend, err := new(BigIntNum).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
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

  divisor, err := new(Decimal).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "divisor, err := new(IntAry).NewNumStr(divisorStr)\n"+
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
    new(BigIntMathDivide).INumMgrFracQuotient(&dividend, &divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).INumMgrFracQuotient(\n"+
      "  &dividend, &divisor, dividendNumSeps, maxPrecision)\n"+
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

func TestBigIntMathDivide_INumMgrFracQuotient_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrFracQuotient_04"

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

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  dividend, err := new(IntAry).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  dividendStr, expectedNumSeps)\n"+
      "dividendStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
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
      "Because expectedNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(BigIntNum).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisor, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  divisorStr, usaNumSeps)\n"+
      "divisorStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, usaNumSeps.String(), err.Error())
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
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr &expectedNumSeps)\n"+
      "expectedQuoStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", expectedQuoStr, expectedNumSeps.String(), ePrefix, err.Error())
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
    new(BigIntMathDivide).INumMgrFracQuotient(&dividend, &divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).INumMgrFracQuotient(\n"+
      "  &dividend, &divisor, expectedNumSeps, maxPrecision)\n"+
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

func TestBigIntMathDivide_INumMgrFracQuotient_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrFracQuotient_05"

  // Dividend		divided by		Divisor			=		Quotient
  //   11,5  				/           2.5				=  	 4,6
  dividendStr := "11,5"
  divisorStr := "2.5"
  expectedQuoStr := "4,6"
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

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  dividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(Decimal).NewNumStrWithNumSeps(\n"+
      "  dividendStr, expectedNumSeps)\n"+
      "dividendStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
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
      "Because expectedNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisor, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  divisorStr, usaNumSeps)\n"+
      "divisorStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, usaNumSeps.String(), err.Error())
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
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr &expectedNumSeps)\n"+
      "expectedQuoStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", expectedQuoStr, expectedNumSeps.String(), ePrefix, err.Error())
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
    new(BigIntMathDivide).INumMgrFracQuotient(&dividend, &divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).INumMgrFracQuotient(\n"+
      "  &dividend, &divisor, expectedNumSeps, maxPrecision)\n"+
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

func TestBigIntMathDivide_INumMgrFracQuotient_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrFracQuotient_06"

  // Dividend		divided by		Divisor			=		Quotient
  //   11.5  				/           2.5				=  	 4.6
  dividendStr := "11.5"
  divisorStr := "2.5"
  expectedQuoStr := "4.6"
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

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  dividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(Decimal).NewNumStrWithNumSeps(\n"+
      "  dividendStr, expectedNumSeps)\n"+
      "dividendStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
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
      "Because expectedNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  divisorStr, &usaNumSeps)\n"+
      "divisorStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, usaNumSeps.String(), err.Error())
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
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr &expectedNumSeps)\n"+
      "expectedQuoStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", expectedQuoStr, expectedNumSeps.String(), ePrefix, err.Error())
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
    new(BigIntMathDivide).INumMgrFracQuotient(&dividend, &divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).INumMgrFracQuotient(\n"+
      "  &dividend, &divisor, expectedNumSeps, maxPrecision)\n"+
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

func TestBigIntMathDivide_INumMgrFracQuotientArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_INumMgrFracQuotient_06"

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

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  lenDividends := len(dividendArrayStr)

  divisor, err := new(IntAry).NewNumStrWithNumSeps(divisorStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaDivisor, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  divisorStr, usaNumSeps)\n"+
      "divisorStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, usaNumSeps.String(), err.Error())
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

  dividends := make([]INumMgr, lenDividends)

  expectedResults := make([]BigIntNum, lenDividends)

  var dec Decimal

  for i := 0; i < lenDividends; i++ {

    dec, err = new(Decimal).NewNumStrWithNumSeps(dividendArrayStr[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "dec, err = new(Decimal).NewNumStrWithNumSeps(\n"+
        "  dividendArrayStr[%d], usaNumSeps)\n"+
        "dividendArrayStr[%v]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, dividendArrayStr[i], usaNumSeps.String(), err.Error())
      return
    }

    dividends[i] = &dec

    expectedResults[i], err = new(BigIntNum).NewNumStr(expectedArrayStr[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        " new(BigIntNum).NewNumStrWithNumSeps(\n"+
        "  expectedArrayStr[%d], expectedNumSeps)\n"+
        "expectedArrayStr[%v]='%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedArrayStr[i],
        expectedNumSeps.String(), err.Error())
      return
    }

  }

  resultArray, err := new(BigIntMathDivide).INumMgrFracQuotientArray(dividends, &divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultArray, err := new(BigIntMathDivide).\n"+
      "  INumMgrFracQuotientArray(dividends, &divisor,\n"+
      "    expectedNumSeps, maxPrecision)\n"+
      "divisor= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorNumStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  lenResultArray := len(resultArray)

  if lenDividends != lenResultArray {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenDividends != lenResultArray\n"+
      "Expected Length of Results Array = '%v'\n"+
      "  Actual Length of Results Array = '%v'\n\n",
      ePrefix, lenDividends, lenResultArray)

    return
  }

  var actualEqualsExpectedResults bool

  var resultsArrayNumStr, expectedResultsNumStr string

  var resultsBigINum, expectedResultsBigINum BigIntNum

  var resultsNumSeps NumericSeparatorDto

  for k := 0; k < lenDividends; k++ {

    resultsArrayNumStr, err = resultArray[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultsArrayNumStr, err = resultArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, k, err.Error())
      return
    }

    resultsBigINum, err = resultArray[k].GetBigIntNum()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultsBigINum, err = resultArray[%d].GetBigIntNum()\n"+
        "Error='%v'\n\n",
        ePrefix, k, err.Error())
      return
    }

    expectedResultsNumStr, err = expectedResults[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResults[%d].GetNumStr()\n"+
        "resultsArrayNumStr= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, k, resultsArrayNumStr, err.Error())
      return
    }

    expectedResultsBigINum, err = expectedResults[k].GetBigIntNum()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsBigINum, err = expectedResults[%d].GetBigIntNum()\n"+
        "expectedResultsNumStr= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, k, expectedResultsNumStr, err.Error())
      return
    }

    actualEqualsExpectedResults, err = expectedResultsBigINum.Equal(resultsBigINum)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualEqualsExpectedResults, err = expectedResultsBigINum.Equal(resultsBigINum)\n"+
        "results[%d]= '%v'\n"+
        "expectedResults[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        k,
        resultsArrayNumStr,
        k,
        expectedResultsNumStr,
        err.Error())

      return
    }

    if !actualEqualsExpectedResults {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because actualEqualsExpectedResults == 'false'\n"+
        "Cycle 'k' Value = '%v'\n"+
        "Expected Results Value = '%v'\n"+
        "  Actual Results Value = '%v'\n\n",
        ePrefix, k, expectedResultsNumStr, resultsArrayNumStr)

      return
    }

    if expectedResultsNumStr != resultsArrayNumStr {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedResultsNumStr != resultsArrayNumStr\n"+
        "Expected resultsArrayNumStr = '%v'\n"+
        "  Actual resultsArrayNumStr = '%v'\n"+
        "Cycle k = '%d'\n\n",
        ePrefix, expectedResultsNumStr, resultsArrayNumStr, k)

      return
    }

    resultsNumSeps, err = resultArray[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultsNumSeps, err = resultArray[%d].GetNumericSeparatorsDto()\n"+
        "expectedResultsNumStr= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, k, expectedResultsNumStr, err.Error())
      return
    }

    if !expectedNumSeps.Equal(resultsNumSeps) {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumSeps != resultsNumSeps\n"+
        "Expected resultsNumSeps = '%v'\n"+
        "  Actual resultsNumSeps = '%v'\n"+
        "Cycle k = '%d'\n\n",
        ePrefix, expectedNumSeps.String(), resultsNumSeps.String(), k)

      return
    }
  }

  return
}

func TestBigIntMathDivide_INumMgrFracQuotientArray_02(t *testing.T) {

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
    t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
      "divisor='%v' Error='%v' ",
      divisorStr, err.Error())
  }

  dividends := make([]INumMgr, lenDividends)
  expectedResults := make([]BigIntNum, lenDividends)

  for i := 0; i < lenDividends; i++ {

    dec, err := new(Decimal).NewNumStr(dividendArrayStr[i])

    if err != nil {
      t.Errorf("Error returned by new(Decimal).NewNumStr(dividendArrayStr[i]). "+
        "dividendArrayStr[%v]='%v' Error='%v' ",
        i, dividendArrayStr[i], err.Error())
    }

    dividends[i] = &dec

    expectedResults[i], err = new(BigIntNum).NewNumStrWithNumSeps(expectedArrayStr[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps"+
        "(expectedArrayStr[i], expectedNumSeps). "+
        "expectedArrayStr[%v]='%v' Error='%v' ",
        i, expectedArrayStr[i], err.Error())
    }

  }

  err = dividends[0].SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by dividends[0].SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v'", err.Error())
  }

  resultArray, err := new(BigIntMathDivide).INumMgrFracQuotientArray(dividends, &divisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by INumMgrFracQuotientArray{}.INumMgrFracQuotientArray"+
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

    expectedNumStr := expectedResults[k].GetNumStr()
    actualNumStr := resultArray[k].GetNumStr()

    if expectedNumStr != actualNumStr {
      t.Errorf("Error: Expected Value='%v'. Actual Value='%v' k='%v'",
        expectedNumStr, actualNumStr, k)
    }

    actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
        expectedNumSeps.String(), actualNumSeps.String())
    }
  }
}

func TestBigIntMathDivide_INumMgrFracQuotientArray_03(t *testing.T) {

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
    t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
      "divisor='%v' Error='%v' ",
      divisorStr, err.Error())
  }

  dividends := make([]INumMgr, lenDividends)
  expectedResults := make([]BigIntNum, lenDividends)

  for i := 0; i < lenDividends; i++ {

    dec, err := new(Decimal).NewNumStr(dividendArrayStr[i])

    if err != nil {
      t.Errorf("Error returned by new(Decimal).NewNumStr(dividendArrayStr[i]). "+
        "dividendArrayStr[%v]='%v' Error='%v' ",
        i, dividendArrayStr[i], err.Error())
    }

    dividends[i] = &dec

    expectedResults[i], err = new(BigIntNum).NewNumStrWithNumSeps(expectedArrayStr[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps"+
        "(expectedArrayStr[i], expectedNumSeps). "+
        "expectedArrayStr[%v]='%v' Error='%v' ",
        i, expectedArrayStr[i], err.Error())
    }

  }

  err = dividends[0].SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by dividends[0].SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v'", err.Error())
  }

  resultArray, err := new(BigIntMathDivide).INumMgrFracQuotientArray(dividends, &divisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by INumMgrFracQuotientArray{}.INumMgrFracQuotientArray"+
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

    expectedNumStr := expectedResults[k].GetNumStr()
    actualNumStr := resultArray[k].GetNumStr()

    if expectedNumStr != actualNumStr {
      t.Errorf("Error: Expected Value='%v'. Actual Value='%v' k='%v'",
        expectedNumStr, actualNumStr, k)
    }

    actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
        expectedNumSeps.String(), actualNumSeps.String())
    }
  }
}

func TestBigIntMathDivide_INumMgrModulo_01(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12.555					%						 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  decimalDividend, err := new(Decimal).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  nDtoDivisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).INumMgrModulo(&decimalDividend, &nDtoDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).INumMgrModulo(&decimalDividend, "+
      "&nDtoDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

}

func TestBigIntMathDivide_INumMgrModulo_02(t *testing.T) {
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

  bINumDivisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).INumMgrModulo(&iaDividend, &bINumDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).INumMgrModulo(&iaDividend, "+
      "&bINumDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

}

func TestBigIntMathDivide_INumMgrModulo_03(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12.555					% 				 - 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "-2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  bINumDividend, err := new(BigIntNum).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  decimalDivisor, err := new(Decimal).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err :=
    new(BigIntMathDivide).INumMgrModulo(&bINumDividend, &decimalDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).INumMgrModulo(&bINumDividend, "+
      "&decimalDivisor, maxPrecision). Error='%v'", err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

}

func TestBigIntMathDivide_INumMgrModulo_04(t *testing.T) {
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

  nDtoDividend, err := new(NumStrDto).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(dividendStr). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  iaDivisor, err := new(IntAry).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err :=
    new(BigIntMathDivide).INumMgrModulo(&nDtoDividend, &iaDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).INumMgrModulo(nDtoDividend, "+
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

func TestBigIntMathDivide_INumMgrModulo_05(t *testing.T) {
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

  decimalDividend, err :=
    Decimal{}.NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStrWithNumSeps("+
      "dividendStr, expectedNumSeps). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  nDtoDivisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).INumMgrModulo(&decimalDividend, &nDtoDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).INumMgrModulo(&decimalDividend, "+
      "&nDtoDivisor, maxPrecision). Error='%v'", err.Error())

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

func TestBigIntMathDivide_INumMgrModulo_06(t *testing.T) {
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

  decimalDividend, err :=
    Decimal{}.NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStrWithNumSeps("+
      "dividendStr, expectedNumSeps). "+
      "dividendStr='%v' error='%v'", dividendStr, err.Error())
  }

  nDtoDivisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(divisorStr). "+
      "divisorStr='%v' error='%v'", divisorStr, err.Error())
  }

  moduloBINum, err := new(BigIntMathDivide).INumMgrModulo(&decimalDividend, &nDtoDivisor, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).INumMgrModulo(&decimalDividend, "+
      "&nDtoDivisor, maxPrecision). Error='%v'", err.Error())

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

func TestBigIntMathDivide_NumStrQuotientMod_01(t *testing.T) {
  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555					/						 2.5			=			 5							 0.055
  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedQuoStr := "5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedModuloStr). "+
      "expectedModuloStr='%v' Error='%v' ",
      expectedModuloStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrQuotientMod(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrQuotientMod"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' Error='%v' ",
      dividendStr, divisorStr, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  if !expectedModulo.Equal(modulo) {
    t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
      expectedModulo.GetNumStr(), modulo.GetNumStr())
  }

  actualNumSeps := modulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathDivide_NumStrQuotientMod_02(t *testing.T) {
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

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedModuloStr). "+
      "expectedModuloStr='%v' Error='%v' ",
      expectedModuloStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrQuotientMod(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrQuotientMod"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' Error='%v' ",
      dividendStr, divisorStr, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  if !expectedModulo.Equal(modulo) {
    t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
      expectedModulo.GetNumStr(), modulo.GetNumStr())
  }

  actualNumSeps := modulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathDivide_NumStrQuotientMod_03(t *testing.T) {
  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //	-12.555 				/ 				   2.5 			= 		-5							-0.055
  dividendStr := "-12.555"
  divisorStr := "2.5"
  expectedQuoStr := "-5"
  expectedModuloStr := "-0.055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedModuloStr). "+
      "expectedModuloStr='%v' Error='%v' ",
      expectedModuloStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrQuotientMod(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrQuotientMod"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' Error='%v' ",
      dividendStr, divisorStr, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  if !expectedModulo.Equal(modulo) {
    t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
      expectedModulo.GetNumStr(), modulo.GetNumStr())
  }

  actualNumSeps := modulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathDivide_NumStrQuotientMod_04(t *testing.T) {
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

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedModuloStr). "+
      "expectedModuloStr='%v' Error='%v' ",
      expectedModuloStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrQuotientMod(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrQuotientMod"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' Error='%v' ",
      dividendStr, divisorStr, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  if !expectedModulo.Equal(modulo) {
    t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
      expectedModulo.GetNumStr(), modulo.GetNumStr())
  }

  actualNumSeps := modulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrQuotientMod_05(t *testing.T) {
  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  // 	 12.555					/ 				 - 2.5			=			-5							 0.055

  dividendStr := "12.555"
  divisorStr := "-2.5"
  expectedQuoStr := "-5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedModuloStr). "+
      "expectedModuloStr='%v' Error='%v' ",
      expectedModuloStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrQuotientMod(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrQuotientMod"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' Error='%v' ",
      dividendStr, divisorStr, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  if !expectedModulo.Equal(modulo) {
    t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
      expectedModulo.GetNumStr(), modulo.GetNumStr())
  }

  actualNumSeps := modulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrQuotientMod_06(t *testing.T) {
  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555 				/ 				 - 2 				= 		-6							 0.555

  dividendStr := "12.555"
  divisorStr := "-2"
  expectedQuoStr := "-6"
  expectedModuloStr := "0.555"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedModuloStr). "+
      "expectedModuloStr='%v' Error='%v' ",
      expectedModuloStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrQuotientMod(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrQuotientMod"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' Error='%v' ",
      dividendStr, divisorStr, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  if !expectedModulo.Equal(modulo) {
    t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
      expectedModulo.GetNumStr(), modulo.GetNumStr())
  }

  actualNumSeps := modulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrQuotientMod_07(t *testing.T) {
  // Dividend			divided by		Divisor		=		Quotient			Modulo/Remainder
  //   12,555					/					 2,5			=			 5							 0,055
  dividendStr := "12,555"
  divisorStr := "2,5"
  expectedQuoStr := "5"
  expectedModuloStr := "0,055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, expectedNumSeps). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, expectedNumSeps). "+
      "expectedModuloStr='%v' Error='%v' ",
      expectedModuloStr, err.Error())
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrQuotientMod(dividendStr, divisorStr, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrQuotientMod"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' Error='%v' ",
      dividendStr, divisorStr, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  if !expectedModulo.Equal(modulo) {
    t.Errorf("Error: Expected Modulo='%v'. Instead Modulo='%v'",
      expectedModulo.GetNumStr(), modulo.GetNumStr())
  }

  actualNumSeps := modulo.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathDivide_NumStrFracQuotient_01(t *testing.T) {
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

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, err :=
    new(BigIntMathDivide).NumStrFracQuotient(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrFracQuotient"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' maxPrecision='%v' Error='%v' ",
      dividendStr, divisorStr, maxPrecision, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  actualNumSeps := quotient.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrFracQuotient_02(t *testing.T) {
  // Dividend		divided by		Divisor			=		Quotient
  //	-12.555 			/ 					2.5 			= 		-5.022

  dividendStr := "-12.555"
  divisorStr := "2.5"
  expectedQuoStr := "-5.022"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, err :=
    new(BigIntMathDivide).NumStrFracQuotient(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrFracQuotient"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' maxPrecision='%v' Error='%v' ",
      dividendStr, divisorStr, maxPrecision, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  actualNumSeps := quotient.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrFracQuotient_03(t *testing.T) {
  // Dividend		divided by		Divisor			=		Quotient
  //  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

  dividendStr := "-2.5"
  divisorStr := "12.555"
  expectedQuoStr := "-0.199123855037834"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, err :=
    new(BigIntMathDivide).NumStrFracQuotient(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrFracQuotient"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' maxPrecision='%v' Error='%v' ",
      dividendStr, divisorStr, maxPrecision, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  actualNumSeps := quotient.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrFracQuotient_04(t *testing.T) {
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

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  quotient, err :=
    new(BigIntMathDivide).NumStrFracQuotient(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrFracQuotient"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' maxPrecision='%v' Error='%v' ",
      dividendStr, divisorStr, maxPrecision, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  actualNumSeps := quotient.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrFracQuotient_05(t *testing.T) {
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

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, expectedNumSeps). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

  quotient, err :=
    new(BigIntMathDivide).NumStrFracQuotient(dividendStr, divisorStr, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrFracQuotient"+
      "(dividendStr, divisorStr, maxPrecision).  "+
      "dividendStr='%v' divisorStr='%v' maxPrecision='%v' Error='%v' ",
      dividendStr, divisorStr, maxPrecision, err.Error())
  }

  if !expectedQuo.Equal(quotient) {
    t.Errorf("Error: Expected Quotient='%v'. Instead Quotient='%v'",
      expectedQuo.GetNumStr(), quotient.GetNumStr())
  }

  actualNumSeps := quotient.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrFracQuotientArray_01(t *testing.T) {

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

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  resultArray, err :=
    new(BigIntMathDivide).NumStrFracQuotientArray(dividendArrayStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrFracQuotientArray"+
      "(dividendArrayStr, divisorStr, maxPrecision). "+
      "divisor='%v' maxPrecision='%v' Error='%v' ",
      divisorStr, maxPrecision, err.Error())
  }

  lenResultArray := len(resultArray)

  if lenDividends != lenResultArray {
    t.Errorf("Error: Expected Results Array Length='%v'. Actual Array Length='%v'.",
      lenDividends, lenResultArray)
  }

  for k := 0; k < lenDividends; k++ {

    if expectedArrayStr[k] != resultArray[k].GetNumStr() {
      t.Errorf("Expected Value='%v'. Actual Value='%v' k='%v'",
        expectedArrayStr[k], resultArray[k].GetNumStr(), k)
    }

    actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
        expectedNumSeps.String(), actualNumSeps.String())
    }

  }
}

func TestBigIntMathDivide_NumStrFracQuotientArray_02(t *testing.T) {

  divisorStr := "2,5"
  maxPrecision := uint(15)

  dividendArrayStr := []string{
    "10,5",
    "10",
    "11,5",
    "2,5",
    "-12,555",
    "-2,5",
    "12,555",
    "-122,783",
    "-6847,231",
    "-2,5",
    "-10",
    "-10,5",
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

  resultArray, err :=
    new(BigIntMathDivide).NumStrFracQuotientArray(
      dividendArrayStr,
      divisorStr,
      expectedNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrFracQuotientArray"+
      "(dividendArrayStr, divisorStr, maxPrecision). "+
      "divisor='%v' maxPrecision='%v' Error='%v' ",
      divisorStr, maxPrecision, err.Error())
  }

  lenResultArray := len(resultArray)

  if lenDividends != lenResultArray {
    t.Errorf("Error: Expected Results Array Length='%v'. Actual Array Length='%v'.",
      lenDividends, lenResultArray)
  }

  for k := 0; k < lenDividends; k++ {

    if expectedArrayStr[k] != resultArray[k].GetNumStr() {
      t.Errorf("Expected Value='%v'. Actual Value='%v' k='%v'",
        expectedArrayStr[k], resultArray[k].GetNumStr(), k)
    }

    actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
        expectedNumSeps.String(), actualNumSeps.String())
    }

  }
}

func TestBigIntMathDivide_NumStrModulo_01(t *testing.T) {
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

  dto := &NumericSeparatorDto{}
  numSeps := dto.New()

  moduloBINum, err := new(BigIntMathDivide).NumStrModulo(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrDtoModulo(dividendStr, "+
      "divisorStr, maxPrecision). "+
      "dividendStr='%v' divisorStr='%v' Error='%v'",
      dividendStr, divisorStr, err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrModulo_02(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   -12.555 				% 				 - 2.5 			= 		-0.055

  dividendStr := "-12.555"
  divisorStr := "-2.5"
  expectedModuloStr := "-0.055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  numSeps := new(NumericSeparatorDto).New()

  moduloBINum, err := new(BigIntMathDivide).NumStrModulo(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrDtoModulo(dividendStr, "+
      "divisorStr, maxPrecision). "+
      "dividendStr='%v' divisorStr='%v' Error='%v'",
      dividendStr, divisorStr, err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrModulo_03(t *testing.T) {
  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12.555					% 				 - 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "-2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  numSeps := new(NumericSeparatorDto).New()

  moduloBINum, err := new(BigIntMathDivide).NumStrModulo(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrDtoModulo(dividendStr, "+
      "divisorStr, maxPrecision). "+
      "dividendStr='%v' divisorStr='%v' Error='%v'",
      dividendStr, divisorStr, err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrModulo_04(t *testing.T) {
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

  numSeps := new(NumericSeparatorDto).New()

  moduloBINum, err := new(BigIntMathDivide).NumStrModulo(dividendStr, divisorStr, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrDtoModulo(dividendStr, "+
      "divisorStr, maxPrecision). "+
      "dividendStr='%v' divisorStr='%v' Error='%v'",
      dividendStr, divisorStr, err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected moduloBINum='%v'. Instead moduloBINum='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathDivide_NumStrModulo_05(t *testing.T) {
  // Dividend			  mod by			Divisor		=			Modulo/Remainder
  // --------				------			-------					----------------
  //  12,555				  % 				 2				= 		   0,555

  dividendStr := "12,555"
  divisorStr := "2"
  expectedModuloStr := "0,555"
  maxPrecision := uint(15)

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  moduloBINum, err := new(BigIntMathDivide).NumStrModulo(
    dividendStr,
    divisorStr,
    expectedNumSeps,
    maxPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathDivide).NumStrDtoModulo(dividendStr, "+
      "divisorStr, maxPrecision). "+
      "dividendStr='%v' divisorStr='%v' Error='%v'",
      dividendStr, divisorStr, err.Error())

  }

  actualModuloStr := moduloBINum.GetNumStr()

  if expectedModuloStr != actualModuloStr {
    t.Errorf("Error: Expected moduloBINum='%v'. Instead moduloBINum='%v'",
      expectedModuloStr, actualModuloStr)
  }

  actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}
