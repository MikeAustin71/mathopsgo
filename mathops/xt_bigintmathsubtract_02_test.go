package mathops

import (
	"fmt"
	"testing"
)

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_01(t *testing.T) {

	ePrefix := "TestBigIntMathSubtract_SubtractDecimalOutputToArray_01"

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
			"err = decMinuend.IsValid('Validating decMinuend')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decMinuendNumStr, err := decMinuend.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if minuendStr != decMinuendNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because minuendStr != decMinuendNumStr \n"+
			"Expected decMinuendNumStr = '%v'\n"+
			"  Actual decMinuendNumStr = '%v'\n\n",
			ePrefix, minuendStr, decMinuendNumStr)

		return
	}

	lenSubtrahends := len(subtrahendStrs)

	subtrahendAry := make([]Decimal, lenSubtrahends)

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

	expectedResultsAry := make([]Decimal, lenSubtrahends)

	var expectedResultsNumStr string

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"subtrahendAry[%d], err = new(Decimal).\n"+
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
				fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
				"Error='%v'\n\n", ePrefix, err.Error())
			return
		}

		expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultsAry[%d], err := new(Decimal).\n"+
				"  NewNumStr(expectedStrs[%d])\n"+
				"expectedStrs[%d]= '%v'\n"+
				"Error='%v'\n\n",
				ePrefix, i, i, i, expectedStrs[i], err.Error())

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
		new(BigIntMathSubtract).SubtractDecimalOutputToArray(
			decMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathSubtract).SubtractDecimalOutputToArray(\n"+
			"  decMinuend, subtrahendAry[...])\n"+
			"decMinuend= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decMinuend,
			err.Error())

		return
	}

	var expectedResultEqualsResult bool

	var expectedResultNumStr, resultNumStr string

	var resultNumSeps NumericSeparatorDto

	for k := 0; k < lenSubtrahends; k++ {

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

		expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"expectedResultEqualsResult, err =\n"+
				"  expectedResultsAry[%d].Equal(result[%d])\n"+
				"expectedResultsAry[%d]= '%v'\n"+
				"result[%d]= '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
			return
		}

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

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_02(t *testing.T) {

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

	minuendDecimal, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractDecimalOutputToArray"+
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_03(t *testing.T) {

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

	minuendDecimal, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractDecimalOutputToArray"+
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}
}

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_04(t *testing.T) {

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

	minuendDecimal, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractDecimalOutputToArray"+
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_05(t *testing.T) {

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

	minuendDecimal, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractDecimalOutputToArray"+
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}
}

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_06(t *testing.T) {

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

	minuendDecimal, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = minuendDecimal.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by minuendDecimal.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(Decimal).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractDecimalOutputToArray"+
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		actualNumStr := resultArray[k].GetNumStr()

		if expectedStrs[k] != actualNumStr {
			t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. Index='%v'",
				expectedStrs[k], actualNumStr, k)
		}

		actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. Index='%v'",
				expectedNumSeps.String(), actualNumSeps.String(), k)
		}

	}
}

func TestBigIntMathSubtract_SubtractDecimalSeries_01(t *testing.T) {

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

	decMinuend, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
		decMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalSeries("+
			"decMinuend, ...). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimalSeries_02(t *testing.T) {

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

	decMinuend, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
		decMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalSeries("+
			"decMinuend, ...). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimalSeries_03(t *testing.T) {

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

	decMinuend, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
		decMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalSeries("+
			"decMinuend, ...). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimalSeries_04(t *testing.T) {

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

	decMinuend, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
		decMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalSeries("+
			"decMinuend, ...). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractDecimalSeries_05(t *testing.T) {

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
	expectedNumStr := "6805998,231036"

	decMinuend, err := new(Decimal).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(Decimal).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = decMinuend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by decMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(Decimal).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
		decMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5])

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalSeries("+
			"decMinuend, ...). Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'. ",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathSubtract_SubtractIntAry_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"
	expectedBigINumSign := 1

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAry(iaMinuend, iaSubtrahend)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractIntAry(iaMinuend, "+
			"iaSubtrahend) iaMinuend='%v' subtrahendStr='%v' Error='%v' ",
			iaMinuend.GetNumStr(), iaSubtrahend.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractIntAry_02(t *testing.T) {
	// minuend = 949321.6712
	minuendStr := "949321.6712"

	// subtrahend = 45678.21
	subtrahendStr := "45678.21"

	// result = 903643.4612
	expectedBigINumStr := "903643.4612"
	expectedBigINumSign := 1

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAry(iaMinuend, iaSubtrahend)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractIntAry(iaMinuend, "+
			"iaSubtrahend) iaMinuend='%v' subtrahendStr='%v' Error='%v' ",
			iaMinuend.GetNumStr(), iaSubtrahend.GetNumStr(), err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractIntAry_03(t *testing.T) {
	// minuend = -5876458.56789012
	minuendStr := "-5876458.56789012"

	// subtrahend = 847129.876
	subtrahendStr := "847129.876"

	// result = -6723588.44389012
	expectedBigINumStr := "-6723588.44389012"
	expectedBigINumSign := -1

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAry(iaMinuend, iaSubtrahend)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractIntAry(iaMinuend, "+
			"iaSubtrahend) iaMinuend='%v' subtrahendStr='%v' Error='%v' ",
			iaMinuend.GetNumStr(), iaSubtrahend.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractIntAry_04(t *testing.T) {
	// minuend = -289.673849
	minuendStr := "-289.673849"

	// subtrahend = -14579.012
	subtrahendStr := "-14579.012"

	// result = 14289.338151
	expectedBigINumStr := "14289.338151"
	expectedBigINumSign := 1

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAry(iaMinuend, iaSubtrahend)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractIntAry(iaMinuend, "+
			"iaSubtrahend) iaMinuend='%v' subtrahendStr='%v' Error='%v' ",
			iaMinuend.GetNumStr(), iaSubtrahend.GetNumStr(), err.Error())
	}

	if !expectedBigINum.Equal(result) {
		t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractIntAry_05(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedNumStr := "99,999"

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStr) "+
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = iaMinuend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iaMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAry(iaMinuend, iaSubtrahend)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractIntAry(iaMinuend, "+
			"iaSubtrahend) iaMinuend='%v' subtrahendStr='%v' Error='%v' ",
			iaMinuend.GetNumStr(), iaSubtrahend.GetNumStr(), err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntMathSubtract_SubtractIntAryArray_01(t *testing.T) {

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

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalArray("+
			"iaMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractIntAryArray_02(t *testing.T) {

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

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractIntAryArray("+
			"iaMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractIntAryArray_03(t *testing.T) {

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

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractIntAryArray("+
			"iaMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractIntAryArray_04(t *testing.T) {

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

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractIntAryArray("+
			"iaMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
		t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
			expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
	}

	if expectedBigINumSign != result.sign {
		t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
			expectedBigINumSign, result.sign)
	}

}

func TestBigIntMathSubtract_SubtractIntAryArray_05(t *testing.T) {

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
	expectedNumStr := "6805998,231036"

	iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = iaMinuend.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by iaMinuend.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = new(IntAry).NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend0). "+
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = new(IntAry).NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend1). "+
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = new(IntAry).NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend2). "+
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = new(IntAry).NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend3). "+
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = new(IntAry).NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend4). "+
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = new(IntAry).NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from new(IntAry).NewNumStr(subtrahend5). "+
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := new(BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathSubtract).SubtractDecimalArray("+
			"iaMinuend, subtrahendAry). Error='%v' ", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Comparison Error: Expected NumStr='%v'. Instead, NumStr= '%v'. ",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := result.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_01(t *testing.T) {

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

	minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractIntAryOutputToArray"+
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equals(&expectedResultsAry[k]) {
			t.Errorf("Inequality Error: Expected ResultsAry='%v'. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_02(t *testing.T) {

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

	minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractIntAryOutputToArray"+
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equals(&expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_03(t *testing.T) {

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

	minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractIntAryOutputToArray"+
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equals(&expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}
}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_04(t *testing.T) {

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

	minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractIntAryOutputToArray"+
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equals(&expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_05(t *testing.T) {

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

	minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractIntAryOutputToArray"+
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equals(&expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}
}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_06(t *testing.T) {

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

	minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by new(IntAry).NewNumStr(minuendStr) "+
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	err = minuendIntAry.SetNumericSeparatorsDto(expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by minuendIntAry.SetNumericSeparatorsDto(expectedNumSeps). "+
			"Error='%v' ", err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i := 0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(subtrahendStrs[i]) "+
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by new(IntAry).NewNumStr(expectedStrs[i]) "+
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned bynew(BigIntMathSubtract).SubtractIntAryOutputToArray"+
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k := 0; k < lenSubtrahends; k++ {

		actualNumStr := resultArray[k].GetNumStr()

		if expectedStrs[k] != actualNumStr {
			t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
				expectedStrs[k], actualNumStr)
		}

		actualNumSeps := resultArray[k].GetNumericSeparatorsDto()

		if !expectedNumSeps.Equal(actualNumSeps) {
			t.Errorf("Error: Expected numSeps='%v'. Instead, NumStr='%v'.",
				expectedNumSeps.String(), actualNumSeps.String())
		}
	}
}
