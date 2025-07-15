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
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModulo(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_DecimalModulo_02(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   -12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "-0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModulo(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_DecimalModulo_03(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModulo(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_DecimalModulo_04(t *testing.T) {
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

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModulo(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

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

func TestBigIntMathDivide_DecimalModulo_05(t *testing.T) {
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

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModulo(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_DecimalModulo_06(t *testing.T) {
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

	decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStrWithNumSeps"+
			"(dividendStr, expectedNumSeps). dividendStr='%v' error='%v'",
			dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModulo(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_DecimalModulo_07(t *testing.T) {
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

	decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStrWithNumSeps"+
			"(dividendStr, expectedNumSeps). dividendStr='%v' error='%v'",
			dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	moduloBINum, err := new(BigIntMathDivide).DecimalModulo(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModulo(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := moduloBINum.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := moduloBINum.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_DecimalModuloToDecimal_01(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					%						 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	decModulo, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := decModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_DecimalModuloToDecimal_02(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   -12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "-0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	decModulo, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := decModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_DecimalModuloToDecimal_03(t *testing.T) {
	// Dividend			  mod by			Divisor			=			Modulo/Remainder
	// --------				------			-------						----------------
	//   12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
	expectedModuloStr := "0.055"
	maxPrecision := uint(15)

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	decModulo, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := decModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

}

func TestBigIntMathDivide_DecimalModuloToDecimal_04(t *testing.T) {
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

	decDividend, err := new(Decimal).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	decModulo, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := decModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected decModulo='%v'. Instead decModulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := decModulo.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_DecimalModuloToDecimal_05(t *testing.T) {
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

	decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStrWithNumSeps"+
			"(dividendStr, expectedNumSeps). dividendStr='%v' error='%v'",
			dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	decModulo, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := decModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected modulo='%v'. Instead modulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := decModulo.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathDivide_DecimalModuloToDecimal_06(t *testing.T) {
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

	decDividend, err := new(Decimal).NewNumStrWithNumSeps(dividendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(dividendStr). "+
			"dividendStr='%v' error='%v'", dividendStr, err.Error())
	}

	decDivisor, err := new(Decimal).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(divisorStr). "+
			"divisorStr='%v' error='%v'", divisorStr, err.Error())
	}

	decModulo, err := new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, decDivisor, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).DecimalModuloToDecimal(decDividend, "+
			"decDivisor, maxPrecision). Error='%v'", err.Error())

	}

	actualModuloStr := decModulo.GetNumStr()

	if expectedModuloStr != actualModuloStr {
		t.Errorf("Error: Expected decModulo='%v'. Instead decModulo='%v'",
			expectedModuloStr, actualModuloStr)
	}

	actualNumSeps := decModulo.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}
