package mathops

import "testing"

func TestBigIntMathDivide_BigIntNumQuotientMod_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_01"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555					/						 2.5			=			 5							 0.055
  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedQuoStr := "5"
  expectedModuloStr := "0.055"
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

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, divisorStr, err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, divisorStr, err.Error())
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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

  numSeps := new(NumericSeparatorDto).NewUSADefaults()

  actualQuotient, actualModulo, err :=
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, numSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotient, actualModulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, numSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "numSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendNumStr,
      divisorNumStr,
      numSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  actualQuotientNumStr, err := actualQuotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuotientNumStr, err := actualQuotient.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := actualModulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := actualModulo.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuotient)\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Quotient is NOT Equal to Actual Quotient.\n"+
      "Expected actualQuotient = '%v'\n"+
      "Instead, actualQuotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(actualModulo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err := expectedModulo.Equal(actualModulo)\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Quotient is NOT Equal to Actual Quotient.\n"+
      "Expected actualModulo = '%v'\n"+
      "Instead, actualModulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected and Actual Number Strings don't match.\n"+
      "Expected actualQuotientNumStr = '%v'\n"+
      "Instead, actualQuotientNumStr = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  if expectedModuloNumStr != actualModuloNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected and Actual Number Strings don't match.\n"+
      "Expected actualModuloNumStr = '%v'\n"+
      "Instead, actualModuloNumStr = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)
  }

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_02"

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
      "dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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
      "dividend= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "dividend= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedModuloStr, err.Error())
    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr,
      dividendNumSeps.String(), maxPrecision, err.Error())
    return
  }

  err = quotient.IsValid(ePrefix + "\nValidating quotient")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = quotient.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = modulo.IsValid(ePrefix + "\nValidating modulo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = modulo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsQuotient, err := expectedQuo.Equal(quotient)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsQuotient , err = expectedQuo.Equal(quotient)\n"+
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

  quotientNumStr, err := quotient.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotientNumStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsQuotient {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuo.Equal(quotient) == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, quotientNumStr)

    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)\n"+
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

  moduloNumStr, err := modulo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloNumStr, err := modulo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuo.Equal(quotient) == 'false'\n"+
      "Expected 'modulo' = '%v'\n"+
      "  Actual 'modulo' = '%v'\n\n",
      ePrefix, expectedModuloNumStr, moduloNumStr)

    return
  }
}

func TestBigIntMathDivide_BigIntNumQuotientMod_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_03"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //    2.5 					/ 				 	12.555		= 	   0							 2.5
  dividendStr := "2.5"
  divisorStr := "12.555"
  expectedQuoStr := "0"
  expectedModuloStr := "2.5"
  maxPrecision := uint(15)

  dividend, err := new(BigIntNum).NewNumStr(dividendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(BigIntNum).NewNumStr(dividendStr)\n"+
      "dividendStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, divisorStr, err.Error())
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "dividend, divisor, dividendNumSeps, maxPrecision)\n"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_04"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //	-12.555 				/ 				   2.5 			= 		-5							-0.055
  dividendStr := "-12.555"
  divisorStr := "2.5"
  expectedQuoStr := "-5"
  expectedModuloStr := "-0.055"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_05"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //  -12.555     		/    			 	 2  			= 		-6							-0.555
  dividendStr := "-12.555"
  divisorStr := "2"
  expectedQuoStr := "-6"
  expectedModuloStr := "-0.555"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_06"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //  - 2.5 					/ 				 	12.555		= 		 0							-2.5
  dividendStr := "-2.5"
  divisorStr := "12.555"
  expectedQuoStr := "0"
  expectedModuloStr := "-2.5"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_07(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_07"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  // 	 12.555					/ 				 - 2.5			=			-5							 0.055

  dividendStr := "12.555"
  divisorStr := "-2.5"
  expectedQuoStr := "-5"
  expectedModuloStr := "0.055"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_08(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_08"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555 				/ 				 - 2 				= 		-6							 0.555

  dividendStr := "12.555"
  divisorStr := "-2"
  expectedQuoStr := "-6"
  expectedModuloStr := "0.555"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_09(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_09"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //    2.5 				  / 				 -12.555		= 		 0							 2.5

  dividendStr := "2.5"
  divisorStr := "-12.555"
  expectedQuoStr := "0"
  expectedModuloStr := "2.5"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_10(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_10"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  // 	-12.555 				/ 				 - 2.5 			= 		 5							-0.055

  dividendStr := "-12.555"
  divisorStr := "-2.5"
  expectedQuoStr := "5"
  expectedModuloStr := "-0.055"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by\n"+
      " new(BigIntNum).NewNumStr(divisorStr).\n"+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_11(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_11"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //  -12.555     		/    			 - 2 				= 		 6							-0.555

  dividendStr := "-12.555"
  divisorStr := "-2"
  expectedQuoStr := "6"
  expectedModuloStr := "-0.555"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_12(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_12"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //  - 2.5	 					/ 				 -12.555		= 		 0							-2.5

  dividendStr := "-2.5"
  divisorStr := "-12.555"
  expectedQuoStr := "0"
  expectedModuloStr := "-2.5"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_13(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_13"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //  0	 					   / 				  -12.555			= 		 0							0

  dividendStr := "0"
  divisorStr := "-12.555"
  expectedQuoStr := "0"
  expectedModuloStr := "0"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_14(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_14"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //  0	 					   / 				   12.555			= 		 0							0

  dividendStr := "0"
  divisorStr := "12.555"
  expectedQuoStr := "0"
  expectedModuloStr := "0"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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
      "Because expectedNumSeps.Equal(actualNumSeps) == 'false'\n"+
      "Expected Numeric Separators = '%v'\n"+
      "  Actual Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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
      "Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
      "Expected Numeric Separators = '%v'\n"+
      "  Actual Numeric Separators = '%v'\n\n",
      ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_15(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_15"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //  4	 					   / 				    2					= 		 2							0

  dividendStr := "4"
  divisorStr := "2"
  expectedQuoStr := "2"
  expectedModuloStr := "0"
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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
      "expectedModuloStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedModuloStr, err.Error())
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
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, dividendNumSeps, maxPrecision)"+
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

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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
      "Because expectedNumSeps.Equal(actualNumSeps) == 'false'\n"+
      "Expected Numeric Separators = '%v'\n"+
      "  Actual Numeric Separators = '%v'\n\n",
      ePrefix, expectedQuoNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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
      "Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
      "Expected Numeric Separators = '%v'\n"+
      "  Actual Numeric Separators = '%v'\n\n",
      ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_16(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_16"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555					/						 2.5			=			 5							 0.055
  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedQuoNumStr := "5"
  expectedModuloNumStr := "0,055"
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

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, divisorStr, err.Error())
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

  quotient, modulo, err :=
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)"+
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

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

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

  if expectedModuloNumStr != actualModuloNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because  expectedModuloNumStr != actualModuloNumStr\n"+
      "Expected modulo = '%v'\n"+
      "  Actual modulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient NumSeps = '%v'\n"+
      "  Actual Quotient NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err = modulo.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
      "Expected Modulo NumSeps = '%v'\n"+
      "  Actual Modulo NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_BigIntNumQuotientMod_17(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_17"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //   12.555					/						 2.5			=			 5							 0.055

  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedQuoNumStr := "5"
  expectedModuloNumStr := "0.055"
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

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err = expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
      "divisorStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, divisorStr, err.Error())
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

  quotient, modulo, err :=
    new(BigIntMathDivide).BigIntNumQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err := new(BigIntMathDivide).BigIntNumQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)"+
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

  if expectedQuoNumStr != actualQuotientNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumStr != actualQuotientNumStr\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

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

  if expectedModuloNumStr != actualModuloNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because  expectedModuloNumStr != actualModuloNumStr\n"+
      "Expected modulo = '%v'\n"+
      "  Actual modulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected Quotient NumSeps = '%v'\n"+
      "  Actual Quotient NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), actualQuoNumSeps.String())

    return
  }

  actualModuloNumSeps, err := modulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err = modulo.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
      "Expected Modulo NumSeps = '%v'\n"+
      "  Actual Modulo NumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), actualModuloNumSeps.String())

    return
  }

  return
}

func TestBigIntMathDivide_BigIntNumIntQuotient_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_01"

  // Dividend	 divided by		Divisor		=		   Quotient
  // 		 5 				/ 				 2 				= 				 2

  dividendStr := "5"
  divisorStr := "2"
  expectedQuoStr := "2"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_02"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //     5.25 		/ 				 2  			= 				 2

  dividendStr := "5.25"
  divisorStr := "2"
  expectedQuoStr := "2"

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

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendNumSeps.SetUSADefaults()

  err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividendNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_03"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //     2 				/ 				 4				= 				 0

  dividendStr := "2"
  divisorStr := "4"
  expectedQuoStr := "0"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_04"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  // 		-5 				/ 				 2 				= 				-2

  dividendStr := "-5"
  divisorStr := "2"
  expectedQuoStr := "-2"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_05"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //    -5.25     /    			 2  			= 				-2

  dividendStr := "-5.25"
  divisorStr := "2"
  expectedQuoStr := "-2"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_06"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //    -2 				/ 				 4				= 				 0

  dividendStr := "-2"
  divisorStr := "4"
  expectedQuoStr := "0"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_07(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_07"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  // 		 5 				/ 				-2 				=					-2

  dividendStr := "5"
  divisorStr := "-2"
  expectedQuoStr := "-2"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_08(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_08"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //     5.25 		/ 				-2 				= 				-2

  dividendStr := "5.25"
  divisorStr := "-2"
  expectedQuoStr := "-2"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_09(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_09"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //     2 				/ 				-4				= 				 0

  dividendStr := "2"
  divisorStr := "-4"
  expectedQuoStr := "0"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_10(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_10"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  // 		-5 				/ 				-2 				= 				 2

  dividendStr := "-5"
  divisorStr := "-2"
  expectedQuoStr := "2"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_11(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_11"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //    -5.25     /    			-2 				= 				 2

  dividendStr := "-5.25"
  divisorStr := "-2"
  expectedQuoStr := "2"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_12(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_12"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //    -2 				/ 				-4				= 				 0

  dividendStr := "-2"
  divisorStr := "-4"
  expectedQuoStr := "0"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_13(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_13"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //  12.555			/ 			  -2.5			=			    -5

  dividendStr := "12.555"
  divisorStr := "-2.5"
  expectedQuoStr := "-5"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_14(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_14"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  // -12.555			/ 			  -2.5			=			     5
  dividendStr := "-12.555"
  divisorStr := "-2.5"
  expectedQuoStr := "5"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_15(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_15"

  //																					 Integer
  // Dividend	 divided by		Divisor		=		     Quotient
  //  12.555			/ 			  -2				=			    -6
  dividendStr := "12.555"
  divisorStr := "-2"
  expectedQuoStr := "-6"

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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_16(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_16"

  // Dividend	 divided by		Divisor		=		   Quotient
  // 		 5 				/ 				 2 				= 				 2

  dividendStr := "5"
  divisorStr := "2"
  expectedQuoStr := "2"

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

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_BigIntNumIntQuotient_17(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_BigIntNumIntQuotient_17"

  // Dividend	 divided by		Divisor		=		   Quotient
  // 		 5 				/ 				 2 				= 				 2

  dividendStr := "5"
  divisorStr := "2"
  expectedQuoStr := "2"

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

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err = expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = dividend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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

  divisor, err := new(BigIntNum).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(divisorStr). "+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  quotient, err :=
    new(BigIntMathDivide).BigIntNumIntQuotient(dividend, divisor, dividendNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).BigIntNumIntQuotient(\n"+
      "  dividend, divisor, dividendNumSeps)"+
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
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuotientNumStr)

    return
  }

  actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := quotient.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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
