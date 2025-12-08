package mathops

import (
	"testing"
)

func TestNumStrDto_CompareAbsoluteVals_01(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareAbsoluteVals_01"

	originalNumberStr1 := "-12567.218956"

	originalNumberStr2 := "-9211.40"

	expectedCompareInt := 1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareAbsoluteVals_02(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareAbsoluteVals_02"

	originalNumberStr1 := "-12567.218956"

	originalNumberStr2 := "9211.40"

	expectedCompareInt := 1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareAbsoluteVals_03(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareAbsoluteVals_03"

	originalNumberStr1 := "-12567.218956"

	originalNumberStr2 := "12567.218956"

	expectedCompareInt := 0

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareAbsoluteVals_04(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareAbsoluteVals_04"

	originalNumberStr1 := "567.21"

	originalNumberStr2 := "12567.218956"

	expectedCompareInt := -1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareAbsoluteVals_05(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareAbsoluteVals_05"

	originalNumberStr1 := "567.21"

	originalNumberStr2 := "-12567.218956"

	expectedCompareInt := -1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareAbsoluteVals_06(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareAbsoluteVals_06"

	originalNumberStr1 := "567.21"

	originalNumberStr2 := "-567.21"

	expectedCompareInt := 0

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareAbsoluteVals_07(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareAbsoluteVals_07"

	originalNumberStr1 := "567.21"

	originalNumberStr2 := "567.21"

	expectedCompareInt := 0

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareAbsoluteVals_08(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareAbsoluteVals_08"

	originalNumberStr1 := "567.21"

	originalNumberStr2 := "1567.21"

	expectedCompareInt := -1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareAbsoluteValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareAbsoluteValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Absolute Comparision Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareSignedVals_01(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareSignedVals_01"

	originalNumberStr1 := "-12567.218956"

	originalNumberStr2 := "-9211.40"

	expectedCompareInt := -1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareSignedValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if actualCompareInt != expectedCompareInt {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareSignedVals_02(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareSignedVals_02"

	originalNumberStr1 := "12567.218956"

	originalNumberStr2 := "9211.40"

	expectedCompareInt := 1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareSignedValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if actualCompareInt != expectedCompareInt {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareSignedVals_03(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareSignedVals_03"

	originalNumberStr1 := "-12567.218956"

	originalNumberStr2 := "9211.40"

	expectedCompareInt := -1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareSignedValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if actualCompareInt != expectedCompareInt {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareSignedVals_04(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareSignedVals_04"

	originalNumberStr1 := "12567.218956"

	originalNumberStr2 := "-9211.40"

	expectedCompareInt := 1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareSignedValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if actualCompareInt != expectedCompareInt {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareSignedVals_05(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareSignedVals_05"

	originalNumberStr1 := "12567.218956"

	originalNumberStr2 := "-12567.218956"

	expectedCompareInt := 1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareSignedValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if actualCompareInt != expectedCompareInt {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareSignedVals_06(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareSignedVals_06"

	originalNumberStr1 := "-12567.218956"

	originalNumberStr2 := "12567.218956"

	expectedCompareInt := -1

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareSignedValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if actualCompareInt != expectedCompareInt {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareSignedVals_07(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareSignedVals_07"

	originalNumberStr1 := "-12567.218956"

	originalNumberStr2 := "-12567.218956"

	expectedCompareInt := 0

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareSignedValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if actualCompareInt != expectedCompareInt {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CompareSignedVals_08(t *testing.T) {

	ePrefix := "TestNumStrDto_CompareSignedVals_08"

	originalNumberStr1 := "12567.218956"

	originalNumberStr2 := "12567.218956"

	expectedCompareInt := 0

	nDto := new(NumStrDto).New()

	numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1, err := nDto.ParseNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDto1.IsValid("Validating numStrDto1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto1.IsValid('Validating numStrDto1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto1NumberStr, err := numStrDto1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto1NumberStr, err := numStrDto1.GetNumStr()\n"+
			"numStrDto1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDto1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDto1NumberStr\n"+
			"Expected numStrDto1NumberStr = '%v'\n"+
			"  Actual numStrDto1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDto1NumberStr)

		return
	}

	numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2, err := nDto.ParseNumStr(originalNumberStr2)\n"+
			"  originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = numStrDto2.IsValid("Validating numStrDto2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDto2.IsValid('Validating numStrDto2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDto2NumberStr, err := numStrDto2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDto2NumberStr, err := numStrDto2.GetNumStr()\n"+
			"numStrDto2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != numStrDto2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDto2 Number String is INVALID!\n"+
			"Because originalNumberStr2 != numStrDto2NumberStr\n"+
			"Expected numStrDto2NumberStr = '%v'\n"+
			"  Actual numStrDto2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, numStrDto2NumberStr)

		return
	}

	actualCompareInt, err := nDto.CompareSignedValues(&numStrDto1, &numStrDto2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualCompareInt, err := nDto.CompareSignedValues(\n"+
			"  &numStrDto1, &numStrDto2)\n"+
			"numStrDto1= '%v'\n"+
			"numStrDto2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDto1NumberStr,
			numStrDto2NumberStr,
			err.Error())

		return
	}

	if actualCompareInt != expectedCompareInt {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", originalNumberStr1, originalNumberStr2, expectedCompareInt, actualCompareInt)
	}

	if expectedCompareInt != actualCompareInt {
		t.Errorf("%v\n"+
			"Error: Expected and Expected Comparision Signed Values ARE NOT EQUAL!\n"+
			"Because expectedCompareInt != actualCompareInt\n"+
			"Expected actualCompareInt = '%v'\n"+
			"  Actual actualCompareInt = '%v'\n\n",
			ePrefix, expectedCompareInt, actualCompareInt)

		return
	}

	return
}

func TestNumStrDto_CopyIn_01(t *testing.T) {

	ePrefix := "TestNumStrDto_CopyIn_01"

	originalNumberStr1 := "123.456"

	expectedNumberStr := "123.456"

	expectedAbsoluteIntStr := "123"

	expectedAbsoluteFracStr := "456"

	expectedSignValue := 1

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedHasNumericDigits, expectedIsFractionalValue bool

	expectedHasNumericDigits = true

	expectedIsFractionalValue = true

	numStrDtoOriginal1, err := new(NumStrDto).ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err := new(NumStrDto).ParseNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoResult := new(NumStrDto).New()

	err = numStrDtoResult.CopyIn(numStrDtoOriginal1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoOriginal1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"   numStrDtoResult.GetAbsIntRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err := \n"+
			"   numStrDtoResult.GetAbsFracRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedAbsoluteIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsoluteIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedAbsoluteIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedAbsoluteFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsoluteFracStr != numStrDtoResultAbsFracStr\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedAbsoluteFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: HasNumericDigits Flag is INVALID!\n"+
			"Because expectedHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag is INVALID!\n"+
			"Because expectedIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	return
}

func TestNumStrDto_CopyIn_02(t *testing.T) {

	ePrefix := "TestNumStrDto_CopyIn_02"

	originalNumberStr1 := "-123.456"

	expectedNumberStr := "-123.456"

	expectedAbsoluteIntStr := "123"

	expectedAbsoluteFracStr := "456"

	expectedSignValue := -1

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedHasNumericDigits, expectedIsFractionalValue bool

	expectedHasNumericDigits = true

	expectedIsFractionalValue = true

	numStrDtoOriginal1, err := new(NumStrDto).ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err := new(NumStrDto).ParseNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoResult := new(NumStrDto).New()

	err = numStrDtoResult.CopyIn(numStrDtoOriginal1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoOriginal1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"   numStrDtoResult.GetAbsIntRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err := \n"+
			"   numStrDtoResult.GetAbsFracRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedAbsoluteIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsoluteIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedAbsoluteIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedAbsoluteFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsoluteFracStr != numStrDtoResultAbsFracStr\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedAbsoluteFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: HasNumericDigits Flag is INVALID!\n"+
			"Because expectedHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag is INVALID!\n"+
			"Because expectedIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	return
}

func TestNumStrDto_CopyOut_01(t *testing.T) {

	ePrefix := "TestNumStrDto_CopyOut_01"

	originalNumberStr1 := "123.456"

	expectedNumberStr := "123.456"

	expectedAbsoluteIntStr := "123"

	expectedAbsoluteFracStr := "456"

	expectedSignValue := 1

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedHasNumericDigits, expectedIsFractionalValue bool

	expectedHasNumericDigits = true

	expectedIsFractionalValue = true

	numStrDtoOriginal1, err := new(NumStrDto).ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err := new(NumStrDto).ParseNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoResult, err := numStrDtoOriginal1.CopyOut()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := numStrDtoOriginal1.CopyOut()\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoOriginal1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"   numStrDtoResult.GetAbsIntRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err := \n"+
			"   numStrDtoResult.GetAbsFracRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedAbsoluteIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsoluteIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedAbsoluteIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedAbsoluteFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsoluteFracStr != numStrDtoResultAbsFracStr\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedAbsoluteFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: HasNumericDigits Flag is INVALID!\n"+
			"Because expectedHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag is INVALID!\n"+
			"Because expectedIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	return
}

func TestNumStrDto_CopyOut_02(t *testing.T) {

	ePrefix := "TestNumStrDto_CopyOut_02"

	originalNumberStr1 := "-123.456"

	expectedNumberStr := "-123.456"

	expectedAbsoluteIntStr := "123"

	expectedAbsoluteFracStr := "456"

	expectedSignValue := -1

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedHasNumericDigits, expectedIsFractionalValue bool

	expectedHasNumericDigits = true

	expectedIsFractionalValue = true

	numStrDtoOriginal1, err := new(NumStrDto).ParseNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err := new(NumStrDto).ParseNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because originalNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoResult, err := numStrDtoOriginal1.CopyOut()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := numStrDtoOriginal1.CopyOut()\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoOriginal1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"   numStrDtoResult.GetAbsIntRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err := \n"+
			"   numStrDtoResult.GetAbsFracRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedAbsoluteIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsoluteIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedAbsoluteIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedAbsoluteFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsoluteFracStr != numStrDtoResultAbsFracStr\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedAbsoluteFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: HasNumericDigits Flag is INVALID!\n"+
			"Because expectedHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag is INVALID!\n"+
			"Because expectedIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	return
}

func TestNumStrDto_Divide_01(t *testing.T) {

	ePrefix := "TestNumStrDto_Divide_01"

	//	  dividend/divisor = quotient
	//	  nDto / n2Dto = quotient

	dividendNumStr := "12"

	divisorNumStr := "3"

	minimumPrecision := 1

	maximumPrecision := 5

	expectedNumberStr := "4.0"

	expectedSignValue := 1

	expectedPrecisionInt := 1

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	dividendNumStrDto, err := new(NumStrDto).NewNumStr(dividendNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendNumStrDto, err := \n"+
			"  new(NumStrDto).NewNumStr(dividendNumStr)\n"+
			"dividendNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, dividendNumStr, err.Error())
		return
	}

	err = dividendNumStrDto.IsValid("Validating initial dividendNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumStrDto.IsValid('Validating dividendNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendNumStrDtoNumberStr, err := dividendNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dividendNumStrDtoNumberStr, err := dividendNumStrDto.GetNumStr()\n"+
			"dividendNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if dividendNumStr != dividendNumStrDtoNumberStr {
		t.Errorf("%v\n"+
			"Error: dividendNumStrDto Number String is INVALID!\n"+
			"Because dividendNumStr != dividendNumStrDtoNumberStr\n"+
			"Expected dividendNumStrDtoNumberStr = '%v'\n"+
			"  Actual dividendNumStrDtoNumberStr = '%v'\n\n",
			ePrefix, dividendNumStr, dividendNumStrDtoNumberStr)

		return
	}

	divisorNumStrDto, err := new(NumStrDto).NewNumStr(divisorNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStrDto, err := new(NumStrDto).\n"+
			"  NewNumStr(divisorNumStr)\n"+
			"divisorNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, divisorNumStr, err.Error())
		return
	}

	err = divisorNumStrDto.IsValid("Validating divisorNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = divisorNumStrDto.IsValid('Validating divisorNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	divisorNumStrDtoNumberStr, err := divisorNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"divisorNumStrDtoNumberStr, err := divisorNumStrDto.GetNumStr()\n"+
			"divisorNumStrDto set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if dividendNumStr != divisorNumStrDtoNumberStr {
		t.Errorf("%v\n"+
			"Error: divisorNumStrDto Number String is INVALID!\n"+
			"Because dividendNumStr != divisorNumStrDtoNumberStr\n"+
			"Expected divisorNumStrDtoNumberStr = '%v'\n"+
			"  Actual divisorNumStrDtoNumberStr = '%v'\n\n",
			ePrefix, dividendNumStr, divisorNumStrDtoNumberStr)

		return
	}

	err = dividendNumStrDto.Divide(divisorNumStrDto, minimumPrecision, maximumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumStrDto.Divide(\n"+
			"  divisorNumStrDto, minimumPrecision, maximumPrecision)\n"+
			"dividendNumStrDto= '%v'\n"+
			"divisorNumStrDto= '%v'\n"+
			"minimumPrecision= '%v'\n"+
			"maximumPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			dividendNumStrDtoNumberStr,
			divisorNumStrDtoNumberStr,
			minimumPrecision,
			maximumPrecision,
			err.Error())

		return
	}

	err = dividendNumStrDto.IsValid("Validating final dividendNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = dividendNumStrDto.IsValid('Validating final dividendNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	dividendNumStrDtoNumberStr, err = dividendNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotientNumStrDto, err := dividendNumStrDto.CopyOut()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumStrDto, err := dividendNumStrDto.CopyOut()\n"+
			"dividendNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, dividendNumStrDtoNumberStr, err.Error())
		return
	}

	err = quotientNumStrDto.IsValid("Validating quotientNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = quotientNumStrDto.IsValid('Validating quotientNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotientNumStrDtoNumberStr, err := quotientNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumStrDtoNumberStr, err := quotientNumStrDto.GetNumStr()\n"+
			"quotientNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	quotientNumStrDtoPrecisionInt := quotientNumStrDto.GetPrecision()

	quotientNumStrDtoPrecisionUint, err := quotientNumStrDto.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumStrDtoPrecisionUint, err :=\n"+
			"  quotientNumStrDto.GetPrecisionUint()\n"+
			"quotientNumStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, quotientNumStrDtoNumberStr, err.Error())
		return
	}

	quotientNumStrDtoSignValue, err := quotientNumStrDto.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumStrDtoSignValue, err := quotientNumStrDto.GetSign()\n"+
			"quotientNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, quotientNumStrDtoNumberStr, err.Error())
		return
	}

	quotientNumStrDtoNumSeps, err := quotientNumStrDto.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"quotientNumStrDtoNumSeps, err := quotientNumStrDto.GetNumericSeparatorsDto()\n"+
			"quotientNumStrDto= '%v\n"+
			"Error= '%v'\n\n", ePrefix, quotientNumStrDtoNumberStr, err.Error())
		return
	}

	if expectedNumberStr != quotientNumStrDtoNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and quotientNumStrDto Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != quotientNumStrDtoNumberStr\n"+
			"Expected quotientNumStrDtoNumberStr = '%v'\n"+
			"  Actual quotientNumStrDtoNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, quotientNumStrDtoNumberStr)

		return
	}

	if expectedPrecisionInt != quotientNumStrDtoPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & quotientNumStrDto Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != quotientNumStrDtoPrecisionInt\n"+
			"Expected quotientNumStrDtoPrecisionInt = '%v'\n"+
			"  Actual quotientNumStrDtoPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, quotientNumStrDtoPrecisionInt)

		return
	}

	if expectedPrecisionUint != quotientNumStrDtoPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & quotientNumStrDto Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != quotientNumStrDtoPrecisionUint\n"+
			"Expected quotientNumStrDtoPrecisionUint = '%v'\n"+
			"  Actual quotientNumStrDtoPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, quotientNumStrDtoPrecisionUint)

		return
	}

	if expectedSignValue != quotientNumStrDtoSignValue {
		t.Errorf("%v\n"+
			"Error: expected & quotientNumStrDto Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != quotientNumStrDtoSignValue\n"+
			"Expected quotientNumStrDtoSignValue = '%v'\n"+
			"  Actual quotientNumStrDtoSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, quotientNumStrDtoSignValue)

		return
	}

	if !expectedNumSeps.Equal(quotientNumStrDtoNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != quotientNumStrDtoNumSeps \n"+
			"Expected quotientNumStrDtoNumSeps = '%v'\n"+
			"  Actual quotientNumStrDtoNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), quotientNumStrDtoNumSeps.String())

		return
	}

	return
}

func TestNumStrDto_FormatForMathOps_01(t *testing.T) {

	ePrefix := "TestNumStrDto_FormatForMathOps_01"

	inputNumberStr1 := "-12567.218956"

	inputNumberStr2 := "-9211.40"

	//                                     1         2         3
	//                          0.1234567890123456789012345678901234567
	expectedNumberStr1 := "-12567.218956"

	expectedNumber1PrecisionInt := 6

	expectedNumber1PrecisionUint := uint(expectedNumber1PrecisionInt)

	expectedNumber1SignValue := -1

	expectedNumSeps1 := new(NumericSeparatorDto).NewUSADefaults()

	//                                     1         2         3
	//                          0.1234567890123456789012345678901234567
	expectedNumberStr2 := "-09211.400000"

	expectedNumber2PrecisionInt := 6

	expectedNumber2PrecisionUint := uint(expectedNumber2PrecisionInt)

	expectedNumber2SignValue := -1

	expectedNumSeps2 := new(NumericSeparatorDto).NewUSADefaults()

	expectedCompareValue := 1

	var expectedIsOrderReversedFlag bool

	expectedIsOrderReversedFlag = false

	nDto := new(NumStrDto).New()

	numStrDtoExpected1, err := nDto.ParseNumStr(expectedNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr1)\n"+
			"expectedNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr1, err.Error())
		return
	}

	err = numStrDtoExpected1.IsValid("Validating numStrDtoExpected1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected1.IsValid('Validating numStrDtoExpected1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()\n"+
			"numStrDtoExpected1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr1 != numStrDtoExpected1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected1 Number String is INVALID!\n"+
			"Because expectedNumberStr1 != numStrDtoExpected1NumberStr\n"+
			"Expected numStrDtoExpected1NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoExpected1NumberStr)

		return
	}

	numStrDtoExpected2, err := nDto.ParseNumStr(expectedNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr2)\n"+
			"expectedNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr2, err.Error())
		return
	}

	err = numStrDtoExpected2.IsValid("Validating numStrDtoExpected2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected2.IsValid('Validating numStrDtoExpected2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()\n"+
			"numStrDtoExpected2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr2 != numStrDtoExpected2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected2 Number String is INVALID!\n"+
			"Because expectedNumberStr2 != numStrDtoExpected2NumberStr\n"+
			"Expected numStrDtoExpected2NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoExpected2NumberStr)

		return
	}

	numStrDtoOriginal1, err := nDto.ParseNumStr(inputNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr1)\n"+
			"inputNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because inputNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoOriginal2, err := nDto.ParseNumStr(inputNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr2)\n"+
			"inputNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr2, err.Error())
		return
	}

	err = numStrDtoOriginal2.IsValid("Validating numStrDtoOriginal2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal2.IsValid('Validating numStrDtoOriginal2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()\n"+
			"numStrDtoOriginal2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr2 != numStrDtoOriginal2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal2 Number String is INVALID!\n"+
			"Because inputNumberStr2 != numStrDtoOriginal2NumberStr\n"+
			"Expected numStrDtoOriginal2NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal2NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr2, numStrDtoOriginal2NumberStr)

		return
	}

	numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err := nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err :=\n"+
			"  nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"numStrDtoOriginal2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoOriginal1NumberStr,
			numStrDtoOriginal2NumberStr,
			err.Error())

		return
	}

	err = numStrDtoResult1.IsValid("Validating numStrDtoResult1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult1.IsValid('Validating numStrDtoResult1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()\n"+
			"numStrDtoResult1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1PrecisionInt := numStrDtoResult1.GetPrecision()

	numStrDtoResult1PrecisionUint, err := numStrDtoResult1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1PrecisionUint, err :=\n"+
			"  numStrDtoResult1.GetPrecisionUint()\n"+
			"numStrDtoResult1Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()\n"+
			"numStrDtoResult1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult2.IsValid("Validating numStrDtoResult2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult2.IsValid('Validating numStrDtoResult2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()\n"+
			"numStrDtoResult2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2PrecisionInt := numStrDtoResult2.GetPrecision()

	numStrDtoResult2PrecisionUint, err := numStrDtoResult2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2PrecisionUint, err :=\n"+
			"  numStrDtoResult2.GetPrecisionUint()\n"+
			"numStrDtoResult2Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()\n"+
			"numStrDtoResult2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	if expectedIsOrderReversedFlag != actualIsOrderReversed {
		t.Errorf("%v\n"+
			"Error: Is Order Reversed Flag IS INVALID!\n"+
			"Because expectedIsOrderReversedFlag != actualIsOrderReversed\n"+
			"Expected actualIsOrderReversed = '%v'\n"+
			"  Actual actualIsOrderReversed = '%v'\n\n",
			ePrefix, expectedIsOrderReversedFlag, actualIsOrderReversed)

		return
	}

	if expectedCompareValue != actualCompareResult {
		t.Errorf("%v\n"+
			"Error: Actual Comparison Value is INVALID!\n"+
			"Because expectedCompareValue != actualCompareResult\n"+
			"Expected actualCompareResult = '%v'\n"+
			"  Actual actualCompareResult = '%v'\n\n",
			ePrefix, expectedCompareValue, actualCompareResult)

		return
	}

	if expectedNumberStr1 != numStrDtoResult1NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult1 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr1 != numStrDtoResult1NumberStr\n"+
			"Expected numStrDtoResult1NumberStr = '%v'\n"+
			"  Actual numStrDtoResult1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoResult1NumberStr)

		return
	}

	if expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt\n"+
			"Expected numStrDtoResult1PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionInt, numStrDtoResult1PrecisionInt)

		return
	}

	if expectedNumber1PrecisionUint != numStrDtoResult1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult1 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult1PrecisionUint\n"+
			"Expected numStrDtoResult1PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionUint, numStrDtoResult1PrecisionUint)

		return
	}

	if expectedNumber1SignValue != numStrDtoResult1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1SignValue != numStrDtoResult1SignValue\n"+
			"Expected numStrDtoResult1SignValue = '%v'\n"+
			"  Actual numStrDtoResult1SignValue = '%v'\n\n",
			ePrefix, expectedNumber1SignValue, numStrDtoResult1SignValue)

		return
	}

	if !expectedNumSeps1.Equal(numStrDtoResult1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps1 != numStrDtoResult1NumSeps \n"+
			"Expected numStrDtoResult1NumSeps = '%v'\n"+
			"  Actual numStrDtoResult1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps1.String(), numStrDtoResult1NumSeps.String())

		return
	}

	if expectedNumberStr2 != numStrDtoResult2NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult2 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr2 != numStrDtoResult2NumberStr\n"+
			"Expected numStrDtoResult2NumberStr = '%v'\n"+
			"  Actual numStrDtoResult2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoResult2NumberStr)

		return
	}

	if expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt\n"+
			"Expected numStrDtoResult2PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionInt, numStrDtoResult2PrecisionInt)

		return
	}

	if expectedNumber2PrecisionUint != numStrDtoResult2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult2 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult2PrecisionUint\n"+
			"Expected numStrDtoResult2PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionUint, numStrDtoResult2PrecisionUint)

		return
	}

	if expectedNumber2SignValue != numStrDtoResult2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2SignValue != numStrDtoResult2SignValue\n"+
			"Expected numStrDtoResult2SignValue = '%v'\n"+
			"  Actual numStrDtoResult2SignValue = '%v'\n\n",
			ePrefix, expectedNumber2SignValue, numStrDtoResult2SignValue)

		return
	}

	if !expectedNumSeps2.Equal(numStrDtoResult2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps2 != numStrDtoResult2NumSeps \n"+
			"Expected numStrDtoResult2NumSeps = '%v'\n"+
			"  Actual numStrDtoResult2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps2.String(), numStrDtoResult2NumSeps.String())

		return
	}

	return
}

func TestNumStrDto_FormatForMathOps_02(t *testing.T) {

	ePrefix := "TestNumStrDto_FormatForMathOps_02"

	inputNumberStr1 := "-9211.40"

	inputNumberStr2 := "-12567.218956"

	//                                     1         2         3
	//                          0.1234567890123456789012345678901234567
	expectedNumberStr1 := "-12567.218956"

	expectedNumber1PrecisionInt := 6

	expectedNumber1PrecisionUint := uint(expectedNumber1PrecisionInt)

	expectedNumber1SignValue := -1

	expectedNumSeps1 := new(NumericSeparatorDto).NewUSADefaults()

	//                                     1         2         3
	//                          0.1234567890123456789012345678901234567
	expectedNumberStr2 := "-09211.400000"

	expectedNumber2PrecisionInt := 6

	expectedNumber2PrecisionUint := uint(expectedNumber2PrecisionInt)

	expectedNumber2SignValue := -1

	expectedNumSeps2 := new(NumericSeparatorDto).NewUSADefaults()

	expectedCompareValue := 1

	var expectedIsOrderReversedFlag bool

	expectedIsOrderReversedFlag = true

	nDto := new(NumStrDto).New()

	numStrDtoExpected1, err := nDto.ParseNumStr(expectedNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr1)\n"+
			"expectedNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr1, err.Error())
		return
	}

	err = numStrDtoExpected1.IsValid("Validating numStrDtoExpected1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected1.IsValid('Validating numStrDtoExpected1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()\n"+
			"numStrDtoExpected1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr1 != numStrDtoExpected1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected1 Number String is INVALID!\n"+
			"Because expectedNumberStr1 != numStrDtoExpected1NumberStr\n"+
			"Expected numStrDtoExpected1NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoExpected1NumberStr)

		return
	}

	numStrDtoExpected2, err := nDto.ParseNumStr(expectedNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr2)\n"+
			"expectedNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr2, err.Error())
		return
	}

	err = numStrDtoExpected2.IsValid("Validating numStrDtoExpected2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected2.IsValid('Validating numStrDtoExpected2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()\n"+
			"numStrDtoExpected2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr2 != numStrDtoExpected2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected2 Number String is INVALID!\n"+
			"Because expectedNumberStr2 != numStrDtoExpected2NumberStr\n"+
			"Expected numStrDtoExpected2NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoExpected2NumberStr)

		return
	}

	numStrDtoOriginal1, err := nDto.ParseNumStr(inputNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr1)\n"+
			"inputNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because inputNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoOriginal2, err := nDto.ParseNumStr(inputNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr2)\n"+
			"inputNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr2, err.Error())
		return
	}

	err = numStrDtoOriginal2.IsValid("Validating numStrDtoOriginal2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal2.IsValid('Validating numStrDtoOriginal2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()\n"+
			"numStrDtoOriginal2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr2 != numStrDtoOriginal2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal2 Number String is INVALID!\n"+
			"Because inputNumberStr2 != numStrDtoOriginal2NumberStr\n"+
			"Expected numStrDtoOriginal2NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal2NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr2, numStrDtoOriginal2NumberStr)

		return
	}

	numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err := nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err :=\n"+
			"  nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"numStrDtoOriginal2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoOriginal1NumberStr,
			numStrDtoOriginal2NumberStr,
			err.Error())

		return
	}

	err = numStrDtoResult1.IsValid("Validating numStrDtoResult1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult1.IsValid('Validating numStrDtoResult1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()\n"+
			"numStrDtoResult1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1PrecisionInt := numStrDtoResult1.GetPrecision()

	numStrDtoResult1PrecisionUint, err := numStrDtoResult1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1PrecisionUint, err :=\n"+
			"  numStrDtoResult1.GetPrecisionUint()\n"+
			"numStrDtoResult1Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()\n"+
			"numStrDtoResult1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult2.IsValid("Validating numStrDtoResult2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult2.IsValid('Validating numStrDtoResult2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()\n"+
			"numStrDtoResult2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2PrecisionInt := numStrDtoResult2.GetPrecision()

	numStrDtoResult2PrecisionUint, err := numStrDtoResult2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2PrecisionUint, err :=\n"+
			"  numStrDtoResult2.GetPrecisionUint()\n"+
			"numStrDtoResult2Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()\n"+
			"numStrDtoResult2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	if expectedIsOrderReversedFlag != actualIsOrderReversed {
		t.Errorf("%v\n"+
			"Error: Is Order Reversed Flag IS INVALID!\n"+
			"Because expectedIsOrderReversedFlag != actualIsOrderReversed\n"+
			"Expected actualIsOrderReversed = '%v'\n"+
			"  Actual actualIsOrderReversed = '%v'\n\n",
			ePrefix, expectedIsOrderReversedFlag, actualIsOrderReversed)

		return
	}

	if expectedCompareValue != actualCompareResult {
		t.Errorf("%v\n"+
			"Error: Actual Comparison Value is INVALID!\n"+
			"Because expectedCompareValue != actualCompareResult\n"+
			"Expected actualCompareResult = '%v'\n"+
			"  Actual actualCompareResult = '%v'\n\n",
			ePrefix, expectedCompareValue, actualCompareResult)

		return
	}

	if expectedNumberStr1 != numStrDtoResult1NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult1 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr1 != numStrDtoResult1NumberStr\n"+
			"Expected numStrDtoResult1NumberStr = '%v'\n"+
			"  Actual numStrDtoResult1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoResult1NumberStr)

		return
	}

	if expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt\n"+
			"Expected numStrDtoResult1PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionInt, numStrDtoResult1PrecisionInt)

		return
	}

	if expectedNumber1PrecisionUint != numStrDtoResult1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult1 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult1PrecisionUint\n"+
			"Expected numStrDtoResult1PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionUint, numStrDtoResult1PrecisionUint)

		return
	}

	if expectedNumber1SignValue != numStrDtoResult1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1SignValue != numStrDtoResult1SignValue\n"+
			"Expected numStrDtoResult1SignValue = '%v'\n"+
			"  Actual numStrDtoResult1SignValue = '%v'\n\n",
			ePrefix, expectedNumber1SignValue, numStrDtoResult1SignValue)

		return
	}

	if !expectedNumSeps1.Equal(numStrDtoResult1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps1 != numStrDtoResult1NumSeps \n"+
			"Expected numStrDtoResult1NumSeps = '%v'\n"+
			"  Actual numStrDtoResult1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps1.String(), numStrDtoResult1NumSeps.String())

		return
	}

	if expectedNumberStr2 != numStrDtoResult2NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult2 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr2 != numStrDtoResult2NumberStr\n"+
			"Expected numStrDtoResult2NumberStr = '%v'\n"+
			"  Actual numStrDtoResult2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoResult2NumberStr)

		return
	}

	if expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt\n"+
			"Expected numStrDtoResult2PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionInt, numStrDtoResult2PrecisionInt)

		return
	}

	if expectedNumber2PrecisionUint != numStrDtoResult2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult2 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult2PrecisionUint\n"+
			"Expected numStrDtoResult2PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionUint, numStrDtoResult2PrecisionUint)

		return
	}

	if expectedNumber2SignValue != numStrDtoResult2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2SignValue != numStrDtoResult2SignValue\n"+
			"Expected numStrDtoResult2SignValue = '%v'\n"+
			"  Actual numStrDtoResult2SignValue = '%v'\n\n",
			ePrefix, expectedNumber2SignValue, numStrDtoResult2SignValue)

		return
	}

	if !expectedNumSeps2.Equal(numStrDtoResult2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps2 != numStrDtoResult2NumSeps \n"+
			"Expected numStrDtoResult2NumSeps = '%v'\n"+
			"  Actual numStrDtoResult2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps2.String(), numStrDtoResult2NumSeps.String())

		return
	}

	return
}

func TestNumStrDto_FormatForMathOps_03(t *testing.T) {

	ePrefix := "TestNumStrDto_FormatForMathOps_03"

	inputNumberStr1 := "-6"

	inputNumberStr2 := "67.521"

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr1 := "67.521"

	expectedNumber1PrecisionInt := 3

	expectedNumber1PrecisionUint := uint(expectedNumber1PrecisionInt)

	expectedNumber1SignValue := 1

	expectedNumSeps1 := new(NumericSeparatorDto).NewUSADefaults()

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	expectedNumberStr2 := "-06.000"

	expectedNumber2PrecisionInt := 3

	expectedNumber2PrecisionUint := uint(expectedNumber2PrecisionInt)

	expectedNumber2SignValue := -1

	expectedNumSeps2 := new(NumericSeparatorDto).NewUSADefaults()

	expectedCompareValue := 1

	var expectedIsOrderReversedFlag bool

	expectedIsOrderReversedFlag = true

	nDto := new(NumStrDto).New()

	numStrDtoExpected1, err := nDto.ParseNumStr(expectedNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr1)\n"+
			"expectedNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr1, err.Error())
		return
	}

	err = numStrDtoExpected1.IsValid("Validating numStrDtoExpected1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected1.IsValid('Validating numStrDtoExpected1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()\n"+
			"numStrDtoExpected1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr1 != numStrDtoExpected1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected1 Number String is INVALID!\n"+
			"Because expectedNumberStr1 != numStrDtoExpected1NumberStr\n"+
			"Expected numStrDtoExpected1NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoExpected1NumberStr)

		return
	}

	numStrDtoExpected2, err := nDto.ParseNumStr(expectedNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr2)\n"+
			"expectedNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr2, err.Error())
		return
	}

	err = numStrDtoExpected2.IsValid("Validating numStrDtoExpected2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected2.IsValid('Validating numStrDtoExpected2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()\n"+
			"numStrDtoExpected2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr2 != numStrDtoExpected2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected2 Number String is INVALID!\n"+
			"Because expectedNumberStr2 != numStrDtoExpected2NumberStr\n"+
			"Expected numStrDtoExpected2NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoExpected2NumberStr)

		return
	}

	numStrDtoOriginal1, err := nDto.ParseNumStr(inputNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr1)\n"+
			"inputNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because inputNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoOriginal2, err := nDto.ParseNumStr(inputNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr2)\n"+
			"inputNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr2, err.Error())
		return
	}

	err = numStrDtoOriginal2.IsValid("Validating numStrDtoOriginal2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal2.IsValid('Validating numStrDtoOriginal2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()\n"+
			"numStrDtoOriginal2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr2 != numStrDtoOriginal2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal2 Number String is INVALID!\n"+
			"Because inputNumberStr2 != numStrDtoOriginal2NumberStr\n"+
			"Expected numStrDtoOriginal2NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal2NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr2, numStrDtoOriginal2NumberStr)

		return
	}

	numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err := nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err :=\n"+
			"  nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"numStrDtoOriginal2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoOriginal1NumberStr,
			numStrDtoOriginal2NumberStr,
			err.Error())

		return
	}

	err = numStrDtoResult1.IsValid("Validating numStrDtoResult1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult1.IsValid('Validating numStrDtoResult1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()\n"+
			"numStrDtoResult1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1PrecisionInt := numStrDtoResult1.GetPrecision()

	numStrDtoResult1PrecisionUint, err := numStrDtoResult1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1PrecisionUint, err :=\n"+
			"  numStrDtoResult1.GetPrecisionUint()\n"+
			"numStrDtoResult1Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()\n"+
			"numStrDtoResult1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult2.IsValid("Validating numStrDtoResult2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult2.IsValid('Validating numStrDtoResult2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()\n"+
			"numStrDtoResult2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2PrecisionInt := numStrDtoResult2.GetPrecision()

	numStrDtoResult2PrecisionUint, err := numStrDtoResult2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2PrecisionUint, err :=\n"+
			"  numStrDtoResult2.GetPrecisionUint()\n"+
			"numStrDtoResult2Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()\n"+
			"numStrDtoResult2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	if expectedIsOrderReversedFlag != actualIsOrderReversed {
		t.Errorf("%v\n"+
			"Error: Is Order Reversed Flag IS INVALID!\n"+
			"Because expectedIsOrderReversedFlag != actualIsOrderReversed\n"+
			"Expected actualIsOrderReversed = '%v'\n"+
			"  Actual actualIsOrderReversed = '%v'\n\n",
			ePrefix, expectedIsOrderReversedFlag, actualIsOrderReversed)

		return
	}

	if expectedCompareValue != actualCompareResult {
		t.Errorf("%v\n"+
			"Error: Actual Comparison Value is INVALID!\n"+
			"Because expectedCompareValue != actualCompareResult\n"+
			"Expected actualCompareResult = '%v'\n"+
			"  Actual actualCompareResult = '%v'\n\n",
			ePrefix, expectedCompareValue, actualCompareResult)

		return
	}

	if expectedNumberStr1 != numStrDtoResult1NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult1 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr1 != numStrDtoResult1NumberStr\n"+
			"Expected numStrDtoResult1NumberStr = '%v'\n"+
			"  Actual numStrDtoResult1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoResult1NumberStr)

		return
	}

	if expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt\n"+
			"Expected numStrDtoResult1PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionInt, numStrDtoResult1PrecisionInt)

		return
	}

	if expectedNumber1PrecisionUint != numStrDtoResult1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult1 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult1PrecisionUint\n"+
			"Expected numStrDtoResult1PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionUint, numStrDtoResult1PrecisionUint)

		return
	}

	if expectedNumber1SignValue != numStrDtoResult1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1SignValue != numStrDtoResult1SignValue\n"+
			"Expected numStrDtoResult1SignValue = '%v'\n"+
			"  Actual numStrDtoResult1SignValue = '%v'\n\n",
			ePrefix, expectedNumber1SignValue, numStrDtoResult1SignValue)

		return
	}

	if !expectedNumSeps1.Equal(numStrDtoResult1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps1 != numStrDtoResult1NumSeps \n"+
			"Expected numStrDtoResult1NumSeps = '%v'\n"+
			"  Actual numStrDtoResult1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps1.String(), numStrDtoResult1NumSeps.String())

		return
	}

	if expectedNumberStr2 != numStrDtoResult2NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult2 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr2 != numStrDtoResult2NumberStr\n"+
			"Expected numStrDtoResult2NumberStr = '%v'\n"+
			"  Actual numStrDtoResult2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoResult2NumberStr)

		return
	}

	if expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt\n"+
			"Expected numStrDtoResult2PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionInt, numStrDtoResult2PrecisionInt)

		return
	}

	if expectedNumber2PrecisionUint != numStrDtoResult2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult2 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult2PrecisionUint\n"+
			"Expected numStrDtoResult2PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionUint, numStrDtoResult2PrecisionUint)

		return
	}

	if expectedNumber2SignValue != numStrDtoResult2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2SignValue != numStrDtoResult2SignValue\n"+
			"Expected numStrDtoResult2SignValue = '%v'\n"+
			"  Actual numStrDtoResult2SignValue = '%v'\n\n",
			ePrefix, expectedNumber2SignValue, numStrDtoResult2SignValue)

		return
	}

	if !expectedNumSeps2.Equal(numStrDtoResult2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps2 != numStrDtoResult2NumSeps \n"+
			"Expected numStrDtoResult2NumSeps = '%v'\n"+
			"  Actual numStrDtoResult2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps2.String(), numStrDtoResult2NumSeps.String())

		return
	}

	return
}

func TestNumStrDto_FormatForMathOps_04(t *testing.T) {

	ePrefix := "TestNumStrDto_FormatForMathOps_04"

	inputNumberStr1 := "67.521"

	inputNumberStr2 := "-6"

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr1 := "67.521"

	expectedNumber1PrecisionInt := 3

	expectedNumber1PrecisionUint := uint(expectedNumber1PrecisionInt)

	expectedNumber1SignValue := 1

	expectedNumSeps1 := new(NumericSeparatorDto).NewUSADefaults()

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	expectedNumberStr2 := "-06.000"

	expectedNumber2PrecisionInt := 3

	expectedNumber2PrecisionUint := uint(expectedNumber2PrecisionInt)

	expectedNumber2SignValue := -1

	expectedNumSeps2 := new(NumericSeparatorDto).NewUSADefaults()

	expectedCompareValue := 1

	var expectedIsOrderReversedFlag bool

	expectedIsOrderReversedFlag = false

	nDto := new(NumStrDto).New()

	numStrDtoExpected1, err := nDto.ParseNumStr(expectedNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr1)\n"+
			"expectedNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr1, err.Error())
		return
	}

	err = numStrDtoExpected1.IsValid("Validating numStrDtoExpected1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected1.IsValid('Validating numStrDtoExpected1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()\n"+
			"numStrDtoExpected1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr1 != numStrDtoExpected1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected1 Number String is INVALID!\n"+
			"Because expectedNumberStr1 != numStrDtoExpected1NumberStr\n"+
			"Expected numStrDtoExpected1NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoExpected1NumberStr)

		return
	}

	numStrDtoExpected2, err := nDto.ParseNumStr(expectedNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr2)\n"+
			"expectedNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr2, err.Error())
		return
	}

	err = numStrDtoExpected2.IsValid("Validating numStrDtoExpected2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected2.IsValid('Validating numStrDtoExpected2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()\n"+
			"numStrDtoExpected2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr2 != numStrDtoExpected2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected2 Number String is INVALID!\n"+
			"Because expectedNumberStr2 != numStrDtoExpected2NumberStr\n"+
			"Expected numStrDtoExpected2NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoExpected2NumberStr)

		return
	}

	numStrDtoOriginal1, err := nDto.ParseNumStr(inputNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr1)\n"+
			"inputNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because inputNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoOriginal2, err := nDto.ParseNumStr(inputNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr2)\n"+
			"inputNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr2, err.Error())
		return
	}

	err = numStrDtoOriginal2.IsValid("Validating numStrDtoOriginal2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal2.IsValid('Validating numStrDtoOriginal2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()\n"+
			"numStrDtoOriginal2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr2 != numStrDtoOriginal2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal2 Number String is INVALID!\n"+
			"Because inputNumberStr2 != numStrDtoOriginal2NumberStr\n"+
			"Expected numStrDtoOriginal2NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal2NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr2, numStrDtoOriginal2NumberStr)

		return
	}

	numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err := nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err :=\n"+
			"  nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"numStrDtoOriginal2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoOriginal1NumberStr,
			numStrDtoOriginal2NumberStr,
			err.Error())

		return
	}

	err = numStrDtoResult1.IsValid("Validating numStrDtoResult1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult1.IsValid('Validating numStrDtoResult1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()\n"+
			"numStrDtoResult1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1PrecisionInt := numStrDtoResult1.GetPrecision()

	numStrDtoResult1PrecisionUint, err := numStrDtoResult1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1PrecisionUint, err :=\n"+
			"  numStrDtoResult1.GetPrecisionUint()\n"+
			"numStrDtoResult1Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()\n"+
			"numStrDtoResult1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult2.IsValid("Validating numStrDtoResult2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult2.IsValid('Validating numStrDtoResult2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()\n"+
			"numStrDtoResult2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2PrecisionInt := numStrDtoResult2.GetPrecision()

	numStrDtoResult2PrecisionUint, err := numStrDtoResult2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2PrecisionUint, err :=\n"+
			"  numStrDtoResult2.GetPrecisionUint()\n"+
			"numStrDtoResult2Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()\n"+
			"numStrDtoResult2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	if expectedIsOrderReversedFlag != actualIsOrderReversed {
		t.Errorf("%v\n"+
			"Error: Is Order Reversed Flag IS INVALID!\n"+
			"Because expectedIsOrderReversedFlag != actualIsOrderReversed\n"+
			"Expected actualIsOrderReversed = '%v'\n"+
			"  Actual actualIsOrderReversed = '%v'\n\n",
			ePrefix, expectedIsOrderReversedFlag, actualIsOrderReversed)

		return
	}

	if expectedCompareValue != actualCompareResult {
		t.Errorf("%v\n"+
			"Error: Actual Comparison Value is INVALID!\n"+
			"Because expectedCompareValue != actualCompareResult\n"+
			"Expected actualCompareResult = '%v'\n"+
			"  Actual actualCompareResult = '%v'\n\n",
			ePrefix, expectedCompareValue, actualCompareResult)

		return
	}

	if expectedNumberStr1 != numStrDtoResult1NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult1 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr1 != numStrDtoResult1NumberStr\n"+
			"Expected numStrDtoResult1NumberStr = '%v'\n"+
			"  Actual numStrDtoResult1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoResult1NumberStr)

		return
	}

	if expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt\n"+
			"Expected numStrDtoResult1PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionInt, numStrDtoResult1PrecisionInt)

		return
	}

	if expectedNumber1PrecisionUint != numStrDtoResult1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult1 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult1PrecisionUint\n"+
			"Expected numStrDtoResult1PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionUint, numStrDtoResult1PrecisionUint)

		return
	}

	if expectedNumber1SignValue != numStrDtoResult1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1SignValue != numStrDtoResult1SignValue\n"+
			"Expected numStrDtoResult1SignValue = '%v'\n"+
			"  Actual numStrDtoResult1SignValue = '%v'\n\n",
			ePrefix, expectedNumber1SignValue, numStrDtoResult1SignValue)

		return
	}

	if !expectedNumSeps1.Equal(numStrDtoResult1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps1 != numStrDtoResult1NumSeps \n"+
			"Expected numStrDtoResult1NumSeps = '%v'\n"+
			"  Actual numStrDtoResult1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps1.String(), numStrDtoResult1NumSeps.String())

		return
	}

	if expectedNumberStr2 != numStrDtoResult2NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult2 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr2 != numStrDtoResult2NumberStr\n"+
			"Expected numStrDtoResult2NumberStr = '%v'\n"+
			"  Actual numStrDtoResult2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoResult2NumberStr)

		return
	}

	if expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt\n"+
			"Expected numStrDtoResult2PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionInt, numStrDtoResult2PrecisionInt)

		return
	}

	if expectedNumber2PrecisionUint != numStrDtoResult2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult2 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult2PrecisionUint\n"+
			"Expected numStrDtoResult2PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionUint, numStrDtoResult2PrecisionUint)

		return
	}

	if expectedNumber2SignValue != numStrDtoResult2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2SignValue != numStrDtoResult2SignValue\n"+
			"Expected numStrDtoResult2SignValue = '%v'\n"+
			"  Actual numStrDtoResult2SignValue = '%v'\n\n",
			ePrefix, expectedNumber2SignValue, numStrDtoResult2SignValue)

		return
	}

	if !expectedNumSeps2.Equal(numStrDtoResult2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps2 != numStrDtoResult2NumSeps \n"+
			"Expected numStrDtoResult2NumSeps = '%v'\n"+
			"  Actual numStrDtoResult2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps2.String(), numStrDtoResult2NumSeps.String())

		return
	}

	return
}

func TestNumStrDto_FormatForMathOps_05(t *testing.T) {

	ePrefix := "TestNumStrDto_FormatForMathOps_05"

	inputNumberStr1 := "-67.521"

	inputNumberStr2 := "6"

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	expectedNumberStr1 := "-67.521"

	expectedNumber1PrecisionInt := 3

	expectedNumber1PrecisionUint := uint(expectedNumber1PrecisionInt)

	expectedNumber1SignValue := -1

	expectedNumSeps1 := new(NumericSeparatorDto).NewUSADefaults()

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr2 := "06.000"

	expectedNumber2PrecisionInt := 3

	expectedNumber2PrecisionUint := uint(expectedNumber2PrecisionInt)

	expectedNumber2SignValue := 1

	expectedNumSeps2 := new(NumericSeparatorDto).NewUSADefaults()

	expectedCompareValue := 1

	var expectedIsOrderReversedFlag bool

	expectedIsOrderReversedFlag = false

	nDto := new(NumStrDto).New()

	numStrDtoExpected1, err := nDto.ParseNumStr(expectedNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr1)\n"+
			"expectedNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr1, err.Error())
		return
	}

	err = numStrDtoExpected1.IsValid("Validating numStrDtoExpected1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected1.IsValid('Validating numStrDtoExpected1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()\n"+
			"numStrDtoExpected1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr1 != numStrDtoExpected1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected1 Number String is INVALID!\n"+
			"Because expectedNumberStr1 != numStrDtoExpected1NumberStr\n"+
			"Expected numStrDtoExpected1NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoExpected1NumberStr)

		return
	}

	numStrDtoExpected2, err := nDto.ParseNumStr(expectedNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr2)\n"+
			"expectedNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr2, err.Error())
		return
	}

	err = numStrDtoExpected2.IsValid("Validating numStrDtoExpected2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected2.IsValid('Validating numStrDtoExpected2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()\n"+
			"numStrDtoExpected2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr2 != numStrDtoExpected2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected2 Number String is INVALID!\n"+
			"Because expectedNumberStr2 != numStrDtoExpected2NumberStr\n"+
			"Expected numStrDtoExpected2NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoExpected2NumberStr)

		return
	}

	numStrDtoOriginal1, err := nDto.ParseNumStr(inputNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr1)\n"+
			"inputNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because inputNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoOriginal2, err := nDto.ParseNumStr(inputNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr2)\n"+
			"inputNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr2, err.Error())
		return
	}

	err = numStrDtoOriginal2.IsValid("Validating numStrDtoOriginal2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal2.IsValid('Validating numStrDtoOriginal2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()\n"+
			"numStrDtoOriginal2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr2 != numStrDtoOriginal2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal2 Number String is INVALID!\n"+
			"Because inputNumberStr2 != numStrDtoOriginal2NumberStr\n"+
			"Expected numStrDtoOriginal2NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal2NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr2, numStrDtoOriginal2NumberStr)

		return
	}

	numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err := nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err :=\n"+
			"  nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"numStrDtoOriginal2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoOriginal1NumberStr,
			numStrDtoOriginal2NumberStr,
			err.Error())

		return
	}

	err = numStrDtoResult1.IsValid("Validating numStrDtoResult1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult1.IsValid('Validating numStrDtoResult1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()\n"+
			"numStrDtoResult1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1PrecisionInt := numStrDtoResult1.GetPrecision()

	numStrDtoResult1PrecisionUint, err := numStrDtoResult1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1PrecisionUint, err :=\n"+
			"  numStrDtoResult1.GetPrecisionUint()\n"+
			"numStrDtoResult1Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()\n"+
			"numStrDtoResult1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult2.IsValid("Validating numStrDtoResult2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult2.IsValid('Validating numStrDtoResult2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()\n"+
			"numStrDtoResult2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2PrecisionInt := numStrDtoResult2.GetPrecision()

	numStrDtoResult2PrecisionUint, err := numStrDtoResult2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2PrecisionUint, err :=\n"+
			"  numStrDtoResult2.GetPrecisionUint()\n"+
			"numStrDtoResult2Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()\n"+
			"numStrDtoResult2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	if expectedIsOrderReversedFlag != actualIsOrderReversed {
		t.Errorf("%v\n"+
			"Error: Is Order Reversed Flag IS INVALID!\n"+
			"Because expectedIsOrderReversedFlag != actualIsOrderReversed\n"+
			"Expected actualIsOrderReversed = '%v'\n"+
			"  Actual actualIsOrderReversed = '%v'\n\n",
			ePrefix, expectedIsOrderReversedFlag, actualIsOrderReversed)

		return
	}

	if expectedCompareValue != actualCompareResult {
		t.Errorf("%v\n"+
			"Error: Actual Comparison Value is INVALID!\n"+
			"Because expectedCompareValue != actualCompareResult\n"+
			"Expected actualCompareResult = '%v'\n"+
			"  Actual actualCompareResult = '%v'\n\n",
			ePrefix, expectedCompareValue, actualCompareResult)

		return
	}

	if expectedNumberStr1 != numStrDtoResult1NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult1 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr1 != numStrDtoResult1NumberStr\n"+
			"Expected numStrDtoResult1NumberStr = '%v'\n"+
			"  Actual numStrDtoResult1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoResult1NumberStr)

		return
	}

	if expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt\n"+
			"Expected numStrDtoResult1PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionInt, numStrDtoResult1PrecisionInt)

		return
	}

	if expectedNumber1PrecisionUint != numStrDtoResult1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult1 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult1PrecisionUint\n"+
			"Expected numStrDtoResult1PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionUint, numStrDtoResult1PrecisionUint)

		return
	}

	if expectedNumber1SignValue != numStrDtoResult1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1SignValue != numStrDtoResult1SignValue\n"+
			"Expected numStrDtoResult1SignValue = '%v'\n"+
			"  Actual numStrDtoResult1SignValue = '%v'\n\n",
			ePrefix, expectedNumber1SignValue, numStrDtoResult1SignValue)

		return
	}

	if !expectedNumSeps1.Equal(numStrDtoResult1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps1 != numStrDtoResult1NumSeps \n"+
			"Expected numStrDtoResult1NumSeps = '%v'\n"+
			"  Actual numStrDtoResult1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps1.String(), numStrDtoResult1NumSeps.String())

		return
	}

	if expectedNumberStr2 != numStrDtoResult2NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult2 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr2 != numStrDtoResult2NumberStr\n"+
			"Expected numStrDtoResult2NumberStr = '%v'\n"+
			"  Actual numStrDtoResult2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoResult2NumberStr)

		return
	}

	if expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt\n"+
			"Expected numStrDtoResult2PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionInt, numStrDtoResult2PrecisionInt)

		return
	}

	if expectedNumber2PrecisionUint != numStrDtoResult2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult2 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult2PrecisionUint\n"+
			"Expected numStrDtoResult2PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionUint, numStrDtoResult2PrecisionUint)

		return
	}

	if expectedNumber2SignValue != numStrDtoResult2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2SignValue != numStrDtoResult2SignValue\n"+
			"Expected numStrDtoResult2SignValue = '%v'\n"+
			"  Actual numStrDtoResult2SignValue = '%v'\n\n",
			ePrefix, expectedNumber2SignValue, numStrDtoResult2SignValue)

		return
	}

	if !expectedNumSeps2.Equal(numStrDtoResult2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps2 != numStrDtoResult2NumSeps \n"+
			"Expected numStrDtoResult2NumSeps = '%v'\n"+
			"  Actual numStrDtoResult2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps2.String(), numStrDtoResult2NumSeps.String())

		return
	}

	return
}

func TestNumStrDto_FormatForMathOps_06(t *testing.T) {

	ePrefix := "TestNumStrDto_FormatForMathOps_06"

	inputNumberStr1 := "-67.521"

	inputNumberStr2 := "67.521"

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	expectedNumberStr1 := "-67.521"

	expectedNumber1PrecisionInt := 3

	expectedNumber1PrecisionUint := uint(expectedNumber1PrecisionInt)

	expectedNumber1SignValue := -1

	expectedNumSeps1 := new(NumericSeparatorDto).NewUSADefaults()

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr2 := "67.521"

	expectedNumber2PrecisionInt := 3

	expectedNumber2PrecisionUint := uint(expectedNumber2PrecisionInt)

	expectedNumber2SignValue := 1

	expectedNumSeps2 := new(NumericSeparatorDto).NewUSADefaults()

	expectedCompareValue := 0

	var expectedIsOrderReversedFlag bool

	expectedIsOrderReversedFlag = false

	nDto := new(NumStrDto).New()

	numStrDtoExpected1, err := nDto.ParseNumStr(expectedNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr1)\n"+
			"expectedNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr1, err.Error())
		return
	}

	err = numStrDtoExpected1.IsValid("Validating numStrDtoExpected1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected1.IsValid('Validating numStrDtoExpected1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected1NumberStr, err := numStrDtoExpected1.GetNumStr()\n"+
			"numStrDtoExpected1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr1 != numStrDtoExpected1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected1 Number String is INVALID!\n"+
			"Because expectedNumberStr1 != numStrDtoExpected1NumberStr\n"+
			"Expected numStrDtoExpected1NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoExpected1NumberStr)

		return
	}

	numStrDtoExpected2, err := nDto.ParseNumStr(expectedNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2, err :=\n"+
			"  nDto.ParseNumStr(expectedNumberStr2)\n"+
			"expectedNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr2, err.Error())
		return
	}

	err = numStrDtoExpected2.IsValid("Validating numStrDtoExpected2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoExpected2.IsValid('Validating numStrDtoExpected2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoExpected2NumberStr, err := numStrDtoExpected2.GetNumStr()\n"+
			"numStrDtoExpected2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr2 != numStrDtoExpected2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoExpected2 Number String is INVALID!\n"+
			"Because expectedNumberStr2 != numStrDtoExpected2NumberStr\n"+
			"Expected numStrDtoExpected2NumberStr = '%v'\n"+
			"  Actual numStrDtoExpected2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoExpected2NumberStr)

		return
	}

	numStrDtoOriginal1, err := nDto.ParseNumStr(inputNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr1)\n"+
			"inputNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr1, err.Error())
		return
	}

	err = numStrDtoOriginal1.IsValid("Validating numStrDtoOriginal1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal1.IsValid('Validating numStrDtoOriginal1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal1NumberStr, err := numStrDtoOriginal1.GetNumStr()\n"+
			"numStrDtoOriginal1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr1 != numStrDtoOriginal1NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal1 Number String is INVALID!\n"+
			"Because inputNumberStr1 != numStrDtoOriginal1NumberStr\n"+
			"Expected numStrDtoOriginal1NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal1NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr1, numStrDtoOriginal1NumberStr)

		return
	}

	numStrDtoOriginal2, err := nDto.ParseNumStr(inputNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2, err :=\n"+
			"  nDto.ParseNumStr(inputNumberStr2)\n"+
			"inputNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr2, err.Error())
		return
	}

	err = numStrDtoOriginal2.IsValid("Validating numStrDtoOriginal2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoOriginal2.IsValid('Validating numStrDtoOriginal2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoOriginal2NumberStr, err := numStrDtoOriginal2.GetNumStr()\n"+
			"numStrDtoOriginal2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if inputNumberStr2 != numStrDtoOriginal2NumberStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoOriginal2 Number String is INVALID!\n"+
			"Because inputNumberStr2 != numStrDtoOriginal2NumberStr\n"+
			"Expected numStrDtoOriginal2NumberStr = '%v'\n"+
			"  Actual numStrDtoOriginal2NumberStr = '%v'\n\n",
			ePrefix, inputNumberStr2, numStrDtoOriginal2NumberStr)

		return
	}

	numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err := nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1, numStrDtoResult2, actualCompareResult, actualIsOrderReversed, err :=\n"+
			"  nDto.FormatForMathOps(numStrDtoOriginal1, numStrDtoOriginal2)\n"+
			"numStrDtoOriginal1= '%v'\n"+
			"numStrDtoOriginal2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoOriginal1NumberStr,
			numStrDtoOriginal2NumberStr,
			err.Error())

		return
	}

	err = numStrDtoResult1.IsValid("Validating numStrDtoResult1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult1.IsValid('Validating numStrDtoResult1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumberStr, err := numStrDtoResult1.GetNumStr()\n"+
			"numStrDtoResult1 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult1PrecisionInt := numStrDtoResult1.GetPrecision()

	numStrDtoResult1PrecisionUint, err := numStrDtoResult1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1PrecisionUint, err :=\n"+
			"  numStrDtoResult1.GetPrecisionUint()\n"+
			"numStrDtoResult1Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1SignValue, err := numStrDtoResult1.GetSign()\n"+
			"numStrDtoResult1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult1NumSeps, err := numStrDtoResult1.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult1= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult1NumberStr, err.Error())
		return
	}

	err = numStrDtoResult2.IsValid("Validating numStrDtoResult2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult2.IsValid('Validating numStrDtoResult2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumberStr, err := numStrDtoResult2.GetNumStr()\n"+
			"numStrDtoResult2 set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult2PrecisionInt := numStrDtoResult2.GetPrecision()

	numStrDtoResult2PrecisionUint, err := numStrDtoResult2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2PrecisionUint, err :=\n"+
			"  numStrDtoResult2.GetPrecisionUint()\n"+
			"numStrDtoResult2Result= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2SignValue, err := numStrDtoResult2.GetSign()\n"+
			"numStrDtoResult2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult2NumSeps, err := numStrDtoResult2.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResult2NumberStr, err.Error())
		return
	}

	if expectedIsOrderReversedFlag != actualIsOrderReversed {
		t.Errorf("%v\n"+
			"Error: Is Order Reversed Flag IS INVALID!\n"+
			"Because expectedIsOrderReversedFlag != actualIsOrderReversed\n"+
			"Expected actualIsOrderReversed = '%v'\n"+
			"  Actual actualIsOrderReversed = '%v'\n\n",
			ePrefix, expectedIsOrderReversedFlag, actualIsOrderReversed)

		return
	}

	if expectedCompareValue != actualCompareResult {
		t.Errorf("%v\n"+
			"Error: Actual Comparison Value is INVALID!\n"+
			"Because expectedCompareValue != actualCompareResult\n"+
			"Expected actualCompareResult = '%v'\n"+
			"  Actual actualCompareResult = '%v'\n\n",
			ePrefix, expectedCompareValue, actualCompareResult)

		return
	}

	if expectedNumberStr1 != numStrDtoResult1NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult1 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr1 != numStrDtoResult1NumberStr\n"+
			"Expected numStrDtoResult1NumberStr = '%v'\n"+
			"  Actual numStrDtoResult1NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr1, numStrDtoResult1NumberStr)

		return
	}

	if expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1PrecisionInt != numStrDtoResult1PrecisionInt\n"+
			"Expected numStrDtoResult1PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionInt, numStrDtoResult1PrecisionInt)

		return
	}

	if expectedNumber1PrecisionUint != numStrDtoResult1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult1 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult1PrecisionUint\n"+
			"Expected numStrDtoResult1PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult1PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber1PrecisionUint, numStrDtoResult1PrecisionUint)

		return
	}

	if expectedNumber1SignValue != numStrDtoResult1SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult1 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber1SignValue != numStrDtoResult1SignValue\n"+
			"Expected numStrDtoResult1SignValue = '%v'\n"+
			"  Actual numStrDtoResult1SignValue = '%v'\n\n",
			ePrefix, expectedNumber1SignValue, numStrDtoResult1SignValue)

		return
	}

	if !expectedNumSeps1.Equal(numStrDtoResult1NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps1 != numStrDtoResult1NumSeps \n"+
			"Expected numStrDtoResult1NumSeps = '%v'\n"+
			"  Actual numStrDtoResult1NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps1.String(), numStrDtoResult1NumSeps.String())

		return
	}

	if expectedNumberStr2 != numStrDtoResult2NumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult2 Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr2 != numStrDtoResult2NumberStr\n"+
			"Expected numStrDtoResult2NumberStr = '%v'\n"+
			"  Actual numStrDtoResult2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr2, numStrDtoResult2NumberStr)

		return
	}

	if expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2PrecisionInt != numStrDtoResult2PrecisionInt\n"+
			"Expected numStrDtoResult2PrecisionInt = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionInt = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionInt, numStrDtoResult2PrecisionInt)

		return
	}

	if expectedNumber2PrecisionUint != numStrDtoResult2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult2 Precision Uint Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResult2PrecisionUint\n"+
			"Expected numStrDtoResult2PrecisionUint = '%v'\n"+
			"  Actual numStrDtoResult2PrecisionUint = '%v'\n\n",
			ePrefix, expectedNumber2PrecisionUint, numStrDtoResult2PrecisionUint)

		return
	}

	if expectedNumber2SignValue != numStrDtoResult2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedNumber2SignValue != numStrDtoResult2SignValue\n"+
			"Expected numStrDtoResult2SignValue = '%v'\n"+
			"  Actual numStrDtoResult2SignValue = '%v'\n\n",
			ePrefix, expectedNumber2SignValue, numStrDtoResult2SignValue)

		return
	}

	if !expectedNumSeps2.Equal(numStrDtoResult2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps2 != numStrDtoResult2NumSeps \n"+
			"Expected numStrDtoResult2NumSeps = '%v'\n"+
			"  Actual numStrDtoResult2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps2.String(), numStrDtoResult2NumSeps.String())

		return
	}

	return
}
