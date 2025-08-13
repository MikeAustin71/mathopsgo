package mathops

import (
	"fmt"
	"testing"
)

func TestBigIntMathSubtract_SubtractIntArySeries_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractIntArySeries_01"

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

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = iaMinuend.IsValid("Validating iaMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMinuend.IsValid('Validating iaMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMinuendNumStr, err := iaMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != iaMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != iaMinuendNumStr\n"+
			"Expected iaMinuendNumStr = '%v'\n"+
			"  Actual iaMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, iaMinuendNumStr)

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

	lenSubtrahends := 6

	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractIntArySeries(\n"+
			" iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
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

func TestBigIntMathSubtract_SubtractIntArySeries_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractIntArySeries_02"

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

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = iaMinuend.IsValid("Validating iaMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMinuend.IsValid('Validating iaMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMinuendNumStr, err := iaMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != iaMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != iaMinuendNumStr\n"+
			"Expected iaMinuendNumStr = '%v'\n"+
			"  Actual iaMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, iaMinuendNumStr)

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

	lenSubtrahends := 6

	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractIntArySeries(\n"+
			" iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
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

func TestBigIntMathSubtract_SubtractIntArySeries_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractIntArySeries_03"

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

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = iaMinuend.IsValid("Validating iaMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMinuend.IsValid('Validating iaMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMinuendNumStr, err := iaMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != iaMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != iaMinuendNumStr\n"+
			"Expected iaMinuendNumStr = '%v'\n"+
			"  Actual iaMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, iaMinuendNumStr)

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

	lenSubtrahends := 6

	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractIntArySeries(\n"+
			" iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
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

func TestBigIntMathSubtract_SubtractIntArySeries_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractIntArySeries_04"

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

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = iaMinuend.IsValid("Validating iaMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMinuend.IsValid('Validating iaMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMinuendNumStr, err := iaMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != iaMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != iaMinuendNumStr\n"+
			"Expected iaMinuendNumStr = '%v'\n"+
			"  Actual iaMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, iaMinuendNumStr)

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

	lenSubtrahends := 6

	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractIntArySeries(\n"+
			" iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
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

func TestBigIntMathSubtract_SubtractIntArySeries_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractIntArySeries_05"

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

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	iaMinuend, err := new(IntAry).NewNumStrWithNumSeps(minuendStr, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuend, err := new(IntAry).\n"+
			"  NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendStr,
			usaNumSeps.NewUSADefaults(),
			err.Error())

		return
	}

	err = iaMinuend.SetNumericSeparatorsDto(expectedNumSeps)

	err = iaMinuend.IsValid("Validating iaMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMinuend.IsValid('Validating iaMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMinuendNumStr, err := iaMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != iaMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != iaMinuendNumStr\n"+
			"Expected iaMinuendNumStr = '%v'\n"+
			"  Actual iaMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, iaMinuendNumStr)

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

	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStrWithNumSeps(subtrahend0, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[0], err = new(IntAry).\n"+
			"  NewNumStrWithNumSeps(subtrahend0, usaNumSeps)\n"+
			"subtrahend0= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, usaNumSeps.String(), err.Error())
		return
	}

	subtrahendAry[1], err = new(IntAry).NewNumStrWithNumSeps(subtrahend1, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[1], err = new(IntAry).\n"+
			"  NewNumStrWithNumSeps(subtrahend1, usaNumSeps)\n"+
			"subtrahend1= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, usaNumSeps.String(), err.Error())
		return
	}

	subtrahendAry[2], err = new(IntAry).NewNumStrWithNumSeps(subtrahend2, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[2], err = new(IntAry).\n"+
			"  NewNumStrWithNumSeps(subtrahend2, usaNumSeps)\n"+
			"subtrahend2= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, usaNumSeps.String(), err.Error())
		return
	}

	subtrahendAry[3], err = new(IntAry).NewNumStrWithNumSeps(subtrahend3, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[3], err = new(IntAry).\n"+
			"  NewNumStrWithNumSeps(subtrahend3, usaNumSeps)\n"+
			"subtrahend3= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, usaNumSeps.String(), err.Error())
		return
	}

	subtrahendAry[4], err = new(IntAry).NewNumStrWithNumSeps(subtrahend4, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[4], err = new(IntAry).\n"+
			"  NewNumStrWithNumSeps(subtrahend4, usaNumSeps)\n"+
			"subtrahend4= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, usaNumSeps.String(), err.Error())
		return
	}

	subtrahendAry[5], err = new(IntAry).NewNumStrWithNumSeps(subtrahend5, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" subtrahendAry[5], err = new(IntAry).\n"+
			"  NewNumStrWithNumSeps(subtrahend5, usaNumSeps)\n"+
			"subtrahend5= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, usaNumSeps.String(), err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractIntArySeries(
		iaMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractIntArySeries(\n"+
			" iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
			"  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, iaMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
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

func TestBigIntMathSubtract_SubtractINumMgr_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgr_01"

	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	decMinuend, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = decMinuend.IsValid("Validating decMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decMinuend.IsValid(Validating decMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decMinuendNumStr, err := decMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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
			"err = nDtoSubtrahend.IsValid(Validating nDtoSubtrahend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	result, err := new(BigIntMathSubtract).SubtractINumMgr(&decMinuend, &nDtoSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgr(&decMinuend, &nDtoSubtrahend)\n"+
			"decMinuend= '%v'\n"+
			"nDtoSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decMinuendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgr_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgr_02"

	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = iaMinuend.IsValid("Validating iaMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaMinuend.IsValid(Validating iaMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaMinuendNumStr, err := iaMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSubtrahend, err := new(BigIntNum).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSubtrahend, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = bINumSubtrahend.IsValid("Validating bINumSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumSubtrahend.IsValid(Validating bINumSubtrahend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumSubtrahendNumStr, err := bINumSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumSubtrahendNumStr, err := bINumSubtrahend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	result, err := new(BigIntMathSubtract).SubtractINumMgr(
		&iaMinuend, &bINumSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgr(&iaMinuend, &bINumSubtrahend\n"+
			"iaMinuend= '%v'\n"+
			"bINumSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			iaMinuendNumStr,
			bINumSubtrahendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgr_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgr_02"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaSubtrahend, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = iaSubtrahend.IsValid("Validating iaSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iaSubtrahend.IsValid(Validating iaSubtrahend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	result, err := new(BigIntMathSubtract).SubtractINumMgr(&nDtoMinuend, &iaSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgr(&nDtoMinuend, &iaSubtrahend\n"+
			"nDtoMinuend= '%v'\n"+
			"iaSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
			iaSubtrahendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgr_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgr_04"

	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"

	expectedBigINumSign := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bINumMinuend, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = bINumMinuend.IsValid("Validating bINumMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumMinuend.IsValid(Validating bINumMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumMinuendNumStr, err := bINumMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumMinuendNumStr, err := bINumMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decSubtrahend, err := new(Decimal).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decSubtrahend, err := new(BigIntNum).NewNumStr(subtrahendStr)\n"+
			"subtrahendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, err.Error())
		return
	}

	err = decSubtrahend.IsValid("Validating decSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decSubtrahend.IsValid(Validating decSubtrahend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decSubtrahendNumStr, err := decSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decSubtrahendNumStr, err := decSubtrahend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	result, err := new(BigIntMathSubtract).SubtractINumMgr(&bINumMinuend, &decSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgr(&bINumMinuend, &decSubtrahend\n"+
			"bINumMinuend= '%v'\n"+
			"decSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumMinuendNumStr,
			decSubtrahendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgr_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgr_05"

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

	decMinuend, err := new(Decimal).NewNumStrWithNumSeps(minuendStr, usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decMinuend, err := new(Decimal).\n"+
			"  NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, usaNumSeps.String(), err.Error())
		return
	}

	err = decMinuend.IsValid("Validating decMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decMinuend.IsValid(Validating decMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decMinuendNumStr, err := decMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoSubtrahend, err := new(NumStrDto).NewNumStrWithNumSeps(subtrahendStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahend, err := new(NumStrDto).\n"+
			" NewNumStrWithNumSeps(subtrahendStr, &usaNumSeps)\n"+
			"subtrahendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahendStr, usaNumSeps.String(), err.Error())
		return
	}

	err = nDtoSubtrahend.IsValid("Validating nDtoSubtrahend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoSubtrahend.IsValid(Validating nDtoSubtrahend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoSubtrahendNumStr, err := nDtoSubtrahend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	result, err := new(BigIntMathSubtract).SubtractINumMgr(&decMinuend, &nDtoSubtrahend)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgr(&decMinuend, &nDtoSubtrahend\n"+
			"decMinuend= '%v'\n"+
			"nDtoSubtrahend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decMinuendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgrArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrArray_01"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	lenSubtrahendsArray := 6

	subtrahendAry := make([]INumMgr, lenSubtrahendsArray)

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[0] = &dec0

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[1] = &bINum1

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[2] = &ia2

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[3] = &nDto3

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[4] = &dec4

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	subtrahendAry[5] = &ia5

	result, err := new(BigIntMathSubtract).SubtractINumMgrArray(&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrArray(\n"+
			"  &nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgrArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrArray_02"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	lenSubtrahends := 6

	subtrahendAry := make([]INumMgr, lenSubtrahends)

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[0] = &dec0

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[1] = &bINum1

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[2] = &ia2

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[3] = &nDto3

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[4] = &dec4

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	subtrahendAry[5] = &ia5

	result, err := new(BigIntMathSubtract).SubtractINumMgrArray(
		&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrArray(\n"+
			"  &nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgrArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrArray_03"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	lenSubtrahends := 6

	subtrahendAry := make([]INumMgr, lenSubtrahends)

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[0] = &dec0

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[1] = &bINum1

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[2] = &ia2

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[3] = &nDto3

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[4] = &dec4

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	subtrahendAry[5] = &ia5

	result, err := new(BigIntMathSubtract).SubtractINumMgrArray(
		&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrArray(\n"+
			"  &nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgrArray_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrArray_04"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	lenSubtrahends := 6

	subtrahendAry := make([]INumMgr, lenSubtrahends)

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[0] = &dec0

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[1] = &bINum1

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[2] = &ia2

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[3] = &nDto3

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[4] = &dec4

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	subtrahendAry[5] = &ia5

	result, err := new(BigIntMathSubtract).SubtractINumMgrArray(
		&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrArray(\n"+
			"  &nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgrArray_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrArray_05"

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
			ePrefix, minuendStr, usaNumSeps.String(), err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	subtrahendAry := make([]INumMgr, lenSubtrahends)

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	subtrahendAry[0] = &dec0

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	subtrahendAry[1] = &bINum1

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	subtrahendAry[2] = &ia2

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	subtrahendAry[3] = &nDto3

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	subtrahendAry[4] = &dec4

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	subtrahendAry[5] = &ia5

	result, err := new(BigIntMathSubtract).SubtractINumMgrArray(
		&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrArray(\n"+
			"  &nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
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

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrOutputToArray_01"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	subtrahendAry := make([]INumMgr, lenSubtrahends)

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

	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	var dec2NumStr string

	var dec, dec2 Decimal

	for i := 0; i < lenSubtrahends; i++ {

		dec, err = new(Decimal).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"dec, err = new(Decimal).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = dec.IsValid(fmt.Sprintf("Validating dec[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = dec[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		subtrahendAry[i] = &dec

		dec2, err = new(Decimal).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"dec2, err = new(Decimal).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		err = dec2.IsValid(fmt.Sprintf("Validating dec2[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = dec2[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		dec2NumStr, err = dec2.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"dec2NumStr, err = dec2[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		if expectedStrs[i] != dec2NumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != dec2NumStr \n"+
				"Expected dec2NumStr[%d] = '%v'\n"+
				"  Actual dec2NumStr[%d] = '%v'\n\n",
				ePrefix, i, i, expectedStrs[i], i, dec2NumStr)

			return
		}

		expectedResultsAry[i] = &dec2

	} // End of Loop

	result, err :=
		new(BigIntMathSubtract).SubtractINumMgrOutputToArray(&nDtoMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrOutputToArray(\n"+
			"  &nDtoMinuend, subtrahendAry[...])\n"+
			"nDtoMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nDtoMinuendNumStr,
			err.Error())

		return
	}

	var expectedResultNumStr, resultNumStr string

	var resultNumSeps NumericSeparatorDto

	for k := 0; k < lenSubtrahends; k++ {

		err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = result[%d].IsValid('Validating result[%d]')\n"+
				"Validation Error= '%v'\n\n",
				ePrefix, k, k, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrOutputToArray_02"

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

	iAryMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iAryMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = iAryMinuend.IsValid("Validating iAryMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = iAryMinuend.IsValid(Validating iAryMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	iAryMinuendNumStr, err := iAryMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iAryMinuendNumStr, err := iAryMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	subtrahendAry := make([]INumMgr, lenSubtrahends)

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

	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	var nDto2NumStr string

	var nDto1, nDto2 NumStrDto

	for i := 0; i < lenSubtrahends; i++ {

		nDto1, err = new(NumStrDto).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDto1, err = new(NumStrDto).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = nDto1.IsValid(fmt.Sprintf("Validating nDto1[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = nDto1[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		subtrahendAry[i] = &nDto1

		nDto2, err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDto2, err = new(NumStrDto).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		err = nDto2.IsValid(fmt.Sprintf("Validating nDto2[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = nDto2[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		nDto2NumStr, err = nDto2.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDto2NumStr, err = nDto2[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		if expectedStrs[i] != nDto2NumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != nDto2NumStr \n"+
				"Expected nDto2NumStr[%d] = '%v'\n"+
				"  Actual nDto2NumStr[%d] = '%v'\n\n",
				ePrefix, i, i, expectedStrs[i], i, nDto2NumStr)

			return
		}

		expectedResultsAry[i] = &nDto2

	} // End of Loop

	result, err :=
		new(BigIntMathSubtract).SubtractINumMgrOutputToArray(
			&iAryMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrOutputToArray(\n"+
			"  &iAryMinuend, subtrahendAry[...])\n"+
			"iAryMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			iAryMinuendNumStr,
			err.Error())

		return
	}

	var expectedResultNumStr, resultNumStr string

	var resultNumSeps NumericSeparatorDto

	for k := 0; k < lenSubtrahends; k++ {

		err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = result[%d].IsValid('Validating result[%d]')\n"+
				"Validation Error= '%v'\n\n",
				ePrefix, k, k, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrOutputToArray_03"

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

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	minuendINumMgr, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendINumMgr, err := new(Decimal).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendINumMgr.IsValid("Validating minuendINumMgr")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendINumMgr.IsValid(Validating minuendINumMgr)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendINumMgrNumStr, err := minuendINumMgr.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendINumMgrNumStr, err := minuendINumMgr.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	subtrahendAry := make([]INumMgr, lenSubtrahends)

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

	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	var ia2NumStr string

	var ia1, ia2 IntAry

	for i := 0; i < lenSubtrahends; i++ {

		ia1, err = new(IntAry).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia1, err = new(IntAry).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = ia1.IsValid(fmt.Sprintf("Validating ia1[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia1[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		subtrahendAry[i] = &ia1

		ia2, err = new(IntAry).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia2, err = new(IntAry).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		err = ia2.IsValid(fmt.Sprintf("Validating ia2[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = ia2[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		ia2NumStr, err = ia2.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"ia2NumStr, err = ia2[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		if expectedStrs[i] != ia2NumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != ia2NumStr \n"+
				"Expected ia2NumStr[%d] = '%v'\n"+
				"  Actual ia2NumStr[%d] = '%v'\n\n",
				ePrefix, i, i, expectedStrs[i], i, ia2NumStr)

			return
		}

		expectedResultsAry[i] = &ia2
	}

	result, err :=
		new(BigIntMathSubtract).SubtractINumMgrOutputToArray(&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrOutputToArray(\n"+
			"  &minuendINumMgr, subtrahendAry[...])\n"+
			"minuendINumMgr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendINumMgrNumStr,
			err.Error())

		return
	}

	var expectedResultNumStr, resultNumStr string

	var resultNumSeps NumericSeparatorDto

	for k := 0; k < lenSubtrahends; k++ {

		err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = result[%d].IsValid('Validating result[%d]')\n"+
				"Validation Error= '%v'\n\n",
				ePrefix, k, k, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrOutputToArray_04"

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

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	minuendINumMgr, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendINumMgr, err := new(Decimal).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendINumMgr.IsValid("Validating minuendINumMgr")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendINumMgr.IsValid(Validating minuendINumMgr)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendINumMgrNumStr, err := minuendINumMgr.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendINumMgrNumStr, err := minuendINumMgr.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	subtrahendAry := make([]INumMgr, lenSubtrahends)

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

	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	var dec2NumStr string

	var dec1, dec2 Decimal

	for i := 0; i < lenSubtrahends; i++ {

		dec1, err = new(Decimal).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"dec1, err = new(Decimal).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = dec1.IsValid(fmt.Sprintf("Validating dec1[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = dec1[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		subtrahendAry[i] = &dec1

		dec2, err = new(Decimal).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"dec2, err = new(Decimal).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		err = dec2.IsValid(fmt.Sprintf("Validating dec2[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = dec2[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		dec2NumStr, err = dec2.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"dec2NumStr, err = dec2[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		if expectedStrs[i] != dec2NumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != dec2NumStr \n"+
				"Expected dec2NumStr[%d] = '%v'\n"+
				"  Actual dec2NumStr[%d] = '%v'\n\n",
				ePrefix, i, i, expectedStrs[i], i, dec2NumStr)

			return
		}

		expectedResultsAry[i] = &dec2

	}

	result, err :=
		new(BigIntMathSubtract).SubtractINumMgrOutputToArray(&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrOutputToArray(\n"+
			"  &minuendINumMgr, subtrahendAry[...])\n"+
			"minuendINumMgr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendINumMgrNumStr,
			err.Error())

		return
	}

	var expectedResultNumStr, resultNumStr string

	var resultNumSeps NumericSeparatorDto

	for k := 0; k < lenSubtrahends; k++ {

		err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = result[%d].IsValid('Validating result[%d]')\n"+
				"Validation Error= '%v'\n\n",
				ePrefix, k, k, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrOutputToArray_05"

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

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	minuendINumMgr, err := new(BigIntNum).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendINumMgr, err := new(BigIntNum).NewNumStr(minuendStr)\n"+
			"minuendStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, err.Error())
		return
	}

	err = minuendINumMgr.IsValid("Validating minuendINumMgr")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendINumMgr.IsValid(Validating minuendINumMgr)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendINumMgrNumStr, err := minuendINumMgr.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendINumMgrNumStr, err := minuendINumMgr.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenSubtrahends := len(subtrahendStrs)

	subtrahendAry := make([]INumMgr, lenSubtrahends)

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

	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	var nDto2NumStr string

	var nDto1, nDto2 NumStrDto

	for i := 0; i < lenSubtrahends; i++ {

		nDto1, err = new(NumStrDto).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDto1, err = new(NumStrDto).\n"+
				"  NewNumStr(subtrahendStrs[%d])\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, subtrahendStrs[i], err.Error())
			return
		}

		err = nDto1.IsValid(fmt.Sprintf("Validating nDto1[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = nDto1[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		subtrahendAry[i] = &nDto1

		nDto2, err = new(NumStrDto).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDto2, err = new(NumStrDto).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, expectedStrs[i], err.Error())
			return
		}

		err = nDto2.IsValid(fmt.Sprintf("Validating nDto2[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = nDto2[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		nDto2NumStr, err = nDto2.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDto2NumStr, err = nDto2[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		if expectedStrs[i] != nDto2NumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != nDto2NumStr \n"+
				"Expected nDto2NumStr[%d] = '%v'\n"+
				"  Actual nDto2NumStr[%d] = '%v'\n\n",
				ePrefix, i, i, expectedStrs[i], i, nDto2NumStr)

			return
		}

		expectedResultsAry[i] = &nDto2

	}

	result, err := new(BigIntMathSubtract).SubtractINumMgrOutputToArray(
		&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrOutputToArray(\n"+
			"  &minuendINumMgr, subtrahendAry[...])\n"+
			"minuendINumMgr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendINumMgrNumStr,
			err.Error())

		return
	}

	var expectedResultNumStr, resultNumStr string

	var resultNumSeps NumericSeparatorDto

	for k := 0; k < lenSubtrahends; k++ {

		err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = result[%d].IsValid('Validating result[%d]')\n"+
				"Validation Error= '%v'\n\n",
				ePrefix, k, k, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrOutputToArray_06(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrOutputToArray_06"

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
			"err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
		return
	}

	usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	minuendINumMgr, err := new(NumStrDto).NewNumStrWithNumSeps(
		minuendStr, &usaNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendINumMgr, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(minuendStr, &usaNumSeps)\n"+
			"minuendStr= '%v'\n"+
			"usaNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, minuendStr, usaNumSeps.String(), err.Error())
		return
	}

	err = minuendINumMgr.IsValid("Validating minuendINumMgr")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = minuendINumMgr.IsValid('Validating minuendINumMgr')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	minuendINumMgrNumStr, err := minuendINumMgr.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"minuendINumMgrNumStr, err := minuendINumMgr.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != minuendINumMgrNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != minuendINumMgrNumStr \n"+
			"Expected minuendINumMgrNumStr = '%v'\n"+
			"  Actual minuendINumMgrNumStr = '%v'\n\n",
			ePrefix, minuendStr, minuendINumMgrNumStr)

		return
	}

	lenSubtrahends := len(subtrahendStrs)

	subtrahendAry := make([]INumMgr, lenSubtrahends)

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

	expectedResultsAry := make([]INumMgr, lenSubtrahends)

	var dec2NumStr string

	var dec1, dec2 Decimal

	for i := 0; i < lenSubtrahends; i++ {

		dec1, err = new(Decimal).NewNumStrWithNumSeps(subtrahendStrs[i], usaNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"nDto1, err = new(NumStrDto).\n"+
				"  NewNumStrWithNumSeps(subtrahendStrs[%d], usaNumSeps))\n"+
				"subtrahendStrs[%d]= '%v'\n"+
				"usaNumSeps= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, subtrahendStrs[i], usaNumSeps.String(), err.Error())
			return
		}

		err = dec1.IsValid(fmt.Sprintf("Validating dec1[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = dec1[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		subtrahendAry[i] = &dec1

		dec2, err = new(Decimal).NewNumStrWithNumSeps(expectedStrs[i], expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"dec2, err = new(NumStrDto).\n"+
				"  NewNumStrWithNumSeps(expectedStrs[%d], expectedNumSeps)\n"+
				"expectedStrs[%d]= '%v'\n"+
				"expectedNumSeps= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, expectedStrs[i], expectedNumSeps.String(), err.Error())
			return
		}

		err = dec2.IsValid(fmt.Sprintf("Validating dec2[%d]", i))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = dec2[%d].IsValid()\n"+
				"Error='%v'\n\n", ePrefix, i, err.Error())
			return
		}

		dec2NumStr, err = dec2.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"dec2NumStr, err = dec2[%d].GetNumStr()\n"+
				"Error= '%v'\n\n", ePrefix, i, err.Error())
			return
		}

		if expectedStrs[i] != dec2NumStr {
			t.Errorf("%v\n"+
				"Error: Number String Values NOT Equal\n"+
				"Because expectedStrs[%d] != dec2NumStr \n"+
				"Expected dec2NumStr[%d] = '%v'\n"+
				"  Actual dec2NumStr[%d] = '%v'\n\n",
				ePrefix, i, i, expectedStrs[i], i, dec2NumStr)

			return
		}

		expectedResultsAry[i] = &dec2
	}

	result, err :=
		new(BigIntMathSubtract).SubtractINumMgrOutputToArray(&minuendINumMgr, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrOutputToArray(\n"+
			"  &minuendINumMgr, subtrahendAry[...])\n"+
			"minuendINumMgr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			minuendINumMgrNumStr,
			err.Error())

		return
	}

	var expectedResultNumStr, resultNumStr string

	var resultNumSeps NumericSeparatorDto

	for k := 0; k < lenSubtrahends; k++ {

		err = result[k].SetNumericSeparatorsDto(expectedNumSeps)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"result[%d].SetNumericSeparatorsDto(expectedNumSeps)\n\n"+
				"expectedNumSeps= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, k, expectedNumSeps.String(), err.Error())
			return
		}

		err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = result[%d].IsValid('Validating result[%d]')\n"+
				"Validation Error= '%v'\n\n",
				ePrefix, k, k, err.Error())
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

func TestBigIntMathSubtract_SubtractINumMgrSeries_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrSeries_01"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(\n"+
			" &nDtoMinuend, &dec0, &bINum1, &ia2,\n"+
			"  &nDto3, &dec4, &ia5)\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahend0, subtrahend1, subtrahend2,
			subtrahend3, subtrahend4, subtrahend5,
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

func TestBigIntMathSubtract_SubtractINumMgrSeries_02(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrSeries_02"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(\n"+
			" &nDtoMinuend, &dec0, &bINum1, &ia2,\n"+
			"  &nDto3, &dec4, &ia5)\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahend0, subtrahend1, subtrahend2,
			subtrahend3, subtrahend4, subtrahend5,
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

func TestBigIntMathSubtract_SubtractINumMgrSeries_03(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrSeries_03"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(\n"+
			" &nDtoMinuend, &dec0, &bINum1, &ia2,\n"+
			"  &nDto3, &dec4, &ia5)\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			" subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahend0, subtrahend1, subtrahend2,
			subtrahend3, subtrahend4, subtrahend5,
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

func TestBigIntMathSubtract_SubtractINumMgrSeries_04(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrSeries_04"

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
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
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

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(\n"+
			" &nDtoMinuend, &dec0, &bINum1, &ia2,\n"+
			"  &nDto3, &dec4, &ia5)\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			"subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahend0, subtrahend1, subtrahend2,
			subtrahend3, subtrahend4, subtrahend5,
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

func TestBigIntMathSubtract_SubtractINumMgrSeries_05(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractINumMgrSeries_05"

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
			ePrefix, minuendStr, usaNumSeps.String(), err.Error())
		return
	}

	err = nDtoMinuend.IsValid("Validating nDtoMinuend")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDtoMinuend.IsValid(Validating nDtoMinuend)\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"nDtoMinuendNumStr, err := nDtoMinuend.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

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

	dec0, err := new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec0, err := new(Decimal).NewNumStr(subtrahend0)\n"+
			"subtrahend0= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend0, err.Error())
		return
	}

	bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bINum1, err := new(BigIntNum).NewNumStr(subtrahend1)\n"+
			"subtrahend1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend1, err.Error())
		return
	}

	ia2, err := new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia2, err := new(IntAry).NewNumStr(subtrahend2)\n"+
			"subtrahend2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend2, err.Error())
		return
	}

	nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" nDto3, err := new(NumStrDto).NewNumStr(subtrahend3)\n"+
			"subtrahend3= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend3, err.Error())
		return
	}

	dec4, err := new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" dec4, err := new(Decimal).NewNumStr(subtrahend4)\n"+
			"subtrahend4= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend4, err.Error())
		return
	}

	ia5, err := new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" ia5, err := new(IntAry).NewNumStr(subtrahend5)\n"+
			"subtrahend5= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, subtrahend5, err.Error())
		return
	}

	result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(
		&nDtoMinuend,
		&dec0,
		&bINum1,
		&ia2,
		&nDto3,
		&dec4,
		&ia5)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractINumMgrSeries(\n"+
			" &nDtoMinuend, &dec0, &bINum1, &ia2,\n"+
			"  &nDto3, &dec4, &ia5)\n"+
			"iaMinuend= '%v'\n"+
			"subtrahendAry[0]= '%v'\n"+
			"subtrahendAry[1]= '%v'\n"+
			"subtrahendAry[2]= '%v'\n"+
			"subtrahendAry[3]= '%v'\n"+
			"subtrahendAry[4]= '%v'\n"+
			"subtrahendAry[5]= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nDtoMinuendNumStr, subtrahend0, subtrahend1, subtrahend2,
			subtrahend3, subtrahend4, subtrahend5,
			err.Error())

		return
	}

	err = result.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.SetNumericSeparatorsDto(\n"+
			"  expectedNumSeps)\n"+
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
