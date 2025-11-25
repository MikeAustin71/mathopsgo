package mathops

import "testing"

func TestIntAryMathAdd_Add_01(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_Add_01"

	originalNumberStr1 := "45.762"

	originalNumberStr2 := "12.851"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "58.613"

	expectedPrecisionInt := 3

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

	intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryFinal.IsValid("Validating intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err := intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err := intAryFinal.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalPrecisionInt := intAryFinal.GetPrecision()

	intAryFinalPrecisionUint, err := intAryFinal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalPrecisionUint, err :=\n"+
			"  intAryFinal.GetPrecisionUint()\n"+
			"intAryFinalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalSignValue, err := intAryFinal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalSignValue, err := intAryFinal.GetSign()\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()\n"+
			"intAryFinal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryFinal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryFinalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryFinalPrecisionInt\n"+
			"Expected intAryFinalPrecisionInt = '%v'\n"+
			"  Actual intAryFinalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryFinalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryFinalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryFinalPrecisionUint\n"+
			"Expected intAryFinalPrecisionUint = '%v'\n"+
			"  Actual intAryFinalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryFinalPrecisionUint)

		return
	}

	if expectedSignValue != intAryFinalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryFinalSignValue\n"+
			"Expected intAryFinalSignValue = '%v'\n"+
			"  Actual intAryFinalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryFinalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryFinalNumSeps \n"+
			"Expected intAryFinalNumSeps = '%v'\n"+
			"  Actual intAryFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryFinalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_Add_02(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_Add_02"

	originalNumberStr1 := "45.762"

	originalNumberStr2 := "-12.851"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "32.911"

	expectedPrecisionInt := 3

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

	intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryFinal.IsValid("Validating intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err := intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err := intAryFinal.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalPrecisionInt := intAryFinal.GetPrecision()

	intAryFinalPrecisionUint, err := intAryFinal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalPrecisionUint, err :=\n"+
			"  intAryFinal.GetPrecisionUint()\n"+
			"intAryFinalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalSignValue, err := intAryFinal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalSignValue, err := intAryFinal.GetSign()\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()\n"+
			"intAryFinal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryFinal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryFinalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryFinalPrecisionInt\n"+
			"Expected intAryFinalPrecisionInt = '%v'\n"+
			"  Actual intAryFinalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryFinalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryFinalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryFinalPrecisionUint\n"+
			"Expected intAryFinalPrecisionUint = '%v'\n"+
			"  Actual intAryFinalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryFinalPrecisionUint)

		return
	}

	if expectedSignValue != intAryFinalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryFinalSignValue\n"+
			"Expected intAryFinalSignValue = '%v'\n"+
			"  Actual intAryFinalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryFinalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryFinalNumSeps \n"+
			"Expected intAryFinalNumSeps = '%v'\n"+
			"  Actual intAryFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryFinalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_Add_03(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_Add_03"

	originalNumberStr1 := "-45.762"

	originalNumberStr2 := "-12.851"

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr := "-58.613"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

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

	intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryFinal.IsValid("Validating intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err := intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err := intAryFinal.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalPrecisionInt := intAryFinal.GetPrecision()

	intAryFinalPrecisionUint, err := intAryFinal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalPrecisionUint, err :=\n"+
			"  intAryFinal.GetPrecisionUint()\n"+
			"intAryFinalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalSignValue, err := intAryFinal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalSignValue, err := intAryFinal.GetSign()\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()\n"+
			"intAryFinal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryFinal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryFinalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryFinalPrecisionInt\n"+
			"Expected intAryFinalPrecisionInt = '%v'\n"+
			"  Actual intAryFinalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryFinalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryFinalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryFinalPrecisionUint\n"+
			"Expected intAryFinalPrecisionUint = '%v'\n"+
			"  Actual intAryFinalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryFinalPrecisionUint)

		return
	}

	if expectedSignValue != intAryFinalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryFinalSignValue\n"+
			"Expected intAryFinalSignValue = '%v'\n"+
			"  Actual intAryFinalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryFinalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryFinalNumSeps \n"+
			"Expected intAryFinalNumSeps = '%v'\n"+
			"  Actual intAryFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryFinalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_Add_04(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_Add_04"

	originalNumberStr1 := "0"

	originalNumberStr2 := "12.851"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "12.851"

	expectedPrecisionInt := 3

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

	intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryFinal.IsValid("Validating intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err := intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err := intAryFinal.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalPrecisionInt := intAryFinal.GetPrecision()

	intAryFinalPrecisionUint, err := intAryFinal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalPrecisionUint, err :=\n"+
			"  intAryFinal.GetPrecisionUint()\n"+
			"intAryFinalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalSignValue, err := intAryFinal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalSignValue, err := intAryFinal.GetSign()\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()\n"+
			"intAryFinal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryFinal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryFinalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryFinalPrecisionInt\n"+
			"Expected intAryFinalPrecisionInt = '%v'\n"+
			"  Actual intAryFinalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryFinalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryFinalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryFinalPrecisionUint\n"+
			"Expected intAryFinalPrecisionUint = '%v'\n"+
			"  Actual intAryFinalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryFinalPrecisionUint)

		return
	}

	if expectedSignValue != intAryFinalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryFinalSignValue\n"+
			"Expected intAryFinalSignValue = '%v'\n"+
			"  Actual intAryFinalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryFinalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryFinalNumSeps \n"+
			"Expected intAryFinalNumSeps = '%v'\n"+
			"  Actual intAryFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryFinalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_Add_05(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_Add_05"

	originalNumberStr1 := "45.762"

	originalNumberStr2 := "0"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "45.762"

	expectedPrecisionInt := 3

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

	intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryFinal.IsValid("Validating intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err := intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err := intAryFinal.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalPrecisionInt := intAryFinal.GetPrecision()

	intAryFinalPrecisionUint, err := intAryFinal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalPrecisionUint, err :=\n"+
			"  intAryFinal.GetPrecisionUint()\n"+
			"intAryFinalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalSignValue, err := intAryFinal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalSignValue, err := intAryFinal.GetSign()\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()\n"+
			"intAryFinal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryFinal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryFinalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryFinalPrecisionInt\n"+
			"Expected intAryFinalPrecisionInt = '%v'\n"+
			"  Actual intAryFinalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryFinalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryFinalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryFinalPrecisionUint\n"+
			"Expected intAryFinalPrecisionUint = '%v'\n"+
			"  Actual intAryFinalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryFinalPrecisionUint)

		return
	}

	if expectedSignValue != intAryFinalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryFinalSignValue\n"+
			"Expected intAryFinalSignValue = '%v'\n"+
			"  Actual intAryFinalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryFinalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryFinalNumSeps \n"+
			"Expected intAryFinalNumSeps = '%v'\n"+
			"  Actual intAryFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryFinalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_Add_06(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_Add_06"

	originalNumberStr1 := "45.762"

	originalNumberStr2 := "12,851"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "58.613"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	intAry1, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, expectedNumSeps)\n"+
			"originalNumberStr1= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr1,
			expectedNumSeps.String(),
			err.Error())

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

	intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
			"intAry1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
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

	intAry2, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)\n"+
			"originalNumberStr2= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr2,
			alternateNumSeps.String(),
			err.Error())

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

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if !alternateNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because alternateNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, alternateNumSeps.String(), intAry2NumSeps.String())

		return
	}

	intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryFinal.IsValid("Validating intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err := intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err := intAryFinal.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalPrecisionInt := intAryFinal.GetPrecision()

	intAryFinalPrecisionUint, err := intAryFinal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalPrecisionUint, err :=\n"+
			"  intAryFinal.GetPrecisionUint()\n"+
			"intAryFinalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalSignValue, err := intAryFinal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalSignValue, err := intAryFinal.GetSign()\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()\n"+
			"intAryFinal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryFinal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryFinalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryFinalPrecisionInt\n"+
			"Expected intAryFinalPrecisionInt = '%v'\n"+
			"  Actual intAryFinalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryFinalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryFinalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryFinalPrecisionUint\n"+
			"Expected intAryFinalPrecisionUint = '%v'\n"+
			"  Actual intAryFinalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryFinalPrecisionUint)

		return
	}

	if expectedSignValue != intAryFinalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryFinalSignValue\n"+
			"Expected intAryFinalSignValue = '%v'\n"+
			"  Actual intAryFinalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryFinalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryFinalNumSeps \n"+
			"Expected intAryFinalNumSeps = '%v'\n"+
			"  Actual intAryFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryFinalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_Add_07(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_Add_06"

	originalNumberStr1 := "45,762"

	originalNumberStr2 := "12.851"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "58,613"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	intAry1, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, expectedNumSeps)\n"+
			"originalNumberStr1= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr1,
			expectedNumSeps.String(),
			err.Error())

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

	intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumSeps, err := intAry1.GetNumericSeparatorsDto()\n"+
			"intAry1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry1NumberStr, err.Error())
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

	intAry2, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)\n"+
			"originalNumberStr2= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr2,
			alternateNumSeps.String(),
			err.Error())

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

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if !alternateNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because alternateNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, alternateNumSeps.String(), intAry2NumSeps.String())

		return
	}

	intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinal, err := new(IntAryMathAdd).Add(&intAry1, &intAry2)\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAry1NumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryFinal.IsValid("Validating intAryFinal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryFinal.IsValid('Validating intAryFinal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalNumberStr, err := intAryFinal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err := intAryFinal.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryFinalPrecisionInt := intAryFinal.GetPrecision()

	intAryFinalPrecisionUint, err := intAryFinal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalPrecisionUint, err :=\n"+
			"  intAryFinal.GetPrecisionUint()\n"+
			"intAryFinalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalSignValue, err := intAryFinal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalSignValue, err := intAryFinal.GetSign()\n"+
			"intAryFinal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumSeps, err := intAryFinal.GetNumericSeparatorsDto()\n"+
			"intAryFinal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryFinalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryFinal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryFinalNumberStr \n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryFinalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryFinalPrecisionInt\n"+
			"Expected intAryFinalPrecisionInt = '%v'\n"+
			"  Actual intAryFinalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryFinalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryFinalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryFinal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryFinalPrecisionUint\n"+
			"Expected intAryFinalPrecisionUint = '%v'\n"+
			"  Actual intAryFinalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryFinalPrecisionUint)

		return
	}

	if expectedSignValue != intAryFinalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryFinal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryFinalSignValue\n"+
			"Expected intAryFinalSignValue = '%v'\n"+
			"  Actual intAryFinalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryFinalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryFinalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryFinalNumSeps \n"+
			"Expected intAryFinalNumSeps = '%v'\n"+
			"  Actual intAryFinalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryFinalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_RunTotal_01(t *testing.T) {
	total := IntAry{}.NewZero(0)

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"
	expectedStr := "27.852"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3, err := IntAry{}.NewNumStr(nStr3)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr3). "+
			"nStr3='%v' Error='%v' ", nStr3, err.Error())
	}

	ia4, err := IntAry{}.NewNumStr(nStr4)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr4). "+
			"nStr4='%v' Error='%v' ", nStr4, err.Error())
	}

	IntAryMathAdd{}.RunTotal(&total, &ia1)

	IntAryMathAdd{}.RunTotal(&total, &ia2)

	IntAryMathAdd{}.RunTotal(&total, &ia3)

	IntAryMathAdd{}.RunTotal(&total, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

}

func TestIntAryMathAdd_RunTotal_02(t *testing.T) {

	expectedNumSeps := new(NumericSeparatorDto).New()
	totalStr := "0"
	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	nStr1 := "5,1"
	nStr2 := "21,452"
	nStr3 := "8"
	nStr4 := "-6,7"
	expectedStr := "27.852"

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3, err := IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	ia4, err := IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.RunTotal(&total, &ia1)

	IntAryMathAdd{}.RunTotal(&total, &ia2)

	IntAryMathAdd{}.RunTotal(&total, &ia3)

	IntAryMathAdd{}.RunTotal(&total, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestIntAryMathAdd_RunTotal_03(t *testing.T) {

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	totalStr := "0"
	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"

	expectedStr := "27,852"

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3, err := IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	ia4, err := IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.RunTotal(&total, &ia1)

	IntAryMathAdd{}.RunTotal(&total, &ia2)

	IntAryMathAdd{}.RunTotal(&total, &ia3)

	IntAryMathAdd{}.RunTotal(&total, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestIntAryMathAdd_AddManyArray_01(t *testing.T) {

	total := IntAry{}.NewZero(0)

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"
	expectedStr := "27.852"

	iaArray := make([]IntAry, 4)

	var err error

	iaArray[0], err = IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	iaArray[1], err = IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	iaArray[2], err = IntAry{}.NewNumStr(nStr3)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr3). "+
			"nStr3='%v' Error='%v' ", nStr3, err.Error())
	}

	iaArray[3], err = IntAry{}.NewNumStr(nStr4)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr4). "+
			"nStr4='%v' Error='%v' ", nStr4, err.Error())
	}

	IntAryMathAdd{}.AddManyArray(&total, iaArray)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

}

func TestIntAryMathAdd_AddManyArray_02(t *testing.T) {

	var err error

	expectedNumSeps := new(NumericSeparatorDto).New()

	totalStr := "0"

	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	nStr1 := "5,1"
	nStr2 := "21,452"
	nStr3 := "8"
	nStr4 := "-6,7"
	expectedStr := "27.852"

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	iaArray := make([]IntAry, 4)

	iaArray[0], err = IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	iaArray[1], err = IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	iaArray[2], err = IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	iaArray[3], err = IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.AddManyArray(&total, iaArray)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestIntAryMathAdd_AddManyArray_03(t *testing.T) {

	var err error

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	totalStr := "0"

	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"
	expectedStr := "27,852"

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	iaArray := make([]IntAry, 4)

	iaArray[0], err = IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	iaArray[1], err = IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	iaArray[2], err = IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	iaArray[3], err = IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.AddManyArray(&total, iaArray)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestIntAryMathAdd_AddMany_01(t *testing.T) {
	total := IntAry{}.NewZero(0)

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"
	expectedStr := "27.852"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3, err := IntAry{}.NewNumStr(nStr3)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr3). "+
			"nStr3='%v' Error='%v' ", nStr3, err.Error())
	}

	ia4, err := IntAry{}.NewNumStr(nStr4)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr4). "+
			"nStr4='%v' Error='%v' ", nStr4, err.Error())
	}

	IntAryMathAdd{}.AddMany(&total, &ia1, &ia2, &ia3, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}
}

func TestIntAryMathAdd_AddMany_02(t *testing.T) {

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	totalStr := "0"

	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	nStr1 := "5,1"
	nStr2 := "21,452"
	nStr3 := "8"
	nStr4 := "-6,7"

	expectedStr := "27.852"

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3, err := IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	ia4, err := IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.AddMany(&total, &ia1, &ia2, &ia3, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestIntAryMathAdd_AddMany_03(t *testing.T) {

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	totalStr := "0"

	total, err := IntAry{}.NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps("+
			"totalStr, expectedNumSeps). "+
			"totalStr='%v' expectedNumSeps='%v' Error='%v'",
			totalStr, expectedNumSeps.String(), err.Error())
	}

	nStr1 := "5.1"
	nStr2 := "21.452"
	nStr3 := "8"
	nStr4 := "-6.7"

	expectedStr := "27,852"

	alternateNumSeps := NumericSeparatorDto{}
	alternateNumSeps.DecimalSeparator = '.'
	alternateNumSeps.ThousandsSeparator = ','
	alternateNumSeps.CurrencySymbol = '$'

	ia1, err := IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr1, alternateNumSeps). "+
			"nStr1='%v' alternateNumSeps='%v' Error='%v' ",
			nStr1, alternateNumSeps.String(), err.Error())
	}

	ia2, err := IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr2, alternateNumSeps). "+
			"nStr2='%v' alternateNumSeps='%v' Error='%v' ",
			nStr2, alternateNumSeps.String(), err.Error())
	}

	ia3, err := IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr3, alternateNumSeps). "+
			"nStr3='%v' alternateNumSeps='%v' Error='%v' ",
			nStr3, alternateNumSeps.String(), err.Error())
	}

	ia4, err := IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStrWithNumSeps(nStr4, alternateNumSeps). "+
			"nStr4='%v' alternateNumSeps='%v' Error='%v' ",
			nStr4, alternateNumSeps.String(), err.Error())
	}

	IntAryMathAdd{}.AddMany(&total, &ia1, &ia2, &ia3, &ia4)

	if expectedStr != total.GetNumStr() {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedStr, total.GetNumStr())
	}

	actualNumSeps := total.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'.  Instead, NumSeps='%v' ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}
