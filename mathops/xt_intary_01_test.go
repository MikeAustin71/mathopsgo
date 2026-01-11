package mathops

import (
	"testing"
)

func TestIntAry_AddMultipleToThis_01(t *testing.T) {

	ePrefix := "TestIntAry_AddMultipleToThis_01"

	nStr1 := "457.3"
	nStr2 := "82.975"
	nStr3 := "94"
	nStr4 := "697.21589"
	nStr5 := "9648.37"

	//                                   1         2         3
	//                        0.1234567890123456789012345678901234567
	expectedNumberStr := "10979.86089"

	expectedPrecisionUint := uint(5)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia0, err := new(IntAry).NewZero(0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia0, err := new(IntAry).NewZero(0)\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1 := new(IntAry).New()

	err = ia1.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2 := new(IntAry).New()

	err = ia2.SetIntAryWithNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.SetIntAryWithNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia3 := new(IntAry).New()

	err = ia3.SetIntAryWithNumStr(nStr3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" err = ia3.SetIntAryWithNumStr(nStr3)\n"+
			"nStr3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr3, err.Error())
		return
	}

	err = ia3.IsValid("Validating ia3")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia3.IsValid('Validating ia3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia3NumberStr, err := ia3.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia3NumberStr, err := ia3.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia4 := new(IntAry).New()

	err = ia4.SetIntAryWithNumStr(nStr4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia4.SetIntAryWithNumStr(nStr4)\n"+
			"nStr4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr4, err.Error())
		return
	}

	err = ia4.IsValid("Validating ia4")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia4.IsValid('Validating ia4')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia4NumberStr, err := ia4.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia4NumberStr, err := ia4.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia5 := new(IntAry).New()

	err = ia5.SetIntAryWithNumStr(nStr5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia5.SetIntAryWithNumStr(nStr5)\n"+
			"nStr5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr5, err.Error())
		return
	}

	err = ia5.IsValid("Validating ia5")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia5.IsValid('Validating ia5')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia5NumberStr, err := ia5.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia5NumberStr, err := ia5.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia0.AddMultipleToThis(&ia1, &ia2, &ia3, &ia4, &ia5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia0.AddMultipleToThis(&ia1, &ia2, &ia3, &ia4, &ia5)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"ia3= '%v'\n"+
			"ia4= '%v'\n"+
			"ia5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			ia3NumberStr,
			ia4NumberStr,
			ia5NumberStr,
			err.Error())

		return
	}

	err = ia0.IsValid("Validating ia0")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia0.IsValid('Validating ia0')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia0NumberStr, err := ia0.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia0NumberStr, err := ia0.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia0PrecisionUint, err := ia0.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia0PrecisionUint, err :=\n"+
			"  ia0.GetPrecisionUint()\n"+
			"ia0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia0NumberStr, err.Error())
		return
	}

	ia0SignValue, err := ia0.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia0SignValue, err := ia0.GetSign()\n"+
			"ia0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia0NumberStr, err.Error())
		return
	}

	ia0NumSeps, err := ia0.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia0NumSeps, err := ia0.GetNumericSeparatorsDto()\n"+
			"ia0= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia0NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia0NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia0NumberStr \n"+
			"Expected ia0NumberStr = '%v'\n"+
			"  Actual ia0NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia0NumberStr)

		return
	}

	if expectedPrecisionUint != ia0PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia0PrecisionUint\n"+
			"Expected ia0PrecisionUint = '%v'\n"+
			"  Actual ia0PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia0PrecisionUint)

		return
	}

	if expectedSignVal != ia0SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia0 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia0SignValue\n"+
			"Expected ia0SignValue = '%v'\n"+
			"  Actual ia0SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia0SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia0NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia0NumSeps \n"+
			"Expected ia0NumSeps = '%v'\n"+
			"  Actual ia0NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia0NumSeps.String())

		return
	}

	return
}

func TestIntAry_AddToThis_01(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_01"

	nStr1 := "457.3"

	nStr2 := "22.2"

	expectedNumberStr := "479.5"

	expectedFinalIntAry := []uint8{4, 7, 9, 5}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(1)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1 := new(IntAry).New()

	err := ia1.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2 := new(IntAry).New()

	err = ia2.SetIntAryWithNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.SetIntAryWithNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_02(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_02"

	nStr1 := "457.325"

	nStr2 := "-22.2"

	expectedNumberStr := "435.125"

	expectedFinalIntAry := []uint8{4, 3, 5, 1, 2, 5}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_03(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_03"

	nStr1 := "-457.325"

	nStr2 := "22.2"

	expectedNumberStr := "-435.125"

	expectedFinalIntAry := []uint8{4, 3, 5, 1, 2, 5}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(3)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_04(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_04"

	nStr1 := "-457.325"

	nStr2 := "-22.2"

	expectedNumberStr := "-479.525"

	expectedFinalIntAry := []uint8{4, 7, 9, 5, 2, 5}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(3)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_05(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_05"

	nStr1 := "0.000"

	nStr2 := "-22.2"

	expectedNumberStr := "-22.200"

	expectedFinalIntAry := []uint8{2, 2, 2, 0, 0}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(3)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_06(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_06"

	nStr1 := "0.000"

	nStr2 := "0"

	expectedNumberStr := "0.000"

	expectedFinalIntAry := []uint8{0, 0, 0, 0}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_07(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_07"

	nStr1 := "99.225"

	nStr2 := "-99.1"

	expectedNumberStr := "0.125"

	expectedFinalIntAry := []uint8{0, 1, 2, 5}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_08(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_08"

	nStr1 := "350"

	nStr2 := "122"

	expectedNumberStr := "472"

	expectedFinalIntAry := []uint8{0, 1, 2, 5}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_09(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_09"

	nStr1 := "-350"

	nStr2 := "122"

	expectedNumberStr := "-228"

	expectedFinalIntAry := []uint8{2, 2, 8}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_10(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_10"

	nStr1 := "-350"

	nStr2 := "-122"

	expectedNumberStr := "-472"

	expectedFinalIntAry := []uint8{4, 7, 2}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_11(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_11"

	nStr1 := "350"

	nStr2 := "-122"

	expectedNumberStr := "228"

	expectedFinalIntAry := []uint8{2, 2, 8}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_12(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_12"

	nStr1 := "350"

	nStr2 := "0"

	expectedNumberStr := "350"

	expectedFinalIntAry := []uint8{3, 5, 0}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_13(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_13"

	nStr1 := "-350"

	nStr2 := "0"

	expectedNumberStr := "-350"

	expectedFinalIntAry := []uint8{3, 5, 0}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_14(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_14"

	nStr1 := "122"

	nStr2 := "350"

	expectedNumberStr := "472"

	expectedFinalIntAry := []uint8{4, 7, 2}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_15(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_15"

	nStr1 := "-122"

	nStr2 := "350"

	expectedNumberStr := "228"

	expectedFinalIntAry := []uint8{2, 2, 8}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_16(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_16"

	nStr1 := "-122"

	nStr2 := "-350"

	expectedNumberStr := "-472"

	expectedFinalIntAry := []uint8{4, 7, 2}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_17(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_17"

	nStr1 := "122"

	nStr2 := "-350"

	expectedNumberStr := "-228"

	expectedFinalIntAry := []uint8{2, 2, 8}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_18(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_18"

	nStr1 := "0"

	nStr2 := "350"

	expectedNumberStr := "350"

	expectedFinalIntAry := []uint8{3, 5, 0}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_19(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_19"

	nStr1 := "0"

	nStr2 := "-350"

	expectedNumberStr := "-350"

	expectedFinalIntAry := []uint8{3, 5, 0}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_20(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_20"

	nStr1 := "122"

	nStr2 := "122"

	expectedNumberStr := "244"

	expectedFinalIntAry := []uint8{2, 4, 4}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_21(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_21"

	nStr1 := "-122"

	nStr2 := "122"

	expectedNumberStr := "0"

	expectedFinalIntAry := []uint8{0}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_22(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_22"

	nStr1 := "-122"

	nStr2 := "-122"

	expectedNumberStr := "-244"

	expectedFinalIntAry := []uint8{2, 4, 4}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_23(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_23"

	nStr1 := "122"

	nStr2 := "-122"

	expectedNumberStr := "0"

	expectedFinalIntAry := []uint8{0}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}

func TestIntAry_AddToThis_24(t *testing.T) {

	ePrefix := "TestIntAry_AddToThis_24"

	nStr1 := "0"

	nStr2 := "0"

	expectedNumberStr := "0"

	expectedFinalIntAry := []uint8{0}

	lenExpectedIntAry := len(expectedFinalIntAry)

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(nStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(nStr1)\n"+
			"nStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr1, err.Error())
		return
	}

	err = ia1.IsValid("Validating initial ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating initial ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err := ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err := ia1.GetNumStr()\n"+
			"ia1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(nStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(nStr2)\n"+
			"nStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nStr2, err.Error())
		return
	}

	err = ia2.IsValid("Validating ia2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia2.IsValid('Validating ia2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia2NumberStr, err := ia2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2NumberStr, err := ia2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = ia1.AddIntAryToThis(&ia2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntAryToThis(&ia2)\n"+
			"ia1= '%v'\n"+
			"ia2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			ia2NumberStr,
			err.Error())

		return
	}

	err = ia1.IsValid("Validating final ia1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.IsValid('Validating final ia1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1NumberStr, err = ia1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumberStr, err = ia1.GetNumStr()\n"+
			"ia1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	ia1PrecisionUint, err := ia1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1PrecisionUint, err :=\n"+
			"  ia1.GetPrecisionUint()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1SignValue, err := ia1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1SignValue, err := ia1.GetSign()\n"+
			"ia1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, ia1NumberStr, err.Error())
		return
	}

	ia1NumSeps, err := ia1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1NumSeps, err := ia1.GetNumericSeparatorsDto()\n"+
			"ia1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, ia1NumberStr, err.Error())
		return
	}

	if expectedNumberStr != ia1NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != ia1NumberStr \n"+
			"Expected ia1NumberStr = '%v'\n"+
			"  Actual ia1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, ia1NumberStr)

		return
	}

	if expectedPrecisionUint != ia1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != ia1PrecisionUint\n"+
			"Expected ia1PrecisionUint = '%v'\n"+
			"  Actual ia1PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, ia1PrecisionUint)

		return
	}

	if expectedSignVal != ia1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & ia1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != ia1SignValue\n"+
			"Expected ia1SignValue = '%v'\n"+
			"  Actual ia1SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, ia1SignValue)

		return
	}

	if !expectedNumSeps.Equal(ia1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != ia1NumSeps \n"+
			"Expected ia1NumSeps = '%v'\n"+
			"  Actual ia1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), ia1NumSeps.String())

		return
	}

	ia1Stats := ia1.GetIntAryStats()

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lenExpectedIntAry, ia1Stats.IntAryLen)
	}

	if lenExpectedIntAry != ia1Stats.IntAryLen {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because lenExpectedIntAry != ia1Stats.IntAryLen\n"+
			"Expected ia1Stats.IntAryLen = '%v'\n"+
			"  Actual ia1Stats.IntAryLen = '%v'\n\n",
			ePrefix, lenExpectedIntAry, ia1Stats.IntAryLen)

		return
	}

	var element uint8

	for i := 0; i < lenExpectedIntAry; i++ {

		element, err = ia1.GetIntAryElement(i)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"element, err = ia1.GetIntAryElement(i)\n"+
				"i= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

		if expectedFinalIntAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}

		if expectedFinalIntAry[i] != element {
			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because expectedFinalIntAry[%v] != element\n"+
				"Expected element = '%v'\n"+
				"  Actual element = '%v'\n\n",
				ePrefix, i, expectedFinalIntAry[i], element)

			return
		}

	}

	return
}
