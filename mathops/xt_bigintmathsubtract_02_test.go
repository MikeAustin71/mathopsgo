package mathops

import "testing"

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_01(t *testing.T) {

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


	minuendDecimal, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = Decimal{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = Decimal{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractDecimalOutputToArray" +
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Inequality Error: Expected ResultsAry='%v'. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}

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

	minuendDecimal, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = Decimal{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = Decimal{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractDecimalOutputToArray" +
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

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

	minuendDecimal, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = Decimal{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = Decimal{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractDecimalOutputToArray" +
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

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

	minuendDecimal, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = Decimal{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = Decimal{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractDecimalOutputToArray" +
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

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

	minuendDecimal, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]Decimal, lenSubtrahends)
	expectedResultsAry := make([]Decimal, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = Decimal{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = Decimal{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractDecimalOutputToArray(minuendDecimal, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractDecimalOutputToArray" +
			"(minuendDecimal, subtrahendAry) minuendDecimal='%v'  Error='%v'. ",
			minuendDecimal.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}
}

func TestBigIntMathSubtract_SubtractDecimalSeries_01(t *testing.T) {

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0:= "123.894000"
	subtrahend1:= "67.1"
	subtrahend2:= "93.0"
	subtrahend3:= "-124498.67158"
	subtrahend4:= "647129.57"
	subtrahend5:= "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"
	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimalSeries(
		decMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5] )

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimalSeries(" +
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

	subtrahend0:= "737.21"
	subtrahend1:= "9637591.879546"
	subtrahend2:= "28"
	subtrahend3:= "5284.9765"
	subtrahend4:= "-189291837.12"
	subtrahend5:= "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"
	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimalSeries(
		decMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5] )

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimalSeries(" +
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

	subtrahend0:= "-28934682.721"
	subtrahend1:= "424.987654321"
	subtrahend2:= "-987"
	subtrahend3:= "62.94"
	subtrahend4:= "-999999999.99999"
	subtrahend5:= "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"
	expectedBigINumSign := 1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimalSeries(
		decMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5] )

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimalSeries(" +
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


	subtrahend0:= "-28934682.721"
	subtrahend1:= "424.987654321"
	subtrahend2:= "-987"
	subtrahend3:= "62.94"
	subtrahend4:= "-999999999.99999"
	subtrahend5:= "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"
	expectedBigINumSign := -1

	decMinuend, err := Decimal{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]Decimal, lenSubtrahends)

	subtrahendAry[0], err = Decimal{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = Decimal{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = Decimal{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = Decimal{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = Decimal{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = Decimal{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from Decimal{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractDecimalSeries(
		decMinuend,
		subtrahendAry[0],
		subtrahendAry[1],
		subtrahendAry[2],
		subtrahendAry[3],
		subtrahendAry[4],
		subtrahendAry[5] )

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimalSeries(" +
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

func TestBigIntMathSubtract_SubtractIntAry_01(t *testing.T) {
	// minuend = 123.32
	minuendStr := "123.32"

	// subtrahend = 23.321
	subtrahendStr := "23.321"

	// result = 99.999
	expectedBigINumStr := "99.999"
	expectedBigINumSign := 1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := IntAry{}.NewNumStr(subtrahendStr)

	if err != nil  {
		t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntAry(iaMinuend, iaSubtrahend)

	if err != nil  {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntAry(iaMinuend, "+
			"iaSubtrahend) iaMinuend='%v' subtrahendStr='%v' Error='%v' ",
			iaMinuend.GetNumStr() , iaSubtrahend.GetNumStr(), err.Error())
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

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := IntAry{}.NewNumStr(subtrahendStr)

	if err != nil  {
		t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntAry(iaMinuend, iaSubtrahend)

	if err != nil  {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntAry(iaMinuend, "+
			"iaSubtrahend) iaMinuend='%v' subtrahendStr='%v' Error='%v' ",
			iaMinuend.GetNumStr() , iaSubtrahend.GetNumStr(), err.Error())
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

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := IntAry{}.NewNumStr(subtrahendStr)

	if err != nil  {
		t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntAry(iaMinuend, iaSubtrahend)

	if err != nil  {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntAry(iaMinuend, "+
			"iaSubtrahend) iaMinuend='%v' subtrahendStr='%v' Error='%v' ",
			iaMinuend.GetNumStr() , iaSubtrahend.GetNumStr(), err.Error())
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

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v'", minuendStr, err.Error())
	}

	iaSubtrahend, err := IntAry{}.NewNumStr(subtrahendStr)

	if err != nil  {
		t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStr) " +
			"subtrahendStr='%v' Error='%v' ", subtrahendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntAry(iaMinuend, iaSubtrahend)

	if err != nil  {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntAry(iaMinuend, "+
			"iaSubtrahend) iaMinuend='%v' subtrahendStr='%v' Error='%v' ",
			iaMinuend.GetNumStr() , iaSubtrahend.GetNumStr(), err.Error())
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

func TestBigIntMathSubtract_SubtractIntAryArray_01(t *testing.T) {

	var err error

	// minuend = 7328941.123456
	minuendStr := "7328941.123456"

	subtrahend0:= "123.894000"
	subtrahend1:= "67.1"
	subtrahend2:= "93.0"
	subtrahend3:= "-124498.67158"
	subtrahend4:= "647129.57"
	subtrahend5:= "28"

	// result = 6805998.231036
	expectedBigINumStr := "6805998.231036"
	expectedBigINumSign := 1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = IntAry{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = IntAry{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = IntAry{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = IntAry{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntAryArray(iaMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractDecimalArray(" +
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

	subtrahend0:= "737.21"
	subtrahend1:= "9637591.879546"
	subtrahend2:= "28"
	subtrahend3:= "5284.9765"
	subtrahend4:= "-189291837.12"
	subtrahend5:= "7638932.12398765"

	// result = 153,035,620.80650965
	expectedBigINumStr := "153035620.80650965"

	expectedBigINumSign := 1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = IntAry{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = IntAry{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = IntAry{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = IntAry{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntAryArray(iaMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntAryArray(" +
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

	subtrahend0:= "-28934682.721"
	subtrahend1:= "424.987654321"
	subtrahend2:= "-987"
	subtrahend3:= "62.94"
	subtrahend4:= "-999999999.99999"
	subtrahend5:= "-9638932.371"

	// Result:  2,757,547,756.287792379
	expectedBigINumStr := "2757547756.287792379"
	expectedBigINumSign := 1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = IntAry{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = IntAry{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = IntAry{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = IntAry{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntAryArray(iaMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntAryArray(" +
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

	subtrahend0:= "-28934682.721"
	subtrahend1:= "424.987654321"
	subtrahend2:= "-987"
	subtrahend3:= "62.94"
	subtrahend4:= "-999999999.99999"
	subtrahend5:= "-9638932.371"

	// Result:   -680,399,527.959121021
	expectedBigINumStr := "-680399527.959121021"
	expectedBigINumSign := -1

	iaMinuend, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v' Error='%v' ", minuendStr, err.Error())
	}

	expectedBigINum, err := BigIntNum{}.NewNumStr(expectedBigINumStr)

	if err!=nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) " +
			"expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
	}

	lenSubtrahends := 6
	subtrahendAry := make([]IntAry, lenSubtrahends)

	subtrahendAry[0], err = IntAry{}.NewNumStr(subtrahend0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend0). " +
			"subtrahend0='%v' Error='%v'. ",
			subtrahend0, err.Error())
	}

	subtrahendAry[1], err = IntAry{}.NewNumStr(subtrahend1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend1). " +
			"subtrahend1='%v' Error='%v'. ",
			subtrahend1, err.Error())
	}

	subtrahendAry[2], err = IntAry{}.NewNumStr(subtrahend2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend2). " +
			"subtrahend2='%v' Error='%v'. ",
			subtrahend2, err.Error())
	}

	subtrahendAry[3], err = IntAry{}.NewNumStr(subtrahend3)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend3). " +
			"subtrahend3='%v' Error='%v'. ",
			subtrahend3, err.Error())
	}

	subtrahendAry[4], err = IntAry{}.NewNumStr(subtrahend4)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend4). " +
			"subtrahend4='%v' Error='%v'. ",
			subtrahend4, err.Error())
	}

	subtrahendAry[5], err = IntAry{}.NewNumStr(subtrahend5)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(subtrahend5). " +
			"subtrahend5='%v' Error='%v'. ",
			subtrahend5, err.Error())
	}

	result, err := BigIntMathSubtract{}.SubtractIntAryArray(iaMinuend, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned by BigIntMathSubtract{}.SubtractIntAryArray(" +
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


	minuendIntAry, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = IntAry{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = IntAry{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractIntAryOutputToArray" +
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
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

	minuendIntAry, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = IntAry{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = IntAry{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractIntAryOutputToArray" +
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
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

	minuendIntAry, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = IntAry{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = IntAry{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractIntAryOutputToArray" +
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
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

	minuendIntAry, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = IntAry{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = IntAry{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractIntAryOutputToArray" +
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
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

	minuendIntAry, err := IntAry{}.NewNumStr(minuendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(minuendStr) " +
			"minuendStr='%v'  Error='%v'. ", minuendStr, err.Error())
	}

	lenSubtrahends := len(subtrahendStrs)
	subtrahendAry := make([]IntAry, lenSubtrahends)
	expectedResultsAry := make([]IntAry, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		subtrahendAry[i], err = IntAry{}.NewNumStr(subtrahendStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(subtrahendStrs[i]) " +
				"subtrahendStrs[%v]='%v'  Error='%v'. ", i, subtrahendStrs[i], err.Error())
		}

		expectedResultsAry[i], err = IntAry{}.NewNumStr(expectedStrs[i])

		if err != nil {
			t.Errorf("Error returned by IntAry{}.NewNumStr(expectedStrs[i]) " +
				"expectedStrs[%v]='%v'  Error='%v'. ", i, expectedStrs[i], err.Error())
		}

	}

	resultArray, err :=
		BigIntMathSubtract{}.SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

	if err != nil {
		t.Errorf("Error returned byBigIntMathSubtract{}.SubtractIntAryOutputToArray" +
			"(minuendIntAry, subtrahendAry) minuendIntAry='%v'  Error='%v'. ",
			minuendIntAry.GetNumStr(), err.Error())
	}

	for k:=0; k < lenSubtrahends; k++ {

		if !resultArray[k].Equal(expectedResultsAry[k]) {
			t.Errorf("Error: Expected ResultsAry='%v' Not Equal. Instead, ResultsAry='%v'. ",
				expectedResultsAry[k].GetNumStr(), resultArray[k].GetNumStr())
		}

	}
}

