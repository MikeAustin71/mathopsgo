package mathops

import "testing"

func TestIntAryMathDivide_Divide_01(t *testing.T) {

	// Terminology
	// dividend / divisor = quotient

	ePrefix := "TestIntAryMathDivide_Divide_01"

	dividendStr := "25.6"

	divisorStr := "2.3"

	maxPrecision := 30

	minPrecision := 5

	//                                1         2         3
	//                     0.123456789012345678901234567890
	expectedNumberStr := "11.130434782608695652173913043478"

	expectedPrecisionInt := 30

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryDividend, err := new(IntAry).NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryDividend, err := new(IntAry).NewNumStr(dividendStr)\n"+
			"dividendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, dividendStr, err.Error())
		return
	}

	err = intAryDividend.IsValid("Validating initial intAryDividend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryDividend.IsValid('Validating intAryDividend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryDividendNumberStr, err := intAryDividend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryDividendNumberStr, err := intAryDividend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if dividendStr != intAryDividendNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because dividendStr != intAryDividendNumberStr\n"+
			"Expected intAryDividendNumberStr = '%v'\n"+
			"  Actual intAryDividendNumberStr = '%v'\n\n",
			ePrefix, dividendStr, intAryDividendNumberStr)

		return
	}

	intAryDivisor, err := new(IntAry).NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryDivisor, err := new(IntAry).NewNumStr(divisorStr)\n"+
			"divisorStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, divisorStr, err.Error())
		return
	}

	err = intAryDivisor.IsValid("Validating initial intAryDivisor")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryDivisor.IsValid('Validating intAryDivisor')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryDivisorNumberStr, err := intAryDivisor.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if divisorStr != intAryDivisorNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because divisorStr != intAryDivisorNumberStr\n"+
			"Expected intAryDivisorNumberStr = '%v'\n"+
			"  Actual intAryDivisorNumberStr = '%v'\n\n",
			ePrefix, divisorStr, intAryDivisorNumberStr)

		return
	}

	intAryQuotient, err := new(IntAryMathDivide).Divide(&intAryDividend, &intAryDivisor, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryQuotient, err := new(IntAryMathDivide).Divide(\n"+
			"  &intAryDividend, &intAryDivisor, minPrecision, maxPrecision)\n"+
			"intAryDividend= '%v'\n"+
			"intAryDivisor= '%v'\n"+
			"minPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryDividendNumberStr,
			intAryDivisorNumberStr,
			minPrecision,
			maxPrecision,
			err.Error())

		return
	}

	err = intAryQuotient.IsValid("Validating intAryQuotient")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryQuotient.IsValid('Validating initial intAryQuotient')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryQuotientNumberStr, err := intAryQuotient.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryQuotientPrecisionInt := intAryQuotient.GetPrecision()

	intAryQuotientPrecisionUint, err := intAryQuotient.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryQuotientPrecisionUint, err :=\n"+
			"  intAryQuotient.GetPrecisionUint()\n"+
			"intAryQuotientResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryQuotientNumberStr, err.Error())
		return
	}

	intAryQuotientSignValue, err := intAryQuotient.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryQuotientSignValue, err := intAryQuotient.GetSign()\n"+
			"intAryQuotient= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryQuotientNumberStr, err.Error())
		return
	}

	intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryQuotientNumSeps, err := intAryQuotient.GetNumericSeparatorsDto()\n"+
			"intAryQuotient= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryQuotientNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryQuotientNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryQuotientNumberStr \n"+
			"Expected intAryQuotientNumberStr = '%v'\n"+
			"  Actual intAryQuotientNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryQuotientNumberStr)

		return
	}

	if expectedPrecisionInt != intAryQuotientPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryQuotient Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryQuotientPrecisionInt\n"+
			"Expected intAryQuotientPrecisionInt = '%v'\n"+
			"  Actual intAryQuotientPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryQuotientPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryQuotientPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryQuotient Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryQuotientPrecisionUint\n"+
			"Expected intAryQuotientPrecisionUint = '%v'\n"+
			"  Actual intAryQuotientPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryQuotientPrecisionUint)

		return
	}

	if expectedSignValue != intAryQuotientSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryQuotient Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryQuotientSignValue\n"+
			"Expected intAryQuotientSignValue = '%v'\n"+
			"  Actual intAryQuotientSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryQuotientSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryQuotientNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryQuotientNumSeps \n"+
			"Expected intAryQuotientNumSeps = '%v'\n"+
			"  Actual intAryQuotientNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryQuotientNumSeps.String())

		return
	}

	return
}

func TestIntAryMathDivide_DivideByTwo_01(t *testing.T) {

	// Terminology
	// dividend / divisor = quotient

	ePrefix := "TestIntAryMathDivide_DivideByTwo_01"

	originalNumberStr := "25.63"

	//                                1         2         3
	//                     0.123456789012345678901234567890
	expectedNumberStr := "12.815"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryResult, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAryResult.IsValid("Validating initial intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating initial intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
			"intAryResult set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result Testing Intial IntAry Number String!\n"+
			"Because originalNumberStr != intAryResultNumberStr\n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryResultNumberStr)

		return
	}

	err = new(IntAryMathDivide).DivideByTwo(&intAryResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = new(IntAryMathDivide).DivideByTwo(\n"+
			"  &intAryResult)\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	err = intAryResult.IsValid("Validating final intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating final intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err = intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
			"intAryResult Number String set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultPrecisionInt := intAryResult.GetPrecision()

	intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultPrecisionUint, err :=\n"+
			"  intAryResult.GetPrecisionUint()\n"+
			"intAryResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultSignValue, err := intAryResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultSignValue, err := intAryResult.GetSign()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
			"intAryResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionInt != intAryResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
			"Expected intAryResultPrecisionInt = '%v'\n"+
			"  Actual intAryResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignValue != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}
