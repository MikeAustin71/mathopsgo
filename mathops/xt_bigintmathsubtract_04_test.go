package mathops

import (
	"fmt"
	"testing"
)

func TestBigIntMathSubtract_SubtractNumStr_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStr_01"

	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"subtrahendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			subtrahendStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStr_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStr_02"

	var err error

	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"subtrahendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			subtrahendStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStr_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStr_03"

	var err error

	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"

	expectedBigINumSign := -1

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"subtrahendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			subtrahendStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStr_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStr_04"

	var err error

	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStr(minuendStr, subtrahendStr, usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"subtrahendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			subtrahendStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStr_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStr_05"

	var err error

	// minuend = -289,673849
	minuendStr := "-289,673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579,012"

	// result = 14289.338151
	expectedBigINumStr := "14289,338151"

	expectedBigINumSign := 1

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

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStr(minuendStr, subtrahendStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStr(minuendStr, subtrahendStr, expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"subtrahendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			subtrahendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_01"

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]string, lenSubtrahends)

	subtrahendAry[0] = subtrahend0
	subtrahendAry[1] = subtrahend1
	subtrahendAry[2] = subtrahend2
	subtrahendAry[3] = subtrahend3
	subtrahendAry[4] = subtrahend4
	subtrahendAry[5] = subtrahend5

	result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_02"

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"

	subtrahend0 := "737.21"
	subtrahend1 := "9637591.879546"
	subtrahend2 := "28"
	subtrahend3 := "5284.9765"
	subtrahend4 := "-189291837.12"
	subtrahend5 := "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]string, lenSubtrahends)

	subtrahendAry[0] = subtrahend0
	subtrahendAry[1] = subtrahend1
	subtrahendAry[2] = subtrahend2
	subtrahendAry[3] = subtrahend3
	subtrahendAry[4] = subtrahend4
	subtrahendAry[5] = subtrahend5

	result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_03"

	var err error

	// minuend =   1,718,973,642.1234567
	minuendStr := "1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"

	expectedBigINumSign := 1

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]string, lenSubtrahends)

	subtrahendAry[0] = subtrahend0
	subtrahendAry[1] = subtrahend1
	subtrahendAry[2] = subtrahend2
	subtrahendAry[3] = subtrahend3
	subtrahendAry[4] = subtrahend4
	subtrahendAry[5] = subtrahend5

	result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrArray_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_04"

	var err error

	// minuend =   -1,718,973,642.1234567
	minuendStr := "-1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"

	expectedBigINumSign := -1

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6
	subtrahendAry := make([]string, lenSubtrahends)

	subtrahendAry[0] = subtrahend0
	subtrahendAry[1] = subtrahend1
	subtrahendAry[2] = subtrahend2
	subtrahendAry[3] = subtrahend3
	subtrahendAry[4] = subtrahend4
	subtrahendAry[5] = subtrahend5

	result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrArray_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrArray_05"

	var err error

	// minuend =   -1 718 973 642,1234567
	minuendStr := "-1718973642,1234567"

	subtrahend0 := "-28934682,721"
	subtrahend1 := "424,987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62,94"
	subtrahend4 := "-999999999,99999"
	subtrahend5 := "-9638932,371"

	// Result:   -680 399 527,959121021
	expectedBigINumStr := "-680399527,959121021"

	expectedBigINumSign := -1

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

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]string, lenSubtrahends)

	subtrahendAry[0] = subtrahend0
	subtrahendAry[1] = subtrahend1
	subtrahendAry[2] = subtrahend2
	subtrahendAry[3] = subtrahend3
	subtrahendAry[4] = subtrahend4
	subtrahendAry[5] = subtrahend5

	result, err := new(BigIntMathSubtract).SubtractNumStrArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_01"

	var err error

	// minuend =   100
	minuendStr := "100"

	subtrahendStrs := []string{
		"5",
		"10",
		"30",
		"60.55",
		"-100.1",
		"-5.6",
	}

	expectedStrs := []string{
		"95",
		"90",
		"70",
		"39.45",
		"200.1",
		"105.6",
	}

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedNumStrsAry := len(expectedStrs)

	if lenExpectedNumStrsAry != lenSubtrahends {
		t.Errorf("%v\n"+
			"Error: Test Data is corrupted!\n"+
			"Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
			"Expected Length of subtrahendStrs Array = '%v'\n"+
			"  Actual Length of subtrahendStrs Array = '%v'\n\n",
			ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

		return
	}

	subtrahendAry := make([]string, lenSubtrahends)

	expectedResultsAry := make([]string, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i] = subtrahendStrs[i]

		expectedResultsAry[i] = expectedStrs[i]

	}

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrOutputToArray(minuendNumStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k] != result[k] {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because result[%d] != expectedResultsAry[%d]\n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

			return
		}

	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_02"

	var err error

	// minuend =   5051
	minuendStr := "5051"

	subtrahendStrs := []string{
		"8000",
		"6051.123456",
		"-30871.25",
		"604.55",
		"9100.123",
		"-115.76",
	}

	expectedStrs := []string{
		"-2949",
		"-1000.123456",
		"35922.25",
		"4446.45",
		"-4049.123",
		"5166.76",
	}

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedNumStrsAry := len(expectedStrs)

	if lenExpectedNumStrsAry != lenSubtrahends {
		t.Errorf("%v\n"+
			"Error: Test Data is corrupted!\n"+
			"Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
			"Expected Length of subtrahendStrs Array = '%v'\n"+
			"  Actual Length of subtrahendStrs Array = '%v'\n\n",
			ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

		return
	}

	subtrahendAry := make([]string, lenSubtrahends)

	expectedResultsAry := make([]string, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i] = subtrahendStrs[i]

		expectedResultsAry[i] = expectedStrs[i]

	}

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k] != result[k] {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because result[%d] != expectedResultsAry[%d]\n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

			return
		}

	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_03"

	var err error

	// minuend =   -20051.974578
	minuendStr := "-20051.974578"

	subtrahendStrs := []string{
		"476.543798",
		"6051.123456",
		"-270871.25",
		"15604.5589321",
		"987100.123",
		"-114555.76",
	}

	expectedStrs := []string{
		"-20528.518376",
		"-26103.098034",
		"250819.275422",
		"-35656.5335101",
		"-1007152.097578",
		"94503.785422",
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedNumStrsAry := len(expectedStrs)

	if lenExpectedNumStrsAry != lenSubtrahends {
		t.Errorf("%v\n"+
			"Error: Test Data is corrupted!\n"+
			"Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
			"Expected Length of subtrahendStrs Array = '%v'\n"+
			"  Actual Length of subtrahendStrs Array = '%v'\n\n",
			ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

		return
	}

	subtrahendAry := make([]string, lenSubtrahends)

	expectedResultsAry := make([]string, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i] = subtrahendStrs[i]

		expectedResultsAry[i] = expectedStrs[i]

	}

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k] != result[k] {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because result[%d] != expectedResultsAry[%d]\n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

			return
		}

	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_04"

	var err error

	// minuend =   0
	minuendStr := "0"

	subtrahendStrs := []string{
		"476.543798",
		"6051.123456",
		"-270871.25",
		"15604.5589321",
		"987100.123",
		"-114555.76",
	}

	expectedStrs := []string{
		"-476.543798",
		"-6051.123456",
		"270871.25",
		"-15604.5589321",
		"-987100.123",
		"114555.76",
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedNumStrsAry := len(expectedStrs)

	if lenExpectedNumStrsAry != lenSubtrahends {
		t.Errorf("%v\n"+
			"Error: Test Data is corrupted!\n"+
			"Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
			"Expected Length of subtrahendStrs Array = '%v'\n"+
			"  Actual Length of subtrahendStrs Array = '%v'\n\n",
			ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

		return
	}

	subtrahendAry := make([]string, lenSubtrahends)

	expectedResultsAry := make([]string, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i] = subtrahendStrs[i]

		expectedResultsAry[i] = expectedStrs[i]

	}

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k] != result[k] {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because result[%d] != expectedResultsAry[%d]\n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

			return
		}

	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_05"

	var err error

	// minuend =   0
	minuendStr := "98.2"

	subtrahendStrs := []string{
		"0",
		"0.000",
		"0",
		"0",
		"0.00000",
		"0.0",
	}

	expectedStrs := []string{
		"98.2",
		"98.200",
		"98.2",
		"98.2",
		"98.20000",
		"98.2",
	}

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedNumStrsAry := len(expectedStrs)

	if lenExpectedNumStrsAry != lenSubtrahends {
		t.Errorf("%v\n"+
			"Error: Test Data is corrupted!\n"+
			"Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
			"Expected Length of subtrahendStrs Array = '%v'\n"+
			"  Actual Length of subtrahendStrs Array = '%v'\n\n",
			ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

		return
	}

	subtrahendAry := make([]string, lenSubtrahends)

	expectedResultsAry := make([]string, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i] = subtrahendStrs[i]

		expectedResultsAry[i] = expectedStrs[i]

	}

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k] != result[k] {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because result[%d] != expectedResultsAry[%d]\n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

			return
		}

	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrOutputToArray_06(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrOutputToArray_06"

	var err error

	// minuend =   5051
	minuendStr := "5051"

	subtrahendStrs := []string{
		"8000",
		"6051,123456",
		"-30871,25",
		"604,55",
		"9100,123",
		"-115,76",
	}

	expectedStrs := []string{
		"-2949",
		"-1000,123456",
		"35922,25",
		"4446,45",
		"-4049,123",
		"5166,76",
	}

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

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedNumStrsAry := len(expectedStrs)

	if lenExpectedNumStrsAry != lenSubtrahends {
		t.Errorf("%v\n"+
			"Error: Test Data is corrupted!\n"+
			"Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
			"Expected Length of subtrahendStrs Array = '%v'\n"+
			"  Actual Length of subtrahendStrs Array = '%v'\n\n",
			ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

		return
	}

	subtrahendAry := make([]string, lenSubtrahends)

	expectedResultsAry := make([]string, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i] = subtrahendStrs[i]

		expectedResultsAry[i] = expectedStrs[i]

	}

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrOutputToArray(minuendStr, subtrahendAry, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrOutputToArray(minuendStr, subtrahendAry[...], expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			expectedNumSeps.String(),
			err.Error())

		return
	}

	for k := 0; k < lenSubtrahends; k++ {

		if expectedResultsAry[k] != result[k] {

			t.Errorf("%v\n"+
				"Error: Unexpected Result!\n"+
				"Because result[%d] != expectedResultsAry[%d]\n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, k, k, expectedResultsAry[k], k, result[k])

			return
		}

	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrSeries_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrSeries_01"

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
		expectedNumSeps,
		minuendStr,
		subtrahend0,
		subtrahend1,
		subtrahend2,
		subtrahend3,
		subtrahend4,
		subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrSeries(\n"+
			"  expectedNumSeps, minuendStr, subtrahend0, subtrahend1,\n"+
			"   subtrahend2, subtrahend3, subtrahend4, subtrahend5)\n"+
			"expectedNumSeps= '%v'\n"+
			"minuendStr= '%v'\n"+
			"subtrahend0= '%v'\n"+
			"subtrahend1= '%v'\n"+
			"subtrahend2= '%v'\n"+
			"subtrahend3= '%v'\n"+
			"subtrahend4= '%v'\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			minuendStr,
			subtrahend0,
			subtrahend1,
			subtrahend2,
			subtrahend3,
			subtrahend4,
			subtrahend5,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrSeries_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrSeries_02"

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"

	subtrahend0 := "737.21"
	subtrahend1 := "9637591.879546"
	subtrahend2 := "28"
	subtrahend3 := "5284.9765"
	subtrahend4 := "-189291837.12"
	subtrahend5 := "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
		expectedNumSeps,
		minuendStr,
		subtrahend0,
		subtrahend1,
		subtrahend2,
		subtrahend3,
		subtrahend4,
		subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrSeries(\n"+
			"  expectedNumSeps, minuendStr, subtrahend0, subtrahend1,\n"+
			"   subtrahend2, subtrahend3, subtrahend4, subtrahend5)\n"+
			"expectedNumSeps= '%v'\n"+
			"minuendStr= '%v'\n"+
			"subtrahend0= '%v'\n"+
			"subtrahend1= '%v'\n"+
			"subtrahend2= '%v'\n"+
			"subtrahend3= '%v'\n"+
			"subtrahend4= '%v'\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			minuendStr,
			subtrahend0,
			subtrahend1,
			subtrahend2,
			subtrahend3,
			subtrahend4,
			subtrahend5,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrSeries_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrSeries_03"

	var err error

	// minuend =   1,718,973,642.1234567
	minuendStr := "1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
		expectedNumSeps,
		minuendStr,
		subtrahend0,
		subtrahend1,
		subtrahend2,
		subtrahend3,
		subtrahend4,
		subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrSeries(\n"+
			"  expectedNumSeps, minuendStr, subtrahend0, subtrahend1,\n"+
			"   subtrahend2, subtrahend3, subtrahend4, subtrahend5)\n"+
			"expectedNumSeps= '%v'\n"+
			"minuendStr= '%v'\n"+
			"subtrahend0= '%v'\n"+
			"subtrahend1= '%v'\n"+
			"subtrahend2= '%v'\n"+
			"subtrahend3= '%v'\n"+
			"subtrahend4= '%v'\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			minuendStr,
			subtrahend0,
			subtrahend1,
			subtrahend2,
			subtrahend3,
			subtrahend4,
			subtrahend5,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrSeries_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrSeries_04"

	var err error

	// minuend =   -1,718,973,642.1234567
	minuendStr := "-1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"

	expectedBigINumSign := -1

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	err = expectedNumSeps.IsValid("Validating expectedNumSeps")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
		expectedNumSeps,
		minuendStr,
		subtrahend0,
		subtrahend1,
		subtrahend2,
		subtrahend3,
		subtrahend4,
		subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrSeries(\n"+
			"  expectedNumSeps, minuendStr, subtrahend0, subtrahend1,\n"+
			"   subtrahend2, subtrahend3, subtrahend4, subtrahend5)\n"+
			"expectedNumSeps= '%v'\n"+
			"minuendStr= '%v'\n"+
			"subtrahend0= '%v'\n"+
			"subtrahend1= '%v'\n"+
			"subtrahend2= '%v'\n"+
			"subtrahend3= '%v'\n"+
			"subtrahend4= '%v'\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			minuendStr,
			subtrahend0,
			subtrahend1,
			subtrahend2,
			subtrahend3,
			subtrahend4,
			subtrahend5,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrSeries_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrSeries_05"

	var err error

	// minuend = 7328941,123456
	minuendStr := "7328941,123456"

	subtrahend0 := "123,894000"
	subtrahend1 := "67,1"
	subtrahend2 := "93,0"
	subtrahend3 := "-124498,67158"
	subtrahend4 := "647129,57"
	subtrahend5 := "28"

	// result = 6805998,231036
	expectedBigINumStr := "6805998,231036"

	expectedBigINumSign := 1

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

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrSeries(
		expectedNumSeps,
		minuendStr,
		subtrahend0,
		subtrahend1,
		subtrahend2,
		subtrahend3,
		subtrahend4,
		subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrSeries(\n"+
			"  expectedNumSeps, minuendStr, subtrahend0, subtrahend1,\n"+
			"   subtrahend2, subtrahend3, subtrahend4, subtrahend5)\n"+
			"expectedNumSeps= '%v'\n"+
			"minuendStr= '%v'\n"+
			"subtrahend0= '%v'\n"+
			"subtrahend1= '%v'\n"+
			"subtrahend2= '%v'\n"+
			"subtrahend3= '%v'\n"+
			"subtrahend4= '%v'\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedNumSeps.String(),
			minuendStr,
			subtrahend0,
			subtrahend1,
			subtrahend2,
			subtrahend3,
			subtrahend4,
			subtrahend5,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDto_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDto_01"

	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	nDtoSubtrahend, err := new(NumStrDto).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahend, err := new(NumStrDto).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = nDtoSubtrahend.IsValid("Validating nDtoSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoSubtrahend.IsValid('Validating nDtoSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != nDtoSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != nDtoSubtrahendNumStr \n"+
			"Expected nDtoSubtrahendNumStr = '%v'\n"+
			"  Actual nDtoSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, nDtoSubtrahendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)\n"+
			"nDtoMinuend= '%v'\n"+
			"nDtoSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
			nDtoSubtrahendNumStr,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDto_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDto_02"

	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	nDtoSubtrahend, err := new(NumStrDto).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahend, err := new(NumStrDto).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = nDtoSubtrahend.IsValid("Validating nDtoSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoSubtrahend.IsValid('Validating nDtoSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != nDtoSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != nDtoSubtrahendNumStr \n"+
			"Expected nDtoSubtrahendNumStr = '%v'\n"+
			"  Actual nDtoSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, nDtoSubtrahendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)\n"+
			"nDtoMinuend= '%v'\n"+
			"nDtoSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
			nDtoSubtrahendNumStr,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDto_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDto_03"

	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"

	expectedBigINumSign := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	nDtoSubtrahend, err := new(NumStrDto).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahend, err := new(NumStrDto).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = nDtoSubtrahend.IsValid("Validating nDtoSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoSubtrahend.IsValid('Validating nDtoSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != nDtoSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != nDtoSubtrahendNumStr \n"+
			"Expected nDtoSubtrahendNumStr = '%v'\n"+
			"  Actual nDtoSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, nDtoSubtrahendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)\n"+
			"nDtoMinuend= '%v'\n"+
			"nDtoSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
			nDtoSubtrahendNumStr,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDto_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDto_04"

	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	nDtoSubtrahend, err := new(NumStrDto).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahend, err := new(NumStrDto).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = nDtoSubtrahend.IsValid("Validating nDtoSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoSubtrahend.IsValid('Validating nDtoSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != nDtoSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != nDtoSubtrahendNumStr \n"+
			"Expected nDtoSubtrahendNumStr = '%v'\n"+
			"  Actual nDtoSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, nDtoSubtrahendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)\n"+
			"nDtoMinuend= '%v'\n"+
			"nDtoSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
			nDtoSubtrahendNumStr,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDto_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDto_05"

	var err error

	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99,999"

	expectedBigINumSign := 1

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

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStrWithNumSeps(minuendStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(minuendStr, &usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	nDtoSubtrahend, err := new(NumStrDto).NewNumStrWithNumSeps(subtrahendStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahend, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(minuendStr, &usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			usaNumSeps.String(),
			err.Error())

		return
	}

	err = nDtoSubtrahend.IsValid("Validating nDtoSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoSubtrahend.IsValid('Validating nDtoSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != nDtoSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != nDtoSubtrahendNumStr \n"+
			"Expected nDtoSubtrahendNumStr = '%v'\n"+
			"  Actual nDtoSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, nDtoSubtrahendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDto(nDtoMinuend, nDtoSubtrahend)\n"+
			"nDtoMinuend= '%v'\n"+
			"nDtoSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
			nDtoSubtrahendNumStr,
			err.Error())

		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoArray_01"

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoArray_02"

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"

	subtrahend0 := "737.21"
	subtrahend1 := "9637591.879546"
	subtrahend2 := "28"
	subtrahend3 := "5284.9765"
	subtrahend4 := "-189291837.12"
	subtrahend5 := "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoArray_03"

	var err error

	// minuend =   1,718,973,642.1234567
	minuendStr := "1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoArray_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoArray_04"

	var err error

	// minuend =   -1,718,973,642.1234567
	minuendStr := "-1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"

	expectedBigINumSign := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigINumStr,
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
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoArray_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoArray_05"

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998,231036
	expectedBigINumStr := "6805998,231036"

	expectedBigINumSign := 1

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

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != nDtoMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != nDtoMinuendNumStr\n"+
			"Expected nDtoMinuendNumStr = '%v'\n"+
			"  Actual nDtoMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, nDtoMinuendNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(
		expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).\n"+
			"  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
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

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).\n"+
			"  SubtractNumStrDtoArray(nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, err.Error())
		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_01"

	var err error

	// minuend =   100
	minuendStr := "100"

	subtrahendStrs := []string{
		"5",
		"10",
		"30",
		"60.55",
		"-100.1",
		"-5.6",
	}

	expectedStrs := []string{
		"95",
		"90",
		"70",
		"39.45",
		"200.1",
		"105.6",
	}

	minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendNumStrDto.IsValid("Validating minuendNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendNumStrDto.IsValid('Validating minuendNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedStrs := len(expectedStrs)

	if lenSubtrahends != lenExpectedStrs {
		t.Errorf("%v\n"+
			"Error: Test is Corrupted!\n"+
			"Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
			"Expected lenExpectedStrs = '%v'\n"+
			"  Actual lenExpectedStrs = '%v'\n\n",
			ePrefix, lenSubtrahends, lenExpectedStrs)

		return
	}

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	expectedResultsAry := make([]NumStrDto, lenSubtrahends)

	var expectedResultsNumStr string

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(NumStrDto).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"subtrahendAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = subtrahendAry[%d].IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		if expectedStrs[i] != expectedResultsNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != expectedResultsNumStr \n"+
				"Expected expectedResultsNumStr = '%v'\n"+
				"  Actual expectedResultsNumStr = '%v'\n\n",
				ePrefix, i, expectedStrs[i], expectedResultsNumStr)

			return
		}

	} // End of Loop

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
			"  minuendNumStrDto, subtrahendAry[...])\n"+
			"minuendNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendNumStrDtoNumStr,
			err.Error())

		return
	}

	var expectedResultEqualsResult bool

	var expectedResultNumStr, resultNumStr string

	for k := 0; k < lenSubtrahends; k++ {

		resultNumStr, err = result[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultEqualsResult = expectedResultsAry[k].Equal(result[k])

		if !expectedResultEqualsResult {
			t.Errorf("%v\n"+
				"Error: Expected and 'result' values NOT Equal\n"+
				"Because expectedResultEqualsResult = 'false' \n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, expectedResultNumStr, k, resultNumStr)

			return
		}

		if expectedResultNumStr != resultNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedResultNumStr != resultNumStr \n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, expectedResultNumStr, resultNumStr)

			return
		}

	} // End of loop

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_02"

	var err error

	// minuend =   5051
	minuendStr := "5051"

	subtrahendStrs := []string{
		"8000",
		"6051.123456",
		"-30871.25",
		"604.55",
		"9100.123",
		"-115.76",
	}

	expectedStrs := []string{
		"-2949",
		"-1000.123456",
		"35922.25",
		"4446.45",
		"-4049.123",
		"5166.76",
	}

	minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendNumStrDto.IsValid("Validating minuendNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendNumStrDto.IsValid('Validating minuendNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedStrs := len(expectedStrs)

	if lenSubtrahends != lenExpectedStrs {
		t.Errorf("%v\n"+
			"Error: Test is Corrupted!\n"+
			"Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
			"Expected lenExpectedStrs = '%v'\n"+
			"  Actual lenExpectedStrs = '%v'\n\n",
			ePrefix, lenSubtrahends, lenExpectedStrs)

		return
	}

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	expectedResultsAry := make([]NumStrDto, lenSubtrahends)

	var expectedResultsNumStr string

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(NumStrDto).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"subtrahendAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = subtrahendAry[%d].IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		if expectedStrs[i] != expectedResultsNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != expectedResultsNumStr \n"+
				"Expected expectedResultsNumStr = '%v'\n"+
				"  Actual expectedResultsNumStr = '%v'\n\n",
				ePrefix, i, expectedStrs[i], expectedResultsNumStr)

			return
		}

	} // End of Loop

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
			"  minuendNumStrDto, subtrahendAry[...])\n"+
			"minuendNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendNumStrDtoNumStr,
			err.Error())

		return
	}

	var expectedResultEqualsResult bool

	var expectedResultNumStr, resultNumStr string

	for k := 0; k < lenSubtrahends; k++ {

		resultNumStr, err = result[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultEqualsResult = expectedResultsAry[k].Equal(result[k])

		if !expectedResultEqualsResult {
			t.Errorf("%v\n"+
				"Error: Expected and 'result' values NOT Equal\n"+
				"Because expectedResultEqualsResult = 'false' \n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, expectedResultNumStr, k, resultNumStr)

			return
		}

		if expectedResultNumStr != resultNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedResultNumStr != resultNumStr \n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, expectedResultNumStr, resultNumStr)

			return
		}

	} // End of loop

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_03"

	var err error

	// minuend =   -20051.974578
	minuendStr := "-20051.974578"

	subtrahendStrs := []string{
		"476.543798",
		"6051.123456",
		"-270871.25",
		"15604.5589321",
		"987100.123",
		"-114555.76",
	}

	expectedStrs := []string{
		"-20528.518376",
		"-26103.098034",
		"250819.275422",
		"-35656.5335101",
		"-1007152.097578",
		"94503.785422",
	}

	minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendNumStrDto.IsValid("Validating minuendNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendNumStrDto.IsValid('Validating minuendNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedStrs := len(expectedStrs)

	if lenSubtrahends != lenExpectedStrs {
		t.Errorf("%v\n"+
			"Error: Test is Corrupted!\n"+
			"Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
			"Expected lenExpectedStrs = '%v'\n"+
			"  Actual lenExpectedStrs = '%v'\n\n",
			ePrefix, lenSubtrahends, lenExpectedStrs)

		return
	}

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	expectedResultsAry := make([]NumStrDto, lenSubtrahends)

	var expectedResultsNumStr string

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(NumStrDto).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"subtrahendAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = subtrahendAry[%d].IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		if expectedStrs[i] != expectedResultsNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != expectedResultsNumStr \n"+
				"Expected expectedResultsNumStr = '%v'\n"+
				"  Actual expectedResultsNumStr = '%v'\n\n",
				ePrefix, i, expectedStrs[i], expectedResultsNumStr)

			return
		}

	} // End of Loop

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
			"  minuendNumStrDto, subtrahendAry[...])\n"+
			"minuendNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendNumStrDtoNumStr,
			err.Error())

		return
	}

	var expectedResultEqualsResult bool

	var expectedResultNumStr, resultNumStr string

	for k := 0; k < lenSubtrahends; k++ {

		resultNumStr, err = result[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultEqualsResult = expectedResultsAry[k].Equal(result[k])

		if !expectedResultEqualsResult {
			t.Errorf("%v\n"+
				"Error: Expected and 'result' values NOT Equal\n"+
				"Because expectedResultEqualsResult = 'false' \n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, expectedResultNumStr, k, resultNumStr)

			return
		}

		if expectedResultNumStr != resultNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedResultNumStr != resultNumStr \n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, expectedResultNumStr, resultNumStr)

			return
		}

	} // End of loop

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_04"

	var err error

	// minuend =   0
	minuendStr := "0"

	subtrahendStrs := []string{
		"476.543798",
		"6051.123456",
		"-270871.25",
		"15604.5589321",
		"987100.123",
		"-114555.76",
	}

	expectedStrs := []string{
		"-476.543798",
		"-6051.123456",
		"270871.25",
		"-15604.5589321",
		"-987100.123",
		"114555.76",
	}

	minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendNumStrDto.IsValid("Validating minuendNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendNumStrDto.IsValid('Validating minuendNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedStrs := len(expectedStrs)

	if lenSubtrahends != lenExpectedStrs {
		t.Errorf("%v\n"+
			"Error: Test is Corrupted!\n"+
			"Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
			"Expected lenExpectedStrs = '%v'\n"+
			"  Actual lenExpectedStrs = '%v'\n\n",
			ePrefix, lenSubtrahends, lenExpectedStrs)

		return
	}

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	expectedResultsAry := make([]NumStrDto, lenSubtrahends)

	var expectedResultsNumStr string

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(NumStrDto).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"subtrahendAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = subtrahendAry[%d].IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		if expectedStrs[i] != expectedResultsNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != expectedResultsNumStr \n"+
				"Expected expectedResultsNumStr = '%v'\n"+
				"  Actual expectedResultsNumStr = '%v'\n\n",
				ePrefix, i, expectedStrs[i], expectedResultsNumStr)

			return
		}

	} // End of Loop

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
			"  minuendNumStrDto, subtrahendAry[...])\n"+
			"minuendNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendNumStrDtoNumStr,
			err.Error())

		return
	}

	var expectedResultEqualsResult bool

	var expectedResultNumStr, resultNumStr string

	for k := 0; k < lenSubtrahends; k++ {

		resultNumStr, err = result[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultEqualsResult = expectedResultsAry[k].Equal(result[k])

		if !expectedResultEqualsResult {
			t.Errorf("%v\n"+
				"Error: Expected and 'result' values NOT Equal\n"+
				"Because expectedResultEqualsResult = 'false' \n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, expectedResultNumStr, k, resultNumStr)

			return
		}

		if expectedResultNumStr != resultNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedResultNumStr != resultNumStr \n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, expectedResultNumStr, resultNumStr)

			return
		}

	} // End of loop

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_05"

	var err error

	// minuend =   0
	minuendStr := "98.2"

	subtrahendStrs := []string{
		"0",
		"0.000",
		"0",
		"0",
		"0.00000",
		"0.0",
	}

	expectedStrs := []string{
		"98.2",
		"98.200",
		"98.2",
		"98.2",
		"98.20000",
		"98.2",
	}

	minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendNumStrDto.IsValid("Validating minuendNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendNumStrDto.IsValid('Validating minuendNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedStrs := len(expectedStrs)

	if lenSubtrahends != lenExpectedStrs {
		t.Errorf("%v\n"+
			"Error: Test is Corrupted!\n"+
			"Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
			"Expected lenExpectedStrs = '%v'\n"+
			"  Actual lenExpectedStrs = '%v'\n\n",
			ePrefix, lenSubtrahends, lenExpectedStrs)

		return
	}

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	expectedResultsAry := make([]NumStrDto, lenSubtrahends)

	var expectedResultsNumStr string

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(NumStrDto).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"subtrahendAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = subtrahendAry[%d].IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = expectedResultsAry[%d].IsValid(ePrefix)\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		if expectedStrs[i] != expectedResultsNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != expectedResultsNumStr \n"+
				"Expected expectedResultsNumStr = '%v'\n"+
				"  Actual expectedResultsNumStr = '%v'\n\n",
				ePrefix, i, expectedStrs[i], expectedResultsNumStr)

			return
		}

	} // End of Loop

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
			"  minuendNumStrDto, subtrahendAry[...])\n"+
			"minuendNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendNumStrDtoNumStr,
			err.Error())

		return
	}

	var expectedResultEqualsResult bool

	var expectedResultNumStr, resultNumStr string

	for k := 0; k < lenSubtrahends; k++ {

		resultNumStr, err = result[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultEqualsResult = expectedResultsAry[k].Equal(result[k])

		if !expectedResultEqualsResult {
			t.Errorf("%v\n"+
				"Error: Expected and 'result' values NOT Equal\n"+
				"Because expectedResultEqualsResult = 'false' \n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, expectedResultNumStr, k, resultNumStr)

			return
		}

		if expectedResultNumStr != resultNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedResultNumStr != resultNumStr \n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, expectedResultNumStr, resultNumStr)

			return
		}

	} // End of loop

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_06(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoOutputToArray_06"

	var err error

	// minuend =   100
	minuendStr := "100"

	subtrahendStrs := []string{
		"5",
		"10",
		"30",
		"60.55",
		"-100.1",
		"-5.6",
	}

	expectedStrs := []string{
		"95",
		"90",
		"70",
		"39,45",
		"200,1",
		"105,6",
	}

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

	minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDto, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendNumStrDto.IsValid("Validating minuendNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendNumStrDto.IsValid('Validating minuendNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendNumStrDtoNumStr, err := minuendNumStrDto.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	lenExpectedStrs := len(expectedStrs)

	if lenSubtrahends != lenExpectedStrs {
		t.Errorf("%v\n"+
			"Error: Test is Corrupted!\n"+
			"Because Lengths of subtrahendStrs and expectedStrs ARE NOT EQUAL!\n"+
			"Expected lenExpectedStrs = '%v'\n"+
			"  Actual lenExpectedStrs = '%v'\n\n",
			ePrefix, lenSubtrahends, lenExpectedStrs)

		return
	}

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	expectedResultsAry := make([]NumStrDto, lenSubtrahends)

	var expectedResultsNumStr string

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(NumStrDto).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"subtrahendAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = subtrahendAry[%d].IsValid(ePrefix)\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		expectedResultsAry[i], err = new(NumStrDto).NewNumStrWithNumSeps(expectedStrs[i], &expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsAry[%d], err = new(NumStrDto).\n"+
				"  NewNumStrWithNumSeps(expectedStrs[%d], &expectedNumSeps)\n"+
				"expectedStrs[%d]= '%v'\n"+
				"expectedNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, i, i, expectedStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err := expectedResultsAry[%d].IsValid()\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, expectedStrs[i], err.Error())

			return
		}

		expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		if expectedStrs[i] != expectedResultsNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != expectedResultsNumStr \n"+
				"Expected expectedResultsNumStr = '%v'\n"+
				"  Actual expectedResultsNumStr = '%v'\n\n",
				ePrefix, i, expectedStrs[i], expectedResultsNumStr)

			return
		}

	} // End of Loop

	result, err :=
		new(BigIntMathSubtract).SubtractNumStrDtoOutputToArray(minuendNumStrDto, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractBigIntNumOutputToArray(\n"+
			"  minuendNumStrDto, subtrahendAry[...])\n"+
			"minuendNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendNumStrDtoNumStr,
			err.Error())

		return
	}

	var expectedResultEqualsResult bool

	var expectedResultNumStr, resultNumStr string

	var resultNumSeps NumericSeparatorDto

	for k := 0; k < lenSubtrahends; k++ {

		err = result[k].SetNumericSeparatorsDto(expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = result[%d].SetNumericSeparatorsDto(expectedNumSeps)\n"+
				"result[%d]= '%v'\n"+
				"expectedNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, k, k, result[k], expectedNumSeps.String(), err.Error())
			return
		}

		err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = result[%d].IsValid('Validating result[%d]')\n"+
				"result[%d]= '%v'\n"+
				"Validation Error= '%v'\n\n",
				ePrefix, k, k, k, result[k], err.Error())
			return
		}

		resultNumStr, err = result[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumStr, err = result[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		resultNumSeps, err = result[k].GetNumericSeparatorsDto()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, k, err.Error())
			return
		}

		expectedResultEqualsResult = expectedResultsAry[k].Equal(result[k])

		if !expectedResultEqualsResult {
			t.Errorf("%v\n"+
				"Error: Expected and 'result' values NOT Equal\n"+
				"Because expectedResultEqualsResult = 'false' \n"+
				"Expected result[%d] = '%v'\n"+
				"  Actual result[%d] = '%v'\n\n",
				ePrefix, k, expectedResultNumStr, k, resultNumStr)

			return
		}

		if expectedResultNumStr != resultNumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedResultNumStr != resultNumStr \n"+
				"Expected resultNumStr = '%v'\n"+
				"  Actual resultNumStr = '%v'\n\n",
				ePrefix, expectedResultNumStr, resultNumStr)

			return
		}

		if !expectedNumSeps.Equal(resultNumSeps) {
			t.Errorf("%v\n"+
				"Error: Number Sign Values NOT Equal\n"+
				"Because expectedNumSeps != resultNumSeps \n"+
				"Expected resultNumSeps[%d] = '%v'\n"+
				"  Actual resultNumSeps[%d] = '%v'\n\n",
				ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

			return
		}

	} // End of loop

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoSeries_01"

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
		nDtoMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(\n"+
			" nDtoMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"nDtoMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
			subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoSeries_02"

	var err error

	// minuend = -18,973,642.1234567
	minuendStr := "-18973642.1234567"

	subtrahend0 := "737.21"
	subtrahend1 := "9637591.879546"
	subtrahend2 := "28"
	subtrahend3 := "5284.9765"
	subtrahend4 := "-189291837.12"
	subtrahend5 := "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
		nDtoMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(\n"+
			" nDtoMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"nDtoMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
			subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoSeries_03"

	var err error

	// minuend =   1,718,973,642.1234567
	minuendStr := "1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
		nDtoMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(\n"+
			" nDtoMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"nDtoMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
			subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoSeries_04"

	var err error

	// minuend =   -1,718,973,642.1234567
	minuendStr := "-1718973642.1234567"

	subtrahend0 := "-28934682.721"
	subtrahend1 := "424.987654321"
	subtrahend2 := "-987"
	subtrahend3 := "62.94"
	subtrahend4 := "-999999999.99999"
	subtrahend5 := "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"

	expectedBigINumSign := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
		nDtoMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(\n"+
			" nDtoMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"nDtoMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
			subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_SubtractNumStrDtoSeries_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractNumStrDtoSeries_05"

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0 := "123.894000"
	subtrahend1 := "67.1"
	subtrahend2 := "93.0"
	subtrahend3 := "-124498.67158"
	subtrahend4 := "647129.57"
	subtrahend5 := "28"

	// result = 6805998,231036
	expectedBigINumStr := "6805998,231036"

	expectedBigINumSign := 1

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

	nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(NumStrDto).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuend, err := new(NumStrDto).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid('Validating nDtoMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	if expectedBigINumStr != expectedBigINumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBigINumberStr \n"+
			"Expected expectedBigINumberStr = '%v'\n"+
			"  Actual expectedBigINumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBigINumberStr)

		return
	}

	expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := 6

	subtrahendAry := make([]NumStrDto, lenSubtrahends)

	subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[0], err = new(NumStrDto).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[1], err = new(NumStrDto).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[2], err = new(NumStrDto).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[3], err = new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[4], err = new(NumStrDto).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(NumStrDto).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(
		nDtoMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractNumStrDtoSeries(\n"+
			" nDtoMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"nDtoMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
			subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
			err.Error())

		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigInt, err := result.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigInt, err := result.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := result.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := result.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(result)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultNumStr)

		return
	}

	if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

		return
	}

	if expectedBigINumSign != resultSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separators Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

		return
	}

	return
}
