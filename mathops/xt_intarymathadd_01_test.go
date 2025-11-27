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

	ePrefix := "TestIntAryMathAdd_RunTotal_01"

	originalNumberStr1 := "5.1"

	originalNumberStr2 := "21.452"

	originalNumberStr3 := "8"

	originalNumberStr4 := "-6.7"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "27.852"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryTotal, err := new(IntAry).NewZero(0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(IntAry).NewZero(0)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryTotal.IsValid("Validating initial intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating initial intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err := intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

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

	err = intAry1.IsValid("Validating intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
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

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
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
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr2 != intAry2NumberStr\n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAry2NumberStr)

		return
	}

	intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)\n"+
			"originalNumberStr3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr3, err.Error())
		return
	}

	err = intAry3.IsValid("Validating intAry3")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry3.IsValid('Validating intAry3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry3NumberStr, err := intAry3.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3NumberStr, err := intAry3.GetNumStr()\n"+
			"intAry3 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr3 != intAry3NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr3 != intAry3NumberStr\n"+
			"Expected intAry3NumberStr = '%v'\n"+
			"  Actual intAry3NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr3, intAry3NumberStr)

		return
	}

	intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)\n"+
			"originalNumberStr4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr4, err.Error())
		return
	}

	err = intAry4.IsValid("Validating intAry4")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry4.IsValid('Validating intAry4')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry4NumberStr, err := intAry4.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4NumberStr, err := intAry4.GetNumStr()\n"+
			"intAry4 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr4 != intAry4NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr4 != intAry4NumberStr\n"+
			"Expected intAry4NumberStr = '%v'\n"+
			"  Actual intAry4NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr4, intAry4NumberStr)

		return
	}

	intAryMathAdd := new(IntAryMathAdd)

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry1)\n"+
			"intAryTotal= '%v'\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry1NumberStr,
			err.Error())

		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry1 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry2)\n"+
			"intAryTotal= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry2 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry3)\n"+
			"intAryTotal= '%v'\n"+
			"intAry3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry3 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry4)\n"+
			"intAryTotal= '%v'\n"+
			"intAry3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating final intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating final intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry4 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalPrecisionInt := intAryTotal.GetPrecision()

	intAryTotalPrecisionUint, err := intAryTotal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalPrecisionUint, err :=\n"+
			"  intAryTotal.GetPrecisionUint()\n"+
			"intAryTotalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalSignValue, err := intAryTotal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalSignValue, err := intAryTotal.GetSign()\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()\n"+
			"intAryTotal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryTotalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Total Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryTotalNumberStr \n"+
			"Expected intAryTotalNumberStr = '%v'\n"+
			"  Actual intAryTotalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryTotalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryTotalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryTotalPrecisionInt\n"+
			"Expected intAryTotalPrecisionInt = '%v'\n"+
			"  Actual intAryTotalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryTotalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryTotalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryTotalPrecisionUint\n"+
			"Expected intAryTotalPrecisionUint = '%v'\n"+
			"  Actual intAryTotalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryTotalPrecisionUint)

		return
	}

	if expectedSignValue != intAryTotalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryTotalSignValue\n"+
			"Expected intAryTotalSignValue = '%v'\n"+
			"  Actual intAryTotalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryTotalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryTotalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryTotalNumSeps \n"+
			"Expected intAryTotalNumSeps = '%v'\n"+
			"  Actual intAryTotalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryTotalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_RunTotal_02(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_RunTotal_02"

	originalNumberStr1 := "5,1"

	originalNumberStr2 := "21,452"

	originalNumberStr3 := "8"

	originalNumberStr4 := "-6,7"

	originalTotalNumberStr := "0"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "27.852"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(originalTotalNumberStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(originalTotalNumberStr, expectedNumSeps)\n"+
			"originalTotalNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalTotalNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating initial intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating initial intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err := intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr1, alternateNumSeps)\n"+
			"originalNumberStr1= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr1,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry1.IsValid("Validating intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr1 != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, intAry1NumberStr)

		return
	}

	intAry2, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr2, alternateNumSeps)\n"+
			"originalNumberStr2= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr2,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
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
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr2 != intAry2NumberStr\n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAry2NumberStr)

		return
	}

	intAry3, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr3, alternateNumSeps)\n"+
			"originalNumberStr3= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr3,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry3.IsValid("Validating intAry3")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry3.IsValid('Validating intAry3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry3NumberStr, err := intAry3.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3NumberStr, err := intAry3.GetNumStr()\n"+
			"intAry3 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr3 != intAry3NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr3 != intAry3NumberStr\n"+
			"Expected intAry3NumberStr = '%v'\n"+
			"  Actual intAry3NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr3, intAry3NumberStr)

		return
	}

	intAry4, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr4, alternateNumSeps)\n"+
			"originalNumberStr4= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr4,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry4.IsValid("Validating intAry4")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry4.IsValid('Validating intAry4')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry4NumberStr, err := intAry4.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4NumberStr, err := intAry4.GetNumStr()\n"+
			"intAry4 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr4 != intAry4NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr4 != intAry4NumberStr\n"+
			"Expected intAry4NumberStr = '%v'\n"+
			"  Actual intAry4NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr4, intAry4NumberStr)

		return
	}

	intAryMathAdd := new(IntAryMathAdd)

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry1)\n"+
			"intAryTotal= '%v'\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry1NumberStr,
			err.Error())

		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry1 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry2)\n"+
			"intAryTotal= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry2 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry3)\n"+
			"intAryTotal= '%v'\n"+
			"intAry3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry3 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry4)\n"+
			"intAryTotal= '%v'\n"+
			"intAry3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating final intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating final intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry4 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalPrecisionInt := intAryTotal.GetPrecision()

	intAryTotalPrecisionUint, err := intAryTotal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalPrecisionUint, err :=\n"+
			"  intAryTotal.GetPrecisionUint()\n"+
			"intAryTotalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalSignValue, err := intAryTotal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalSignValue, err := intAryTotal.GetSign()\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()\n"+
			"intAryTotal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryTotalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Total Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryTotalNumberStr \n"+
			"Expected intAryTotalNumberStr = '%v'\n"+
			"  Actual intAryTotalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryTotalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryTotalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryTotalPrecisionInt\n"+
			"Expected intAryTotalPrecisionInt = '%v'\n"+
			"  Actual intAryTotalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryTotalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryTotalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryTotalPrecisionUint\n"+
			"Expected intAryTotalPrecisionUint = '%v'\n"+
			"  Actual intAryTotalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryTotalPrecisionUint)

		return
	}

	if expectedSignValue != intAryTotalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryTotalSignValue\n"+
			"Expected intAryTotalSignValue = '%v'\n"+
			"  Actual intAryTotalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryTotalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryTotalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryTotalNumSeps \n"+
			"Expected intAryTotalNumSeps = '%v'\n"+
			"  Actual intAryTotalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryTotalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_RunTotal_03(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_RunTotal_03"

	originalNumberStr1 := "5.1"

	originalNumberStr2 := "21.452"

	originalNumberStr3 := "8"

	originalNumberStr4 := "-6.7"

	originalTotalNumberStr := "0"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "27,852"

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

	intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(originalTotalNumberStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(originalTotalNumberStr, expectedNumSeps)\n"+
			"originalTotalNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalTotalNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating initial intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating initial intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err := intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr1, alternateNumSeps)\n"+
			"originalNumberStr1= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr1,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry1.IsValid("Validating intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr1 != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, intAry1NumberStr)

		return
	}

	intAry2, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr2, alternateNumSeps)\n"+
			"originalNumberStr2= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr2,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
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
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr2 != intAry2NumberStr\n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAry2NumberStr)

		return
	}

	intAry3, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr3, alternateNumSeps)\n"+
			"originalNumberStr3= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr3,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry3.IsValid("Validating intAry3")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry3.IsValid('Validating intAry3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry3NumberStr, err := intAry3.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3NumberStr, err := intAry3.GetNumStr()\n"+
			"intAry3 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr3 != intAry3NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr3 != intAry3NumberStr\n"+
			"Expected intAry3NumberStr = '%v'\n"+
			"  Actual intAry3NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr3, intAry3NumberStr)

		return
	}

	intAry4, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr4, alternateNumSeps)\n"+
			"originalNumberStr4= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr4,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry4.IsValid("Validating intAry4")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry4.IsValid('Validating intAry4')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry4NumberStr, err := intAry4.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4NumberStr, err := intAry4.GetNumStr()\n"+
			"intAry4 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr4 != intAry4NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr4 != intAry4NumberStr\n"+
			"Expected intAry4NumberStr = '%v'\n"+
			"  Actual intAry4NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr4, intAry4NumberStr)

		return
	}

	intAryMathAdd := new(IntAryMathAdd)

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry1)\n"+
			"intAryTotal= '%v'\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry1NumberStr,
			err.Error())

		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry1 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry2)\n"+
			"intAryTotal= '%v'\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry2 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry3)\n"+
			"intAryTotal= '%v'\n"+
			"intAry3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry3 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryMathAdd.RunTotal(&intAryTotal, &intAry4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.RunTotal(&intAryTotal, &intAry4)\n"+
			"intAryTotal= '%v'\n"+
			"intAry3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry2NumberStr,
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating final intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating final intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry4 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalPrecisionInt := intAryTotal.GetPrecision()

	intAryTotalPrecisionUint, err := intAryTotal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalPrecisionUint, err :=\n"+
			"  intAryTotal.GetPrecisionUint()\n"+
			"intAryTotalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalSignValue, err := intAryTotal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalSignValue, err := intAryTotal.GetSign()\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()\n"+
			"intAryTotal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryTotalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Total Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryTotalNumberStr \n"+
			"Expected intAryTotalNumberStr = '%v'\n"+
			"  Actual intAryTotalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryTotalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryTotalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryTotalPrecisionInt\n"+
			"Expected intAryTotalPrecisionInt = '%v'\n"+
			"  Actual intAryTotalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryTotalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryTotalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryTotalPrecisionUint\n"+
			"Expected intAryTotalPrecisionUint = '%v'\n"+
			"  Actual intAryTotalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryTotalPrecisionUint)

		return
	}

	if expectedSignValue != intAryTotalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryTotalSignValue\n"+
			"Expected intAryTotalSignValue = '%v'\n"+
			"  Actual intAryTotalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryTotalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryTotalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryTotalNumSeps \n"+
			"Expected intAryTotalNumSeps = '%v'\n"+
			"  Actual intAryTotalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryTotalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_AddManyArray_01(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_AddManyArray_01"

	originalNumberStr1 := "5.1"

	originalNumberStr2 := "21.452"

	originalNumberStr3 := "8"

	originalNumberStr4 := "-6.7"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "27.852"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryTotal, err := new(IntAry).NewZero(0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(IntAry).NewZero(0)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryTotal.IsValid("Validating initial intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating initial intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err := intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaArray := make([]IntAry, 4)

	iaArray[0], err = new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[0], err = new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	iaArray[1], err = new(IntAry).NewNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[1], err = new(IntAry).NewNumStr(originalNumberStr2)\n"+
			"originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	iaArray[2], err = new(IntAry).NewNumStr(originalNumberStr3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[2], err = new(IntAry).NewNumStr(originalNumberStr3)\n"+
			"originalNumberStr3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr3, err.Error())
		return
	}

	iaArray[3], err = new(IntAry).NewNumStr(originalNumberStr4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[2], err = new(IntAry).NewNumStr(originalNumberStr3)\n"+
			"originalNumberStr3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr3, err.Error())
		return
	}

	err = new(IntAryMathAdd).AddManyArray(&intAryTotal, iaArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = new(IntAryMathAdd).AddManyArray(&intAryTotal, iaArray)\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating final intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating final intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry4 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalPrecisionInt := intAryTotal.GetPrecision()

	intAryTotalPrecisionUint, err := intAryTotal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalPrecisionUint, err :=\n"+
			"  intAryTotal.GetPrecisionUint()\n"+
			"intAryTotalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalSignValue, err := intAryTotal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalSignValue, err := intAryTotal.GetSign()\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()\n"+
			"intAryTotal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryTotalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Total Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryTotalNumberStr \n"+
			"Expected intAryTotalNumberStr = '%v'\n"+
			"  Actual intAryTotalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryTotalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryTotalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryTotalPrecisionInt\n"+
			"Expected intAryTotalPrecisionInt = '%v'\n"+
			"  Actual intAryTotalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryTotalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryTotalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryTotalPrecisionUint\n"+
			"Expected intAryTotalPrecisionUint = '%v'\n"+
			"  Actual intAryTotalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryTotalPrecisionUint)

		return
	}

	if expectedSignValue != intAryTotalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryTotalSignValue\n"+
			"Expected intAryTotalSignValue = '%v'\n"+
			"  Actual intAryTotalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryTotalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryTotalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryTotalNumSeps \n"+
			"Expected intAryTotalNumSeps = '%v'\n"+
			"  Actual intAryTotalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryTotalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_AddManyArray_02(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_AddManyArray_02"

	originalNumberStr1 := "5,1"

	originalNumberStr2 := "21,452"

	originalNumberStr3 := "8"

	originalNumberStr4 := "-6,7"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "27.852"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	totalStr := "0"

	intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(totalStr, expectedNumSeps)\n"+
			"totalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			totalStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating initial intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating initial intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err := intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaArray := make([]IntAry, 4)

	iaArray[0], err = new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[0], err = new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr1, alternateNumSeps)\n"+
			"originalNumberStr1= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr1,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	iaArray[1], err = new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[1], err = new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr2, alternateNumSeps)\n"+
			"originalNumberStr2= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr2,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	iaArray[2], err = new(IntAry).NewNumStrWithNumSeps(originalNumberStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[2], err = new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr3, alternateNumSeps)\n"+
			"originalNumberStr3= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr3,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	iaArray[3], err = new(IntAry).NewNumStrWithNumSeps(originalNumberStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[3], err = new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr4, alternateNumSeps)\n"+
			"originalNumberStr4= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr4,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = new(IntAryMathAdd).AddManyArray(&intAryTotal, iaArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = new(IntAryMathAdd).AddManyArray(&intAryTotal, iaArray)\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating final intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating final intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry4 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalPrecisionInt := intAryTotal.GetPrecision()

	intAryTotalPrecisionUint, err := intAryTotal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalPrecisionUint, err :=\n"+
			"  intAryTotal.GetPrecisionUint()\n"+
			"intAryTotalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalSignValue, err := intAryTotal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalSignValue, err := intAryTotal.GetSign()\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()\n"+
			"intAryTotal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryTotalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Total Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryTotalNumberStr \n"+
			"Expected intAryTotalNumberStr = '%v'\n"+
			"  Actual intAryTotalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryTotalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryTotalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryTotalPrecisionInt\n"+
			"Expected intAryTotalPrecisionInt = '%v'\n"+
			"  Actual intAryTotalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryTotalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryTotalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryTotalPrecisionUint\n"+
			"Expected intAryTotalPrecisionUint = '%v'\n"+
			"  Actual intAryTotalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryTotalPrecisionUint)

		return
	}

	if expectedSignValue != intAryTotalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryTotalSignValue\n"+
			"Expected intAryTotalSignValue = '%v'\n"+
			"  Actual intAryTotalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryTotalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryTotalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryTotalNumSeps \n"+
			"Expected intAryTotalNumSeps = '%v'\n"+
			"  Actual intAryTotalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryTotalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_AddManyArray_03(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_AddManyArray_03"

	originalNumberStr1 := "5.1"

	originalNumberStr2 := "21.452"

	originalNumberStr3 := "8"

	originalNumberStr4 := "-6.7"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "27,852"

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

	alternateNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	totalStr := "0"

	intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(totalStr, expectedNumSeps)\n"+
			"totalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			totalStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating initial intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating initial intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err := intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaArray := make([]IntAry, 4)

	iaArray[0], err = new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[0], err = new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr1, alternateNumSeps)\n"+
			"originalNumberStr1= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr1,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	iaArray[1], err = new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[1], err = new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr2, alternateNumSeps)\n"+
			"originalNumberStr2= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr2,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	iaArray[2], err = new(IntAry).NewNumStrWithNumSeps(originalNumberStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[2], err = new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr3, alternateNumSeps)\n"+
			"originalNumberStr3= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr3,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	iaArray[3], err = new(IntAry).NewNumStrWithNumSeps(originalNumberStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaArray[3], err = new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr4, alternateNumSeps)\n"+
			"originalNumberStr4= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr4,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = new(IntAryMathAdd).AddManyArray(&intAryTotal, iaArray)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = new(IntAryMathAdd).AddManyArray(&intAryTotal, iaArray)\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating final intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating final intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry4 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalPrecisionInt := intAryTotal.GetPrecision()

	intAryTotalPrecisionUint, err := intAryTotal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalPrecisionUint, err :=\n"+
			"  intAryTotal.GetPrecisionUint()\n"+
			"intAryTotalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalSignValue, err := intAryTotal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalSignValue, err := intAryTotal.GetSign()\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()\n"+
			"intAryTotal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryTotalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Total Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryTotalNumberStr \n"+
			"Expected intAryTotalNumberStr = '%v'\n"+
			"  Actual intAryTotalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryTotalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryTotalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryTotalPrecisionInt\n"+
			"Expected intAryTotalPrecisionInt = '%v'\n"+
			"  Actual intAryTotalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryTotalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryTotalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryTotalPrecisionUint\n"+
			"Expected intAryTotalPrecisionUint = '%v'\n"+
			"  Actual intAryTotalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryTotalPrecisionUint)

		return
	}

	if expectedSignValue != intAryTotalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryTotalSignValue\n"+
			"Expected intAryTotalSignValue = '%v'\n"+
			"  Actual intAryTotalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryTotalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryTotalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryTotalNumSeps \n"+
			"Expected intAryTotalNumSeps = '%v'\n"+
			"  Actual intAryTotalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryTotalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_AddMany_01(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_AddMany_01"

	originalNumberStr1 := "5.1"

	originalNumberStr2 := "21.452"

	originalNumberStr3 := "8"

	originalNumberStr4 := "-6.7"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "27.852"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryTotal, err := new(IntAry).NewZero(0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"total, err := new(IntAry).NewZero(0)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAryTotal.IsValid("Validating initial intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating initial intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err := intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

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

	err = intAry1.IsValid("Validating intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
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

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
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
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr2 != intAry2NumberStr\n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAry2NumberStr)

		return
	}

	intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3, err := new(IntAry).NewNumStr(originalNumberStr3)\n"+
			"originalNumberStr3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr3, err.Error())
		return
	}

	err = intAry3.IsValid("Validating intAry3")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry3.IsValid('Validating intAry3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry3NumberStr, err := intAry3.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3NumberStr, err := intAry3.GetNumStr()\n"+
			"intAry3 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr3 != intAry3NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr3 != intAry3NumberStr\n"+
			"Expected intAry3NumberStr = '%v'\n"+
			"  Actual intAry3NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr3, intAry3NumberStr)

		return
	}

	intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4, err := new(IntAry).NewNumStr(originalNumberStr4)\n"+
			"originalNumberStr4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr4, err.Error())
		return
	}

	err = intAry4.IsValid("Validating intAry4")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry4.IsValid('Validating intAry4')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry4NumberStr, err := intAry4.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4NumberStr, err := intAry4.GetNumStr()\n"+
			"intAry4 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr4 != intAry4NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr4 != intAry4NumberStr\n"+
			"Expected intAry4NumberStr = '%v'\n"+
			"  Actual intAry4NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr4, intAry4NumberStr)

		return
	}

	intAryMathAdd := new(IntAryMathAdd)

	err = intAryMathAdd.AddMany(&intAryTotal, &intAry1, &intAry2, &intAry3, &intAry4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.AddMany(\n"+
			"  intAryTotal, &intAry1, &intAry2, &intAry3, &intAry4)\n"+
			"intAryTotal= '%v'\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"intAry3= '%v'\n"+
			"intAry4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry1NumberStr,
			intAry2NumberStr,
			intAry3NumberStr,
			intAry4NumberStr,
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating final intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating final intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry4 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalPrecisionInt := intAryTotal.GetPrecision()

	intAryTotalPrecisionUint, err := intAryTotal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalPrecisionUint, err :=\n"+
			"  intAryTotal.GetPrecisionUint()\n"+
			"intAryTotalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalSignValue, err := intAryTotal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalSignValue, err := intAryTotal.GetSign()\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()\n"+
			"intAryTotal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryTotalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Total Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryTotalNumberStr \n"+
			"Expected intAryTotalNumberStr = '%v'\n"+
			"  Actual intAryTotalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryTotalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryTotalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryTotalPrecisionInt\n"+
			"Expected intAryTotalPrecisionInt = '%v'\n"+
			"  Actual intAryTotalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryTotalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryTotalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryTotalPrecisionUint\n"+
			"Expected intAryTotalPrecisionUint = '%v'\n"+
			"  Actual intAryTotalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryTotalPrecisionUint)

		return
	}

	if expectedSignValue != intAryTotalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryTotalSignValue\n"+
			"Expected intAryTotalSignValue = '%v'\n"+
			"  Actual intAryTotalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryTotalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryTotalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryTotalNumSeps \n"+
			"Expected intAryTotalNumSeps = '%v'\n"+
			"  Actual intAryTotalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryTotalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_AddMany_02(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_AddMany_02"

	originalNumberStr1 := "5,1"

	originalNumberStr2 := "21,452"

	originalNumberStr3 := "8"

	originalNumberStr4 := "-6,7"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "27.852"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	alternateNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	alternateNumSeps.DecimalSeparator = frenchDecSeparator
	alternateNumSeps.ThousandsSeparator = frenchThousandsSeparator
	alternateNumSeps.CurrencySymbol = frenchCurrencySymbol

	totalStr := "0"

	intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(totalStr, expectedNumSeps)\n"+
			"totalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			totalStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating initial intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating initial intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err := intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr1, alternateNumSeps)\n"+
			"originalNumberStr1= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr1,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry1.IsValid("Validating intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr1 != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, intAry1NumberStr)

		return
	}

	intAry2, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr2, alternateNumSeps)\n"+
			"originalNumberStr2= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr2,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
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
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr2 != intAry2NumberStr\n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAry2NumberStr)

		return
	}

	intAry3, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr3, alternateNumSeps)\n"+
			"originalNumberStr3= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr3,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry3.IsValid("Validating intAry3")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry3.IsValid('Validating intAry3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry3NumberStr, err := intAry3.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3NumberStr, err := intAry3.GetNumStr()\n"+
			"intAry3 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr3 != intAry3NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr3 != intAry3NumberStr\n"+
			"Expected intAry3NumberStr = '%v'\n"+
			"  Actual intAry3NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr3, intAry3NumberStr)

		return
	}

	intAry4, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr4, alternateNumSeps)\n"+
			"originalNumberStr4= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr4,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry4.IsValid("Validating intAry4")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry4.IsValid('Validating intAry4')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry4NumberStr, err := intAry4.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4NumberStr, err := intAry4.GetNumStr()\n"+
			"intAry4 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr4 != intAry4NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr4 != intAry4NumberStr\n"+
			"Expected intAry4NumberStr = '%v'\n"+
			"  Actual intAry4NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr4, intAry4NumberStr)

		return
	}

	intAryMathAdd := new(IntAryMathAdd)

	err = intAryMathAdd.AddMany(&intAryTotal, &intAry1, &intAry2, &intAry3, &intAry4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.AddMany(\n"+
			"  intAryTotal, &intAry1, &intAry2, &intAry3, &intAry4)\n"+
			"intAryTotal= '%v'\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"intAry3= '%v'\n"+
			"intAry4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry1NumberStr,
			intAry2NumberStr,
			intAry3NumberStr,
			intAry4NumberStr,
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating final intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating final intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry4 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalPrecisionInt := intAryTotal.GetPrecision()

	intAryTotalPrecisionUint, err := intAryTotal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalPrecisionUint, err :=\n"+
			"  intAryTotal.GetPrecisionUint()\n"+
			"intAryTotalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalSignValue, err := intAryTotal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalSignValue, err := intAryTotal.GetSign()\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()\n"+
			"intAryTotal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryTotalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Total Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryTotalNumberStr \n"+
			"Expected intAryTotalNumberStr = '%v'\n"+
			"  Actual intAryTotalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryTotalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryTotalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryTotalPrecisionInt\n"+
			"Expected intAryTotalPrecisionInt = '%v'\n"+
			"  Actual intAryTotalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryTotalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryTotalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryTotalPrecisionUint\n"+
			"Expected intAryTotalPrecisionUint = '%v'\n"+
			"  Actual intAryTotalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryTotalPrecisionUint)

		return
	}

	if expectedSignValue != intAryTotalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryTotalSignValue\n"+
			"Expected intAryTotalSignValue = '%v'\n"+
			"  Actual intAryTotalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryTotalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryTotalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryTotalNumSeps \n"+
			"Expected intAryTotalNumSeps = '%v'\n"+
			"  Actual intAryTotalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryTotalNumSeps.String())

		return
	}

	return
}

func TestIntAryMathAdd_AddMany_03(t *testing.T) {

	ePrefix := "TestIntAryMathAdd_AddMany_02"

	originalNumberStr1 := "5.1"

	originalNumberStr2 := "21.452"

	originalNumberStr3 := "8"

	originalNumberStr4 := "-6.7"

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "27,852"

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

	totalStr := "0"

	intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(totalStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotal, err := new(IntAry).NewNumStrWithNumSeps(totalStr, expectedNumSeps)\n"+
			"totalStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			totalStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating initial intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating initial intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err := intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr1, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr1, alternateNumSeps)\n"+
			"originalNumberStr1= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr1,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry1.IsValid("Validating intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr1 != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, intAry1NumberStr)

		return
	}

	intAry2, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr2, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr2, alternateNumSeps)\n"+
			"originalNumberStr2= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr2,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
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
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr2 != intAry2NumberStr\n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, intAry2NumberStr)

		return
	}

	intAry3, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr3, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr3, alternateNumSeps)\n"+
			"originalNumberStr3= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr3,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry3.IsValid("Validating intAry3")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry3.IsValid('Validating intAry3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry3NumberStr, err := intAry3.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry3NumberStr, err := intAry3.GetNumStr()\n"+
			"intAry3 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr3 != intAry3NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr3 != intAry3NumberStr\n"+
			"Expected intAry3NumberStr = '%v'\n"+
			"  Actual intAry3NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr3, intAry3NumberStr)

		return
	}

	intAry4, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr4, alternateNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4, err := new(IntAry).NewNumStrWithNumSeps(\n"+
			"  originalNumberStr4, alternateNumSeps)\n"+
			"originalNumberStr4= '%v'\n"+
			"alternateNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberStr4,
			alternateNumSeps.String(),
			err.Error())

		return
	}

	err = intAry4.IsValid("Validating intAry4")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry4.IsValid('Validating intAry4')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry4NumberStr, err := intAry4.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry4NumberStr, err := intAry4.GetNumStr()\n"+
			"intAry4 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr4 != intAry4NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr4 != intAry4NumberStr\n"+
			"Expected intAry4NumberStr = '%v'\n"+
			"  Actual intAry4NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr4, intAry4NumberStr)

		return
	}

	intAryMathAdd := new(IntAryMathAdd)

	err = intAryMathAdd.AddMany(&intAryTotal, &intAry1, &intAry2, &intAry3, &intAry4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryMathAdd.AddMany(\n"+
			"  intAryTotal, &intAry1, &intAry2, &intAry3, &intAry4)\n"+
			"intAryTotal= '%v'\n"+
			"intAry1= '%v'\n"+
			"intAry2= '%v'\n"+
			"intAry3= '%v'\n"+
			"intAry4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryTotalNumberStr,
			intAry1NumberStr,
			intAry2NumberStr,
			intAry3NumberStr,
			intAry4NumberStr,
			err.Error())

		return
	}

	err = intAryTotal.IsValid("Validating final intAryTotal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryTotal.IsValid('Validating final intAryTotal')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalNumberStr, err = intAryTotal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumberStr, err := intAryTotal.GetNumStr()\n"+
			"intAryTotal set after intAry4 addition.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryTotalPrecisionInt := intAryTotal.GetPrecision()

	intAryTotalPrecisionUint, err := intAryTotal.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalPrecisionUint, err :=\n"+
			"  intAryTotal.GetPrecisionUint()\n"+
			"intAryTotalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalSignValue, err := intAryTotal.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalSignValue, err := intAryTotal.GetSign()\n"+
			"intAryTotal= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryTotalNumSeps, err := intAryTotal.GetNumericSeparatorsDto()\n"+
			"intAryTotal= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryTotalNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryTotalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Total Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryTotalNumberStr \n"+
			"Expected intAryTotalNumberStr = '%v'\n"+
			"  Actual intAryTotalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryTotalNumberStr)

		return
	}

	if expectedPrecisionInt != intAryTotalPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryTotalPrecisionInt\n"+
			"Expected intAryTotalPrecisionInt = '%v'\n"+
			"  Actual intAryTotalPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryTotalPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryTotalPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryTotal Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryTotalPrecisionUint\n"+
			"Expected intAryTotalPrecisionUint = '%v'\n"+
			"  Actual intAryTotalPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryTotalPrecisionUint)

		return
	}

	if expectedSignValue != intAryTotalSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryTotal Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryTotalSignValue\n"+
			"Expected intAryTotalSignValue = '%v'\n"+
			"  Actual intAryTotalSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryTotalSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryTotalNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryTotalNumSeps \n"+
			"Expected intAryTotalNumSeps = '%v'\n"+
			"  Actual intAryTotalNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryTotalNumSeps.String())

		return
	}

	return
}
