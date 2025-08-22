package mathops

import (
	"testing"
)

func TestBigIntMathSubtract_SubtractPair_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_01"

	// minuend = 123.32
	minuendStr := "123.32"

	minuendPrecision := uint(2)

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	subtrahendPrecision := uint(3)

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_SubtractPair_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_02"

	// minuend = 949321.6712
	minuendStr := "949321.6712"

	minuendPrecision := uint(4)

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	subtrahendPrecision := uint(2)

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_SubtractPair_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_03"

	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	minuendPrecision := uint(8)

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	subtrahendPrecision := uint(3)

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"

	expectedBigINumSign := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_SubtractPair_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_04"

	// minuend = -289.673849
	minuendStr := "-289.673849"

	minuendPrecision := uint(6)

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	subtrahendPrecision := uint(3)

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_SubtractPair_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_05"

	// minuend = 0
	minuendStr := "0"

	minuendPrecision := uint(0)

	// subtrahend = 0
	subtrahendStr := "0"

	subtrahendPrecision := uint(0)

	// result = 0
	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_SubtractPair_06(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_06"

	// minuend = 270.1
	minuendStr := "270.1"

	minuendPrecision := uint(1)

	// subtrahend = 0
	subtrahendStr := "0"

	subtrahendPrecision := uint(0)

	// result = 270.1
	expectedBigINumStr := "270.1"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_SubtractPair_07(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_07"

	// minuend = 0
	minuendStr := "0"

	minuendPrecision := uint(0)

	// subtrahend = 270.1
	subtrahendStr := "270.1"

	subtrahendPrecision := uint(1)

	// result = -270.1
	expectedBigINumStr := "-270.1"

	expectedBigINumSign := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_SubtractPair_08(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_08"

	// minuend = 0
	minuendStr := "0"

	minuendPrecision := uint(0)

	// subtrahend = -270.1
	subtrahendStr := "-270.1"

	subtrahendPrecision := uint(1)

	// result = 270.1
	expectedBigINumStr := "270.1"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_SubtractPair_09(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_09"

	// minuend = 2.5
	minuendStr := "2.5"

	minuendPrecision := uint(1)

	// subtrahend = 2.5
	subtrahendStr := "2.5"

	subtrahendPrecision := uint(1)

	// result = 0.0
	expectedBigINumStr := "0.0"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_SubtractPair_10(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractPair_10"

	var err error

	// minuend = 123.32
	minuendStr := "123.32"

	minuendPrecision := uint(2)

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	subtrahendPrecision := uint(3)

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

	maxPrecision := minuendPrecision

	if subtrahendPrecision > maxPrecision {
		maxPrecision = subtrahendPrecision
	}

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = expectedBigINum.IsValid("Validating expectedBigINum")

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

	bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair, err := new(BigIntPair).NewBigIntNum(minuendBiNum, subtrahendBiNum)\n"+
			"minuendBiNum= '%v'\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBiNumStr,
			subtrahendBiNumStr,
			err.Error())

		return
	}

	err = bPair.IsValid("Validating bPair")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.IsValid('Validating bPair')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair1NumStr, err := bPair.Big1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair1NumStr, err := bPair.Big1.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPair2NumStr, err := bPair.Big2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPair2NumStr, err := bPair.Big2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bPair.MakePrecisionsEqual()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig1PrecisionUint, err := bPair.Big1.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig1PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig1PrecisionUint\n"+
			"Expected bPairBig1PrecisionUint = '%v'\n"+
			"  Actual bPairBig1PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig1PrecisionUint)

		return
	}

	bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bPairBig2PrecisionUint, err := bPair.Big2.GetPrecisionUint()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if maxPrecision != bPairBig2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because maxPrecision != bPairBig2PrecisionUint\n"+
			"Expected bPairBig2PrecisionUint = '%v'\n"+
			"  Actual bPairBig2PrecisionUint = '%v'\n\n",
			ePrefix, maxPrecision, bPairBig2PrecisionUint)

		return
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractPair(bPair)\n"+
			"bPair.Big1= '%v'\n"+
			"bPair.Big2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bPair1NumStr,
			bPair2NumStr,
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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_01"

	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_02"

	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643,4612"

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_03"

	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"

	expectedBigINumSign := -1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_04"

	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_05"

	minuendStr := "5"

	subtrahendStr := "5"

	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_06(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_06"

	minuendStr := "-5"

	subtrahendStr := "5"

	expectedBigINumStr := "-10"

	expectedBigINumSign := -1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_07(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_07"

	minuendStr := "0"

	subtrahendStr := "0"

	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_08(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_08"

	minuendStr := "50.0"

	subtrahendStr := "2.60134"

	expectedBigINumStr := "47.39866"

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_09(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_09"

	minuendStr := "50.0"

	subtrahendStr := "50.0000"

	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_10(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_10"

	minuendStr := "0.0"

	subtrahendStr := "0.0000"

	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	var expectedNumSeps = new(NumericSeparatorDto).NewUSADefaults()

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_BigIntSubtract_11(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_BigIntSubtract_11"

	var err error

	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643,4612"

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

	minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNum, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendBiNum.IsValid("Validating minuendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendBiNum.IsValid('Validating minuendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendBiNumStr, err := minuendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBiNumStr, err := minuendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != minuendBiNumStr\n"+
			"Expected minuendBiNumStr = '%v'\n"+
			"  Actual minuendBiNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendBiNumStr)

		return
	}

	minuendPrecisionBigInt, err := minuendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendPrecisionBigInt, err := \n"+
			"  minuendBiNum.GetPrecisionBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	minuendBigInt, err := minuendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendBigInt, err := \n"+
			"   minuendBiNum.GetBigInt()\n"+
			"minuendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendBiNumStr, err.Error())
		return
	}

	subtrahendBiNum, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNum, err := \n"+
			" new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = subtrahendBiNum.IsValid("Validating subtrahendBiNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = subtrahendBiNum.IsValid('Validating subtrahendBiNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBiNumStr, err := subtrahendBiNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != subtrahendBiNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected subtrahendBiNumStr = '%v'\n"+
			"  Actual subtrahendBiNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, subtrahendBiNumStr)

		return
	}

	subtrahendPrecisionBigInt, err := subtrahendBiNum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendPrecisionBigInt, err := \n"+
			"  subtrahendBiNum.GetPrecisionBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	subtrahendBigInt, err := subtrahendBiNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"subtrahendBigInt, err := \n"+
			"  subtrahendBiNum.GetBigInt()\n"+
			"subtrahendBiNum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendBiNumStr, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr, &expectedNumSeps)\n"+
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
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigINumPrecisionBigInt, err := \n"+
			"  expectedBigINum.GetPrecisionBigInt()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
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

	resultBigInt, resultPrecisionBigInt, err := new(BigIntMathSubtract).BigIntSubtract(
		minuendBigInt,
		minuendPrecisionBigInt,
		subtrahendBigInt,
		subtrahendPrecisionBigInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathSubtract).BigIntSubtract(\n"+
			"  minuendBigInt, minuendPrecisionBigInt, subtrahendBigInt, subtrahendPrecisionBigInt)"+
			"minuendBigInt= '%v'\n"+
			"minuendPrecisionBigInt= '%v'\n"+
			"subtrahendBigInt= '%v'\n"+
			" subtrahendPrecisionBigInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendBigInt.Text(10),
			minuendPrecisionBigInt.Text(10),
			subtrahendBigInt.Text(10),
			subtrahendPrecisionBigInt.Text(10),
			err.Error())

		return
	}

	resultBigINum, err := new(BigIntNum).NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINum, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecisionNumSeps(resultBigInt, resultPrecisionBigInt, expectedNumSeps)\n"+
			"resultBigInt= '%v'\n"+
			"resultPrecisionBigInt= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			resultBigInt.Text(10),
			resultPrecisionBigInt.Text(10),
			expectedNumSeps.String(),
			err.Error())

		return
	}

	err = resultBigINum.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = resultBigINum.IsValid("Validating resultBigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBigINum.IsValid('Validating resultBigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBigINumStr, err := resultBigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultSignValue, err := resultBigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultSignValue, err := resultBigINum.GetSign()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	if expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionBigInt.Cmp(resultPrecisionBigInt) != 0 \n"+
			"Expected resultPrecisionBigInt = '%v'\n"+
			"  Actual resultPrecisionBigInt = '%v'\n\n",
			ePrefix, expectedBigINumPrecisionBigInt.Text(10), resultPrecisionBigInt.Text(10))

		return
	}

	expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.Equal(resultBigINum)\n"+
			"expectedBigINum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBigINumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values NOT Equal\n"+
			"Because expectedEqualsResult = 'false' \n"+
			"Expected result = '%v'\n"+
			"  Actual result = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

		return
	}

	if expectedBigINumStr != resultBigINumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != resultBigINumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBigINumStr)

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
			"Error: Numeric Separator Values NOT Equal\n"+
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

func TestBigIntMathSubtract_FixedDecimalSubtract_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_FixedDecimalSubtract_01"

	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bIFixDecMinuend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		minuendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecMinuend.IsValid("Validating bIFixDecMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecMinuend.IsValid('Validating bIFixDecMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if minuendStr != bIFixDecMinuendNumStr {

		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != bIFixDecMinuendNumStr\n"+
			"Expected bIFixDecMinuendNumStr = '%v'\n"+
			"  Actual bIFixDecMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, bIFixDecMinuendNumStr)

		return
	}

	bIFixDecSubtrahend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		subtrahendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(subtrahendStr, expectedNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecSubtrahend.IsValid("Validating bIFixDecSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecSubtrahend.IsValid('Validating bIFixDecSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != bIFixDecSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected bIFixDecSubtrahendNumStr = '%v'\n"+
			"  Actual bIFixDecSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, bIFixDecSubtrahendNumStr)

		return
	}

	expectedBIFixDecNum, err := new(BigIntFixedDecimal).
		NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBIFixDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBIFixDecNumberStr \n"+
			"Expected expectedBIFixDecNumberStr = '%v'\n"+
			"  Actual expectedBIFixDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBIFixDecNumberStr)

		return
	}

	expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBigIFixDecNumSeps, err := expectedBIFixDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIFixDecNumSeps, err := \n"+
			"  expectedBIFixDecNum.GetNumericSeparatorsDto()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBIFixDecNumPrecisionUint, err := expectedBIFixDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumPrecisionUint, err := \n"+
			"  expectedBIFixDecNum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	resultBIFixDec, err := new(BigIntMathSubtract).FixedDecimalSubtract(
		bIFixDecMinuend,
		bIFixDecSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDec, err := new(BigIntMathSubtract).\n"+
			"  FixedDecimalSubtract(bIFixDecMinuend, bIFixDecSubtrahend)\n"+
			"bIFixDecMinuend= '%v'\n"+
			"bIFixDecSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bIFixDecMinuendNumStr,
			bIFixDecSubtrahendNumStr,
			err.Error())

		return
	}

	err = resultBIFixDec.IsValid("Validating resultBIFixDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.IsValid('Validating resultBIFixDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecSignValue, err := resultBIFixDec.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecSignValue, err := resultBIFixDec.GetSign()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecNumSeps, err := resultBIFixDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumSeps, err := \n"+
			"  resultBIFixDec.GetNumericSeparatorsDto()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecPrecisionUint, err := resultBIFixDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecPrecisionUint, err := \n"+
			"  resultBIFixDec.GetPrecisionUint()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBIFixDecNum.EqualValue(resultBIFixDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.EqualValue(result)\n"+
			"expectedBIFixDecNum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBIFixDecNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values ARE NOT Equal\n"+
			"Because expectedBIFixDecNum != resultBIFixDec \n"+
			"Expected resultBIFixDec = '%v'\n"+
			"  Actual resultBIFixDec = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumBigInt.Cmp(resultBIFixDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBIFixDecNumBigInt.Text(10), resultBIFixDecBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultBIFixDecNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumPrecisionUint != resultBIFixDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedBIFixDecNumPrecisionUint, resultBIFixDecPrecisionUint)

		return
	}

	if expectedBigINumSign != resultBIFixDecSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultBIFixDecSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultBIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBIFixDecNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigIFixDecNumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_FixedDecimalSubtract_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_FixedDecimalSubtract_02"

	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bIFixDecMinuend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		minuendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecMinuend.IsValid("Validating bIFixDecMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecMinuend.IsValid('Validating bIFixDecMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if minuendStr != bIFixDecMinuendNumStr {

		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != bIFixDecMinuendNumStr\n"+
			"Expected bIFixDecMinuendNumStr = '%v'\n"+
			"  Actual bIFixDecMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, bIFixDecMinuendNumStr)

		return
	}

	bIFixDecSubtrahend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		subtrahendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(subtrahendStr, expectedNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecSubtrahend.IsValid("Validating bIFixDecSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecSubtrahend.IsValid('Validating bIFixDecSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != bIFixDecSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected bIFixDecSubtrahendNumStr = '%v'\n"+
			"  Actual bIFixDecSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, bIFixDecSubtrahendNumStr)

		return
	}

	expectedBIFixDecNum, err := new(BigIntFixedDecimal).
		NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBIFixDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBIFixDecNumberStr \n"+
			"Expected expectedBIFixDecNumberStr = '%v'\n"+
			"  Actual expectedBIFixDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBIFixDecNumberStr)

		return
	}

	expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBigIFixDecNumSeps, err := expectedBIFixDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIFixDecNumSeps, err := \n"+
			"  expectedBIFixDecNum.GetNumericSeparatorsDto()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBIFixDecNumPrecisionUint, err := expectedBIFixDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumPrecisionUint, err := \n"+
			"  expectedBIFixDecNum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	resultBIFixDec, err := new(BigIntMathSubtract).FixedDecimalSubtract(
		bIFixDecMinuend,
		bIFixDecSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDec, err := new(BigIntMathSubtract).\n"+
			"  FixedDecimalSubtract(bIFixDecMinuend, bIFixDecSubtrahend)\n"+
			"bIFixDecMinuend= '%v'\n"+
			"bIFixDecSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bIFixDecMinuendNumStr,
			bIFixDecSubtrahendNumStr,
			err.Error())

		return
	}

	err = resultBIFixDec.IsValid("Validating resultBIFixDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.IsValid('Validating resultBIFixDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecSignValue, err := resultBIFixDec.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecSignValue, err := resultBIFixDec.GetSign()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecNumSeps, err := resultBIFixDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumSeps, err := \n"+
			"  resultBIFixDec.GetNumericSeparatorsDto()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecPrecisionUint, err := resultBIFixDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecPrecisionUint, err := \n"+
			"  resultBIFixDec.GetPrecisionUint()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBIFixDecNum.EqualValue(resultBIFixDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.EqualValue(result)\n"+
			"expectedBIFixDecNum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBIFixDecNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values ARE NOT Equal\n"+
			"Because expectedBIFixDecNum != resultBIFixDec \n"+
			"Expected resultBIFixDec = '%v'\n"+
			"  Actual resultBIFixDec = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumBigInt.Cmp(resultBIFixDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBIFixDecNumBigInt.Text(10), resultBIFixDecBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultBIFixDecNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumPrecisionUint != resultBIFixDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedBIFixDecNumPrecisionUint, resultBIFixDecPrecisionUint)

		return
	}

	if expectedBigINumSign != resultBIFixDecSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultBIFixDecSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultBIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBIFixDecNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigIFixDecNumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_FixedDecimalSubtract_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_FixedDecimalSubtract_03"

	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"

	expectedBigINumSign := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bIFixDecMinuend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		minuendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecMinuend.IsValid("Validating bIFixDecMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecMinuend.IsValid('Validating bIFixDecMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if minuendStr != bIFixDecMinuendNumStr {

		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != bIFixDecMinuendNumStr\n"+
			"Expected bIFixDecMinuendNumStr = '%v'\n"+
			"  Actual bIFixDecMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, bIFixDecMinuendNumStr)

		return
	}

	bIFixDecSubtrahend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		subtrahendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(subtrahendStr, expectedNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecSubtrahend.IsValid("Validating bIFixDecSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecSubtrahend.IsValid('Validating bIFixDecSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != bIFixDecSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected bIFixDecSubtrahendNumStr = '%v'\n"+
			"  Actual bIFixDecSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, bIFixDecSubtrahendNumStr)

		return
	}

	expectedBIFixDecNum, err := new(BigIntFixedDecimal).
		NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBIFixDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBIFixDecNumberStr \n"+
			"Expected expectedBIFixDecNumberStr = '%v'\n"+
			"  Actual expectedBIFixDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBIFixDecNumberStr)

		return
	}

	expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBigIFixDecNumSeps, err := expectedBIFixDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIFixDecNumSeps, err := \n"+
			"  expectedBIFixDecNum.GetNumericSeparatorsDto()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBIFixDecNumPrecisionUint, err := expectedBIFixDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumPrecisionUint, err := \n"+
			"  expectedBIFixDecNum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	resultBIFixDec, err := new(BigIntMathSubtract).FixedDecimalSubtract(
		bIFixDecMinuend,
		bIFixDecSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDec, err := new(BigIntMathSubtract).\n"+
			"  FixedDecimalSubtract(bIFixDecMinuend, bIFixDecSubtrahend)\n"+
			"bIFixDecMinuend= '%v'\n"+
			"bIFixDecSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bIFixDecMinuendNumStr,
			bIFixDecSubtrahendNumStr,
			err.Error())

		return
	}

	err = resultBIFixDec.IsValid("Validating resultBIFixDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.IsValid('Validating resultBIFixDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecSignValue, err := resultBIFixDec.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecSignValue, err := resultBIFixDec.GetSign()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecNumSeps, err := resultBIFixDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumSeps, err := \n"+
			"  resultBIFixDec.GetNumericSeparatorsDto()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecPrecisionUint, err := resultBIFixDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecPrecisionUint, err := \n"+
			"  resultBIFixDec.GetPrecisionUint()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBIFixDecNum.EqualValue(resultBIFixDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.EqualValue(result)\n"+
			"expectedBIFixDecNum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBIFixDecNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values ARE NOT Equal\n"+
			"Because expectedBIFixDecNum != resultBIFixDec \n"+
			"Expected resultBIFixDec = '%v'\n"+
			"  Actual resultBIFixDec = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumBigInt.Cmp(resultBIFixDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBIFixDecNumBigInt.Text(10), resultBIFixDecBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultBIFixDecNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumPrecisionUint != resultBIFixDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedBIFixDecNumPrecisionUint, resultBIFixDecPrecisionUint)

		return
	}

	if expectedBigINumSign != resultBIFixDecSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultBIFixDecSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultBIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBIFixDecNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigIFixDecNumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_FixedDecimalSubtract_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_FixedDecimalSubtract_04"

	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bIFixDecMinuend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		minuendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecMinuend.IsValid("Validating bIFixDecMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecMinuend.IsValid('Validating bIFixDecMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if minuendStr != bIFixDecMinuendNumStr {

		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != bIFixDecMinuendNumStr\n"+
			"Expected bIFixDecMinuendNumStr = '%v'\n"+
			"  Actual bIFixDecMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, bIFixDecMinuendNumStr)

		return
	}

	bIFixDecSubtrahend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		subtrahendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(subtrahendStr, expectedNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecSubtrahend.IsValid("Validating bIFixDecSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecSubtrahend.IsValid('Validating bIFixDecSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != bIFixDecSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected bIFixDecSubtrahendNumStr = '%v'\n"+
			"  Actual bIFixDecSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, bIFixDecSubtrahendNumStr)

		return
	}

	expectedBIFixDecNum, err := new(BigIntFixedDecimal).
		NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBIFixDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBIFixDecNumberStr \n"+
			"Expected expectedBIFixDecNumberStr = '%v'\n"+
			"  Actual expectedBIFixDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBIFixDecNumberStr)

		return
	}

	expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBigIFixDecNumSeps, err := expectedBIFixDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIFixDecNumSeps, err := \n"+
			"  expectedBIFixDecNum.GetNumericSeparatorsDto()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBIFixDecNumPrecisionUint, err := expectedBIFixDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumPrecisionUint, err := \n"+
			"  expectedBIFixDecNum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	resultBIFixDec, err := new(BigIntMathSubtract).FixedDecimalSubtract(
		bIFixDecMinuend,
		bIFixDecSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDec, err := new(BigIntMathSubtract).\n"+
			"  FixedDecimalSubtract(bIFixDecMinuend, bIFixDecSubtrahend)\n"+
			"bIFixDecMinuend= '%v'\n"+
			"bIFixDecSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bIFixDecMinuendNumStr,
			bIFixDecSubtrahendNumStr,
			err.Error())

		return
	}

	err = resultBIFixDec.IsValid("Validating resultBIFixDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.IsValid('Validating resultBIFixDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecSignValue, err := resultBIFixDec.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecSignValue, err := resultBIFixDec.GetSign()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecNumSeps, err := resultBIFixDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumSeps, err := \n"+
			"  resultBIFixDec.GetNumericSeparatorsDto()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecPrecisionUint, err := resultBIFixDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecPrecisionUint, err := \n"+
			"  resultBIFixDec.GetPrecisionUint()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBIFixDecNum.EqualValue(resultBIFixDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.EqualValue(result)\n"+
			"expectedBIFixDecNum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBIFixDecNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values ARE NOT Equal\n"+
			"Because expectedBIFixDecNum != resultBIFixDec \n"+
			"Expected resultBIFixDec = '%v'\n"+
			"  Actual resultBIFixDec = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumBigInt.Cmp(resultBIFixDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBIFixDecNumBigInt.Text(10), resultBIFixDecBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultBIFixDecNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumPrecisionUint != resultBIFixDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedBIFixDecNumPrecisionUint, resultBIFixDecPrecisionUint)

		return
	}

	if expectedBigINumSign != resultBIFixDecSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultBIFixDecSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultBIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBIFixDecNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigIFixDecNumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_FixedDecimalSubtract_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_FixedDecimalSubtract_05"

	minuendStr := "5"

	subtrahendStr := "5"

	expectedBigINumStr := "0"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bIFixDecMinuend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		minuendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecMinuend.IsValid("Validating bIFixDecMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecMinuend.IsValid('Validating bIFixDecMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if minuendStr != bIFixDecMinuendNumStr {

		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != bIFixDecMinuendNumStr\n"+
			"Expected bIFixDecMinuendNumStr = '%v'\n"+
			"  Actual bIFixDecMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, bIFixDecMinuendNumStr)

		return
	}

	bIFixDecSubtrahend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		subtrahendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(subtrahendStr, expectedNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecSubtrahend.IsValid("Validating bIFixDecSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecSubtrahend.IsValid('Validating bIFixDecSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != bIFixDecSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected bIFixDecSubtrahendNumStr = '%v'\n"+
			"  Actual bIFixDecSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, bIFixDecSubtrahendNumStr)

		return
	}

	expectedBIFixDecNum, err := new(BigIntFixedDecimal).
		NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBIFixDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBIFixDecNumberStr \n"+
			"Expected expectedBIFixDecNumberStr = '%v'\n"+
			"  Actual expectedBIFixDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBIFixDecNumberStr)

		return
	}

	expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBigIFixDecNumSeps, err := expectedBIFixDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIFixDecNumSeps, err := \n"+
			"  expectedBIFixDecNum.GetNumericSeparatorsDto()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBIFixDecNumPrecisionUint, err := expectedBIFixDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumPrecisionUint, err := \n"+
			"  expectedBIFixDecNum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	resultBIFixDec, err := new(BigIntMathSubtract).FixedDecimalSubtract(
		bIFixDecMinuend,
		bIFixDecSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDec, err := new(BigIntMathSubtract).\n"+
			"  FixedDecimalSubtract(bIFixDecMinuend, bIFixDecSubtrahend)\n"+
			"bIFixDecMinuend= '%v'\n"+
			"bIFixDecSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bIFixDecMinuendNumStr,
			bIFixDecSubtrahendNumStr,
			err.Error())

		return
	}

	err = resultBIFixDec.IsValid("Validating resultBIFixDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.IsValid('Validating resultBIFixDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecSignValue, err := resultBIFixDec.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecSignValue, err := resultBIFixDec.GetSign()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecNumSeps, err := resultBIFixDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumSeps, err := \n"+
			"  resultBIFixDec.GetNumericSeparatorsDto()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecPrecisionUint, err := resultBIFixDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecPrecisionUint, err := \n"+
			"  resultBIFixDec.GetPrecisionUint()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBIFixDecNum.EqualValue(resultBIFixDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.EqualValue(result)\n"+
			"expectedBIFixDecNum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBIFixDecNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values ARE NOT Equal\n"+
			"Because expectedBIFixDecNum != resultBIFixDec \n"+
			"Expected resultBIFixDec = '%v'\n"+
			"  Actual resultBIFixDec = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumBigInt.Cmp(resultBIFixDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBIFixDecNumBigInt.Text(10), resultBIFixDecBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultBIFixDecNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumPrecisionUint != resultBIFixDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedBIFixDecNumPrecisionUint, resultBIFixDecPrecisionUint)

		return
	}

	if expectedBigINumSign != resultBIFixDecSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultBIFixDecSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultBIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBIFixDecNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigIFixDecNumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_FixedDecimalSubtract_06(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_FixedDecimalSubtract_06"

	minuendStr := "-5"

	subtrahendStr := "5"

	expectedBigINumStr := "-10"

	expectedBigINumSign := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bIFixDecMinuend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		minuendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecMinuend.IsValid("Validating bIFixDecMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecMinuend.IsValid('Validating bIFixDecMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if minuendStr != bIFixDecMinuendNumStr {

		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != bIFixDecMinuendNumStr\n"+
			"Expected bIFixDecMinuendNumStr = '%v'\n"+
			"  Actual bIFixDecMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, bIFixDecMinuendNumStr)

		return
	}

	bIFixDecSubtrahend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		subtrahendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(subtrahendStr, expectedNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecSubtrahend.IsValid("Validating bIFixDecSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecSubtrahend.IsValid('Validating bIFixDecSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != bIFixDecSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected bIFixDecSubtrahendNumStr = '%v'\n"+
			"  Actual bIFixDecSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, bIFixDecSubtrahendNumStr)

		return
	}

	expectedBIFixDecNum, err := new(BigIntFixedDecimal).
		NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBIFixDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBIFixDecNumberStr \n"+
			"Expected expectedBIFixDecNumberStr = '%v'\n"+
			"  Actual expectedBIFixDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBIFixDecNumberStr)

		return
	}

	expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBigIFixDecNumSeps, err := expectedBIFixDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIFixDecNumSeps, err := \n"+
			"  expectedBIFixDecNum.GetNumericSeparatorsDto()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBIFixDecNumPrecisionUint, err := expectedBIFixDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumPrecisionUint, err := \n"+
			"  expectedBIFixDecNum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	resultBIFixDec, err := new(BigIntMathSubtract).FixedDecimalSubtract(
		bIFixDecMinuend,
		bIFixDecSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDec, err := new(BigIntMathSubtract).\n"+
			"  FixedDecimalSubtract(bIFixDecMinuend, bIFixDecSubtrahend)\n"+
			"bIFixDecMinuend= '%v'\n"+
			"bIFixDecSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bIFixDecMinuendNumStr,
			bIFixDecSubtrahendNumStr,
			err.Error())

		return
	}

	err = resultBIFixDec.IsValid("Validating resultBIFixDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.IsValid('Validating resultBIFixDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecSignValue, err := resultBIFixDec.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecSignValue, err := resultBIFixDec.GetSign()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecNumSeps, err := resultBIFixDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumSeps, err := \n"+
			"  resultBIFixDec.GetNumericSeparatorsDto()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecPrecisionUint, err := resultBIFixDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecPrecisionUint, err := \n"+
			"  resultBIFixDec.GetPrecisionUint()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBIFixDecNum.EqualValue(resultBIFixDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.EqualValue(result)\n"+
			"expectedBIFixDecNum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBIFixDecNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values ARE NOT Equal\n"+
			"Because expectedBIFixDecNum != resultBIFixDec \n"+
			"Expected resultBIFixDec = '%v'\n"+
			"  Actual resultBIFixDec = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumBigInt.Cmp(resultBIFixDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBIFixDecNumBigInt.Text(10), resultBIFixDecBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultBIFixDecNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumPrecisionUint != resultBIFixDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedBIFixDecNumPrecisionUint, resultBIFixDecPrecisionUint)

		return
	}

	if expectedBigINumSign != resultBIFixDecSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultBIFixDecSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultBIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBIFixDecNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigIFixDecNumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_FixedDecimalSubtract_07(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_FixedDecimalSubtract_07"

	minuendStr := "0"

	subtrahendStr := "0"

	expectedBigINumStr := "0"

	expectedBigINumSign := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bIFixDecMinuend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		minuendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecMinuend.IsValid("Validating bIFixDecMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecMinuend.IsValid('Validating bIFixDecMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if minuendStr != bIFixDecMinuendNumStr {

		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != bIFixDecMinuendNumStr\n"+
			"Expected bIFixDecMinuendNumStr = '%v'\n"+
			"  Actual bIFixDecMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, bIFixDecMinuendNumStr)

		return
	}

	bIFixDecSubtrahend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		subtrahendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(subtrahendStr, expectedNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecSubtrahend.IsValid("Validating bIFixDecSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecSubtrahend.IsValid('Validating bIFixDecSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != bIFixDecSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected bIFixDecSubtrahendNumStr = '%v'\n"+
			"  Actual bIFixDecSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, bIFixDecSubtrahendNumStr)

		return
	}

	expectedBIFixDecNum, err := new(BigIntFixedDecimal).
		NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBIFixDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBIFixDecNumberStr \n"+
			"Expected expectedBIFixDecNumberStr = '%v'\n"+
			"  Actual expectedBIFixDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBIFixDecNumberStr)

		return
	}

	expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBigIFixDecNumSeps, err := expectedBIFixDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIFixDecNumSeps, err := \n"+
			"  expectedBIFixDecNum.GetNumericSeparatorsDto()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBIFixDecNumPrecisionUint, err := expectedBIFixDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumPrecisionUint, err := \n"+
			"  expectedBIFixDecNum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	resultBIFixDec, err := new(BigIntMathSubtract).FixedDecimalSubtract(
		bIFixDecMinuend,
		bIFixDecSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDec, err := new(BigIntMathSubtract).\n"+
			"  FixedDecimalSubtract(bIFixDecMinuend, bIFixDecSubtrahend)\n"+
			"bIFixDecMinuend= '%v'\n"+
			"bIFixDecSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bIFixDecMinuendNumStr,
			bIFixDecSubtrahendNumStr,
			err.Error())

		return
	}

	err = resultBIFixDec.IsValid("Validating resultBIFixDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.IsValid('Validating resultBIFixDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecSignValue, err := resultBIFixDec.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecSignValue, err := resultBIFixDec.GetSign()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecNumSeps, err := resultBIFixDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumSeps, err := \n"+
			"  resultBIFixDec.GetNumericSeparatorsDto()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecPrecisionUint, err := resultBIFixDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecPrecisionUint, err := \n"+
			"  resultBIFixDec.GetPrecisionUint()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBIFixDecNum.EqualValue(resultBIFixDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.EqualValue(result)\n"+
			"expectedBIFixDecNum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBIFixDecNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values ARE NOT Equal\n"+
			"Because expectedBIFixDecNum != resultBIFixDec \n"+
			"Expected resultBIFixDec = '%v'\n"+
			"  Actual resultBIFixDec = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumBigInt.Cmp(resultBIFixDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBIFixDecNumBigInt.Text(10), resultBIFixDecBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultBIFixDecNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumPrecisionUint != resultBIFixDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedBIFixDecNumPrecisionUint, resultBIFixDecPrecisionUint)

		return
	}

	if expectedBigINumSign != resultBIFixDecSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultBIFixDecSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultBIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBIFixDecNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigIFixDecNumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_FixedDecimalSubtract_08(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_FixedDecimalSubtract_08"

	minuendStr := "50.0"

	subtrahendStr := "2.60134"

	expectedBigINumStr := "47.39866"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bIFixDecMinuend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		minuendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, expectedNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecMinuend.IsValid("Validating bIFixDecMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecMinuend.IsValid('Validating bIFixDecMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if minuendStr != bIFixDecMinuendNumStr {

		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != bIFixDecMinuendNumStr\n"+
			"Expected bIFixDecMinuendNumStr = '%v'\n"+
			"  Actual bIFixDecMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, bIFixDecMinuendNumStr)

		return
	}

	bIFixDecSubtrahend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		subtrahendStr,
		expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(subtrahendStr, expectedNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecSubtrahend.IsValid("Validating bIFixDecSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecSubtrahend.IsValid('Validating bIFixDecSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != bIFixDecSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected bIFixDecSubtrahendNumStr = '%v'\n"+
			"  Actual bIFixDecSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, bIFixDecSubtrahendNumStr)

		return
	}

	expectedBIFixDecNum, err := new(BigIntFixedDecimal).
		NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
			"expectedBigINumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, err.Error())
		return
	}

	expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBIFixDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBIFixDecNumberStr \n"+
			"Expected expectedBIFixDecNumberStr = '%v'\n"+
			"  Actual expectedBIFixDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBIFixDecNumberStr)

		return
	}

	expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBigIFixDecNumSeps, err := expectedBIFixDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIFixDecNumSeps, err := \n"+
			"  expectedBIFixDecNum.GetNumericSeparatorsDto()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBIFixDecNumPrecisionUint, err := expectedBIFixDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumPrecisionUint, err := \n"+
			"  expectedBIFixDecNum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	resultBIFixDec, err := new(BigIntMathSubtract).FixedDecimalSubtract(
		bIFixDecMinuend,
		bIFixDecSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDec, err := new(BigIntMathSubtract).\n"+
			"  FixedDecimalSubtract(bIFixDecMinuend, bIFixDecSubtrahend)\n"+
			"bIFixDecMinuend= '%v'\n"+
			"bIFixDecSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bIFixDecMinuendNumStr,
			bIFixDecSubtrahendNumStr,
			err.Error())

		return
	}

	err = resultBIFixDec.IsValid("Validating resultBIFixDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.IsValid('Validating resultBIFixDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecSignValue, err := resultBIFixDec.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecSignValue, err := resultBIFixDec.GetSign()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecNumSeps, err := resultBIFixDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumSeps, err := \n"+
			"  resultBIFixDec.GetNumericSeparatorsDto()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecPrecisionUint, err := resultBIFixDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecPrecisionUint, err := \n"+
			"  resultBIFixDec.GetPrecisionUint()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBIFixDecNum.EqualValue(resultBIFixDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.EqualValue(result)\n"+
			"expectedBIFixDecNum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBIFixDecNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values ARE NOT Equal\n"+
			"Because expectedBIFixDecNum != resultBIFixDec \n"+
			"Expected resultBIFixDec = '%v'\n"+
			"  Actual resultBIFixDec = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumBigInt.Cmp(resultBIFixDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBIFixDecNumBigInt.Text(10), resultBIFixDecBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultBIFixDecNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumPrecisionUint != resultBIFixDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedBIFixDecNumPrecisionUint, resultBIFixDecPrecisionUint)

		return
	}

	if expectedBigINumSign != resultBIFixDecSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultBIFixDecSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultBIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBIFixDecNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigIFixDecNumSeps.String())

		return
	}

	return
}

func TestBigIntMathSubtract_FixedDecimalSubtract_09(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_FixedDecimalSubtract_09"

	var err error

	minuendStr := "50.0"

	subtrahendStr := "2.60134"

	expectedBigINumStr := "47,39866"

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

	bIFixDecMinuend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		minuendStr,
		usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, usaNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecMinuend.IsValid("Validating bIFixDecMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecMinuend.IsValid('Validating bIFixDecMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecMinuendNumStr, err := bIFixDecMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if minuendStr != bIFixDecMinuendNumStr {

		t.Errorf("%v\n"+
			"Error: Number String Values Are Not Equal\n"+
			"Because minuendStr != bIFixDecMinuendNumStr\n"+
			"Expected bIFixDecMinuendNumStr = '%v'\n"+
			"  Actual bIFixDecMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, bIFixDecMinuendNumStr)

		return
	}

	bIFixDecSubtrahend, err := new(BigIntFixedDecimal).NewNumStrWithNumSeps(
		subtrahendStr,
		usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahend, err := new(BigIntFixedDecimal).\n"+
			"  NewNumStrWithNumSeps(subtrahendStr, usaNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, usaNumSeps.String(), err.Error())
		return
	}

	err = bIFixDecSubtrahend.IsValid("Validating bIFixDecSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bIFixDecSubtrahend.IsValid('Validating bIFixDecSubtrahend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bIFixDecSubtrahendNumStr, err := bIFixDecSubtrahend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if subtrahendStr != bIFixDecSubtrahendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because subtrahendStr != subtrahendBiNumStr \n"+
			"Expected bIFixDecSubtrahendNumStr = '%v'\n"+
			"  Actual bIFixDecSubtrahendNumStr = '%v'\n\n",
			ePrefix, subtrahendStr, bIFixDecSubtrahendNumStr)

		return
	}

	expectedBIFixDecNum, err := new(BigIntFixedDecimal).
		NewNumStrWithNumSeps(expectedBigINumStr, expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNum, err := new(BigIntNum).NewNumStr(\n"+
			"  expectedBigINumStr, expectedNumSeps\n"+
			"expectedBigINumStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
		return
	}

	expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumberStr, err := expectedBIFixDecNum.GetBigInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedBigINumStr != expectedBIFixDecNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Number String Values NOT Equal\n"+
			"Because expectedBigINumStr != expectedBIFixDecNumberStr \n"+
			"Expected expectedBIFixDecNumberStr = '%v'\n"+
			"  Actual expectedBIFixDecNumberStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, expectedBIFixDecNumberStr)

		return
	}

	expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumBigInt, err := expectedBIFixDecNum.GetBigInt()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBigIFixDecNumSeps, err := expectedBIFixDecNum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIFixDecNumSeps, err := \n"+
			"  expectedBIFixDecNum.GetNumericSeparatorsDto()\n"+
			"expectedBIFixDecNum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	expectedBIFixDecNumPrecisionUint, err := expectedBIFixDecNum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBIFixDecNumPrecisionUint, err := \n"+
			"  expectedBIFixDecNum.GetPrecisionUint()\n"+
			"expectedBigINum= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBIFixDecNumberStr, err.Error())
		return
	}

	resultBIFixDec, err := new(BigIntMathSubtract).FixedDecimalSubtract(
		bIFixDecMinuend,
		bIFixDecSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDec, err := new(BigIntMathSubtract).\n"+
			"  FixedDecimalSubtract(bIFixDecMinuend, bIFixDecSubtrahend)\n"+
			"bIFixDecMinuend= '%v'\n"+
			"bIFixDecSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bIFixDecMinuendNumStr,
			bIFixDecSubtrahendNumStr,
			err.Error())

		return
	}

	err = resultBIFixDec.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.SetNumericSeparatorsDto(expectedNumSeps)\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	err = resultBIFixDec.IsValid("Validating resultBIFixDec")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = resultBIFixDec.IsValid('Validating resultBIFixDec')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumStr, err := resultBIFixDec.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecBigInt, err := resultBIFixDec.GetBigInt()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecSignValue, err := resultBIFixDec.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecSignValue, err := resultBIFixDec.GetSign()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecNumSeps, err := resultBIFixDec.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecNumSeps, err := \n"+
			"  resultBIFixDec.GetNumericSeparatorsDto()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	resultBIFixDecPrecisionUint, err := resultBIFixDec.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultBIFixDecPrecisionUint, err := \n"+
			"  resultBIFixDec.GetPrecisionUint()\n"+
			"resultBIFixDec= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultBIFixDecNumStr, err.Error())
		return
	}

	expectedEqualsResult, err := expectedBIFixDecNum.EqualValue(resultBIFixDec)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResult, err := expectedBigINum.EqualValue(result)\n"+
			"expectedBIFixDecNum= '%v'\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultBIFixDecNumStr, err.Error())
		return
	}

	if !expectedEqualsResult {
		t.Errorf("%v\n"+
			"Error: Expected and 'result' values ARE NOT Equal\n"+
			"Because expectedBIFixDecNum != resultBIFixDec \n"+
			"Expected resultBIFixDec = '%v'\n"+
			"  Actual resultBIFixDec = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumBigInt.Cmp(resultBIFixDecBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Big Int Values NOT Equal\n"+
			"Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
			"Expected resultBigInt = '%v'\n"+
			"  Actual resultBigInt = '%v'\n\n",
			ePrefix, expectedBIFixDecNumBigInt.Text(10), resultBIFixDecBigInt.Text(10))

		return
	}

	if expectedBigINumStr != resultBIFixDecNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values ARE NOT Equal\n"+
			"Because expectedBigINumStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedBigINumStr, resultBIFixDecNumStr)

		return
	}

	if expectedBIFixDecNumPrecisionUint != resultBIFixDecPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Precision Values ARE NOT Equal\n"+
			"Because expectedBigINumPrecisionUint != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedBIFixDecNumPrecisionUint, resultBIFixDecPrecisionUint)

		return
	}

	if expectedBigINumSign != resultBIFixDecSignValue {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedBigINumSign != resultSignValue \n"+
			"Expected resultSignValue = '%v'\n"+
			"  Actual resultSignValue = '%v'\n\n",
			ePrefix, expectedBigINumSign, resultBIFixDecSignValue)

		return
	}

	if !expectedNumSeps.Equal(resultBIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultBIFixDecNumSeps.String())

		return
	}

	if !expectedNumSeps.Equal(expectedBigIFixDecNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != expectedBigINumSeps \n"+
			"Expected expectedBigINumSeps = '%v'\n"+
			"  Actual expectedBigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), expectedBigIFixDecNumSeps.String())

		return
	}

	return
}
