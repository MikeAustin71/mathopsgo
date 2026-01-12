package mathops

import (
  "fmt"
  "testing"
)

func TestBigIntMathDivide_NumStrDtoFracQuotient_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotient_01"

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

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

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
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid(ePrefix + "\nValidating expectedQuo")

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

  quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(\n"+
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

func TestBigIntMathDivide_NumStrDtoFracQuotient_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotient_02"

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

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

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

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

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
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid(ePrefix + "\nValidating expectedQuo")

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

  quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(\n"+
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

func TestBigIntMathDivide_NumStrDtoFracQuotient_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotient_03"

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

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

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

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

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
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid(ePrefix + "\nValidating expectedQuo")

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

  quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(\n"+
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

func TestBigIntMathDivide_NumStrDtoFracQuotient_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotient_04"

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

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

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
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid(ePrefix + "\nValidating expectedQuo")

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

  quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(\n"+
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

func TestBigIntMathDivide_NumStrDtoFracQuotient_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotient_05"

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

  if !expectedNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

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
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

  err = expectedQuo.IsValid(ePrefix + "\nValidating expectedQuo")

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

  quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(\n"+
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

func TestBigIntMathDivide_NumStrDtoFracQuotient_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotient_06"

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

  dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  dividendStr, &expectedNumSeps)\n"+
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

  divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  divisorStr, &expectedNumSeps)\n"+
      "divisorStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, divisorStr, expectedNumSeps.String(), err.Error())
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

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
      "expectedQuoStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedQuo.IsValid(ePrefix + "\nValidating expectedQuo")

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

  quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(\n"+
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

func TestBigIntMathDivide_NumStrDtoFracQuotient_07(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotient_07"

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

  dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  dividendStr, &expectedNumSeps)\n"+
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

  divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  divisorStr, &expectedNumSeps)\n"+
      "divisorStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, divisorStr, expectedNumSeps.String(), err.Error())
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

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
      "expectedQuoStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedQuo.IsValid(ePrefix + "\nValidating expectedQuo")

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

  quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(BigIntMathDivide).NumStrDtoFracQuotient(\n"+
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

func TestBigIntMathDivide_NumStrDtoFracQuotientArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotientArray_01"

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

  lenDividends := len(dividendArrayStr)

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividends := make([]NumStrDto, lenDividends)

  expectedResults := make([]NumStrDto, lenDividends)

  for i := 0; i < lenDividends; i++ {

    dividends[i], err = new(NumStrDto).NewNumStrWithNumSeps(dividendArrayStr[i], &expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "  new(NumStrDto).NewNumStrWithNumSeps(\n"+
        "  dividendArrayStr[%d], &expectedNumSeps)\n"+
        "dividendArrayStr[%v]= '%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, dividendArrayStr[i], expectedNumSeps.String(), err.Error())
      return
    }

    expectedResults[i], err = new(NumStrDto).NewNumStrWithNumSeps(expectedArrayStr[i], &expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        " new(NumStrDto).NewNumStrWithNumSeps(\n"+
        "  expectedArrayStr[%d], &expectedNumSeps)\n"+
        "expectedArrayStr[%v]='%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedArrayStr[i],
        expectedNumSeps.String(), err.Error())
      return
    }

  }

  resultArray, err := new(BigIntMathDivide).NumStrDtoFracQuotientArray(dividends, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultArray, err := new(BigIntMathDivide).\n"+
      "  NumStrDtoFracQuotientArray(dividends, divisor,\n"+
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

  var actualNumSeps NumericSeparatorDto

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

    expectedResultsNumStr, err = expectedResults[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResults[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, k, err.Error())
      return
    }

    actualEqualsExpectedResults = resultArray[k].Equal(expectedResults[k])

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

    actualNumSeps, err = resultArray[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualNumSeps, err = resultArray[%d].GetNumericSeparatorsDto()\n"+
        "Error='%v'\n\n",
        ePrefix, k, err.Error())
      return
    }

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumSeps.Equal(actualNumSeps) == 'false'\n"+
        "Expected Numeric Separators = '%v'\n"+
        "  Actual Numeric Separators = '%v'\n\n",
        ePrefix, expectedNumSeps.String(), actualNumSeps.String())

      return
    }

  }

  return
}

func TestBigIntMathDivide_NumStrDtoFracQuotientArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotientArray_02"

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

  divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("Error returned by\n"+
      "divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)\n"+
      "divisorStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, usaNumSeps.String(), err.Error())
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

  dividends := make([]NumStrDto, lenDividends)

  expectedResults := make([]NumStrDto, lenDividends)

  for i := 0; i < lenDividends; i++ {

    dividends[i], err = new(NumStrDto).NewNumStrWithNumSeps(dividendArrayStr[i], &usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "  new(NumStrDto).NewNumStrWithNumSeps(\n"+
        "  dividendArrayStr[%d], &expectedNumSeps)\n"+
        "dividendArrayStr[%v]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, dividendArrayStr[i], usaNumSeps.String(), err.Error())
      return
    }

    err = dividends[i].IsValid(ePrefix + fmt.Sprintf(" Validating dividends[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err =dividends[%d].IsValid(ePrefix)\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResults[i], err = new(NumStrDto).NewNumStrWithNumSeps(expectedArrayStr[i], &expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        " new(NumStrDto).NewNumStrWithNumSeps(\n"+
        "  expectedArrayStr[%d], &expectedNumSeps)\n"+
        "expectedArrayStr[%v]='%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedArrayStr[i],
        expectedNumSeps.String(), err.Error())
      return
    }

    err = expectedResults[i].IsValid(ePrefix + fmt.Sprintf(" Validating expectedResults[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = expectedResults[%d].IsValid(ePrefix)\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

  }

  resultArray, err := new(BigIntMathDivide).NumStrDtoFracQuotientArray(dividends, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultArray, err := new(BigIntMathDivide).\n"+
      "  NumStrDtoFracQuotientArray(dividends, divisor,\n"+
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

  var actualNumSeps NumericSeparatorDto

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

    expectedResultsNumStr, err = expectedResults[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResults[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, k, err.Error())
      return
    }

    actualEqualsExpectedResults = resultArray[k].Equal(expectedResults[k])

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

    actualNumSeps, err = resultArray[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualNumSeps, err = resultArray[%d].GetNumericSeparatorsDto()\n"+
        "Error='%v'\n\n",
        ePrefix, k, err.Error())
      return
    }

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumSeps.Equal(actualNumSeps) == 'false'\n"+
        "Expected Numeric Separators = '%v'\n"+
        "  Actual Numeric Separators = '%v'\n\n",
        ePrefix, expectedNumSeps.String(), actualNumSeps.String())

      return
    }

  }

  return
}

func TestBigIntMathDivide_NumStrDtoFracQuotientArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoFracQuotientArray_03"

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

  divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("Error returned by\n"+
      "divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)\n"+
      "divisorStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, usaNumSeps.String(), err.Error())
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

  dividends := make([]NumStrDto, lenDividends)

  expectedResults := make([]NumStrDto, lenDividends)

  for i := 0; i < lenDividends; i++ {

    dividends[i], err = new(NumStrDto).NewNumStrWithNumSeps(dividendArrayStr[i], &usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "  new(NumStrDto).NewNumStrWithNumSeps(\n"+
        "  dividendArrayStr[%d], &expectedNumSeps)\n"+
        "dividendArrayStr[%v]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, dividendArrayStr[i], usaNumSeps.String(), err.Error())
      return
    }

    err = dividends[i].IsValid(ePrefix + fmt.Sprintf(" Validating dividends[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err =dividends[%d].IsValid(ePrefix)\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResults[i], err = new(NumStrDto).NewNumStrWithNumSeps(expectedArrayStr[i], &expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        " new(NumStrDto).NewNumStrWithNumSeps(\n"+
        "  expectedArrayStr[%d], &expectedNumSeps)\n"+
        "expectedArrayStr[%v]='%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedArrayStr[i],
        expectedNumSeps.String(), err.Error())
      return
    }

    err = expectedResults[i].IsValid(ePrefix + fmt.Sprintf(" Validating expectedResults[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = expectedResults[%d].IsValid(ePrefix)\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

  }

  resultArray, err := new(BigIntMathDivide).NumStrDtoFracQuotientArray(dividends, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultArray, err := new(BigIntMathDivide).\n"+
      "  NumStrDtoFracQuotientArray(dividends, divisor,\n"+
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

  var actualNumSeps NumericSeparatorDto

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

    expectedResultsNumStr, err = expectedResults[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResults[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, k, err.Error())
      return
    }

    actualEqualsExpectedResults = resultArray[k].Equal(expectedResults[k])

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

    actualNumSeps, err = resultArray[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualNumSeps, err = resultArray[%d].GetNumericSeparatorsDto()\n"+
        "Error='%v'\n\n",
        ePrefix, k, err.Error())
      return
    }

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumSeps.Equal(actualNumSeps) == 'false'\n"+
        "Expected Numeric Separators = '%v'\n"+
        "  Actual Numeric Separators = '%v'\n\n",
        ePrefix, expectedNumSeps.String(), actualNumSeps.String())

      return
    }

  }

  return
}

func TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_01"

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

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
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

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(\n"+
      "  dividendStr, divisorStr, expectedNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloNDto.IsValid("Validating moduloNDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloNDto.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloNDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloNDto.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo := expectedModulo.Equal(moduloNDto)

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_02"

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

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
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

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(\n"+
      "  dividendStr, divisorStr, expectedNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloNDto.IsValid("Validating moduloNDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloNDto.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloNDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloNDto.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo := expectedModulo.Equal(moduloNDto)

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_03"

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

  dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &usaNumSeps)\n"+
      "dividendStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, usaNumSeps.String(), err.Error())
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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
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
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)\n"+
      "divisorStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, usaNumSeps.String(), err.Error())
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

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(\n"+
      "  dividendStr, divisorStr, expectedNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloNDto.IsValid("Validating moduloNDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloNDto.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloNDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloNDto.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo := expectedModulo.Equal(moduloNDto)

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_04"

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

  dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &usaNumSeps)\n"+
      "dividendStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, usaNumSeps.String(), err.Error())
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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
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
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)\n"+
      "divisorStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, usaNumSeps.String(), err.Error())
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

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(
    dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(\n"+
      "  dividendStr, divisorStr, expectedNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloNDto.IsValid("Validating moduloNDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloNDto.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloNDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloNDto.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo := expectedModulo.Equal(moduloNDto)

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_05"

  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //   12,555					%					 2.5				=			 0,055

  dividendStr := "12,555"
  divisorStr := "2.5"
  expectedModuloStr := "0,055"
  maxPrecision := uint(15)

  frenchNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  frenchNumSeps.DecimalSeparator = frenchDecSeparator
  frenchNumSeps.ThousandsSeparator = frenchThousandsSeparator
  frenchNumSeps.CurrencySymbol = frenchCurrencySymbol

  err := frenchNumSeps.IsValid("Validating frenchNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = frenchNumSeps.IsValid('Validating frenchNumSeps')\n"+
      "frenchNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, frenchNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &frenchNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &frenchNumSeps)\n"+
      "dividendStr='%v'\n"+
      "frenchNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, frenchNumSeps.String(), err.Error())
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

  if dividendStr != dividendNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because dividendStr != dividendNumStr\n"+
      "Expected dividendNumStr = '%v'\n"+
      "  Actual dividendNumStr = '%v'\n\n",
      ePrefix, dividendStr, dividendNumStr)

    return
  }

  dividendNumSeps, err := dividend.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendNumSeps, err := dividend.GetNumericSeparatorsDto()\n"+
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !frenchNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because frenchNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, frenchNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &usaNumSeps)\n"+
      "divisorStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, usaNumSeps.String(), err.Error())
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

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(expectedModuloStr, &frenchNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &frenchNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "frenchNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, frenchNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !frenchNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because frenchNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, frenchNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(
    dividend, divisor, frenchNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(\n"+
      "  dividendStr, divisorStr, frenchNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "frenchNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      frenchNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloNDto.IsValid("Validating moduloNDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloNDto.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloNDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloNDto.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo := expectedModulo.Equal(moduloNDto)

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = frenchNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := frenchNumSeps.CopyOut(false)\n"+
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

func TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModuloToNumStrDto_06"

  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //  12.555					%					 2.5				=			 0.055

  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedModuloStr := "0.055"
  maxPrecision := uint(15)

  frenchNumSeps := NumericSeparatorDto{}

  frenchNumSeps.DecimalSeparator = '.'
  frenchNumSeps.ThousandsSeparator = ','
  frenchNumSeps.CurrencySymbol = '$'

  err := frenchNumSeps.IsValid("Validating frenchNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = frenchNumSeps.IsValid('Validating frenchNumSeps')\n"+
      "frenchNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, frenchNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &usaNumSeps)\n"+
      "dividendStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, dividendStr, usaNumSeps.String(), err.Error())
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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !frenchNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because frenchNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, frenchNumSeps.String(), dividendNumSeps.String())

    return
  }

  if dividendStr != dividendNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because dividendStr != dividendNumStr\n"+
      "Expected dividendNumStr = '%v'\n"+
      "  Actual dividendNumStr = '%v'\n\n",
      ePrefix, dividendStr, dividendNumStr)

    return
  }

  divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &usaNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &usaNumSeps)\n"+
      "divisorStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      divisorStr, usaNumSeps.String(), err.Error())
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

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if divisorStr != divisorNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because divisorStr != divisorNumStr\n"+
      "Expected divisorNumStr = '%v'\n"+
      "  Actual divisorNumStr = '%v'\n\n",
      ePrefix, divisorStr, divisorNumStr)

    return
  }

  if !frenchNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because frenchNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, frenchNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(expectedModuloStr, &frenchNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &frenchNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "frenchNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, frenchNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !frenchNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because frenchNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, frenchNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(
    dividend, divisor, frenchNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloNDto, err := new(BigIntMathDivide).NumStrDtoModuloToNumStrDto(\n"+
      "  dividendStr, divisorStr, frenchNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "frenchNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      frenchNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloNDto.IsValid("Validating moduloNDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloNDto.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNDtoNumStr, err := moduloNDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNDtoNumStr, err := moduloNDto.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo := expectedModulo.Equal(moduloNDto)

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNDtoNumStr)

    return
  }

  if expectedModuloNumStr != actualModuloNDtoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloNumStr != actualModuloNDtoNumStr\n"+
      "Expected actualModuloNDtoNumStr = '%v'\n"+
      "  Actual actualModuloNDtoNumStr = '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNDtoNumStr)

    return
  }

  actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloNDto.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNDtoNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = frenchNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := frenchNumSeps.CopyOut(false)\n"+
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

func TestBigIntMathDivide_NumStrDtoModulo_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModulo_01"

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloBINum.IsValid("Validating moduloBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloBINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloBINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(moduloBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err :=\n"+
      "  expectedModulo.Equal(moduloBINum)\n"+
      "expectedModulo= '%v'\n"+
      "moduloBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoModulo_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModulo_02"

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloBINum.IsValid("Validating moduloBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloBINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloBINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(moduloBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err :=\n"+
      "  expectedModulo.Equal(moduloBINum)\n"+
      "expectedModulo= '%v'\n"+
      "moduloBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoModulo_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModulo_03"

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloBINum.IsValid("Validating moduloBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloBINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloBINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(moduloBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err :=\n"+
      "  expectedModulo.Equal(moduloBINum)\n"+
      "expectedModulo= '%v'\n"+
      "moduloBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoModulo_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModulo_04"

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloBINum.IsValid("Validating moduloBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloBINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloBINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(moduloBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err :=\n"+
      "  expectedModulo.Equal(moduloBINum)\n"+
      "expectedModulo= '%v'\n"+
      "moduloBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoModulo_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModulo_05"

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
      " dividendStr, &expectedNumSeps)\n"+
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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
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

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloBINum.IsValid("Validating moduloBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloBINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloBINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(moduloBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err :=\n"+
      "  expectedModulo.Equal(moduloBINum)\n"+
      "expectedModulo= '%v'\n"+
      "moduloBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoModulo_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoModulo_06"

  // Dividend			  mod by			Divisor			=			Modulo/Remainder
  // --------				------			-------						----------------
  //  12.555					%						 2.5			=			 0.055

  dividendStr := "12.555"
  divisorStr := "2.5"
  expectedModuloStr := "0.055"
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
      " dividendStr, &expectedNumSeps)\n"+
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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStrWithNumSeps(divisorStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(NumStrDto).NewNumStrWithNumSeps(\n"+
      " divisorStr, &expectedNumSeps)\n"+
      "divisorStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, divisorStr, expectedNumSeps.String(), err.Error())
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

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "moduloBINum, err := new(BigIntMathDivide).NumStrDtoModulo(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividendStr= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  err = moduloBINum.IsValid("Validating moduloBINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = moduloBINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualModuloNumStr, err := moduloBINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumStr, err := moduloBINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedModuloEqualsActualModulo, err := expectedModulo.Equal(moduloBINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloEqualsActualModulo, err :=\n"+
      "  expectedModulo.Equal(moduloBINum)\n"+
      "expectedModulo= '%v'\n"+
      "moduloBINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedModuloNumStr, actualModuloNumStr, err.Error())
    return
  }

  if !expectedModuloEqualsActualModulo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedModuloEqualsActualModulo == 'false'\n"+
      "Expected Modulo = '%v'\n"+
      "  Actual Modulo = '%v'\n\n",
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

  actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualModuloNumSeps, err := moduloBINum.GetNumericSeparatorsDto()\n"+
      "modulo= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, actualModuloNumStr, err.Error())

    return
  }

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

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

func TestBigIntMathDivide_NumStrDtoQuotientMod_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoQuotientMod_01"

  // Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
  //  - 2.5 					/ 				 	12.555		= 		 0							-2.5
  dividendStr := "-2.5"
  divisorStr := "12.555"
  expectedQuoStr := "0"
  expectedModuloStr := "-2.5"
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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()\n"+
      "expectedQuo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedQuoNumSeps\n"+
      "Expected expectedQuoNumSeps = '%v'\n"+
      "  Actual expectedQuoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedQuoNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrDtoQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err :=\n"+
      "  new(BigIntMathDivide).NumStrDtoQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

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

  expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoQuotientMod_02"

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()\n"+
      "expectedQuo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedQuoNumSeps\n"+
      "Expected expectedQuoNumSeps = '%v'\n"+
      "  Actual expectedQuoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedQuoNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrDtoQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err :=\n"+
      "  new(BigIntMathDivide).NumStrDtoQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

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

  expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoQuotientMod_03"

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()\n"+
      "expectedQuo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedQuoNumSeps\n"+
      "Expected expectedQuoNumSeps = '%v'\n"+
      "  Actual expectedQuoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedQuoNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrDtoQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err :=\n"+
      "  new(BigIntMathDivide).NumStrDtoQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

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

  expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoQuotientMod_04"

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()\n"+
      "expectedQuo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedQuoNumSeps\n"+
      "Expected expectedQuoNumSeps = '%v'\n"+
      "  Actual expectedQuoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedQuoNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrDtoQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err :=\n"+
      "  new(BigIntMathDivide).NumStrDtoQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

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

  expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoQuotientMod_05"

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()\n"+
      "expectedQuo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedQuoNumSeps\n"+
      "Expected expectedQuoNumSeps = '%v'\n"+
      "  Actual expectedQuoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedQuoNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrDtoQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err :=\n"+
      "  new(BigIntMathDivide).NumStrDtoQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

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

  expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoQuotientMod_06"

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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(dividendNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != dividendNumSeps\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual dividendNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), dividendNumSeps.String())

    return
  }

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()\n"+
      "expectedQuo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedQuoNumSeps\n"+
      "Expected expectedQuoNumSeps = '%v'\n"+
      "  Actual expectedQuoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedQuoNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrDtoQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err :=\n"+
      "  new(BigIntMathDivide).NumStrDtoQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

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

  expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_07(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoQuotientMod_07"

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

  dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).\n"+
      "  NewNumStrWithNumSeps(dividendStr, &expectedNumSeps)\n"+
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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
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

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()\n"+
      "expectedQuo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedQuoNumSeps\n"+
      "Expected expectedQuoNumSeps = '%v'\n"+
      "  Actual expectedQuoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedQuoNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrDtoQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err :=\n"+
      "  new(BigIntMathDivide).NumStrDtoQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

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

  expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

func TestBigIntMathDivide_NumStrDtoQuotientMod_08(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_NumStrDtoQuotientMod_08"

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

  dividend, err := new(NumStrDto).NewNumStrWithNumSeps(dividendStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(NumStrDto).\n"+
      "  NewNumStrWithNumSeps(dividendStr, &expectedNumSeps)\n"+
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
      "dividend= '%v'"+
      "Error='%v'\n\n", ePrefix, dividendNumStr, err.Error())
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

  divisor, err := new(NumStrDto).NewNumStr(divisorStr)

  if err != nil {
    t.Errorf("Error returned bynew(NumStrDto).NewNumStr(divisorStr). "+
      "divisorStr='%v' Error='%v' ",
      divisorStr, err.Error())
  }

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "new(NumStrDto).NewNumStr(divisorStr)\n"+
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
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorNumSeps, err := divisor.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorNumSeps, err := divisor.GetNumericSeparatorsDto()\n"+
      "divisor= '%v'"+
      "Error='%v'\n\n", ePrefix, divisorNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(divisorNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because usaNumSeps != divisorNumSeps\n"+
      "Expected divisorNumSeps = '%v'\n"+
      "  Actual divisorNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), divisorNumSeps.String())

    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedQuoStr, &expectedNumSeps)\n"+
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

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err := expectedQuo.GetNumericSeparatorsDto()\n"+
      "expectedQuo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedQuoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedQuoNumSeps\n"+
      "Expected expectedQuoNumSeps = '%v'\n"+
      "  Actual expectedQuoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedQuoNumSeps.String())

    return
  }

  expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedModuloStr, &expectedNumSeps)\n"+
      "expectedModuloStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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

  expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err := expectedModulo.GetNumericSeparatorsDto()\n"+
      "expectedModulo= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedModuloNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(expectedModuloNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumSeps != expectedModuloNumSeps\n"+
      "Expected expectedModuloNumSeps = '%v'\n"+
      "  Actual expectedModuloNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedModuloNumSeps.String())

    return
  }

  quotient, modulo, err :=
    new(BigIntMathDivide).NumStrDtoQuotientMod(dividend, divisor, expectedNumSeps, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "quotient, modulo, err :=\n"+
      "  new(BigIntMathDivide).NumStrDtoQuotientMod(\n"+
      "  dividend, divisor, expectedNumSeps, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisorStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      divisorStr,
      expectedNumSeps.String(),
      maxPrecision,
      err.Error())

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

  expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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

  expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedModuloNumSeps, err = expectedNumSeps.CopyOut(false)\n"+
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
