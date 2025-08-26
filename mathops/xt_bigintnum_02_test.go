package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntNum_Floor_01(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_01"

	var err error

	originalNumStr := "5.95"

	expectedNumStr := "5"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_02(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_02"

	var err error

	originalNumStr := "5.05"

	expectedNumStr := "5"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_03(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_03"

	var err error

	originalNumStr := "5"

	expectedNumStr := "5"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_04(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_04"

	var err error

	originalNumStr := "-5.05"

	expectedNumStr := "-6"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_05(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_05"

	var err error

	originalNumStr := "2.4"

	expectedNumStr := "2"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_06(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_06"

	var err error

	originalNumStr := "2.9"

	expectedNumStr := "2"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_07(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_07"

	var err error

	originalNumStr := "-2.7"

	expectedNumStr := "-3"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_08(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_08"

	var err error

	originalNumStr := "-2"

	expectedNumStr := "-2"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_09(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_09"

	var err error

	originalNumStr := "0"

	expectedNumStr := "0"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_10(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_10"

	var err error

	originalNumStr := "18972.0000000000001"

	expectedNumStr := "18972"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_11(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_11"

	var err error

	originalNumStr := "-18972.0000000000001"

	expectedNumStr := "-18973"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_12(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_12"

	var err error

	originalNumStr := "0.0000000000001"

	expectedNumStr := "0"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_13(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_13"

	var err error

	originalNumStr := "-0.0000000000001"

	expectedNumStr := "-1"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_14(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_14"

	var err error

	originalNumStr := "-189765342891.0000000000001"

	expectedNumStr := "-189765342892"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_15(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_15"

	var err error

	originalNumStr := "189765342891.0000000000001"

	expectedNumStr := "189765342891"

	expectedPrecision := uint(0)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_Floor_16(t *testing.T) {

	ePrefix := "TestBigIntNum_Floor_15"

	var err error

	originalNumStr := "189765342891,0000000000001"

	expectedNumStr := "189765342891"

	expectedPrecision := uint(0)

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr, &expectedNumSeps)\n"+
			"expectedNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumSeps.String(), err.Error())
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedNumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedBigINumberStr)

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

	if expectedPrecision != expectedBigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal!\n"+
			"Because  expectedPrecision != expectedBigINumPrecisionUint\n"+
			"Expected expectedBigINumPrecisionUint = '%v'\n"+
			"  Actual expectedBigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, expectedBigINumPrecisionUint)

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
			ePrefix, originalNumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumOriginalNumStr \n"+
			"Expected bINumOriginalNumStr = '%v'\n"+
			"  Actual bINumOriginalNumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	floorBINum, err := bINum.Floor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINum, err := bINum.Floor()\n"+
			"bINumOriginalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bINumStr, err.Error())
		return
	}

	err = floorBINum.IsValid("Validating Final floorBINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = floorBINum.IsValid('Validating floorBINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumStr, err := floorBINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumStr, err := floorBINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumSeps, err := floorBINum.GetNumericSeparatorsDto()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, floorBINumStr, err.Error())
		return
	}

	floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"floorBINumPrecisionUint, err := floorBINum.GetPrecisionUint()\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, floorBINumStr, err.Error())
		return
	}

	expectedEqualsFloorBINum, err := expectedBigINum.Equal(floorBINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsFloorBINum, err :=\n"+
			"  expectedBigINum.Equal(floorBINum)\n"+
			"expectedBigINum= '%v'\n"+
			"floorBINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumberStr, floorBINumStr, err.Error())
		return
	}

	if !expectedEqualsFloorBINum {
		t.Errorf("%v\n"+
			"Error: Expected and 'floorBINum' values ARE NOT Equal\n"+
			"Because expectedEqualsBINum = 'false' \n"+
			"Expected floorBINum = '%v'\n"+
			"  Actual floorBINum = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedNumStr != floorBINumStr {
		t.Errorf("%v\n"+
			"Error: Expected and bINumFinal number strings NOT EQUAL!\n"+
			"Because expectedNumStr != floorBINumStr\n"+
			"Expected floorBINumStr = '%v'\n"+
			"  Actual floorBINumStr = '%v'\n\n",
			ePrefix, expectedNumStr, floorBINumStr)

		return
	}

	if expectedPrecision != floorBINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected Precision Values ARE NOT Equal!\n"+
			"Because expectedPrecision != floorBINumPrecisionUint\n"+
			"Expected floorBINumPrecisionUint = '%v'\n"+
			"  Actual floorBINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, floorBINumPrecisionUint)

		return
	}

	if !expectedNumSeps.Equal(floorBINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal\n"+
			"Because expectedNumSeps != floorBINumSeps \n"+
			"Expected floorBINumSeps = '%v'\n"+
			"  Actual floorBINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), floorBINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_01(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_01"

	var err error

	originalNumStr := "-123.45"

	expectedNumStr := "-123.45"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_02(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_02"

	var err error

	originalNumStr := "123.45"

	expectedNumStr := "123.45"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_03(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_03"

	var err error

	originalNumStr := "-123.45"

	expectedNumStr := "(123.45)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_04(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_04"

	var err error

	originalNumStr := "-1234.56"

	expectedNumStr := "-1234.56"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_05(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_05"

	var err error

	originalNumStr := "1234.56"

	expectedNumStr := "1234.56"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_06(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_06"

	var err error

	originalNumStr := "-1234.56"

	expectedNumStr := "(1234.56)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_07(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_07"

	var err error

	originalNumStr := "0"

	expectedNumStr := "0"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_08(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_08"

	var err error

	originalNumStr := "0.000"

	expectedNumStr := "0.000"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_09(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_09"

	var err error

	originalNumStr := "0.000"

	expectedNumStr := "0.000"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_12(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_12"

	var err error

	originalNumStr := "12345"

	expectedNumStr := "12345"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_13(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_13"

	var err error

	originalNumStr := "-12345"

	expectedNumStr := "(12345)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_14(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_14"

	var err error

	originalNumStr := "-12345"

	expectedNumStr := "-12345"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_15(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_15"

	var err error

	originalBInt := big.NewInt(12345)

	precisionUint := uint(8)

	expectedNumStr := "0.00012345"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_16(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_16"

	var err error

	originalBInt := big.NewInt(12345)

	precisionUint := uint(5)

	expectedNumStr := "0.12345"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_17(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_17"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(8)

	expectedNumStr := "-0.00012345"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_18(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_18"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(5)

	expectedNumStr := "-0.12345"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_19(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_19"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(8)

	expectedNumStr := "(0.00012345)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_20(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_20"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(5)

	expectedNumStr := "(0.12345)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_21(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_21"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(5)

	expectedNumStr := "12345"

	mode := ABSOLUTEPURENUMSTRFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_22(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_22"

	var err error

	originalBInt := big.NewInt(12345)

	precisionUint := uint(2)

	expectedNumStr := "12345"

	mode := ABSOLUTEPURENUMSTRFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_23(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_23"

	var err error

	originalBInt := big.NewInt(12345)

	precisionUint := uint(7)

	expectedNumStr := "0012345"

	mode := ABSOLUTEPURENUMSTRFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatNumStr_24(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatNumStr_24"

	var err error

	originalBInt := big.NewInt(12345)

	precisionUint := uint(2)

	expectedNumStr := "12345"

	mode := ABSOLUTEPURENUMSTRFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatNumStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatNumStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_01(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_01"

	var err error

	originalNumStr := "1234"

	expectedNumStr := "1,234"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_02(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_02"

	var err error

	originalNumStr := "123"

	expectedNumStr := "123"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_03(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_03"

	var err error

	originalNumStr := "-1234"

	expectedNumStr := "(1,234)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_04(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_04"

	var err error

	originalNumStr := "-1234"

	expectedNumStr := "-1,234"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_05(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_05"

	var err error

	originalNumStr := "1234.567"

	expectedNumStr := "1,234.567"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_06(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_06"

	var err error

	originalNumStr := "-1234.567"

	expectedNumStr := "-1,234.567"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_07(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_07"

	var err error

	originalNumStr := "-1234.567"

	expectedNumStr := "(1,234.567)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_08(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_08"

	var err error

	originalNumStr := "0"

	expectedNumStr := "0"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_09(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_09"

	var err error

	originalNumStr := "0.0000"

	expectedNumStr := "0.0000"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_10(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_10"

	var err error

	originalNumStr := "0.0000"

	expectedNumStr := "0.0000"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_11(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_11"

	var err error

	originalNumStr := "1234567890.12"

	expectedNumStr := "1,234,567,890.12"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_12(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_12"

	var err error

	originalNumStr := "-1234567890.12"

	expectedNumStr := "-1,234,567,890.12"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_13(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_13"

	var err error

	originalNumStr := "-1234567890.12"

	expectedNumStr := "(1,234,567,890.12)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_14(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_14"

	var err error

	originalNumStr := "1234567890"

	expectedNumStr := "1,234,567,890"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_15(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_15"

	var err error

	originalNumStr := "-1234567890"

	expectedNumStr := "-1,234,567,890"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_16(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_16"

	var err error

	originalBInt := big.NewInt(12345)

	precisionUint := uint(8)

	expectedNumStr := "0.00012345"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_17(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_17"

	var err error

	originalBInt := big.NewInt(12345)

	precisionUint := uint(5)

	expectedNumStr := "0.12345"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_18(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_18"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(8)

	expectedNumStr := "-0.00012345"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_19(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_19"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(5)

	expectedNumStr := "-0.12345"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_20(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_20"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(8)

	expectedNumStr := "(0.00012345)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_21(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_21"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(5)

	expectedNumStr := "(0.12345)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_22(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_22"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(5)

	expectedNumStr := "12345"

	mode := ABSOLUTEPURENUMSTRFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_23(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_23"

	var err error

	originalBInt := big.NewInt(12345)

	precisionUint := uint(0)

	expectedNumStr := "12,345"

	mode := ABSOLUTEPURENUMSTRFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_24(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_24"

	var err error

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(0)

	expectedNumStr := "12,345"

	mode := ABSOLUTEPURENUMSTRFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_25(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_25"

	var err error

	originalBInt := big.NewInt(123)

	expectedNumStr := "123"

	precisionUint := uint(0)

	mode := ABSOLUTEPURENUMSTRFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	return
}

func TestBigIntNum_FormatThousandsStr_26(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatThousandsStr_26"

	var err error

	originalNumStr := "1234"

	expectedNumStr := "1 234"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, err.Error())
		return
	}

	err = bINum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalNumStr= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalNumStr, err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumFormattedNumStr, err := bINum.FormatThousandsStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFormattedNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFormattedNumStr\n"+
			"Expected bINumFormattedNumStr = '%v'\n"+
			"  Actual bINumFormattedNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFormattedNumStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_01(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_01"

	originalNumStr := "1234"

	expectedNumStr := "$1,234"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_02(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_02"

	originalNumStr := "123"

	expectedNumStr := "$123"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_03(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_03"

	originalNumStr := "-1234"

	expectedNumStr := "($1,234)"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_04(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_04"

	originalNumStr := "-1234"

	expectedNumStr := "-$1,234"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_05(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_05"

	originalNumStr := "1234.567"

	expectedNumStr := "$1,234.567"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_06(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_06"

	originalNumStr := "-1234.567"

	expectedNumStr := "-$1,234.567"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, expectedNumSeps.String(), err.Error())
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_07(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_07"

	originalNumStr := "-1234.567"

	expectedNumStr := "($1,234.567)"

	mode := PARENTHESESNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, expectedNumSeps.String(), err.Error())
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_08(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_08"

	originalNumStr := "0"

	expectedNumStr := "$0"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_09(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_09"

	originalNumStr := "0.0000"

	expectedNumStr := "$0.0000"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_10(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_10"

	originalNumStr := "0.0000"

	expectedNumStr := "$0.0000"

	mode := PARENTHESESNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_11(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_11"

	originalNumStr := "1234567890.12"

	expectedNumStr := "$1,234,567,890.12"

	mode := PARENTHESESNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_12(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_12"

	originalNumStr := "-1234567890.12"

	expectedNumStr := "-$1,234,567,890.12"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_13(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_13"

	originalNumStr := "-1234567890.12"

	expectedNumStr := "($1,234,567,890.12)"

	mode := PARENTHESESNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_14(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_14"

	originalNumStr := "1234567890"

	expectedNumStr := "$1,234,567,890"

	mode := PARENTHESESNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_15(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_15"

	originalNumStr := "-1234567890"

	expectedNumStr := "-$1,234,567,890"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_16(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_16"

	originalBInt := big.NewInt(12345)

	precisionUint := uint(8)

	expectedNumStr := "$0.00012345"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_17(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_17"

	originalBInt := big.NewInt(12345)

	precisionUint := uint(5)

	expectedNumStr := "$0.12345"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_18(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_18"

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(8)

	expectedNumStr := "-$0.00012345"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_19(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_19"

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(5)

	expectedNumStr := "-$0.12345"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_20(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_20"

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(8)

	expectedNumStr := "($0.00012345)"

	mode := PARENTHESESNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_21(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_21"

	originalBInt := big.NewInt(-12345)

	precisionUint := uint(5)

	expectedNumStr := "($0.12345)"

	mode := PARENTHESESNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewBigInt(originalBInt, precisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).\n"+
			"  NewBigInt(originalBInt, precisionUint)\n"+
			"originalBInt= '%v'\n"+
			"precisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBInt.Text(10),
			precisionUint,
			err.Error())

		return
	}

	err = bINum.IsValid("Validating bINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum')\n"+
			"originalBInt= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, originalBInt.Text(10), err.Error())
		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_22(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_22"

	originalNumStr := "-1234567890"

	expectedNumStr := "-£1,234,567,890"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	err = bINum.SetCurrencySymbol('\U000000a3')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetCurrencySymbol('U000000a3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Set Currency")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Set Currency')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_23(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_23"

	originalNumStr := "1234567890"

	expectedNumStr := "£1,234,567,890"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	err = bINum.SetCurrencySymbol('\U000000a3')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetCurrencySymbol('U000000a3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Set Currency")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Set Currency')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_24(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_24"

	originalNumStr := "1234567890.12"

	expectedNumStr := "£1,234,567,890.12"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	err = bINum.SetCurrencySymbol('\U000000a3')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetCurrencySymbol('U000000a3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Set Currency")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Set Currency')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_25(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_25"

	originalNumStr := "-1234567890.12"

	expectedNumStr := "-£1,234,567,890.12"

	mode := LEADMINUSNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	err = bINum.SetCurrencySymbol('\U000000a3')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetCurrencySymbol('U000000a3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Set Currency")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Set Currency')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_26(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_26"

	originalNumStr := "-1234567890.12"

	expectedNumStr := "(£1,234,567,890.12)"

	mode := PARENTHESESNEGVALFMTMODE

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	err = bINum.SetCurrencySymbol('\U000000a3')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.SetCurrencySymbol('U000000a3')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bINum.IsValid("Validating bINum After Set Currency")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bINum After Set Currency')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_27(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_27"

	originalNumStr := "-1234567890.12"

	expectedNumStr := "($1,234,567,890.12)"

	mode := PARENTHESESNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_28(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_28"

	originalNumStr := "-1234567890.12"

	expectedNumStr := "-$1,234,567,890.12"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_29(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_29"

	originalNumStr := "1234567890.12"

	expectedNumStr := "$1,234,567,890.12"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr)\n"+
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_30(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_30"

	var err error

	originalNumStr := "1234567890,12"

	expectedNumStr := "€1 234 567 890,12"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, expectedNumSeps.String(), err.Error())
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}

func TestBigIntNum_FormatCurrencyStr_31(t *testing.T) {

	ePrefix := "TestBigIntNum_FormatCurrencyStr_31"

	var err error

	originalNumStr := "-1234567890,12"

	expectedNumStr := "-€1 234 567 890,12"

	mode := LEADMINUSNEGVALFMTMODE

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	var modeStr string

	labelLen := len(NegativeValueFmtModeLabels)

	intModeValue := int(mode)

	if intModeValue < 0 || intModeValue >= labelLen {

		t.Errorf("%v\n"+
			"Error: Test Data is Corrupted!\n"+
			"Because intModeValue < 0 || intModeValue >= labelLen\n"+
			"Actual intModeValue = '%v'\n\n",
			ePrefix, intModeValue)

		return
	}

	modeStr = NegativeValueFmtModeLabels[mode]

	bINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalNumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum, err := new(BigIntNum).NewNumStr(originalNumStr, &expectedNumSeps)\n"+
			"originalNumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumStr, expectedNumSeps.String(), err.Error())
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

	bINumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumStr, err := bINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumStr != bINumStr {
		t.Errorf("%v\n"+
			"Error: Original Number String Values ARE NOT Equal\n"+
			"Because originalNumStr != bINumStr \n"+
			"Expected bINumStr = '%v'\n"+
			"  Actual bINumStr = '%v'\n\n",
			ePrefix, originalNumStr, bINumStr)

		return
	}

	bINumSeps, err := bINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINumFmtCurrencyStr, err := bINum.FormatCurrencyStr(mode)\n"+
			"mode String= '%v'\n"+
			"mode Value= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, modeStr, intModeValue, err.Error())
		return
	}

	if expectedNumStr != bINumFmtCurrencyStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedNumStr != bINumFmtCurrencyStr\n"+
			"Expected bINumFmtCurrencyStr = '%v'\n"+
			"  Actual bINumFmtCurrencyStr = '%v'\n\n",
			ePrefix, expectedNumStr, bINumFmtCurrencyStr)

		return
	}

	if !expectedNumSeps.Equal(bINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected bINumSeps = '%v'\n"+
			"  Actual bINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bINumSeps.String())

		return
	}

	return
}
