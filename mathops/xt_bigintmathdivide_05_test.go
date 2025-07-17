package mathops

import "testing"

func TestBigIntMathDivide_FixedDecimalFracQuotient_01(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_01"

  // Dividend		divided by		Divisor			=		Quotient
  // 	 10.5  				/ 					2 				= 	 5.25

  dividendStr := "10.5"
  divisorStr := "2"
  expectedQuoStr := "5.25"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_02(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_02"

  // Dividend		divided by		Divisor			=		Quotient
  // 	 10    				/ 					2 				= 			5

  dividendStr := "10"
  divisorStr := "2"
  expectedQuoStr := "5"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_03(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_03"

  // Dividend		divided by		Divisor			=		Quotient
  //   11.5  				/         	2.5				=  		4.6

  dividendStr := "11.5"
  divisorStr := "2.5"
  expectedQuoStr := "4.6"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_04(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_04"

  // Dividend		divided by		Divisor			=		Quotient
  //    2.5					/				 	12.555			=		0.199123855037834

  dividendStr := "2.5"
  divisorStr := "12.555"
  expectedQuoStr := "0.199123855037834"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_05(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_05"

  // Dividend		divided by		Divisor			=		Quotient
  //	-12.555 			/ 					2.5 			= 		-5.022

  dividendStr := "-12.555"
  divisorStr := "2.5"
  expectedQuoStr := "-5.022"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_06(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_06"

  // Dividend		divided by		Divisor			=		Quotient
  //  -12.555     	/    				2  			  = 	 -6.2775

  dividendStr := "-12.555"
  divisorStr := "2"
  expectedQuoStr := "-6.2775"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_07(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_07"

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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_08(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_08"

  // Dividend		divided by		Divisor			=		Quotient
  // 	 12.555				/ 				- 2.5			  =		 -5.022

  dividendStr := "12.555"
  divisorStr := "-2.5"
  expectedQuoStr := "-5.022"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_09(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_09"

  // Dividend		divided by		Divisor			=		Quotient
  //   12.555 			/ 				 -2 				=    -6.2775

  dividendStr := "12.555"
  divisorStr := "-2"
  expectedQuoStr := "-6.2775"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_10(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_10"

  // Dividend		divided by		Divisor			=		Quotient
  //    2.5 				/ 				-12.555		  = 	-0.199123855037834

  dividendStr := "2.5"
  divisorStr := "-12.555"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_11(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_11"

  // Dividend		divided by		Divisor			=		Quotient
  // 	-12.555 			/ 				 -2.5 			= 	 5.022

  dividendStr := "-12.555"
  divisorStr := "-2.5"
  expectedQuoStr := "5.022"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_12(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_12"

  // Dividend		divided by		Divisor			=		Quotient
  //  -12.555     	/    			 -2 				= 		6.2775

  dividendStr := "-12.555"
  divisorStr := "-2"
  expectedQuoStr := "6.2775"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_13(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_13"

  // Dividend		divided by		Divisor			=					Quotient
  //  - 2.5	 				/ 				-12.555		  = 		0.199123855037834

  dividendStr := "-2.5"
  divisorStr := "-12.555"
  expectedQuoStr := "0.199123855037834"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_14(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_14"

  // Dividend		divided by		Divisor			=					Quotient
  //  -10						/					- 2					=						5

  dividendStr := "-10"
  divisorStr := "-2"
  expectedQuoStr := "5"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_15(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_15"

  // Dividend		divided by		Divisor			=					Quotient
  //  0							/					- 2					=						0

  dividendStr := "0"
  divisorStr := "-2"
  expectedQuoStr := "0"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_16(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_16"

  // Dividend		divided by		Divisor			=					Quotient
  //  0							/					  2					=						0

  dividendStr := "0"
  divisorStr := "2"
  expectedQuoStr := "0"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_17(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_17"

  // Dividend		divided by		Divisor			=					Quotient
  //  11.5				  /				  2.5					=						4.6

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

  err = dividend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "Because expectedNumSeps.Equal(dividendNumSeps) == 'false'\n"+
      "Expected dividendNumSeps = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), dividendNumSeps.String())

    return
  }

  err = dividend.IsValid(ePrefix + "\nValidating dividend #2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = dividend.IsValid(ePrefix)\n"+
      "ID #2\n"+
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
    return
  }

  expectedQuo, err := new(BigIntNum).NewNumStr(expectedQuoStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedQuoStr). "+
      "expectedQuoStr='%v' Error='%v' ",
      expectedQuoStr, err.Error())
  }

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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_18(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_18"

  // Dividend		divided by		Divisor			=					Quotient
  //  11.5				  /				  2.5					=						4.6

  dividendStr := "11.5"
  divisorStr := "2.5"
  expectedQuoStr := "4.6"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_19(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_19"

  // Dividend		divided by		Divisor			=		Quotient
  //    0 				 / 	  			 12.555		  = 	0

  dividendStr := "0"
  divisorStr := "12.555"
  expectedQuoStr := "0"
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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_20(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_20"

  // Dividend		divided by		Divisor			=		Quotient
  //    15.8 			 / 	  			 0		  		= 	ERROR

  dividendStr := "15.8"
  divisorStr := "0"
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

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

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

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  _, err =
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Error: Expected Divide By zero Error.\n"+
      "However, NO ERROR was returned!\n",
      ePrefix)
    return
  }

  return
}

func TestBigIntMathDivide_FixedDecimalFracQuotient_21(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_19"

  // Dividend		 divided by		Divisor							=		Quotient
  // 0.000009218 		 /        35829.8234	     		= 2.572717118108932683156903307539e-10

  dividendStr := "0.000009218"
  divisorStr := "35829.8234"
  expectedQuoStr := "0.0000000002572717118108932683156903307539"

  expectedQuo, err :=
    new(BigIntNum).NewNumStr(expectedQuoStr)

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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoNumStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

  maxPrecision, err := expectedQuo.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "maxPrecision, err := expectedQuo.GetPrecisionUint()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedQuoStr, err.Error())
    return
  }

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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "dividendNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendNumSeps.String(), err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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

func TestBigIntMathDivide_FixedDecimalFracQuotient_22(t *testing.T) {

  ePrefix := "TestBigIntMathDivide_FixedDecimalFracQuotient_22"

  // Dividend		 divided by		Divisor							=		Quotient
  // 35829.8234 	 /        	 0.000009218     		= 3886941136.9060533738338034280755

  dividendStr := "35829.8234"
  divisorStr := "0.000009218"
  expectedQuoStr := "3886941136.9060533738338034280755"

  expectedQuo, err :=
    new(BigIntNum).NewNumStr(expectedQuoStr)

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
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  expectedQuoNumStr, err := expectedQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
      "expectedQuoStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoStr, err.Error())
    return
  }

  maxPrecision, err := expectedQuo.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "maxPrecision, err := expectedQuo.GetPrecisionUint()\n"+
      "expectedQuoNumStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedQuoNumStr, err.Error())
    return
  }

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
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, dividendStr, err.Error())
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
      "divisorStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, divisorStr, err.Error())
    return
  }

  dividendIntValue, err := dividend.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendIntValue, err := dividend.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendPrecisionUint, err := dividend.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendPrecisionUint , err := dividend.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  dividendFixDec, err :=
    new(BigIntFixedDecimal).New(
      dividendIntValue,
      dividendPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  dividendIntValue, dividendPrecisionUint)\n"+
      "dividendIntValue= '%v'\n"+
      "dividendPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendIntValue.Text(10),
      dividendPrecisionUint,
      err.Error())

    return
  }

  dividendFixDecNumStr, err := dividendFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dividendFixDecNumStr, err := dividendFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorIntValue, err := divisor.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorIntValue, err := divisor.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorPrecisionUint, err := divisor.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorPrecisionUint , err := divisor.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  divisorFixDec, err := new(BigIntFixedDecimal).New(
    divisorIntValue,
    divisorPrecisionUint)

  if err != nil {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDec, err := new(BigIntFixedDecimal).New(\n"+
      "  divisorIntValue, divisorPrecisionUint)\n"+
      "divisorIntValue= '%v'\n"+
      "divisorPrecisionUint= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorIntValue.Text(10),
      divisorPrecisionUint,
      err.Error())

    return
  }

  divisorFixDecNumStr, err := divisorFixDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "divisorFixDecNumStr, err := divisorFixDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotient, err :=
    new(BigIntMathDivide).FixedDecimalFracQuotient(
      dividendFixDec,
      divisorFixDec,
      dividendNumSeps,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotient, err := new(BigIntMathDivide).\n"+
      "  FixedDecimalFracQuotient(dividendFixDec, divisorFixDec,\n"+
      "  dividendNumSeps, maxPrecision)\n"+
      "dividendFixDec= '%v'\n"+
      "divisorFixDec= '%v'\n"+
      "dividendNumSeps= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendFixDecNumStr,
      divisorFixDecNumStr,
      dividendNumSeps.String(),
      maxPrecision,
      err.Error())

    return
  }

  fracQuotientIntValue, err := fracQuotient.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientIntValue, err := fracQuotient.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fracQuotientPrecisionUint, err := fracQuotient.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuo, err := new(BigIntNum).NewBigInt(
    fracQuotientIntValue,
    fracQuotientPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuo, err := new(BigIntNum).NewBigInt(\n"+
      "  fracQuotientIntValue, fracQuotientPrecisionUint)\n"+
      "fracQuotientIntValue= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      fracQuotientIntValue.Text(10),
      fracQuotientPrecisionUint,
      err.Error())

    return
  }

  err = actualQuo.IsValid(ePrefix + "\nValidating actualQuo")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = actualQuo.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualQuoNumStr, err := actualQuo.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumStr, err := actualQuo.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedQuoEqualsActualQuo, err := expectedQuo.Equal(actualQuo)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedQuoEqualsActualQuo {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because 'expectedQuoEqualsActualQuo' == 'false'\n"+
      "Expected quotient = '%v'\n"+
      "  Actual quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  if expectedQuoNumStr != actualQuoNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected Quotient = '%v'\n"+
      "  Actual Quotient = '%v'\n\n",
      ePrefix, expectedQuoNumStr, actualQuoNumStr)

    return
  }

  actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualQuoNumSeps, err := actualQuo.GetNumericSeparatorsDto()\n"+
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
