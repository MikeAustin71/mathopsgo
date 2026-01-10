package mathops

import "testing"

func TestBigIntMathDivide_BigIntNumModulo_01(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_01"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//   12.555					%						 2.5			=			 0.055
	dividendStr := "12.555"
	divisorStr := "2.5"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_02(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_02"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//   12.555  	 			% 				 	 2  			= 		 0.555

	dividendStr := "12.555"
	divisorStr := "2"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_03(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_03"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//    2.5 					% 				 	12.555		= 	   2.500

	dividendStr := "2.5"
	divisorStr := "12.555"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_04(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_04"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//	-12.555 				% 				   2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "2.5"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_05(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_05"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//  -12.555     		%    			 	 2  			= 		-0.555

	dividendStr := "-12.555"
	divisorStr := "2"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_06(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_06"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//  - 2.5 					% 				 	12.555		= 		-2.500

	dividendStr := "-2.5"
	divisorStr := "12.555"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_07(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_07"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	// 	 12.555					% 				 - 2.5			=			 0.055

	dividendStr := "12.555"
	divisorStr := "-2.5"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_08(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_08"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//   12.555 				% 				 - 2 				= 		 0.555

	dividendStr := "12.555"
	divisorStr := "-2"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_09(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_09"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//    2.5 				  % 				 -12.555		= 		 2.5

	dividendStr := "2.5"
	divisorStr := "-12.555"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_10(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_10"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	// 	-12.555 				% 				 - 2.5 			= 		-0.055

	dividendStr := "-12.555"
	divisorStr := "-2.5"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_11(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_11"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//  -12.555     		%    			 - 2 				= 		-0.555

	dividendStr := "-12.555"
	divisorStr := "-2"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_12(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_12"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//  - 2.5	 					% 				 -12.555		= 		-2.5

	dividendStr := "-2.5"
	divisorStr := "-12.555"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_13(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_13"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//    0	 					% 				  -12.555		  = 		 0

	dividendStr := "0"
	divisorStr := "-12.555"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).BigIntNumModulo"+
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_14(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_14"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//    0	 					% 				   12.555		  = 		0

	dividendStr := "0"
	divisorStr := "12.555"
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathDivide).BigIntNumModulo"+
			"(dividend, divisor, maxPrecision).  Error='%v' ",
			err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_15(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_15"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//   12.555					%						 2.5			=			 0.055
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

	err := frenchNumSeps.IsValid(ePrefix + "\nValidating frenchNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = frenchNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividend, err := new(BigIntNum).NewNumStrWithNumSeps(dividendStr, &frenchNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(dividendStr, &frenchNumSeps)\n"+
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

	expectedModulo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedModuloStr, &frenchNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedModulo, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedModuloStr, &frenchNumSeps)\n"+
			"expectedModuloStr='%v'\n"+
			"frenchNumSeps='%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedModuloStr, frenchNumSeps.String(), err.Error())
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, frenchNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, frenchNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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

	expectedModuloNumSeps, err := frenchNumSeps.CopyOut(false)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_16(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_16"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//   12.555					%						 2.5			=			 0.055
	dividendStr := "12.555"
	divisorStr := "2.5"
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

	frenchNumSeps := NumericSeparatorDto{}

	frenchNumSeps.DecimalSeparator = '.'
	frenchNumSeps.ThousandsSeparator = ','
	frenchNumSeps.CurrencySymbol = '$'

	err = frenchNumSeps.IsValid(ePrefix + "\nValidating frenchNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = frenchNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = dividend.SetNumericSeparatorsDto(frenchNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividend.SetNumericSeparatorsDto(frenchNumSeps)\n"+
			"frenchNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, frenchNumSeps.String(), err.Error())
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumModulo_17(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumModulo_17"

	// Dividend			 mod by				Divisor			=		Modulo/Remainder
	//   12.555  	 			% 				 	 2  			= 		 0.555

	dividendStr := "12.555"
	divisorStr := "2"
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

	divisorNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	divisorNumSeps.DecimalSeparator = frenchDecSeparator
	divisorNumSeps.ThousandsSeparator = frenchThousandsSeparator
	divisorNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = divisorNumSeps.IsValid(ePrefix + "\nValidating divisorNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = divisorNumSeps.IsValid(ePrefix)\n"+
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

	err = divisor.IsValid(ePrefix + "\nValidating divisor #1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = divisor.IsValid(ePrefix)\n"+
			"Validating divisor #1\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = divisor.SetNumericSeparatorsDto(divisorNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = divisor.SetNumericSeparatorsDto(divisorNumSeps)\n"+
			"divisorNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, divisorNumSeps.String(), err.Error())
		return
	}

	err = divisor.IsValid(ePrefix + "\nValidating divisor #2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = divisor.IsValid(ePrefix)\n"+
			"Validating divisor #2\n"+
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

	modulo, err :=
		new(BigIntMathDivide).BigIntNumModulo(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"modulo, err := new(BigIntMathDivide).BigIntNumModulo(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)"+
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
			"Expected modulo = '%v'\n"+
			"  Actual modulo = '%v'\n\n",
			ePrefix, expectedModuloNumStr, actualModuloNumStr)

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
			"Expected Modulo Numeric Separators = '%v'\n"+
			"  Actual Modulo Numeric Separators = '%v'\n\n",
			ePrefix, expectedModuloNumSeps.String(), actualModuloNumSeps.String())

		return
	}

	return
}

func TestBigIntMathDivide_BigIntNumFracQuotient_01(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_01"

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

	quotient, err :=
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
			"  dividend, divisor, dividendNumSeps, maxPrecision)\n"+
			"dividend= '%v'\n"+
			"divisor= '%v'\n"+
			"dividendNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, dividendNumStr, divisorNumStr, dividendNumSeps.String(), err.Error())
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

func TestBigIntMathDivide_BigIntNumFracQuotient_02(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_02"

	// Dividend		divided by		Divisor			=		Quotient
	// 	 10    				/ 					2 				= 			5

	dividendStr := "10"
	divisorStr := "2"
	expectedQuoStr := "5"
	maxPrecision := uint(15)

	expectedDividendNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	err := expectedDividendNumSeps.IsValid("Validating expectedDividendNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedDividendNumSeps.IsValid(\n"+
			"  'Validating expectedDividendNumSeps')\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividend, err := new(BigIntNum).NewNumStrWithNumSeps(
		dividendStr, &expectedDividendNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
			"  dividendStr, &expectedDividendNumSeps)\n"+
			"dividendStr='%v'\n"+
			"expectedDividendNumSeps= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix, dividendStr, expectedDividendNumSeps.String(), err.Error())
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

	if !expectedDividendNumSeps.Equal(dividendNumSeps) {

		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedDividendNumSeps.Equal(dividendNumSeps) == 'false'\n"+
			"Expected Dividend Numeric Separators = '%v'\n"+
			"  Actual Dividend Numeric Separators = '%v'\n\n",
			ePrefix, expectedDividendNumSeps.String(), dividendNumSeps.String())

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_03(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_03"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_04(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_04"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_05(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_05"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_06(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_06"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_07(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_07"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_08(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_08"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_09(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_09"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_10(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_10"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_11(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_11"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_12(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_12"

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
		t.Errorf("Error returned by:\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_13(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_13"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_14(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_14"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_15(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_15"

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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_16(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_16"

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

	frenchNumSeps := NumericSeparatorDto{}

	frenchNumSeps.DecimalSeparator = '.'
	frenchNumSeps.ThousandsSeparator = ','
	frenchNumSeps.CurrencySymbol = '$'

	err = frenchNumSeps.IsValid(ePrefix + "\nValidating frenchNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = frenchNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = dividend.SetNumericSeparatorsDto(frenchNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividend.SetNumericSeparatorsDto(frenchNumSeps)\n"+
			"frenchNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, frenchNumSeps.String(), err.Error())
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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_17(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_17"

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

	frenchNumSeps := NumericSeparatorDto{}

	frenchNumSeps.DecimalSeparator = '.'
	frenchNumSeps.ThousandsSeparator = ','
	frenchNumSeps.CurrencySymbol = '$'

	err = frenchNumSeps.IsValid(ePrefix + "\nValidating frenchNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = frenchNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	err = dividend.SetNumericSeparatorsDto(frenchNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividend.SetNumericSeparatorsDto(frenchNumSeps)\n"+
			"frenchNumSeps= '%v'\n"+
			"Error='%v'\n\n", ePrefix, frenchNumSeps.String(), err.Error())
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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotient_18(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotient_18"

	// Dividend		divided by		Divisor			=					Quotient
	//  11.5				  /				  2.5					=						4.6

	dividendStr := "11,5"
	divisorStr := "2.5"
	expectedQuoStr := "4,6"
	maxPrecision := uint(15)

	frenchNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	frenchNumSeps.DecimalSeparator = frenchDecSeparator
	frenchNumSeps.ThousandsSeparator = frenchThousandsSeparator
	frenchNumSeps.CurrencySymbol = frenchCurrencySymbol

	err := frenchNumSeps.IsValid(ePrefix + "\nValidating frenchNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = frenchNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividend, err := new(BigIntNum).NewNumStrWithNumSeps(dividendStr, &frenchNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividend, err := new(BigIntNum).NewNumStrWithNumSeps(dividendStr, &frenchNumSeps)\n"+
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

	expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(expectedQuoStr, &frenchNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedQuo, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
			"  expectedQuoStr, &frenchNumSeps)\n"+
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
		new(BigIntMathDivide).BigIntNumFracQuotient(dividend, divisor, dividendNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotient, err := new(BigIntMathDivide).BigIntNumFracQuotient(\n"+
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

func TestBigIntMathDivide_BigIntNumFracQuotientArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotientArray_01"

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

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
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

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	dividends := make([]BigIntNum, lenDividends)

	expectedResults := make([]BigIntNum, lenDividends)

	for i := 0; i < lenDividends; i++ {

		dividends[i], err = new(BigIntNum).NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"  new(BigIntNum).NewNumStr(dividendArrayStr[%d])\n"+
				"dividendArrayStr[%v]='%v'\nError='%v'\n\n",
				ePrefix, i, i, dividendArrayStr[i], err.Error())
			return
		}

		expectedResults[i], err = new(BigIntNum).NewNumStr(expectedArrayStr[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				" new(BigIntNum).NewNumStr(expectedArrayStr[%d])\n"+
				"expectedArrayStr[%v]='%v'\nError='%v'\n\n",
				ePrefix, i, i, expectedArrayStr[i], err.Error())
			return
		}

	}

	resultArray, err := new(BigIntMathDivide).BigIntNumFracQuotientArray(dividends, divisor, usaNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultArray, err := new(BigIntMathDivide).\n"+
			"  BigIntNumFracQuotientArray(\n"+
			"    dividends, divisor, usaNumSeps, maxPrecision)\n"+
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
	}

	return
}

func TestBigIntMathDivide_BigIntNumFracQuotientArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotientArray_02"

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

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
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

	frenchNumSeps := NumericSeparatorDto{}
	frenchNumSeps.DecimalSeparator = '.'
	frenchNumSeps.ThousandsSeparator = ','
	frenchNumSeps.CurrencySymbol = '$'

	err = frenchNumSeps.IsValid(ePrefix + "\nValidating frenchNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = frenchNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	dividends := make([]BigIntNum, lenDividends)

	expectedResults := make([]BigIntNum, lenDividends)

	for i := 0; i < lenDividends; i++ {

		dividends[i], err = new(BigIntNum).NewNumStr(dividendArrayStr[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"  new(BigIntNum).NewNumStr(dividendArrayStr[%d])\n"+
				"dividendArrayStr[%v]='%v'\nError='%v'\n\n",
				ePrefix, i, i, dividendArrayStr[i], err.Error())
			return
		}

		expectedResults[i], err = new(BigIntNum).NewNumStr(expectedArrayStr[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				" new(BigIntNum).NewNumStr(expectedArrayStr[%d])\n"+
				"expectedArrayStr[%v]='%v'\nError='%v'\n\n",
				ePrefix, i, i, expectedArrayStr[i], err.Error())
			return
		}

	}

	resultArray, err := new(BigIntMathDivide).BigIntNumFracQuotientArray(dividends, divisor, frenchNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultArray, err := new(BigIntMathDivide).\n"+
			"  BigIntNumFracQuotientArray(\n"+
			"    dividends, divisor, frenchNumSeps, maxPrecision)\n"+
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

		if !frenchNumSeps.Equal(actualNumSeps) {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because frenchNumSeps.Equal(actualNumSeps) == 'false'\n"+
				"Expected Numeric Separators = '%v'\n"+
				"  Actual Numeric Separators = '%v'\n\n",
				ePrefix, frenchNumSeps.String(), actualNumSeps.String())

			return
		}
	}

	return
}

func TestBigIntMathDivide_BigIntNumFracQuotientArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathDivide_BigIntNumFracQuotientArray_03"

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

	divisor, err := new(BigIntNum).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by\n"+
			"divisor, err := new(BigIntNum).NewNumStr(divisorStr)\n"+
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

	frenchNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	frenchNumSeps.DecimalSeparator = frenchDecSeparator
	frenchNumSeps.ThousandsSeparator = frenchThousandsSeparator
	frenchNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = frenchNumSeps.IsValid(ePrefix + "\nValidating frenchNumSeps")

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = frenchNumSeps.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	dividends := make([]BigIntNum, lenDividends)

	expectedResults := make([]BigIntNum, lenDividends)

	for i := 0; i < lenDividends; i++ {

		dividends[i], err = new(BigIntNum).NewNumStrWithNumSeps(dividendArrayStr[i], &usaNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"  new(BigIntNum).NewNumStrWithNumSeps(\n"+
				"  dividendArrayStr[%d], &usaNumSeps)\n"+
				"dividendArrayStr[%v]= '%v'\n"+
				"usaNumSeps= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, dividendArrayStr[i], usaNumSeps.String(), err.Error())
			return
		}

		expectedResults[i], err = new(BigIntNum).NewNumStrWithNumSeps(expectedArrayStr[i], &frenchNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				" new(BigIntNum).NewNumStrWithNumSeps(\n"+
				"  expectedArrayStr[%d], &frenchNumSeps)\n"+
				"expectedArrayStr[%v]='%v'\n"+
				"frenchNumSeps= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, expectedArrayStr[i],
				frenchNumSeps.String(), err.Error())
			return
		}

	}

	resultArray, err := new(BigIntMathDivide).BigIntNumFracQuotientArray(dividends, divisor, frenchNumSeps, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultArray, err := new(BigIntMathDivide).\n"+
			"  BigIntNumFracQuotientArray(dividends, divisor,\n"+
			"    frenchNumSeps, maxPrecision)\n"+
			"divisor= '%v'\n"+
			"frenchNumSeps= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix,
			divisorNumStr,
			frenchNumSeps.String(),
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

		if !frenchNumSeps.Equal(actualNumSeps) {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because frenchNumSeps.Equal(actualNumSeps) == 'false'\n"+
				"Expected Numeric Separators = '%v'\n"+
				"  Actual Numeric Separators = '%v'\n\n",
				ePrefix, frenchNumSeps.String(), actualNumSeps.String())

			return
		}
	}

	return
}
