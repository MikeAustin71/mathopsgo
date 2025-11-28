package mathops

import "testing"

func TestIntAryMathSubtract_SubtractTotal_01(t *testing.T) {

	ePrefix := "TestIntAryMathSubtract_SubtractTotal_01"

	originalNumberStr1 := "25.72"

	originalNumberStr2 := "8.4"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "17.32"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result Testing Intial IntAry Number String!\n"+
			"Because originalNumberStr1 != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, intAry1NumberStr)

		return
	}

	intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
			"originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = intAry2.IsValid("Validating initial intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating initial intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"intAry2 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result Testing Intial IntAry Number String!\n"+
			"Because originalNumberStr2 != intAry2NumberStr\n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAry2NumberStr)

		return
	}

	err = new(IntAryMathSubtract).SubtractTotal(&intAry1, true, &intAry2, true, true)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = new(IntAryMathSubtract).SubtractTotal(\n"+
			"  &intAry1, validateintAry1=true, &intAry2,\n"+
			"  validateintAry2=true, validateResult=true)\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error='%v'\n\n", ePrefix,
			intAry1NumberStr,
			intAry2NumberStr,
			err.Error())
		return
	}

	err = intAry1.IsValid("Validating final intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating final intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err = intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 Number String set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1PrecisionInt := intAry1.GetPrecision()

	intAry1PrecisionUint, err := intAry1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1PrecisionUint, err :=\n"+
			"  intAry1.GetPrecisionUint()\n"+
			"intAry1Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry1SignValue, err := intAry1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1SignValue, err := intAry1.GetSign()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
			"intAry1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry1NumberStr \n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry1NumberStr)

		return
	}

	if expectedPrecisionInt != intAry1PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry1PrecisionInt\n"+
			"Expected intAry1PrecisionInt = '%v'\n"+
			"  Actual intAry1PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry1PrecisionInt)

		return
	}

	if expectedPrecisionUint != intAry1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry1 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry1PrecisionUint\n"+
			"Expected intAry1PrecisionUint = '%v'\n"+
			"  Actual intAry1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry1PrecisionUint)

		return
	}

	if expectedSignValue != intAry1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry1SignValue\n"+
			"Expected intAry1SignValue = '%v'\n"+
			"  Actual intAry1SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry1SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry1NumSeps \n"+
			"Expected intAry1NumSeps = '%v'\n"+
			"  Actual intAry1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry1NumSeps.String())

		return
	}

	return
}

func TestIntAryMathSubtract_Subtract_01(t *testing.T) {

	ePrefix := "TestIntAryMathSubtract_Subtract_01"

	originalNumberStr1 := "25.72"

	originalNumberStr2 := "8.4"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "17.32"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result Testing Intial IntAry Number String!\n"+
			"Because originalNumberStr1 != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, intAry1NumberStr)

		return
	}

	intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
			"originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = intAry2.IsValid("Validating initial intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating initial intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"intAry2 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result Testing Intial IntAry Number String!\n"+
			"Because originalNumberStr2 != intAry2NumberStr\n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAry2NumberStr)

		return
	}

	intAry3, err := new(IntAryMathSubtract).Subtract(&intAry1, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3, err := new(IntAryMathSubtract).Subtract(&intAry1, &intAry2)\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAry3.IsValid("Validating final intAry3")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry3.IsValid('Validating final intAry3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry3NumberStr, err := intAry3.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3NumberStr, err := intAry3.GetNumStr()\n"+
			"intAry3 Number String set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry3PrecisionInt := intAry3.GetPrecision()

	intAry3PrecisionUint, err := intAry3.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3PrecisionUint, err :=\n"+
			"  intAry3.GetPrecisionUint()\n"+
			"intAry3Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry3NumberStr, err.Error())
		return
	}

	intAry3SignValue, err := intAry3.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3SignValue, err := intAry3.GetSign()\n"+
			"intAry3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry3NumberStr, err.Error())
		return
	}

	intAry3NumSeps, err := intAry3.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3NumSeps, err := intAry3.GetNumericSeparatorsDto()\n"+
			"intAry3= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry3NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry3NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry3NumberStr \n"+
			"Expected intAry3NumberStr = '%v'\n"+
			"  Actual intAry3NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry3NumberStr)

		return
	}

	if expectedPrecisionInt != intAry3PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry3PrecisionInt\n"+
			"Expected intAry3PrecisionInt = '%v'\n"+
			"  Actual intAry3PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry3PrecisionInt)

		return
	}

	if expectedPrecisionUint != intAry3PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAry3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry3PrecisionUint\n"+
			"Expected intAry3PrecisionUint = '%v'\n"+
			"  Actual intAry3PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry3PrecisionUint)

		return
	}

	if expectedSignValue != intAry3SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry3 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry3SignValue\n"+
			"Expected intAry3SignValue = '%v'\n"+
			"  Actual intAry3SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry3SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry3NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry3NumSeps \n"+
			"Expected intAry3NumSeps = '%v'\n"+
			"  Actual intAry3NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry3NumSeps.String())

		return
	}

	return
}
