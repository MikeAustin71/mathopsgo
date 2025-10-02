package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntNum_SetPrecision_01(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_01"

	originalNumStr := "654.123456"

	newPrecision := uint(3)

	expectedNumberStr := "654.123"

	expectedPrecisionUint := uint(3)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_02(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_02"

	originalNumStr := "654.123456"

	newPrecision := uint(4)

	expectedNumberStr := "654.1235"

	expectedPrecisionUint := uint(4)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_03(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_03"

	originalNumStr := "654.123456"

	newPrecision := uint(0)

	expectedNumberStr := "654"

	expectedPrecisionUint := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_04(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_04"

	originalNumStr := "-654.123456"

	newPrecision := uint(3)

	expectedNumberStr := "-654.123"

	expectedPrecisionUint := uint(3)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_05(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_05"

	originalNumStr := "-654.123456"

	newPrecision := uint(4)

	expectedNumberStr := "-654.1235"

	expectedPrecisionUint := uint(4)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_06(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_06"

	originalNumStr := "654"

	newPrecision := uint(3)

	expectedNumberStr := "654.000"

	expectedPrecisionUint := uint(3)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_07(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_07"

	originalNumStr := "654.123456"

	newPrecision := uint(9)

	expectedNumberStr := "654.123456000"

	expectedPrecisionUint := uint(9)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_08(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_08"

	originalNumStr := "-654"

	newPrecision := uint(9)

	expectedNumberStr := "-654.000000000"

	expectedPrecisionUint := uint(9)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_09(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_09"

	originalNumStr := "-654.123456"

	newPrecision := uint(9)

	expectedNumberStr := "-654.123456000"

	expectedPrecisionUint := uint(9)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_10(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_10"

	originalNumStr := "0"

	newPrecision := uint(4)

	expectedNumberStr := "0.0000"

	expectedPrecisionUint := uint(4)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_SetPrecision_11(t *testing.T) {

	ePrefix := "TestBigIntNum_SetPrecision_11"

	originalNumStr := "0.0000"

	newPrecision := uint(0)

	expectedNumberStr := "0"

	expectedPrecisionUint := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedNumberStr, &expectedNumSeps)\n"+
			"expectedNumberStr='%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumberStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumberStr, err := expectedBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumberStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigINumberStr)

		return
	}

	expectedBigINumPrecisionUint, err := expectedBigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumPrecisionUint, err :=\n"+
			"  expectedBigINum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumberStr, err.Error())

		return
	}

	if expectedPrecisionUint != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/ expectedBigINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, expectedBigINumPrecisionUint)

		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.SetPrecision(newPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetPrecision(newPrecision)\n"+
			"newPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, newPrecision, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Precision")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Precision')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"After new 'bINum' Precision\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: 'bINumNumberStr' NOT EQUAL to expectedBigINumberStr\n"+
			"Because expectedBigINumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigINumPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedBigINumPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionUint, bINumPrecisionUint)

		return
	}

	expectedBigINumEqualsBINum, err := expectedBigINum.Equal(bINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumEqualsBINum, err :=\n"+
			"  expectedBigINum.Equal(bINum)\n"+
			"expectedBigINum= '%v'\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, bINumNumberStr, err.Error())
		return
	}

	if !expectedBigINumEqualsBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'bINum' values ARE NOT Equal\n"+
			"Because expectedBigINumEqualsBINum = 'false' \n"+
			"Expected bINum = '%v'\n"+
			"  Actual bINum = '%v'\n\n",
			ePrefix, expectedBigINumberStr, bINumNumberStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TrimTrailingFracZeros_01(t *testing.T) {

	ePrefix := "TestBigIntNum_TrimTrailingFracZeros_01"

	originalNumStr := "-123.000"

	expectedNumberStr := "-123"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "-123"

	expectedAbsBigIntNumStr := "123"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := -1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TrimTrailingFracZeros()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TrimTrailingFracZeros()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TrimTrailingFracZeros_02(t *testing.T) {

	ePrefix := "TestBigIntNum_TrimTrailingFracZeros_02"

	originalNumStr := "123.000"

	expectedNumberStr := "123"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "123"

	expectedAbsBigIntNumStr := "123"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TrimTrailingFracZeros()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TrimTrailingFracZeros()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TrimTrailingFracZeros_03(t *testing.T) {

	ePrefix := "TestBigIntNum_TrimTrailingFracZeros_03"

	originalNumStr := "123.0090"

	expectedNumberStr := "123.009"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "123009"

	expectedAbsBigIntNumStr := "123009"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TrimTrailingFracZeros()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TrimTrailingFracZeros()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TrimTrailingFracZeros_04(t *testing.T) {

	ePrefix := "TestBigIntNum_TrimTrailingFracZeros_04"

	originalNumStr := "-123.0090"

	expectedNumberStr := "-123.009"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(3)

	expectedBigIntNumStr := "-1230090"

	expectedAbsBigIntNumStr := "1230090"

	expectedScaleFactor := big.NewInt(1000)

	expectedSignVal := -1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TrimTrailingFracZeros()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TrimTrailingFracZeros()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TrimTrailingFracZeros_05(t *testing.T) {

	ePrefix := "TestBigIntNum_TrimTrailingFracZeros_05"

	originalNumStr := "0.000"

	expectedNumberStr := "0"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := uint(0)

	expectedBigIntNumStr := "0"

	expectedAbsBigIntNumStr := "0"

	expectedScaleFactor := big.NewInt(1)

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TrimTrailingFracZeros()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TrimTrailingFracZeros()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_01(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_01"

	originalNumStr := "-123.567"

	truncToDecPlace := uint(2)

	expectedNumberStr := "-123.56"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "-12356"

	expectedAbsBigIntNumStr := "12356"

	expectedSignVal := -1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_02(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_02"

	originalNumStr := "123.567"

	truncToDecPlace := uint(2)

	expectedNumberStr := "123.56"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "12356"

	expectedAbsBigIntNumStr := "12356"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_03(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_03"

	originalNumStr := "123.567"

	truncToDecPlace := uint(3)

	expectedNumberStr := "123.567"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "123567"

	expectedAbsBigIntNumStr := "123567"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_04(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_04"

	originalNumStr := "123.567"

	truncToDecPlace := uint(4)

	expectedNumberStr := "123.5670"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "123567"

	expectedAbsBigIntNumStr := "123567"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_05(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_05"

	originalNumStr := "-123.567"

	truncToDecPlace := uint(4)

	expectedNumberStr := "-123.5670"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "-1235670"

	expectedAbsBigIntNumStr := "1235670"

	expectedSignVal := -1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_06(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_06"

	originalNumStr := "0.000"

	truncToDecPlace := uint(2)

	expectedNumberStr := "0.00"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "000"

	expectedAbsBigIntNumStr := "000"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_07(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_07"

	originalNumStr := "654.123456"

	truncToDecPlace := uint(3)

	expectedNumberStr := "654.123"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "654123"

	expectedAbsBigIntNumStr := "654123"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_08(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_08"

	originalNumStr := "654.123456789"

	truncToDecPlace := uint(4)

	expectedNumberStr := "654.1234"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "6541234"

	expectedAbsBigIntNumStr := "6541234"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_09(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_08"

	originalNumStr := "654.123456789"

	truncToDecPlace := uint(0)

	expectedNumberStr := "654"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "654"

	expectedAbsBigIntNumStr := "654"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_10(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_10"

	truncToDecPlace := uint(5)

	originalNumStr := "654"

	expectedNumberStr := "654.00000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "65400000"

	expectedAbsBigIntNumStr := "65400000"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_11(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_11"

	originalNumStr := "654.123"

	truncToDecPlace := uint(9)

	expectedNumberStr := "654.123000000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "654123000000"

	expectedAbsBigIntNumStr := "654123000000"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_12(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_12"

	originalNumStr := "0"

	truncToDecPlace := uint(6)

	expectedNumberStr := "0.000000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "0000000"

	expectedAbsBigIntNumStr := "0000000"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_13(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_13"

	originalNumStr := "0.000000"

	truncToDecPlace := uint(0)

	expectedNumberStr := "0"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "0"

	expectedAbsBigIntNumStr := "0"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_TruncToDecPlace_14(t *testing.T) {

	ePrefix := "TestBigIntNum_TruncToDecPlace_14"

	originalNumStr := "654.123456789015"

	truncToDecPlace := uint(11)

	expectedNumberStr := "654.12345678901"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedPrecisionUint := truncToDecPlace

	expectedScaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(expectedPrecisionUint)), nil)

	expectedBigIntNumStr := "65412345678901"

	expectedAbsBigIntNumStr := "65412345678901"

	expectedSignVal := 1

	expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedBigIntNumStr, 10)\n"+
			"expectedBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedBigIntNumStr)
		return
	}

	expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

	if !isOk {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
			"   SetString(expectedAbsBigIntNumStr, 10)\n"+
			"expectedAbsBigIntNumStr= '%v'\n"+
			"Error= 'isOk == false'\n\n",
			ePrefix, expectedAbsBigIntNumStr)
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNumberStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Original & bINum Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumNumberStr)

		return
	}

	err = bINum.TruncToDecPlace(truncToDecPlace)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.TruncToDecPlace(truncToDecPlace)\n"+
			"bINum= '%v'\n"+
			"truncToDecPlace= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, truncToDecPlace, err.Error())
		return
	}

	bINumNumberStr, err = bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNumberStr, err = bINum.GetNumStr()\n"+
			"After Truncating Fractional Zeros\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != bINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != bINumNumberStr \n"+
			"Expected bINumNumberStr = '%v'\n"+
			"  Actual bINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bINumNumberStr)

		return
	}

	bINumBigInt, err := bINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBigInt, err := bINum.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
			"Expected bINumBigInt = '%v'\n"+
			"  Actual bINumBigInt = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

		return
	}

	bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
			"Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
			"Expected bINumAbsBigInt = '%v'\n"+
			"  Actual bINumAbsBigInt = '%v'\n\n",
			ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

		return
	}

	bINumPrecisionUint, err := bINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
			"bINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedPrecisionUint != bINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bINumPrecisionUint\n"+
			"Expected bINumPrecisionUint = '%v'\n"+
			"  Actual bINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bINumPrecisionUint)

		return
	}

	bINumScaleFactor, err := bINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
			"Expected bINumScaleFactor = '%v'\n"+
			"  Actual bINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

		return
	}

	bINumSignValue, err := bINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSignValue, err := bINum.GetSign()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if expectedSignVal != bINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
			"Because  expectedSignVal != bINumSignValue\n"+
			"Expected bINumSignValue = '%v'\n"+
			"  Actual bINumSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bINumSignValue)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"bINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != bINumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}
