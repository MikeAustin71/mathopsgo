package mathops

import (
	"math/big"
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
	ia1, _ := IntAry{}.NewNumStr("100")

	num := int64(-25)
	precision := uint(0)

	ia1.AddInt64ToThis(num, precision)

	expected := "75"

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}
}

func TestIntAry_AddInt64ToThis_03(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("-100")

	num := int64(-25)
	precision := uint(0)

	ia1.AddInt64ToThis(num, precision)

	expected := "-125"

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecisionUint() != precision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", precision, ia1.GetPrecision())
	}
}

func TestIntAry_AddInt64ToThis_04(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("25.75")

	num := int64(5050)
	precision := uint(2)
	expected := "76.25"

	ia1.AddInt64ToThis(num, precision)

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}
}

func TestIntAry_AddInt64ToThis_05(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("100.925")

	num := int64(-25967)
	precision := uint(3)

	ia1.AddInt64ToThis(num, precision)

	expected := "74.958"

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if precision != ia1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'.  Instead precision='%v'", precision, ia1.GetPrecision())
	}

}

func TestIntAry_AddInt64ToThis_06(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("-100.35842")

	num := int64(-1256984)
	precision := uint(4)
	outPrecision := uint(5)

	ia1.AddInt64ToThis(num, precision)

	expected := "-226.05682"

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecisionUint() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddInt64ToThis_07(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("0")

	num := int64(0)
	precision := uint(0)
	outPrecision := uint(0)

	ia1.AddInt64ToThis(num, precision)

	expected := "0"

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecisionUint() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddInt64ToThis_08(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("0.00")

	num := int64(000)
	precision := uint(2)
	outPrecision := uint(2)

	ia1.AddInt64ToThis(num, precision)

	expected := "0.00"

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecisionUint() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddBigIntToThis_01(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("25")

	num := big.NewInt(50)
	precision := 0
	expected := "75"

	err := ia1.AddBigIntToThis(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia1.AddBigIntToThis(50, 0). err='%v'", err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}
}

func TestIntAry_AddBigIntToThis_02(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("100")

	num := big.NewInt(-25)
	precision := 0

	err := ia1.AddBigIntToThis(num, precision)

	expected := "75"

	if err != nil {
		t.Errorf("Error returned from ia1.AddBigIntToThis(num, 0). num='%v' err='%v'", num, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}
}

func TestIntAry_AddBigIntToThis_03(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("-100")

	num := big.NewInt(-25)
	precision := 0

	err := ia1.AddBigIntToThis(num, precision)

	expected := "-125"

	if err != nil {
		t.Errorf("Error returned from ia1.AddBigIntToThis(num, 0). num='%v' err='%v'", num, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != precision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", precision, ia1.GetPrecision())
	}
}

func TestIntAry_AddBigIntToThis_04(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("25.75")

	num := big.NewInt(5050)
	precision := 2
	expected := "76.25"

	err := ia1.AddBigIntToThis(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia1.AddBigIntToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}
}

func TestIntAry_AddBigIntToThis_05(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("100.925")

	num := big.NewInt(-25967)
	precision := 3

	err := ia1.AddBigIntToThis(num, precision)

	expected := "74.958"

	if err != nil {
		t.Errorf("Error returned from ia1.AddBigIntToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if precision != ia1.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'.  Instead precision='%v'", precision, ia1.GetPrecision())
	}

}

func TestIntAry_AddBigIntToThis_06(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("-100.35842")

	num := big.NewInt(-1256984)
	precision := 4
	outPrecision := 5

	err := ia1.AddBigIntToThis(num, precision)

	expected := "-226.05682"

	if err != nil {
		t.Errorf("Error returned from ia1.AddBigIntToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddBigIntToThis_07(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("0")

	num := big.NewInt(0)
	precision := 0
	outPrecision := 0

	err := ia1.AddBigIntToThis(num, precision)

	expected := "0"

	if err != nil {
		t.Errorf("Error returned from ia1.AddBigIntToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddBigIntToThis_08(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("0.00")

	num := big.NewInt(000)
	precision := 2
	outPrecision := 2

	err := ia1.AddBigIntToThis(num, precision)

	expected := "0.00"

	if err != nil {
		t.Errorf("Error returned from ia1.AddBigIntToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddBigIntNumToThis_01(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1.05")

	bINum, _ := BigIntNum{}.NewNumStr("2.37")

	expectedStr := "3.42"

	err := ia1.AddBigIntNumToThis(bINum)

	if err != nil {
		t.Errorf("Error returned by ia1.AddBigIntNumToThis(bINum). "+
			"bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())
	}

	if expectedStr != ia1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, ia1.GetNumStr())
	}

}

func TestIntAry_AddBigIntNumToThis_02(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("-1.05")

	bINum, _ := BigIntNum{}.NewNumStr("2.37")

	expectedStr := "1.32"

	err := ia1.AddBigIntNumToThis(bINum)

	if err != nil {
		t.Errorf("Error returned by ia1.AddBigIntNumToThis(bINum). "+
			"bINum='%v' Error='%v' ", bINum.GetNumStr(), err.Error())
	}

	if expectedStr != ia1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, ia1.GetNumStr())
	}

}

func TestIntAry_AddFloat32ToThis_01(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("1.00")

	num := float32(50.00)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "51.00"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_02(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1.25")

	num := float32(50.50)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "51.75"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_03(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("-1.25")

	num := float32(50.50)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "49.25"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_04(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("-5.25787")

	num := float32(-60.324)
	precision := 3
	outPrecision := 5

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "-65.58187"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_05(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("5.25787")

	num := float32(-34.324)
	precision := 3
	outPrecision := 5

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "-29.06613"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_06(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1245.25787")

	num := float32(350.324)
	precision := 3
	outPrecision := 5

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "1595.58187"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_07(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1245.25")

	num := float32(84350.320000000)
	precision := -1
	outPrecision := 2

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "85595.57"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_08(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1245.25")

	num := float32(84350.325)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "85595.58"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_09(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("0")

	num := float32(0.000)
	precision := 0
	outPrecision := 0

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "0"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_10(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("0.00")

	num := float32(0.000)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "0.00"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat32ToThis_11(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1.25")

	num := float32(50.545)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat32ToThis(num, precision)

	expected := "51.80"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat32ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_01(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("1.00")

	num := 50.00
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "51.00"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_02(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1.25")

	num := 50.50
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "51.75"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_03(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("-1.25")

	num := 50.50
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "49.25"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_04(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("-5.25787")

	num := -60.324
	precision := 3
	outPrecision := 5

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "-65.58187"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_05(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("5.25787")

	num := -34.324
	precision := 3
	outPrecision := 5

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "-29.06613"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_06(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1245.25787")

	num := 350.324
	precision := 3
	outPrecision := 5

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "1595.58187"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_07(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1245.25")

	num := 84350.320000000
	precision := -1
	outPrecision := 2

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "85595.57"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_08(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1245.25")

	num := 84350.325
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "85595.58"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_09(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("0")

	num := 0.000
	precision := 0
	outPrecision := 0

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "0"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloat64ToThis_10(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("0.00")

	num := 0.000
	precision := 2
	outPrecision := 2

	err := ia1.AddFloat64ToThis(num, precision)

	expected := "0.00"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloat64ToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloatBigToThis_01(t *testing.T) {
	ia1, _ := IntAry{}.NewNumStr("1.00")

	num := big.NewFloat(50.00)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "51.00"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloatBigToThis_02(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1.25")

	num := big.NewFloat(50.50)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "51.75"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloatBigToThis_03(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("-1.25")

	num := big.NewFloat(50.50)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "49.25"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloatBigToThis_04(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("-5.25787")

	num := big.NewFloat(-60.324)
	precision := 3
	outPrecision := 5

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "-65.58187"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloatBigToThis_05(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("5.25787")

	num := big.NewFloat(-34.324)
	precision := 3
	outPrecision := 5

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "-29.06613"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloatBigToThis_06(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1245.25787")

	num := big.NewFloat(350.324)
	precision := 3
	outPrecision := 5

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "1595.58187"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}

}

func TestIntAry_AddFloatBigToThis_07(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1245.25")

	num := big.NewFloat(84350.320000000)
	precision := -1
	outPrecision := 2

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "85595.57"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloatBigToThis_08(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("1245.25")

	num := big.NewFloat(84350.325)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "85595.58"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloatBigToThis_09(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("0")

	num := big.NewFloat(0.000)
	precision := 0
	outPrecision := 0

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "0"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_AddFloatBigToThis_10(t *testing.T) {

	ia1, _ := IntAry{}.NewNumStr("0.00")

	num := big.NewFloat(0.000)
	precision := 2
	outPrecision := 2

	err := ia1.AddFloatBigToThis(num, precision)

	expected := "0.00"

	if err != nil {
		t.Errorf("Error returned from ia1.AddFloatBigToThis(num, precision). num='%v' precision='%v' err='%v'", num, precision, err)
	}

	result := ia1.GetNumStr()

	if expected != result {
		t.Errorf("Expected result='%v'. Instead, result='%v' ", expected, result)
	}

	if ia1.GetPrecision() != outPrecision {
		t.Errorf("Expected ia1.GetPrecisionInt= '%v'.  Instead, ia1.GetPrecisionInt='%v' ", outPrecision, ia1.GetPrecision())
	}
}

func TestIntAry_Ceiling_01(t *testing.T) {
	nStr1 := "0.925"
	expected := "1.000"
	precision := 3

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Ceiling_02(t *testing.T) {
	nStr1 := "-2.7"
	expected := "-2.0"
	precision := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Ceiling_03(t *testing.T) {
	nStr1 := "2.9"
	expected := "3.0"
	precision := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}
func TestIntAry_Ceiling_04(t *testing.T) {
	nStr1 := "2.0"
	expected := "2.0"
	precision := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Ceiling_05(t *testing.T) {
	nStr1 := "2.4"
	expected := "3.0"
	precision := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Ceiling_06(t *testing.T) {
	nStr1 := "2.9"
	expected := "3.0"
	precision := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Ceiling_07(t *testing.T) {
	nStr1 := "-2"
	expected := "-2"
	precision := 0

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Ceiling_08(t *testing.T) {
	nStr1 := "-5.05"
	expected := "-5.00"
	precision := 2

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Ceiling_09(t *testing.T) {
	nStr1 := "5.05"
	expected := "6.00"
	precision := 2

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Ceiling_10(t *testing.T) {
	nStr1 := "5.95"
	expected := "6.00"
	precision := 2

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Ceiling_11(t *testing.T) {
	nStr1 := "5"
	expected := "5"
	precision := 0

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Ceiling()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_ChangeSign_01(t *testing.T) {
	nStr := "-572"
	expectedStr := "572"

	ia, err := IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	ia.ChangeSign()

	if expectedStr != ia.GetNumStr() {
		t.Errorf("Error: Expected ia.GetNumStr='%v'. Instead, ia.GetNumStr='%v' ",
			expectedStr, ia.GetNumStr())
	}

}

func TestIntAry_ChangeSign_02(t *testing.T) {
	nStr := "572"
	expectedStr := "-572"

	ia, err := IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	ia.ChangeSign()

	if expectedStr != ia.GetNumStr() {
		t.Errorf("Error: Expected ia.GetNumStr='%v'. Instead, ia.GetNumStr='%v' ",
			expectedStr, ia.GetNumStr())
	}

}

func TestIntAry_ChangeSign_03(t *testing.T) {
	nStr := "0"
	expectedStr := "0"

	ia, err := IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	ia.ChangeSign()

	if expectedStr != ia.GetNumStr() {
		t.Errorf("Error: Expected ia.GetNumStr='%v'. Instead, ia.GetNumStr='%v' ",
			expectedStr, ia.GetNumStr())
	}

}

func TestIntAry_ChangeSign_04(t *testing.T) {
	nStr := "-12.456"
	expectedStr := "12.456"

	ia, err := IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	ia.ChangeSign()

	if expectedStr != ia.GetNumStr() {
		t.Errorf("Error: Expected ia.GetNumStr='%v'. Instead, ia.GetNumStr='%v' ",
			expectedStr, ia.GetNumStr())
	}

}

func TestIntAry_ChangeSign_05(t *testing.T) {
	nStr := "12.456"
	expectedStr := "-12.456"

	ia, err := IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	ia.ChangeSign()

	if expectedStr != ia.GetNumStr() {
		t.Errorf("Error: Expected ia.GetNumStr='%v'. Instead, ia.GetNumStr='%v' ",
			expectedStr, ia.GetNumStr())
	}

}

func TestIntAry_Equals_01(t *testing.T) {
	nStr1 := "000549721.32178000"

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	ia2 := IntAry{}.New()
	ia2.CopyIn(&ia, false)

	if !ia.Equals(&ia2) {
		t.Error("Error: ia NOT EQUAL to ia2!")
	}

}

func TestIntAry_Equals_02(t *testing.T) {
	nStr1 := "-000549721.32178000"

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	ia.CopyToBackUp()

	ia2 := IntAry{}.New()
	ia2.CopyIn(&ia, true)

	if !ia.Equals(&ia2) {
		t.Error("Error: ia NOT EQUAL to ia2!")
	}

	if !ia.BackUp.Equals(&ia.BackUp) {
		t.Error("Error: ia.Backup != ia2.Backup!")
	}

}

func TestIntAry_Equals_03(t *testing.T) {
	nStr1 := "-000549721.32178000"

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr1)
	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	ia.CopyToBackUp()

	ia2 := IntAry{}.New()

	ia2.CopyIn(&ia, true)

	err = ia2.SetSign(1)

	if err != nil {
		t.Errorf("Received Error from ia.SetSign(1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	if ia.Equals(&ia2) {
		t.Error("Error: ia EQUALS ia2!")
	}

}

func TestIntAry_Equals_04(t *testing.T) {
	nStr1 := "-000549721.32178000"

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	ia.CopyToBackUp()

	ia2 := IntAry{}.New()
	ia2.CopyIn(&ia, true)

	ia2.BackUp.SetSignValue(1)

	if !ia.Equals(&ia2) {
		t.Error("Error: ia NOT EQUAL ia2!")
	}

	if ia.BackUp.Equals(&ia2.BackUp) {
		t.Error("Error: ia.BackUp SHOULD NOT EQUAL ia2.BackUp")
	}

}

func TestIntAry_Floor_01(t *testing.T) {
	nStr1 := "99.925"
	expected := "99.000"
	precision := 3

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_02(t *testing.T) {
	nStr1 := "-99.925"
	expected := "-100.000"
	precision := 3

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_03(t *testing.T) {
	nStr1 := "0.925"
	expected := "0.000"
	precision := 3

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_04(t *testing.T) {
	nStr1 := "2.0"
	expected := "2.0"
	precision := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_05(t *testing.T) {
	nStr1 := "-2.7"
	expected := "-3.0"
	precision := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_06(t *testing.T) {
	nStr1 := "-2"
	expected := "-2"
	precision := 0

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_07(t *testing.T) {
	nStr1 := "2.9"
	expected := "2.0"
	precision := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_08(t *testing.T) {
	nStr1 := "5.05"
	expected := "5.00"
	precision := 2

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_09(t *testing.T) {
	nStr1 := "-5.05"
	expected := "-6.00"
	precision := 2

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_10(t *testing.T) {
	nStr1 := "-2.7"
	expected := "-3.0"
	precision := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_11(t *testing.T) {
	nStr1 := "-2"
	expected := "-2"
	precision := 0

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_12(t *testing.T) {
	nStr1 := "2"
	expected := "2"
	precision := 0

	ia := IntAry{}.New()

	err := ia.SetIntAryWithNumStr(nStr1)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr1)\n"+
			"nStr1= '%v'\n Error= %v",
			nStr1, err.Error())
	}

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStrDto= '%v'. Instead, got numStrDto='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}
