package mathops

import "testing"

func TestBigIntMathDivide_DecimalQuotientMod_01(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumQuotientMod_01"

	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555					/						 2.5			=			 5							 0.055
	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedQuoStr := "5"
	expectedModuloStr := "0.055"
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, modulo, err :=
		new(BigIntMathDivide).DecimalQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalQuotientMod(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
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

	actualModuloNumStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := modulo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

	expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedModuloEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModulo.Equal(modulo) == 'false'\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalQuotientMod_02(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalQuotientMod_02"

	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555  	 			/ 				 	 2  			= 		 6							 0.555
	dividendStr := "12.555"
	divisorStr := "2"
	expectedQuoStr := "6"
	expectedModuloStr := "0.555"
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, modulo, err :=
		new(BigIntMathDivide).DecimalQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalQuotientMod(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
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

	actualModuloNumStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := modulo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

	expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedModuloEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModulo.Equal(modulo) == 'false'\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalQuotientMod_03(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalQuotientMod_03"

	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//	-12.555 				/ 				   2.5 			= 		-5							-0.055
	dividendStr := "-12.555"
	divisorStr := "2.5"
	expectedQuoStr := "-5"
	expectedModuloStr := "-0.055"
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, modulo, err :=
		new(BigIntMathDivide).DecimalQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalQuotientMod(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
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

	actualModuloNumStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := modulo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

	expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedModuloEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModulo.Equal(modulo) == 'false'\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalQuotientMod_04(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalQuotientMod_04"

	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//  -12.555     		/    			 	 2  			= 		-6							-0.555
	dividendStr := "-12.555"
	divisorStr := "2"
	expectedQuoStr := "-6"
	expectedModuloStr := "-0.555"
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, modulo, err :=
		new(BigIntMathDivide).DecimalQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalQuotientMod(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
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

	actualModuloNumStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := modulo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

	expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedModuloEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModulo.Equal(modulo) == 'false'\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalQuotientMod_05(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalQuotientMod_05"

	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	// 	 12.555					/ 				 - 2.5			=			-5							 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedQuoStr := "-5"
	expectedModuloStr := "0.055"
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, modulo, err :=
		new(BigIntMathDivide).DecimalQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalQuotientMod(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
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

	actualModuloNumStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := modulo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

	expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedModuloEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModulo.Equal(modulo) == 'false'\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalQuotientMod_06(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalQuotientMod_06"

	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//   12.555 				/ 				 - 2 				= 		-6							 0.555

	dividendStr := "12.555"
	divisorStr := "-2"
	expectedQuoStr := "-6"
	expectedModuloStr := "0.555"
	maxPrecision := uint(15)

	dividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' Error='%v' ",
			dividendStr, err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())

		return
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = dividend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, modulo, err :=
		new(BigIntMathDivide).DecimalQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalQuotientMod(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
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

	actualModuloNumStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := modulo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

	expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedModuloEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModulo.Equal(modulo) == 'false'\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalQuotientMod_07(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalQuotientMod_07"

	// Dividend			divided by		Divisor			=		Quotient			Modulo/Remainder
	//	-12.555 				/ 				   2.5 			= 		-5							-0.055
	dividendStr := "-12,555"
	divisorStr := "2.5"
	expectedQuoStr := "-5"
	expectedModuloStr := "-0,055"
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

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = dividend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
			"  expectedModuloStr, &expectedNumSeps)\n"+
			"expectedModuloStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
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
		new(BigIntMathDivide).DecimalQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalQuotientMod(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
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

	actualModuloNumStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := modulo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

	expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedModuloEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModulo.Equal(modulo) == 'false'\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalQuotientMod_08(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalQuotientMod_08"

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

	dividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(Decimal).NewNumStrWithNumSeps(\n"+
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, modulo, err :=
		new(BigIntMathDivide).DecimalQuotientMod(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalQuotientMod(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
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

	actualModuloNumStr, err := modulo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := modulo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

	expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloEqualsActualModulo, err := expectedModulo.Equal(modulo)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedModuloEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModulo.Equal(modulo) == 'false'\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected 'modulo' = '%v'\n"+
			"  Actual 'modulo' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalFracQuotient_01(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalFracQuotient_01"

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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, err :=
		new(BigIntMathDivide).DecimalFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalFracQuotient(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalFracQuotient_02(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalFracQuotient_02"

	// Dividend		divided by		Divisor			=		Quotient
	//	-12.555 			/ 					2.5 			= 		-5.022

	dividendStr := "-12.555"
	divisorStr := "2.5"
	expectedQuoStr := "-5.022"
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, err :=
		new(BigIntMathDivide).DecimalFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalFracQuotient(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalFracQuotient_03(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalFracQuotient_03"

	// Dividend		divided by		Divisor			=		Quotient
	//  - 2.5 				/ 			 	12.555		  = 	-0.199123855037834

	dividendStr := "-2.5"
	divisorStr := "12.555"
	expectedQuoStr := "-0.199123855037834"
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, err :=
		new(BigIntMathDivide).DecimalFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalFracQuotient(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalFracQuotient_04(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalFracQuotient_04"

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

	err = dividend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	quotient, err :=
		new(BigIntMathDivide).DecimalFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalFracQuotient(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalFracQuotient_05(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalFracQuotient_05"

	// Dividend		divided by		Divisor			=		Quotient
	// 	 10.5  				/ 					2 				= 	 5.25

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

	dividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(Decimal).NewNumStrWithNumSeps(\n"+
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	quotient, err :=
		new(BigIntMathDivide).DecimalFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalFracQuotient(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalFracQuotient_06(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalFracQuotient_06"

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

	dividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(Decimal).NewNumStrWithNumSeps(\n"+
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
			"dividend= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, dividendNumStr, err.Error())

		return
	}

	err = dividendNumSeps.IsValid(ePrefix + "\nValidating dividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumSeps.IsValid(ePrefix)\n"+
			"dividend= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr,
			dividendNumSeps.String(), err.Error())
		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
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

	expectedQuoNumStr, err := expectedQuo.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoNumStr, err := expectedQuo.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	quotient, err :=
		new(BigIntMathDivide).DecimalFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, modulo, err := new(BigIntMathDivide).DecimalFracQuotient(\n"+
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

	actualQuotientNumStr, err := quotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualQuotientNumStr, err := quotient.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	expectedQuoEqualsActualQuotient, err := expectedQuo.Equal(quotient)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuoEqualsActualQuotient , err =\n"+
			"  expectedQuo.Equal(quotient)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if !expectedQuoEqualsActualQuotient {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedQuo.Equal(quotient) == 'false'\n"+
			"Expected 'quotient' = '%v'\n"+
			"  Actual 'quotient' = '%v'\n\n",
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

func TestBigIntMathDivide_DecimalFracQuotientArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalFracQuotientArray_01"

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

	if len(expectedArrayStr) != lenDividends {

		t.Errorf("%v\n"+
			"Test Configuration Error!\n"+
			"Lengths of 'expectedArrayStr' and 'dividendArrayStr' are not equal.\n"+
			"Length of 'dividendArrayStr' = '%v'\n"+
			"Length of 'expectedArrayStr' = '%v'\n\n",
			ePrefix, lenDividends, len(expectedArrayStr))

		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by\n"+
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
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividends := make([]Decimal, lenDividends)

	expectedResults := make([]Decimal, lenDividends)

	for i := 0; i < lenDividends; i++ {

		dividends[i], err = new(Decimal).NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"  new(Decimal).NewNumStr(dividendArrayStr[%d])\n"+
				"dividendArrayStr[%v]='%v'\nError='%v'\n\n",
				ePrefix, i, i, dividendArrayStr[i], err.Error())
			return
		}

		expectedResults[i], err = new(Decimal).NewNumStr(expectedArrayStr[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				" new(Decimal).NewNumStr(expectedArrayStr[%d])\n"+
				"expectedArrayStr[%v]='%v'\nError='%v'\n\n",
				ePrefix, i, i, expectedArrayStr[i], err.Error())
			return
		}

	}

	resultArray, err := new(BigIntMathDivide).DecimalFracQuotientArray(dividends, divisor, expectedNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultArray, err := new(BigIntMathDivide).\n"+
			"  DecimalFracQuotientArray(\n"+
			"    dividends, divisor, expectedNumSeps, maxPrecision)\n"+
			"divisor= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			divisorNumStr,
			maxPrecision,
			err.Error())

		return
	}

	lenResultArray := len(resultArray)

	if lenDividends != lenResultArray {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenDividends != lenResultArray\n"+
			"Expected Results Array Length = '%v'\n"+
			"  Actual Results Array Length = '%v'\n\n",
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

		actualEqualsExpectedResults, err = resultArray[k].Equal(expectedResults[k])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"actualEqualsExpectedResults, err = resultArray[%d].Equal(expectedResults[%d])\n"+
				"resultsArray[%d]= '%v'\n"+
				"expectedResults[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix,
				k,
				k,
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

func TestBigIntMathDivide_DecimalFracQuotientArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalFracQuotientArray_02"

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

	lenDividends := len(dividendArrayStr)

	if len(expectedArrayStr) != lenDividends {

		t.Errorf("%v\n"+
			"Test Configuration Error!\n"+
			"Lengths of 'expectedArrayStr' and 'dividendArrayStr' are not equal.\n"+
			"Length of 'dividendArrayStr' = '%v'\n"+
			"Length of 'expectedArrayStr' = '%v'\n\n",
			ePrefix, lenDividends, len(expectedArrayStr))

		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by\n"+
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
			"divisorNumStr, err := divisor.GetNumStr()\n"+
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

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	dividends := make([]Decimal, lenDividends)

	expectedResults := make([]Decimal, lenDividends)

	for i := 0; i < lenDividends; i++ {

		dividends[i], err = new(Decimal).NewNumStrWithNumSeps(dividendArrayStr[i], usaNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"  new(Decimal).NewNumStrWithNumSeps(\n"+
				"  dividendArrayStr[%v], usaNumSeps)\n"+
				"dividendArrayStr[%v]= '%v'\n"+
				"usaNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, dividendArrayStr[i], usaNumSeps.String(), err.Error())
			return
		}

		expectedResults[i], err = new(Decimal).NewNumStrWithNumSeps(expectedArrayStr[i], expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				" new(Decimal).NewNumStrWithNumSeps(expectedArrayStr[%v], expectedNumSeps)\n"+
				"expectedArrayStr[%v]= '%v'\n"+
				"expectedNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, expectedArrayStr[i], expectedNumSeps.String(), err.Error())
			return
		}

	}

	resultArray, err := new(BigIntMathDivide).DecimalFracQuotientArray(dividends, divisor, expectedNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultArray, err := new(BigIntMathDivide).\n"+
			"  DecimalFracQuotientArray(dividends, divisor,\n"+
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

		actualEqualsExpectedResults, err = resultArray[k].Equal(expectedResults[k])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"actualEqualsExpectedResults, err = resultArray[%d].Equal(expectedResults[%d])\n"+
				"resultsArray[%d]= '%v'\n"+
				"expectedResults[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix,
				k,
				k,
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

func TestBigIntMathDivide_DecimalFracQuotientArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalFracQuotientArray_02"

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

	lenDividends := len(dividendArrayStr)

	if len(expectedArrayStr) != lenDividends {

		t.Errorf("%v\n"+
			"Test Configuration Error!\n"+
			"Lengths of 'expectedArrayStr' and 'dividendArrayStr' are not equal.\n"+
			"Length of 'dividendArrayStr' = '%v'\n"+
			"Length of 'expectedArrayStr' = '%v'\n\n",
			ePrefix, lenDividends, len(expectedArrayStr))

		return
	}

	divisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by\n"+
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
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividends := make([]Decimal, lenDividends)

	expectedResults := make([]Decimal, lenDividends)

	for i := 0; i < lenDividends; i++ {

		dividends[i], err = new(Decimal).NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"  new(Decimal).NewNumStr(dividendArrayStr[%v])\n"+
				"dividendArrayStr[%v]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, dividendArrayStr[i], err.Error())
			return
		}

		expectedResults[i], err = new(Decimal).NewNumStr(expectedArrayStr[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				" new(Decimal).NewNumStr(expectedArrayStr[%v])\n"+
				"expectedArrayStr[%v]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, expectedArrayStr[i], err.Error())
			return
		}

	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	resultArray, err := new(BigIntMathDivide).DecimalFracQuotientArray(dividends, divisor, expectedNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultArray, err := new(BigIntMathDivide).\n"+
			"  DecimalFracQuotientArray(dividends, divisor,\n"+
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

		actualEqualsExpectedResults, err = resultArray[k].Equal(expectedResults[k])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"actualEqualsExpectedResults, err = resultArray[%d].Equal(expectedResults[%d])\n"+
				"resultsArray[%d]= '%v'\n"+
				"expectedResults[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix,
				k,
				k,
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

func TestBigIntMathDivide_DecimalModulo_01(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModulo_01"

	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloBINum.IsValid(ePrefix + "\nValidating expectedModulo")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloBINum.IsValid(ePrefix + "\nValidating moduloBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)\n"+
			"expectedModuloBINum= '%v'\n"+
			"actualModuloBINum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Modulo = '%v'\n"+
			"  Actual Modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Modulo Number String = '%v'\n"+
			"  Actual Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModulo_02(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModulo_02"

	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   -12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "-0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloBINum.IsValid(ePrefix + "\nValidating expectedModulo")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloBINum.IsValid(ePrefix + "\nValidating moduloBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)\n"+
			"expectedModuloBINum= '%v'\n"+
			"actualModuloBINum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Modulo = '%v'\n"+
			"  Actual Modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Modulo Number String = '%v'\n"+
			"  Actual Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModulo_03(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModulo_03"

	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloBINum.IsValid(ePrefix + "\nValidating expectedModuloBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloBINum.IsValid(ePrefix + "\nValidating actualModuloBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)\n"+
			"expectedModuloBINum= '%v'\n"+
			"actualModuloBINum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Modulo = '%v'\n"+
			"  Actual Modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Modulo Number String = '%v'\n"+
			"  Actual Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModulo_04(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModulo_04"

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

	err := expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloBINum.IsValid(ePrefix + "\nValidating expectedModulo")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, expectedNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(\n"+
			"  decDividend, decDivisor, expectedNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			expectedNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloBINum.IsValid(ePrefix + "\nValidating moduloBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)\n"+
			"expectedModuloBINum= '%v'\n"+
			"actualModuloBINum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Modulo = '%v'\n"+
			"  Actual Modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Modulo Number String = '%v'\n"+
			"  Actual Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
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
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModulo_05(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModulo_05"

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

	err := expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)\n"+
			"dividendStr='%v'\n"+
			"expectedNumSeps='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBINum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloBINum.IsValid(ePrefix + "\nValidating expectedModulo")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloBINum.IsValid(ePrefix + "\nValidating moduloBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)\n"+
			"expectedModuloBINum= '%v'\n"+
			"actualModuloBINum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Modulo = '%v'\n"+
			"  Actual Modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Modulo Number String = '%v'\n"+
			"  Actual Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModulo_06(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModulo_06"

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

	err := expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)\n"+
			"dividendStr='%v'\n"+
			"expectedNumSeps='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(decDividendNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps.Equal(decDividendNumSeps) == 'false'\n"+
			"Expected decDividendNumSeps = '%v'\n"+
			"  Actual decDividendNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decDividendNumSeps.String())

		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloBINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)\n"+
			"expectedModuloStr='%v'\n"+
			"expectedNumSeps='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedModuloBINum.IsValid(ePrefix + "\nValidating expectedModulo")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedModuloStr != expectedModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloStr != expectedModuloNumStr"+
			"The extracted BigIntNum Num String != Original Expected Modulo String\n"+
			"Expected BigInt Num String = '%v'\n"+
			"  Actual BigInt Num String = '%v'\n\n",
			ePrefix, expectedModuloStr, expectedModuloNumStr)

		return
	}

	actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloBINum.IsValid(ePrefix + "\nValidating moduloBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)\n"+
			"expectedModuloBINum= '%v'\n"+
			"actualModuloBINum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Modulo = '%v'\n"+
			"  Actual Modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Modulo Number String = '%v'\n"+
			"  Actual Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModulo_07(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModulo_07"

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

	err := expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)\n"+
			"dividendStr='%v'\n"+
			"expectedNumSeps='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(decDividendNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps.Equal(decDividendNumSeps) == 'false'\n"+
			"Expected decDividendNumSeps = '%v'\n"+
			"  Actual decDividendNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decDividendNumSeps.String())

		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloBINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloBINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &expectedNumSeps)\n"+
			"expectedModuloStr='%v'\n"+
			"expectedNumSeps='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedModuloBINum.IsValid(ePrefix + "\nValidating expectedModulo")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedModuloStr != expectedModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloStr != expectedModuloNumStr"+
			"The extracted BigIntNum Num String != Original Expected Modulo String\n"+
			"Expected BigInt Num String = '%v'\n"+
			"  Actual BigInt Num String = '%v'\n\n",
			ePrefix, expectedModuloStr, expectedModuloNumStr)

		return
	}

	actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModulo(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloBINum.IsValid(ePrefix + "\nValidating moduloBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloBINum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloBINum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloBINum.Equal(actualModuloBINum)\n"+
			"expectedModuloBINum= '%v'\n"+
			"actualModuloBINum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Modulo = '%v'\n"+
			"  Actual Modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Modulo Number String = '%v'\n"+
			"  Actual Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloBINum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModuloToDecimal_01(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModuloToDecimal_01"

	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloDecNum, err := new(Decimal).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloDecNum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloDecNum.IsValid(ePrefix + "\nValidating expectedModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloDecNum, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModuloToDecimal(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloDecNum.IsValid(ePrefix + "\nValidating actualModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)\n"+
			"expectedModuloDecNum= '%v'\n"+
			"actualModuloDecNum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Decimal Modulo = '%v'\n"+
			"  Actual Decimal = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Decimal Modulo Number String = '%v'\n"+
			"  Actual Decimal Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Decimal Modulo Numeric Separators = '%v'\n"+
			"  Actual Decimal Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModuloToDecimal_02(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModuloToDecimal_02"

	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   -12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "-0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloDecNum, err := new(Decimal).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloDecNum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloDecNum.IsValid(ePrefix + "\nValidating expectedModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloDecNum, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModuloToDecimal(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloDecNum.IsValid(ePrefix + "\nValidating actualModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)\n"+
			"expectedModuloDecNum= '%v'\n"+
			"actualModuloDecNum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Decimal Modulo = '%v'\n"+
			"  Actual Decimal = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Decimal Modulo Number String = '%v'\n"+
			"  Actual Decimal Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Decimal Modulo Numeric Separators = '%v'\n"+
			"  Actual Decimal Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModuloToDecimal_03(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModuloToDecimal_02"

	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
			"dividendStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloDecNum, err := new(Decimal).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloDecNum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloDecNum.IsValid(ePrefix + "\nValidating expectedModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloDecNum, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModuloToDecimal(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloDecNum.IsValid(ePrefix + "\nValidating actualModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)\n"+
			"expectedModuloDecNum= '%v'\n"+
			"actualModuloDecNum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Decimal Modulo = '%v'\n"+
			"  Actual Decimal = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Decimal Modulo Number String = '%v'\n"+
			"  Actual Decimal Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumSeps, err := decDividendNumSeps.CopyOut(false)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedModuloNumSeps.Equal(actualModuloNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Decimal Modulo Numeric Separators = '%v'\n"+
			"  Actual Decimal Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModuloToDecimal_04(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModuloToDecimal_04"

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

	err := expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStr(dividendStr)\n"+
			"dividendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	err = decDividend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(decDividendNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps.Equal(decDividendNumSeps) = 'false'\n"+
			"Expected decDividendNumSeps = '%v'\n"+
			"  Actual decDividendNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decDividendNumSeps.String())

		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloDecNum, err := new(Decimal).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloDecNum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloDecNum.IsValid(ePrefix + "\nValidating expectedModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloDecNum, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModuloToDecimal(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloDecNum.IsValid(ePrefix + "\nValidating actualModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)\n"+
			"expectedModuloDecNum= '%v'\n"+
			"actualModuloDecNum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Decimal Modulo = '%v'\n"+
			"  Actual Decimal = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Decimal Modulo Number String = '%v'\n"+
			"  Actual Decimal Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
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
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Decimal Modulo Numeric Separators = '%v'\n"+
			"  Actual Decimal Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModuloToDecimal_05(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModuloToDecimal_05"

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

	err := expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)\n"+
			"dividendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, dividendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(decDividendNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps.Equal(decDividendNumSeps) = 'false'\n"+
			"Expected decDividendNumSeps = '%v'\n"+
			"  Actual decDividendNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decDividendNumSeps.String())

		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloDecNum, err := new(Decimal).NewNumStrWithNumSeps(expectedModuloStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloDecNum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, expectedNumSeps)\n"+
			"expectedModuloStr='%v'\n"+
			"expectedNumSeps='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedModuloDecNum.IsValid(ePrefix + "\nValidating expectedModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloDecNum, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModuloToDecimal(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloDecNum.IsValid(ePrefix + "\nValidating actualModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)\n"+
			"expectedModuloDecNum= '%v'\n"+
			"actualModuloDecNum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Decimal Modulo = '%v'\n"+
			"  Actual Decimal = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Decimal Modulo Number String = '%v'\n"+
			"  Actual Decimal Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
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
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Decimal Modulo Numeric Separators = '%v'\n"+
			"  Actual Decimal Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_DecimalModuloToDecimal_06(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_DecimalModuloToDecimal_06"

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

	err := expectedNumSeps.IsValid(ePrefix + "\nValidating expectedNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)\n"+
			"dividendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, dividendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = decDividend.IsValid(ePrefix + "\nValidating decDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividend.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumStr, err := decDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumStr, err := decDividend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDividendNumSeps, err := decDividend.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = decDividendNumSeps.IsValid(ePrefix + "\nValidating decDividendNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDividendNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if !expectedNumSeps.Equal(decDividendNumSeps) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumSeps.Equal(decDividendNumSeps) = 'false'\n"+
			"Expected decDividendNumSeps = '%v'\n"+
			"  Actual decDividendNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decDividendNumSeps.String())

		return
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decDivisor, err := new(Decimal).NewNumStr(divisorStr)\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = decDivisor.IsValid(ePrefix + "\nValidating decDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decDivisor.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	decDivisorNumStr, err := decDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStr, err := divisor.GetNumStr()\n"+
			"divisorStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	expectedModuloDecNum, err := new(Decimal).NewNumStr(expectedModuloStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloDecNum, err := new(BigIntNum).NewNumStr(expectedModuloStr)\n"+
			"expectedModuloStr='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, err.Error())
		return
	}

	err = expectedModuloDecNum.IsValid(ePrefix + "\nValidating expectedModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModuloNumStr, err := expectedModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloDecNum, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, decDividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloBINum, err := new(BigIntMathDivide).DecimalModuloToDecimal(\n"+
			"  decDividend, decDivisor, decDividendNumSeps, maxPrecision)\n"+
			"decDividend= '%v'\n"+
			"decDivisor= '%v'\n"+
			"decDividendNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, decDividendNumStr, decDivisorNumStr,
			decDividendNumSeps.String(), maxPrecision, err.Error())
		return
	}

	err = actualModuloDecNum.IsValid(ePrefix + "\nValidating actualModuloDecNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualModuloDecNum.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	actualModuloNumStr, err := actualModuloDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumStr, err := actualModuloDecNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsActualModulo, err := expectedModuloDecNum.Equal(actualModuloDecNum)\n"+
			"expectedModuloDecNum= '%v'\n"+
			"actualModuloDecNum= '%v'\n"+
			"Error='%v'\n\n", ePrefix, expectedModuloNumStr,
			actualModuloNumStr, err.Error())
		return
	}

	if !expectedEqualsActualModulo {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedEqualsActualModulo == 'false'\n"+
			"Expected Decimal Modulo = '%v'\n"+
			"  Actual Decimal = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	if expectedModuloNumStr != actualModuloNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedModuloNumStr != actualModuloNumStr\n"+
			"Expected Decimal Modulo Number String = '%v'\n"+
			"  Actual Decimal Modulo Number String = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

		return
	}

	actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualModuloNumSeps, err := actualModuloDecNum.GetNumericSeparatorsDto()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
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
			"Because expectedModuloNumSeps.Equal(actualModuloNumSeps) == 'false'\n"+
			"Expected Decimal Modulo Numeric Separators = '%v'\n"+
			"  Actual Decimal Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}
