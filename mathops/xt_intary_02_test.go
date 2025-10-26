package mathops

import (
	"math/big"
	"strconv"
	"testing"
)

func TestIntAry_AddIntToThis_01(t *testing.T) {

	ePrefix := "TestIntAry_AddIntToThis_01"

	originalNumberStr1 := "25"

	originalNumberInt2 := 50

	originalNumberPrecisionUint2 := uint(0)

	expectedNumberStr := "75"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddIntToThis(originalNumberInt2, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionUint2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddIntToThis_02(t *testing.T) {

	ePrefix := "TestIntAry_AddIntToThis_02"

	originalNumberStr1 := "100"

	originalNumberInt2 := -25

	originalNumberPrecisionUint2 := uint(0)

	expectedNumberStr := "75"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddIntToThis(originalNumberInt2, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionUint2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddIntToThis_03(t *testing.T) {

	ePrefix := "TestIntAry_AddIntToThis_03"

	originalNumberStr1 := "-100"

	originalNumberInt2 := -25

	originalNumberPrecisionUint2 := uint(0)

	expectedNumberStr := "-125"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddIntToThis(originalNumberInt2, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionUint2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddIntToThis_04(t *testing.T) {

	ePrefix := "TestIntAry_AddIntToThis_04"

	originalNumberStr1 := "25.75"

	originalNumberInt2 := 5050

	originalNumberPrecisionUint2 := uint(2)

	expectedNumberStr := "76.25"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddIntToThis(originalNumberInt2, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionUint2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddIntToThis_05(t *testing.T) {

	ePrefix := "TestIntAry_AddIntToThis_05"

	originalNumberStr1 := "100.925"

	originalNumberInt2 := -25967

	originalNumberPrecisionUint2 := uint(3)

	expectedNumberStr := "74.958"

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddIntToThis(originalNumberInt2, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionUint2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddIntToThis_06(t *testing.T) {

	ePrefix := "TestIntAry_AddIntToThis_06"

	originalNumberStr1 := "-100.35842"

	originalNumberInt2 := -1256984

	originalNumberPrecisionUint2 := uint(4)

	expectedNumberStr := "-226.05682"

	expectedPrecisionUint := uint(5)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddIntToThis(originalNumberInt2, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionUint2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddIntToThis_07(t *testing.T) {

	ePrefix := "TestIntAry_AddIntToThis_07"

	originalNumberStr1 := "0"

	originalNumberInt2 := 0

	originalNumberPrecisionUint2 := uint(0)

	expectedNumberStr := "0"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddIntToThis(originalNumberInt2, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionUint2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionUint2,
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

	return
}

// TODO - Try returning '0.000'
func TestIntAry_AddIntToThis_08(t *testing.T) {

	ePrefix := "TestIntAry_AddIntToThis_08"

	originalNumberStr1 := "0.00"

	originalNumberInt2 := 0

	originalNumberPrecisionUint2 := uint(0)

	expectedNumberStr := "0"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddIntToThis(originalNumberInt2, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionUint2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddInt64ToThis_01(t *testing.T) {

	ePrefix := "TestIntAry_AddInt64ToThis_01"

	originalNumberStr1 := "25"

	originalNumberInt642 := int64(50)

	originalNumberPrecisionUint2 := uint(0)

	expectedNumberStr := "75"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddInt64ToThis(originalNumberInt642, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddInt64ToThis(\n"+
			"  originalNumberInt642, originalNumberPrecisionUint2)\n"+
			"originalNumberInt642= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt642,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddInt64ToThis_02(t *testing.T) {

	ePrefix := "TestIntAry_AddInt64ToThis_02"

	originalNumberStr1 := "100"

	originalNumberInt642 := int64(-25)

	originalNumberPrecisionUint2 := uint(0)

	expectedNumberStr := "75"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddInt64ToThis(originalNumberInt642, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddInt64ToThis(\n"+
			"  originalNumberInt642, originalNumberPrecisionUint2)\n"+
			"originalNumberInt642= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt642,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddInt64ToThis_03(t *testing.T) {

	ePrefix := "TestIntAry_AddInt64ToThis_03"

	originalNumberStr1 := "-100"

	originalNumberInt642 := int64(-25)

	originalNumberPrecisionUint2 := uint(0)

	expectedNumberStr := "-125"

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddInt64ToThis(originalNumberInt642, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddInt64ToThis(\n"+
			"  originalNumberInt642, originalNumberPrecisionUint2)\n"+
			"originalNumberInt642= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt642,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddInt64ToThis_04(t *testing.T) {

	ePrefix := "TestIntAry_AddInt64ToThis_04"

	originalNumberStr1 := "25.75"

	originalNumberInt642 := int64(5050)

	originalNumberPrecisionUint2 := uint(2)

	expectedNumberStr := "76.25"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddInt64ToThis(originalNumberInt642, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddInt64ToThis(\n"+
			"  originalNumberInt642, originalNumberPrecisionUint2)\n"+
			"originalNumberInt642= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt642,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddInt64ToThis_05(t *testing.T) {

	ePrefix := "TestIntAry_AddInt64ToThis_05"

	originalNumberStr1 := "100.925"

	originalNumberInt642 := int64(-25967)

	originalNumberPrecisionUint2 := uint(3)

	expectedNumberStr := "74.958"

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddInt64ToThis(originalNumberInt642, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddInt64ToThis(\n"+
			"  originalNumberInt642, originalNumberPrecisionUint2)\n"+
			"originalNumberInt642= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt642,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddInt64ToThis_06(t *testing.T) {

	ePrefix := "TestIntAry_AddInt64ToThis_06"

	originalNumberStr1 := "-100.35842"

	originalNumberInt642 := int64(-1256984)

	originalNumberPrecisionUint2 := uint(4)

	expectedNumberStr := "-226.05682"

	expectedPrecisionUint := uint(5)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddInt64ToThis(originalNumberInt642, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddInt64ToThis(\n"+
			"  originalNumberInt642, originalNumberPrecisionUint2)\n"+
			"originalNumberInt642= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt642,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddInt64ToThis_07(t *testing.T) {

	ePrefix := "TestIntAry_AddInt64ToThis_07"

	originalNumberStr1 := "0"

	originalNumberInt642 := int64(0)

	originalNumberPrecisionUint2 := uint(0)

	expectedNumberStr := "0"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddInt64ToThis(originalNumberInt642, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddInt64ToThis(\n"+
			"  originalNumberInt642, originalNumberPrecisionUint2)\n"+
			"originalNumberInt642= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt642,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddInt64ToThis_08(t *testing.T) {

	ePrefix := "TestIntAry_AddInt64ToThis_08"

	originalNumberStr1 := "0.00"

	originalNumberInt642 := int64(0)

	originalNumberPrecisionUint2 := uint(2)

	expectedNumberStr := "0.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddInt64ToThis(originalNumberInt642, originalNumberPrecisionUint2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddInt64ToThis(\n"+
			"  originalNumberInt642, originalNumberPrecisionUint2)\n"+
			"originalNumberInt642= '%v'\n"+
			"originalNumberPrecisionUint2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt642,
			originalNumberPrecisionUint2,
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

	return
}

func TestIntAry_AddBigIntToThis_01(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntToThis_01"

	originalNumberStr1 := "25"

	originalNumberInt2 := big.NewInt(50)

	originalNumberPrecisionInt2 := 0

	expectedNumberStr := "75"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddBigIntToThis(originalNumberInt2, originalNumberPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionInt2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionInt2,
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

	return
}

func TestIntAry_AddBigIntToThis_02(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntToThis_02"

	originalNumberStr1 := "100"

	originalNumberInt2 := big.NewInt(-25)

	originalNumberPrecisionInt2 := 0

	expectedNumberStr := "75"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddBigIntToThis(originalNumberInt2, originalNumberPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionInt2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionInt2,
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

	return
}

func TestIntAry_AddBigIntToThis_03(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntToThis_03"

	originalNumberStr1 := "-100"

	originalNumberInt2 := big.NewInt(-25)

	originalNumberPrecisionInt2 := 0

	expectedNumberStr := "-125"

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddBigIntToThis(originalNumberInt2, originalNumberPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionInt2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionInt2,
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

	return
}

func TestIntAry_AddBigIntToThis_04(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntToThis_04"

	originalNumberStr1 := "25.75"

	originalNumberInt2 := big.NewInt(5050)

	originalNumberPrecisionInt2 := 2

	expectedNumberStr := "76.25"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddBigIntToThis(originalNumberInt2, originalNumberPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionInt2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionInt2,
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

	return
}

func TestIntAry_AddBigIntToThis_05(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntToThis_05"

	originalNumberStr1 := "100.925"

	originalNumberInt2 := big.NewInt(-25967)

	originalNumberPrecisionInt2 := 3

	expectedNumberStr := "74.958"

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddBigIntToThis(originalNumberInt2, originalNumberPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionInt2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionInt2,
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

	return
}

func TestIntAry_AddBigIntToThis_06(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntToThis_06"

	originalNumberStr1 := "-100.35842"

	originalNumberInt2 := big.NewInt(-1256984)

	originalNumberPrecisionInt2 := 4

	expectedNumberStr := "-226.05682"

	expectedPrecisionUint := uint(5)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddBigIntToThis(originalNumberInt2, originalNumberPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionInt2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionInt2,
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

	return
}

func TestIntAry_AddBigIntToThis_07(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntToThis_07"

	originalNumberStr1 := "0"

	originalNumberInt2 := big.NewInt(0)

	originalNumberPrecisionInt2 := 0

	expectedNumberStr := "0"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddBigIntToThis(originalNumberInt2, originalNumberPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionInt2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionInt2,
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

	return
}

func TestIntAry_AddBigIntToThis_08(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntToThis_08"

	originalNumberStr1 := "0.00"

	originalNumberInt2 := big.NewInt(0)

	originalNumberPrecisionInt2 := 2

	expectedNumberStr := "0.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddBigIntToThis(originalNumberInt2, originalNumberPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntToThis(\n"+
			"  originalNumberInt2, originalNumberPrecisionInt2)\n"+
			"originalNumberInt2= '%v'\n"+
			"originalNumberPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumberInt2,
			originalNumberPrecisionInt2,
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

	return
}

func TestIntAry_AddBigIntNumToThis_01(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntNumToThis_01"

	originalNumberStr1 := "1.05"

	originalNumberStr2 := "2.37"

	expectedNumberStr := "3.42"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	bINum2, err := new(BigIntNum).NewNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).NewNumStr(originalNumberStr2)\n"+
			"originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumberStr, err := bINum2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumberStr, err := bINum2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != bINum2NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr2 != bINum2NumberStr\n"+
			"Expected bINum2NumberStr = '%v'\n"+
			"  Actual bINum2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, bINum2NumberStr)

		return
	}

	err = ia1.AddBigIntNumToThis(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntNumToThis(bINum2)\n"+
			"ia1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			bINum2NumberStr,
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

	return
}

func TestIntAry_AddBigIntNumToThis_02(t *testing.T) {

	ePrefix := "TestIntAry_AddBigIntNumToThis_02"

	originalNumberStr1 := "-1.05"

	originalNumberStr2 := "2.37"

	expectedNumberStr := "1.32"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	bINum2, err := new(BigIntNum).NewNumStr(originalNumberStr2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2, err := new(BigIntNum).NewNumStr(originalNumberStr2)\n"+
			"originalNumberStr2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr2, err.Error())
		return
	}

	err = bINum2.IsValid("Validating bINum2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum2.IsValid('Validating bINum2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINum2NumberStr, err := bINum2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum2NumberStr, err := bINum2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr2 != bINum2NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr2 != bINum2NumberStr\n"+
			"Expected bINum2NumberStr = '%v'\n"+
			"  Actual bINum2NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr2, bINum2NumberStr)

		return
	}

	err = ia1.AddBigIntNumToThis(bINum2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddBigIntNumToThis(bINum2)\n"+
			"ia1= '%v'\n"+
			"bINum2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			bINum2NumberStr,
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

	return
}

func TestIntAry_AddFloat32ToThis_01(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_01"

	originalNumberStr1 := "1.00"

	originalFloat32Number2 := float32(50.00)

	originalPrecisionInt2 := 2

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "51.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_02(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_02"

	originalNumberStr1 := "1.25"

	originalFloat32Number2 := float32(50.50)

	originalPrecisionInt2 := 2

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "51.75"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_03(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_03"

	originalNumberStr1 := "-1.25"

	originalFloat32Number2 := float32(50.50)

	originalPrecisionInt2 := 2

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "49.25"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_04(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_04"

	originalNumberStr1 := "-5.25787"

	originalFloat32Number2 := float32(-60.324)

	originalPrecisionInt2 := 3

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "-65.58187"

	expectedPrecisionUint := uint(5)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_05(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_05"

	originalNumberStr1 := "5.25787"

	originalFloat32Number2 := float32(-34.324)

	originalPrecisionInt2 := 3

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "-29.06613"

	expectedPrecisionUint := uint(5)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_06(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_06"

	originalNumberStr1 := "1245.25787"

	originalFloat32Number2 := float32(350.324)

	originalPrecisionInt2 := 3

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "1595.58187"

	expectedPrecisionUint := uint(5)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_07(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_07"

	originalNumberStr1 := "1245.25"

	originalFloat32Number2 := float32(84350.320000000)

	originalPrecisionInt2 := -1

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "85595.57"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_08(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_08"

	originalNumberStr1 := "1245.25"

	originalFloat32Number2 := float32(84350.325)

	originalPrecisionInt2 := 2

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "85595.58"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_09(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_09"

	originalNumberStr1 := "0"

	originalFloat32Number2 := float32(0.000)

	originalPrecisionInt2 := 0

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "0"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_10(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_10"

	originalNumberStr1 := "0.00"

	originalFloat32Number2 := float32(0.000)

	originalPrecisionInt2 := 2

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "0.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat32ToThis_11(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat32ToThis_11"

	originalNumberStr1 := "1.25"

	originalFloat32Number2 := float32(50.545)

	originalPrecisionInt2 := 2

	originalFloat32Number2Str :=
		strconv.FormatFloat(float64(originalFloat32Number2), 'f', originalPrecisionInt2, 32)

	expectedNumberStr := "51.80"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat32ToThis(originalFloat32Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat32ToThis(\n"+
			"  originalFloat32Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat32Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_01(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_01"

	originalNumberStr1 := "1.00"

	originalFloat64Number2 := 50.00

	originalPrecisionInt2 := 2

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "51.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_02(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_02"

	originalNumberStr1 := "1.25"

	originalFloat64Number2 := 50.50

	originalPrecisionInt2 := 2

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "51.75"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_03(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_03"

	originalNumberStr1 := "-1.25"

	originalFloat64Number2 := 50.50

	originalPrecisionInt2 := 2

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "49.25"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_04(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_04"

	originalNumberStr1 := "-5.25787"

	originalFloat64Number2 := -60.324

	originalPrecisionInt2 := 3

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "-65.58187"

	expectedPrecisionUint := uint(5)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_05(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_05"

	originalNumberStr1 := "5.25787"

	originalFloat64Number2 := -34.324

	originalPrecisionInt2 := 3

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "-29.06613"

	expectedPrecisionUint := uint(5)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_06(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_06"

	originalNumberStr1 := "1245.25787"

	originalFloat64Number2 := 350.324

	originalPrecisionInt2 := 3

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "1595.58187"

	expectedPrecisionUint := uint(5)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_07(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_07"

	originalNumberStr1 := "1245.25"

	originalFloat64Number2 := 84350.320000000

	originalPrecisionInt2 := -1

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "85595.57"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_08(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_08"

	originalNumberStr1 := "1245.25"

	originalFloat64Number2 := 84350.325

	originalPrecisionInt2 := 2

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "85595.58"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_09(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_09"

	originalNumberStr1 := "0"

	originalFloat64Number2 := 0.000

	originalPrecisionInt2 := 0

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "0"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloat64ToThis_10(t *testing.T) {

	ePrefix := "TestIntAry_AddFloat64ToThis_10"

	originalNumberStr1 := "0.00"

	originalFloat64Number2 := 0.000

	originalPrecisionInt2 := 2

	originalFloat64Number2Str :=
		strconv.FormatFloat(originalFloat64Number2, 'f', originalPrecisionInt2, 64)

	expectedNumberStr := "0.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloat64ToThis(originalFloat64Number2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloat64ToThis(\n"+
			"  originalFloat64Number2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalFloat32Number2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalFloat64Number2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_01(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_01"

	originalNumberStr1 := "1.00"

	originalBigFloatNumber2 := big.NewFloat(50.00)

	originalPrecisionInt2 := 2

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "51.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_02(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_02"

	originalNumberStr1 := "1.25"

	originalBigFloatNumber2 := big.NewFloat(50.50)

	originalPrecisionInt2 := 2

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "51.75"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_03(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_03"

	originalNumberStr1 := "-1.25"

	originalBigFloatNumber2 := big.NewFloat(50.50)

	originalPrecisionInt2 := 2

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "49.25"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_04(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_04"

	originalNumberStr1 := "-5.25787"

	originalBigFloatNumber2 := big.NewFloat(-60.324)

	originalPrecisionInt2 := 3

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "-65.58187"

	expectedPrecisionUint := uint(5)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_05(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_05"

	originalNumberStr1 := "5.25787"

	originalBigFloatNumber2 := big.NewFloat(-34.324)

	originalPrecisionInt2 := 3

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "-29.06613"

	expectedPrecisionUint := uint(5)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_06(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_06"

	originalNumberStr1 := "1245.25787"

	originalBigFloatNumber2 := big.NewFloat(350.324)

	originalPrecisionInt2 := 3

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "1595.58187"

	expectedPrecisionUint := uint(5)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_07(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_07"

	originalNumberStr1 := "1245.25"

	originalBigFloatNumber2 := big.NewFloat(84350.320000000)

	originalPrecisionInt2 := -1

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "85595.57"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_08(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_08"

	originalNumberStr1 := "1245.25"

	originalBigFloatNumber2 := big.NewFloat(84350.325)

	originalPrecisionInt2 := 2

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "85595.58"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_09(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_09"

	originalNumberStr1 := "0"

	originalBigFloatNumber2 := big.NewFloat(0.000)

	originalPrecisionInt2 := 0

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "0"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_AddFloatBigToThis_10(t *testing.T) {

	ePrefix := "TestIntAry_AddFloatBigToThis_10"

	originalNumberStr1 := "0.00"

	originalBigFloatNumber2 := big.NewFloat(0.000)

	originalPrecisionInt2 := 2

	originalBigFloatNumber2Str := originalBigFloatNumber2.Text('f', originalPrecisionInt2)

	expectedNumberStr := "0.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	ia1, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia1, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
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

	err = ia1.AddFloatBigToThis(originalBigFloatNumber2, originalPrecisionInt2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = ia1.AddFloatBigToThis(\n"+
			"  originalBigFloatNumber2, originalPrecisionInt2)\n"+
			"ia1= '%v'\n"+
			"originalBigFloatNumber2= '%v'\n"+
			"originalPrecisionInt2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			ia1NumberStr,
			originalBigFloatNumber2Str,
			originalPrecisionInt2,
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

	return
}

func TestIntAry_Ceiling_01(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_01"

	originalNumberStr1 := "0.925"

	expectedNumberStr := "1.000"

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Ceiling_02(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_02"

	originalNumberStr1 := "-2.7"

	expectedNumberStr := "-2.0"

	expectedPrecisionUint := uint(1)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Ceiling_03(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_03"

	originalNumberStr1 := "2.9"

	expectedNumberStr := "3.0"

	expectedPrecisionUint := uint(1)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}
func TestIntAry_Ceiling_04(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_04"

	originalNumberStr1 := "2.0"

	expectedNumberStr := "2.0"

	expectedPrecisionUint := uint(1)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Ceiling_05(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_05"

	originalNumberStr1 := "2.4"

	expectedNumberStr := "3.0"

	expectedPrecisionUint := uint(1)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Ceiling_06(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_06"

	originalNumberStr1 := "2.9"

	expectedNumberStr := "3.0"

	expectedPrecisionUint := uint(1)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Ceiling_07(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_07"

	originalNumberStr1 := "-2"

	expectedNumberStr := "-2"

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Ceiling_08(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_08"

	originalNumberStr1 := "-5.05"

	expectedNumberStr := "-5.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Ceiling_09(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_09"

	originalNumberStr1 := "5.05"

	expectedNumberStr := "6.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Ceiling_10(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_10"

	originalNumberStr1 := "5.95"

	expectedNumberStr := "6.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Ceiling_11(t *testing.T) {

	ePrefix := "TestIntAry_Ceiling_11"

	originalNumberStr1 := "5"

	expectedNumberStr := "5"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Ceiling()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Ceiling()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_ChangeSign_01(t *testing.T) {

	ePrefix := "TestIntAry_ChangeSign_01"

	originalNumberStr1 := "-572"

	expectedNumberStr := "572"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.ChangeSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating changesign intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating changesign intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAryNumberStr after changesign\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignVal != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_ChangeSign_02(t *testing.T) {

	ePrefix := "TestIntAry_ChangeSign_02"

	originalNumberStr1 := "572"

	expectedNumberStr := "-572"

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.ChangeSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating changesign intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating changesign intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAryNumberStr after changesign\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignVal != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_ChangeSign_03(t *testing.T) {

	ePrefix := "TestIntAry_ChangeSign_03"

	originalNumberStr1 := "0"

	expectedNumberStr := "0"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.ChangeSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating changesign intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating changesign intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAryNumberStr after changesign\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignVal != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_ChangeSign_04(t *testing.T) {

	ePrefix := "TestIntAry_ChangeSign_04"

	originalNumberStr1 := "-12.456"

	expectedNumberStr := "12.456"

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.ChangeSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating changesign intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating changesign intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAryNumberStr after changesign\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignVal != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_ChangeSign_05(t *testing.T) {

	ePrefix := "TestIntAry_ChangeSign_05"

	originalNumberStr1 := "12.456"

	expectedNumberStr := "-12.456"

	expectedPrecisionUint := uint(3)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr1)\n"+
			"originalNumberStr1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr1, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry.ChangeSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.ChangeSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating changesign intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating changesign intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err = intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err = intAry.GetNumStr()\n"+
			"intAryNumberStr after changesign\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignVal != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_Equals_01(t *testing.T) {

	ePrefix := "TestIntAry_Equals_01"

	originalNumberStr1 := "000549721.32178000"

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry1.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2 := new(IntAry).New()

	err = intAry2.CopyIn(&intAry1, false)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.CopyIn(&intAry1, false)\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !intAry1.Equals(&intAry2) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAry1 NOT EQUAL to intAry2\n"+
			"Expected intAry2 = '%v'\n"+
			"  Actual intAry2 = '%v'\n\n",
			ePrefix, intAry1NumberStr, intAry2NumberStr)

		return
	}

	return
}

func TestIntAry_Equals_02(t *testing.T) {

	ePrefix := "TestIntAry_Equals_02"

	originalNumberStr1 := "-000549721.32178000"

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry1.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	err = intAry1.CopyToBackUp()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry2 := new(IntAry).New()

	err = intAry2.CopyIn(&intAry1, true)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.CopyIn(&intAry1, false)\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !intAry1.Equals(&intAry2) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAry1 NOT EQUAL to intAry2\n"+
			"Expected intAry2 = '%v'\n"+
			"  Actual intAry2 = '%v'\n\n",
			ePrefix, intAry1NumberStr, intAry2NumberStr)

		return
	}

	intAry1BackupNumberStr := intAry1.BackUp.GetNumStr()

	intAry2BackupNumberStr := intAry2.BackUp.GetNumStr()

	if !intAry1.BackUp.Equals(&intAry2.BackUp) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAry1.BackUp NOT EQUAL to intAry2.BackUp\n"+
			"Expected intAry2.BackUp = '%v'\n"+
			"  Actual intAry2.BackUp = '%v'\n\n",
			ePrefix, intAry1BackupNumberStr, intAry2BackupNumberStr)

		return
	}

	return
}

func TestIntAry_Equals_03(t *testing.T) {

	ePrefix := "TestIntAry_Equals_01"

	originalNumberStr1 := "-000549721.32178000"

	expectedNumberStr := "000549721.32178000"

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry1.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	err = intAry1.CopyToBackUp()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.CopyToBackUp()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry2 := new(IntAry).New()

	err = intAry2.CopyIn(&intAry1, true)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.CopyIn(&intAry1, true)\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = intAry2.SetSign(1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2 After Sign Change")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2 After Sign Change')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err = intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"intAry2NumberStr after sign change\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if intAry1.Equals(&intAry2) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAry1 is EQUAL TO intAry2\n"+
			"  after sign change.\n"+
			"Expected intAry2 = '%v'\n"+
			"  Actual intAry2 = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	return
}

func TestIntAry_Equals_04(t *testing.T) {

	ePrefix := "TestIntAry_Equals_04"

	originalNumberStr1 := "-000549721.32178000"

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry1.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	err = intAry1.CopyToBackUp()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry2 := new(IntAry).New()

	err = intAry2.CopyIn(&intAry1, true)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.CopyIn(&intAry1, true)\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	intAry2.BackUp.SetSignValue(1)

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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if !intAry1.Equals(&intAry2) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAry1 NOT EQUAL to intAry2\n"+
			"Expected intAry2 = '%v'\n"+
			"  Actual intAry2 = '%v'\n\n",
			ePrefix, intAry1NumberStr, intAry2NumberStr)

		return
	}

	intAry1BackupNumberStr := intAry1.BackUp.GetNumStr()

	intAry2BackupNumberStr := intAry2.BackUp.GetNumStr()

	if intAry1.BackUp.Equals(&intAry2.BackUp) {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because intAry1.BackUp IS EQUAL TO intAry2.BackUp\n"+
			" after sign change.\n"+
			"Expected intAry2.BackUp = '%v'\n"+
			"  Actual intAry2.BackUp = '%v'\n\n",
			ePrefix, intAry1BackupNumberStr, intAry2BackupNumberStr)

		return
	}

	return
}

func TestIntAry_Floor_01(t *testing.T) {

	ePrefix := "TestIntAry_Floor_01"

	originalNumberStr1 := "99.925"

	expectedNumberStr := "99.000"

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_02(t *testing.T) {

	ePrefix := "TestIntAry_Floor_02"

	originalNumberStr1 := "-99.925"

	expectedNumberStr := "-100.000"

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_03(t *testing.T) {

	ePrefix := "TestIntAry_Floor_03"

	originalNumberStr1 := "0.925"

	expectedNumberStr := "0.000"

	expectedPrecisionUint := uint(3)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_04(t *testing.T) {

	ePrefix := "TestIntAry_Floor_04"

	originalNumberStr1 := "2.0"

	expectedNumberStr := "2.0"

	expectedPrecisionUint := uint(1)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_05(t *testing.T) {

	ePrefix := "TestIntAry_Floor_05"

	originalNumberStr1 := "-2.7"

	expectedNumberStr := "-3.0"

	expectedPrecisionUint := uint(1)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_06(t *testing.T) {

	ePrefix := "TestIntAry_Floor_06"

	originalNumberStr1 := "-2"

	expectedNumberStr := "-2"

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_07(t *testing.T) {

	ePrefix := "TestIntAry_Floor_07"

	originalNumberStr1 := "2.9"

	expectedNumberStr := "2.0"

	expectedPrecisionUint := uint(1)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_08(t *testing.T) {

	ePrefix := "TestIntAry_Floor_08"

	originalNumberStr1 := "5.05"

	expectedNumberStr := "5.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_09(t *testing.T) {

	ePrefix := "TestIntAry_Floor_09"

	originalNumberStr1 := "-5.05"

	expectedNumberStr := "-6.00"

	expectedPrecisionUint := uint(2)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_10(t *testing.T) {

	ePrefix := "TestIntAry_Floor_10"

	originalNumberStr1 := "-2.7"

	expectedNumberStr := "-3.0"

	expectedPrecisionUint := uint(1)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_11(t *testing.T) {

	ePrefix := "TestIntAry_Floor_11"

	originalNumberStr1 := "-2"

	expectedNumberStr := "-2"

	expectedPrecisionUint := uint(0)

	expectedSignVal := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_Floor_12(t *testing.T) {

	ePrefix := "TestIntAry_Floor_12"

	originalNumberStr1 := "2"

	expectedNumberStr := "2"

	expectedPrecisionUint := uint(0)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr1)\n"+
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

	intAry2, err := intAry1.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.Floor()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
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

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignVal != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}
